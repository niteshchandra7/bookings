package main

import (
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/niteshchandra7/bookings/internals/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing: test passes
	default:
		t.Errorf("type is not *chi.Mut but %T", v)
	}

}
