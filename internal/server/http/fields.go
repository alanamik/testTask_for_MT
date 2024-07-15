package server

type ResponseSuccess struct {
	Cypher     string `json:"cypher"`
	StatusCode int    `json:"code"`
}

type ResponseError struct {
	Error      string `json:"error"`
	StatusCode int    `json:"code"`
}

type GetEncryptedStrIn struct {
	Str       string `json:"str"`
	Algorithm string `json:"algorithm"`
}

type GetEncryptedStrOut struct {
	EncryptedStr string `json:"encryptedStr"`
}
