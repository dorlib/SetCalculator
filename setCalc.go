package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	fmt.Println("Hello, welcome to the sets calculator")
	fmt.Println("set will be writen like that: {1,2,...,n}")

	var operation int

	fmt.Println("first choose the operation\nenter the number of the operation you want to use")
	fmt.Printf("1.union\n2.intersection\n3.complement\n4.symetricDifference\n5.IsContainedIn\n6.ContainTheElement\n")
	fmt.Scan(&operation)

	if operation == 6 {
		fmt.Println("first write the set and then the element")
		fmt.Println("please write the number you want to check like a set, for exemple {4}")
	}

	var firstSet string
	fmt.Println("enter the first set:")
	fmt.Scan(&firstSet)

	firstSetList := setHandler(firstSet)

	var secondSet string
	fmt.Println("enter the second set:")
	fmt.Scan(&secondSet)

	secondSetList := setHandler(secondSet)

	var result []float64
	var res bool

	switch {
	case operation == 1:
		result = Union(firstSetList, secondSetList)
	case operation == 2:
		result = Intersection(firstSetList, secondSetList)
	case operation == 3:
		result = Complement(firstSetList, secondSetList)
	case operation == 4:
		result = SymetricDifference(firstSetList, secondSetList)
	case operation == 5:
		res = IsContainedIn(firstSetList, secondSetList)
	case operation == 6:
		res = ContainTheElement(firstSetList, secondSetList)
	}

	final := FinalCheck(result)

	var f bool
	for i := 0; i < len(final); i++ {
		f = false
		for j := 0; j < i; j++ {
			if final[j] == final[i] {
				f = true
			}
			if f == true {
				log.Fatalln("opps, somthing bad happend, please try again.")
			}
		}
	}

	finalFinal := ReverseSet(final)

	if finalFinal == "{}" {
		finalFinal = "empty group"
	}

	if operation == 5 {
		fmt.Println(res)
	} else if operation == 6 {
		fmt.Println(res)
	} else {
		fmt.Println(finalFinal)
	}

	return
}

func setHandler(set string) []float64 {
	var sliceSet []string
	for _, set := range set {
		sliceSet = append(sliceSet, string(set))
	}

	var S []int
	for i := 0; i < len(sliceSet); i++ {
		switch num := sliceSet[i]; num {
		case "{", "}", ",":
			S = append(S, i)
		default:
			continue
		}
	}

	var sliceString []string
	var str string

	for i := 1; i < len(S); i++ {
		for k := 0; k < S[i]-S[i-1]; k++ {
			str = str + sliceSet[k+S[i-1]]
		}
		sliceString = append(sliceString, str)
		str = ""
	}

	for i := 0; i < len(sliceString); i++ {
		sliceString[i] = sliceString[i][1:]
	}
	var sliceFloat []float64

	for j := 0; j < len(sliceString); j++ {
		f, _ := strconv.ParseFloat(string(sliceString[j]), 64)
		sliceFloat = append(sliceFloat, f)
	}

	var found bool
	var sorted []float64

	for i := 0; i < len(sliceFloat); i++ {
		found = false
		for j := 0; j < i; j++ {
			if sliceFloat[i] == sorted[j] {
				found = true
				break
			}

		}
		if found == false {
			sorted = append(sorted, sliceFloat[i])
		}

	}
	return sorted
}

func Union(firstSetList, secondSetList []float64) []float64 {
	var result []float64
	var found bool
	for i := 0; i < len(firstSetList); i++ {
		found = false
		for j := 0; j < i; j++ {
			if firstSetList[i] == result[j] {
				found = true
				break
			}
		}
		if found == false {
			result = append(result, firstSetList[i])
		}
	}

	for i := 0; i < len(secondSetList); i++ {
		found = false
		for j := 0; j < i; j++ {
			if secondSetList[i] == result[j] {
				found = true
				break
			}
		}
		if found == false {
			result = append(result, secondSetList[i])

		}

	}
	return result
}

func Intersection(firstSetList, secondSetList []float64) []float64 {
	var result []float64
	for i := 0; i < len(firstSetList); i++ {
		for j := 0; j < len(secondSetList); j++ {
			if firstSetList[i] == secondSetList[j] {
				result = append(result, firstSetList[i])
			}
		}
	}
	return result
}

func Complement(firstSetList, secondSetList []float64) []float64 {
	var result []float64
	result = firstSetList
	for i := 0; i < len(firstSetList); i++ {
		for j := 0; j < len(secondSetList); j++ {
			if firstSetList[i] == secondSetList[j] {
				result = append(result[:i], result[i+1:]...)
			}
		}
	}
	return result
}

func SymetricDifference(firstSetList, secondSetList []float64) []float64 {
	firstRes := []float64{}
	secondRes := []float64{}
	finalRes := []float64{}
	firstRes = Complement(firstSetList, secondSetList)
	secondRes = Complement(secondSetList, firstSetList)
	finalRes = Union(firstRes, secondRes)
	return finalRes

}

func ReverseSet(result []float64) string {
	str := "{"
	for i := 0; i < len(result); i++ {
		if i == len(result)-1 {
			str = str + strconv.FormatFloat(result[i], 'f', -1, 64)
			break
		}
		str = str + strconv.FormatFloat(result[i], 'f', -1, 64) + ","
	}
	str += "}"
	return str
}

func FinalCheck(res []float64) []float64 {
	var foo bool
	finalAfterCheck := []float64{}
	for i := 0; i < len(res); i++ {
		foo = false
		for j := 0; j < i; j++ {
			if res[i] == res[j] {
				foo = true
				break
			}
		}
		if foo == false {
			finalAfterCheck = append(finalAfterCheck, res[i])
		}

	}
	return finalAfterCheck
}

func IsElementIn(firstSetList []float64, num []float64) bool {
	var res bool
	res = false
	for i := 0; i < len(firstSetList); i++ {
		if num[0] == firstSetList[i] {
			res = true
			break
		}
	}
	return res
}

func ContainTheElement(firstSetList, secondSetList []float64) bool {
	var found bool
	for i := 0; i < len(firstSetList); i++ {
		found = false
		for j := 0; j < len(secondSetList); j++ {
			if firstSetList[i] == secondSetList[j] {
				found = true

			}
		}
	}
	return found
}
