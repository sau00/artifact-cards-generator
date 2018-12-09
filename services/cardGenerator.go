package services

import (
	"encoding/json"
	"fmt"
	"github.com/h2non/filetype"
	"image"
	"image/draw"
	"image/jpeg"
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
	sourceImg, err := png.Decode(source)
	//
	if err != nil {
		fmt.Println(err)
	}

	// Algo:
	// I will generate the first title image. After that I will put this title image on source image

	// 20 symbols per title
	titleDst := getTitleImage("Heck3rman")
	// 26 symbols per string
	descriptionDst := getDescriptionImage("Promises to solve rubick's cube but doesn't do this")

	imageDst := getImageImage("image3.jpg")

	numberImage1 := getNumberImage("99")
	numberImage2 := getNumberImage("1")


	sr := image.Rectangle{image.Point{0, 0}, image.Point{sourceImg.Bounds().Dx(), sourceImg.Bounds().Dy()}}
	rgba := image.NewRGBA(sr)

	var titleOffsetX = 35
	var titleOffsetY = 15

	var maxTitleWidth = 465

	var descriptionOffsetX = 35
	var descriptionOffsetY = 500
	var maxDescriptionWidth = 465
	var maxDescriptionHeight = 230

	var imageOffsetX = 4
	var imageOffsetY = 89

	var number1ImageOffsetX = 437
	var number1ImageOffsetY = 799

	var number2ImageOffsetX = 117
	var number2ImageOffsetY = 799

	tr := image.Rectangle{image.Point{titleOffsetX, titleOffsetY}, image.Point{maxTitleWidth + titleOffsetX, titleDst.Bounds().Dy() + titleOffsetY}}
	dr := image.Rectangle{image.Point{descriptionOffsetX, descriptionOffsetY + (maxDescriptionHeight-descriptionDst.Bounds().Dy())/2}, image.Point{maxDescriptionWidth + descriptionOffsetX, descriptionDst.Bounds().Dy() + descriptionOffsetY + (maxDescriptionHeight-descriptionDst.Bounds().Dy())/2}}

	ir := image.Rectangle{image.Point{imageOffsetX, imageOffsetY}, image.Point{imageDst.Bounds().Dx() + imageOffsetX, imageDst.Bounds().Dy() + imageOffsetY}}

	n1r := image.Rectangle{image.Point{number1ImageOffsetX, number1ImageOffsetY}, image.Point{numberImage1.Bounds().Dx() + number1ImageOffsetX, number1ImageOffsetY + numberImage1.Bounds().Dy()}}
	n2r := image.Rectangle{image.Point{number2ImageOffsetX, number2ImageOffsetY}, image.Point{numberImage1.Bounds().Dx() + number2ImageOffsetX, number2ImageOffsetY + numberImage1.Bounds().Dy()}}

	draw.Draw(rgba, sourceImg.Bounds(), sourceImg, image.Point{0, 0}, draw.Over)
	draw.Draw(rgba, tr, titleDst, image.Point{0, 0}, draw.Over)
	draw.Draw(rgba, dr, descriptionDst, image.Point{0, 0}, draw.Over)
	draw.Draw(rgba, ir, imageDst, image.Point{0, 0}, draw.Over)
	draw.Draw(rgba, n1r, numberImage1, image.Point{0, 0}, draw.Over)
	draw.Draw(rgba, n2r, numberImage2, image.Point{0, 0}, draw.Over)


	out, err := os.Create("output.png")
	if err != nil {
		fmt.Println(err)
	}

	//var opt jpeg.Options
	//opt.Quality = 100

	png.Encode(out, rgba)
}

func getImageImage(filename string) *image.RGBA {
	var width = 519
	var height = 390

	imgSrc, err := ioutil.ReadFile(filename)
	imgOs, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	imageType, unknown := filetype.Match(imgSrc)

	if unknown != nil {
		fmt.Println(unknown)

		return nil
	}

	var imageData image.Image

	if filetype.IsImage(imgSrc) {

		if imageType.Extension == "jpg" {
			imageData, err = jpeg.Decode(imgOs)
		} else if imageType.Extension == "png" {
			imageData, err = png.Decode(imgOs)
		}

		if err != nil {
			fmt.Println(err)
		}
	}

	dstImage := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	ir := image.Rectangle{image.Point{0, 0}, image.Point{width, height}}

	draw.Draw(dstImage, ir, imageData, image.Point{0, 0}, draw.Over)

	return dstImage
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
	var height = len(descriptionStrings)*strHeight + len(descriptionStrings)*verticalWhitespace

	//for _, char := range description {
	//	var letter = letterPoint(string(char), "description.json")
	//	width += letter.Max.X - letter.Min.X + whitespace
	//}

	//fmt.Println(width)
	//fmt.Println(height)

	dstDescription := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	for i, str := range descriptionStrings {
		var yOffset = i*strHeight + i*verticalWhitespace
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

func getNumberImage(numberString string) *image.RGBA {
	font, err := os.Open("numbers.png")

	if err != nil {
		fmt.Println(err)
	}

	fontImg, err := png.Decode(font)

	if err != nil {
		fmt.Println(err)
	}

	var whitespace = 1

	var titleWidth = 0
	var titleHeight = 62

	for _, char := range numberString {
		var letter = letterPoint(string(char), "numbers.json")
		titleWidth += letter.Max.X - letter.Min.X + whitespace
	}

	//fmt.Println(titleWidth)
	//fmt.Println(titleHeight)

	dstTitle := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{titleWidth, titleHeight}})

	var cTitleWidth = 0

	for _, char := range numberString {
		var letter = letterPoint(string(char), "numbers.json")
		//fmt.Println(cTitleWidth)
		st := image.Rectangle{image.Point{cTitleWidth, 0}, image.Point{cTitleWidth + letter.Max.X - letter.Min.X + whitespace, titleHeight}}

		cTitleWidth = cTitleWidth + letter.Max.X - letter.Min.X + whitespace

		draw.Draw(dstTitle, st, fontImg, image.Point{letter.Min.X, 0}, draw.Over)
	}

	return dstTitle
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
