package presentation

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	pp_command "github.com/octoposprime/op-be-graphql/internal/application/presentation/port/command"
	pp_query "github.com/octoposprime/op-be-graphql/internal/application/presentation/port/query"
)

type Resolver struct {
	QueryHandler   pp_query.QueryPort
	CommandHandler pp_command.CommandPort
}
