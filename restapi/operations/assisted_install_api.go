// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/openshift/assisted-service/restapi/operations/events"
	"github.com/openshift/assisted-service/restapi/operations/installer"
	"github.com/openshift/assisted-service/restapi/operations/managed_domains"
	"github.com/openshift/assisted-service/restapi/operations/versions"
)

// NewAssistedInstallAPI creates a new AssistedInstall instance
func NewAssistedInstallAPI(spec *loads.Document) *AssistedInstallAPI {
	return &AssistedInstallAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		BinProducer:  runtime.ByteStreamProducer(),
		JSONProducer: runtime.JSONProducer(),

		InstallerCancelInstallationHandler: installer.CancelInstallationHandlerFunc(func(params installer.CancelInstallationParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.CancelInstallation has not yet been implemented")
		}),
		InstallerCompleteInstallationHandler: installer.CompleteInstallationHandlerFunc(func(params installer.CompleteInstallationParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.CompleteInstallation has not yet been implemented")
		}),
		InstallerDeregisterClusterHandler: installer.DeregisterClusterHandlerFunc(func(params installer.DeregisterClusterParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.DeregisterCluster has not yet been implemented")
		}),
		InstallerDeregisterHostHandler: installer.DeregisterHostHandlerFunc(func(params installer.DeregisterHostParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.DeregisterHost has not yet been implemented")
		}),
		InstallerDisableHostHandler: installer.DisableHostHandlerFunc(func(params installer.DisableHostParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.DisableHost has not yet been implemented")
		}),
		InstallerDownloadClusterFilesHandler: installer.DownloadClusterFilesHandlerFunc(func(params installer.DownloadClusterFilesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.DownloadClusterFiles has not yet been implemented")
		}),
		InstallerDownloadClusterISOHandler: installer.DownloadClusterISOHandlerFunc(func(params installer.DownloadClusterISOParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.DownloadClusterISO has not yet been implemented")
		}),
		InstallerDownloadClusterKubeconfigHandler: installer.DownloadClusterKubeconfigHandlerFunc(func(params installer.DownloadClusterKubeconfigParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.DownloadClusterKubeconfig has not yet been implemented")
		}),
		InstallerEnableHostHandler: installer.EnableHostHandlerFunc(func(params installer.EnableHostParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.EnableHost has not yet been implemented")
		}),
		InstallerGenerateClusterISOHandler: installer.GenerateClusterISOHandlerFunc(func(params installer.GenerateClusterISOParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.GenerateClusterISO has not yet been implemented")
		}),
		InstallerGetClusterHandler: installer.GetClusterHandlerFunc(func(params installer.GetClusterParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.GetCluster has not yet been implemented")
		}),
		InstallerGetCredentialsHandler: installer.GetCredentialsHandlerFunc(func(params installer.GetCredentialsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.GetCredentials has not yet been implemented")
		}),
		InstallerGetFreeAddressesHandler: installer.GetFreeAddressesHandlerFunc(func(params installer.GetFreeAddressesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.GetFreeAddresses has not yet been implemented")
		}),
		InstallerGetHostHandler: installer.GetHostHandlerFunc(func(params installer.GetHostParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.GetHost has not yet been implemented")
		}),
		InstallerGetNextStepsHandler: installer.GetNextStepsHandlerFunc(func(params installer.GetNextStepsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.GetNextSteps has not yet been implemented")
		}),
		InstallerInstallClusterHandler: installer.InstallClusterHandlerFunc(func(params installer.InstallClusterParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.InstallCluster has not yet been implemented")
		}),
		InstallerListClustersHandler: installer.ListClustersHandlerFunc(func(params installer.ListClustersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.ListClusters has not yet been implemented")
		}),
		VersionsListComponentVersionsHandler: versions.ListComponentVersionsHandlerFunc(func(params versions.ListComponentVersionsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation versions.ListComponentVersions has not yet been implemented")
		}),
		EventsListEventsHandler: events.ListEventsHandlerFunc(func(params events.ListEventsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation events.ListEvents has not yet been implemented")
		}),
		InstallerListHostsHandler: installer.ListHostsHandlerFunc(func(params installer.ListHostsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.ListHosts has not yet been implemented")
		}),
		ManagedDomainsListManagedDomainsHandler: managed_domains.ListManagedDomainsHandlerFunc(func(params managed_domains.ListManagedDomainsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation managed_domains.ListManagedDomains has not yet been implemented")
		}),
		InstallerPostStepReplyHandler: installer.PostStepReplyHandlerFunc(func(params installer.PostStepReplyParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.PostStepReply has not yet been implemented")
		}),
		InstallerRegisterClusterHandler: installer.RegisterClusterHandlerFunc(func(params installer.RegisterClusterParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.RegisterCluster has not yet been implemented")
		}),
		InstallerRegisterHostHandler: installer.RegisterHostHandlerFunc(func(params installer.RegisterHostParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.RegisterHost has not yet been implemented")
		}),
		InstallerResetClusterHandler: installer.ResetClusterHandlerFunc(func(params installer.ResetClusterParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.ResetCluster has not yet been implemented")
		}),
		InstallerSetDebugStepHandler: installer.SetDebugStepHandlerFunc(func(params installer.SetDebugStepParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.SetDebugStep has not yet been implemented")
		}),
		InstallerUpdateClusterHandler: installer.UpdateClusterHandlerFunc(func(params installer.UpdateClusterParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.UpdateCluster has not yet been implemented")
		}),
		InstallerUpdateHostInstallProgressHandler: installer.UpdateHostInstallProgressHandlerFunc(func(params installer.UpdateHostInstallProgressParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.UpdateHostInstallProgress has not yet been implemented")
		}),
		InstallerUploadClusterIngressCertHandler: installer.UploadClusterIngressCertHandlerFunc(func(params installer.UploadClusterIngressCertParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation installer.UploadClusterIngressCert has not yet been implemented")
		}),

		// Applies when the "X-Secret-Key" header is set
		AgentAuthAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (agentAuth) X-Secret-Key from header param [X-Secret-Key] has not yet been implemented")
		},
		// Applies when the "Authorization" header is set
		UserAuthAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (userAuth) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*AssistedInstallAPI Assisted installation */
type AssistedInstallAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// BinProducer registers a producer for the following mime types:
	//   - application/octet-stream
	BinProducer runtime.Producer
	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// AgentAuthAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key X-Secret-Key provided in the header
	AgentAuthAuth func(string) (interface{}, error)

	// UserAuthAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	UserAuthAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// InstallerCancelInstallationHandler sets the operation handler for the cancel installation operation
	InstallerCancelInstallationHandler installer.CancelInstallationHandler
	// InstallerCompleteInstallationHandler sets the operation handler for the complete installation operation
	InstallerCompleteInstallationHandler installer.CompleteInstallationHandler
	// InstallerDeregisterClusterHandler sets the operation handler for the deregister cluster operation
	InstallerDeregisterClusterHandler installer.DeregisterClusterHandler
	// InstallerDeregisterHostHandler sets the operation handler for the deregister host operation
	InstallerDeregisterHostHandler installer.DeregisterHostHandler
	// InstallerDisableHostHandler sets the operation handler for the disable host operation
	InstallerDisableHostHandler installer.DisableHostHandler
	// InstallerDownloadClusterFilesHandler sets the operation handler for the download cluster files operation
	InstallerDownloadClusterFilesHandler installer.DownloadClusterFilesHandler
	// InstallerDownloadClusterISOHandler sets the operation handler for the download cluster i s o operation
	InstallerDownloadClusterISOHandler installer.DownloadClusterISOHandler
	// InstallerDownloadClusterKubeconfigHandler sets the operation handler for the download cluster kubeconfig operation
	InstallerDownloadClusterKubeconfigHandler installer.DownloadClusterKubeconfigHandler
	// InstallerEnableHostHandler sets the operation handler for the enable host operation
	InstallerEnableHostHandler installer.EnableHostHandler
	// InstallerGenerateClusterISOHandler sets the operation handler for the generate cluster i s o operation
	InstallerGenerateClusterISOHandler installer.GenerateClusterISOHandler
	// InstallerGetClusterHandler sets the operation handler for the get cluster operation
	InstallerGetClusterHandler installer.GetClusterHandler
	// InstallerGetCredentialsHandler sets the operation handler for the get credentials operation
	InstallerGetCredentialsHandler installer.GetCredentialsHandler
	// InstallerGetFreeAddressesHandler sets the operation handler for the get free addresses operation
	InstallerGetFreeAddressesHandler installer.GetFreeAddressesHandler
	// InstallerGetHostHandler sets the operation handler for the get host operation
	InstallerGetHostHandler installer.GetHostHandler
	// InstallerGetNextStepsHandler sets the operation handler for the get next steps operation
	InstallerGetNextStepsHandler installer.GetNextStepsHandler
	// InstallerInstallClusterHandler sets the operation handler for the install cluster operation
	InstallerInstallClusterHandler installer.InstallClusterHandler
	// InstallerListClustersHandler sets the operation handler for the list clusters operation
	InstallerListClustersHandler installer.ListClustersHandler
	// VersionsListComponentVersionsHandler sets the operation handler for the list component versions operation
	VersionsListComponentVersionsHandler versions.ListComponentVersionsHandler
	// EventsListEventsHandler sets the operation handler for the list events operation
	EventsListEventsHandler events.ListEventsHandler
	// InstallerListHostsHandler sets the operation handler for the list hosts operation
	InstallerListHostsHandler installer.ListHostsHandler
	// ManagedDomainsListManagedDomainsHandler sets the operation handler for the list managed domains operation
	ManagedDomainsListManagedDomainsHandler managed_domains.ListManagedDomainsHandler
	// InstallerPostStepReplyHandler sets the operation handler for the post step reply operation
	InstallerPostStepReplyHandler installer.PostStepReplyHandler
	// InstallerRegisterClusterHandler sets the operation handler for the register cluster operation
	InstallerRegisterClusterHandler installer.RegisterClusterHandler
	// InstallerRegisterHostHandler sets the operation handler for the register host operation
	InstallerRegisterHostHandler installer.RegisterHostHandler
	// InstallerResetClusterHandler sets the operation handler for the reset cluster operation
	InstallerResetClusterHandler installer.ResetClusterHandler
	// InstallerSetDebugStepHandler sets the operation handler for the set debug step operation
	InstallerSetDebugStepHandler installer.SetDebugStepHandler
	// InstallerUpdateClusterHandler sets the operation handler for the update cluster operation
	InstallerUpdateClusterHandler installer.UpdateClusterHandler
	// InstallerUpdateHostInstallProgressHandler sets the operation handler for the update host install progress operation
	InstallerUpdateHostInstallProgressHandler installer.UpdateHostInstallProgressHandler
	// InstallerUploadClusterIngressCertHandler sets the operation handler for the upload cluster ingress cert operation
	InstallerUploadClusterIngressCertHandler installer.UploadClusterIngressCertHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *AssistedInstallAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *AssistedInstallAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *AssistedInstallAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *AssistedInstallAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *AssistedInstallAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *AssistedInstallAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *AssistedInstallAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *AssistedInstallAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *AssistedInstallAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the AssistedInstallAPI
