package iso

import ("testing"
	"fmt"
)

func TestGreetingASCII(t *testing.T) {
	//ascii := ASCII // når denne er aktiv funker den
	ascii := ASCIIExtended //Når denne er aktiv får du feilmelling
	if !(isASCII(ascii)) {

		t.Fail()

	}
}

func isASCII(s string) bool {

	for _, c := range s {
		if c > 127 {
			fmt.Println("Inneholder bare extended ASCII tegn\n")
			return true
		}

	}
	return false
}
