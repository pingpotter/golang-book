package main

import (
	"fmt"
	"time"
)

func main() {

	volumn := 200
	start := time.Now()

	order(volumn)
	/*container := order(volumn)
	for _, cup := range container {
		fmt.Println(cup)
	}
	*/
	end := time.Now()

	start2 := time.Now()
	order2(volumn)
	/*container2 := order2(volumn)
	for _, cup := range container2 {
		fmt.Println(cup)
	}
	*/
	end2 := time.Now()

	fmt.Println("Cafe 1", end.Sub(start))
	fmt.Println("Cafe 2", end2.Sub(start2))

}

func order2(volumn int) (container []string) {

	order := make(chan string)
	bar := make(chan string)

	go reCive2(order, volumn)
	go brew2(order, bar)
	go brew2(order, bar)
	go brew2(order, bar)

	for x := range bar {
		container = append(container, serve(x))
	}

	return
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

	for i := 1; i <= volumn; i++ {
		coffee := reCive(i)
		coffee = brew(coffee)
		container = append(container, serve(coffee))
	}

	return
}

func reCive2(out chan<- string, volumn int) {
	for i := 0; i <= volumn; i++ {

		time.Sleep(5 * time.Millisecond)
		out <- fmt.Sprintf("order: %d", i)
	}
	close(out)
}

func brew2(in <-chan string, out chan<- string) {

	for x := range in {
		time.Sleep(100 * time.Millisecond)
		out <- fmt.Sprintf("%s %s", x, "espresso")
	}
	close(out)
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
