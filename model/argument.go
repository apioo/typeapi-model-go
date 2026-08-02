package model

import "github.com/apioo/typeschema-model-go/model"

// Describes an argument passed to an operation.
type Argument struct {
    ContentType string `json:"contentType"`
    In string `json:"in"`
    Name string `json:"name"`
    Schema *PropertyType `json:"schema"`
}

