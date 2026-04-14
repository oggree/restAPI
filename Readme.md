# Oggree REST API

This module provides the foundation for the REST API server in the Oggree project using the [Echo web framework](https://echo.labstack.com/). It manages basic server initialization, robust middleware wrapping, standard application configuration logic (via the `config` wrapper), and unified HTTP response formatting.

## Core Components

### 1. `Init()`
Initializes the global Echo application instance (`Api`). It also registers critical routing paths and cross-cutting middleware.

- **Routes:** 
  - `GET /` and `GET /health`: Standard healthcheck endpoints that return a 200 OK `I'm Healthy!` text response.
  - `GET /swagger/*`: Serves the generated API documentation interface.
- **Middleware:**
  - **CORS:** A permissive configuration allowing cross-origin requests from any origin (`*`) with extensive method coverage.
  - **Logger:** Implements global request logging for observability.
  - **Recover:** Safely catches panics to prevent full process crashes. This is dynamically injected only if `config.GetString("env")` is set to `"production"`.

### 2. `Start()`
Pulls the target starting port from configuration (`config.GetString("port")`), defaults to `8080`, and runs the blocking Echo HTTP server loop.

### 3. Structured Responses
Exposes consistent structural formats to ensure REST clients always process a predictable schema.
- **`ResponseModel`**: A unified JSON payload structure.
  ```json
  {
    "status": true,
    "error": null,
    "data": { ... }
  }
  ```
- **`ResponseSuccessful(payload)`**: A quick utility function that wraps arbitrary payload data into a standardized successful `ResponseModel`.

## Configuration Properties
Uses the `config` package to manage the following application configuration bindings:
- `env`: Indicates the deployment environment. If `production`, enables panic-recovery mechanisms.
- `port`: Indicates the listening port for the web server (defaults to `"8080"`).

## Example API Start Usage
```go
package main

import (
	"github.com/oggree/config"
	"github.com/oggree/restAPI"
)

func main() {
    // Optional: Setup config values (depending on wrapper implementation)
    config.Set("port", "3000")
    config.Set("env", "development")
    
    // 1. Initialize the API
    restAPI.Init()
    
    // 2. Add sub-routes directly to the global instance
    // restAPI.Api.POST("/login", myLoginHandler)

    // 3. Start the blocking server
    restAPI.Start()
}
```
