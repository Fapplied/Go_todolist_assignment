package main

import (
	"net/http"
	"testing"
)

func TestCreate(t *testing.T) {


resp, err := http.Get("localhost:8080")
if err != nil {
	t.Errorf("Server is not responding with todos")
}

defer resp.Body.Close()
	

}
