package lexer

import "testing"

func TestServiceLiterals(t *testing.T) {
	t.Run("servicePrefix", func(t *testing.T) {
		want := "service"
		if got := servicePrefix; got != want {
			t.Errorf("servicePrefix = %q, want %q", got, want)
		}
	})
	t.Run("serviceTokenValue", func(t *testing.T) {
		want := "service"
		if got := serviceTokenValue; got != want {
			t.Errorf("serviceTokenValue = %q, want %q", got, want)
		}
	})
	t.Run("serviceNamePrefix", func(t *testing.T) {
		want := "name:"
		if got := serviceNamePrefix; got != want {
			t.Errorf("serviceNamePrefix = %q, want %q", got, want)
		}
	})
	t.Run("serviceTypePrefix", func(t *testing.T) {
		want := "type:"
		if got := serviceTypePrefix; got != want {
			t.Errorf("serviceTypePrefix = %q, want %q", got, want)
		}
	})
}
