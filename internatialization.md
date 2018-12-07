````
package main

import (
	"fmt"
	"image"
	"unicode/utf8"
)



// this should be in a package imagefontpack

// Type - Language Type
type Type int8

// Language Types
const (
	English Type = iota
	Russian
)

// ImageFontPack - an image font pack
type ImageFontPack struct {
	Height        int
	Language      Type
	UnknownSymbol image.Point
	Symbols       map[rune]image.Point
}

// GetImageHorizontalPoint - returns the horizontal placement of a rune in the font pack
func (p ImageFontPack) GetImageHorizontalPoint(r rune) image.Point {
	point, ok := p.Symbols[r]

	if !ok {
		return p.UnknownSymbol
	}

	return point
}

// GetImageHorizontalPoint - returns the horizontal placement of each string's rune in the font pack
func (p ImageFontPack) GetImageHorizontalPoints(s string) []image.Point {
	points := make([]image.Point, utf8.RuneCountInString(s))
	i := 0
	for _, r := range s {
		points[i] = p.GetImageHorizontalPoint(r)
		i++
	}

	return points
}

func main() {
	RussianFontPack24 := ImageFontPack{
		Height:        62,
		Language:      Russian,
		UnknownSymbol: image.Point{1, 2},
		Symbols: map[rune]image.Point{
			'A': image.Point{3, 4},
			'B': image.Point{3, 4},
			'D': image.Point{3, 4},
			'д': image.Point{5, 6},
		},
	}

	fmt.Println(RussianFontPack24.GetImageHorizontalPoint('A'))
	fmt.Println(RussianFontPack24.GetImageHorizontalPoint('д'))
	fmt.Println(RussianFontPack24.GetImageHorizontalPoint('X'))

	fmt.Println(RussianFontPack24.GetImageHorizontalPoints("дроф ABC !"))
}
``