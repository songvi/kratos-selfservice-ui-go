package driver

import (
	"github.com/songvi/kratos-selfservice-ui-go/handler"
	"github.com/songvi/kratos-selfservice-ui-go/driver/configuration"
	"github.com/sirupsen/logrus"
)

type Registry interface {
	handler.HandlerProvider
	WithConfig(cfg configuration.ConfigProvider) Registry
	WithLogger(l logrus.FieldLogger) Registry
}