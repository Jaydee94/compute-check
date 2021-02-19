package app

import "fmt"

func Run() {
  fmt.Println("works")

  // initialize logging (maybe just stick to standard logger)

  // load config

  // setup server(s) + routes based on config
  // if we ever decide to server using TLS, we should setup
  // two server, one for business logic, and one for healthchecks etc

  // start server(s)

  // set readinessprobe to true

  // wait for SIGTERM or SIGKILL

  // set readinessprobe to false

  // wait some seconds, maybe we still receive new connections

  // gracefully terminate servers and exit

}
