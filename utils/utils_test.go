package utils

import (
	"net/http"
	"testing"
)

func TestGetIDFromRequest(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/products/1", nil)
	id, _ := GetIDFromRequest(r)
	if id != 1 {
		t.Errorf("Expect 1 get %d", id)
	}
}
