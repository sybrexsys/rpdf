package rpdf

type Canvas struct {
	rm ResourceManager
}

func newCanvas(rm ResourceManager) *Canvas {
	return &Canvas{
		rm: rm,
	}

}
