package rpdf

type DocumentInformation struct {
	Creator  string
	Subject  string
	Keywords string
	Author   string
	Title    string
}

type Setting struct {
	Info            DocumentInformation
	DefaultPageSize CanvasSize
	PackContent     bool
}

func DefaultSettings() Setting {
	return Setting{
		DefaultPageSize: PageSizeA4,
		PackContent:     true,
		Info: DocumentInformation{
			Creator: "golang application",
		},
	}
}
