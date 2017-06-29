package task3

/*
Вывести треугольники в порядке убывания их площади.
Входные параметры : срез объектов треугольник
Выход : упорядоченный массив имён треугольников

Примечание:
• Расчёт площади треугольника должен производится по формуле Герона.
• Каждый треугольник определяется именами вершин и длинами его сторон.
• Приложение должно обрабатывать ввод чисел с плавающей точкой.
Пример определения треугольника:
{
	vertices: ‘ABC’,
	a: 10,
	b: 20,
	c:
*/

import (
	"math"
	"sort"
	"fmt"
	"errors"
)

type Params struct {
	Triangles []Triangle `json:"triangles"`
}

func Demo(params []Params) {
	for _, param := range params {
		fmt.Printf("Received triangles:%#v\r\n", param)
		if result, err := Run(param); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:\r\n", result)
		}
	}
}

// Returns error when params can't pass validation
func Validate(params Params) (err error) {
	for _, t := range params.Triangles {
		if !t.Validate() {
			return errors.New(fmt.Sprintf("Detected not valid triangle %s A:%f B:%f C:%f", t.Vertices, t.A, t.B, t.C))
		}
	}
	return nil
}

func Run(params Params) (names []string, err error) {
	if err := Validate(params); err != nil {
		return nil, err
	}
	return SortTriangles(params.Triangles), nil
}

type Triangle struct {
	Vertices string
	A        float64
	B        float64
	C        float64
}

func (t Triangle) Validate() bool {
	isPositive := t.A > 0 && t.B > 0 && t.C > 0
	isCorrectSides := (t.A < t.B+t.C) && (t.B < t.A+t.C) && (t.C < t.A+t.B)
	return isPositive && isCorrectSides
}

type TriangleWithArea struct {
	Triangle *Triangle
	Area     float64
}

type Areas []TriangleWithArea

func (s Areas) Len() int      { return len(s) }
func (s Areas) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Areas) Less(i, j int) bool {
	return s[i].Area < s[j].Area
}

func (t *Triangle) Square() float64 {
	p := (t.A + t.B + t.C) / 2
	return math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
}

func SortTriangles(triangles []Triangle) (names []string) {
	var areas Areas
	for k := range triangles {
		areas = append(areas, TriangleWithArea{
			Triangle: &triangles[k],
			Area:     triangles[k].Square(),
		})
		fmt.Printf("%#v Area: %v\r\n", triangles[k], triangles[k].Square())
	}

	sort.Sort(sort.Reverse(areas)) // sort.Reverse() according to "Вывести треугольники в порядке _убывания_ их площади"

	for _, v := range areas {
		names = append(names, v.Triangle.Vertices)
	}
	return
}
