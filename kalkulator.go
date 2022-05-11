package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// *** *** POBIERANIE PARY LICZB CAŁKOWITYCH OD UŻYTKOWNIKA *** *** -- funkcja prosi użytkownika o dwie liczby, konwertuje je na int i zwraca dwie wartości userInput1 i userInput2

func pobieranieLiczb() (int64, int64) {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Podaj pierwszą liczbę całkowitą: \n")
	scanner.Scan()
	userInput1, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("Podaj drugą liczbę całkowitą: \n")
	scanner.Scan()
	userInput2, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	return userInput1, userInput2
}

// *** *** BAZA DOSTĘPNYCH W PROGRAMIE KALKULATOR DZIAŁAŃ ( OPERACJI MATEMATYCZNYCH) *** ***

func dodawanie(x, y int64) int64 {
	return x + y
}

func odejmowanie(x, y int) int {
	return x - y
}
func mnozenie(x, y int) int {
	return x * y

}

func dzielenie(x, y int) int {
	return x / y
}

// *** **** MENU WYBORU OPERACJI MATEMATYCZNYCH *** **** funkcja pyta usera jakie działanie chce wykonać, i zwraca odpowiadajacy działaniu string

func wyborDzialania() string {

	var wybor string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Jakie działanie chcesz wykonać? \n Dla dodawania wpisz: dod \n Dla odejmowania wpisz: od \n Dla mnożenia wpisz: mno \n Dla dzielenia wpisz: dz\n\n")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("Wpisałeś: %q\n\n", input)

	wybor = input
	return wybor

}

// *** *** HISTORIA WYKONANYCH OBLICZEŃ *** ***

func main() {

	x, y := pobieranieLiczb()
	fmt.Printf("Liczby, które wybrałes to: %v  oraz: %v\n\n", x, y)

	var wynik int64
	wybraneDzialanie := wyborDzialania()

	if wybraneDzialanie == "dod" {
		wynik = dodawanie(x, y)
		fmt.Printf("Wynik działania to: %v\n", wynik)

	} else {
		fmt.Println("Jeszcze nie znam tego działania")
	}

	//fmt.Printf("Wynik dodawania to: %v \n", dodawanie(x, y))

}
