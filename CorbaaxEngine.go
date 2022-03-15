/**
	* Author 		: Corbetta Luca
	* Project 		: CorbaaxEngine
	* Description	: the project aim to create a free and openSource Bi-Dimensional game engine
**/

package CorbaaxEngine

import (
	"image/color"

	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

//#region VectorialForce
type VectorialForce struct {
	Force, Theta float64
}

// Method

// Return the Converted Value of the current object in Algebric Form ;
// start 	: F(cos(O)-sin(O))
// ret 		: A + B
// A 		= F * cos(O * Math*Pi)
// O 		= F * -sin(O * Math*Pi)
func (vec VectorialForce) ToAlgebricForm() AlgebricForce {
	return AlgebricForce{
		Xaxsis: vec.Force * math.Cos(vec.Theta*math.Pi),
		Yaxsis: vec.Force * math.Sin(vec.Theta*math.Pi),
	}
}

// Func

// Return the Generate VectorialForce of the value given: Force, Theta
func NewVectorialForce(Force, Theta float64) VectorialForce {
	return VectorialForce{
		Force: Force,
		Theta: Theta,
	}
}

// Return a VectorialForce with the value Sums of the two Forces given in VectorialForm A,B
func SumVectorForces(A, B VectorialForce) VectorialForce {
	return SumAlgebricForces(A.ToAlgebricForm(), B.ToAlgebricForm()).ToVectorialForm()
}

func SumOfVForces(Forces []VectorialForce) VectorialForce {
	// BlankVector
	V := NewVectorialForce(0, 0)
	// Sum Loop
	for i := range Forces {
		V = SumVectorForces(V, Forces[i])
	}
	return V
}

//#endregion

//#region AlgebricForce
type AlgebricForce struct {
	Xaxsis, Yaxsis float64
}

// Method

// Return the Converted Value of the current object in Vectorial Form ;
// start 	: A + B
// ret 		: F(cos(O)-sin(O))
// F 		= √(A^2 + B^2)
// O 		= arcTg(B / A)
func (alg AlgebricForce) ToVectorialForm() VectorialForce {
	return VectorialForce{
		Force: math.Sqrt(math.Pow(alg.Xaxsis, 2) + math.Pow(alg.Yaxsis, 2)),
		Theta: math.Atan(alg.Yaxsis/alg.Xaxsis) / math.Pi,
	}
}

//Func

// Return the Generate AlgebricalForce of the value given: X, Y
func NewAlgebricForce(X, Y float64) AlgebricForce {
	return AlgebricForce{
		Xaxsis: X,
		Yaxsis: Y,
	}
}

// Return a AlgebricForce with the value Sums two Forces given in AlgebricForm A,B
func SumAlgebricForces(A, B AlgebricForce) AlgebricForce {
	return AlgebricForce{
		Xaxsis: A.Xaxsis + B.Xaxsis,
		Yaxsis: A.Yaxsis + B.Yaxsis,
	}
}

// Return the Sums of the forces contained in a []AlgebricForce
func SumOfAForces(Forces []AlgebricForce) AlgebricForce {
	// BlankVector
	A := NewAlgebricForce(0, 0)
	// Sum Loop
	for i := range Forces {
		A = SumAlgebricForces(A, Forces[i])
	}
	return A
}

//#endregion

//#region Point
type Point struct {
	X, Y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		X: x,
		Y: y,
	}
}

//#endregion

//#region HitBox

type HitBox struct {
	PosA, PosB Point
}

func NewHitBoxFromPoint(a, b Point) (HitBox, bool) {
	if a.X < b.X && a.Y < b.Y {
		return HitBox{
			PosA: a,
			PosB: b,
		}, true
	} else {
		return HitBox{}, false
	}

}

func NewHitBox(x1, y1, x2, y2 float64) (HitBox, bool) {
	return NewHitBoxFromPoint(NewPoint(x1, y1), NewPoint(x2, y2))
}

func (bx HitBox) RenderFullHitbox() (disp *ebiten.Image) {
	disp = ebiten.NewImage(int(bx.PosB.X-bx.PosA.X), int(bx.PosB.Y-bx.PosA.Y))
	disp.Fill(color.White)
	return
}

func (bx *HitBox) PositionUpdate(offsetX, offsetY float64) {
	bx.PosA.X += offsetX
	bx.PosB.X += offsetX
	bx.PosA.Y += offsetY
	bx.PosB.Y += offsetY
}

func betweenValue(a, b, c float64) bool {
	if a < c && c > b {
		return true
	}
	return false
}

func (bx HitBox) IsGoingToCollide(bx2 HitBox) bool {
	if betweenValue(bx.PosA.X, bx.PosB.X, bx2.PosA.X) || betweenValue(bx.PosA.X, bx.PosB.X, bx2.PosB.X) && betweenValue(bx.PosA.Y, bx.PosB.Y, bx2.PosA.Y) || betweenValue(bx.PosA.Y, bx.PosB.Y, bx2.PosB.Y) {
		return true
	}
	return false
}

//#endregion

//#region Player

type Player struct {
	Force  VectorialForce
	HitBox HitBox
	Image  *ebiten.Image
	X, Y   float64
}

func NewPlayer(vec VectorialForce, bx HitBox, sprite *ebiten.Image, x, y float64) Player {
	return Player{
		Force:  vec,
		HitBox: bx,
		Image:  sprite,
		X:      x,
		Y:      y,
	}
}

func (px Player) PositionUpdate() Player {
	offX := px.Force.Force * math.Cos(px.Force.Theta)
	offY := px.Force.Force * (-math.Sin(px.Force.Theta))
	px.X += offX
	px.Y += offY
	px.HitBox.PositionUpdate(offX, offY)
	return px
}

//#endregion
/*
func main() {
	a := VectorialForce{
		Force: 2,
		Theta: 0.5,
	}
	b := a.ToAlgebricForm()
	fmt.Println("alg:", b.Xaxsis, b.Yaxsis)
	c := b.ToVectorialForm()
	fmt.Println("Vec:", c.Force, c.Theta)
}
*/
