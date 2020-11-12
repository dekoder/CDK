package lib

import "fmt"

var (
	Plugins map[string]PluginInterface
)

func init() {
	Plugins = make(map[string]PluginInterface)
}

type PluginInterface interface {
	Desc() string
	Run() bool
}

func ListAllPlugin() {
	for name, plugin := range Plugins {
		fmt.Println(name, "\t", plugin.Desc())
	}
}

func RunSinglePlugin(name string) {
	Plugins[name].Run()
}

//每个插件在初始化时必须注册
func Register(name string, plugin PluginInterface) {
	Plugins[name] = plugin
}
