package cours

import "fmt"

func boucleWhile(){
	i := 0
	for i < 5{
		fmt.Println(i)
		i++
	}
}

func boucleDoWhile(){
	i := 0
	for {
		fmt.Println(i)
		i++
		if i < 5{
			break
		}
	}
}
func boucleFor(){
	for i:= 0; i < 5; i++{
		fmt.Println(i)
	}
}
func boucleForEach(array []string){
	for index:=range array{
		fmt.Println(index)
	}
	for index, value:=range array{
		fmt.Println(index, value)
	}
	/* 
	*** si on veut récuperer juste la valeur sans l'index
	*******/
	for _, value:=range array{
		fmt.Println(value)
	}
}