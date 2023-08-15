//go:build tools
// +build tools

package main

import (
	_ "github.com/google/wire/cmd/wire"
	_ "go.uber.org/mock/mockgen"
)
