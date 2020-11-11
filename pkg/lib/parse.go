package lib

import (
	"fmt"
	"github.com/Xyntax/CDK/pkg/exploit"
	"github.com/Xyntax/CDK/pkg/netcat"
	"github.com/Xyntax/CDK/pkg/network"
	"github.com/Xyntax/CDK/pkg/ps"
	"github.com/Xyntax/CDK/pkg/search"
	"github.com/Xyntax/CDK/pkg/ven"
	"github.com/docopt/docopt-go"
	"os"
)

func ParseCmds() map[string]interface{} {
	usage := `
Container Duck

The zero-dependency docker/k8s exploit toolkit.

Usage:
  cdk search [--full]
  cdk run (--list | <exploit> [options]) 
  cdk filescan <dir>
  cdk ps
  cdk ifconfig
  cdk nc [options]
  cdk vi <file>
  cdk kubectl [options]
  cdk -h | --help
  cdk -v | --version

Options:
  -h --help     Show this screen.
  -v --version  Show version.
`
	ver := "cdk v0.1.1"
	arguments, _ := docopt.ParseArgs(usage, os.Args[1:], ver)
	return arguments
}

func PassInnerArgs() {
	os.Args = os.Args[1:]
}

func ParseDocopt() {
	Args := ParseCmds()
	fmt.Println(Args)

	if Args["search"].(bool) {

		fmt.Printf("\n[Information Gathering - System Info]\n")
		search.BasicSysInfo()

		fmt.Printf("\n[Information Gathering - Services]\n")
		search.SearchSensitiveEnv()
		search.SearchSensitiveService()

		fmt.Printf("\n[Information Gathering - Commands and Capabilities]\n")
		search.SearchAvailableCommands()

		fmt.Printf("\n[Information Gathering - Mounts]\n")
		search.MountEscape()

		if Args["--full"].(bool) {
			fmt.Printf("\n[Information Gathering - Sensitive Files]\n")
			search.SearchLocalFilePath()
		}
	}
	if Args["run"].(bool) {
		if Args["--list"].(bool) {
			return
		}
		if Args["run"].(bool) {
			if Args["<exploit>"].(string) == "mount-escape" {
				//exploit.MountEscape()
			}
		}
	}
	if Args["filescan"].(bool) {
		StartDir := Args["<dir>"].(string)
		fmt.Printf("\n[Scan Secrets from Dir: %s]\n", StartDir)
		exploit.SearchLocalFileText(StartDir)
	}
}

func ParseArgsMain() {
	// parse tools first
	switch os.Args[1] {
	case "nc":
		PassInnerArgs()
		netcat.RunVendorNetcat()
	case "vi":
		PassInnerArgs()
		ven.RunVendorVen()
	//case "kubectl":
	//	PassInnerArgs()
	//	kubectl.RunKubectl()
	case "ifconfig":
		network.GetLocalAddresses()
	// use docopt to parse CDK original args
	case "ps":
		ps.RunPs()
	default:
		ParseDocopt()
	}

}
