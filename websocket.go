package websocket

import (
    "crypto/tls"
    "crypto/x509"
    "errors"
    "log"
    "net/url"
    `sync`
    `time`

    "github.com/gorilla/websocket"
)

var syncLock = new(sync.Mutex)

type Ws struct {
    // websocket connection
    conn *websocket.Conn

    // certificate pool used for secure connections
    caPool *x509.CertPool

    // set to true to use certificates
    secure bool

    // url contains the url to connect to
    url url.URL

    // set to true to send the initMsg when a connection is mademv6l.tar.g
    sendInitMsg bool

    // message that is to be send  when a connection is made
    initMsg []byte

    // set to true to automatically try to to reconnect
    reconnect bool

    // close handler is called when a connection ends
    closeHandler func(int, string) error
}

// create a nwe caPool, this is needed since we can not add new certs to an empty cert pool
func init() {
    Websocket.caPool = x509.NewCertPool()
}

// semver 2.0
const version = "1.1.0"

// return the current version number
func (w *Ws) Version() string {
    return version
}

// read a websocket message
func (w *Ws) Read() (int, []byte, error) {
    if w.conn == nil {
        _ = w.Connect()
        return 0, []byte{}, errors.New("can not read when there is no connection, trying to reconnect")
    }
    t, d, err := w.conn.ReadMessage()
    go w.errCheck(err)
    return t, d, err
}

// read a websocket message in json format
func (w *Ws) ReadJSON(v interface{}) error {
    if w.conn == nil {
        _ = w.Connect()
        return errors.New("can not read when there is no connection, trying to reconnect")
    }
    err := w.conn.ReadJSON(v)
    go w.errCheck(err)
    return err
}

// write a message
func (w *Ws) WriteMessage(messageType int, data []byte) error {
    if w.conn == nil {
        _ = w.Connect()
        return errors.New("can not write when there is no connection, trying to reconnect")
    }
    err := w.conn.WriteMessage(messageType, data)
    go w.errCheck(err)
    return err
}

// write a message in json format
func (w *Ws) WriteJSON(v interface{}) error {
    if w.conn == nil {
        _ = w.Connect()
        return errors.New("can not write when there is no connection, trying to reconnect")
    }
    err := w.conn.WriteJSON(v)
    go w.errCheck(err)
    return err
}

// add a certificate to the certificate pool
func (w *Ws) AppendCertsFromPem(pemCerts []byte) bool {
    return w.caPool.AppendCertsFromPEM(pemCerts)
}

// set the url to connect to
func (w *Ws) SetUrl(scheme, host, path string) {
    w.url = url.URL{Scheme: scheme, Host: host, Path: path}
}

// connect to the websocket server
func (w *Ws) Connect() error {
    syncLock.Lock()
    defer syncLock.Unlock()
    var d websocket.Dialer
    if w.secure {
        config := tls.Config{RootCAs: w.caPool}
        d = websocket.Dialer{TLSClientConfig: &config}
    }
    c, _, err := d.Dial(w.url.String(), nil)
    if err != nil {
        return err
    }
    if w.conn != nil {
        w.Close()
    }
    w.conn = c
    w.conn.SetCloseHandler(w.closeHandler)
    if w.sendInitMsg {
        return w.WriteMessage(1, w.initMsg)
    }
    return nil
}

// set a message to be send when a connection is established
func (w *Ws) SetInitMsg(msg []byte) {
    w.sendInitMsg = true
    w.initMsg = msg
}

// set a close handler to call when a connection ends
func (w *Ws) SetCloseHandler(f func(int, string) error) {
    w.closeHandler = f
}

// set to true for automatic reconnecting
func (w *Ws) Reconnect(b bool) {
    w.reconnect = b
}

// close the websocket connection
func (w *Ws) Close() error {
    if w.conn == nil {
        return nil
    }
    w.WriteMessage(websocket.CloseMessage, []byte{})
    return w.conn.Close()
}

// check for network problems
func (w *Ws) errCheck(err error) {
    if err != nil {
        log.Println(err)
        if w.reconnect {
            for err != nil {
                err = w.Connect()
                time.Sleep(1 * time.Second)
            }
        }
    }
}

// set the secure bit
func (w *Ws) SetSecure(b bool) {
    w.secure = b
}

// writeQueue requires  a channel te read message from and a channel to send errors to
// if wil requeue failed messages until the queue is filled, then it will throw them away
func (w *Ws) WriteQueue(c chan []byte, e chan error) {
    go func() {
        for bytes := range c {
            err := w.WriteMessage(1, bytes)
            go w.errCheck(err)
            if err != nil {
                e <- err
                c <- bytes
                time.Sleep(200 * time.Millisecond)
            }
        }
    }()
}

// exported as symbol named "Websocket"
var Websocket Ws
