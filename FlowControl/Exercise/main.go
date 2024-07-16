package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%3 != 0 {
			continue
		}
		if i >= 7 {
			break
		}
		fmt.Println(i)
	}
}

// i = 1, 1 % 3 != 0, continue -> will skip everything below
// i = 2, 2 % 3 is not 0 → continue → skip fmt.Println(i)
// i = 3, 3 % 3 is 0 → check -> prints 3
// 4 % 3 is not 0 → continue → skip fmt.Println(i)
// 5 % 3 is not 0 → continue → skip fmt.Println(i)
// 6 % 3 is 0 → check i >= 7 → 6 < 7 → fmt.Println(6) → prints 6
//7 % 3 is 0 → check i >= 7 → 7 >= 7 → break → loop terminates

/*By using the break statement, the loop exits immediately when i reaches 7,
so the numbers 3 and 6 are printed, and the loop does not process further values.*/
