package pages

import (
	_ "embed"
)

//go:embed site.html
var Site string

//go:embed login.html
var Login string
