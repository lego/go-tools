package pkg

import "net/http"

func fn() {
	// Seen in actual code
	_ = http.ListenAndServe("localhost:8080/", nil) // MATCH /invalid port or service name in host:port pair/
	_ = http.ListenAndServe("localhost", nil)       // MATCH /invalid port or service name in host:port pair/
	_ = http.ListenAndServe("localhost:8080", nil)
	_ = http.ListenAndServe(":8080", nil)
	_ = http.ListenAndServe(":http", nil)
	_ = http.ListenAndServe("localhost:http", nil)
	_ = http.ListenAndServe("local_host:8080", nil)
}
