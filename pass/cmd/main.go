package main

import (
	"github.com/gptscript-ai/gptscript-credential-helpers/credentials"
	"github.com/gptscript-ai/gptscript-credential-helpers/pass"
)

func main() {
	credentials.Serve(pass.Pass{})
}
