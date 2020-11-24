package lib

import (
	"fmt"
	"github.com/Xyntax/CDK/pkg/evaluate"
	"github.com/Xyntax/CDK/pkg/kubectl"
	"github.com/Xyntax/CDK/pkg/netcat"
	"github.com/Xyntax/CDK/pkg/network"
	"github.com/Xyntax/CDK/pkg/ps"
	"github.com/Xyntax/CDK/pkg/util"
	"github.com/Xyntax/CDK/pkg/vi"
	"github.com/docopt/docopt-go"
	"log"
	"os"
)

var Args map[string]interface{} // global for scripts to parse inner args

func ParseCmds() map[string]interface{} {
	usage := `
Container DucK  
zero-dependency docker/k8s exploit toolkit by <i@cdxy.me>

Find tutorial, configuration and use-case in https://github.com/Xyntax/CDK/wiki

Usage:
  cdk evaluate [--full]
  cdk run (--list | <exploit> [<args>...])
  cdk ps
  cdk ifconfig
  cdk nc [options]
  cdk vi <file>
  cdk kcurl [options]
  cdk ucurl [options]
  cdk -h | --help
  cdk -v | --version

Examples:
  cdk evaluate --full     							Run all information gathering scripts to find vulnerability inside container.
  cdk run --list									List all available exploits of docker/k8s.
  cdk run mount_cgroup "touch /tmp/exp_success"    	Automated escape privileged container then let target host run shell command.
  cdk vi /root/abk    								Edit files in container like "vi" command.
  cdk ps  											Show process information like "ps -ef" command.
  
Options:
  -h --help     Show this help msg.
  -v --version  Show version.
`
	ver := "cdk v0.1.3"
	arguments, _ := docopt.ParseArgs(usage, os.Args[1:], ver)
	return arguments
}

func PassInnerArgs() {
	os.Args = os.Args[1:]
}

func ParseDocopt() {
	Args = ParseCmds()
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

		fmt.Printf("\n[Information Gathering - K8s API Server]\n")
		evaluate.CheckK8sAnonymousLogin()

		fmt.Printf("\n[Information Gathering - K8s Service Account]\n")
		evaluate.CheckK8sServiceAccount()

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
	case "kcurl":
		PassInnerArgs()
		kubectl.KubectlMain()
	case "ucurl":
		PassInnerArgs()
		if len(os.Args) != 5 {
			log.Fatal("invalid input args, Example: ./cdk ucurl get /var/run/docker.sock http://127.0.0.1/info \"\"")
		}
		util.UnixHttpSend(os.Args[1], os.Args[2], os.Args[3], os.Args[4]) // test
	case "ifconfig":
		network.GetLocalAddresses()
	// use docopt to parse CDK original args
	case "ps":
		ps.RunPs()
	default:
		ParseDocopt()
	}

}
