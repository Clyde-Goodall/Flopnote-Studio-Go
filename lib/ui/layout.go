package main

type ElementTypeEnum int

const (
	Container = iota
	DecoratedContainer
	Button
	RoundedButton
	Text
)

func (e ElementTypeEnum) String() string {
	return [...]string{
		"container",
		"decorated_container",
		"button",
		"rounded_button",
		"text",
	}[e]
}

type Element struct {
	context            *Element
	elementType        ElementTypeEnum
	anchorX, anchorY   int
	paddingX, paddingY int
	autoScale          bool
	children           []Element
}
