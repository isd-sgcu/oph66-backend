package database

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type JSONB json.RawMessage

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *JSONB) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	*j = JSONB(json.RawMessage(bytes))
	return nil
}

// Value return json value, implement driver.Valuer interface
func (j JSONB) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

// I wish there's some kind of inheritance here ):
func (j JSONB) MarshalJSON() ([]byte, error) {
	return json.RawMessage(j).MarshalJSON()
}

func (j *JSONB) UnmarshalJSON(value []byte) error {
	return json.Unmarshal(value, j)
}
