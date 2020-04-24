package handler

import (
	err "github.com/songvi/kratos-selfservice-ui-go/handler/error"
	"github.com/songvi/kratos-selfservice-ui-go/handler/login"
	"github.com/songvi/kratos-selfservice-ui-go/handler/logout"
	"github.com/songvi/kratos-selfservice-ui-go/handler/profile"
	"github.com/songvi/kratos-selfservice-ui-go/handler/registration"
	"github.com/songvi/kratos-selfservice-ui-go/handler/whoami"
	"github.com/songvi/kratos-selfservice-ui-go/handler/hydra"

)

type HandlerProvider interface {
	login.LoginHandlerProvider
	registration.RegistrationHandlerProvider
	profile.ProfileHandlerProvider
	err.ErrorHandlerProvider
	logout.LogoutHandlerProvider
	whoami.WhoamiHandlerProvider
	hydra.HydraConsentHandlerProvider
	hydra.HydraLoginHandlerProvider
}
