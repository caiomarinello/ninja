package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// Initializes gin test context with a recorder and makes
// a request to the given path with the provided request body info.
func setupTestContextRec[T interface{}](method, path string, bodyInfo T) (*gin.Context, *httptest.ResponseRecorder, error) {
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	
	jsonData, err := json.Marshal(bodyInfo)
	if err != nil {
		return nil, nil, err
	}
	c.Request, err = http.NewRequest(method, path, bytes.NewReader(jsonData))
	if err != nil {
		return nil, nil, err
	}
	return c, rec, nil
}