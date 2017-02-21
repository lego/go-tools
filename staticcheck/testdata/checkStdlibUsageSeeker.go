package pkg

import "io"

func fn() {
	const SeekStart = 0
	var s io.Seeker
	_, _ = s.Seek(0, 0)
	_, _ = s.Seek(0, io.SeekStart)
	_, _ = s.Seek(io.SeekStart, 0) // MATCH /the first argument of io.Seeker is the offset/
	_, _ = s.Seek(SeekStart, 0)
}
