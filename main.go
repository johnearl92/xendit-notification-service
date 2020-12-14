package main

import "github.com/johnearl92/xendit-notification-service.git/cmd"

var version = "dev"

func main() {
	cmd.Execute(version)
}
