package rpdf

import "io"

type Canvas struct {
	closed  bool
	started bool
	width   float32
	height  float32
	writer  io.Writer
	rm      ResourceManager
}

func newCanvas(rm ResourceManager, pageWidth, pageHeight float32) *Canvas {
	return &Canvas{
		rm:     rm,
		width:  pageWidth,
		height: pageHeight,
	}
}

func (canvas *Canvas) CloseCanvas() {
	canvas.closed = true
}
