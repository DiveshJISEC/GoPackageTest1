package public

import (
	"fmt"
)

func CheckBalance(matchAct string) string {
	return fmt.Sprintf("API Pvt1. Requested balance for %s is 5105.14", matchAct)
}

func IsSufficientBalcnce(matchAct string, amt int) string {

	s := "not sufficent"
	if amt > 10000 {
		s = "sufficent"
	}

	return fmt.Sprintf("API Pvt2. Balance for %s for amount %d is %s", matchAct, amt, s)
}

func CheckLien(matchAct string) string {
	return fmt.Sprintf("API Pvt3. Requested lien for %s is 8453.010", matchAct)
}
