package models

type TranslationCreateRequest struct {
	EntityType  string `json:"entity"`
	EntityID    int64  `json:"entityID"`
	Language    string `json:"language"`
	FieldName   string `json:"fieldName"`
	Translation string `json:"translation"`
}
