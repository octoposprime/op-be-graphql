package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/golobby/container/v3"

	pa_command "github.com/octoposprime/op-be-graphql/internal/application/presentation/adapter/command"
	pa_query "github.com/octoposprime/op-be-graphql/internal/application/presentation/adapter/query"
	as "github.com/octoposprime/op-be-graphql/internal/application/service"
	ds "github.com/octoposprime/op-be-graphql/internal/domain/service"
	ia_ebus "github.com/octoposprime/op-be-graphql/pkg/infrastructure/adapter/ebus"
	ia_repo "github.com/octoposprime/op-be-graphql/pkg/infrastructure/adapter/repository"
	ia_service "github.com/octoposprime/op-be-graphql/pkg/infrastructure/adapter/service"
	pc_graphql "github.com/octoposprime/op-be-graphql/pkg/presentation/controller/graphql"
	pc_probe "github.com/octoposprime/op-be-graphql/pkg/presentation/controller/probe"
	tseed "github.com/octoposprime/op-be-graphql/tool/config"
	tconfig "github.com/octoposprime/op-be-shared/tool/config"
	tredis "github.com/octoposprime/op-be-shared/tool/redis"
)

var internalConfig tconfig.InternalConfig
var redisConfig tconfig.RedisConfig
var seedConfig tseed.SeedConfig

func main() {
	internalConfig.ReadConfig()
	redisConfig.ReadConfig()
	seedConfig.ReadConfig()
	var err error

	fmt.Println("Starting Graphql Service...")
	redisClient := tredis.NewRedisClient(redisConfig.Redis.Host, redisConfig.Redis.Port, redisConfig.Redis.Password, redisConfig.Redis.Db)
	_, err = redisClient.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Redis")

	cont := container.New()

	//Domain Graphql Service
	err = cont.Singleton(func() *ds.Service {
		return ds.NewService()
	})
	if err != nil {
		panic(err)
	}

	//Infrastructure Graphql Micro Service Adapter
	err = cont.Singleton(func() ia_service.ServiceAdapter {
		return ia_service.NewServiceAdapter()
	})
	if err != nil {
		panic(err)
	}

	//Infrastructure Graphql Redis Repository Adapter
	err = cont.Singleton(func() ia_repo.RedisAdapter {
		return ia_repo.NewRedisAdapter(redisClient)
	})
	if err != nil {
		panic(err)
	}

	//Infrastructure Graphql EBus Adapter
	err = cont.Singleton(func() ia_ebus.EBusAdapter {
		return ia_ebus.NewEBusAdapter(redisClient)
	})
	if err != nil {
		panic(err)
	}

	//Application Graphql Service
	err = cont.Singleton(func(s *ds.Service, r ia_repo.RedisAdapter, serv ia_service.ServiceAdapter, e ia_ebus.EBusAdapter) *as.Service {
		return as.NewService(s, &r, &serv, &e)
	})
	if err != nil {
		panic(err)
	}

	//Application Graphql Query Adapter
	err = cont.Singleton(func(s *as.Service) pa_query.QueryAdapter {
		return pa_query.NewQueryAdapter(s)
	})
	if err != nil {
		panic(err)
	}

	//Application Graphql Command Adapter
	err = cont.Singleton(func(s *as.Service) pa_command.CommandAdapter {
		return pa_command.NewCommandAdapter(s)
	})
	if err != nil {
		panic(err)
	}

	var queryHandler pa_query.QueryAdapter
	err = cont.Resolve(&queryHandler)
	if err != nil {
		panic(err)
	}

	var commandHandler pa_command.CommandAdapter
	err = cont.Resolve(&commandHandler)
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	if !internalConfig.Local {
		wg.Add(1)
		go pc_probe.NewProbeAPI().Serve(internalConfig.Restapi.ProbePort)
	}
	wg.Add(1)
	go pc_graphql.NewGraphql(queryHandler, commandHandler).Serve(internalConfig.Restapi.GraphqlPort)
	wg.Wait()

}
