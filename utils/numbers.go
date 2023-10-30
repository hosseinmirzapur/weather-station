package utils

import (
	"errors"
	"strconv"
)

func StringToF64(str string) (*float64, error) {
	f, err := strconv.ParseFloat(str, 64)

	if err != nil {
		return nil, errors.New("error converting string to float64")
	}

	return &f, nil
}
