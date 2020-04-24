package driver

import (
	"github.com/songvi/kratos-selfservice-ui-go/driver/configuration"
)

type DefaultDriver struct {
	c configuration.ConfigProvider
	r Registry
}

func (d *DefaultDriver) Registry() Registry {
	return d.r
}

func (d *DefaultDriver) Configuration() configuration.ConfigProvider {
	return d.c
}

func NewDefaultDriver(cfg configuration.ConfigProvider, reg Registry) Driver {
	return &DefaultDriver{c: cfg, r: reg}
}
