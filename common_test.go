package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []article

// function is used for setip before executing the test function
func TestMain(m *testing.M) {
	// set gin to test mode
	gin.SetMode(gin.TestMode)

	// run the other tests
	os.Exit(m.Run())
}

func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()

	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}

	return r
}

// helper function to process a request and test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// create a response recorder
	w := httptest.NewRecorder()

	// create the service and process the above request
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// this function is used to store the main lists into temporary one for testing
func saveLists() {
	tmpArticleList = articleList
}

// restore the lists
func restoreLists() {
	articleList = tmpArticleList
}
