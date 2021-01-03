package main

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/ttf"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}

func run() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("Could not initialize SDL: %v", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		return fmt.Errorf("Could not initialize TTF: %v", err)
	}
	defer ttf.Quit()

	w, r, err := sdl.CreateWindowAndRenderer(1200, 800, sdl.WINDOW_SHOWN)

	if err != nil {
		return fmt.Errorf("Cannot create window %v ", err)
	}
	defer w.Destroy()

	if err := drawTitle(r); err != nil {
		return fmt.Errorf("Couldnt draw %v", err)
	}

	time.Sleep(5 * time.Second)

	if err := drawBackground(r); err != nil {
		return fmt.Errorf("Couldnt draw background %v", err)
	}

	time.Sleep(5 * time.Second)

	return nil
}

func drawTitle(r *sdl.Renderer) error {
	r.Clear()
	f, err := ttf.OpenFont("Fonts/beech/JANGKIDS.ttf", 80)

	if err != nil {
		return fmt.Errorf("Couldnt load font %v", err)
	}
	defer f.Close()
	c := sdl.Color{R: 10, G: 100, B: 255, A: 225}

	s, err := f.RenderUTF8Solid("THE GREAT WHITE", c)

	if err != nil {
		return fmt.Errorf("Error rendering %v", err)
	}

	defer s.Free()

	t, err := r.CreateTextureFromSurface(s)

	if err != nil {
		return fmt.Errorf("Error creating texture%v", err)
	}

	defer t.Destroy()

	if err := r.Copy(t, nil, nil); err != nil {
		return fmt.Errorf("Couldnt copy texture%v", err)
	}

	r.Present()

	return nil
}

func drawBackground(r *sdl.Renderer) error {
	r.Clear()

	t, err := img.LoadTexture(r, "background/sea.jpg")

	if err != nil {
		return fmt.Errorf("Couldnt load background image %v", err)
	}
	defer t.Destroy()

	if err := r.Copy(t, nil, nil); err != nil {
		return fmt.Errorf("Couldnt copy bg texture%v", err)
	}

	r.Present()
	return nil
}
