package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"

	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

const padding = 8
const fontSize = 10

type Data struct {
	Label string
	Value string

	LabelBoxWidth   int
	LabelTextOffset int
	ValueBoxWidth   int
	ValueBoxOffset  int
	ValueTextOffset int
	BadgeWidth      int

	CSS        template.CSS
}

func textWidth(face font.Face, s string) int {
	var adv int
	for _, r := range s {
		a, _ := face.GlyphAdvance(r)
		adv  += int(a >> 6)
	}
	return adv
}

func main() {
	if len(os.Args) < 5 {
		fmt.Println("usage: badgegen <label> <value> <theme.css> <template.svg>")
		os.Exit(1)
 	}

	label 			:= os.Args[1]
	value 			:= os.Args[2]
	themePath 		:= os.Args[3]
	templatePath 	:= os.Args[4]

	font, _ := opentype.Parse(goregular.TTF)
	face, err := opentype.NewFace(font, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72,
	})
	if err != nil {
		panic(err)
	}

	labelW := textWidth(face, label) + padding*2
	valueW := textWidth(face, value) + padding*2

	data := Data{
		Label: label,
		Value: value,

		LabelBoxWidth:      labelW,
		LabelTextOffset:    padding/2,
		ValueBoxWidth:      valueW,
		ValueBoxOffset:     labelW,
		ValueTextOffset:    labelW + padding/2,
		BadgeWidth:         labelW + valueW,

		CSS: template.CSS(mustRead(themePath)),
	}

	tmpl := template.Must(template.ParseFiles(templatePath))

	var out bytes.Buffer
	_ = tmpl.Execute(&out, data)

	fmt.Print(out.String())
}

func mustRead(p string) string {
	b, err := os.ReadFile(p)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}
