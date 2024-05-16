package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// 定义常量
const Pi = 3.14159

// 定义接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 定义结构体
type Circle struct {
	Radius float64
}

// 定义结构体方法
func (c Circle) Area() float64 {
	return Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * Pi * c.Radius
}

// 定义结构体
type Rectangle struct {
	Width, Height float64
}

// 定义结构体方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 定义泛型函数
func PrintShapeInfo[T Shape](s T) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// 定义错误处理函数
func CheckPositive(value float64) error {
	if value <= 0 {
		return errors.New("value must be positive")
	}
	return nil
}

func main() {
	// 使用 if-else 语句
	if err := CheckPositive(-5); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value is positive")
	}

	// 使用 switch 语句
	shapeType := "circle"
	switch shapeType {
	case "circle":
		fmt.Println("It's a circle")
	case "rectangle":
		fmt.Println("It's a rectangle")
	default:
		fmt.Println("Unknown shape")
	}

	// 使用 for 循环
	for i := 0; i < 3; i++ {
		fmt.Println(strings.Repeat("Go ", i+1))
	}

	// 使用结构体和接口
	c := Circle{Radius: 5}
	r := Rectangle{Width: 10, Height: 5}

	PrintShapeInfo(c)
	PrintShapeInfo(r)

	// 使用 map 和 slice
	shapeMap := map[string]Shape{
		"circle":    c,
		"rectangle": r,
	}

	shapeList := []Shape{c, r}

	for name, shape := range shapeMap {
		fmt.Printf("%s: Area = %.2f, Perimeter = %.2f\n", name, shape.Area(), shape.Perimeter())
	}

	for _, shape := range shapeList {
		fmt.Printf("Shape: Area = %.2f, Perimeter = %.2f\n", shape.Area(), shape.Perimeter())
	}

	// 使用匿名函数
	square := func(x float64) float64 {
		return x * x
	}
	fmt.Println("Square of 4 is", square(4))

	// 使用闭包
	power := func(exponent float64) func(float64) float64 {
		return func(base float64) float64 {
			return math.Pow(base, exponent)
		}
	}
	cube := power(3)
	fmt.Println("Cube of 3 is", cube(3))
}
