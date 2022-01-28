// Kryssing 0 =								"[ kylling rev korn ---v\\ \\hs__/ ______________ /ø---]"
// Kryssing 1 = hs og kylling i båt,		"[ rev korn ---v\\ \\hs kylling/ ______________ /ø---]"
// Kryssing 2 = hs drar alene tilbake 		"[ korn rev ---v\\  ______________\\hs__/ /ø--- kylling ]"
// Kryssing 3 = hs og rev i båt 			"[ korn ---v\\ \\hs rev/ ______________ /ø---   kylling ]"
// Kryssing 4 = hs tar kylling med tilbake 	"[ korn ---v\\  ______________\\hs kylling/ /ø--- rev ]"
// Kryssing 5 = hs og korn i båt, 			"[ kylling ---v\\ \\hs korn/ ______________ /ø---   rev ]"
// Kryssing 6 = hs drar alene tilbake 		"[ kylling ---v\\  ______________ \\hs_//ø---   rev korn]"
// Kryssing 7 = hs og kylling i båt, 		"[ ---v\\  ______________ \\___//ø--- hs kylling rev korn]"

package main

import (
	"fmt"
)

func main() {
	fmt.Println(rivercrossing.InitialState())
	rivercrossing.PutInBoat("kylling")
	rivercrossing.CrossRiver(0)
	rivercrossing.ViewState()
	rivercrossing.CrossRiver(1)
	rivercrossing.ViewState()
	rivercrossing.TakeOut()
	rivercrossing.CrossRiver(1)
	rivercrossing.ViewState()

}
