package model

// Describes OAuth 2.0 authentication, defining endpoints and scopes required by the API.
type SecurityOAuth struct {
    Type string `json:"type"`
    AuthorizationUrl string `json:"authorizationUrl"`
    Scopes []string `json:"scopes"`
    TokenUrl string `json:"tokenUrl"`
}

