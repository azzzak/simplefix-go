package tests

import (
	"context"
	"fmt"
	simplefixgo "github.com/b2broker/simplefix-go"
	"github.com/b2broker/simplefix-go/session"
	fixgen "github.com/b2broker/simplefix-go/tests/fix44"
	"net"
	"testing"
	"time"
)

func RunAcceptor(port int, t *testing.T, storage session.MessageStorage) *simplefixgo.Acceptor {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		t.Fatalf("listen error: %s", err)
	}

	handlerFactory := simplefixgo.NewAcceptorHandlerFactory(fixgen.FieldMsgType, 10)

	server := simplefixgo.NewAcceptor(listener, handlerFactory, func(handler simplefixgo.AcceptorHandler) {
		s, err := session.NewAcceptorSession(
			context.Background(),
			&pseudoGeneratedOpts,
			handler,
			&session.LogonSettings{
				LogonTimeout: time.Second * 30,
				HeartBtLimits: &session.IntLimits{
					Min: 5,
					Max: 60,
				},
			},
			func(request *session.LogonSettings) (err error) {
				t.Logf(
					"user '%s' connected with password '%s'",
					request.Username,
					request.Password,
				)

				return nil
			},
		)
		if err != nil {
			panic(err)
		}

		err = s.Run()
		if err != nil {
			t.Fatalf("run s: %s", err)
		}

		s.SetMessageStorage(storage)
	})

	return server
}
