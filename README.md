# Zyphi - An Advanced Conversational LLM Framework

<div align="center">
  <img src="./img/logo3.jpeg" alt="Zyphria Banner" width="100%" />
</div>

## Table of Contents

- [Overview](#Overview)
- [Core Features](#core-features)
- [Extension Points](#extension-points)
- [Quick Start](#quick-start)
- [Using Zyphi as a Module](#using-zyphi-as-a-module)

## Overview

Zyphi is a highly modular AI conversation engine built in Go that emphasizes pluggable architecture and platform independence. It provides a flexible foundation for building conversational systems through:

- Plugin-based architecture with hot-swappable components
- Multi-provider LLM support (OpenAI, custom providers)
- Cross-platform conversation management
- Extensible manager system for custom behaviors
- Vector-based semantic storage with pgvector

## Core Features

### Plugin Architecture

- **Manager System**: Extend functionality through custom managers
  - Insight Manager: Extracts and maintains conversation insights
  - Personality Manager: Handles response behavior and style
  - Custom Managers: Add your own specialized behaviors

### State Management

- **Shared State System**: Centralized state management across components
  - Manager-specific data storage
  - Custom data injection
  - Cross-manager communication

### LLM Integration

- **Provider Abstraction**: Support for multiple LLM providers
  - Built-in OpenAI support
  - Extensible provider interface for custom LLMs
  - Configurable model selection per operation
  - Automatic fallback and retry handling

### Platform Support

- **Platform Agnostic Core**:
  - Abstract conversation engine independent of platforms
  - Built-in support for CLI chat and Twitter
  - Extensible platform manager interface
  - Example implementations for new platform integration

### Storage Layer

- **Flexible Data Storage**:
  - PostgreSQL with pgvector for semantic search
  - GORM-based data models
  - Customizable fragment storage
  - Vector embedding support

### Toolkit/Function System

- **Pluggable Tool/Function Integration**:
  - Support for custom tool implementations
  - Built-in toolkit management
  - Function calling capabilities
  - Automatic tool response handling
  - State-aware tool execution

## Extension Points

1. **LLM Providers**: Add new AI providers by implementing the LLM interface

```go
type Provider interface {
    GenerateCompletion(context.Context, CompletionRequest) (string, error)
    GenerateJSON(context.Context, JSONRequest, interface{}) error
    EmbedText(context.Context, string) ([]float32, error)
}
```

2. **Managers**: Create new behaviors by implementing the Manager interface

```go
type Manager interface {
    GetID() ManagerID
    GetDependencies() []ManagerID
    Process(state *state.State) error
    PostProcess(state *state.State) error
    Context(state *state.State) ([]state.StateData, error)
    Store(fragment *db.Fragment) error
    StartBackgroundProcesses()
    StopBackgroundProcesses()
    RegisterEventHandler(callback EventCallbackFunc)
    triggerEvent(eventData EventData)
}
```

## Quick Start

1. Clone the repository

```bash
git clone https://github.com/Zyphria/zyphi
```

2. Copy `.env.example` to `.env` and configure your environment variables
3. Install dependencies:

```bash
go mod download
```

4. Run the chat example:

```bash
go run examples/chat/main.go
```

6. Run the Twitter bot:

```bash
go run examples/twitter/main.go
```

## Environment Variables

```env
DB_URL=postgresql://user:password@localhost:5432/zyphi
OPENAI_API_KEY=your_openai_api_key

Platform-specific credentials as needed
```

## Architecture

The project follows a clean, modular architecture:

- `engine`: Core conversation engine
- `manager`: Plugin manager system
- `managers/*`: Built-in manager implementations
- `state`: Shared state management
- `llm`: LLM provider interfaces
- `stores`: Data storage implementations
- `tools/*`: Built-in tool implementations
- `examples/`: Reference implementations

## Using Zyphi as a Module

1. Add Zyphi to your Go project:

```bash
go get github.com/Zyphria/zyphi
```

2. Import Zyphi in your code:

```go
import (
  "github.com/Zyphria/zyphi/engine"
  "github.com/Zyphria/zyphi/llm"
  "github.com/Zyphria/zyphi/manager"
  "github.com/Zyphria/zyphi/managers/personality"
  "github.com/Zyphria/zyphi/managers/insight"
  ... etc
)
```

3. Basic usage example:

```go
// Initialize LLM client
llmClient, err := llm.NewLLMClient(llm.Config{
  ProviderType: llm.ProviderOpenAI,
  APIKey: os.Getenv("OPENAI_API_KEY"),
  ModelConfig: map[llm.ModelType]string{
    llm.ModelTypeDefault: openai.GPT4,
  },
  Logger: logger,
  Context: ctx,
})

// Create engine instance
engine, err := engine.New(
  engine.WithContext(ctx),
  engine.WithLogger(logger),
  engine.WithDB(db),
  engine.WithLLM(llmClient),
)

// Process input
state, err := engine.NewState(actorID, sessionID, "Your input text here")
if err != nil {
  log.Fatal(err)
}

response, err := engine.Process(state)
if err != nil {
  log.Fatal(err)
}
```

4. Available packages:

- `zyphi/engine`: Core conversation engine
- `zyphi/llm`: LLM provider interfaces and implementations
- `zyphi/manager`: Base manager system
- `zyphi/managers/*`: Built-in manager implementations
- `zyphi/state`: State management utilities
- `zyphi/stores`: Data storage implementations

For detailed examples, see the `examples/` directory in the repository.
