package evaluate

import (
	"golang.org/x/sys/unix"
	"log"
	"os"
	"os/user"
	"runtime"
	"strings"
)

func BasicSysInfo() {
	// current dir(pwd)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("current dir:", dir)

	// current user(id)
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("current user:", u.Username, "uid:", u.Uid, "gid:", u.Gid, "home:", u.HomeDir)

	// system os and arch
	log.Println("os:", runtime.GOOS, "arch:", runtime.GOARCH, "cpus:", runtime.NumCPU())

	// kernel version
	version, _ := unix.Sysctl("kern.osrelease")
	log.Println("kernel version:", strings.ToLower(version))
}
