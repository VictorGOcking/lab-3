package painter

import (
	"golang.org/x/exp/shiny/screen"
	"image"
	"image/color"
)

// Operation змінює вхідну текстуру.
type Operation interface {
	// Do виконує зміну операції, повертаючи true, якщо текстура вважається готовою для відображення.
	Do(t screen.Texture) (ready bool)
}

// OperationList групує список операції в одну.
type OperationList []Operation

func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}

// UpdateOp операція, яка не змінює текстуру, але сигналізує, що текстуру потрібно розглядати як готову.
var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture) bool { return true }

// OperationFunc використовується для перетворення функції оновлення текстури в Operation.
type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

type textureData struct {
	Bgc     color.Color
	BRec    []float64
	Figures [][]float64
}

var tData = textureData{
	Bgc: color.Black,
}

// WhiteFill зафарбовує текстуру у білий колір. Може бути використана як Operation через OperationFunc(WhiteFill).
func WhiteFill(t screen.Texture) {
	tData.Bgc = color.White
}

// GreenFill зафарбовує текстуру у зелений колір. Може бути використана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture) {
	tData.Bgc = color.RGBA{G: 0xff}
}

// DrawBgRect малює на фоні прямокутник чорного кольору у вказаних координатах.
func DrawBgRect(t screen.Texture, coords []float64) {
	// Малювання чорного прямокутника в координатах x1,y1,x2,y2
	tData.BRec = coords
}

// DrawFigure малює нову фігуру варіанта (буква Т) з центром у вказаних координатах поверх сформованого фону.
func DrawFigure(t screen.Texture, coords []float64) {
	// Малювання букви Т в координатах x1,y1
	tData.Figures = append(tData.Figures, coords)
}

// Move переміщає усі фігури, попередньо намальовані за допомогою команди figure, у вказані координати.
func Move(t screen.Texture, coords []float64) {
	// Перенесення фігур (букв Т) за координатами x1,y1
	for _, figure := range tData.Figures {
		figure[0] = coords[0]
		figure[1] = coords[1]
	}
}

// Reset очищає весь поточний стан текстури (інформацію про колір фону, чорний прямокутник, усі фігури додані через команду figure). Залишає лише фон з чорним кольором.
func Reset(t screen.Texture) {
	// 1. Очищуємо інформацію про поточний стан текстури.
	// 2. Замальовуємо фон чорним кольором
	tData.Bgc = color.Black
	tData.BRec = tData.BRec[:0]
	tData.Figures = tData.Figures[:0]

	// Важливо: якщо знадобляться й інші аргументи, то у файлі parser.go треба змінити виклик
}

func CreateTexture(t screen.Texture) {
	t.Fill(t.Bounds(), tData.Bgc, screen.Src)

	if len(tData.BRec) > 0 {
		rectBody := image.Rectangle{
			Min: image.Point{
				X: int(tData.BRec[0] * float64(t.Size().X)),
				Y: int(tData.BRec[1] * float64(t.Size().Y)),
			},
			Max: image.Point{
				X: int(tData.BRec[2] * float64(t.Size().X)),
				Y: int(tData.BRec[3] * float64(t.Size().Y)),
			},
		}
		t.Fill(rectBody, color.Black, screen.Src)
	}

	for _, figure := range tData.Figures {
		figureBody1 := image.Rectangle{
			Min: image.Point{
				X: int(figure[0]*float64(t.Size().X)) - 200,
				Y: int(figure[1]*float64(t.Size().Y)) - 200,
			},
			Max: image.Point{
				X: int(figure[0]*float64(t.Size().X)) + 200,
				Y: int(figure[1] * float64(t.Size().Y)),
			},
		}
		figureColor1 := color.RGBA{R: 0xff, G: 0xff}
		t.Fill(figureBody1, figureColor1, screen.Src)

		figureBody2 := image.Rectangle{
			Min: image.Point{
				X: int(figure[0]*float64(t.Size().X)) - 67,
				Y: int(figure[1] * float64(t.Size().Y)),
			},
			Max: image.Point{
				X: int(figure[0]*float64(t.Size().X)) + 67,
				Y: int(figure[1]*float64(t.Size().Y)) + 200,
			},
		}
		figureColor2 := color.RGBA{R: 0xff, G: 0xff}
		t.Fill(figureBody2, figureColor2, screen.Src)
	}
}
