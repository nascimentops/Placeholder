package stringconverter

//ToBytes recebe uma string e devolve uma slice de bytes
func ToBytes(x string) []byte {
	var bytes []byte = make([]byte, len(x))
	bytes = []byte(x)
	return bytes
}

//ToString recebe uma slice de bytes e retorna uma string
func ToString(x []byte) string {
	var str string = string(x)
	return str
}
