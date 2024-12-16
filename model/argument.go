package model

import "github.com/apioo/typeschema-model-go/model"

// Describes arguments of the operation
type Argument struct {
    ContentType string `json:"contentType"`
    In string `json:"in"`
    Name string `json:"name"`
    Schema *PropertyType `json:"schema"`
}

