package rpdf

import "io"

type Document struct {
	pages    []*Page
	settings Setting
	stdEngine
	rm ResourceManager
}

func NewDocument(writer io.WriteSeeker, encryptor Encryptor, settings Setting) (*Document, error) {
	return &Document{
		settings: settings,
		stdEngine: stdEngine{
			writer:    writer,
			encryptor: encryptor,
		},
		pages: []*Page{},
		rm:    newStdRecourceManager(),
	}, nil
}

func (doc *Document) NewPage() (*Page, error) {
	page := newPage(doc.rm)
	doc.pages = append(doc.pages, page)
	return page, nil
}

func (doc *Document) FlushAndDone() error {
	return nil
}

func (doc *Document) GetResourceManager() ResourceManager {
	return doc.rm
}
