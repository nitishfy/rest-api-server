// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetUsersUserHandlerFunc turns a function with the right signature into a get users user handler
type GetUsersUserHandlerFunc func(GetUsersUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUsersUserHandlerFunc) Handle(params GetUsersUserParams) middleware.Responder {
	return fn(params)
}

// GetUsersUserHandler interface for that can handle valid get users user params
type GetUsersUserHandler interface {
	Handle(GetUsersUserParams) middleware.Responder
}

// NewGetUsersUser creates a new http.Handler for the get users user operation
func NewGetUsersUser(ctx *middleware.Context, handler GetUsersUserHandler) *GetUsersUser {
	return &GetUsersUser{Context: ctx, Handler: handler}
}

/*
	GetUsersUser swagger:route GET /users/{user} getUsersUser

Returns greeting to the user
*/
type GetUsersUser struct {
	Context *middleware.Context
	Handler GetUsersUserHandler
}

func (o *GetUsersUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetUsersUserParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
