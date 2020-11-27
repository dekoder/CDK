package lib

import (
	"fmt"
	"github.com/Xyntax/CDK/conf"
	"github.com/Xyntax/CDK/pkg/evaluate"
	"github.com/Xyntax/CDK/pkg/kubectl"
	"github.com/Xyntax/CDK/pkg/netcat"
	"github.com/Xyntax/CDK/pkg/network"
	"github.com/Xyntax/CDK/pkg/nmap"
	"github.com/Xyntax/CDK/pkg/ps"
	"github.com/Xyntax/CDK/pkg/util"
	"github.com/Xyntax/CDK/pkg/vi"
	"github.com/docopt/docopt-go"
	"log"
	"os"
)

var Args map[string]interface{} // global for scripts to parse inner args

var usage = `
Container DucK
Zero-dependency docker/k8s penetration toolkit by <i@cdxy.me>
Find tutorial, configuration and use-case in https://github.com/Xyntax/CDK/wiki

Usage:
  cdk evaluate [--full]
  cdk run (--list | <exploit> [<args>...])
  cdk <tool> [<args>...]

Evaluate:
  cdk evaluate                              Gather information to find weekness inside container.
  cdk evaluate --full                       Enable file scan during information gathering.

Exploit:
  cdk run --list                            List all available exploits.
  cdk run <exploit> [<args>...]             Run single exploit, docs in https://github.com/Xyntax/CDK/wiki

Tool:
  vi <file>                                 Edit files in container like "vi" command.
  ps                                        Show process information like "ps -ef" command.
  nc [options]                              Create TCP tunnel.
  ifconfig                                  Show network information.
  kcurl	(get|post) <url> <data>             Make request to K8s api-server.
  ucurl (get|post) <socket> <uri> <data>    Make request to docker unix socket.

Options:
  -h --help     Show this help msg.
  -v --version  Show version.
`
var version = "cdk v0.1.4"

func ParseCmds() map[string]interface{} {
	arguments, _ := docopt.ParseArgs(usage, os.Args[1:], version)
	return arguments
}

func PassInnerArgs() {
	os.Args = os.Args[1:]
}

func ParseDocopt() {
	if len(os.Args) == 1 {
		docopt.PrintHelpAndExit(nil, usage)
	}
	Args = ParseCmds()

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
		evaluate.CheckK8sServiceAccount(conf.K8sSATokenDefaultPath)

		if Args["--full"].(bool) {
			fmt.Printf("\n[Information Gathering - Sensitive Files]\n")
			evaluate.SearchLocalFilePath()
		}
		return
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
			return
		}
		RunSinglePlugin(name)
		return
	}

	if Args["<tool>"] != nil {
		switch Args["<tool>"] {
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
		case "probe":
			args := Args["<args>"].([]string)
			if len(args) != 1 {
				log.Println("Invalid input args.")
				log.Fatal("usage: cdk probe 192.168.1.0-255")
			}
			nmap.TCPScanToolAPI(args[0])
		default:
			docopt.PrintHelpAndExit(nil, usage)
		}
	}
}
