package api_test

import (
	"bytes"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"rielta/intertal/api"
	"testing"
)

func TestApi_GetNumber(t *testing.T) {
	t.Parallel()

	testCase := map[string]struct {
		res        int64
		wantStatus int
		wantRes    string
	}{
		"success": {0, http.StatusOK, "0"},
	}

	for name, tc := range testCase {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			application := NewMockApplication(ctrl)
			handlers := api.New(application)

			application.EXPECT().GetNumber().Return(tc.res)
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/number/"), nil)

			handlers.ServeHTTP(w, req)

			assert.Equal(t, w.Code, tc.wantStatus)
			assert.JSONEq(t, w.Body.String(), tc.wantRes)
		})
	}
}

func TestApi_PutNumber(t *testing.T) {
	t.Parallel()

	testCase := map[string]struct {
		reqString  string
		reqInt     int64
		res        int64
		wantStatus int
		wantRes    string
	}{
		"success":            {`{"number" : 2}`, 2, 2, http.StatusOK, "2"},
		"err":                {`{"number" : "2"}`, 2, 2, http.StatusBadRequest, `{"error":"json: cannot unmarshal string into Go struct field CreateRequest.number of type int64"}`},
		"err_negative_value": {`{"number" : -2}`, 2, 2, http.StatusBadRequest, `{"error": "Key: 'CreateRequest.Number' Error:Field validation for 'Number' failed on the 'gte' tag"}`},
	}

	for name, tc := range testCase {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			application := NewMockApplication(ctrl)
			handlers := api.New(application)

			if tc.wantStatus == http.StatusOK {
				application.EXPECT().PutNumber(tc.reqInt).Return(tc.res)
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/number/"), bytes.NewBufferString(tc.reqString))

			handlers.ServeHTTP(w, req)

			assert.Equal(t, w.Code, tc.wantStatus)
			assert.JSONEq(t, w.Body.String(), tc.wantRes)
		})
	}
}
