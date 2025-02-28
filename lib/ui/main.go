package main

func InitUI(e Element) Element {
	return Element{
		elementType: e.elementType,
		anchorX:     e.anchorX,
		anchorY:     e.anchorY,
		paddingX:    e.paddingX,
		paddingY:    e.paddingY,
		autoScale:   e.autoScale,
		children:    e.children,
	}
}
