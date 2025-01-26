package entities

type OuterAPIHandler struct {
	OuterAPIURL        string
	OuterAPIToken      string
	OuterAPIAuthHeader string
}

func NewOuterAPIHandler(url, token, header string) *OuterAPIHandler {
	return &OuterAPIHandler{
		OuterAPIURL:        url,
		OuterAPIToken:      token,
		OuterAPIAuthHeader: header,
	}
}
