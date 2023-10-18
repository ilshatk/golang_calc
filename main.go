package main

import (
	"fmt"
	"strconv"
	"strings"
)

func intToRoman(num int) string {
	numbers := map[int]string{
		1:   "I",
		4:   "IV",
		5:   "V",
		9:   "IX",
		10:  "X",
		40:  "XL",
		50:  "L",
		90:  "XC",
		100: "C",
	}
	str := ""
	//единицы-----------------------------------
	if num%10 == 4 {
		str = numbers[4] + str
		num = num - 4
	}

	if num%10 == 9 {
		str = numbers[9] + str
		num = num - 9
	}

	if 4 < num%10 && num%10 < 9 {
		str = numbers[5] + str
		num = num - 5
	}

	if num%10 < 4 {
		str = str + strings.Repeat(numbers[1], num%10)
		num = num - num%10
	}
	s1 := str
	str = ""
	//fmt.Println("Строчка после единиц", s1)
	// десятки---------------------------
	if num%100 == 40 {
		str = numbers[40] + str
		num = num - 40
	}

	if num%100 == 90 {
		str = numbers[90] + str
		num = num - 90
	}

	if 40 < num%100 && num%100 < 90 {
		str = numbers[50] + str
		num = num - 50
	}

	if num%100 < 40 {
		str = str + strings.Repeat(numbers[10], (num%100)/10)
		num = num - num%100
	}
	s2 := str
	str = ""

	//сотни-------------------------------------

	if num/100 == 1 {
		str = numbers[100]
	}

	s3 := str
	str = ""

	str = s3 + s2 + s1

	return str
}

func main() {
	romeToIntMap := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	var a, operation, b string
	var romeA, romeB bool
	var result int
	//fmt.Scan(&a, &operation, &b)
	_, err := fmt.Scanln(&a, &operation, &b)

	if err != nil {
		fmt.Println("[ERROR] Incorrect quantity numbers")
		return
	}

	numA, errA := strconv.Atoi(a)

	if errA != nil { // проверка на ошибку конвертации
		if val, ok := romeToIntMap[a]; ok { // проверка на то есть ли она в мапе
			//fmt.Printf("String %s arabic\n", a)
			numA = val
			romeA = true
		} else {
			fmt.Printf("[ERROR] String %s not rome and not arabic\n", a)
			return
		}
	}

	numB, errB := strconv.Atoi(b)
	if errB != nil { // проверка на ошибку конвертации
		if val, ok := romeToIntMap[b]; ok { // проверка на то есть ли она в мапе
			//fmt.Printf("String %s arabic \n", b)
			numB = val
			romeB = true
		} else {
			fmt.Printf("[ERROR] String %s not rome and not arabic\n", b)
			return
		}
	}

	if romeA != romeB {
		fmt.Printf("[ERROR] Nums different")
		return
	}

	if numA > 10 || numB > 10 || numA < 1 || numB < 1 { // здесь проверим лежит ли число в нужном диапозоне
		fmt.Printf("[ERROR] Num should be 1 < Num < 10 \n")
		return
	}

	switch operation {
	case "+":
		result = numA + numB
	case "-":
		result = numA - numB
	case "/":
		result = numA / numB
	case "*":
		result = numA * numB
	default:
		fmt.Println("[ERROR] Non-existent operator")
		return
	}

	if romeA {
		if result < 1 {
			fmt.Println("[ERROR] You entered Roman numerals, the result cannot be less than 1")
		} else {
			fmt.Println(intToRoman(result))
		}
	} else {
		fmt.Println(result)
	}

}
