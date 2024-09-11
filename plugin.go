package roadrunner_demo_plugin

import (
	"context"
	"github.com/roadrunner-server/errors"
	"go.uber.org/zap"
	"time"
)

const PluginName = "roadrunner_demo_plugin"

type Plugin struct {
	cfg    *Config
	logger *zap.Logger
	ticker *time.Ticker
}

type Configurer interface {
	// UnmarshalKey takes a single key and unmarshal it into a Struct.
	UnmarshalKey(name string, out any) error
	// Has checks if config section exists.
	Has(name string) bool
}

type Logger interface { // <-- logger plugin implements
	NamedLogger(name string) *zap.Logger
}

func (p *Plugin) Init(cfg Configurer, log Logger) error {
	p.logger = log.NamedLogger("roadrunner_demo_plugin")
	p.logger.Info("start initializing demo plugin")
	const op = errors.Op("custom_plugin_init")

	if !cfg.Has(PluginName) {
		return errors.E(op, errors.Disabled)
	}

	err := cfg.UnmarshalKey(PluginName, &p.cfg)
	if err != nil {
		// Error will stop execution
		return errors.E(op, err)
	}

	p.cfg.InitDefaults()

	p.logger.Info("end initializing demo plugin")
	return nil
}

func (p *Plugin) Serve() chan error {
	const op = errors.Op(PluginName)
	errCh := make(chan error, 1)

	p.ticker = time.NewTicker(10 * time.Second)

	err := p.DoSomeWork()
	if err != nil {
		errCh <- errors.E(op, err)
		return errCh
	}

	return nil
}

func (p *Plugin) Stop(_ context.Context) error {
	p.logger.Info("Stopping plugin")
	p.ticker.Stop()
	return nil
}

func (p *Plugin) DoSomeWork() error {
	p.logger.Info("DoSomeWork")
	for {
		select {
		case <-p.ticker.C:
			// Code to execute every 10 seconds
			p.logger.Info("DoSomeWork")
		}
	}
}
