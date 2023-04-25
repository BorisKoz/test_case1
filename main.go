package main

import (
	"github.com/Knetic/govaluate"
	"math"
	"strconv"
)

var LEN = 9

/*
Технически - задача найти все значения для перестановок
9x8x7x6x5x4x3x2x1x0, где x - некий знак или ничего.
Очевидно, что всего таких - 3^9. Поскольку каждое значение
уравнения уникально, создадим map: вычисление - результат,
в который запишем посчитанное значение. Для вычисления используем
модуль github.com/Knetic/govaluate, альтернативно -
возможно написать простой парсер, запоминающий последний знак,
и считывающий число до следующего знака.
 */

func FindPermutations() map[string]int {
	all := make(map[string]int)
	for i := 0; i < int(math.Pow(3, 9)); i++ {
		operators := OperatorsForEquation(i)
		key := ""
		for j := 0; j < LEN; j++ {
			switch operators[j] {
			case '+':
				key += strconv.Itoa(9 - j) + string(operators[j])
			case '-':
				key += strconv.Itoa(9 - j) + string(operators[j])
			case ' ':
				key += strconv.Itoa(9 - j)
			}
		}
		key += "0"
		expression, _ := govaluate.NewEvaluableExpression(key)
		val, _ := expression.Evaluate(nil)
		res := val.(float64)
		all[key] = int(res)
	}
	return all
}

func OperatorsForEquation(val int) string {
	res := ""
	for i := 0; len(res) < LEN; i++ {
		switch val % 3 {
		case 0:
			res += "+"
		case 1:
			res += "-"
		case 2:
			res += " "
		}
		val = val / 3
	}
	return res
}


func main() {
	result := 200
	perms := FindPermutations()
	// печать результатов
	println("найдены выражения:")
	for key, val := range perms {
		if val == result {
			println(key + " = " + strconv.Itoa(val))
		}
	}
}
