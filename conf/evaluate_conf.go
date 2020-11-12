package conf

// check useful linux commands in container
var LinuxCommandChecklist = []string{
	"curl",
	"wget",
	"nc",
	"netcat",
	"kubectl",
	"docker",
	"find",
	"ps",
	"java",
	"python",
	"python3",
	"php",
	"node",
	"npm",
	"apt",
	"yum",
	"dpkg",
	"nginx",
	"httpd",
	"apache",
	"apache2",
	"ssh",
	"mysql",
	"mysql-client",
	"git",
	"svn",
	"vi",
	"capsh",
	"mount",
	"fdisk",
}

// match ENV to find useful service
var SensitiveEnvRegex = "(?i)\\bssh_|k8s|kubernetes|docker|gopath"

// match process name to find useful service
var SensitiveProcessRegex = "(?i)ssh|ftp|http|tomcat|nginx|engine|php|java|python|perl|ruby|kube|docker|\\bgo\\b"

// match local file path to find sensitive file
// walk starts from StartDir and match substring(AbsFilePath,<names in NameList>)
type sensitiveFileRules struct {
	StartDir string
	NameList []string
}

var SensitiveFileConf = sensitiveFileRules{
	StartDir: "/",
	NameList: []string{
		`docker.sock`,
		`.kube/`,
		`.git/`,
		`.svn/`,
		`.pip/`,
		`/.bash_history`,
		`/.bash_profile`,
		`/.bashrc`,
		`/.ssh/`,
		`.token`,
		`/serviceaccount`,
		`.dockerenv`,
		`/config.json`,
	},
}
