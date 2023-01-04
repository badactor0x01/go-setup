
package main

import "fmt"
import log "github.com/sirupsen/logrus"

func main() {
    fmt.Println("Hello, World!")
	// Adding logging instead of printing.
	log.Info("Just an INFO log, no worries")
	log.Warn("A WARN log might make us a bit nervous...")
	log.Error("Now something is really wrong!")
}