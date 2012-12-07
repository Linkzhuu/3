package render

import (
	"code.google.com/p/nimble-cube/core"
	"code.google.com/p/nimble-cube/dump"
	gl "github.com/chsc/gogl/gl21"
	"math/rand"
	"os"
)

func Load(fname string) *dump.Frame {
	f, err := os.Open(fname)
	core.Fatal(err)
	defer f.Close()

	r := dump.NewReader(f, dump.CRC_ENABLED)
	core.Fatal(r.Read())
	return &(r.Frame)
}

func Render(frame *dump.Frame) {
	ClearScene()

	size := frame.MeshSize
	cell := frame.MeshStep
	maxworld := 0.
	for i := range size {
		world := float64(size[i]) * cell[i]
		if world > maxworld {
			maxworld = world
		}
	}
	scale := 1 / maxworld
	rx, ry, rz := float32(0.5*scale*cell[0]), float32(0.5*scale*cell[1]), float32(0.5*scale*cell[2])

	rand.Seed(0)
	m := frame.Vectors()
	i := 0
	{
		//for i := range m[0] {
		x := float32(scale * cell[0] * (float64(i-size[0]/2) + 0.5))
		for j := range m[0][i] {
			y := float32(scale * cell[1] * (float64(j-size[1]/2) + 0.5))
			for k := range m[0][i][j] {
				z := float32(scale * cell[2] * (float64(k-size[2]/2) + 0.5))
				rnd := gl.Float(rand.Float32()*0.5 + 0.5)
				gl.Color3f(rnd, rnd, rnd)
				(&Cube{Vertex{z, y, x}, Vertex{rz, ry, rx}}).Render()
			}
		}
	}

}

func ClearScene() {
	ambient := []gl.Float{0.7, 0.7, 0.7, 1}
	diffuse := []gl.Float{1, 1, 1, 1}
	lightpos := []gl.Float{0.2, 0.5, 1, 1}
	gl.Lightfv(gl.LIGHT0, gl.AMBIENT, &ambient[0])
	gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, &diffuse[0])
	gl.Lightfv(gl.LIGHT0, gl.POSITION, &lightpos[0])
	gl.Enable(gl.LIGHT0)
	ambdiff := []gl.Float{0.5, 0.5, 0.5, 1}
	gl.Materialfv(gl.FRONT_AND_BACK, gl.AMBIENT_AND_DIFFUSE, &ambdiff[0])

	gl.ClearColor(1, 1, 1, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}
