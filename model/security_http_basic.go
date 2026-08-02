package model

// Describes HTTP Basic authentication, requiring a base64-encoded username and password.
type SecurityHttpBasic struct {
    Type string `json:"type"`
}

