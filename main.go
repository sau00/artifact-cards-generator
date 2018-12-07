package main

import (
	"acg/alphabets"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	//source, err := os.Open("source.png")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	////sourceImg, err := png.Decode(source)
	//sourceImg, err := png.Decode(source)
	////
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//// Algo:
	//// I will generate the first title image. After that I will put this title image on source image
	//
	//titleDst := getTitleImage("HELLO WORLD")
	//
	//sr := image.Rectangle{image.Point{0, 0}, image.Point{sourceImg.Bounds().Dx(), sourceImg.Bounds().Dy()}}
	//rgba := image.NewRGBA(sr)
	//
	//var titleOffsetX = 34
	//var titleOffsetY = 25
	//
	//tr := image.Rectangle{image.Point{titleOffsetX, titleOffsetY}, image.Point{titleDst.Bounds().Dx() + titleOffsetX, titleDst.Bounds().Dy() + titleOffsetY}}
	//
	//draw.Draw(rgba, sourceImg.Bounds(), sourceImg, image.Point{0, 0}, draw.Over)
	//draw.Draw(rgba, tr, titleDst, image.Point{0, 0}, draw.Over)
	//
	//out, err := os.Create("output.png")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	////var opt jpeg.Options
	////opt.Quality = 100
	//
	//png.Encode(out, rgba)

	engFontTitle, err := os.Open("eng_title.png")

	if err != nil {
		fmt.Println(err)
	}

	//sourceImg, err := png.Decode(source)
	engFontTitleImg, err := png.Decode(engFontTitle)
	//
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(GetLetterHorizontalPointsFromImage(engFontTitleImg, 1))
}

func alphabet2Cords(image image.Image) {
	alphabet, err := os.Open("alphabet.png")

	if err != nil {
		fmt.Println(err)
	}
	defer alphabet.Close()

	alphabetImg, err := png.Decode(alphabet)

	if err != nil {
		fmt.Println(err)
	}

	width := alphabetImg.Bounds().Dx()
	height := alphabetImg.Bounds().Dy()

	fmt.Println(height, " height ")
	fmt.Println(width, " width ")

	var alphabetStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	fmt.Println(string(alphabetStr[0]))
	fmt.Println("----------------------")

	var alphabetJSON alphabets.Alphabet
	alphabetJSON.Filename = "alphabet.png"
	alphabetJSON.Height = height
	
	var previousIsLineTransparent = false
	var letterCounter = 0
	for x := 0; x <= width; x++ {
		// isLineTransparent = false -> start of letter or letter; true - end of previous letter
		var isLineTransparent = true

		for y := 0; y <= height; y++ {
			_, _, _, a := alphabetImg.At(x, y).RGBA()

			if a == 0 {
				//fmt.Println("empty pixel - ", x, y, r, g, b, a)
			} else {
				//fmt.Println("filled pixel - ", x, y, r, g, b, a)
				isLineTransparent = false
			}
		}

		//if isLineTransparent {
		//	fmt.Println(x, "No letter", previousIsLineTransparent)
		//} else {
		//	fmt.Println(x, "Letter", previousIsLineTransparent)
		//}

		if previousIsLineTransparent == false && isLineTransparent == true {
			fmt.Println("Letter ", string(alphabetStr[letterCounter]), "ended", (x - 1))
			//fmt.Println((x - 1), "Letter", previousIsLineTransparent)
			letterCounter++
		} else if previousIsLineTransparent == true && isLineTransparent == false {
			fmt.Println("Letter", string(alphabetStr[letterCounter]), "started", x)
			//fmt.Println(x, "Letter", previousIsLineTransparent)
		}

		previousIsLineTransparent = isLineTransparent
	}
}

func GetLetterHorizontalPointsFromImage(img image.Image, minWidthThreshold int) []image.Point {
	bounds := img.Bounds()

	letterHorizontalPoints := []image.Point{}

	startPoint := -1
	for i := bounds.Min.X; i <= bounds.Max.X; i++ {
		emptySpace := true
		for j := bounds.Min.Y; j <= bounds.Max.Y; j++ {
			if _, _, _, a := img.At(i, j).RGBA(); a != 0 {
				if startPoint == -1 {
					startPoint = i
				}

				emptySpace = false
				break
			}
		}

		if startPoint != -1 && (emptySpace == true || i == bounds.Max.X) {
			endPoint := i
			if emptySpace == true {
				endPoint--
			}

			if endPoint-startPoint+1 > minWidthThreshold {
				letterHorizontalPoints = append(letterHorizontalPoints, image.Point{startPoint, endPoint})
			}

			startPoint = -1
		}
	}

	return letterHorizontalPoints
}

func getTitleImage(titleString string) *image.RGBA {
	engFontTitle, err := os.Open("eng_title.png")

	if err != nil {
		fmt.Println(err)
	}

	//sourceImg, err := png.Decode(source)
	engFontTitleImg, err := png.Decode(engFontTitle)
	//
	if err != nil {
		fmt.Println(err)
	}

	var whitespace = 1

	var titleWidth = 0
	var titleHeight = 52

	for _, char := range titleString {
		var letter = letterPoint(string(char))
		titleWidth += letter.Max.X - letter.Min.X + whitespace
	}

	fmt.Println(titleWidth)
	fmt.Println(titleHeight)

	dstTitle := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{titleWidth, titleHeight}})

	var cTitleWidth = 0

	for _, char := range titleString {
		var letter = letterPoint(string(char))
		fmt.Println(cTitleWidth)
		st := image.Rectangle{image.Point{cTitleWidth, 0}, image.Point{cTitleWidth + letter.Max.X - letter.Min.X + whitespace, titleHeight}}

		cTitleWidth = cTitleWidth + letter.Max.X - letter.Min.X + whitespace

		draw.Draw(dstTitle, st, engFontTitleImg, image.Point{letter.Min.X, 0}, draw.Over)
	}

	return dstTitle
}

func letterPoint(s string) image.Rectangle {
	switch s {
	case "A":
		return image.Rectangle{
			image.Point{0, 0},
			image.Point{29, 52},
		}

	case "B":
		return image.Rectangle{
			image.Point{40, 0},
			image.Point{60, 52},
		}

	case "C":
		return image.Rectangle{
			image.Point{71, 0},
			image.Point{95, 52},
		}

	case "D":
		return image.Rectangle{
			image.Point{109, 0},
			image.Point{135, 52},
		}

	case "E":
		return image.Rectangle{
			image.Point{146, 0},
			image.Point{164, 52},
		}

	case "H":
		return image.Rectangle{
			image.Point{241, 0},
			image.Point{270, 52},
		}

	case "L":
		return image.Rectangle{
			image.Point{362, 0},
			image.Point{381, 52},
		}

	case "O":
		return image.Rectangle{
			image.Point{473, 0},
			image.Point{504, 52},
		}

	case "W":
		return image.Rectangle{
			image.Point{761, 0},
			image.Point{801, 52},
		}

	case "R":
		return image.Rectangle{
			image.Point{590, 0},
			image.Point{614, 52},
		}

	// Return A letter by default
	default:
		return image.Rectangle{
			image.Point{0, 0},
			image.Point{30, 52},
		}
	}
}

//func makeTitle(s string) {
//
//}
