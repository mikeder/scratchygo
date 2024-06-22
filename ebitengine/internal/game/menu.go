package game

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	fontSize      = 24
	titleFontSize = fontSize * 1.5
	smallFontSize = fontSize / 2
)

var (
	textFaceSource *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}
	textFaceSource = s
}

func StartMenu(screen *ebiten.Image) {
	title := "CRUSTY CRAWLERS"
	body := "\n\n\n\n\n\nCLICK or SPACE\n\nTO START"

	// title text
	op := &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2, 3*titleFontSize)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = titleFontSize
	op.PrimaryAlign = text.AlignCenter
	text.Draw(screen, title, &text.GoTextFace{
		Source: textFaceSource,
		Size:   titleFontSize,
	}, op)

	// body text
	op = &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2, 3*titleFontSize)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = fontSize
	op.PrimaryAlign = text.AlignCenter
	text.Draw(screen, body, &text.GoTextFace{
		Source: textFaceSource,
		Size:   fontSize,
	}, op)
}

func PlayMenu(counter, wave uint, screen *ebiten.Image) {
	title := fmt.Sprintf("WAVE: %d", wave)
	body := fmt.Sprintf("\n\n\nCRABS DISPATCHED: %d\n\n\n", counter)

	// title text
	op := &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2, 3*titleFontSize)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = titleFontSize
	op.PrimaryAlign = text.AlignCenter
	text.Draw(screen, title, &text.GoTextFace{
		Source: textFaceSource,
		Size:   titleFontSize,
	}, op)

	// body text
	op = &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2, 3*titleFontSize)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = fontSize
	op.PrimaryAlign = text.AlignCenter
	text.Draw(screen, body, &text.GoTextFace{
		Source: textFaceSource,
		Size:   fontSize,
	}, op)
}

func OverMenu(counter, wave uint, screen *ebiten.Image) {
	title := "GAME OVER!"
	body := fmt.Sprintf("\n\n\nWAVE: %d\nCRABS DISPATCHED: %d\n\n\n", wave, counter)

	// title text
	op := &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2, 3*titleFontSize)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = titleFontSize
	op.PrimaryAlign = text.AlignCenter
	text.Draw(screen, title, &text.GoTextFace{
		Source: textFaceSource,
		Size:   titleFontSize,
	}, op)

	// body text
	op = &text.DrawOptions{}
	op.GeoM.Translate(ScreenWidth/2, 3*titleFontSize)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = fontSize
	op.PrimaryAlign = text.AlignCenter
	text.Draw(screen, body, &text.GoTextFace{
		Source: textFaceSource,
		Size:   fontSize,
	}, op)
}
