package main

func main() {
	server, err := newServer()

	if err != nil {
		panic(err)
	}

	_ = server.Run(":8080")
}
