# CDK
Container Duck - Zero Dependency Docker/K8s Penetration Toolkit

# 介绍
本工具适用于攻入容器环境后的横向移动场景，解决以下问题：
  
1. 生产环境的容器是缩减后的linux系统，往往没有常用的linux命令和python等脚本环境，传统渗透工具无法使用。本工具提供golang实现的原生渗透工具集。
2. 集成docker/k8s场景特有的 逃逸、横向移动、持久化利用方式，插件化管理。
  
目前大部分功能仍在开发中，建议/反馈请提issue或mail: `i[at]cdxy.me`

# 功能

evaluate模块——容器内部弱点评估

```
cdk evaluate [--full]
```

|类别|功能|已支持|用例|
|---|---|---|---|
|信息收集|os基础指纹|✔||
|信息收集|容器内可用的capabilities|✔||
|信息收集|容器内可用的linux命令|✔||
|信息收集|mount到容器中的目录|✔||
|信息收集|env中的敏感服务|✔||
|信息收集|进程中的敏感服务|✔||
|信息收集|本地敏感文件扫描|✔||
|漏洞扫描|集成-版本比对|||

run模块——执行指定的脚本（插件化维护poc/exp）

```
cdk run --list
cdk run <script-name> [options]
```

|类别|功能|已支持|用例|
|---|---|---|---|
|逃逸|docker-runc CVE-2019-5736|||
|逃逸|dirtycow CVE-2016-5159|||
|逃逸|CVE-2017-7308|||
|逃逸|docker.sock逃逸|||
|逃逸|挂载device逃逸|✔||
|逃逸|cgroups逃逸|✔||
|逃逸|ptrace逃逸|||
|横向移动|本地K8s service account证书利用|||
|横向移动|K8s api-server本地未授权|||
|横向移动|etcd本地未授权|||
|信息窃取|代码库AK扫描|✔||
|信息窃取|K8s secrets dump|||
|信息窃取|K8s config dump|||
|持久化|webshell植入|||
|持久化|分发K8s后门Pod|||
|持久化|部署shadow K8s api-server|||
|持久化|部署K8s cronjob|||
|痕迹清理|K8s audit-log清理|||

工具命令——还原部分linux指令及常见的渗透工具

```
cdk nc [options]
cdk ps
```

|命令|功能|已支持|用例|
|---|---|---|---|
|nc|文件/shell通道|✔||
|ps|进程信息|✔||
|ifconfig|网络信息|✔||
|vi|文件编辑|✔||
|curl|HTTP发包|||
|kubectl|轻量级K8s管理|||
|probe|内网扫描|||
|tunnel|隧道|||

TODO

1. loader
2. 插件化
3. EDR检测与躲避