/* package main

import "fmt"

func sum(a int, b int) int {
    return a + b
}

func main() {
    result := sum(5, 3)
    fmt.Println(result)
} */

/* package main

import "fmt"

func divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("It cant divide between 0)
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    fmt.Println(result, err)
} */

/* package main

import "fmt"

type Task struct {
    ID        int
    Title     string
    Completed bool
}

func main() {
    tasks := []Task{} //forEach

    tasks = append(tasks, Task{ID: 1, Title: "Learn Go", Completed: false})
    tasks = append(tasks, Task{ID: 2, Title: "Build API", Completed: false})
    tasks = append(tasks, Task{ID: 3, Title: "Use Redis", Completed: false})

    for _, task := range tasks {
        fmt.Println(task.ID, task.Title, task.Completed)
    }
} */
/*
package main

import "fmt"

type Task struct {
    ID        int
    Title     string
    Completed bool
}

func main(){
	tasks:= map[int]Task{} //map
		tasks[1] =Task{ID: 1, Title: "Learning Go", Completed: false}
		tasks[2] = Task{ID: 2, Title: "Learning Go2", Completed: false}
		tasks[3] = Task{ID: 3, Title: "Learning Go3", Completed: false}

	fmt.Println(tasks[1])
	fmt.Println(tasks[2])
} */
/*
package main

import "fmt"

type Task struct {
    ID        int
    Title     string
    Completed bool
}

func (t*Task) Complete(){
	t.Completed = true
}

func main(){
	task := Task{ID: 1, Title: "Hi Go", Completed:false}

	fmt.Println(task.Completed)

	task.Complete()
	fmt.Println(task)
}
*/

package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func rest(a int, b int) (int, error) {
	if a <= 0 {
		return 0, fmt.Errorf("first number must be greater than 0")
	}
	return a - b, nil
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func mult(a int, b int) int {
	return a * b
}

func mainA() {
	//sum
	fmt.Println("Sum:", sum(2, 3))
//Rest
	result, err := rest(5, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Rest:", result)
	}
//Divide
	result2, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Divide:", result2)
	}

	fmt.Println("Mult:", mult(4, 5))
}