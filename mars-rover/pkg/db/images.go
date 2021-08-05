package db

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type images []string

// implementing Value() sql driver function
func (i images) Value() (driver.Value, error) {
	if len(i) == 0 {
		return "[]", nil
	}

	return fmt.Sprintf(`["%s"]`, strings.Join(i, `","`)), nil
}

// implementing Scan() sql driver function
func (i *images) Scan(src interface{}) (err error) {
	var images []string
	switch src.(type) {
	case string:
		err = json.Unmarshal([]byte(src.(string)), &images)
	case []byte:
		err = json.Unmarshal(src.([]byte), &images)
	default:
		return errors.New("Incompatible type for images")
	}
	if err != nil {
		return
	}
	*i = images
	return nil
}
