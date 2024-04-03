//go:build windows

package main

import (
	"github.com/gptscript-ai/gptscript-credential-helpers/credentials"
	"github.com/gptscript-ai/gptscript-credential-helpers/wincred"
)

func main() {
	credentials.Serve(wincred.Wincred{})
}
