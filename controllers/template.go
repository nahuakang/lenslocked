package controllers

import "net/http"

// Template interface decouples controllers package's dependency on views package's Template struct
type Template interface {
	Execute(w http.ResponseWriter, data interface{})
}
