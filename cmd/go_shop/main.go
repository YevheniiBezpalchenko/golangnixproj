package main

import (
	"project/pkg/shop_package/filesystem"
)

func main() {
	repo := filesystem.UserFileRepository{}
	//user := models.User{2, "test1@gmail.com", "Password", "17.12.2021 16:18"}
	repo.GetByEmail()
	//text := "hello"
	//file, err := os.Create("hello.txt")
	//if err != nil {
	//	fmt.Println("Unable to create file", err)
	//	os.Exit(1)
	//}
	//defer file.Close()
	//
	//file.WriteString(text)
	//fmt.Println("Done.")
}
