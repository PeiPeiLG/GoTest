package main
import "fmt"
func main(){
	// 宣告用來存放輸入資料的變數
	var x int
	var y int
	fmt.Print("請輸入兩個值，用空格分開:")
	// 接收輸入資料
	fmt.Scanln(&x,&y)
	// 輸出輸入資料
	fmt.Println("兩數字相乘是:", x*y)
}