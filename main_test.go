// main_test.go
package main_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	main "github.com/mahdiZarepoor/simple-api"
)


const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)`

var a main.App

// TestMain is a special name for function which will be run before all other
// tests . it will intialize and connect to postgres and also ensure the table 
// exists , then run the tests and after that clear the table which used to test 
func TestMain(m *testing.M) {
	// connect to db
    a.Initialize(
        os.Getenv("APP_DB_USERNAME"),
        os.Getenv("APP_DB_PASSWORD"),
        os.Getenv("APP_DB_NAME"))

	// ensure if table we wanna use exist
    ensureTableExists()

	code := m.Run()
    clearTable()
    os.Exit(code)
}

func ensureTableExists() {
    if _, err := a.DB.Exec(tableCreationQuery); err != nil {
        log.Fatal(err)
    }
}

func clearTable() {
    a.DB.Exec("DELETE FROM products")
    a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

func TestEmptyTable(t *testing.T) {
    clearTable()

    req, _ := http.NewRequest("GET", "/products", nil)
    response := executeRequest(req)

    checkResponseCode(t, http.StatusOK, response.Code)

    if body := response.Body.String(); body != "[]" {
        t.Errorf("Expected an empty array. Got %s", body)
    }
}

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

// hey hey hey , what exactly is a test ? a test is not just something 
// routine, it would be different base on you code, get it ? 

// instead of testing the different part of you code inside your main function
// (because the only place where you can run line by line is inside your main func)
// you can use the go test command and run test and test functions . 

