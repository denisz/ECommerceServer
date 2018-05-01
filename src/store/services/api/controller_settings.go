package api

import (
	. "store/models"
)

type ControllerSettings struct {
	Controller
}

func (p *ControllerSettings) GetSettings() (*Settings, error) {
	var settings Settings
	err := p.GetStore().
		From(NodeNamedSettings).
		Get("settings", "754-3010", &settings)

	if err != nil {
		return nil, err
	}

	return &settings, nil
}

