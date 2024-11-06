package cartridge

import "errors"

var (
	ReadingHeaderError     = errors.New("Failed to read header data")
	UnsupportedFileTypeErr = errors.New("Unsupported file type")
)
