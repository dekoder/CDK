package evaluate

import (
	"fmt"
	"github.com/Xyntax/CDK/pkg/kubectl"
	"log"
	"strings"
)

func CheckK8sServiceAccount() {

	// get api-server connection conf in ENV
	addr := kubectl.ApiServerAddr()

	// check if we can login service-account with /run/secrets/kubernetes.io/serviceaccount/token
	log.Println("trying to login service-account with /run/secrets/kubernetes.io/serviceaccount/token")
	token, err := kubectl.GetServiceAccountToken()
	if err != nil {
		fmt.Println("\terr: ", err)
		return // exit this script
	}
	resp := kubectl.ServerAccountRequest(token, "get", addr+"/apis", "")
	if len(resp) > 0 && strings.Contains(resp, "APIGroupList") {
		fmt.Println("\tservice-account is available")

		// check if the current service-account can list namespaces
		log.Println("trying to list namespaces")
		resp := kubectl.ServerAccountRequest(token, "get", addr+"/api/v1/namespaces", "")
		if len(resp) > 0 && strings.Contains(resp, "kube-system") {
			fmt.Println("\tsuccess, the service-account have a high authority.")
			fmt.Println("\tnow you can make your own request to takeover the entire k8s cluster with `./cdk kcurl` command\n\tgood luck and have fun.")
		} else {
			fmt.Println("\tfailed")
			fmt.Println("\tresponse:" + resp)
		}
	} else {
		fmt.Println("\tservice-account is not available")
		fmt.Println("\tresponse:" + resp)
	}

}