func (o *AssistedInstallAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.BinProducer == nil {
		unregistered = append(unregistered, "BinProducer")
	}
	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.AgentAuthAuth == nil {
		unregistered = append(unregistered, "XSecretKeyAuth")
	}
	if o.UserAuthAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.InstallerCancelInstallationHandler == nil {
		unregistered = append(unregistered, "installer.CancelInstallationHandler")
	}
	if o.InstallerCompleteInstallationHandler == nil {
		unregistered = append(unregistered, "installer.CompleteInstallationHandler")
	}
	if o.InstallerDeregisterClusterHandler == nil {
		unregistered = append(unregistered, "installer.DeregisterClusterHandler")
	}
	if o.InstallerDeregisterHostHandler == nil {
		unregistered = append(unregistered, "installer.DeregisterHostHandler")
	}
	if o.InstallerDisableHostHandler == nil {
		unregistered = append(unregistered, "installer.DisableHostHandler")
	}
	if o.InstallerDownloadClusterFilesHandler == nil {
		unregistered = append(unregistered, "installer.DownloadClusterFilesHandler")
	}
	if o.InstallerDownloadClusterISOHandler == nil {
		unregistered = append(unregistered, "installer.DownloadClusterISOHandler")
	}
	if o.InstallerDownloadClusterKubeconfigHandler == nil {
		unregistered = append(unregistered, "installer.DownloadClusterKubeconfigHandler")
	}
	if o.InstallerEnableHostHandler == nil {
		unregistered = append(unregistered, "installer.EnableHostHandler")
	}
	if o.InstallerGenerateClusterISOHandler == nil {
		unregistered = append(unregistered, "installer.GenerateClusterISOHandler")
	}
	if o.InstallerGetClusterHandler == nil {
		unregistered = append(unregistered, "installer.GetClusterHandler")
	}
	if o.InstallerGetCredentialsHandler == nil {
		unregistered = append(unregistered, "installer.GetCredentialsHandler")
	}
	if o.InstallerGetFreeAddressesHandler == nil {
		unregistered = append(unregistered, "installer.GetFreeAddressesHandler")
	}
	if o.InstallerGetHostHandler == nil {
		unregistered = append(unregistered, "installer.GetHostHandler")
	}
	if o.InstallerGetNextStepsHandler == nil {
		unregistered = append(unregistered, "installer.GetNextStepsHandler")
	}
	if o.InstallerInstallClusterHandler == nil {
		unregistered = append(unregistered, "installer.InstallClusterHandler")
	}
	if o.InstallerListClustersHandler == nil {
		unregistered = append(unregistered, "installer.ListClustersHandler")
	}
	if o.VersionsListComponentVersionsHandler == nil {
		unregistered = append(unregistered, "versions.ListComponentVersionsHandler")
	}
	if o.EventsListEventsHandler == nil {
		unregistered = append(unregistered, "events.ListEventsHandler")
	}
	if o.InstallerListHostsHandler == nil {
		unregistered = append(unregistered, "installer.ListHostsHandler")
	}
	if o.ManagedDomainsListManagedDomainsHandler == nil {
		unregistered = append(unregistered, "managed_domains.ListManagedDomainsHandler")
	}
	if o.InstallerPostStepReplyHandler == nil {
		unregistered = append(unregistered, "installer.PostStepReplyHandler")
	}
	if o.InstallerRegisterClusterHandler == nil {
		unregistered = append(unregistered, "installer.RegisterClusterHandler")
	}
	if o.InstallerRegisterHostHandler == nil {
		unregistered = append(unregistered, "installer.RegisterHostHandler")
	}
	if o.InstallerResetClusterHandler == nil {
		unregistered = append(unregistered, "installer.ResetClusterHandler")
	}
	if o.InstallerSetDebugStepHandler == nil {
		unregistered = append(unregistered, "installer.SetDebugStepHandler")
	}
	if o.InstallerUpdateClusterHandler == nil {
		unregistered = append(unregistered, "installer.UpdateClusterHandler")
	}
	if o.InstallerUpdateHostInstallProgressHandler == nil {
		unregistered = append(unregistered, "installer.UpdateHostInstallProgressHandler")
	}
	if o.InstallerUploadClusterIngressCertHandler == nil {
		unregistered = append(unregistered, "installer.UploadClusterIngressCertHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *AssistedInstallAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *AssistedInstallAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "agentAuth":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.AgentAuthAuth)

		case "userAuth":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.UserAuthAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *AssistedInstallAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *AssistedInstallAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *AssistedInstallAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/octet-stream":
			result["application/octet-stream"] = o.BinProducer
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *AssistedInstallAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the assisted install API
func (o *AssistedInstallAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *AssistedInstallAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/clusters/{cluster_id}/actions/cancel"] = installer.NewCancelInstallation(o.context, o.InstallerCancelInstallationHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/clusters/{cluster_id}/actions/complete_installation"] = installer.NewCompleteInstallation(o.context, o.InstallerCompleteInstallationHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/clusters/{cluster_id}"] = installer.NewDeregisterCluster(o.context, o.InstallerDeregisterClusterHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/clusters/{cluster_id}/hosts/{host_id}"] = installer.NewDeregisterHost(o.context, o.InstallerDeregisterHostHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/clusters/{cluster_id}/hosts/{host_id}/actions/enable"] = installer.NewDisableHost(o.context, o.InstallerDisableHostHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clusters/{cluster_id}/downloads/files"] = installer.NewDownloadClusterFiles(o.context, o.InstallerDownloadClusterFilesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clusters/{cluster_id}/downloads/image"] = installer.NewDownloadClusterISO(o.context, o.InstallerDownloadClusterISOHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clusters/{cluster_id}/downloads/kubeconfig"] = installer.NewDownloadClusterKubeconfig(o.context, o.InstallerDownloadClusterKubeconfigHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/clusters/{cluster_id}/hosts/{host_id}/actions/enable"] = installer.NewEnableHost(o.context, o.InstallerEnableHostHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/clusters/{cluster_id}/downloads/image"] = installer.NewGenerateClusterISO(o.context, o.InstallerGenerateClusterISOHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clusters/{cluster_id}"] = installer.NewGetCluster(o.context, o.InstallerGetClusterHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clusters/{cluster_id}/credentials"] = installer.NewGetCredentials(o.context, o.InstallerGetCredentialsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clusters/{cluster_id}/free_addresses"] = installer.NewGetFreeAddresses(o.context, o.InstallerGetFreeAddressesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clusters/{cluster_id}/hosts/{host_id}"] = installer.NewGetHost(o.context, o.InstallerGetHostHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clusters/{cluster_id}/hosts/{host_id}/instructions"] = installer.NewGetNextSteps(o.context, o.InstallerGetNextStepsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/clusters/{cluster_id}/actions/install"] = installer.NewInstallCluster(o.context, o.InstallerInstallClusterHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clusters"] = installer.NewListClusters(o.context, o.InstallerListClustersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/component_versions"] = versions.NewListComponentVersions(o.context, o.VersionsListComponentVersionsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/events/{entity_id}"] = events.NewListEvents(o.context, o.EventsListEventsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clusters/{cluster_id}/hosts"] = installer.NewListHosts(o.context, o.InstallerListHostsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/domains"] = managed_domains.NewListManagedDomains(o.context, o.ManagedDomainsListManagedDomainsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/clusters/{cluster_id}/hosts/{host_id}/instructions"] = installer.NewPostStepReply(o.context, o.InstallerPostStepReplyHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/clusters"] = installer.NewRegisterCluster(o.context, o.InstallerRegisterClusterHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/clusters/{cluster_id}/hosts"] = installer.NewRegisterHost(o.context, o.InstallerRegisterHostHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/clusters/{cluster_id}/actions/reset"] = installer.NewResetCluster(o.context, o.InstallerResetClusterHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/clusters/{cluster_id}/hosts/{host_id}/actions/debug"] = installer.NewSetDebugStep(o.context, o.InstallerSetDebugStepHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/clusters/{cluster_id}"] = installer.NewUpdateCluster(o.context, o.InstallerUpdateClusterHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/clusters/{cluster_id}/hosts/{host_id}/progress"] = installer.NewUpdateHostInstallProgress(o.context, o.InstallerUpdateHostInstallProgressHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/clusters/{cluster_id}/uploads/ingress-cert"] = installer.NewUploadClusterIngressCert(o.context, o.InstallerUploadClusterIngressCertHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *AssistedInstallAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *AssistedInstallAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *AssistedInstallAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *AssistedInstallAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *AssistedInstallAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
