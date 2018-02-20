package main

import (
	"fmt"
	"log"
	"strconv"

	owm "github.com/briandowns/openweathermap"
)

func main() {

	fmt.Println(weatherCelsius(25, "Bangkok few cloud"))
	fmt.Println(weatherCelsius(34, "Tak many"))
	fmt.Println(weatherCelsius(17, "Phuket rainy"))
	fmt.Println(weatherCelsius(9, "Chiang-mai cold"))

	fmt.Println(weatherCelsius2(1234567890, "Locations"))
	fmt.Println(weatherCelsius2(1234567890.239, "Locations2"))
	fmt.Println(weatherCelsius2(-1234567890.234, "Locations2"))

	//API by open weather map

	w, err := owm.NewCurrent("C", "EN", "bfbbb8a64f577f6c416c3e012e579d28")
	if err != nil {
		log.Fatalln(err)
		return
	}

	w.CurrentByID(1609350) //Bangkok
	fmt.Println(weatherCelsius2(w.Main.Temp, fmt.Sprintf("%s,%s is %s", w.Name, w.Sys.Country, w.Weather[0].Description)))

	w.CurrentByID(1608528) //Changwat Nakhon Ratchasima
	fmt.Println(weatherCelsius2(w.Main.Temp, fmt.Sprintf("%s,%s is %s", w.Name, w.Sys.Country, w.Weather[0].Description)))

	w.CurrentByID(524901) //Moscow
	fmt.Println(weatherCelsius2(w.Main.Temp, fmt.Sprintf("%s,%s is %s", w.Name, w.Sys.Country, w.Weather[0].Description)))
}

func weatherCelsius(celsius int, desc string) string {

	var ret string

	switch celsius {
	case 25:
		ret = " _  _" + "\n" + " _||_ " + "\n" + "|_  _| c" + "\n" + desc
	case 34:
		ret = " _   " + "\n" + " _||_|" + "\n" + " _|  | c" + "\n" + desc
	case 17:
		ret = "    _" + "\n" + "  |  |" + "\n" + "  |  | c" + "\n" + desc
	case 9:
		ret = " _   " + "\n" + "|_|   " + "\n" + " _| c" + "\n" + desc
	default:
		ret = " _   " + "\n" + "|_|   " + "\n" + " |_| c" + "\n" + desc

	}
	return ret
}

func weatherCelsius2(celsius float64, desc string) string {

	strCel := strconv.FormatFloat(celsius, 'f', 2, 64)

	var str [][]string

	for _, point := range strCel {
		str = append(str, pattern(string(point)))
	}

	str = append(str, pattern("c"))

	ret := printOut(str)
	ret += desc

	return ret

}

func printOut(digit [][]string) string {

	ret := ""
	for n := 0; n < 3; n++ {
		for m := 0; m < len(digit); m++ {
			ret += digit[m][n]
		}
		ret += "\n"
	}
	return ret
}
func pattern(number string) []string {

	ret := []string{}

	switch number {
	case "-":
		ret = []string{" ", "_", " "}
	case ".":
		ret = []string{" ", " ", "."}
	case "c":
		ret = []string{"   ", "   ", " C "}
	case "1":
		ret = []string{"   ", "  |", "  |"}
	case "2":
		ret = []string{" _ ", " _|", "|_ "}
	case "3":
		ret = []string{" _ ", " _|", " _|"}
	case "4":
		ret = []string{"   ", "|_|", "  |"}
	case "5":
		ret = []string{" _ ", "|_ ", " _|"}
	case "6":
		ret = []string{" _ ", "|_ ", "|_|"}
	case "7":
		ret = []string{" _ ", "  |", "  |"}
	case "8":
		ret = []string{" _ ", "|_|", "|_|"}
	case "9":
		ret = []string{" _ ", "|_|", " _|"}
	default:
		ret = []string{" _ ", "| |", "|_|"}
	}
	return ret
}
