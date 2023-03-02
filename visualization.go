package main

import (
	"fmt"
	"image/color"
	"log"
	"strings"
	"time"

	"github.com/fstanis/screenresolution"
	"github.com/fzipp/canvas"
)

type Pair struct {
	previous string
	current  string
	dx       float64
	dy       float64
	x        float64
	y        float64
	v0       *Vertex
	v1       *Vertex
}

const RECT_SIDE = 20

var resolutionWidth int
var resolutionHeight int

func visualization(g *Graph) {
	fmt.Println("\nStarting server at http://127.0.0.1:8080/")
	fmt.Println("Quit the server with CONTROL-C.")

	resolution := screenresolution.GetPrimary()
	if resolution == nil {
		resolutionWidth = 1920
		resolutionHeight = 1080
	} else {
		resolutionWidth = resolution.Width
		resolutionHeight = resolution.Height
	}

	err := canvas.ListenAndServe(":8080", func(ctx *canvas.Context) {
		run(ctx, g) // pass variable to run function
	}, canvas.Size(resolutionWidth, resolutionHeight), canvas.Title("Running ants"))
	if err != nil {
		log.Fatal(err)
	}
}

func drawEdges(ctx *canvas.Context, g *Graph) {
	s := readFile()

	ctx.SetFillStyle(color.RGBA{R: 0, G: 0, B: 0, A: 255})
	for _, links := range theLinks(s) {
		vertex := strings.Split(links, ", ")
		x0 := scale(float64(g.getVertex(vertex[0]).x)) + RECT_SIDE/2
		y0 := scale(float64(g.getVertex(vertex[0]).y)) + RECT_SIDE/2
		x1 := scale(float64(g.getVertex(vertex[1]).x)) + RECT_SIDE/2
		y1 := scale(float64(g.getVertex(vertex[1]).y)) + RECT_SIDE/2
		drawLine(ctx, x0, y0, x1, y1)
	}
}

func drawPoints(ctx *canvas.Context, g *Graph) {
	ctx.SetFillStyle(color.RGBA{R: 180, G: 180, B: 180, A: 255})
	for _, vertex := range g.vertices {
		switch vertex.name {
		case g.start:
			ctx.SetFillStyle(color.RGBA{G: 200, A: 255})
			x := scale(float64(vertex.x))
			y := scale(float64(vertex.y))
			ctx.FillRect(x, y, RECT_SIDE, RECT_SIDE)
		case g.end:
			ctx.SetFillStyle(color.RGBA{R: 200, A: 255})
			x := scale(float64(vertex.x))
			y := scale(float64(vertex.y))
			ctx.FillRect(x, y, RECT_SIDE, RECT_SIDE)
		default:
			x := scale(float64(vertex.x))
			y := scale(float64(vertex.y))
			ctx.FillRect(x, y, RECT_SIDE, RECT_SIDE)
		}
		ctx.SetFillStyle(color.RGBA{R: 0, G: 0, B: 0, A: 255})
		ctx.FillText(vertex.name, scale(float64(vertex.x)), scale(float64(vertex.y))-10)
		ctx.SetFillStyle(color.RGBA{R: 180, G: 180, B: 180, A: 255})
	}
}

// stretches given coordinates
func scale(n float64) float64 {
	return n*70 + 30
}

// draw edge
func drawLine(ctx *canvas.Context, x0, y0, x1, y1 float64) {
	ctx.BeginPath()
	ctx.MoveTo(x0, y0)
	ctx.LineTo(x1, y1)
	ctx.Stroke()
}

func run(ctx *canvas.Context, g *Graph) {
	// go through answers's steps
	for i := 1; i < len(g.visualization); i++ {
		var p []Pair
		for _, pair := range g.visualization[i] {
			room := strings.Split(pair, "-")
			vertex := room[1]
			ant := room[0]
			var previousVertex string
			// find ant on previous step
			for _, prevPair := range g.visualization[i-1] {
				prevRoom := strings.Split(prevPair, "-")
				prevVertex := prevRoom[1]
				prevAnt := prevRoom[0]
				if ant == prevAnt {
					previousVertex = prevVertex
				}
			}
			// if didn't find - ant previous room was the start room
			if previousVertex == "" {
				previousVertex = g.start
			}
			// add pair previous-current to slice for rendering movings by step
			p = append(p, Pair{current: vertex, previous: previousVertex})
		}
		moveAnts(g, ctx, p)
	}
}

// render ants moving at one step
func moveAnts(g *Graph, ctx *canvas.Context, p []Pair) {

	duration := time.Second * 1
	steps := 60

	// set init values for each pair of vertexes
	for i := 0; i < len(p); i++ {
		p[i].v0 = g.getVertex(p[i].previous)
		p[i].v1 = g.getVertex(p[i].current)

		p[i].dx = float64(p[i].v1.x-p[i].v0.x) / float64(steps)
		p[i].dy = float64(p[i].v1.y-p[i].v0.y) / float64(steps)

		p[i].x = float64(p[i].v0.x)
		p[i].y = float64(p[i].v0.y)
	}

	// render moving by edge
	for i := 0; i < steps; i++ {
		clear(ctx)
		drawEdges(ctx, g)
		drawPoints(ctx, g)
		for i := 0; i < len(p); i++ {
			p[i].x += p[i].dx
			p[i].y += p[i].dy

			// Draw the rectangular object
			x := scale(float64(p[i].x))
			y := scale(float64(p[i].y))
			ctx.SetFillStyle(color.RGBA{B: 200, A: 255})
			ctx.FillRect(x, y, RECT_SIDE, RECT_SIDE)
		}

		ctx.Flush()

		// Wait for the next frame
		time.Sleep(duration / time.Duration(steps))
	}
}

// clear screen
func clear(ctx *canvas.Context) {
	ctx.SetFillStyle(color.RGBA{R: 229, G: 222, B: 202, A: 255})
	ctx.FillRect(0, 0, float64(resolutionWidth), float64(resolutionHeight))
}
