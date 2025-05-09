package stringconverter

import (
	"reflect"
	"testing"
)
//Teste da função
func TestStringToBytes(t *testing.T) {
	input := ToBytes("teste")
	resultado := []byte{116, 101, 115, 116, 101}
	if !reflect.DeepEqual(input, resultado) {
		t.Error("Expected:", resultado, "Got:", input)
	}
}
