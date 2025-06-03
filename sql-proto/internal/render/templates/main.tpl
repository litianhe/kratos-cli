﻿package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"

	"github.com/tx7do/go-utils/trans"
	"github.com/tx7do/kratos-bootstrap/bootstrap"
{{.ServerImportPath}}
	"{{.Project}}/pkg/service"
)

var version string

// go build -ldflags "-X main.version=x.y.z"

func newApp(
	lg log.Logger,
	re registry.Registrar,
{{.ServerFormalParameters}}
) *kratos.App {
	return bootstrap.NewApp(
		lg,
		re,
{{.ServerTransferParameters}}
	)
}

func main() {
	bootstrap.Bootstrap(initApp, trans.Ptr(service.{{.Service}}Service), trans.Ptr(version))
}
