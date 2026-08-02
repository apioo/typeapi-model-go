package model

// Describes API key authentication passed via a header or query parameter.
type SecurityApiKey struct {
    Type string `json:"type"`
    In string `json:"in"`
    Name string `json:"name"`
}

