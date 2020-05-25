package websocket

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/url"
)

// todo make writing thread safe

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

// create a nwe certpool, this is needed since we can not add new certs to an empty cert pool
func init() {
	Websocket.caPool = x509.NewCertPool()
}

// semver 2.0
const version = "1.0.0"

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
	w.errCheck(err)
	log.Println(w.conn.UnderlyingConn().RemoteAddr().String())
	return t, d, err
}

// read a websocket message in json format
func (w *Ws) ReadJSON(v interface{}) error {
	if w.conn == nil {
		_ = w.Connect()
		return errors.New("can not read when there is no connection, trying to reconnect")
	}
	err := w.conn.ReadJSON(v)
	w.errCheck(err)
	return err
}

// write a message
func (w *Ws) WriteMessage(messageType int, data []byte) error {
	if w.conn == nil {
		_ = w.Connect()
		return errors.New("can not write when there is no connection, trying to reconnect")
	}
	err := w.conn.WriteMessage(messageType, data)
	w.errCheck(err)
	return err
}

// write a message in json format
func (w *Ws) WriteJSON(v interface{}) error {
	if w.conn == nil {
		_ = w.Connect()
		return errors.New("can not write when there is no connection, trying to reconnect")
	}
	err := w.conn.WriteJSON(v)
	w.errCheck(err)
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
	var d websocket.Dialer
	if w.secure {
		config := tls.Config{RootCAs: w.caPool}
		d = websocket.Dialer{TLSClientConfig: &config}
	}
	c, _, err := d.Dial(w.url.String(), nil)
	if err != nil {
		return err
	}
	w.conn = c
	w.conn.SetCloseHandler(w.closeHandler)
	return nil
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
	return w.conn.Close()
}

// check for network problems
func (w *Ws) errCheck(err error) {
	var reset bool
	if w.reconnect {
		if err != nil && websocket.IsCloseError(err) {
			reset = true
		}
		_, ok := err.(*net.OpError)
		if ok {
			reset = true
		}
		if reset {
			_ = w.Close()
			_ = w.Connect()
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
			w.errCheck(err)
			if err != nil {
				select {
				case e <- err:
				default:
				}
				select {
				case c <- bytes:
				default:
				}
			}
		}
	}()
}

// exported as symbol named "Websocket"
var Websocket Ws
