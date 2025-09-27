package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Box struct {
	X      float64     // Posição X do canto superior esquerdo
	Y      float64     // Posição Y do canto superior esquerdo
	Width  float64     // Largura do retângulo
	Height float64     // Altura do retângulo
	Color  color.Color // Cor do retângulo
}

// NewBox cria um novo Box com as dimensões e cor especificadas.
func NewBox(x, y, width, height float64, clr color.Color) *Box {
	return &Box{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Color:  clr,
	}
}

// Draw desenha o Box em um ebiten.Image.
func (b *Box) Draw(img *ebiten.Image) {
	rect := ebiten.NewImage(int(b.Width), int(b.Height))
	rect.Fill(b.Color) // Preenche o retângulo com a cor especificada
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.X, b.Y) // Translada o retângulo para a posição correta
	img.DrawImage(rect, op)     // Desenha o retângulo na imagem fornecida
}

// SetPosition atualiza a posição do Box.
func (b *Box) SetPosition(x, y float64) {
	b.X = x
	b.Y = y
}

// SetSize atualiza as dimensões do Box.
func (b *Box) SetSize(width, height float64) {
	b.Width = width
	b.Height = height
}

// SetColor atualiza a cor do Box.
func (b *Box) SetColor(clr color.Color) {
	b.Color = clr
}

func (b *Box) MoveRight() {
	b.X += 1
}
