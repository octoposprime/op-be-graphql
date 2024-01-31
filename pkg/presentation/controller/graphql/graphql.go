package presentation

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	pp_command "github.com/octoposprime/op-be-graphql/internal/application/presentation/port/command"
	pp_query "github.com/octoposprime/op-be-graphql/internal/application/presentation/port/query"
	mo "github.com/octoposprime/op-be-graphql/internal/domain/model/object"
	resolver "github.com/octoposprime/op-be-graphql/pkg/presentation/controller/graphql/resolver"
	graph "github.com/octoposprime/op-be-graphql/pkg/presentation/dto"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb_user "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"
	tjwt "github.com/octoposprime/op-be-shared/tool/jwt"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Graphql is the GraphQL API for the application
type Graphql struct {
	commandHandler pp_command.CommandPort
	queryHandler   pp_query.QueryPort
}

// NewGraphql creates a new instance of Graphql
func NewGraphql(qh pp_query.QueryPort, ch pp_command.CommandPort) *Graphql {
	api := &Graphql{
		commandHandler: ch,
		queryHandler:   qh,
	}
	return api
}

// Serve starts the API server
func (api *Graphql) Serve(port string) {
	conf := graph.Config{Resolvers: &resolver.Resolver{
		CommandHandler: api.commandHandler,
		QueryHandler:   api.queryHandler,
	}}
	conf.Directives.Auth = api.Auth
	srv := handler.New(graph.NewExecutableSchema(conf))
	var mb int64 = 1 << 20
	srv.AddTransport(transport.MultipartForm{
		MaxMemory:     5 * mb,
		MaxUploadSize: 5 * mb,
	})

	srv.AddTransport(transport.SSE{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	srv.Use(extension.Introspection{})

	srv.SetErrorPresenter(MyErrorPresenter)
	http.Handle("/", corsMwFunc(playground.Handler("GraphQL playground", "/query")))
	http.Handle("/query", corsMwHandler(srv))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

var MyErrorPresenter graphql.ErrorPresenterFunc = func(ctx context.Context, err error) *gqlerror.Error {
	errGql := graphql.DefaultErrorPresenter(ctx, err)
	if strings.Contains(err.Error(), smodel.ErrBase+smodel.ErrSep) {
		errMsg := err.Error()[strings.LastIndex(err.Error(), smodel.ErrBase+smodel.ErrSep):]
		errGql = graphql.DefaultErrorPresenter(ctx, errors.New(errMsg))
	}
	return errGql
}

func corsMwHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func corsMwFunc(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h(w, r)
	}
}

func (api *Graphql) Auth(ctx context.Context, obj interface{}, next graphql.Resolver, policy string) (interface{}, error) {
	token := graphql.GetOperationContext(ctx).Headers.Get(string(smodel.HeaderKeyAuthorization))
	var err error
	var userId string
	var userType pb_user.UserType
	if token != "" {
		userId, _, err = tjwt.DecodeJWT(token)
		if err != nil {
			return nil, smodel.ErrorDecodeJwtFailed
		}
	} else {
		return nil, smodel.ErrorDecodeJwtFailed
	}
	if policy != mo.PermissionNone.Policy {
		users, err := api.queryHandler.GetUsersByFilter(ctx, &pb_user.UserFilter{Id: &userId})
		if err != nil {
			return nil, smodel.ErrorUserUnAuthorized
		}
		if len(users.Users) == 0 {
			return nil, smodel.ErrorUserUnAuthorized
		}
		isAuth := false
		for _, role := range mo.GetRoles() {
			if role.RoleName == users.Users[0].Role {
				for _, permissionGroup := range role.PermissionGroups {
					for _, permission := range permissionGroup.Permissions {
						if permission.Policy == policy {
							isAuth = true
							break
						}
					}
				}
				break
			}
		}
		if !isAuth {
			return nil, smodel.ErrorUserUnAuthorized
		}
		userType = users.Users[0].UserType
	}
	ctx = context.WithValue(ctx, smodel.QueryKeyUid, userId)
	ctx = context.WithValue(ctx, smodel.QueryKeyUType, userType)

	return next(ctx)
}
