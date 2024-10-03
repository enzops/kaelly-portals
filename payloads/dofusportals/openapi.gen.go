// Package dofusportals provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package dofusportals

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oapi-codegen/runtime"
)

const (
	ApiKeyAuthScopes = "ApiKeyAuth.Scopes"
)

// Defines values for IntendedErrorError.
const (
	AcceptLanguageIsBad    IntendedErrorError = "accept_language.is_bad"
	AcceptLanguageNotFound IntendedErrorError = "accept_language.not_found"
	DimensionNotFound      IntendedErrorError = "dimension.not_found"
	IdIsBad                IntendedErrorError = "id.is_bad"
	LoginNotFound          IntendedErrorError = "login.not_found"
	MailIsBad              IntendedErrorError = "mail.is_bad"
	MailNotFound           IntendedErrorError = "mail.not_found"
	NameIsBad              IntendedErrorError = "name.is_bad"
	NameNotFound           IntendedErrorError = "name.not_found"
	PasswordIsBad          IntendedErrorError = "password.is_bad"
	PasswordNotFound       IntendedErrorError = "password.not_found"
	PortalIsBad            IntendedErrorError = "portal.is_bad"
	PortalPositionIsBad    IntendedErrorError = "portal.position.is_bad"
	ServerNotFound         IntendedErrorError = "server.not_found"
	TokenNotFound          IntendedErrorError = "token.not_found"
)

// Defines values for ServerCommunity.
const (
	All ServerCommunity = "all"
	Es  ServerCommunity = "es"
	Fr  ServerCommunity = "fr"
	Pt  ServerCommunity = "pt"
)

// Defines values for ServerType.
const (
	Epic   ServerType = "epic"
	Event  ServerType = "event"
	Heroic ServerType = "heroic"
	Mono   ServerType = "mono"
	Multi  ServerType = "multi"
)

// Defines values for TransportType.
const (
	Brigandin   TransportType = "brigandin"
	CharAVoile  TransportType = "char_a_voile"
	Diligence   TransportType = "diligence"
	Foreuse     TransportType = "foreuse"
	Frigostien  TransportType = "frigostien"
	Scaeroplane TransportType = "scaeroplane"
	Skis        TransportType = "skis"
	Zaap        TransportType = "zaap"
)

// Dimension Dofus game dimension
type Dimension struct {
	// Id Name of the dimension, used as an id.
	Id string `json:"id"`
}

// IntendedError Error that comes when trying to validate and sanitize inputs
type IntendedError struct {
	// Error Code of the functional error
	Error IntendedErrorError `json:"error"`
}

// IntendedErrorError Code of the functional error
type IntendedErrorError string

// Portal Portal position at a specific moment. If the position is not filled, then the portal position is considered as unknown and some fields will not be present.
type Portal struct {
	// CreatedAt Date of the last updated position.
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// CreatedBy User that has added or updated the last position
	CreatedBy *User `json:"createdBy,omitempty"`

	// Dimension Name of the dimension
	Dimension string `json:"dimension"`

	// Position Position [x, y] in Dofus map
	Position *Position `json:"position,omitempty"`

	// RemainingUses Remaining number of uses.
	RemainingUses *float32 `json:"remainingUses,omitempty"`

	// Server Name of the server
	Server string `json:"server"`

	// UpdatedAt Date of the last updated portal.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// UpdatedBy User that has added or updated the last position
	UpdatedBy *User `json:"updatedBy,omitempty"`
}

// PortalHistory Portal position value at a specific period.
type PortalHistory struct {
	// IsInCanopy Is the portal position in the canopy or not.
	IsInCanopy *bool `json:"isInCanopy,omitempty"`

	// Position Position [x, y] in Dofus map
	Position *Position `json:"position,omitempty"`

	// RemainingUses Remaining number of uses.
	RemainingUses *float32 `json:"remainingUses,omitempty"`
}

// Position Position [x, y] in Dofus map
type Position struct {
	// ConditionalTransport Transport used to access to a position easily
	ConditionalTransport *Transport `json:"conditionalTransport,omitempty"`

	// IsInCanopy Is the position in the canopy or not.
	IsInCanopy *bool `json:"isInCanopy,omitempty"`

	// Transport Transport used to access to a position easily
	Transport *Transport `json:"transport,omitempty"`

	// X x value in [x, y].
	X float32 `json:"x"`

	// Y y value in [x, y].
	Y float32 `json:"y"`
}

// Server Dofus game server
type Server struct {
	// Active Define if the server is accessible for players or not
	Active bool `json:"active"`

	// Community Community of the server. Most of the time, it is a specific region in Europe.
	Community ServerCommunity `json:"community"`

	// Id Name of the server, used as an id.
	Id string `json:"id"`

	// Type Type of the server. Most of the servers are multi-account.
	Type ServerType `json:"type"`
}

// ServerCommunity Community of the server. Most of the time, it is a specific region in Europe.
type ServerCommunity string

// ServerType Type of the server. Most of the servers are multi-account.
type ServerType string

