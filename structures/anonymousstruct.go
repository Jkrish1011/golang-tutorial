package structures

func anonymousStructure() bool {
	anony := struct {
		make  string
		model string
	}{
		make:  "AUDI",
		model: "A4",
	}
	if anony.make == "AUDI" && anony.model == "A4" {
		return true
	} else {
		return false
	}
}
