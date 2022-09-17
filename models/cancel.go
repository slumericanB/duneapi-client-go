package models

import (
	"errors"
)

type CancelResponse struct {
	Success bool `json:"success"`
}

func (r CancelResponse) HasError() error {
	if !r.Success {
		return errors.New("Failed to cancel query execution")
	}

	return nil
}
