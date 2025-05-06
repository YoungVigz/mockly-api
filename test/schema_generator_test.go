package test

import (
	"os"
	"testing"

	"github.com/YoungVigz/mockly-api/internal/repository"
	"github.com/YoungVigz/mockly-api/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestGenerateFromSchema_Success(t *testing.T) {

	err := os.MkdirAll("./schemas", os.ModePerm)
	assert.NoError(t, err, "should be able to create ./schemas directory")

	f, err := os.CreateTemp("./schemas", "*.json")
	assert.NoError(t, err, "should be able to create file inside ./schemas")
	f.Close()
	os.Remove(f.Name())

	serv := services.NewSchemaService(&repository.SchemaRepository{})

	jsonData := `{
		"models": {
			"point": {
				"fields": {
					"x": { "type": "number" },
					"y": { "type": "number" }
				}
			}
		}
	}`

	data, name, err := serv.GenerateFromSchema([]byte(jsonData), 0)

	assert.NoError(t, err)
	assert.Len(t, name, 10)
	assert.Contains(t, string(data), `"point"`)
}

func TestGenerateFromSchema_NotSuccesful(t *testing.T) {

	err := os.MkdirAll("./schemas", os.ModePerm)
	assert.NoError(t, err, "should be able to create ./schemas directory")

	f, err := os.CreateTemp("./schemas", "*.json")
	assert.NoError(t, err, "should be able to create file inside ./schemas")
	f.Close()
	os.Remove(f.Name())

	serv := services.NewSchemaService(&repository.SchemaRepository{})

	jsonData := `{
		"models": {
			"point": {
			}
		}
	}`

	data, _, err := serv.GenerateFromSchema([]byte(jsonData), 0)

	assert.Error(t, err)
	assert.NotContains(t, string(data), `"point"`)
}
