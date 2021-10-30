package main

import (
	"EntregableGo/model"

	"fmt"
)

func main() {
	nuevaCadena := "TX03SYL"
	result, err := model.ObtenerResultado(nuevaCadena)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
