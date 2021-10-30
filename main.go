package main

import (
	"EntregableGo/model"
	"errors"
	"fmt"
)

func generarResultado(cadena string) (model.Result, error) {
	if len(cadena) > 4 {
		cadena := model.CadenaEntrante(cadena)
		largo := model.Longitud(cadena.Value)
		tipo := model.Tipo(cadena.Value)
		valor := model.Valor(cadena.Value)

		if model.ChequeoDeCadena(tipo, valor, largo) == true {
			return model.NewResult(tipo, largo, valor), nil
		}
	}
	return model.Result{}, errors.New("La cadena no es valida")
}

func main() {
	nuevaCadena := "TX03SYL"
	result, err := generarResultado(nuevaCadena)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

}
