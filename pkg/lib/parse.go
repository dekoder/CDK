package lib

import (
	"fmt"
	"github.com/Xyntax/CDK/pkg/evaluate"
	"github.com/Xyntax/CDK/pkg/netcat"
	"github.com/Xyntax/CDK/pkg/network"
	"github.com/Xyntax/CDK/pkg/ps"
	"github.com/Xyntax/CDK/pkg/vi"
	"github.com/docopt/docopt-go"
	"os"
)

var Args = ParseCmds()

func ParseCmds() map[string]interface{} {
	usage := `
Container Duck

The zero-dependency docker/k8s exploit toolkit.

Usage:
  cdk evaluate [--full]
  cdk run (--list | <exploit> [<args>...])
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
	//fmt.Println(Args)

	if Args["evaluate"].(bool) {

		fmt.Printf("\n[Information Gathering - System Info]\n")
		evaluate.BasicSysInfo()

		fmt.Printf("\n[Information Gathering - Services]\n")
		evaluate.SearchSensitiveEnv()
		evaluate.SearchSensitiveService()

		fmt.Printf("\n[Information Gathering - Commands and Capabilities]\n")
		evaluate.SearchAvailableCommands()
		evaluate.GetProcCapabilities()

		fmt.Printf("\n[Information Gathering - Mounts]\n")
		evaluate.MountEscape()

		if Args["--full"].(bool) {
			fmt.Printf("\n[Information Gathering - Sensitive Files]\n")
			evaluate.SearchLocalFilePath()
		}
	}
	if Args["run"].(bool) {
		if Args["--list"].(bool) {
			ListAllPlugin()
			os.Exit(0)
		}
		name := Args["<exploit>"].(string)
		if Plugins[name] == nil {
			fmt.Printf("\nInvalid script name: %s , available scripts:\n", name)
			ListAllPlugin()
			os.Exit(0)
		}
		RunSinglePlugin(name)
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
		vi.RunVendorVi()
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
