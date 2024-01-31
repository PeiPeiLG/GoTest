package main
import "fmt"
func main(){
	numbers := []int{1, 2, 3, 4, 5}


	fruits := map[string]int{"apple": 5, "banana": 7, "orange": 9}


    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
		
    for key, value := range fruits {
        fmt.Printf("Key: %s, Value: %d\n", key, value)
    }
}


