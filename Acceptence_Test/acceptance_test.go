package Accteptance_Test

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestUserAcceptance(t *testing.T) {
    handler := setupRoutes()

    // 1. Create a new user
    userPayload := `{"name": "Alice"}`
    req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(userPayload))
    w := httptest.NewRecorder()
    handler.ServeHTTP(w, req)

    if w.Code != http.StatusCreated {
        t.Fatalf("expected status 201, got %d", w.Code)
    }

    var createdUser map[string]interface{}
    err := json.NewDecoder(w.Body).Decode(&createdUser)
    if err != nil {
        t.Fatalf("failed to decode response: %v", err)
    }

    if createdUser["name"] != "Alice" {
        t.Errorf("expected name Alice, got %v", createdUser["name"])
    }

    // 2. Retrieve all users
    req = httptest.NewRequest(http.MethodGet, "/users", nil)
    w = httptest.NewRecorder()
    handler.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Fatalf("expected status 200, got %d", w.Code)
    }

    body, _ := io.ReadAll(w.Body)
    if !bytes.Contains(body, []byte("Alice")) {
        t.Errorf("expected response to contain 'Alice', got %s", string(body))
    }
}
