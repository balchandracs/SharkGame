package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type background struct {
	texture *sdl.Texture
}

func newBackGround(r *sdl.Renderer) (*background, error) {
	t, err := img.LoadTexture(r, "background/sea.jpg")

	if err != nil {
		return nil, fmt.Errorf("Couldnt load background image %v", err)
	}

	return &background{texture: t}, nil
}

func (b *background) draw(r *sdl.Renderer) error {
	r.Clear()

	if err := r.Copy(b.texture, nil, nil); err != nil {
		return fmt.Errorf("Couldnt copy bg texture%v", err)
	}
	r.Present()
	return nil
}

func (b *background) destroy() {
	b.texture.Destroy()
}
