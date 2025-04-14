package ReadingFiles

import "bytes"

// Process transforms input data to uppercase and trims whitespace.
func Process(input []byte) []byte {
	return bytes.ToUpper(bytes.TrimSpace(input))
}
