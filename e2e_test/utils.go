package e2etest

import (
	"encoding/json"
)

func StringifyData(output any) string {
	b, _ := json.Marshal(output)
	return string(b)
}
