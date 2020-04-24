package configuration

type ConfigProvider interface {
	ListenOn() string // ie: 127.0.0.1:1983
	TemplateDir() string

	ProfileUrl() string
	LoginUrl() string
	VerificationUrl() string	
	MultiFactorsUrl() string
	RegistrerUrl() string
	ErrorUrl() string
	LogoutUrl() string

	KratosProfileFlowUrl() string
	KratosLoginFlowUrl() string
	KratosRegistrationFlowUrl() string
	KratosPublicFlowsUrl() string
	KratosLogoutFlowsUrl() string
	KratosAdminUrl() string

	HydraAPILoginUrl() string
	HydraAPIConsentUrl() string
	HydraLoginProviderUrl()	string
	HydraConsentProviderUrl() string
}
