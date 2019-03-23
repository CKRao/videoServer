package utils

import "github.com/satori/go.uuid"

func NewUUID() (string, error) {
	uuids, err := uuid.NewV4()

	if err != nil {
		return "", err
	}

	return uuids.String(), nil
}
