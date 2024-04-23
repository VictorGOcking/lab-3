package ui

import (
	"image"
	"image/color"
	"log"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/imageutil"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/draw"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

type Visualizer struct {
	Title         string
	Debug         bool
	OnScreenReady func(s screen.Screen)

	w    screen.Window
	tx   chan screen.Texture
	done chan struct{}

	sz  size.Event
	pos image.Point
	bgc color.RGBA
}

func (pw *Visualizer) Main() {
	pw.tx = make(chan screen.Texture)
	pw.done = make(chan struct{})
	pw.pos.X = 400
	pw.pos.Y = 400
	driver.Main(pw.run)
}

func (pw *Visualizer) Update(t screen.Texture) {
	pw.tx <- t
}

func (pw *Visualizer) run(s screen.Screen) {
	w, err := s.NewWindow(&screen.NewWindowOptions{
		Title:  pw.Title,
		Width:  800,
		Height: 800,
	})
	if err != nil {
		log.Fatal("Failed to initialize the app window:", err)
	}
	defer func() {
		w.Release()
		close(pw.done)
	}()

	if pw.OnScreenReady != nil {
		pw.OnScreenReady(s)
	}

	pw.w = w

	events := make(chan any)
	go func() {
		for {
			e := w.NextEvent()
			if pw.Debug {
				log.Printf("new event: %v", e)
			}
			if detectTerminate(e) {
				close(events)
				break
			}
			events <- e
		}
	}()

	var t screen.Texture

	for {
		select {
		case e, ok := <-events:
			if !ok {
				return
			}
			pw.handleEvent(e, t)

		case t = <-pw.tx:
			w.Send(paint.Event{})
		}
	}
}

func detectTerminate(e any) bool {
	switch e := e.(type) {
	case lifecycle.Event:
		if e.To == lifecycle.StageDead {
			return true // Window destroy initiated.
		}
	case key.Event:
		if e.Code == key.CodeEscape {
			return true // Esc pressed.
		}
	}
	return false
}

func (pw *Visualizer) handleEvent(e any, t screen.Texture) {
	switch e := e.(type) {

	case size.Event: // Оновлення даних про розмір вікна.
		pw.sz = e

	case error:
		log.Printf("ERROR: %s", e)

	case mouse.Event:
		if e.Button != mouse.ButtonLeft {
			break
		}
		if e.Direction != mouse.DirPress {
			break
		}
		if t == nil {
			pw.pos.X = int(e.X)
			pw.pos.Y = int(e.Y)
			pw.w.Send(paint.Event{})
			//pw.drawDefaultUI()
			// DONE: Реалізувати реакцію на натискання кнопки миші.
		}

	case paint.Event:
		// Малювання контенту вікна.
		if t == nil {
			pw.drawDefaultUI()
		} else {
			// Використання текстури отриманої через виклик Update.
			pw.w.Scale(pw.sz.Bounds(), t, t.Bounds(), draw.Src, nil)
		}
		pw.w.Publish()
	}
}

func (pw *Visualizer) drawDefaultUI() {
	pw.w.Fill(pw.sz.Bounds(), color.RGBA{R: 0, G: 255, B: 0, A: 0}, draw.Src) // Фон.

	figureBody1 := image.Rectangle{
		Min: image.Point{
			X: pw.pos.X - 200,
			Y: pw.pos.Y - 200,
		},
		Max: image.Point{
			X: pw.pos.X + 200,
			Y: pw.pos.Y,
		},
	}
	figureColor1 := color.RGBA{R: 255, G: 255, B: 0, A: 0}
	pw.w.Fill(figureBody1, figureColor1, draw.Src)

	figureBody2 := image.Rectangle{
		Min: image.Point{
			X: pw.pos.X - 67,
			Y: pw.pos.Y,
		},
		Max: image.Point{
			X: pw.pos.X + 67,
			Y: pw.pos.Y + 200,
		},
	}
	figureColor2 := color.RGBA{R: 255, G: 255, B: 0, A: 0}
	pw.w.Fill(figureBody2, figureColor2, draw.Src)

	// Малювання білої рамки.
	for _, br := range imageutil.Border(pw.sz.Bounds(), 0) {
		pw.w.Fill(br, color.White, draw.Src)
	}
}
