package pkg

import "time"

const c1 = "12345"
const c2 = "2006"

func fn() {
	_, _ = time.Parse("12345", "") // MATCH /parsing time/
	_, _ = time.Parse(c1, "")      // MATCH /parsing time/
	_, _ = time.Parse(c2, "")
	_, _ = time.Parse(time.RFC3339Nano, "")
	_, _ = time.Parse(time.Kitchen, "")
}
