package entity

import "testing"

func TestKey_Validate(t *testing.T) {
	var test = Key{Text: "test", Key: "test123"}
	err := test.Validate()
	if err != nil {
		t.Fatal(err)
	}
}

func TestKey_Validate_error(t *testing.T) {
	tests := []struct {
		name string
		key  Key
	}{
		{
			name: "empty key",
			key:  Key{Text: "123test", Key: ""},
		},
		{
			name: "empty text",
			key:  Key{Text: "", Key: "test123"},
		},
	}
	for _, tc := range tests {
		t.Run(
			tc.name, func(t *testing.T) {
				err := tc.key.Validate()
				if err == nil {
					t.Fatal("err is nil")
				}
			},
		)
	}
}
