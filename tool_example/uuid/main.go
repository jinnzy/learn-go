package main

import (
	"fmt"
	"github.com/google/uuid"
)

//func GenerateUUID() (string, error) {
//	u, err := uuid.NewV4()
//	if err != nil {
//		fmt.Println(err)
//		return "", err
//	}
//	fmt.Println(u)
//	return string(u[:]), nil
//}

func main()  {
	fmt.Println(uuid.New())
}
