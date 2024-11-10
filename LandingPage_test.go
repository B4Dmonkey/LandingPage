package landingpage

import "testing"

func TestLandingPage(t *testing.T) {
	if New() != "Landing Page" {
		t.Fatal("Expected 'Landing Page' but got something else")
	}
}