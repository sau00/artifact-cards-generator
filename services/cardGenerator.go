package services

import (
	"encoding/json"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"
)

func GenerateCard() {
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

	// 20 symbols per title
	titleDst := getTitleImage("Владимир Иванович")
	// 26 symbols per string
	descriptionDst := getDescriptionImage("I will set up the maximum length of the string for that description")

	sr := image.Rectangle{image.Point{0, 0}, image.Point{sourceImg.Bounds().Dx(), sourceImg.Bounds().Dy()}}
	rgba := image.NewRGBA(sr)

	var titleOffsetX = 35
	var titleOffsetY = 15

	var maxTitleWidth = 465

	var descriptionOffsetX = 35
	var  descriptionOffsetY = 500
	var maxDescriptionWidth = 465
	var maxDescriptionHeight = 230

	tr := image.Rectangle{image.Point{titleOffsetX, titleOffsetY}, image.Point{maxTitleWidth + titleOffsetX, titleDst.Bounds().Dy() + titleOffsetY}}
	dr := image.Rectangle{image.Point{descriptionOffsetX, descriptionOffsetY + (maxDescriptionHeight - descriptionDst.Bounds().Dy()) / 2}, image.Point{maxDescriptionWidth + descriptionOffsetX, descriptionDst.Bounds().Dy() + descriptionOffsetY + (maxDescriptionHeight - descriptionDst.Bounds().Dy()) / 2}}

	draw.Draw(rgba, sourceImg.Bounds(), sourceImg, image.Point{0, 0}, draw.Over)
	draw.Draw(rgba, tr, titleDst, image.Point{0, 0}, draw.Over)
	draw.Draw(rgba, dr, descriptionDst, image.Point{0, 0}, draw.Over)

	out, err := os.Create("output.png")
	if err != nil {
		fmt.Println(err)
	}

	//var opt jpeg.Options
	//opt.Quality = 100

	png.Encode(out, rgba)
}

func getDescriptionImage(description string) *image.RGBA {
	font, err := os.Open("description.png")

	if err != nil {
		fmt.Println(err)
	}

	fontImage, err := png.Decode(font)

	if err != nil {
		fmt.Println(err)
	}

	var maxSymbolsPerString = 27
	//var maxDescriptionWidth = 465
	//var maxDescriptionHeight = 500


	words := strings.Fields(description)

	var descriptionStrings []string

	var stringLength = 0
	var stringVal = ""

	for _, word := range words {
		if (stringLength + 1 + utf8.RuneCountInString(word)) < maxSymbolsPerString {
			stringLength += utf8.RuneCountInString(word) + 1
			stringVal += word + " "
		} else {
			descriptionStrings = append(descriptionStrings, stringVal)
			stringVal = word + " "
			stringLength = utf8.RuneCountInString(word) + 1
		}
	}

	if stringVal != " " {
		descriptionStrings = append(descriptionStrings, stringVal)
	}

	var whitespace = 1
	var verticalWhitespace = 5
	var strHeight = 42
	var width = 465
	var height = len(descriptionStrings) * strHeight + len(descriptionStrings) * verticalWhitespace


	//for _, char := range description {
	//	var letter = letterPoint(string(char), "description.json")
	//	width += letter.Max.X - letter.Min.X + whitespace
	//}

	//fmt.Println(width)
	//fmt.Println(height)

	dstDescription := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	for i, str := range descriptionStrings {
		var yOffset = i * strHeight + i * verticalWhitespace
		var cDescriptionWidth = 0

		for _, char := range str {
			var letter = letterPoint(string(char), "description.json")
			cDescriptionWidth += (letter.Max.X - letter.Min.X) + whitespace
		}

		cDescriptionWidth = (width - cDescriptionWidth) / 2

		for _, char := range str {
			var letter = letterPoint(string(char), "description.json")
			st := image.Rectangle{image.Point{cDescriptionWidth, yOffset}, image.Point{cDescriptionWidth + letter.Max.X - letter.Min.X + whitespace, strHeight + yOffset}}
			cDescriptionWidth = cDescriptionWidth + letter.Max.X - letter.Min.X + whitespace
			draw.Draw(dstDescription, st, fontImage, image.Point{letter.Min.X, 0}, draw.Over)
		}
	}

	return dstDescription
}

func getTitleImage(titleString string) *image.RGBA {
	font, err := os.Open("title.png")

	if err != nil {
		fmt.Println(err)
	}

	//sourceImg, err := png.Decode(source)
	engFontTitleImg, err := png.Decode(font)
	//
	if err != nil {
		fmt.Println(err)
	}

	var whitespace = 1

	var titleWidth = 0
	var titleHeight = 62

	for _, char := range titleString {
		var letter = letterPoint(string(char), "title.json")
		titleWidth += letter.Max.X - letter.Min.X + whitespace
	}

	//fmt.Println(titleWidth)
	//fmt.Println(titleHeight)

	dstTitle := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{titleWidth, titleHeight}})

	var cTitleWidth = 0

	for _, char := range titleString {
		var letter = letterPoint(string(char), "title.json")
		//fmt.Println(cTitleWidth)
		st := image.Rectangle{image.Point{cTitleWidth, 0}, image.Point{cTitleWidth + letter.Max.X - letter.Min.X + whitespace, titleHeight}}

		cTitleWidth = cTitleWidth + letter.Max.X - letter.Min.X + whitespace

		draw.Draw(dstTitle, st, engFontTitleImg, image.Point{letter.Min.X, 0}, draw.Over)
	}

	return dstTitle
}

func letterPoint(s, f string) image.Rectangle {
	alphabetJSONFile, err := os.Open(f)

	if err != nil {
		fmt.Println(err)
	}

	//r := []rune(s)

	defer alphabetJSONFile.Close()

	var alphabet ImageFont

	alphabetStr, _ := ioutil.ReadAll(alphabetJSONFile)

	json.Unmarshal([]byte(alphabetStr), &alphabet)

	//fmt.Println(string(r))
	//
	//fmt.Println(alphabet.Letters[string(r)])

	if s == " " {
		return image.Rectangle{
			image.Point{-10, 0},
			image.Point{0, 50},
		}
	} else {
		return image.Rectangle{
			image.Point{alphabet.Letters[s].StartPos, 0},
			image.Point{alphabet.Letters[s].EndPos, alphabet.Height},
		}
	}
}
