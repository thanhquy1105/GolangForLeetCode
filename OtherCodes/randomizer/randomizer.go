package randomizer

import (
	"fmt"
	"math/rand"
	"time"
)

func Main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	guess := 10

	for i := 0; i != guess; {
		i = r.Intn(guess + 1)
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
