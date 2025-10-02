<h1 align="center">🛠️ Wells-Go: Modular Go Backend with Clean Architecture</h1>

<p align="center">
  <img src="https://your-image-url.com/project-structure.png" alt="Project Structure" width="600">
</p>

<p align="center">
  A modular Go backend project following <strong>Clean Architecture</strong> principles. Scalable, maintainable, and testable.
</p>

<h2>📌 Table of Contents</h2>
<ul>
  <li><a href="#about-the-project">About The Project</a></li>
  <li><a href="#features">Features</a></li>
  <li><a href="#project-structure">Project Structure</a></li>
  <li><a href="#technologies-used">Technologies Used</a></li>
  <li><a href="#getting-started">Getting Started</a></li>
  <li><a href="#api-documentation">API Documentation</a></li>
  <li><a href="#testing">Testing</a></li>
  <li><a href="#contributing">Contributing</a></li>
  <li><a href="#license">License</a></li>
</ul>

<h2 id="about-the-project">ℹ️ About The Project</h2>
<p>
  <strong>Wells-Go</strong> is a backend application built with Go, following Clean Architecture principles. It handles complex business logic while keeping layers decoupled for easier maintenance and scalability.
</p>

<h2 id="features">🚀 Features</h2>
<ul>
  <li>Modular Architecture with clear separation of concerns</li>
  <li>Snake_case API responses for consistency</li>
  <li>Built-in paging support for list endpoints</li>
  <li>Testable components: unit & integration tests</li>
  <li>Extensible for database, cache, and external service integration</li>
</ul>

<h2 id="project-structure">🗂️ Project Structure</h2>
<pre>
wells-go/
├── application/       # Usecases, services, mappers, DTOs
├── domain/            # Entities, repository interfaces
├── interfaces/http/   # HTTP handlers, routing
├── infrastructure/    # Database, Redis, external services
├── response/          # Response helpers (JSON, paging, error)
├── util/              # Utility functions
├── main.go            # Entry point
└── go.mod             # Go modules
</pre>

<h2 id="technologies-used">⚙️ Technologies Used</h2>
<ul>
  <li>Go 1.18+</li>
  <li>GORM (ORM)</li>
  <li>UUID (Unique Identifiers)</li>
  <li>Gin/Gorilla/Mux (HTTP routing)</li>
  <li>Redis (optional caching)</li>
  <li>Docker (optional containerization)</li>
</ul>

<h2 id="getting-started">🛠️ Getting Started</h2>

<h3>Prerequisites</h3>
<ul>
  <li>Go 1.18 or higher</li>
  <li>Git</li>
  <li>Docker (optional)</li>
</ul>

<h3>Installation</h3>
<pre><code>git clone https://github.com/welliardiansyah/wells-go.git
cd wells-go
go mod tidy
</code></pre>

<h3>Running the Application</h3>
<pre><code>go run main.go
</code></pre>
<p>Or using Docker:</p>
<pre><code>docker-compose up
</code></pre>

<h2 id="api-documentation">📚 API Documentation</h2>
<p>RESTful API endpoints for managing resources, supporting GET, POST, PUT, DELETE methods. See <a href="https://github.com/welliardiansyah/wells-go/tree/master">full documentation</a>.</p>

<h2 id="testing">🧪 Testing</h2>
<pre><code>go test ./...
</code></pre>

<h2 id="contributing">🤝 Contributing</h2>
<p>Contributions are welcome! Please fork the repo and submit a pull request. Make sure to:</p>
<ul>
  <li>Follow coding style</li>
  <li>Write tests for new features</li>
  <li>Update documentation as needed</li>
</ul>

<h2 id="license">📄 License</h2>
<p>MIT License - see <a href="https://github.com/welliardiansyah/wells-go/blob/master/LICENSE.md"</a> file for details.</p>
