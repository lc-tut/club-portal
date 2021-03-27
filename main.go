package main

func main() {
	logger, err := newZapLogger()

	if err != nil {
		panic(err)
	}

	engine, err := newGinEngine(logger)

	if err != nil {
		panic(err)
	}

	_ = engine.Run(":8080")
}
