package main

import (
	"EntregableGo/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChequeoDeCadena(t *testing.T) {
	cadena := model.CadenaEntrante("NN03123")
	tipo := model.Tipo(cadena.Value)
	valor := model.Valor(cadena.Value)
	longitud := model.Longitud(cadena.Value)
	assert.Equal(t, model.ChequeoDeCadena(tipo, valor, longitud), true, "La cadena no es valida")
}

/*func TestResultado(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Type    string // the input type
		Value   string // the input value
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX06ABCDE", false, "", "", 0},
		{"NN04000A", false, "", "", 0},
	}

	for _, testData := range cases {
		data, err := generarResultado(testData.Input)
		error2 := (err == nil)
		assert.Equal(t, data.Type, testData.Type)
		assert.Equal(t, data.Value, testData.Value)
		assert.Equal(t, data.Length, testData.Length)
		assert.Equal(t, error2, testData.Success)

	}
}*/

func TestNewResultado(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Type    string // the input type
		Value   string // the input value
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX06ABCDE", false, "", "", 0},
		{"NN04000A", false, "", "", 0},
		{"TX01A", false, "TX", "A", 1},
	}

	for _, testData := range cases {
		data, err := generarResultado(testData.Input)
		assert.Equal(t, data.Type, testData.Type)
		assert.Equal(t, data.Value, testData.Value)
		assert.Equal(t, data.Length, testData.Length)
		assert.Equal(t, err != nil, testData.Success)

	}
}
