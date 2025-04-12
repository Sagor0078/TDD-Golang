package Reflection

import (
	"testing"
)

type User struct {
	Name  string `required:"true"`
	Email string
	Age   int `required:"true"`
}

func TestValidateRequiredFields(t *testing.T) {
	t.Run("valid user struct", func(t *testing.T) {
		user := User{Name: "Alice", Age: 25}
		err := ValidateRequiredFields(user)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("missing required field", func(t *testing.T) {
		user := User{Name: "", Age: 25}
		err := ValidateRequiredFields(user)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("non-struct input", func(t *testing.T) {
		notStruct := 123
		err := ValidateRequiredFields(notStruct)
		if err == nil {
			t.Errorf("expected error for non-struct input, got nil")
		}
	})
}
