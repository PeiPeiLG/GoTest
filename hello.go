package main // 可執行程式必須用 main 做為 package 名稱
import "fmt" // 載入 fmt 套件
func main(){ // main 函式為程式進入點
	fmt.Println("Hello, Golang!")
}

// 撰寫程式 > 建置(build) > 執行程式
// 建置: go build hello.go
// 執行: 輸入檔名(可不用輸入副檔名)