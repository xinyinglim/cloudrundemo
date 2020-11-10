package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	a.Initialize() //if there is anything to initialize
	code := m.Run()
	os.Exit(code)
}

func TestHelloWorld(t *testing.T) {
	req, _ := http.NewRequest("GET", "/helloworld", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != helloworldCheckMessage {
		t.Errorf("Expected '%s', got %s", helloworldCheckMessage, body)
	}
}

//repondWithError
// func respondWithError(w http.ResponseWriter, code int, message string) {
//     respondWithJSON(w, code, map[string]string{"error": message})
// }

// func TestGetNonExistentProduct(t *testing.T) {
// 	req, _ := http.NewRequest("GET", "/helloworld/11", nil)
// 	response := executeRequest(req)

// 	checkResponseCode(t, http.StatusNotFound, response.Code)

// 	var m map[string]string
// 	json.Unmarshal(response.Body.Bytes(), &m)
// 	if m["error"] != "Product not found" {
// 		t.Errorf("Expected the 'error key of the response to be set to 'Product not found'. Got '%s'", m["error"])
// 	}
// 	//tests that status code is 404, indicating that the product was not found
// 	//response should contain an error with the message "Product not found"
// }

// func TestPOSTRequest(t *testing.T) {
// 	var jsonStr = []byte(`{"name": "test product", "price" : 11.22}`)
// 	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonStr))
// 	req.Header.Set("Content-Type", "application/json")

// 	response := executeRequest(req)
// 	checkResponseCode(t, http.StatusCreated, response.Code)

// 	var m map[string]interface{}
// 	json.Unmarshal(response.Body.Bytes(), &m)

// 	if m["name"] != "test product" {
// 		t.Errorf("Expected product name to be 'test product'. Got '%v'", m["name"])
// 	}

// 	if m["price"] != 11.22 {
//         t.Errorf("Expected product price to be '11.22'. Got '%v'", m["price"])
//     }
// // the id is compared to 1.0 because JSON unmarshaling converts numbers to
// 	// floats, when the target is a map[string]interface{}
// 	if m["id"] != 1.0 {
// 		t.Errorf("Expected product ID to be '1'. Got '%v'", m["id"])
// 	}
// }

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}

}
