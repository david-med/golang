package iteraciones

import (
	"fmt"
)

func Iteraciones() {
	for i := 0; i <= 10; i += 2 {
		if i == 6 {
			continue
		}
		fmt.Println(i)
		if i == 8 {
			break
		}
	}
}
