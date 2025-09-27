package adapters

import (
	"goGravity/internal/domain"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const scale = 1e-9

type EbitenRenderer struct {
	screen *ebiten.Image
}

func (r *EbitenRenderer) SetScreen(screen *ebiten.Image) {
	r.screen = screen
}

func (r *EbitenRenderer) DrawBodies(bodies []*domain.Body) {
	ebitenutil.DrawRect(r.screen, 0, 0, 10, 10, color.White) // Clear screen with black
	for _, b := range bodies {
		//ebitenutil.DrawRect(r.screen, b.Position.X, b.Position.Y, b.Size, b.Size, b.Color)
		ebitenutil.DrawCircle(r.screen, b.Position.X, b.Position.Y, b.Size, b.Color)
		println("Drawing body:", b.Name, "at position:", b.Position.X, b.Position.Y)
	}
}
