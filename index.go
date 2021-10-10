package main

func main() {

	err := config.Init()
	if err != nil {
		fmt.Println("Failed to get application config: ", err.Error())
		log.Fatal("Failed to get application config: ", err.Error())
	}

	cfg := config.Get()
	
	redis.Init(cfg.Redis, cfg.Main.Server.CompressedCache)

	err = auth.Init(cfg)
	if err != nil {
		fmt.Println("Failed to init session: ", err.Error())
		tlog.Fatal("Failed to init session: ", err.Error())
	}

	err = database.Init(cfg.Database)
	if err != nil {
		fmt.Println("Failed to init database: ", err.Error())
		tlog.Fatal("Failed to init database: ", err.Error())
	}

	err = monitoring.Init(cfg.Main.Monitoring)
	if err != nil && util.GetEnv() != "development" {
		fmt.Println("Failed to init datadog: ", err.Error())
		tlog.Fatal("Failed to init datadog: ", err.Error())
	}

	module.Init()

	publisher.Init()

	monitoring.Run()

	monitoring.Done()
}