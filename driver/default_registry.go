package driver

import (
	"github.com/sirupsen/logrus"
	"github.com/songvi/kratos-selfservice-ui-go/driver/configuration"
	"github.com/songvi/kratos-selfservice-ui-go/handler/dashboard"
	err "github.com/songvi/kratos-selfservice-ui-go/handler/error"
	"github.com/songvi/kratos-selfservice-ui-go/handler/hydra"
	"github.com/songvi/kratos-selfservice-ui-go/handler/login"
	"github.com/songvi/kratos-selfservice-ui-go/handler/logout"
	"github.com/songvi/kratos-selfservice-ui-go/handler/profile"
	"github.com/songvi/kratos-selfservice-ui-go/handler/registration"
	"github.com/songvi/kratos-selfservice-ui-go/handler/whoami"
)

type DefaultRegistry struct {
	l logrus.FieldLogger
	c configuration.ConfigProvider
}

func NewDefaultRegistry() *DefaultRegistry {
	return &DefaultRegistry{}
}

func (dr *DefaultRegistry) WithConfig(cfg configuration.ConfigProvider) Registry {
	dr.c = cfg
	return dr
}

func (dr *DefaultRegistry) WithLogger(logger logrus.FieldLogger) Registry {
	dr.l = logger
	return dr
}

func (dr *DefaultRegistry) LoginHandler() *login.Handler {
	return login.NewLoginHandler(dr.c, dr.l)
}

func (dr *DefaultRegistry) LogoutHandler() *logout.Handler {
	return logout.NewLogoutHandler(dr.c)
}

func (dr *DefaultRegistry) ProfileHandler() *profile.Handler {
	return profile.NewProfileHandler(dr.c)
}

func (dr *DefaultRegistry) RegistrationHandler() *registration.Handler {
	return registration.NewRegistrationHandler(dr.c, dr.l)
}

func (dr *DefaultRegistry) DashBoardHandler() *dashboard.Handler {
	return dashboard.NewDashBoardHandler(dr.c, dr.l)
}

func (dr *DefaultRegistry) ErrorHandler() *err.Handler {
	return err.NewErrorHandler(dr.c)
}

func (dr *DefaultRegistry) WhoamiHandler() *whoami.Handler {
	return whoami.NewWhoamiHandler(dr.c)
}

func (dr *DefaultRegistry) HydraLoginHandler() *hydra.LoginHandler {
	return hydra.NewHydraLoginHandler(dr.c)
}

func (dr *DefaultRegistry) HydraConsentHandler() *hydra.ConsentHandler {
	return hydra.NewHydraConsentHandler(dr.c)
}
