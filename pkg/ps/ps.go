package ps

import (
	"fmt"
	gops "github.com/mitchellh/go-ps"
	"log"
)

func RunPs() {
	processList, err := gops.Processes()
	if err != nil {
		log.Fatal("ps.Processes() Failed, are you using windows?")
		return
	}
	for _, proc := range processList {
		fmt.Printf("%d\t%d\t%s\n", proc.Pid(), proc.PPid(), proc.Executable())
	}
}
