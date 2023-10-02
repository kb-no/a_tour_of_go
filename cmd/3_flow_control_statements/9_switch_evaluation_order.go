// switch caseは、上から下へcaseを評価する。 caseの条件が一致すれば、そこで停止(自動的にbreak)する。

// (例えば、

// switch i {
// case 0:
// case f():
// }
// では、 i==0 であれば、 case 0 でbreakされるため f は呼び出されない。)

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default: // today + 3, today + 4, today + 5:
		fmt.Println("Too far away.")
	}
}