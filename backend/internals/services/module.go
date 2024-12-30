package services

import "backend/internals/entities/payload"

type ModuleServices interface {
	GetModuleInfo(moduleId string) (*payload.ModuleResponse, error)
}
