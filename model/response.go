package model

import "github.com/apioo/typeschema-model-go/model"

// Describes the response of the operation
type Response struct {
    Code int `json:"code"`
    ContentType string `json:"contentType"`
    Schema *PropertyType `json:"schema"`
}

