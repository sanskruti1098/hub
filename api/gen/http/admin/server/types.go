// Code generated by goa v3.15.2, DO NOT EDIT.
//
// admin HTTP server types
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package server

import (
	admin "github.com/tektoncd/hub/api/gen/admin"
	goa "goa.design/goa/v3/pkg"
)

// UpdateAgentRequestBody is the type of the "admin" service "UpdateAgent"
// endpoint HTTP request body.
type UpdateAgentRequestBody struct {
	// Name of Agent
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Scopes required for Agent
	Scopes []string `form:"scopes,omitempty" json:"scopes,omitempty" xml:"scopes,omitempty"`
}

// UpdateAgentResponseBody is the type of the "admin" service "UpdateAgent"
// endpoint HTTP response body.
type UpdateAgentResponseBody struct {
	// Agent JWT
	Token string `form:"token" json:"token" xml:"token"`
}

// RefreshConfigResponseBody is the type of the "admin" service "RefreshConfig"
// endpoint HTTP response body.
type RefreshConfigResponseBody struct {
	// Config file checksum
	Checksum string `form:"checksum" json:"checksum" xml:"checksum"`
}

// UpdateAgentInvalidPayloadResponseBody is the type of the "admin" service
// "UpdateAgent" endpoint HTTP response body for the "invalid-payload" error.
type UpdateAgentInvalidPayloadResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateAgentInvalidTokenResponseBody is the type of the "admin" service
// "UpdateAgent" endpoint HTTP response body for the "invalid-token" error.
type UpdateAgentInvalidTokenResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateAgentInvalidScopesResponseBody is the type of the "admin" service
// "UpdateAgent" endpoint HTTP response body for the "invalid-scopes" error.
type UpdateAgentInvalidScopesResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateAgentInternalErrorResponseBody is the type of the "admin" service
// "UpdateAgent" endpoint HTTP response body for the "internal-error" error.
type UpdateAgentInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RefreshConfigInvalidTokenResponseBody is the type of the "admin" service
// "RefreshConfig" endpoint HTTP response body for the "invalid-token" error.
type RefreshConfigInvalidTokenResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RefreshConfigInvalidScopesResponseBody is the type of the "admin" service
// "RefreshConfig" endpoint HTTP response body for the "invalid-scopes" error.
type RefreshConfigInvalidScopesResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RefreshConfigInternalErrorResponseBody is the type of the "admin" service
// "RefreshConfig" endpoint HTTP response body for the "internal-error" error.
type RefreshConfigInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewUpdateAgentResponseBody builds the HTTP response body from the result of
// the "UpdateAgent" endpoint of the "admin" service.
func NewUpdateAgentResponseBody(res *admin.UpdateAgentResult) *UpdateAgentResponseBody {
	body := &UpdateAgentResponseBody{
		Token: res.Token,
	}
	return body
}

// NewRefreshConfigResponseBody builds the HTTP response body from the result
// of the "RefreshConfig" endpoint of the "admin" service.
func NewRefreshConfigResponseBody(res *admin.RefreshConfigResult) *RefreshConfigResponseBody {
	body := &RefreshConfigResponseBody{
		Checksum: res.Checksum,
	}
	return body
}

// NewUpdateAgentInvalidPayloadResponseBody builds the HTTP response body from
// the result of the "UpdateAgent" endpoint of the "admin" service.
func NewUpdateAgentInvalidPayloadResponseBody(res *goa.ServiceError) *UpdateAgentInvalidPayloadResponseBody {
	body := &UpdateAgentInvalidPayloadResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateAgentInvalidTokenResponseBody builds the HTTP response body from
// the result of the "UpdateAgent" endpoint of the "admin" service.
func NewUpdateAgentInvalidTokenResponseBody(res *goa.ServiceError) *UpdateAgentInvalidTokenResponseBody {
	body := &UpdateAgentInvalidTokenResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateAgentInvalidScopesResponseBody builds the HTTP response body from
// the result of the "UpdateAgent" endpoint of the "admin" service.
func NewUpdateAgentInvalidScopesResponseBody(res *goa.ServiceError) *UpdateAgentInvalidScopesResponseBody {
	body := &UpdateAgentInvalidScopesResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateAgentInternalErrorResponseBody builds the HTTP response body from
// the result of the "UpdateAgent" endpoint of the "admin" service.
func NewUpdateAgentInternalErrorResponseBody(res *goa.ServiceError) *UpdateAgentInternalErrorResponseBody {
	body := &UpdateAgentInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRefreshConfigInvalidTokenResponseBody builds the HTTP response body from
// the result of the "RefreshConfig" endpoint of the "admin" service.
func NewRefreshConfigInvalidTokenResponseBody(res *goa.ServiceError) *RefreshConfigInvalidTokenResponseBody {
	body := &RefreshConfigInvalidTokenResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRefreshConfigInvalidScopesResponseBody builds the HTTP response body from
// the result of the "RefreshConfig" endpoint of the "admin" service.
func NewRefreshConfigInvalidScopesResponseBody(res *goa.ServiceError) *RefreshConfigInvalidScopesResponseBody {
	body := &RefreshConfigInvalidScopesResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRefreshConfigInternalErrorResponseBody builds the HTTP response body from
// the result of the "RefreshConfig" endpoint of the "admin" service.
func NewRefreshConfigInternalErrorResponseBody(res *goa.ServiceError) *RefreshConfigInternalErrorResponseBody {
	body := &RefreshConfigInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateAgentPayload builds a admin service UpdateAgent endpoint payload.
func NewUpdateAgentPayload(body *UpdateAgentRequestBody, token string) *admin.UpdateAgentPayload {
	v := &admin.UpdateAgentPayload{
		Name: *body.Name,
	}
	v.Scopes = make([]string, len(body.Scopes))
	for i, val := range body.Scopes {
		v.Scopes[i] = val
	}
	v.Token = token

	return v
}

// NewRefreshConfigPayload builds a admin service RefreshConfig endpoint
// payload.
func NewRefreshConfigPayload(token string) *admin.RefreshConfigPayload {
	v := &admin.RefreshConfigPayload{}
	v.Token = token

	return v
}

// ValidateUpdateAgentRequestBody runs the validations defined on
// UpdateAgentRequestBody
func ValidateUpdateAgentRequestBody(body *UpdateAgentRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Scopes == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("scopes", "body"))
	}
	return
}