// Transport Transport used to access to a position easily
type Transport struct {
	// Area Area where the transport is localized.
	Area string `json:"area"`

	// SubArea Sub area where the transport is localized.
	SubArea string `json:"subArea"`

	// Type Type of the transport. Most of the transport are zaaps.
	Type TransportType `json:"type"`

	// X x value in [x, y].
	X float32 `json:"x"`

	// Y y value in [x, y].
	Y float32 `json:"y"`
}

// TransportType Type of the transport. Most of the transport are zaaps.
type TransportType string

// User User that has added or updated the last position
type User struct {
	// Name Name of the user
	Name string `json:"name"`
}

// IntendedErrorResponse Error that comes when trying to validate and sanitize inputs
type IntendedErrorResponse = IntendedError

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetExternalV1Dimensions request
	GetExternalV1Dimensions(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetExternalV1Servers request
	GetExternalV1Servers(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetExternalV1ServersServerIdPortals request
	GetExternalV1ServersServerIdPortals(ctx context.Context, serverId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetExternalV1ServersServerIdPortalsDimensionId request
	GetExternalV1ServersServerIdPortalsDimensionId(ctx context.Context, serverId string, dimensionId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetExternalV1ServersServerIdPortalsDimensionIdHistory request
	GetExternalV1ServersServerIdPortalsDimensionIdHistory(ctx context.Context, serverId string, dimensionId string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetExternalV1Dimensions(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetExternalV1DimensionsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetExternalV1Servers(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetExternalV1ServersRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetExternalV1ServersServerIdPortals(ctx context.Context, serverId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetExternalV1ServersServerIdPortalsRequest(c.Server, serverId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetExternalV1ServersServerIdPortalsDimensionId(ctx context.Context, serverId string, dimensionId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetExternalV1ServersServerIdPortalsDimensionIdRequest(c.Server, serverId, dimensionId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetExternalV1ServersServerIdPortalsDimensionIdHistory(ctx context.Context, serverId string, dimensionId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetExternalV1ServersServerIdPortalsDimensionIdHistoryRequest(c.Server, serverId, dimensionId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetExternalV1DimensionsRequest generates requests for GetExternalV1Dimensions
func NewGetExternalV1DimensionsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/external/v1/dimensions")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetExternalV1ServersRequest generates requests for GetExternalV1Servers
func NewGetExternalV1ServersRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/external/v1/servers")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetExternalV1ServersServerIdPortalsRequest generates requests for GetExternalV1ServersServerIdPortals
func NewGetExternalV1ServersServerIdPortalsRequest(server string, serverId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "serverId", runtime.ParamLocationPath, serverId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/external/v1/servers/%s/portals", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetExternalV1ServersServerIdPortalsDimensionIdRequest generates requests for GetExternalV1ServersServerIdPortalsDimensionId
func NewGetExternalV1ServersServerIdPortalsDimensionIdRequest(server string, serverId string, dimensionId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "serverId", runtime.ParamLocationPath, serverId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "dimensionId", runtime.ParamLocationPath, dimensionId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/external/v1/servers/%s/portals/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetExternalV1ServersServerIdPortalsDimensionIdHistoryRequest generates requests for GetExternalV1ServersServerIdPortalsDimensionIdHistory
func NewGetExternalV1ServersServerIdPortalsDimensionIdHistoryRequest(server string, serverId string, dimensionId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "serverId", runtime.ParamLocationPath, serverId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "dimensionId", runtime.ParamLocationPath, dimensionId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/external/v1/servers/%s/portals/%s/history", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetExternalV1DimensionsWithResponse request
	GetExternalV1DimensionsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetExternalV1DimensionsResponse, error)

	// GetExternalV1ServersWithResponse request
	GetExternalV1ServersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetExternalV1ServersResponse, error)

	// GetExternalV1ServersServerIdPortalsWithResponse request
	GetExternalV1ServersServerIdPortalsWithResponse(ctx context.Context, serverId string, reqEditors ...RequestEditorFn) (*GetExternalV1ServersServerIdPortalsResponse, error)

	// GetExternalV1ServersServerIdPortalsDimensionIdWithResponse request
	GetExternalV1ServersServerIdPortalsDimensionIdWithResponse(ctx context.Context, serverId string, dimensionId string, reqEditors ...RequestEditorFn) (*GetExternalV1ServersServerIdPortalsDimensionIdResponse, error)

	// GetExternalV1ServersServerIdPortalsDimensionIdHistoryWithResponse request
	GetExternalV1ServersServerIdPortalsDimensionIdHistoryWithResponse(ctx context.Context, serverId string, dimensionId string, reqEditors ...RequestEditorFn) (*GetExternalV1ServersServerIdPortalsDimensionIdHistoryResponse, error)
}

type GetExternalV1DimensionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Dimension
	JSON400      *IntendedErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetExternalV1DimensionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetExternalV1DimensionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetExternalV1ServersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Server
	JSON400      *IntendedErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetExternalV1ServersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetExternalV1ServersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetExternalV1ServersServerIdPortalsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Portal
	JSON400      *IntendedErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetExternalV1ServersServerIdPortalsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetExternalV1ServersServerIdPortalsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetExternalV1ServersServerIdPortalsDimensionIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Portal
	JSON400      *IntendedErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetExternalV1ServersServerIdPortalsDimensionIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetExternalV1ServersServerIdPortalsDimensionIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetExternalV1ServersServerIdPortalsDimensionIdHistoryResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]PortalHistory
	JSON400      *IntendedErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetExternalV1ServersServerIdPortalsDimensionIdHistoryResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetExternalV1ServersServerIdPortalsDimensionIdHistoryResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetExternalV1DimensionsWithResponse request returning *GetExternalV1DimensionsResponse
func (c *ClientWithResponses) GetExternalV1DimensionsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetExternalV1DimensionsResponse, error) {
	rsp, err := c.GetExternalV1Dimensions(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetExternalV1DimensionsResponse(rsp)
}

// GetExternalV1ServersWithResponse request returning *GetExternalV1ServersResponse
func (c *ClientWithResponses) GetExternalV1ServersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetExternalV1ServersResponse, error) {
	rsp, err := c.GetExternalV1Servers(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetExternalV1ServersResponse(rsp)
}

// GetExternalV1ServersServerIdPortalsWithResponse request returning *GetExternalV1ServersServerIdPortalsResponse
func (c *ClientWithResponses) GetExternalV1ServersServerIdPortalsWithResponse(ctx context.Context, serverId string, reqEditors ...RequestEditorFn) (*GetExternalV1ServersServerIdPortalsResponse, error) {
	rsp, err := c.GetExternalV1ServersServerIdPortals(ctx, serverId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetExternalV1ServersServerIdPortalsResponse(rsp)
}

// GetExternalV1ServersServerIdPortalsDimensionIdWithResponse request returning *GetExternalV1ServersServerIdPortalsDimensionIdResponse
func (c *ClientWithResponses) GetExternalV1ServersServerIdPortalsDimensionIdWithResponse(ctx context.Context, serverId string, dimensionId string, reqEditors ...RequestEditorFn) (*GetExternalV1ServersServerIdPortalsDimensionIdResponse, error) {
	rsp, err := c.GetExternalV1ServersServerIdPortalsDimensionId(ctx, serverId, dimensionId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetExternalV1ServersServerIdPortalsDimensionIdResponse(rsp)
}

// GetExternalV1ServersServerIdPortalsDimensionIdHistoryWithResponse request returning *GetExternalV1ServersServerIdPortalsDimensionIdHistoryResponse
func (c *ClientWithResponses) GetExternalV1ServersServerIdPortalsDimensionIdHistoryWithResponse(ctx context.Context, serverId string, dimensionId string, reqEditors ...RequestEditorFn) (*GetExternalV1ServersServerIdPortalsDimensionIdHistoryResponse, error) {
	rsp, err := c.GetExternalV1ServersServerIdPortalsDimensionIdHistory(ctx, serverId, dimensionId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetExternalV1ServersServerIdPortalsDimensionIdHistoryResponse(rsp)
}

// ParseGetExternalV1DimensionsResponse parses an HTTP response from a GetExternalV1DimensionsWithResponse call
func ParseGetExternalV1DimensionsResponse(rsp *http.Response) (*GetExternalV1DimensionsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetExternalV1DimensionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Dimension
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest IntendedErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseGetExternalV1ServersResponse parses an HTTP response from a GetExternalV1ServersWithResponse call
func ParseGetExternalV1ServersResponse(rsp *http.Response) (*GetExternalV1ServersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetExternalV1ServersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Server
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest IntendedErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseGetExternalV1ServersServerIdPortalsResponse parses an HTTP response from a GetExternalV1ServersServerIdPortalsWithResponse call
func ParseGetExternalV1ServersServerIdPortalsResponse(rsp *http.Response) (*GetExternalV1ServersServerIdPortalsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetExternalV1ServersServerIdPortalsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Portal
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest IntendedErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseGetExternalV1ServersServerIdPortalsDimensionIdResponse parses an HTTP response from a GetExternalV1ServersServerIdPortalsDimensionIdWithResponse call
func ParseGetExternalV1ServersServerIdPortalsDimensionIdResponse(rsp *http.Response) (*GetExternalV1ServersServerIdPortalsDimensionIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetExternalV1ServersServerIdPortalsDimensionIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Portal
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest IntendedErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseGetExternalV1ServersServerIdPortalsDimensionIdHistoryResponse parses an HTTP response from a GetExternalV1ServersServerIdPortalsDimensionIdHistoryWithResponse call
func ParseGetExternalV1ServersServerIdPortalsDimensionIdHistoryResponse(rsp *http.Response) (*GetExternalV1ServersServerIdPortalsDimensionIdHistoryResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetExternalV1ServersServerIdPortalsDimensionIdHistoryResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []PortalHistory
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest IntendedErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"Hf8TAAD//0AbVjaIIgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
