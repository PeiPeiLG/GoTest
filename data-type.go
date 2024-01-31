package main // 可執行程式必須用 main 做為 package 名稱
import "fmt" // 載入 fmt 套件
func main(){ // main 函式為程式進入點
	var xNumber int = 10
	var xFloat float64 = 3.14
	var xStr string = "Hello, Golang!"
	var xBool bool = true
	var xRune rune = 'A'
	fmt.Println(xNumber, xFloat, xStr, xBool, xRune)
}