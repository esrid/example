package templates

import "embed"

//go:embed **/*.html
var TPL embed.FS
