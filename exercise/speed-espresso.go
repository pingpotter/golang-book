package main

import (
	"fmt"
	"time"
)

func main() {

	volumn := 200
	start := time.Now()

	container := order(volumn)
	for _, cup := range container {
		fmt.Println(cup)
	}

	end := time.Now()
	fmt.Println(end.Sub(start))

}

func order(volumn int) (container []string) {

	/*for i := 1; i <= volumn; i++ {

		time.Sleep(5 * time.Millisecond)
		coffee := fmt.Sprintf("order: %d", i)

		time.Sleep(100 * time.Millisecond)

		coffee = fmt.Sprintf("%s %s", coffee, "espresso")

		time.Sleep(5 * time.Millisecond)
		container = append(container, fmt.Sprintf("%s %s", coffee, "ready :)"))
	}*/

	order := make(chan string, 1)
	bar := make(chan string, 4)
	serv := make(chan string, 1)

	for i := 1; i <= volumn; i++ {
		coffee := reCive(i)
		coffee = brew(coffee)
		container = append(container, serve(coffee))
	}

	return
}

func reCive(number int) string {
	time.Sleep(5 * time.Millisecond)
	return fmt.Sprintf("order: %d", number)
}

func brew(order string) string {
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("%s %s", order, "espresso")
}

func serve(coffee string) string {
	time.Sleep(5 * time.Millisecond)
	return fmt.Sprintf("%s %s", coffee, "ready :)")
}
