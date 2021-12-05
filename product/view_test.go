package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type response struct {
	Merchant, Product string
	Autoship_Enabled bool
	Live bool
}


func TestIsEligibleEndpoint(t *testing.T) {
	// Setup the test cases
	tests := []struct{
		name, merchant, product string
		expectedCode int
		expectedResponse response
	}{
		{"Found", "TestM", "TestP", http.StatusOK, response{"TestM", "TestP", true, true}},
		{"NotFound", "1", "", http.StatusNotFound, response{"", "", false, false}},
		{"NotFound", "1", "TestP", http.StatusNotFound, response{"", "", false, false}},
		{"NotFound", "TestM", "1", http.StatusNotFound, response{"", "", false, false}},
	}

	// Setup the test DB
	dw := DBWrapper{}
	db := dw.InitDB()
	defer db.Close()

	// Setup the gin router
	r := gin.Default()
	r.GET("/product", func(c *gin.Context) {
		IsEligibleView(db, c)
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uri := fmt.Sprintf("/product?merchant=%s&product=%s", tt.merchant, tt.product)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", uri, nil)
			r.ServeHTTP(w, req)

			if tt.expectedCode != w.Code {
				t.Errorf("Wanted %d but got %d", tt.expectedCode, w.Code)
			}

			var resp response
			json.NewDecoder(w.Body).Decode(&resp)
			if tt.expectedResponse != resp {
				t.Errorf("Wanted %+v but got %+v", tt.expectedResponse, resp)
			}
		})
	}
	dw.tearDown()
}
