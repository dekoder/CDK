package main

import (
	goflag "flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/docopt/docopt-go"
	"github.com/spf13/pflag"

	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/kubectl/pkg/util/logs"
	"k8s.io/kubernetes/pkg/kubectl/cmd"

	// Import to initialize client auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

/*
Copyright 2014 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/


func RunKubectl() {
	os.Args = os.Args[1:]
	rand.Seed(time.Now().UnixNano())

	command := cmd.NewDefaultKubectlCommand()

	// TODO: once we switch everything over to Cobra commands, we can go back to calling
	// cliflag.InitFlags() (by removing its pflag.Parse() call). For now, we have to set the
	// normalize func and add the go flag set by hand.
	pflag.CommandLine.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	// cliflag.InitFlags()
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}


func parseCmds() map[string]interface{} {
	usage := `
Container Duck Project

The zero-dependency docker/k8s exploit toolkit

Usage:
  cdk run info/<SCRIPT> [options] 
  cdk run escape/<SCRIPT> [options] 
  x-fingerprint sslprobe [options] (--probe-file=<path_to_probe-file>) (--readtimes=READTIMES) (--input-file=<input-file>) (--output-file=<output-file>)
  x-fingerprint nmap [options] (--services-probe-file=<SERVICEs-PROBES-FILE>) (--input-file=<input-file>) (--output-file=<output-file>)
  x-fingerprint -h | --help
  x-fingerprint --version

Example
  cdk run info/all		run all scripts about local information gathering
  cdk run exploit/all   run all scripts about container escape and 

The most commonly used x-fingerprint grabber-types are:
  info       like nmap to load nmap-services-probes to scan
  exploit    Read the banners only when connected to the port-specified of target
  tool       Send the Probe payload file to grab the fingerprint

Require Options:
  -p=PORT, --port=PORT                         The Port-specified to grab fingerprint
  -M=PROTOCOL, --protocol=PROTOCOL             The protocol module to use [default: tcp] (Optionals: tcp/udp)
  --format=FORMAT                              The output-format (ASCII, JSON, Qpcode(Quoted-printable)) [default: qpcode]
  --workers=WORKERS                            The number of concurrent workers [default: 1000]
  --timeout=TIMEOUT                            Timeout for connection [default: 4]
  --input-file=INPUT-FILE                      Input file name [default: -] stdin
  --output-file=OUTPUT-FILE                    Output file name [default: -] stdout
  --readtimes=READTIMES                        The times of reading the socket-buffer[default: 2]
  --services-probe-file=SERVICE-PROBES-FILE    The serivces-probe-file

Info Module Options:
  --version-intensity=INTENSITY                The version intensity for nmap [default: 7]
  --send-timeout=SEND-TIMEOUT                  Send timeout for grab response [default: 7]
  --read-timeout=READ-TIMEOUT                  Read timeout version for grabbing [default: 7]
  --verbose=VERBOSE                            Verbose for log out details [default: 0]

Exploit Module Options:
  --script=SCRIPT                    The Nmap-script to specified
  --script-args=<n1=v1,[n2=v2,...]>  Provide the args to script-specified above

Tool Module Options:
  --probe-file=PROBE-FILE            To send the probe-payload to the port of target
`
	arguments, _ := docopt.Parse(usage, nil, true, "cdk v1.0", false)
	return arguments
}
func main() {
	//Args := parseCmds()
	if os.Args[1] == "kubectl" {
		fmt.Println(os.Args[2:])
		RunKubectl()
	}
	//f := new(xlib.Factory)
	//f.Initialize(Args)
	//f.Run()
}
