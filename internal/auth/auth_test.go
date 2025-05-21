package auth

import (
	"testing"
	"net/http"
)


func TestErrAuth(t *testing.T) {
	var authHeader http.Header // Empty/zero'd value
	_, err := GetAPIKey(authHeader)
	t.Fatalf("Code is supposed to fail as per boot dev instruct")
	if err == nil {
		t.Fatalf("expected: %v, got: %v",nil,err)
	}
}
