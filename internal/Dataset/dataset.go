package dataset

import (
	"encoding/json"
	"fmt"
)

type Datasets struct {
	Input  string
	Target string
}

func DecodeJson(x []byte) []Datasets {
	var decoded []Datasets
	err := json.Unmarshal(x, &decoded)
	if err != nil {
		fmt.Println(err)
	}
	return decoded
}
