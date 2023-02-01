// Package mtool is just here to embed the README.md file.
package mtool

import _ "embed"

// ReadmeBytes contains the README.md file
//
//go:embed README.md
var ReadmeBytes []byte
