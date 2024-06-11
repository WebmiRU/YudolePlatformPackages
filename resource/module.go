package resource

import "github.com/WebmiRU/YudolePlatformPackages/module"

type Module struct {
	Payload *module.Module `json:"payload"`
}

type ModuleIndex struct {
	Payload map[string]*module.Module `json:"payload"`
}
