package structures

import (
	"fmt"
	"testing"
)

func TestStructWithMethod(t *testing.T) {
	rect := rectangle{
		x: 5,
		y: 10,
	}
	if rect.area() != 50 {
		t.Errorf("Error: Area calculation is not correct!")
	}
}

func TestAuthenticationInfo(t *testing.T) {
	var expectedOutput string = "Authorization: Basic %s:%s"
	authentication := authenticationInfo{
		username: "JK",
		password: "JK@123",
	}

	if authentication.getBasicAuth() != fmt.Sprintf(expectedOutput, authentication.username, authentication.password) {
		t.Errorf("Error: Test authentication failed! Check the inputs")
	}
}
