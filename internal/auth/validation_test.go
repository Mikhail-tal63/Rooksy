package auth

import "testing"

func TestNormlizeUsername(t *testing.T) {

	tests := []struct {
		Name     string
		Input    string
		Expected string
	}{{
		Name:     "lowercase username",
		Input:    "mikhail",
		Expected: "mikhail",
	},
		{
			Name:     "uppercase username",
			Input:    "MIKHAIL",
			Expected: "mikhail",
		},
		{
			Name:     "mixed case username",
			Input:    "miKhaIl",
			Expected: "mikhail",
		},
		{
			Name:     "spaces username",
			Input:    "       mikhail      ",
			Expected: "mikhail",
		},
		{
			Name:     "empty username",
			Input:    "",
			Expected: "",
		},
		{
			Name:     "whitespace only",
			Input:    "           ",
			Expected: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := UsernameValidate(tt.Input)

			if got != tt.Expected {
				t.Errorf(
					"NormalizeUsername(%q) = %q, want %q",
					tt.Input,
					got,
					tt.Expected,
				)
			}
		})
	}

}

func TestNormlizeEmail(t *testing.T) {
	tests := []struct {
		Name     string
		Input    string
		Expected string
	}{{
		Name:     "lowercase email",
		Input:    "idk@gmail.com",
		Expected: "idk@gmail.com",
	},
		{
			Name:     "uppercase email",
			Input:    "IDK@GMAIL.COM",
			Expected: "idk@gmail.com",
		},
		{
			Name:     "mixed case email",
			Input:    "IdK@GmAiL.CoM",
			Expected: "idk@gmail.com",
		},
		{
			Name:     "spaces email",
			Input:    "       idk@gmail.com      ",
			Expected: "idk@gmail.com",
		},
		{
			Name:     "empty email",
			Input:    "",
			Expected: "",
		},
		{
			Name:     "whitespace only",
			Input:    "           ",
			Expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			goodEmail := EmailValidate(tt.Input)
			if tt.Expected != goodEmail {
				t.Errorf(
					"email normalize %q : expacted %q got %q",
					tt.Input,
					tt.Expected,
					goodEmail,
				)
			}
		})
	}
}
