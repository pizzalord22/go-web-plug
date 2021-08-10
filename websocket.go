package websocket

import (
    "crypto/tls"
    "crypto/x509"
    "errors"
    "log"
    "net/url"
    "sync"
    "time"

    "github.com/gorilla/websocket"
)

var syncLock = new(sync.Mutex)

// var reconnectLock = new(sync.Mutex)

type Ws struct {
    // websocket connection
    conn *websocket.Conn

    // certificate pool used for secure connections
    caPool *x509.CertPool

    // set to true to use certificates
    secure bool

    // url contains the url to connect to
    url url.URL

    // set to true to send the initMsg when a connection is made
    sendInitMsg bool

    // message that is to be sent when a connection is made
    initMsg []byte

    // set to true to automatically try to reconnect
    reconnect    bool
    reconnecting bool

    // close handler is called when a connection ends
    closeHandler func(int, string) error
}

// create a nwe caPool, this is needed since we can not add new certs to an empty cert pool
func init() {
    Websocket.caPool = x509.NewCertPool()
}

// semver 2.0
const version = "1.2.0"

// Version return the current version number
func (w *Ws) Version() string {
    return version
}

// Read a websocket message
func (w *Ws) Read() (int, []byte, error) {
    if w.conn == nil {
        err := w.Connect()
        if err != nil {
            return 0, []byte{}, errors.New("can not read when there is no connection, and could not create a connection")
        }
    }
    t, d, err := w.conn.ReadMessage()
    go w.errCheck(err)
    return t, d, err
}

// ReadJSON read a websocket message in json format
func (w *Ws) ReadJSON(v interface{}) error {
    if w.conn == nil {
        err := w.Connect()
        if err != nil {
            return errors.New("can not read when there is no connection, and could not create a connection")
        }
    }
    err := w.conn.ReadJSON(v)
    go w.errCheck(err)
    return err
}

// WriteMessage write a message
func (w *Ws) WriteMessage(messageType int, data []byte) error {
    if w.conn == nil {
        err := w.Connect()
        if err != nil {
            return errors.New("can not read when there is no connection, and could not create a connection")
        }
    }
    err := w.conn.WriteMessage(messageType, data)
    go w.errCheck(err)
    return err
}

// WriteJSON write a message in json format
func (w *Ws) WriteJSON(v interface{}) error {
    if w.conn == nil {
        err := w.Connect()
        if err != nil {
            return errors.New("can not read when there is no connection, and could not create a connection")
        }
    }
    err := w.conn.WriteJSON(v)
    go w.errCheck(err)
    return err
}

// AppendCertsFromPem add a certificate to the certificate pool
func (w *Ws) AppendCertsFromPem(pemCerts []byte) bool {
    return w.caPool.AppendCertsFromPEM(pemCerts)
}

// SetUrl set the url to connect to
func (w *Ws) SetUrl(scheme, host, path string) {
    w.url = url.URL{Scheme: scheme, Host: host, Path: path}
}

// Connect to the websocket server
func (w *Ws) Connect() error {
    syncLock.Lock()
    defer syncLock.Unlock()
    log.Println("locked connect mutex")
    var d websocket.Dialer
    if w.secure {
        config := tls.Config{RootCAs: w.caPool}
        d = websocket.Dialer{TLSClientConfig: &config}
    }
    log.Println("attempting to make connection")
    c, _, err := d.Dial(w.url.String(), nil)
    if err != nil {
        return err
    }
    if w.conn != nil {
        err = w.Close()
        if err != nil {
            log.Println(err)
        }
    }
    log.Println("made a connection")
    w.conn = c
    w.conn.SetCloseHandler(w.closeHandler)
    if w.sendInitMsg {
        return w.WriteMessage(1, w.initMsg)
    }
    log.Println("send innit message exiting connect")
    return nil
}

// SetInitMsg set a message to be sent when a connection is established
func (w *Ws) SetInitMsg(msg []byte) {
    w.sendInitMsg = true
    w.initMsg = msg
}

// SetCloseHandler set a close handler to call when a connection ends
func (w *Ws) SetCloseHandler(f func(int, string) error) {
    w.closeHandler = f
}

// Reconnect set to true for automatic reconnecting
func (w *Ws) Reconnect(b bool) {
    w.reconnect = b
}

// Close the websocket connection
func (w *Ws) Close() error {
    if w.conn == nil {
        return nil
    }
    // w.WriteMessage(websocket.CloseMessage, []byte{})
    return w.conn.Close()
}

// check for network problems
func (w *Ws) errCheck(err error) {
    if err != nil {
        log.Println(err)
    }
    if w.reconnecting {
        return
    }
    if w.reconnect && err != nil {
        w.reconnecting = true
        for err != nil {
            err = w.Connect()
            time.Sleep(1 * time.Second)
        }
        w.reconnecting = false
        return
    }
}

// SetSecure set the secure bit
func (w *Ws) SetSecure(b bool) {
    w.secure = b
}

// WriteQueue requires  a channel te read message from and a channel to send errors to
// if wil requeue failed messages until the queue is filled, then it will throw them away
func (w *Ws) WriteQueue(c chan []byte, e chan error) {
    go func() {
        for bytes := range c {
            err := w.WriteMessage(1, bytes)
            go w.errCheck(err)
            if err != nil {
                e <- err
                // when the buffer is full we remove an old element and insert a new one,
                // so that we always have the last ~5 minutes in the buffer,
                // this only works if there is one routine reading from the channel
                if len(c) == cap(c) {
                    <-c
                }
                c <- bytes
                time.Sleep(1 * time.Second)
            }
        }
    }()
}

func Exit(){
    // todo cleanly exit the plug
}


// Websocket exported as symbol named "Websocket"
var Websocket Ws
