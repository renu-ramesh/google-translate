package models

type ReqBody struct {
	SourceLang string `json:"sourceLang"`
	TargetLang string `json:"targetLang"`
	SourceText string `json:"sourceText"`
}
