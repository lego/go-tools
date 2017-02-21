package pkg

import "runtime"

func fn() {
	var x *int
	foo := func(y *int) { println(x) }
	runtime.SetFinalizer(x, foo)
	runtime.SetFinalizer(x, nil)
	runtime.SetFinalizer(x, func(_ *int) {
		println(x)
	})

	foo = func(y *int) { println(y) }
	runtime.SetFinalizer(x, foo)
	runtime.SetFinalizer(x, func(y *int) {
		println(y)
	})
}

// MATCH:8 /the finalizer closes over the object, preventing the finalizer from ever running \(at .+:7:9/
// MATCH:10 /the finalizer closes over the object, preventing the finalizer from ever running \(at .+:10:26/
