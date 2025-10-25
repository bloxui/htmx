package js

import _ "embed"

// HtmxMinJS contains the minified HTMX JavaScript library (v1.9.12).
// This can be served directly or embedded in HTML pages.
//
//go:embed htmx.min.js
var HtmxMinJS []byte

// IdiomorphExtMinJS contains the minified idiomorph extension for HTMX (v0.7.4).
// This extension provides intelligent DOM morphing for smoother updates.
//
//go:embed idiomorph-ext.min.js
var IdiomorphExtMinJS []byte
