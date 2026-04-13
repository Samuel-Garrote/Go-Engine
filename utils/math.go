package utils

import "fmt"

func Sum(a int, b int) int {
    return a + b
}

func Divide(a int, b int) (int , error){
	if(b==0){
		return 0, fmt.Errorf("It cant divide by 0")
	}
	return a/b, nil
}