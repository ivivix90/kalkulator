package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// *** *** POBIERANIE PARY LICZB CAŁKOWITYCH OD UŻYTKOWNIKA *** *** -- funkcja prosi użytkownika o dwie liczby, konwertuje je na int i zwraca dwie wartości userInput1 i userInput2

func PobieranieLiczb() (int64, int64) {

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

func Dodawanie(x, y int64) int64 {
	return x + y
}

func Odejmowanie(x, y int64) int64 {
	return x - y
}
func Mnozenie(x, y int64) int64 {
	return x * y

}

func Dzielenie(x, y int64) int64 {
	return x / y
}

// *** **** MENU WYBORU OPERACJI MATEMATYCZNYCH *** **** funkcja pyta usera jakie działanie chce wykonać, i zwraca odpowiadajacy działaniu string

func WyborDzialania() string {

	var wybor string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Jakie działanie chcesz wykonać? \n Dla dodawania wpisz: dod \n Dla odejmowania wpisz: od \n Dla mnożenia wpisz: mno \n Dla dzielenia wpisz: dz\n\n")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("Wpisałeś: %q\n\n", input)

	wybor = input
	return wybor

}

func Oblicz() int64 {

	x, y := PobieranieLiczb()
	var wynik int64 = 0
	wybraneDzialanie := WyborDzialania()

	if wybraneDzialanie == "dod" {
		wynik = Dodawanie(x, y)
		fmt.Printf("Wynik działania to: %v\n", wynik)

	} else if wybraneDzialanie == "od" {
		wynik = Odejmowanie(x, y)
		fmt.Printf("Wynik działania to: %v\n", wynik)

	} else if wybraneDzialanie == "mno" {
		wynik = Mnozenie(x, y)
		fmt.Printf("Wynik działania to: %v\n", wynik)

	} else if wybraneDzialanie == "dz" {
		wynik = Dzielenie(x, y)
		fmt.Printf("Wynik działania to: %v\n", wynik)
	} else {
		fmt.Printf("wpisałeś niepoprawny symbol działania.")
	}

	return wynik

}

// *** *** HISTORIA WYKONANYCH OBLICZEŃ *** ***
// pustoooooo

//*** *** FUNC MAIN *** ***
func main() {

	Oblicz()

}
