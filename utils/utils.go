package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Return id from request
// Ex /api/books/1 -> 1
func GetIDFromRequest(r *http.Request) (int, error) {
	parts := strings.Split(r.URL.String(), "/")
	if (len(parts)) != 4 {
		return 0, errors.New("Nothing to do")
	}
	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		fmt.Println("Invalid ID. Nothing to do")
		return 0, errors.New("Invalid ID. Nothing to do")
	}
	return id, nil
}
