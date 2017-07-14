package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteCallsCallLog deletes call log entry

Delete call log entry
*/
func (a *Client) DeleteCallsCallLog(params *DeleteCallsCallLogParams) (*DeleteCallsCallLogAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCallsCallLogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteCallsCallLog",
		Method:             "DELETE",
		PathPattern:        "/calls/{call}/log",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteCallsCallLogReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCallsCallLogAccepted), nil

}

/*
GetCallsCallLog gets call logs

Get call logs
*/
func (a *Client) GetCallsCallLog(params *GetCallsCallLogParams) (*GetCallsCallLogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCallsCallLogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCallsCallLog",
		Method:             "GET",
		PathPattern:        "/calls/{call}/log",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCallsCallLogReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCallsCallLogOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}