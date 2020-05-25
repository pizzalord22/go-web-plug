package websocket

import (
	"crypto/x509"
	"github.com/gorilla/websocket"
	"net/url"
	"reflect"
	"testing"
)

func TestWs_AppendCertsFromPem(t *testing.T) {
	type fields struct {
		conn         *websocket.Conn
		caPool       *x509.CertPool
		secure       bool
		url          url.URL
		sendInitMsg  bool
		initMsg      []byte
		reconnect    bool
		closeHandler func(int, string) error
	}
	type args struct {
		pemCerts []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "append cert test 1", fields: fields{}, args: args{}, want: false},
		{name: "append cert test 2", fields: fields{caPool: new(x509.CertPool)}, args: args{}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Ws{
				conn:         tt.fields.conn,
				caPool:       tt.fields.caPool,
				secure:       tt.fields.secure,
				url:          tt.fields.url,
				sendInitMsg:  tt.fields.sendInitMsg,
				initMsg:      tt.fields.initMsg,
				reconnect:    tt.fields.reconnect,
				closeHandler: tt.fields.closeHandler,
			}
			if got := w.AppendCertsFromPem(tt.args.pemCerts); got != tt.want {
				t.Errorf("AppendCertsFromPem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWs_Close(t *testing.T) {
	type fields struct {
		conn         *websocket.Conn
		caPool       *x509.CertPool
		secure       bool
		url          url.URL
		sendInitMsg  bool
		initMsg      []byte
		reconnect    bool
		closeHandler func(int, string) error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "close test", fields: fields{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Ws{
				conn:         tt.fields.conn,
				caPool:       tt.fields.caPool,
				secure:       tt.fields.secure,
				url:          tt.fields.url,
				sendInitMsg:  tt.fields.sendInitMsg,
				initMsg:      tt.fields.initMsg,
				reconnect:    tt.fields.reconnect,
				closeHandler: tt.fields.closeHandler,
			}
			if err := w.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWs_Connect(t *testing.T) {
	type fields struct {
		conn         *websocket.Conn
		caPool       *x509.CertPool
		secure       bool
		url          url.URL
		sendInitMsg  bool
		initMsg      []byte
		reconnect    bool
		closeHandler func(int, string) error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Ws{
				conn:         tt.fields.conn,
				caPool:       tt.fields.caPool,
				secure:       tt.fields.secure,
				url:          tt.fields.url,
				sendInitMsg:  tt.fields.sendInitMsg,
				initMsg:      tt.fields.initMsg,
				reconnect:    tt.fields.reconnect,
				closeHandler: tt.fields.closeHandler,
			}
			if err := w.Connect(); (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWs_Read(t *testing.T) {
	type fields struct {
		conn         *websocket.Conn
		caPool       *x509.CertPool
		secure       bool
		url          url.URL
		sendInitMsg  bool
		initMsg      []byte
		reconnect    bool
		closeHandler func(int, string) error
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		want1   []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Ws{
				conn:         tt.fields.conn,
				caPool:       tt.fields.caPool,
				secure:       tt.fields.secure,
				url:          tt.fields.url,
				sendInitMsg:  tt.fields.sendInitMsg,
				initMsg:      tt.fields.initMsg,
				reconnect:    tt.fields.reconnect,
				closeHandler: tt.fields.closeHandler,
			}
			got, got1, err := w.Read()
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Read() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestWs_ReadJSON(t *testing.T) {
	type fields struct {
		conn         *websocket.Conn
		caPool       *x509.CertPool
		secure       bool
		url          url.URL
		sendInitMsg  bool
		initMsg      []byte
		reconnect    bool
		closeHandler func(int, string) error
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Ws{
				conn:         tt.fields.conn,
				caPool:       tt.fields.caPool,
				secure:       tt.fields.secure,
				url:          tt.fields.url,
				sendInitMsg:  tt.fields.sendInitMsg,
				initMsg:      tt.fields.initMsg,
				reconnect:    tt.fields.reconnect,
				closeHandler: tt.fields.closeHandler,
			}
			if err := w.ReadJSON(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("ReadJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWs_Reconnect(t *testing.T) {
	type fields struct {
		conn         *websocket.Conn
		caPool       *x509.CertPool
		secure       bool
		url          url.URL
		sendInitMsg  bool
		initMsg      []byte
		reconnect    bool
		closeHandler func(int, string) error
	}
	type args struct {
		b bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Ws{
				conn:         tt.fields.conn,
				caPool:       tt.fields.caPool,
				secure:       tt.fields.secure,
				url:          tt.fields.url,
				sendInitMsg:  tt.fields.sendInitMsg,
				initMsg:      tt.fields.initMsg,
				reconnect:    tt.fields.reconnect,
				closeHandler: tt.fields.closeHandler,
			}
			if got := w.reconnect; got != tt.want {
				t.Errorf("Reconnect() error = %v, wantErr %v", got, tt.want)
			}
		})
	}
}

func TestWs_SetCloseHandler(t *testing.T) {
	type fields struct {
		conn         *websocket.Conn
		caPool       *x509.CertPool
		secure       bool
		url          url.URL
		sendInitMsg  bool
		initMsg      []byte
		reconnect    bool
		closeHandler func(int, string) error
	}
	type args struct {
		f func(int, string) error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Ws{
				conn:         tt.fields.conn,
				caPool:       tt.fields.caPool,
				secure:       tt.fields.secure,
				url:          tt.fields.url,
				sendInitMsg:  tt.fields.sendInitMsg,
				initMsg:      tt.fields.initMsg,
				reconnect:    tt.fields.reconnect,
				closeHandler: tt.fields.closeHandler,
			}
			_ = w
		})
	}
}

func TestWs_SetUrl(t *testing.T) {
	type fields struct {
		conn         *websocket.Conn
		caPool       *x509.CertPool
		secure       bool
		url          url.URL
		sendInitMsg  bool
		initMsg      []byte
		reconnect    bool
		closeHandler func(int, string) error
	}
	type args struct {
		scheme string
		host   string
		path   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Ws{
				conn:         tt.fields.conn,
				caPool:       tt.fields.caPool,
				secure:       tt.fields.secure,
				url:          tt.fields.url,
				sendInitMsg:  tt.fields.sendInitMsg,
				initMsg:      tt.fields.initMsg,
				reconnect:    tt.fields.reconnect,
				closeHandler: tt.fields.closeHandler,
			}
			_ = w
		})
	}
}

func TestWs_Version(t *testing.T) {
	type fields struct {
		conn         *websocket.Conn
		caPool       *x509.CertPool
		secure       bool
		url          url.URL
		sendInitMsg  bool
		initMsg      []byte
		reconnect    bool
		closeHandler func(int, string) error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Ws{
				conn:         tt.fields.conn,
				caPool:       tt.fields.caPool,
				secure:       tt.fields.secure,
				url:          tt.fields.url,
				sendInitMsg:  tt.fields.sendInitMsg,
				initMsg:      tt.fields.initMsg,
				reconnect:    tt.fields.reconnect,
				closeHandler: tt.fields.closeHandler,
			}
			if got := w.Version(); got != tt.want {
				t.Errorf("Version() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWs_WriteJSON(t *testing.T) {
	type fields struct {
		conn         *websocket.Conn
		caPool       *x509.CertPool
		secure       bool
		url          url.URL
		sendInitMsg  bool
		initMsg      []byte
		reconnect    bool
		closeHandler func(int, string) error
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Ws{
				conn:         tt.fields.conn,
				caPool:       tt.fields.caPool,
				secure:       tt.fields.secure,
				url:          tt.fields.url,
				sendInitMsg:  tt.fields.sendInitMsg,
				initMsg:      tt.fields.initMsg,
				reconnect:    tt.fields.reconnect,
				closeHandler: tt.fields.closeHandler,
			}
			if err := w.WriteJSON(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("WriteJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWs_WriteMessage(t *testing.T) {
	type fields struct {
		conn         *websocket.Conn
		caPool       *x509.CertPool
		secure       bool
		url          url.URL
		sendInitMsg  bool
		initMsg      []byte
		reconnect    bool
		closeHandler func(int, string) error
	}
	type args struct {
		messageType int
		data        []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Ws{
				conn:         tt.fields.conn,
				caPool:       tt.fields.caPool,
				secure:       tt.fields.secure,
				url:          tt.fields.url,
				sendInitMsg:  tt.fields.sendInitMsg,
				initMsg:      tt.fields.initMsg,
				reconnect:    tt.fields.reconnect,
				closeHandler: tt.fields.closeHandler,
			}
			if err := w.WriteMessage(tt.args.messageType, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("WriteMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWs_WriteQueue(t *testing.T) {
	type fields struct {
		conn         *websocket.Conn
		caPool       *x509.CertPool
		secure       bool
		url          url.URL
		sendInitMsg  bool
		initMsg      []byte
		reconnect    bool
		closeHandler func(int, string) error
	}
	type args struct {
		c chan []byte
		e chan error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Ws{
				conn:         tt.fields.conn,
				caPool:       tt.fields.caPool,
				secure:       tt.fields.secure,
				url:          tt.fields.url,
				sendInitMsg:  tt.fields.sendInitMsg,
				initMsg:      tt.fields.initMsg,
				reconnect:    tt.fields.reconnect,
				closeHandler: tt.fields.closeHandler,
			}
			_ = w
		})
	}
}

func TestWs_errCheck(t *testing.T) {
	type fields struct {
		conn         *websocket.Conn
		caPool       *x509.CertPool
		secure       bool
		url          url.URL
		sendInitMsg  bool
		initMsg      []byte
		reconnect    bool
		closeHandler func(int, string) error
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Ws{
				conn:         tt.fields.conn,
				caPool:       tt.fields.caPool,
				secure:       tt.fields.secure,
				url:          tt.fields.url,
				sendInitMsg:  tt.fields.sendInitMsg,
				initMsg:      tt.fields.initMsg,
				reconnect:    tt.fields.reconnect,
				closeHandler: tt.fields.closeHandler,
			}
			_ = w
		})
	}
}
