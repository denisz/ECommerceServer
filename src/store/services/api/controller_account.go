package api

import (
	. "store/models"
)

type ControllerAccount struct {
	Controller
}

// Тек. пользователь
func (p *ControllerAccount) Me() error {
	return nil
}
