package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", HandleRoot)
	server.Handle("/api", server.AddMiddleware(HandleApi, CheckAuth(), Logging()))
	_ = server.Listen()
}
