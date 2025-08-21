package adapters

import (
	"goGravity/domain"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type EbitenRenderer struct {
	screen *ebiten.Image
}

func (r *EbitenRenderer) SetScreen(screen *ebiten.Image) {
	r.screen = screen
}

func (r *EbitenRenderer) DrawBodies(bodies []*domain.Body) {
	for _, b := range bodies {
		ebitenutil.DrawRect(r.screen, b.Position.X, b.Position.Y, 5, 5, color.White)
	}
}
