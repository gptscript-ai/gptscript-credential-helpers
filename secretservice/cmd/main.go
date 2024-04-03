//go:build linux && cgo

package main

import (
	"github.com/gptscript-ai/gptscript-credential-helpers/credentials"
	"github.com/gptscript-ai/gptscript-credential-helpers/secretservice"
)

func main() {
	credentials.Serve(secretservice.Secretservice{})
}
