package htmx

import (
	"strings"
	"testing"
)

func TestJavaScript(t *testing.T) {
	js := JavaScript()

	if len(js) == 0 {
		t.Fatal("JavaScript() returned empty content")
	}

	jsString := string(js)

	// Check for HTMX-specific content
	if !strings.Contains(jsString, "htmx") {
		t.Error("JavaScript content does not appear to contain HTMX library")
	}

	// Should be minified (no excessive whitespace)
	lines := strings.Split(jsString, "\n")
	if len(lines) > 10 {
		t.Error("JavaScript appears to not be minified (too many lines)")
	}

	// Check it starts with expected pattern for minified HTMX
	if !strings.HasPrefix(jsString, "var htmx=function()") {
		t.Error("JavaScript does not start with expected HTMX pattern")
	}
}

func TestJavaScriptSize(t *testing.T) {
	js := JavaScript()

	// HTMX minified is typically around 40-50KB
	if len(js) < 30000 {
		t.Error("JavaScript content seems too small to be complete HTMX library")
	}

	if len(js) > 100000 {
		t.Error("JavaScript content seems too large for minified HTMX library")
	}
}