package apex_test

import (
	apexadapter "logur.dev/adapter/apex"
)

func ExampleNew() {
	var l interface{}
	logger := apexadapter.New(l)

	// Output:
	_ = logger
}

// If logger is nil, a default instance is created.
func ExampleNew_default() {
	logger := apexadapter.New(nil)

	// Output:
	_ = logger
}
