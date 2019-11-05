package rpdf

type Page struct {
	Canvas
}

func newPage(rm ResourceManager, pageWidth, pageHeight float32) *Page {
	return &Page{
		Canvas: *newCanvas(rm, pageWidth, pageHeight),
	}
}

func (page *Page) Close() error {
	page.Canvas.CloseCanvas()
	return nil
}
