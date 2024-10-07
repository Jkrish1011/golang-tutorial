package structures

import "testing"

func TestAnonymousStructures(t *testing.T) {
	var result bool = anonymousStructure()
	if result != true {
		t.Errorf("Error: Anonymous Make and Model is not correct")
	}
	t.Logf("Success: Anonymous Make and Model is correct")
}
