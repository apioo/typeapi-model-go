package model

type Argument struct {
    ContentType string `json:"contentType"`
    In string `json:"in"`
    Name string `json:"name"`
    Schema *TypeSchema.Model.PropertyType `json:"schema"`
}

