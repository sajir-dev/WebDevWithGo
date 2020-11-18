package rect

type Rectangle struct {
	width  float32
	height float32
	edge   float32
}

func (r Rectangle) Area() float32 {
	return r.width * r.height
}

func (r Rectangle) PartialArea(partialHeight float32) float32 {
	return (partialHeight * r.width)
}

func MakeRect(ht float32, wt float32) Rectangle {
	return Rectangle{width: wt, height: ht}
}
