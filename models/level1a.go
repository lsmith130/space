package models

import (
	"log"

	"github.com/lsmith130/space/draw"
	"github.com/lsmith130/space/univ"
)

type Level1A struct {
	Body *univ.Body
	u    *univ.Universe
}

func NewLevel1A(u *univ.Universe) *Level1A {
	tex, err := draw.NewTexture("models/Material Diffuse Color.png")
	if err != nil {
		log.Fatal(err)
	}

	b, err := u.NewBody("models/introlevel1.fbx", draw.ProgramTypeStandard, []*draw.Texture{tex, tex})
	if err != nil {
		log.Fatal(err)
	}

	return &Level1A{
		Body: b,
		u:    u,
	}
}

func (l *Level1A) Destroy() {
	l.u.RemoveBody(l.Body)
}