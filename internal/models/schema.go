package models

import "encoding/json"

type Schema struct {
	Id      int             `json:"schema_id"`
	Title   string          `json:"title"`
	UserId  int             `json:"user_id"`
	Content json.RawMessage `json:"content"`
}

type SchemaRequest struct {
	Title   string
	Content json.RawMessage
}

type SchemaResponse struct {
	Id      int             `json:"schema_id"`
	Title   string          `json:"title"`
	Content json.RawMessage `json:"content"`
}
