package apex_test

import (
	"github.com/apex/log"

	apexadapter "logur.dev/adapter/apex"
)

func ExampleNew() {
	logger := apexadapter.New()

	// Output:
	_ = logger
}

func ExampleNewFromFields() {
	logger := apexadapter.NewFromFields(log.Fields{
		"name": "mal",
	})

	// Output:
	_ = logger
}

func ExampleNewFromEntry() {
	entry := log.WithFields(log.Fields{
		"name": "mal",
	})

	logger := apexadapter.NewFromEntry(entry)

	// Output:
	_ = logger
}
