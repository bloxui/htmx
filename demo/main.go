package main

import (
	"fmt"
	"net/http"
	"time"

	. "github.com/bloxui/blox"
	"github.com/bloxui/htmx"
	icons "github.com/bloxui/icons/lucide"
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
				HeadStyle(T(`
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
					/* Icon size utilities */
					.size-3 { width: 12px; height: 12px; }
					.size-4 { width: 16px; height: 16px; }
					.size-5 { width: 20px; height: 20px; }
					.size-6 { width: 24px; height: 24px; }
					.size-8 { width: 32px; height: 32px; }
					/* Icon alignment */
					svg { vertical-align: middle; }
					h1, h2, h3, h4, h5, h6 { display: flex; align-items: center; line-height: 1.2; }
					h1 svg, h2 svg, h3 svg, h4 svg, h5 svg, h6 svg { flex-shrink: 0; }
					p { display: flex; align-items: center; }
					.btn { display: inline-flex; align-items: center; }
				`)),
			),
			Body(
				H1(
					icons.Zap(Class("size-8"), Style("margin-right", "10px")),
					T("HTMX + Blox Integration Demo"),
				),
				P(T("This demo showcases HTMX attributes working with Blox's type-safe HTML generation, using embedded JavaScript (no CDN required).")),

				Main(
					// Content area
					Div(
						Id("content"),
						Class("container"),
						H2(
							icons.Monitor(Class("size-6"), Style("margin-right", "8px")),
							T("Dynamic Content Area"),
						),
						P(T("Click the button below to load content dynamically.")),
					),

					// Load content button
					Button(
						icons.Download(Class("size-4"), Style("margin-right", "6px")),
						T("Load Content"),
						htmx.HxGet("/api/content"),
						htmx.HxTarget("#content"),
						htmx.HxSwap("innerHTML"),
						Class("btn"),
						Style("display", "inline-flex"),
						Style("align-items", "center"),
					),

					// Todo section
					H2(
						icons.ListTodo(Class("size-6"), Style("margin-right", "8px")),
						T("Todo List Demo"),
					),
					Div(
						Id("todo-list"),
						Class("todo-container"),
						htmx.HxGet("/api/todos"),
						htmx.HxTrigger("load"),
						htmx.HxSwap("innerHTML"),
					),

					// Todo form
					Form(
						htmx.HxPost("/api/todos"),
						htmx.HxTarget("#todo-list"),
						htmx.HxSwap("beforeend"),
						Style("display", "flex"),
						Style("gap", "10px"),
						Style("margin", "10px 0"),

						Input(
							InputType("text"),
							InputName("todo"),
							Placeholder("Add a todo..."),
							Required(),
							Class("form-input"),
							Style("flex", "1"),
						),

						Button(
							ButtonType("submit"),
							icons.Plus(Class("size-4"), Style("margin-right", "6px")),
							T("Add Todo"),
							Class("btn btn-success"),
							Style("display", "inline-flex"),
							Style("align-items", "center"),
						),
					),

					// Clear all button
					Button(
						icons.Trash2(Class("size-4"), Style("margin-right", "6px")),
						T("Clear All"),
						htmx.HxDelete("/api/todos"),
						htmx.HxTarget("#todo-list"),
						htmx.HxConfirm("Delete all todos?"),
						htmx.HxSwap("innerHTML"),
						Class("btn btn-danger"),
						Style("margin-top", "10px"),
						Style("display", "inline-flex"),
						Style("align-items", "center"),
					),

					// Auto-refresh demo
					H2(
						icons.RefreshCw(Class("size-6"), Style("margin-right", "8px")),
						T("Auto-refresh Demo"),
					),
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
			H2(
				icons.CircleCheck(Class("size-6"), Style("margin-right", "8px"), Style("color", "#10b981")),
				T("Content Loaded Successfully!"),
			),
			P(T("This content was loaded dynamically using HTMX and rendered with Blox.")),
			P(T("The HTMX JavaScript is served directly from the Go application using embedded files.")),
			Button(
				icons.ArrowRight(Class("size-4"), Style("margin-right", "6px")),
				T("Load More Content"),
				htmx.HxGet("/api/more"),
				htmx.HxTarget("#content"),
				htmx.HxSwap("innerHTML"),
				Class("btn"),
				Style("display", "inline-flex"),
				Style("align-items", "center"),
			),
		)
		_, _ = fmt.Fprint(w, Render(response))
	})

	http.HandleFunc("/api/more", func(w http.ResponseWriter, r *http.Request) {
		response := Div(
			H2(
				icons.Rocket(Class("size-6"), Style("margin-right", "8px"), Style("color", "#f59e0b")),
				T("Even More Content!"),
			),
			P(T("This demonstrates nested HTMX requests working seamlessly.")),
			Ul(
				Li(
					icons.Shield(Class("size-4"), Style("margin-right", "6px"), Style("color", "#10b981")),
					T("Type-safe HTML with compile-time validation"),
				),
				Li(
					icons.Zap(Class("size-4"), Style("margin-right", "6px"), Style("color", "#f59e0b")),
					T("Zero runtime overhead"),
				),
				Li(
					icons.Package(Class("size-4"), Style("margin-right", "6px"), Style("color", "#3b82f6")),
					T("Embedded JavaScript assets"),
				),
				Li(
					icons.Heart(Class("size-4"), Style("margin-right", "6px"), Style("color", "#ef4444")),
					T("Perfect integration with HTMX"),
				),
			),
		)
		_, _ = fmt.Fprint(w, Render(response))
	})

	// In-memory todo storage
	type Todo struct {
		ID   int
		Text string
	}

	var todos []Todo
	var todoCounter int

	http.HandleFunc("/api/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			// Return all todos
			var todoArgs []DivArg
			for _, todo := range todos {
				todoItem := Div(
					Class("todo-item"),
					Span(
						icons.Circle(Class("size-4"), Style("margin-right", "8px"), Style("color", "#6b7280")),
						T(fmt.Sprintf("%d. %s", todo.ID, todo.Text)),
					),
					Button(
						icons.X(Class("size-3"), Style("margin-right", "4px")),
						T("Delete"),
						htmx.HxDelete(fmt.Sprintf("/api/todos/%d", todo.ID)),
						htmx.HxTarget("closest .todo-item"),
						htmx.HxSwap("outerHTML"),
						Class("btn btn-danger"),
						Style("font-size", "12px"),
						Style("padding", "4px 8px"),
						Style("display", "inline-flex"),
						Style("align-items", "center"),
					),
				)
				todoArgs = append(todoArgs, Child(todoItem))
			}
			response := Div(todoArgs...)
			_, _ = fmt.Fprint(w, Render(response))

		case "POST":
			todoText := r.FormValue("todo")
			if todoText != "" {
				todoCounter++
				newTodo := Todo{ID: todoCounter, Text: todoText}
				todos = append(todos, newTodo)

				response := Div(
					Class("todo-item"),
					Span(
						icons.Circle(Class("size-4"), Style("margin-right", "8px"), Style("color", "#6b7280")),
						T(fmt.Sprintf("%d. %s", newTodo.ID, newTodo.Text)),
					),
					Button(
						icons.X(Class("size-3"), Style("margin-right", "4px")),
						T("Delete"),
						htmx.HxDelete(fmt.Sprintf("/api/todos/%d", newTodo.ID)),
						htmx.HxTarget("closest .todo-item"),
						htmx.HxSwap("outerHTML"),
						Class("btn btn-danger"),
						Style("font-size", "12px"),
						Style("padding", "4px 8px"),
						Style("display", "inline-flex"),
						Style("align-items", "center"),
					),
				)
				_, _ = fmt.Fprint(w, Render(response))
			}
		case "DELETE":
			// Clear all todos
			todos = []Todo{}
			_, _ = fmt.Fprint(w, "")
		}
	})

	http.HandleFunc("/api/todos/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			// Extract todo ID from URL path
			path := r.URL.Path
			if len(path) > len("/api/todos/") {
				idStr := path[len("/api/todos/"):]
				var id int
				if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
					// Remove todo from slice
					for i, todo := range todos {
						if todo.ID == id {
							todos = append(todos[:i], todos[i+1:]...)
							break
						}
					}
				}
			}
			// Return empty to remove element from DOM
			_, _ = fmt.Fprint(w, "")
		}
	})

	http.HandleFunc("/api/time", func(w http.ResponseWriter, r *http.Request) {
		response := Div(
			P(
				icons.Clock(Class("size-5"), Style("margin-right", "8px"), Style("color", "#3b82f6")),
				Strong(T("Current Time: ")),
				T(time.Now().Format("2006-01-02 15:04:05")),
			),
			P(
				icons.RotateCcw(Class("size-4"), Style("margin-right", "6px"), Style("color", "#10b981")),
				T("This updates every 2 seconds automatically!"),
			),
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
