package services

import "backend/internals/entities/payload"

type ModuleStepServices interface {
	GetModuleSteps(moduleId string) ([]payload.ModuleStepResponse, error)
}
