package auth

import (
	"net/http"
	"testing"
)

func TestErrAuth(t *testing.T) {
	var authHeader http.Header // Empty/zero'd value
	_, err := GetAPIKey(authHeader)
	if err == nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}
}
