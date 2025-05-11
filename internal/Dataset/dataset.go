package dataset

import (
	"encoding/json"
	"fmt"
)

type datasets struct {
	input  string
	target string
}

func DecodeJson(x []byte) []datasets {
	var decoded []datasets
	err := json.Unmarshal(x, &decoded)
	if err != nil {
		fmt.Println(err)
	}
	return decoded
}
