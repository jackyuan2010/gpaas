package model

type FieldType int

const (
	Text = iota
	LongText
	Numerric
	SelectOne
	SelectMany
	Image
)
