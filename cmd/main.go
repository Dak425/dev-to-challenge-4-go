package main

import (
	"fmt"

	"github.com/Dak425/dev-to-challenge-4-go/pkg/checkbook/memory"
)

func main() {
	raw := `1233.00
125 Hardware;! 24.8?;
123 Flowers 93.5
127 Meat 120.90
120 Picture 34.00
124 Gasoline 11.00
123 Photos;! 71.4?;
122 Picture 93.5
132 Tires;! 19.00,?;
129 Stamps 13.6
129 Fruits{} 17.6
129 Market;! 128.00?;
121 Gasoline;! 13.6?;`

	cb := memory.NewInMemoryCheckBook(raw)

	fmt.Print(cb.FullReport())
}
