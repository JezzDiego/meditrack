package model

type NCM struct {
	Code            string  `json:"code"`
	Description     string  `json:"description"`
	FullDescription string  `json:"full_description"`
	Ex              *string `json:"ex"`
}
