package configuration

import (
	"github.com/ory/viper"
)

type ViperProvider struct{}

var (
	ViperKeyListenOn    = "listen"
	ViperKeyTemplateDir = "templatedir"

	ViperKeyEndpointProfile      = "endpoints.profile"
	ViperKeyEndpointLogin        = "endpoints.login"
	ViperKeyEndpointVerification = "endpoints.verification"
	ViperKeyEndpointError        = "endpoints.error"
	ViperKeyEndpointMultiFactors = "endpoints.multifactors"
	ViperKeyEndpointRegistration = "endpoints.registration"
	ViperKeyEndpointLogout       = "endpoints.logout"

	ViperKeyKratosAdminUrl            = "kratos.services.admin"
	ViperKeyKratosPublicUrl            = "kratos.services.public"

	ViperKeyKratosBrowserInitPath      = "kratos.browser.init-path"
	ViperKeyKratosBrowserRequestPath      = "kratos.browser.request-path"

	ViperKeyHydraLoginProviderUrl   = "kratos.hydra.login_provider"
	ViperKeyHydraConsentProviderUrl = "kratos.hydra.consent_provider"
	ViperKeyHydraAPILoginUrl        = "kratos.hydra.api_login_url"
	ViperKeyHydraAPIConsentUrl      = "kratos.hydra.api_consent_url"
)

func NewViperProvider() *ViperProvider {
	return &ViperProvider{}
}

func (v *ViperProvider) ListenOn() string {
	return viper.GetString(ViperKeyListenOn)
}

func (v *ViperProvider) TemplateDir() string {
	return viper.GetString(ViperKeyTemplateDir)
}

func (v *ViperProvider) ProfileUrl() string {
	return parseURLFromViper(ViperKeyEndpointProfile)
}

func (v *ViperProvider) LogoutUrl() string {
	return parseURLFromViper(ViperKeyEndpointLogout)
}


func (v *ViperProvider) LoginUrl() string {
	return parseURLFromViper(ViperKeyEndpointLogin)
}

func (v *ViperProvider) MultiFactorsUrl() string {
	return parseURLFromViper(ViperKeyEndpointMultiFactors)
}

func (v *ViperProvider) VerificationUrl() string {
	return parseURLFromViper(ViperKeyEndpointVerification)
}

func (v *ViperProvider) ErrorUrl() string {
	return parseURLFromViper(ViperKeyEndpointError)
}

func (v *ViperProvider) RegistrerUrl() string {
	return parseURLFromViper(ViperKeyEndpointRegistration)
}

func (v *ViperProvider) KratosAdminUrl() string {
	return parseURLFromViper(ViperKeyKratosAdminUrl)
}

func (v *ViperProvider) KratosPublicUrl() string {
	return parseURLFromViper(ViperKeyKratosPublicUrl)
}

func (v *ViperProvider) KratosBrowserInitPath() string {
	return parseURLFromViper(ViperKeyKratosBrowserInitPath)
}

func (v *ViperProvider) KratosBrowserRequestPath() string {
	return parseURLFromViper(ViperKeyKratosBrowserRequestPath)
}

func (v *ViperProvider) HydraAPILoginUrl() string {
	return parseURLFromViper(ViperKeyHydraAPILoginUrl)
}

func (v *ViperProvider) HydraAPIConsentUrl() string {
	return parseURLFromViper(ViperKeyHydraAPIConsentUrl)
}

func (v *ViperProvider) HydraLoginProviderUrl() string {
	return parseURLFromViper(ViperKeyHydraLoginProviderUrl)
}

func (v *ViperProvider) HydraConsentProviderUrl() string {
	return parseURLFromViper(ViperKeyHydraConsentProviderUrl)
}

func parseURLFromViper(key string) string {
	return viper.GetString(key)
}
