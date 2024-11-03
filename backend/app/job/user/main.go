package main

import (
	"backend/app/job/user/conf"
	"backend/app/job/user/logic"
	"flag"
	"github.com/zeromicro/go-zero/core/service"
	"log"
)

func main() {
	env := flag.String("env", "test", "Specify the environment (e.g., dev, test, online)")
	flag.Parse()

	// 加载配置
	config, err := conf.LoadConfig(*env)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	svc := service.NewServiceGroup()
	svc.Add(logic.NewUserRelationshipService(&config.UserRelevantCountService))
	svc.Start()

	defer svc.Stop()
}
