package view_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/limpidchart/lc-api/internal/serverhttp/v0/view"
)

func TestNewError(t *testing.T) {
	t.Parallel()

	msg := "some bad thing happened"

	expected := view.Error{
		Body: struct {
			Error struct {
				Message string `json:"message"`
			} `json:"error"`
		}{
			Error: struct {
				Message string `json:"message"`
			}{
				Message: msg,
			},
		},
	}

	assert.Equal(t, expected, view.NewError(msg))
}

func TestErrorMarshalJSON(t *testing.T) {
	t.Parallel()

	expected := []byte(`{"error":{"message":"your chart is gone"}}`)

	e := view.NewError("your chart is gone")

	actual, err := e.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestNewNotFoundError(t *testing.T) {
	t.Parallel()

	id := "uuid_0"
	resource := "crab"

	expected := view.NotFoundError{
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
				Message: fmt.Sprintf("%s not found", resource),
			},
		},
	}

	assert.Equal(t, expected, view.NewNotFoundError(resource, id))
}

func TestNotFoundErrorMarshalJSON(t *testing.T) {
	t.Parallel()

	expected := []byte(`{"error":{"id":"u_u_i_d","message":"train not found"}}`)

	e := view.NewNotFoundError("train", "u_u_i_d")

	actual, err := e.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
