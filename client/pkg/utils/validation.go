package utils

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type BigIntString big.Int

func (b *BigIntString) UnmarshalJSON(data []byte) error {
	if data[0] == '"' {
		var str string

		if err := json.Unmarshal(data, &str); err != nil {
			return err
		}

		value, success := new(big.Int).SetString(str, 10)

		if !success {
			return fmt.Errorf("invalid value, please use an integer value")
		}

		*b = BigIntString(*value)

		return nil
	}

	var num float64

	if err := json.Unmarshal(data, &num); err != nil {
		return err
	}

	value := new(big.Int)
	_, success := value.SetString(fmt.Sprintf("%.0f", num), 10)

	if !success {
		return fmt.Errorf("invalid value, please use an integer value")
	}

	*b = BigIntString(*value)

	return nil
}
