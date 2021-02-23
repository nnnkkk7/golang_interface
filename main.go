package main

import (
	"fmt"
)

//インターフェースの実装
type shape interface {
	getArea() float64
}

func outputArea(s shape) {
	fmt.Println(s.getArea())
}

// 四角形
type rectangle struct {
	tate float64
	yoko float64
}

// 三角形
type triangle struct {
	takasa float64
	teihen float64
}

func (r *rectangle) getArea() float64 {
	return r.tate * r.yoko
}

func (t *triangle) getArea() float64 {
	return t.teihen * t.takasa / 2
}

func main() {
	t := &rectangle{tate: 2, yoko: 5}
	r := &triangle{takasa: 9, teihen: 4}

	//四角形と三角形の面積
	outputArea(t)
	outputArea(r)
}
