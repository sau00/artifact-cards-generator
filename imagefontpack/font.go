package imagefontpack

import (
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"
)

type Font struct {
	Filename string            `json:"filename"`
	Height   int               `json:"height"`
	Letters  map[string]Letter `json:"letters"`
}

type Letter struct {
	StartPos int `json:"start_pos"`
	EndPos   int `json:"end_pos"`
}

// input - png file with alphabet, output - json file with letter's cords
func ParseAlphabet(i, o string) {
	alphabet, err := os.Open(i)

	if err != nil {
		fmt.Println(err)
	}
	defer alphabet.Close()

	alphabetImg, err := png.Decode(alphabet)

	if err != nil {
		fmt.Println(err)
	}

	height := alphabetImg.Bounds().Dy()

	alphabetStr := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя-+=';\"/?!\\")

	var alphabetJSON Font
	alphabetJSON.Filename = "alphabet.png"
	alphabetJSON.Height = height

	points := GetLetterHorizontalPointsFromImage(alphabetImg, 1)

	alphabetJSON.Letters = make(map[string]Letter)

	for i, rune := range alphabetStr {
		fmt.Println(string(rune))
		alphabetJSON.Letters[string(rune)] = Letter{points[i].X, points[i].Y}
	}

	alphabetJSONMarshaled, _ := json.Marshal(alphabetJSON)
	err = ioutil.WriteFile(o, alphabetJSONMarshaled, 0644)
	if err != nil {
		fmt.Println(err)
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
