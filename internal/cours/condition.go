package cours

import "fmt"

func condition(name string) bool {
	a := 0
	if a > 0 {
		return true
	}

	return false
}

func ConditionavecElse() string {
	a := 0
	if a > 0 {
		return "superior"
	} else if a == 0 {
		return "equal"
	} else {
		return "inferior"
	}
}
func ConditionavecswitchCase(name string){
	switch name{
		case "toto":
		 	fmt.Println("Hello", name, "admin")
		case "tata":
			fmt.Println("Hello", name, "adminSuper")
		default: 
			fmt.Println("Hello", name, "user")
	}
}