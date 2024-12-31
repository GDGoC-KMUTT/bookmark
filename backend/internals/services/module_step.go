package services

import "backend/internals/entities/payload"

type ModuleStepServices interface {
	GetModuleSteps(userID uint, moduleID string) ([]payload.ModuleStep, error)
}
