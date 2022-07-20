package misc

import (
	"encoding/json"
	"fmt"
)

func DeepCopy(dst interface{}, src interface{}) error {
	jsonData, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("json.Marshal(src): %v", err)
	}
	err = json.Unmarshal(jsonData, dst)
	if err != nil {
		return fmt.Errorf("json.Unmarshal(jsonData, dst): %v", err)
	}
	return nil
}
