# CDK
Container Duck - Penetration Tool for Docker &amp; K8s Environment

# functions

evaluate

* local sensitive file
* available linux commands
* available linux capabilities
* network information
* ssh keys
* known hosts
* docker.sock
* K8s service account
* K8s Pod infomation
* K8s API server
* cloud service AK/APIs
* image registry

attack/docker

* runc exploit(CVE)
* dirtycow exploit(CVE)
* container escape with docker.sock abuse
* container escape with mounted filepath
* container escape with linux capabilities

attack/k8s

* exploit privileged K8s service account
* exploit unauthorized etcd
* exploit unauthorized K8s api-server
* exploit unauthorized K8s dashboard

deploy

* deploy backdoor pod
* deploy ssh backdoor
* deploy webserver backdoor
* deploy K8s backdoor in MasterNode
* deploy K8s backdoor with K8s Cronjob
* dump K8s Secrets

network tool(nc)

* download
* upload
* reverse shell
* bind shell

editor(vi)

* vi 

K8s management(kubectl)

* exec kubectl commands

probe

* network IP/port scan