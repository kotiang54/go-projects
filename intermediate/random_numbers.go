package intermediate

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	unixTime := time.Now().Unix()
	val := rand.New(rand.NewSource(unixTime))

	// use Go's automatic seeding
	fmt.Println("Auto-seeding:", rand.Intn(100)+1)
	fmt.Println(rand.Float64()) // between 0.0 and 1.0

	// fixed the seed for the random number generation
	fmt.Println("Fixed seeding:", val.Intn(100)+1)
}
