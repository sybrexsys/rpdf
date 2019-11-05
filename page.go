package rpdf

type Page struct {
	Canvas
}

func newPage(rm ResourceManager) *Page {
	return &Page{
		Canvas: *newCanvas(rm),
	}
}
