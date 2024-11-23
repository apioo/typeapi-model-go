package model

// The TypeAPI Root
type TypeAPI struct {
    Definitions map[string]TypeSchema.Model.DefinitionType `json:"definitions"`
    Import map[string]string `json:"import"`
    Root string `json:"root"`
    BaseUrl string `json:"baseUrl"`
    Operations map[string]Operation `json:"operations"`
    Security *Security `json:"security"`
}

