package utils

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(i interface{}) {
	bytes, _ := json.MarshalIndent(i, "", "  ")
	fmt.Println(string(bytes))
}
