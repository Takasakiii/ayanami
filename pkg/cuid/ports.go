package cuid

type Generator interface {
	Generate() string
}
