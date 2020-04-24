package driver

import (
	"github.com/songvi/kratos-selfservice-ui-go/driver/configuration"
)
type Driver interface {
	Registry() Registry
	Configuration() configuration.ConfigProvider
}