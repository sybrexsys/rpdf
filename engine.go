package rpdf

import "io"

type idGetter interface {
}

type stdEngine struct {
	writer    io.WriteSeeker
	encryptor Encryptor
}
