package service

import (
	"GoSamples/ServiceSample/internal/pkg/config/types"
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type Service struct {
	// muxer    *Mux         // dispatch connections to different handlers listen on same port
	listener net.Listener // accept connections from client

	cfg *types.ServiceConfig

	ctx    context.Context    // service context
	cancel context.CancelFunc // call cancel to stop service
}

func (svr *Service) Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	svr.ctx = ctx
	svr.cancel = cancel

	// TODO: run web server and service
	svr.HandleListener(svr.listener, false)

	<-svr.ctx.Done()

	if svr.listener != nil {
		svr.Close()
	}
}

func (svr *Service) HandleListener(l net.Listener, internal bool) {
	for {
		time.Sleep(time.Second * 5)
		logrus.Infoln("service listener running")
	}
}

func NewService(cfg *types.ServiceConfig) (*Service, error) {
	svr := &Service{
		cfg: cfg,
		ctx: context.Background(),
	}

	address := net.JoinHostPort(cfg.TcpCfg.Host, strconv.Itoa(cfg.TcpCfg.Port))
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("create listener error, %v", err)
	}

	svr.listener = ln
	logrus.Infof("tcp listen on %s", address)
	return svr, nil
}

func (svr *Service) Close() error {

	if svr.listener != nil {
		svr.listener.Close()
	}

	if svr.cancel != nil {
		svr.cancel()
	}

	return nil
}
