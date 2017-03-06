package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteIPAMIPHandlerFunc turns a function with the right signature into a delete IP a m IP handler
type DeleteIPAMIPHandlerFunc func(DeleteIPAMIPParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteIPAMIPHandlerFunc) Handle(params DeleteIPAMIPParams) middleware.Responder {
	return fn(params)
}

// DeleteIPAMIPHandler interface for that can handle valid delete IP a m IP params
type DeleteIPAMIPHandler interface {
	Handle(DeleteIPAMIPParams) middleware.Responder
}

// NewDeleteIPAMIP creates a new http.Handler for the delete IP a m IP operation
func NewDeleteIPAMIP(ctx *middleware.Context, handler DeleteIPAMIPHandler) *DeleteIPAMIP {
	return &DeleteIPAMIP{Context: ctx, Handler: handler}
}

/*DeleteIPAMIP swagger:route DELETE /ipam/{ip} ipam deleteIpAMIp

Release an allocated IP address

*/
type DeleteIPAMIP struct {
	Context *middleware.Context
	Handler DeleteIPAMIPHandler
}

func (o *DeleteIPAMIP) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDeleteIPAMIPParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}