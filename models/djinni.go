package models

type Djinni struct {
	Position string `json:"position"`
	Href     string `json:"href"`
	Date     string `json:"date_published"`
	Views    string `json:"views"`
	Reviews  string `json:"reviews"`
}
