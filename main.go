package main

import (
	"acg/imagefontpack"
	"encoding/json"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) > 1 {

		args := os.Args[1:]

		switch args[0] {

		case "--parse":
			imagefontpack.ParseAlphabet("title.png", "title.json")

		case "--generate":
			source, err := os.Open("source.png")

			if err != nil {
				fmt.Println(err)
			}

			//sourceImg, err := png.Decode(source)
			sourceImg, err := png.Decode(source)
			//
			if err != nil {
				fmt.Println(err)
			}

			// Algo:
			// I will generate the first title image. After that I will put this title image on source image

			titleDst := getTitleImage("cykablyat!!!")

			sr := image.Rectangle{image.Point{0, 0}, image.Point{sourceImg.Bounds().Dx(), sourceImg.Bounds().Dy()}}
			rgba := image.NewRGBA(sr)

			var titleOffsetX = 34
			var titleOffsetY = 15

			tr := image.Rectangle{image.Point{titleOffsetX, titleOffsetY}, image.Point{titleDst.Bounds().Dx() + titleOffsetX, titleDst.Bounds().Dy() + titleOffsetY}}

			draw.Draw(rgba, sourceImg.Bounds(), sourceImg, image.Point{0, 0}, draw.Over)
			draw.Draw(rgba, tr, titleDst, image.Point{0, 0}, draw.Over)

			out, err := os.Create("output.png")
			if err != nil {
				fmt.Println(err)
			}

			//var opt jpeg.Options
			//opt.Quality = 100

			png.Encode(out, rgba)

		}
	} else {
		fmt.Println("Wake up Neo")
	}

	//alphabetToCords()
}
func getTitleImage(titleString string) *image.RGBA {
	engFontTitle, err := os.Open("alphabet.png")

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
	var titleHeight = 62

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

	alphabetJSONFile, err := os.Open("alphabet.json")

	if err != nil {
		fmt.Println(err)
	}

	r := []rune(s)

	defer alphabetJSONFile.Close()

	var alphabet imagefontpack.Font

	alphabetStr, _ := ioutil.ReadAll(alphabetJSONFile)

	json.Unmarshal([]byte(alphabetStr), &alphabet)

	fmt.Println(string(r))

	fmt.Println(alphabet.Letters[string(r)])

	if s == " " {
		return image.Rectangle{
			image.Point{-10, 0},
			image.Point{0, 52},
		}
	} else {
		return image.Rectangle{
			image.Point{alphabet.Letters[s].StartPos, 0},
			image.Point{alphabet.Letters[s].EndPos, alphabet.Height},
		}
	}
}
