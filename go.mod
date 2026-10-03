module github.com/julianorchard/website

go 1.26.4

require (
	github.com/yuin/goldmark v1.8.2
	golang.org/x/net v0.57.0
	gopkg.in/yaml.v2 v2.4.0
)

require github.com/mangoumbrella/goldmark-figure v1.4.0

replace github.com/mangoumbrella/goldmark-figure => github.com/julianorchard/goldmark-figure v1.5.0-altcaption.0
