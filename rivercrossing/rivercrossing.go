// put(item), getin(), getout(), crossriver(), takeout(item)
//Start() or setup(), reset() osv.

package event

import "fmt"

var west = [3]string{"rev ", "kylling ", "korn "}
var ost = [3]string{" ", " ", " "}
var boat = [2]string{"hs", ""}

var river = "_____"

var state = ""

func InitialState() string {
	boat[0] = "hs_"
	state = west[0] + west[1] + west[2] + "--v" + "\\_" + boat[0] + boat[1] + "_//" + river + "ø--" + ost[0] + ost[1] + ost[2]
	return state
}
func ViewState() {
	fmt.Println(state)
}

func PutInBoat(item string) (string, string, string, string) {

	if boat[1] == "" {

		if item == "rev" {
			west[0] = ""
			boat[1] = "rev"

		} else if item == "kylling" {
			west[1] = ""
			boat[1] = "kylling"
		} else {
			west[2] = ""
			boat[1] = "korn"

		}
	} else {
		fmt.Println("boat full")
	}
	return boat[1], west[0], west[1], west[2]
}

func CrossRiver(i int) string {
	if i == 1 {
		state = west[0] + west[1] + west[2] + "--v" + river + "\\_" + boat[0] + boat[1] + "_//" + "ø--" + ost[0] + ost[1] + ost[2]
	} else {
		state = west[0] + west[1] + west[2] + "--v" + "\\_" + boat[0] + boat[1] + "_//" + river + "ø--" + ost[0] + ost[1] + ost[2]
	}
	return state
}

func TakeOut() (string, string) {
	ost[1] = "kylling "
	boat[1] = ""
	return ost[1], boat[1]
}

//lage de ulike items og plassere de (rev, kylling, korn og hs)
//plassere item + hs i båt (putInBoat())
// krysse elven CrossRiver() (sjekker at gjenværende item kan være i lag uten at en blir spist)
// ta ut item og plasser på østsida
// dra tilbake, CrossRiver()
// nytt item i båt PutInBoat()
// CrossRiver()
// ta ut item plasser på land (kan gjenværende item være i lag?)
// tar med en tilbake fordi matkjede er rev->kylling->korn og kylling kan ikke være med en av de andre
// tar ut item på vestsida, legger in item som var på vestsida
// CrossRiver()
//Ta ut item
// CrossRiver()
// Legger inn item som var tatt med tilbake (kylling)
// CrossRiver()
// tar ut item, alt komm over ingen ble spist
// Konditsjons, båt kan maks ha 2: hs(sjofør) og 1 item, rev og kylling kan ikke være sammen, kylling og korn kan ikke være sammen
