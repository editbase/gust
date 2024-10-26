# 🪁 Gust
Gust is a lightweight, server-driven web framework for Go that seamlessly integrates with HTMX. It provides a modern approach to building dynamic web applications while maintaining the simplicity and performance of Go.

## Features
- 🚀 Zero-JS by default (only HTMX required)
- 🔄 Smart routing with automatic partial updates
- 🎨 Template inheritance and layouts
- 🛠️ Swappable engine architecture
- 🔒 Built-in security headers
- 🎯 Active link tracking
- 🔥 Development mode with hot reloading

## Getting Started
```go
package main

import "github.com/editbase/gust"

func main() {
    app := gust.New().
        WithPort("3000").
        WithTemplateDir("./templates").
        WithStaticDir("./static")

    app.GET("/", func(c *gust.Context) error {
        return c.Render("index.html", nil)
    })

    app.POST("/api/users", func(c *gust.Context) error {
        // Handle user creation
    })

    app.Use(gust.Logger(), gust.Recover())

    app.Run()
}
```

## Package Structure
```text
github.com/editbase/gust/
├── examples/
│   ├── blog/
│   └── saas/
├── internal/
│   ├── engine/
│   │   └── template.go
│   ├── handler/
│   │   └── static.go
│   └── util/
│       └── path.go
├── app.go
├── config.go
├── context.go
├── middleware.go
├── render.go
├── router.go
├── server.go
└── template.go
```

## Roadmap
Core Features
- [ ] Hot Reloading with HTMX partial updates
- [ ] Type-safe routing with path parameter
- [ ] Component-based templating with Templ integration
- [ ] Router integrations (Chi, Echo, Fiber)
- [ ] State management with context providers
- [ ] Server-side rendering (SSR) support

Framework Integrations
- [ ] React/Vue components for complex UI
- [ ] TailwindCSS integration
- [ ] WebSocket support
- [ ] Database integrations (GORM, SQLx)
- [ ] Authentication providers
- [ ] Background job processing

Developer Experience
- [ ] CLI tool for scaffolding
- [ ] Testing utilities and mocks
- [ ] API documentation generation
- [ ] Docker and deployment tooling
- [ ] GitHub Actions workflows
- [ ] Development tools and debugging

Production Features
- [ ] Edge computing support
- [ ] Monitoring and logging
- [ ] Security features and headers
- [ ] Caching and performance
- [ ] Database migrations
- [ ] Load balancing support

## Contributing
Contributions are welcome! If you find a bug or have a suggestion, please open an issue or submit a pull request.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.