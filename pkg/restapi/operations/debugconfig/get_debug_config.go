package debugconfig

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetDebugConfigHandlerFunc turns a function with the right signature into a get debug config handler
type GetDebugConfigHandlerFunc func(GetDebugConfigParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDebugConfigHandlerFunc) Handle(params GetDebugConfigParams) middleware.Responder {
	return fn(params)
}

// GetDebugConfigHandler interface for that can handle valid get debug config params
type GetDebugConfigHandler interface {
	Handle(GetDebugConfigParams) middleware.Responder
}

// NewGetDebugConfig creates a new http.Handler for the get debug config operation
func NewGetDebugConfig(ctx *middleware.Context, handler GetDebugConfigHandler) *GetDebugConfig {
	return &GetDebugConfig{Context: ctx, Handler: handler}
}

/*GetDebugConfig swagger:route GET /debugconfig/{debugConfigId} debugconfig getDebugConfig

Retrun a debug config

Retrun a debug config

*/
type GetDebugConfig struct {
	Context *middleware.Context
	Handler GetDebugConfigHandler
}

func (o *GetDebugConfig) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetDebugConfigParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}