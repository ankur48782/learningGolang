package condition

import (
	"fmt"
	"time"
)

func LearnCondition() {
	/*
		comparison operators
		Less than <
		Less than or equal <=
		Greater than >
		Greater than or equal >=
		Equal to ==
		Not equal to !=

		logical operators
		Logical AND &&
		Logical OR ||
		Logical NOT !
	*/
	a := 30
	b := 40
	if a == b {
		fmt.Println("a is equal than b")
	} else {
		fmt.Println("Any condition is not match.")
	}

	// else if a < b {
	// 	fmt.Println("a is less than b")
	// } else if a > b {
	// 	fmt.Println("a is greater than b")
	// }

	// else if a <= b {
	// 	fmt.Println("a is less than or equal to b")
	// }
	// else if a <= b {
	// 	fmt.Println("a is less than or equal to b")
	// }
	// else if a != b {
	// 	fmt.Println("a is not equal to b")
	// }
}

func SwitchCondition() {
	day := 7

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}

}

func ForLoop() {
	// for statement1; statement2; statement3 {
	// 	code to be executed for each iteration
	// }

	// i:=0; - Initialize the loop counter (i), and set the start value to 0
	// i < 5; - Continue the loop as long as i is less than 5
	// i++ - Increase the loop counter value by 1 for each iteration
	t := time.Second

	for i := 0; i <= 10; i-- {
		// fmt.Println("before Sleep()")
		if i == -10000 {
			break
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println(t)
}
