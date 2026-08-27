//go:build prod

package internal

import "embed"

//go:embed frontend-build
var EmbeddedFiles embed.FS
