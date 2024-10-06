package variables

import "testing"

func TestVariables(t *testing.T) {
	variables()
	t.Logf("Test passed successfully!")
}

func TestShortHandVariables(t *testing.T) {
	retValue := shortHandVariables()
	if retValue != "Happy birthday! You are now 21 years old!" {
		t.Errorf("Error: Did not get the expected answer!")
	}
}

func TestSameLineDeclaration(t *testing.T) {
	sameLineDeclaration()
	t.Logf("sameLineDeclarationTest passed!")
}

func TestConversions(t *testing.T) {
	float64Variable := convertions()
	if float64Variable != 88.0 {
		t.Errorf("Error: Did not pass")
	}
	t.Logf("sameLineDeclarationTest passed!")
}
