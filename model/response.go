package model

type Response struct {
    Code int `json:"code"`
    ContentType string `json:"contentType"`
    Schema *TypeSchema.Model.PropertyType `json:"schema"`
}

