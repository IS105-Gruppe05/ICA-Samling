package ascii

import ("testing"
	"fmt"
)

func TestGreetingASCII(t *testing.T) {
	//ascii := ASCII // når denne er aktiv funker den
	ascii := ASCIIGreeting //Når denne er aktiv får du feilmelling
	if !(isASCII(ascii)) {

		t.Fail()

	}
}

func isASCII(s string) bool {

	for _, c := range s {
		if c > 127 {
			fmt.Println("Inneholder ikke bare ASCII tegn\n")
			return false
		}

	}
	return true
}
