package model

// Describes HTTP Bearer authentication, typically using a bearer token (e.g., JWT).
type SecurityHttpBearer struct {
    Type string `json:"type"`
}

