package model

import "github.com/apioo/typeschema-model-go/model"

// The root specification object of TypeAPI.
type TypeAPI struct {
    Definitions map[string]DefinitionType `json:"definitions"`
    Import map[string]string `json:"import"`
    Root string `json:"root"`
    BaseUrl string `json:"baseUrl"`
    Operations map[string]Operation `json:"operations"`
    Security *Security `json:"security"`
}

