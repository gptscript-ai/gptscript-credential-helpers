//go:build darwin && cgo

package main

import (
	"github.com/gptscript-ai/gptscript-credential-helpers/credentials"
	"github.com/gptscript-ai/gptscript-credential-helpers/osxkeychain"
)

func main() {
	credentials.Serve(osxkeychain.Osxkeychain{})
}
