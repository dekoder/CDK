# CDK - Zero Dependency Docker/K8s Penetration Toolkit

简体中文 ｜ [English](https://github.com/Xyntax/CDK/README_EN.md)

## 免责声明

未经授权许可使用CDK攻击目标是非法的。
本程序应仅用于安全测试与研究目的。


## 介绍

CDK是一款为容器环境定制的渗透测试工具，在已攻陷的容器内部提供零依赖的常用命令及PoC/EXP。集成Docker/K8s场景特有的 逃逸、横向移动、持久化利用方式，插件化管理。

目前仍在施工中，欢迎 [反馈](https://github.com/Xyntax/CDK/issues) 及建议 <i@cdxy.me>。 

## 功能

CDK包括三个功能模块

1. Evaluate: 容器内部信息收集，以发现潜在的弱点便于后续利用。
2. Exploit: 提供容器逃逸、持久化、横向移动等利用方式。
3. Tool: 修复渗透过程中常用的linux命令以及与Docker/K8s API交互的命令。

### Evaluate Module

Usage 
```
cdk evaluate [--full]
```
该命令默认将执行下表中除"本地文件扫描"以外的全部探测组件，使用`--full`参数开启本地文件扫描。

|阶段|组件|已支持|用例|
|---|---|---|---|
|本地信息收集|os基础指纹|✔|[link](https://github.com/Xyntax/CDK/wiki/Evaluate:-System-Info)|
|本地信息收集|容器内可用的capabilities|✔|[link](https://github.com/Xyntax/CDK/wiki/Evaluate:-Commands-and-Capabilities)|
|本地信息收集|容器内可用的linux命令|✔|[link](https://github.com/Xyntax/CDK/wiki/Evaluate:-Commands-and-Capabilities)|
|本地信息收集|mount到容器中的目录|✔|[link](https://github.com/Xyntax/CDK/wiki/Evaluate:-Mounts)|
|本地信息收集|env中的敏感信息|✔|[link](https://github.com/Xyntax/CDK/wiki/Evaluate:-Services)|
|本地信息收集|进程中的敏感服务|✔|[link](https://github.com/Xyntax/CDK/wiki/Evaluate:-Services)|
|本地信息收集|本地敏感文件|✔|[link](https://github.com/Xyntax/CDK/wiki/Evaluate:-Sensitive-Files)|
|网络探测|K8s api-server信息|✔|[link](https://github.com/Xyntax/CDK/wiki/Evaluate:-K8s-API-Server)|
|网络探测|K8s service-account信息|✔|[link](https://github.com/Xyntax/CDK/wiki/Evaluate:-K8s-Service-Account)|

### Exploit Module

列举全部exploit
```
cdk run --list
```

执行指定的exploit
```
cdk run <script-name> [options]
```

|阶段|组件|调用名|已支持|用例|
|---|---|---|---|---|
|逃逸|docker-runc CVE-2019-5736||||
|逃逸|dirtycow CVE-2016-5159||||
|逃逸|docker.sock PoC (DIND attack)|docker-sock-check|✔|[link](https://github.com/Xyntax/CDK/wiki/Exploit:-docker-sock-check)|
|逃逸|docker.sock 部署后门镜像|docker-sock-deploy|✔|[link](https://github.com/Xyntax/CDK/wiki/Exploit:-docker-sock-deploy)|
|逃逸|设备挂载逃逸|mount-disk|✔|[link](https://github.com/Xyntax/CDK/wiki/Exploit:-mount-disk)|
|逃逸|共享cgroups逃逸|mount-cgroup|✔|[link](https://github.com/Xyntax/CDK/wiki/Exploit:-mount-cgroup)|
|逃逸|挂载procfs逃逸|mount-procfs|✔|[link](https://github.com/Xyntax/CDK/wiki/Exploit:-mount-procfs)|
|逃逸|ptrace逃逸PoC|check-ptrace|✔|[link](https://github.com/Xyntax/CDK/wiki/Exploit:-check-ptrace)|
|横向移动|K8s Service Account投递指令||||
|横向移动|攻击K8s api-server||||
|横向移动|攻击K8s Kubelet||||
|横向移动|攻击K8s Dashboard||||
|横向移动|攻击K8s Helm||||
|横向移动|攻击K8s Etcd||||
|横向移动|攻击Docker私有镜像库||||
|远程控制|反弹Shell|reverse-shell|✔|[link](https://github.com/Xyntax/CDK/wiki/Exploit:-reverse-shell)|
|凭证窃取|AK泄露扫描|ak-leakage|✔|[link](https://github.com/Xyntax/CDK/wiki/Exploit:-ak-leakage)|
|凭证窃取|下载K8s Secrets||||
|凭证窃取|下载K8s Config||||
|持久化|部署WebShell||||
|持久化|部署后门K8s Pod||||
|持久化|部署影子K8s api-server||||
|持久化|部署K8s Cronbob||||
|检测躲避|清理K8s审计日志||||


### Tool Module

修复部分Linux命令，提供一些API访问工具，参数略有不同详见用例文档。

Usage
```
cdk <command> [options]
```

|命令|描述|已支持|用例|
|---|---|---|---|
|nc|TCP传输|✔||
|ps|进程信息|✔||
|ifconfig|网络信息|✔||
|vi|文件编辑|✔||
|kcurl|访问K8s api-server|||
|dcurl|访问Docker HTTP API|||
|ucurl|访问Docker Unix Socket|||
|rcurl|访问Docker Registry API|||
|probe|IP/端口扫描|||

## 下载
将可执行文件投递到已攻入的容器内部开始使用

https://github.com/Xyntax/CDK/tree/main/release

## TODO

1. echo load便于通过web RCE植入CDK
2. EDR检测对抗
3. 条件编译
4. 开发文档
