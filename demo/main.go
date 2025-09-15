package main

import (
	"fmt"
	"net/http"
	"time"

	. "github.com/bloxui/blox"
	"github.com/bloxui/htmx"
)

func main() {
	// Serve embedded HTMX JavaScript with compression
	http.HandleFunc("/js/htmx.min.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		w.Header().Set("Cache-Control", "public, max-age=31536000") // Cache for 1 year

		js := htmx.JavaScript()
		_, _ = w.Write(js)
	})

	// Serve robots.txt
	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		_, _ = fmt.Fprint(w, "User-agent: *\nAllow: /\n")
	})

	// Serve demo page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := Html(
			Lang("en"),
			Head(
				HeadTitle(T("HTMX + Blox Demo")),
				Meta(Charset("UTF-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
				Meta(Name("description"), Content("Interactive demo showcasing HTMX attributes with Blox's type-safe HTML generation in Go. Features embedded JavaScript, todo list, and real-time updates.")),
				Style(T(`
					body { font-family: system-ui, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
					.container { background: #f5f5f5; padding: 20px; border-radius: 8px; margin: 20px 0; }
					.btn { background: #0056b3; color: white; border: none; padding: 8px 16px; border-radius: 4px; cursor: pointer; margin: 5px; }
					.btn:hover { background: #004085; }
					.btn-success { background: #1e7e34; color: white; }
					.btn-success:hover { background: #155724; }
					.btn-danger { background: #c82333; color: white; }
					.btn-danger:hover { background: #a71e2a; }
					.form-input { border: 1px solid #ccc; padding: 8px; border-radius: 4px; margin: 5px; }
					.todo-container { min-height: 100px; background: white; border: 1px solid #ddd; padding: 10px; margin: 10px 0; border-radius: 4px; }
					.todo-item { padding: 10px; border-bottom: 1px solid #eee; display: flex; justify-content: space-between; align-items: center; }
					.htmx-indicator { display: none; }
					.htmx-request .htmx-indicator { display: inline; }
				`)),
			),
			Body(
				H1(T("HTMX + Blox Integration Demo")),
				P(T("This demo showcases HTMX attributes working with Blox's type-safe HTML generation, using embedded JavaScript (no CDN required).")),

				Main(
					// Content area
					Div(
						Id("content"),
						Class("container"),
						H2(T("Dynamic Content Area")),
						P(T("Click the button below to load content dynamically.")),
					),

					// Load content button
					Button(
						T("Load Content"),
						htmx.HxGet("/api/content"),
						htmx.HxTarget("#content"),
						htmx.HxSwap("innerHTML"),
						Class("btn"),
					),

					// Todo section
					H2(T("Todo List Demo")),
					Div(Id("todo-list"), Class("todo-container")),

					// Todo form
					Form(
						htmx.HxPost("/api/todos"),
						htmx.HxTarget("#todo-list"),
						htmx.HxSwap("beforeend"),
						StyleKV("display", "flex"),
						StyleKV("gap", "10px"),
						StyleKV("margin", "10px 0"),

						Input(
							InputType("text"),
							InputName("todo"),
							Placeholder("Add a todo..."),
							Required(),
							Class("form-input"),
							StyleKV("flex", "1"),
						),

						Button(
							ButtonType("submit"),
							T("Add Todo"),
							Class("btn btn-success"),
						),
					),

					// Clear all button
					Button(
						T("Clear All"),
						htmx.HxDelete("/api/todos"),
						htmx.HxTarget("#todo-list"),
						htmx.HxConfirm("Delete all todos?"),
						htmx.HxSwap("innerHTML"),
						Class("btn btn-danger"),
						StyleKV("margin-top", "10px"),
					),

					// Auto-refresh demo
					H2(T("Auto-refresh Demo")),
					Div(
						Id("live-time"),
						htmx.HxGet("/api/time"),
						htmx.HxTrigger("every 2s"),
						htmx.HxSwap("innerHTML"),
						Class("container"),
						T("Loading time..."),
					),
				),

				// Load embedded HTMX JavaScript with defer for non-blocking
				Script(ScriptSrc("/js/htmx.min.js"), Defer()),
			),
		)

		w.Header().Set("Content-Type", "text/html")
		_, _ = fmt.Fprint(w, "<!DOCTYPE html>\n")
		_, _ = fmt.Fprint(w, Render(page))
	})

	// API endpoints
	http.HandleFunc("/api/content", func(w http.ResponseWriter, r *http.Request) {
		response := Div(
			H2(T("✨ Content Loaded Successfully!")),
			P(T("This content was loaded dynamically using HTMX and rendered with Blox.")),
			P(T("The HTMX JavaScript is served directly from the Go application using embedded files.")),
			Button(
				T("Load More Content"),
				htmx.HxGet("/api/more"),
				htmx.HxTarget("#content"),
				htmx.HxSwap("innerHTML"),
				Class("btn"),
			),
		)
		_, _ = fmt.Fprint(w, Render(response))
	})

	http.HandleFunc("/api/more", func(w http.ResponseWriter, r *http.Request) {
		response := Div(
			H2(T("🚀 Even More Content!")),
			P(T("This demonstrates nested HTMX requests working seamlessly.")),
			Ul(
				Li(T("Type-safe HTML with compile-time validation")),
				Li(T("Zero runtime overhead")),
				Li(T("Embedded JavaScript assets")),
				Li(T("Perfect integration with HTMX")),
			),
		)
		_, _ = fmt.Fprint(w, Render(response))
	})

	todoCounter := 0
	http.HandleFunc("/api/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			todoText := r.FormValue("todo")
			if todoText != "" {
				todoCounter++
				response := Div(
					Class("todo-item"),
					Span(T(fmt.Sprintf("%d. %s", todoCounter, todoText))),
					Button(
						T("Delete"),
						htmx.HxDelete(fmt.Sprintf("/api/todos/%d", todoCounter)),
						htmx.HxTarget("closest .todo-item"),
						htmx.HxSwap("outerHTML"),
						Class("btn btn-danger"),
						StyleKV("font-size", "12px"),
						StyleKV("padding", "4px 8px"),
					),
				)
				_, _ = fmt.Fprint(w, Render(response))
			}
		case "DELETE":
			// Clear all todos
			_, _ = fmt.Fprint(w, "")
		}
	})

	http.HandleFunc("/api/todos/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			// Delete individual todo (return empty to remove element)
			_, _ = fmt.Fprint(w, "")
		}
	})

	http.HandleFunc("/api/time", func(w http.ResponseWriter, r *http.Request) {
		response := Div(
			Strong(T("Current Time: ")),
			T(time.Now().Format("2006-01-02 15:04:05")),
			Br(),
			T("This updates every 2 seconds automatically!"),
		)
		_, _ = fmt.Fprint(w, Render(response))
	})

	fmt.Println("🚀 HTMX + Blox Demo Server starting on :8080")
	fmt.Println("📦 Using embedded HTMX JavaScript (no CDN required)")
	fmt.Println("🔗 Open http://localhost:8080 to view the demo")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
