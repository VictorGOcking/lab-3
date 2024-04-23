package painter

import (
	"image/color"

	"golang.org/x/exp/shiny/screen"
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

// WhiteFill зафарбовує текстуру у білий колір. Може бути використана як Operation через OperationFunc(WhiteFill).
func WhiteFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.White, screen.Src)
}

// GreenFill зафарбовує текстуру у зелений колір. Може бути використана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.RGBA{G: 0xff, A: 0xff}, screen.Src)
}

// DrawBgRect малює на фоні прямокутник чорного кольору у вказаних координатах.
func DrawBgRect(t screen.Texture, coords []float64) {
	// Малювання чорного прямокутника в координатах x1,y1,x2,y2
}

// DrawFigure малює нову фігуру варіанта (буква Т) з центром у вказаних координатах поверх сформованого фону.
func DrawFigure(t screen.Texture, coords []float64) {
	// Малювання букви Т в координатах x1,y1
}

// Move переміщає усі фігури, попередньо намальовані за допомогою команди figure, у вказані координати.
func Move(t screen.Texture, coords []float64) {
	// Перенесення фігур (букв Т) за координатами x1,y1
}

// Reset очищає весь поточний стан текстури (інформацію про колір фону, чорний прямокутник, усі фігури додані через команду figure). Залишає лише фон з чорним кольором.
func Reset(t screen.Texture) {
	// 1. Очищуємо інформацію про поточний стан текстури.
	// 2. Замальовуємо фон чорним кольором

	// Важливо: якщо знадобляться й інші аргументи, то у файлі parser.go треба змінити виклик
}
