package apiserver

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/shitikovkirill/auth-service/internal/app/apiserver/dto"
)

func TestServer_GetStatus(t *testing.T) {
	s := newServer(nil)

	testCases := []struct {
		name         string
		expectedCode int
		status       interface{}
	}{
		{
			name:         "get ok status",
			expectedCode: http.StatusOK,
			status:       "ok",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/api/status", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)

			result := rec.Result().Body
			status := &dto.StatusResponse{}
			if err := json.NewDecoder(result).Decode(status); err != nil {
				log.Fatal(err)
			}
			assert.Equal(t, status.Status, "ok")
		})
	}
}
