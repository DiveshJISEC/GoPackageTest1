package private

import "fmt"

func BlockFund(matchAct string, amt int) string {
	return fmt.Sprintf("API Pub1. Fund for %s is blocked by %d", matchAct, amt)
}

func UnblockFund(matchAct string, amt int) string {
	return fmt.Sprintf("API Pub2. Fund for %s is unblocked by %d", matchAct, amt)
}

func MarkLien(matchAct string, amt int) string {
	return fmt.Sprintf("API Pub3. Fund for %s is Marked as Lien by %d", matchAct, amt)
}

func UnMarkLien(matchAct string, amt int) string {
	return fmt.Sprintf("API Pub4. Fund for %s is unMarked as Lien by %d", matchAct, amt)
}
