package GoPackageTest1

import (
	pvt "goPackageTest1/internal"

	pub "goPackageTest1/public"
)

func CheckBalanceV1(matchAct string) string {
	return pub.CheckBalance(matchAct)
}

func isSufficientBalcnce(matchAct string, amt int) string {
	return pub.IsSufficientBalcnce(matchAct, amt)
}

func CheckLien(matchAct string) string {
	return pub.CheckLien(matchAct)
}

func BlockFund(matchAct string, amt int) string {
	return pvt.BlockFund(matchAct, amt)
}

func UnblockFund(matchAct string, amt int) string {
	return pvt.UnblockFund(matchAct, amt)
}

func MarkLien(matchAct string, amt int) string {
	return pvt.MarkLien(matchAct, amt)
}

func unMarkLien(matchAct string, amt int) string {
	return pvt.UnMarkLien(matchAct, amt)
}
