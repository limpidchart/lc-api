package view

import (
	"encoding/json"
	"fmt"
)

// Error represents error message.
//
// swagger:response error
type Error struct {
	// Error message.
	//
	// in: body
	Body struct {
		Error struct {
			// Message of the error.
			Message string `json:"message"`
		} `json:"error"`
	}
}

// NewError returns a new Error.
func NewError(message string) *Error {
	return &Error{
		Body: struct {
			Error struct {
				Message string `json:"message"`
			} `json:"error"`
		}{
			Error: struct {
				Message string `json:"message"`
			}{
				Message: message,
			},
		},
	}
}

// MarshalJSON implements the json.Marshaller interface.
func (r *Error) MarshalJSON() ([]byte, error) {
	res, err := json.Marshal(r.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal error body into JSON: %w", err)
	}

	return res, nil
}

// NotFoundError represents not found error for any resource.
//
// swagger:response notFoundError
type NotFoundError struct {
	// Error message
	//
	// in: body
	Body struct {
		Error struct {
			// Resource ID.
			//
			// swagger:strfmt uuid4
			ID string `json:"id"`

			// Message of the error.
			Message string `json:"message"`
		} `json:"error"`
	}
}

// NewNotFoundError returns a NotFoundError.
func NewNotFoundError(resourceType string, id string) *NotFoundError {
	return &NotFoundError{
		Body: struct {
			Error struct {
				ID      string `json:"id"`
				Message string `json:"message"`
			} `json:"error"`
		}{
			Error: struct {
				ID      string `json:"id"`
				Message string `json:"message"`
			}{
				ID:      id,
				Message: fmt.Sprintf("%s not found", resourceType),
			},
		},
	}
}

// MarshalJSON implements the json.Marshaller interface.
func (r *NotFoundError) MarshalJSON() ([]byte, error) {
	res, err := json.Marshal(r.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal not found error body into JSON: %w", err)
	}

	return res, nil
}
