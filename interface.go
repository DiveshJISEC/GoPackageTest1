package GoPackageTest1

import (
	pvt "github.com/DiveshJISEC/GoPackageTest1/internal"

	pub "github.com/DiveshJISEC/GoPackageTest1/public"
)

func CheckBalanceV1(matchAct string) string {
	return pub.CheckBalance(matchAct)
}

func IsSufficientBalcnce(matchAct string, amt int) string {
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

func GetVersion() string {
	return pub.GetVersion()
}
