package timer

// This is sample from Rob Pike preso
import (
	"fmt"
	"math/rand"
	"time"
)

func timer() {
	timerChan := make(chan time.Time)
	fmt.Println("Starting go routine at", time.Now())
	rand.Seed(time.Now().UnixNano()) // seed with current time to get rand working
	someInt := rand.Intn(10)
	deltaT := time.Duration(someInt) * time.Second
	go func() {
		time.Sleep(deltaT)
		timerChan <- time.Now()
	}() // last parenthesis are important for closure to work

	fmt.Println("Started go routine at", time.Now())
	completedAt := <-timerChan
	fmt.Println("Completed go routine at", completedAt, "waited", deltaT, someInt)

}
