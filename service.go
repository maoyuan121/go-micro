package micro

import (
	"os"
	"os/signal"
	rtime "runtime"
	"strings"
	"sync"

	"go-micro.dev/v4/client"
	"go-micro.dev/v4/cmd"
	"go-micro.dev/v4/debug/handler"
	"go-micro.dev/v4/debug/stats"
	"go-micro.dev/v4/debug/trace"
	"go-micro.dev/v4/logger"
	plugin "go-micro.dev/v4/plugins"
	"go-micro.dev/v4/server"
	"go-micro.dev/v4/store"
	signalutil "go-micro.dev/v4/util/signal"
	"go-micro.dev/v4/util/wrapper"
)

type service struct {
	opts Options

	once sync.Once
}

func newService(opts ...Option) Service {
	service := new(service)
	options := newOptions(opts...)

	// 服务miservice name
	serviceName := options.Server.Options().Name

	// 包装 client，注入 From-Service header 到每次调用
	options.Client = wrapper.FromService(serviceName, options.Client)
	options.Client = wrapper.TraceCall(serviceName, trace.DefaultTracer, options.Client)

	// 包装 server，提供 handler 统计
	err := options.Server.Init(
		server.WrapHandler(wrapper.HandlerStats(stats.DefaultStats)),
		server.WrapHandler(wrapper.TraceHandler(trace.DefaultTracer)),
	)
	if err != nil {
		logger.Fatal(err)
	}

	// set opts
	service.opts = options

	return service
}

// 返回服务名
func (s *service) Name() string {
	return s.opts.Server.Options().Name
}

// Init initialises options. Additionally it calls cmd.Init
// which parses command line flags. cmd.Init is only called
// on first Init.
func (s *service) Init(opts ...Option) {
	// process options
	for _, o := range opts {
		o(&s.opts)
	}

	s.once.Do(func() {
		// 安装插件
		for _, p := range strings.Split(os.Getenv("MICRO_PLUGIN"), ",") {
			if len(p) == 0 {
				continue
			}

			// 加载插件
			c, err := plugin.Load(p)
			if err != nil {
				logger.Fatal(err)
			}

			// 初始化插件
			if err := plugin.Init(c); err != nil {
				logger.Fatal(err)
			}
		}

		// set cmd name
		if len(s.opts.Cmd.App().Name) == 0 {
			s.opts.Cmd.App().Name = s.Server().Options().Name
		}

		// Initialise the command flags, overriding new service
		if err := s.opts.Cmd.Init(
			cmd.Auth(&s.opts.Auth),
			cmd.Broker(&s.opts.Broker),
			cmd.Registry(&s.opts.Registry),
			cmd.Runtime(&s.opts.Runtime),
			cmd.Transport(&s.opts.Transport),
			cmd.Client(&s.opts.Client),
			cmd.Config(&s.opts.Config),
			cmd.Server(&s.opts.Server),
			cmd.Store(&s.opts.Store),
			cmd.Profile(&s.opts.Profile),
		); err != nil {
			logger.Fatal(err)
		}

		// 用服务名设置表名
		name := s.opts.Cmd.App().Name
		err := s.opts.Store.Init(store.Table(name))
		if err != nil {
			logger.Fatal(err)
		}
	})
}

func (s *service) Options() Options {
	return s.opts
}

func (s *service) Client() client.Client {
	return s.opts.Client
}

func (s *service) Server() server.Server {
	return s.opts.Server
}

func (s *service) String() string {
	return "micro"
}

func (s *service) Start() error {
	for _, fn := range s.opts.BeforeStart {
		if err := fn(); err != nil {
			return err
		}
	}

	if err := s.opts.Server.Start(); err != nil {
		return err
	}

	for _, fn := range s.opts.AfterStart {
		if err := fn(); err != nil {
			return err
		}
	}

	return nil
}

func (s *service) Stop() error {
	var err error

	for _, fn := range s.opts.BeforeStop {
		err = fn()
	}

	if err = s.opts.Server.Stop(); err != nil {
		return err
	}

	for _, fn := range s.opts.AfterStop {
		err = fn()
	}

	return err
}

func (s *service) Run() (err error) {
	// exit when help flag is provided
	for _, v := range os.Args[1:] {
		if v == "-h" || v == "--help" {
			os.Exit(0)
		}
	}

	// register the debug handler
	s.opts.Server.Handle(
		s.opts.Server.NewHandler(
			handler.NewHandler(s.opts.Client),
			server.InternalHandler(true),
		),
	)

	// start the profiler
	if s.opts.Profile != nil {
		// to view mutex contention
		rtime.SetMutexProfileFraction(5)
		// to view blocking profile
		rtime.SetBlockProfileRate(1)

		if err = s.opts.Profile.Start(); err != nil {
			return err
		}
		defer func() {
			err = s.opts.Profile.Stop()
			if err != nil {
				logger.Error(err)
			}
		}()
	}

	if logger.V(logger.InfoLevel, logger.DefaultLogger) {
		logger.Infof("Starting [service] %s", s.Name())
	}

	if err = s.Start(); err != nil {
		return err
	}

	ch := make(chan os.Signal, 1)
	if s.opts.Signal {
		signal.Notify(ch, signalutil.Shutdown()...)
	}

	select {
	// wait on kill signal
	case <-ch:
	// wait on context cancel
	case <-s.opts.Context.Done():
	}

	return s.Stop()
}
