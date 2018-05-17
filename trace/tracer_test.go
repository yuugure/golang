package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("return New is nil")
	} else {
		tracer.Trace("Hello Trace")
		if buf.String() != "Hello Trace\n" {
			t.Errorf("Trace should not write '%s'.", buf.String())
		}
	}
}
