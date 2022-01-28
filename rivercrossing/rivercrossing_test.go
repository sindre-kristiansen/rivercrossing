// put(item), getin(), getout(), crossriver(), takeout(item)
//Start() or setup(), reset() osv.

package rivercrossing

import (
	"fmt"
	"testing"
)

func TestPutInBoat(t *testing.T) {
	wanted := "kylling"
	PutInBoat("kylling")
	got := boat[1]
	if got != wanted {
		t.Errorf("feil, fikk %v, ønsket %v.", got, wanted)
	}
}

func TestTakeOut(t *testing.T) {
	wanted := "kylling "
	TakeOut()
	got := ost[1]
	fmt.Println(boat[1])
	if got != wanted {
		t.Errorf("feil, fikk %v, ønsket %v.", got, wanted)
	}
}

func TestViewState(t *testing.T) {
	wanted := "rev korn --v_____\\_hs__//ø-- kylling  "
	InitialState()
	PutInBoat("kylling")
	CrossRiver(0)
	ViewState()
	CrossRiver(1)
	ViewState()
	TakeOut()
	CrossRiver(1)
	ViewState()
	got := state
	if got != wanted {
		t.Errorf("feil, fikk %v, ønsket %v.", got, wanted)
	}
}
