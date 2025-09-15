package js

import _ "embed"

// HtmxMinJS contains the minified HTMX JavaScript library (v1.9.12).
// This can be served directly or embedded in HTML pages.
//
//go:embed htmx.min.js
var HtmxMinJS []byte
