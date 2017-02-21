package pkg

import "os"

func fn() {
	_, _ = os.OpenFile("", 0, 644) // MATCH /file mode.+/
}
