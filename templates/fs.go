package templates

// https://pkg.go.dev/embed
import "embed"

//go:embed *
var FS embed.FS
