package engine

import (
	"context"

	"github.com/Zyphria/zyphi/id"
	"github.com/Zyphria/zyphi/llm"
	"github.com/Zyphria/zyphi/logger"
	"github.com/Zyphria/zyphi/manager"
	"github.com/Zyphria/zyphi/options"
	"github.com/Zyphria/zyphi/stores"

	"gorm.io/gorm"
)

type Engine struct {
	options.RequiredFields

	ctx context.Context

	db *gorm.DB

	logger *logger.Logger

	ID   id.ID
	Name string

	// State management
	managers     []manager.Manager
	managerOrder []manager.ManagerID

	// stores
	actorStore   *stores.ActorStore
	sessionStore *stores.SessionStore

	interactionFragmentStore *stores.FragmentStore

	// LLM client
	llmClient *llm.LLMClient
}
