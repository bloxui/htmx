// Package htmx provides HTMX attribute functions for the Plain component library.
package htmx

import (
	"github.com/plainkit/html"
	"github.com/plainkit/htmx/js"
)

// Core HTMX HTTP methods

// HxGet sets the hx-get attribute for HTMX GET requests.
func HxGet(url string) html.Global {
	return html.Custom("hx-get", url)
}

// HxPost sets the hx-post attribute for HTMX POST requests.
func HxPost(url string) html.Global {
	return html.Custom("hx-post", url)
}

// HxPut sets the hx-put attribute for HTMX PUT requests.
func HxPut(url string) html.Global {
	return html.Custom("hx-put", url)
}

// HxDelete sets the hx-delete attribute for HTMX DELETE requests.
func HxDelete(url string) html.Global {
	return html.Custom("hx-delete", url)
}

// HxPatch sets the hx-patch attribute for HTMX PATCH requests.
func HxPatch(url string) html.Global {
	return html.Custom("hx-patch", url)
}

// Targeting and content manipulation

// HxTarget sets the hx-target attribute to specify where to place the response.
func HxTarget(selector string) html.Global {
	return html.Custom("hx-target", selector)
}

// HxSwap sets the hx-swap attribute to specify how the response will be swapped in.
// Common values: "innerHTML", "outerHTML", "beforebegin", "afterbegin", "beforeend", "afterend", "delete", "none"
func HxSwap(strategy string) html.Global {
	return html.Custom("hx-swap", strategy)
}

// HxSwapOob sets the hx-swap-oob attribute for out-of-band swaps.
func HxSwapOob(value string) html.Global {
	return html.Custom("hx-swap-oob", value)
}

// Event handling

// HxTrigger sets the hx-trigger attribute to specify what triggers the request.
// Examples: "click", "submit", "change", "load", "revealed", "intersect"
func HxTrigger(event string) html.Global {
	return html.Custom("hx-trigger", event)
}

// Loading states and indicators

// HxIndicator sets the hx-indicator attribute to show/hide loading indicators.
func HxIndicator(selector string) html.Global {
	return html.Custom("hx-indicator", selector)
}

// HxDisabledElt sets the hx-disabled-elt attribute to disable elements during requests.
func HxDisabledElt(selector string) html.Global {
	return html.Custom("hx-disabled-elt", selector)
}

// Request configuration

// HxHeaders sets the hx-headers attribute to send additional headers.
// Value should be a JSON object as string, e.g., `{"X-Custom": "value"}`
func HxHeaders(headers string) html.Global {
	return html.Custom("hx-headers", headers)
}

// HxVals sets the hx-vals attribute to include additional values in the request.
// Value should be a JSON object as string, e.g., `{"key": "value"}`
func HxVals(vals string) html.Global {
	return html.Custom("hx-vals", vals)
}

// HxInclude sets the hx-include attribute to include additional form elements.
func HxInclude(selector string) html.Global {
	return html.Custom("hx-include", selector)
}

// HxParams sets the hx-params attribute to filter which parameters to include.
// Values: "all", "none", or comma-separated list of parameter names
func HxParams(params string) html.Global {
	return html.Custom("hx-params", params)
}

// Navigation and history

// HxBoost sets the hx-boost attribute to progressively enhance links and forms.
func HxBoost(enabled bool) html.Global {
	if enabled {
		return html.Custom("hx-boost", "true")
	}
	return html.Custom("hx-boost", "false")
}

// HxPushUrl sets the hx-push-url attribute to update the browser URL.
// Use "true" to push the request URL, "false" to prevent, or a custom URL string.
func HxPushUrl(url string) html.Global {
	return html.Custom("hx-push-url", url)
}

// HxReplaceUrl sets the hx-replace-url attribute to replace the current browser URL.
func HxReplaceUrl(url string) html.Global {
	return html.Custom("hx-replace-url", url)
}

// User interaction

// HxConfirm sets the hx-confirm attribute to show a confirmation dialog.
func HxConfirm(message string) html.Global {
	return html.Custom("hx-confirm", message)
}

// HxPrompt sets the hx-prompt attribute to show a prompt dialog.
func HxPrompt(message string) html.Global {
	return html.Custom("hx-prompt", message)
}

// Advanced features

// HxExt sets the hx-ext attribute to enable HTMX extensions.
// Examples: "json-enc", "morphdom", "alpine-morph", "response-targets"
func HxExt(extensions string) html.Global {
	return html.Custom("hx-ext", extensions)
}

// HxSelect sets the hx-select attribute to select a subset of the response.
func HxSelect(selector string) html.Global {
	return html.Custom("hx-select", selector)
}

// HxSelectOob sets the hx-select-oob attribute for out-of-band content selection.
func HxSelectOob(selector string) html.Global {
	return html.Custom("hx-select-oob", selector)
}

// Synchronization

// HxSync sets the hx-sync attribute to control request synchronization.
// Examples: "drop", "abort", "replace", "queue"
func HxSync(strategy string) html.Global {
	return html.Custom("hx-sync", strategy)
}

// Form handling

// HxEncoding sets the hx-encoding attribute to specify form encoding.
// Common value: "multipart/form-data"
func HxEncoding(encoding string) html.Global {
	return html.Custom("hx-encoding", encoding)
}

// Error handling and validation

// HxValidate sets the hx-validate attribute to force validation before submission.
func HxValidate() html.Global {
	return html.Custom("hx-validate", "true")
}

// Server-sent events

// HxSse sets the hx-sse attribute to connect to a server-sent events source.
// Example: "connect:/events"
func HxSse(value string) html.Global {
	return html.Custom("hx-sse", value)
}

// WebSockets

// HxWs sets the hx-ws attribute to connect to a WebSocket.
// Example: "connect:/ws"
func HxWs(value string) html.Global {
	return html.Custom("hx-ws", value)
}

// Preservation

// HxPreserve sets the hx-preserve attribute to preserve elements during swaps.
func HxPreserve() html.Global {
	return html.Custom("hx-preserve", "true")
}

// Disinheritance (preventing inheritance of HTMX attributes)

// HxDisinherit sets the hx-disinherit attribute to prevent inheritance of specific attributes.
// Examples: "*", "hx-target hx-get", "hx-*"
func HxDisinherit(attrs string) html.Global {
	return html.Custom("hx-disinherit", attrs)
}

// JavaScript returns the embedded HTMX JavaScript content.
// This can be used to serve the HTMX library directly from your Go application
// without requiring external CDN dependencies.
//
// Example usage:
//
//	http.HandleFunc("/js/htmx.min.js", func(w http.ResponseWriter, r *http.Request) {
//	    w.Header().Set("Content-Type", "application/javascript")
//	    w.Write(htmx.JavaScript())
//	})
func JavaScript() []byte {
	return js.HtmxMinJS
}
