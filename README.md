# HTMX for Plain

Type-safe HTMX attributes for the Plain component library.

## Installation

```bash
go get github.com/plainkit/blox
go get github.com/plainkit/htmx
```

## Quick Start

```go
package main

import (
    "fmt"
    x "github.com/plainkit/blox"
    "github.com/plainkit/htmx"
)

func main() {
    button := x.Button(
        x.Text("Click me"),
        htmx.HxGet("/api/data"),
        htmx.HxTarget("#result"),
        htmx.HxSwap("innerHTML"),
    )

    fmt.Println(x.Render(button))
    // Output: <button hx-get="/api/data" hx-target="#result" hx-swap="innerHTML">Click me</button>
}
```

## Core HTMX Attributes

### HTTP Methods
- `HxGet(url)` - GET request
- `HxPost(url)` - POST request
- `HxPut(url)` - PUT request
- `HxDelete(url)` - DELETE request
- `HxPatch(url)` - PATCH request

### Targeting
- `HxTarget(selector)` - Where to place response
- `HxSwap(strategy)` - How to swap content ("innerHTML", "outerHTML", etc.)
- `HxSwapOob(value)` - Out-of-band swaps

### Events
- `HxTrigger(event)` - What triggers the request ("click", "submit", etc.)

### Loading States
- `HxIndicator(selector)` - Loading indicator
- `HxDisabledElt(selector)` - Elements to disable during request

### Request Configuration
- `HxHeaders(json)` - Additional headers as JSON
- `HxVals(json)` - Additional values as JSON
- `HxInclude(selector)` - Include additional form elements
- `HxParams(filter)` - Filter parameters ("all", "none", or names)

### Navigation
- `HxBoost(enabled)` - Progressive enhancement
- `HxPushUrl(url)` - Update browser URL
- `HxReplaceUrl(url)` - Replace browser URL

### User Interaction
- `HxConfirm(message)` - Confirmation dialog
- `HxPrompt(message)` - Prompt dialog

## Complete Example

```go
func todoApp() x.Component {
    return x.Div(
        x.Class("max-w-md mx-auto p-6"),

        // Add todo form
        x.Form(
            htmx.HxPost("/todos"),
            htmx.HxTarget("#todo-list"),
            htmx.HxSwap("beforeend"),
            htmx.HxTrigger("submit"),
            x.Class("mb-4 flex gap-2"),

            x.Input(
                x.InputType("text"),
                x.InputName("todo"),
                x.Placeholder("Add a todo..."),
                x.Required(),
                x.Class("flex-1 px-3 py-2 border rounded"),
            ),

            x.Button(
                x.ButtonType("submit"),
                x.Text("Add"),
                x.Class("px-4 py-2 bg-blue-500 text-white rounded"),
            ),
        ),

        // Todo list container
        x.Div(
            x.Id("todo-list"),
            x.Class("space-y-2"),
        ),

        // Clear all button
        x.Button(
            x.Text("Clear All"),
            htmx.HxDelete("/todos"),
            htmx.HxTarget("#todo-list"),
            htmx.HxConfirm("Delete all todos?"),
            htmx.HxSwap("innerHTML"),
            x.Class("mt-4 px-4 py-2 bg-red-500 text-white rounded"),
        ),
    )
}
```

## Advanced Features

### Extensions
- `HxExt(extensions)` - Enable HTMX extensions ("json-enc", "morphdom", etc.)

### Content Selection
- `HxSelect(selector)` - Select subset of response
- `HxSelectOob(selector)` - Out-of-band content selection

### Synchronization
- `HxSync(strategy)` - Control request synchronization ("drop", "abort", etc.)

### Real-time
- `HxSse(connection)` - Server-sent events
- `HxWs(connection)` - WebSocket connection

### Validation
- `HxValidate()` - Force validation before submission

## Type Safety

All HTMX attributes are type-safe and work with Plain's compile-time validation:

```go
// This works
input := x.Input(
    x.InputType("text"),
    htmx.HxPost("/submit"),
    htmx.HxTrigger("change"),
)

// This also works - HTMX attributes are Global and work with any element
div := x.Div(
    htmx.HxGet("/content"),
    htmx.HxTrigger("revealed"),
    x.Text("Load more..."),
)
```

## License

MIT
