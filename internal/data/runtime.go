package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

// Return the JSON-enconded valuie for the m,ovie runtime field
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

  quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
