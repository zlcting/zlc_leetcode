package main

import "fmt"

func main() {
	J := "aA"
	S := "aAAbbbb"
	sum := numJewelsInStones(J, S)
	fmt.Printf("%v", sum)

}

//宝石与石头
func numJewelsInStones(J string, S string) int {
	var sum int

	for _, v := range J {

		for _, x := range S {
			if v == x {
				sum = sum + 1
			}
		}
	}

	return sum
}
