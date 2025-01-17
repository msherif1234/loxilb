// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedHandlerFunc turns a function with the right signature into a delete config vlan vlan ID member if name tagged tagged handler
type DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedHandlerFunc func(DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedHandlerFunc) Handle(params DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedParams) middleware.Responder {
	return fn(params)
}

// DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedHandler interface for that can handle valid delete config vlan vlan ID member if name tagged tagged params
type DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedHandler interface {
	Handle(DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedParams) middleware.Responder
}

// NewDeleteConfigVlanVlanIDMemberIfNameTaggedTagged creates a new http.Handler for the delete config vlan vlan ID member if name tagged tagged operation
func NewDeleteConfigVlanVlanIDMemberIfNameTaggedTagged(ctx *middleware.Context, handler DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedHandler) *DeleteConfigVlanVlanIDMemberIfNameTaggedTagged {
	return &DeleteConfigVlanVlanIDMemberIfNameTaggedTagged{Context: ctx, Handler: handler}
}

/* DeleteConfigVlanVlanIDMemberIfNameTaggedTagged swagger:route DELETE /config/vlan/{vlan_id}/member/{if_name}/tagged/{tagged} deleteConfigVlanVlanIdMemberIfNameTaggedTagged

Remove a vlan member from a vlan interface

Remove a vlan member from a vlan interface which is defined by vlan_id. If the Vlan interface does not exist on LoxiLB OR a vlan member 'if_name' is not present on the interface the API will return '404'. If the vlan_id passed is less than 2 or greater than 4094 the API will respond with error '400'.

*/
type DeleteConfigVlanVlanIDMemberIfNameTaggedTagged struct {
	Context *middleware.Context
	Handler DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedHandler
}

func (o *DeleteConfigVlanVlanIDMemberIfNameTaggedTagged) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteConfigVlanVlanIDMemberIfNameTaggedTaggedParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
