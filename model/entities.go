package model

import (
	"fmt"
	"unicode"
)

// Creo la estructura del resultado
// tipo, longitud y el valor
type Result struct {
	Type   string
	Length int
	Value  string
}

// Creo la estructura de la cadena
type Cadena struct {
	Value string
}

func CadenaEntrante(value string) Cadena {
	return Cadena{value}
}

func NewResult(t string, l int, v string) Result {
	return Result{t, l, v}
}

// Chequeo que la cadena sea un valor valido
// chequeo que si es numerico sea un valor entero
// Chequeo que si es texto sea una letra (mayuscula o minuscula)
func ChequeoDeCadena(tipo, valor string, largo int) bool {
	if tipo == "NN" && esValorEntero(valor) {
		return true
	}
	if tipo == "TX" && esLetra(valor) {
		return true
	}
	return false
}

// Metodo que valida o no si es un valor entero
func esValorEntero(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// Metodo que valida o no si es una letra
func esLetra(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

// Retorna el tipo de la cadena
func Tipo(cadena string) string {
	return cadena[:2]
}

// Retorna la longitud de la cadena
func Longitud(cadena string) int {
	var aux int
	if _, err := fmt.Sscanf(cadena[2:4], "%2d", &aux); err == nil {

	}
	return aux
}

// Retorna la cantidad de caracteres de la cadena
func CantidadCaracteres(cadena string) int {
	return len(cadena[4:])
}

// Retorna el valor de la cadena
func Valor(cadena string) string {
	return cadena[4:]
}
