package apierrors

import (
	"fmt"
	"reflect"
)

var Version string

// Feature represents a single functionality, associated to a unique name.
// Every feature describes a Request and a Response message, specified in the OCPP protocol.
type Feature interface {
	// Returns the unique name of the feature.
	GetFeatureName() string
	// Returns the type of the request message.
	GetRequestType() reflect.Type
	// Returns the type of the response message.
	GetResponseType() reflect.Type
}

type Request interface {
	Valid(uniqueId string) *Error
	GetFeatureName() string
}

type Response interface {
	Valid(uniqueId string) *Error
	GetFeatureName() string
}

const (
	NotImplemented                ErrorCode = "NotImplemented"                // Requested Action is not known by receiver.
	NotSupported                  ErrorCode = "NotSupported"                  // Requested Action is recognized but not supported by the receiver.
	InternalError                 ErrorCode = "InternalError"                 // An internal error occurred and the receiver was not able to process the requested Action successfully.
	MessageTypeNotSupported       ErrorCode = "MessageTypeNotSupported"       // A message with an Message Type Number received that is not supported by this implementation.
	ProtocolError                 ErrorCode = "ProtocolError"                 // Payload for Action is incomplete.
	SecurityError                 ErrorCode = "SecurityError"                 // During the processing of Action a security issue occurred preventing receiver from completing the Action successfully.
	FormationViolation            ErrorCode = "FormationViolation"            // Payload for Action is syntactically incorrect or not conform the PDU structure for Action.
	PropertyConstraintViolation   ErrorCode = "PropertyConstraintViolation"   // Payload is syntactically correct but at least one field contains an invalid value.
	OccurrenceConstraintViolation ErrorCode = "OccurrenceConstraintViolation" // Payload for Action is syntactically correct but at least one of the fields violates occurrence constraints.
	TypeConstraintViolation       ErrorCode = "TypeConstraintViolation"       // Payload for Action is syntactically correct but at least one of the fields violates data type constraints (e.g. “somestring”: 12).
	GenericError                  ErrorCode = "GenericError"                  // Any other error not covered by the previous ones.
)

// ErrorCode defines a common code name for an error.
type ErrorCode string

// Error wraps an OCPP error, containing an ErrorCode, a Description and the ID of the message.
type Error struct {
	Code        ErrorCode
	Description string
	MessageId   string
}

// NewError creates a new OCPP Error.
func NewError(errorCode ErrorCode, description string, messageId string) *Error {
	return &Error{Code: errorCode, Description: description, MessageId: messageId}
}

func (err *Error) Error() string {
	return fmt.Sprintf("ocpp message (%s): %v - %v", err.MessageId, err.Code, err.Description)
}

// Profile defines a specific set of features, grouped by functionality.
//
// Some vendor may want to keep the protocol as slim as possible, and only support some feature profiles.
// This can easily be achieved by only registering certain profiles, while remaining compliant with the specifications.
type Profile struct {
	Name     string
	Features map[string]Feature
}

// NewProfile creates a new profile, identified by a name and a set of features.
func NewProfile(name string, features ...Feature) *Profile {
	profile := Profile{Name: name, Features: make(map[string]Feature)}
	for _, feature := range features {
		profile.AddFeature(feature)
	}
	return &profile
}

// AddFeature adds a feature to the profile.
func (p *Profile) AddFeature(feature Feature) {
	p.Features[feature.GetFeatureName()] = feature
}

// SupportsFeature returns true if a feature matching the the passed name is registered with this profile, false otherwise.
func (p *Profile) SupportsFeature(name string) bool {
	_, ok := p.Features[name]
	return ok
}

// GetFeature retrieves a feature, identified by a unique name.
// Returns nil in case the feature is not registered with this profile.
func (p *Profile) GetFeature(name string) Feature {
	return p.Features[name]
}

// ParseRequest checks whether a feature is supported and passes the rawRequest message to the requestParser function.
// The type of the request message is passed to the requestParser function, which has to perform type assertion.
func (p *Profile) ParseRequest(featureName string, rawRequest any, requestParser func(raw any, requestType reflect.Type) (Request, error)) (Request, error) {
	feature, ok := p.Features[featureName]
	if !ok {
		return nil, fmt.Errorf("Feature %s not found", featureName)
	}
	requestType := feature.GetRequestType()
	return requestParser(rawRequest, requestType)
}

// ParseResponse checks whether a feature is supported and passes the rawResponse message to the responseParser function.
// The type of the response message is passed to the responseParser function, which has to perform type assertion.
func (p *Profile) ParseResponse(featureName string, rawResponse any, responseParser func(raw any, responseType reflect.Type) (Response, error)) (Response, error) {
	feature, ok := p.Features[featureName]
	if !ok {
		return nil, fmt.Errorf("Feature %s not found", featureName)
	}
	responseType := feature.GetResponseType()
	return responseParser(rawResponse, responseType)
}
