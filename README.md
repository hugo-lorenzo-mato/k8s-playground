# K8S-Playground

<p align="center">
:construction: Under construction :construction:
</p>

## Motivation <a name="motivation"></a>

This project is only for learning purposes. Feel free to use the code provided here, but make sure that it suits your 
goal. On many occasions, the fastest way to perform certain operations is through the OC CLI itself with a little 
shell scripting.

This repository is intended to test the k8s and ocp libraries and also to be used for other tasks where the previous 
advice does not apply for various reasons (one may be the lack of permissions on the resources of the entire cluster).

Also, for many operations, being able to implement the logic in Go and not shell, gives much more flexibility when 
it grows.

## Table of Content

1. [Motivation](#motivation)
1. [Usage](#usage)   
1. [CLI Commands](#clicommands)  
   1. [networking](#networking)  
      1. [checkConnection](#checkConnection)  
   1. [networkpolicy](#networkpolicy)  
      1. [backupAll](#backupAll)  
   1. [route](#route)  
      1. [getNamespace](#getNamespace)    
   1. [cluster](#cluster)  
      1. [getNamespace](#login)  
1. [Changelog](#changelog)  

## Usage <a name="usage"></a>

1. Clone the project: `git clone`
2. Use makefile: `make compile` or `make help` if you want to test another options. The makefile compiles for Linux, 
   Windows and BSD. Cross-platform compilation is also implemented.

```shell
➜  k8s-playground git:(main) ✗ make help

Choose a command run in k8s-playground:

  crosscompile   cross platform compilation.
  compile        execute compilation for the current platform and architecture.
```   

3. Execute the CLI: `./bin/k8s-playground --help`


```shell
➜  k8s-playground git:(main) ✗ ./bin/k8s-playground --help                                                                                                           

██╗  ██╗ █████╗ ███████╗      ██████╗ ██╗      █████╗ ██╗   ██╗ ██████╗ ██████╗  ██████╗ ██╗   ██╗███╗   ██╗██████╗ 
██║ ██╔╝██╔══██╗██╔════╝      ██╔══██╗██║     ██╔══██╗╚██╗ ██╔╝██╔════╝ ██╔══██╗██╔═══██╗██║   ██║████╗  ██║██╔══██╗
█████╔╝ ╚█████╔╝███████╗█████╗██████╔╝██║     ███████║ ╚████╔╝ ██║  ███╗██████╔╝██║   ██║██║   ██║██╔██╗ ██║██║  ██║
██╔═██╗ ██╔══██╗╚════██║╚════╝██╔═══╝ ██║     ██╔══██║  ╚██╔╝  ██║   ██║██╔══██╗██║   ██║██║   ██║██║╚██╗██║██║  ██║
██║  ██╗╚█████╔╝███████║      ██║     ███████╗██║  ██║   ██║   ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║ ╚████║██████╔╝
╚═╝  ╚═╝ ╚════╝ ╚══════╝      ╚═╝     ╚══════╝╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝╚═════╝ 
                                                                                                   Version-0.3.0
NAME:
   k8s-playground - This CLI brings together some personal tests using the K8S and OCP libraries.

USAGE:
   k8s-playground [global options] command [command options] [arguments...]

VERSION:
   0.3.0

COMMANDS:
   networkpolicy  networkpolicy related actions
   networking     networking related actions
   route          Route related actions
   version        Displays K8S-Playground CLI version
   help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
INFO[0000] Cli Command Execution took 518.154µs
```

The connection info yaml file should be formed as follows:

```yaml
openshift:
  des:
    myclusterdes1:
      url: cluster-uri:port
      user: username
      password: password
    myclusterdes2:
      url: cluster-uri:port
      user: username
      password: password
  pre:
    myclusterpre:
      url: cluster-uri:port
      user: username
      password: password
  pro:
    myclusterpro:
      url: cluster-uri:port
      user: username
      password: password
```

The authentication and client creation model can be adapted to another incluster type depending on the needs or 
the part of the library that each person wants to test.

## CLI Commands <a name="clicommands"></a>

#### networking <a name="networking"></a>

##### checkConnection <a name="checkConnection"></a>

Passing a project or namespace source and another target to the CLI, check if there is connectivity between them.

The goal is not so much to test the connectivity as to test the elements involved: consult the list of services of a
given project, access its specification, list the pods to identify a running one on which to launch the command, etc.


```shell
NAME:
   k8s-playground networking checkConnection - Backup of all the networkpolicies of a cluster will be created.

USAGE:
   k8s-playground networking checkConnection [command options] [arguments...]

OPTIONS:
   --cluster value                 Cluster to iterate through. E.G: mycluster1 [$CLUSTER]
   --supraenvironment value        Supraenvironment. E.G: des [$SUPRAENVIRONMENT]
   --connectionDataYamlPath value  Path to the yaml file containing the credentials and URIs to connect to the different clusters. E.G: example.yml (default: "data/cluster.yml") [$CREDENTIALS_YAML_PATH]
   --namespaceSource value         Namespace/project source of the connection request. [$NAMESPACE_SOURCE]
   --namespaceTarget value         Namespace/project target of the connection request. [$NAMESPACE_TARGET]
   --help, -h                      show help (default: false)
```

Output example:


```shell
➜  k8s-playground git:(main) ✗ ./bin/k8s-playground networking checkConnection --cluster mycluster --supraenvironment pre --namespaceSource ns-example-pre --namespaceTarget ns-exampletarget-pre

██╗  ██╗ █████╗ ███████╗      ██████╗ ██╗      █████╗ ██╗   ██╗ ██████╗ ██████╗  ██████╗ ██╗   ██╗███╗   ██╗██████╗ 
██║ ██╔╝██╔══██╗██╔════╝      ██╔══██╗██║     ██╔══██╗╚██╗ ██╔╝██╔════╝ ██╔══██╗██╔═══██╗██║   ██║████╗  ██║██╔══██╗
█████╔╝ ╚█████╔╝███████╗█████╗██████╔╝██║     ███████║ ╚████╔╝ ██║  ███╗██████╔╝██║   ██║██║   ██║██╔██╗ ██║██║  ██║
██╔═██╗ ██╔══██╗╚════██║╚════╝██╔═══╝ ██║     ██╔══██║  ╚██╔╝  ██║   ██║██╔══██╗██║   ██║██║   ██║██║╚██╗██║██║  ██║
██║  ██╗╚█████╔╝███████║      ██║     ███████╗██║  ██║   ██║   ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║ ╚████║██████╔╝
╚═╝  ╚═╝ ╚════╝ ╚══════╝      ╚═╝     ╚══════╝╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝╚═════╝ 
                                                                                                   Version-0.2.0
INFO[0000]      ---------------------------- flags values ---------------------------- 
INFO[0000]      Key: nsOrigin                  --> ns-example-pre                    
INFO[0000]      Key: nsTarget                  --> ns-exampletarget-pre                
INFO[0000]      Key: dryRun                    --> false                               
INFO[0000]      Key: cluster                   --> mycluster                     
INFO[0000]      Key: supraenvironment          --> pre                                 
INFO[0000]      Key: connectionDataYamlPath    --> data/cluster.yml                    
INFO[0000]      ---------------------------------------------------------------------- 
INFO[0000] + command: curl -o /dev/null --max-time 3 -ks my-service.ns-exampletarget-pre.svc.cluster.local:8080 && echo Connection is available from $(uname -n) in ns-example-pre namespace to my-service.ns-exampletarget-pre.svc.cluster.local:8080 || echo Connection is NOT available from $(uname -n) in ns-example-pre namespace to my-service.ns-exampletarget-pre.svc.cluster.local:8080 
INFO[0001] Output: Connection is available from mypod-7cfd88c87-lc5gd in ns-example-pre namespace to my-svc.ns-exampletarget-pre.svc.cluster.local:8080

```

#### networkpolicy <a name="networkpolicy"></a>

##### backupAll <a name="backupAll"></a>

Very basic example to show the clients process creation and how to iterate with the data returned by them. The command 
backs up all existing networkpolicies in the cluster by iterating through all projects (--all-namespace flag not 
available).

```shell
➜  k8s-playground ./bin/k8s-playground-linux networkpolicy backupAll --help

██╗  ██╗ █████╗ ███████╗      ██████╗ ██╗      █████╗ ██╗   ██╗ ██████╗ ██████╗  ██████╗ ██╗   ██╗███╗   ██╗██████╗ 
██║ ██╔╝██╔══██╗██╔════╝      ██╔══██╗██║     ██╔══██╗╚██╗ ██╔╝██╔════╝ ██╔══██╗██╔═══██╗██║   ██║████╗  ██║██╔══██╗
█████╔╝ ╚█████╔╝███████╗█████╗██████╔╝██║     ███████║ ╚████╔╝ ██║  ███╗██████╔╝██║   ██║██║   ██║██╔██╗ ██║██║  ██║
██╔═██╗ ██╔══██╗╚════██║╚════╝██╔═══╝ ██║     ██╔══██║  ╚██╔╝  ██║   ██║██╔══██╗██║   ██║██║   ██║██║╚██╗██║██║  ██║
██║  ██╗╚█████╔╝███████║      ██║     ███████╗██║  ██║   ██║   ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║ ╚████║██████╔╝
╚═╝  ╚═╝ ╚════╝ ╚══════╝      ╚═╝     ╚══════╝╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝╚═════╝ 
                                                                                                   Version-0.1.0
NAME:
   k8s-playground-linux networkpolicy backupAll - Backup of all the networkpolicies of a cluster will be created.

USAGE:
   k8s-playground-linux networkpolicy backupAll [command options] [arguments...]

OPTIONS:
   --cluster value                 Cluster to iterate through. E.G: micluster [$CLUSTER]
   --supraenvironment value        Supraenvironment. E.G: des [$SUPRAENVIRONMENT]
   --connectionDataYamlPath value  Path to the yaml file containing the credentials and URIs to connect to the different clusters. E.G: example.yml (default: "data/cluster.yml") [$CREDENTIALS_YAML_PATH]
   --dryRun, -d                    dryRun execution (default: false) [$DRY_RUN]
   --help, -h                      show help (default: false)
   
INFO[0000] Cli Command Execution took 527.735µs
```

For each project or namespace it generates a new directory within TARGET and for each networkpolicy, a yaml file with 
its definition.

```shell
➜  k8s-playground git:(main) ✗ ls -lha TARGET/project-des-test/
total 44K
drwxrwxr-x   2 vagrant vagrant 4.0K Aug  1 15:32 .
drwxrwxr-x 591 vagrant vagrant  24K Aug  1 15:34 ..
-rwxrwxr-x   1 vagrant vagrant  999 Aug  1 15:32 allow-from-ingress-namespace-1627824762.yml
-rwxrwxr-x   1 vagrant vagrant  933 Aug  1 15:32 allow-from-same-namespace-1627824762.yml
```

It shows via log how many projects were analyzed and network policies were exported to files

```shell
INFO[0116] -> Number of projects analyzed: 589. Total number of Network Policies: 4815. 
INFO[0116] Cli Command Execution took 1m56.572868891s 
```

#### route <a name="route"></a>

##### getNamespace <a name="getNamespace"></a>

Given a host and a path, prints the namespace/project name where route is applied. As I show in the example, it can be 
useful to find duplicate routes in the cluster.

![image](https://user-images.githubusercontent.com/22169531/127917820-d4c47adf-42f6-44e6-a83c-01de85b7ace0.png)

Shows the status so that, in the case of duplicates, the one that is active can be identified.


```shell
NAME:
   k8s-playground route checkDuplicates - Check for duplicates for a specific path.

USAGE:
   k8s-playground route checkDuplicates [command options] [arguments...]

OPTIONS:
   --cluster value                 Cluster to iterate through. E.G: mycluster [$CLUSTER]
   --connectionDataYamlPath value  Path to the yaml file containing the credentials and URIs to connect to the different clusters. E.G: example.yml (default: "data/cluster.yml") [$CREDENTIALS_YAML_PATH]
   --host value                    Route host without protocol. E.G: myruta.ruta.com/custom-path [$host]
   --path value                    Route Path.E.G. custom-path [$path]
   --supraenvironment value        Supraenvironment. E.G: des [$SUPRAENVIRONMENT]
   --help, -h                      show help (default: false)
   
INFO[0000] Cli Command Execution took 823.046µs  
```

Output example:


```shell
➜  k8s-playground git:(main) ✗ ./bin/k8s-playground route getNamespace --cluster my-cluster --supraenvironment pre --host 'test-path.com' --path '/test-path'
██╗  ██╗ █████╗ ███████╗      ██████╗ ██╗      █████╗ ██╗   ██╗ ██████╗ ██████╗  ██████╗ ██╗   ██╗███╗   ██╗██████╗ 
██║ ██╔╝██╔══██╗██╔════╝      ██╔══██╗██║     ██╔══██╗╚██╗ ██╔╝██╔════╝ ██╔══██╗██╔═══██╗██║   ██║████╗  ██║██╔══██╗
█████╔╝ ╚█████╔╝███████╗█████╗██████╔╝██║     ███████║ ╚████╔╝ ██║  ███╗██████╔╝██║   ██║██║   ██║██╔██╗ ██║██║  ██║
██╔═██╗ ██╔══██╗╚════██║╚════╝██╔═══╝ ██║     ██╔══██║  ╚██╔╝  ██║   ██║██╔══██╗██║   ██║██║   ██║██║╚██╗██║██║  ██║
██║  ██╗╚█████╔╝███████║      ██║     ███████╗██║  ██║   ██║   ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║ ╚████║██████╔╝
╚═╝  ╚═╝ ╚════╝ ╚══════╝      ╚═╝     ╚══════╝╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝╚═════╝ 
                                                                                                   Version-0.3.0
INFO[0000]      ---------------------------- flags values ---------------------------- 
INFO[0000]      Key: supraenvironment          --> pre                                 
INFO[0000]      Key: cluster                   --> mycluster                     
INFO[0000]      Key: connectionDataYamlPath    --> data/cluster.yml                    
INFO[0000]      Key: host                      --> test-path.com                       
INFO[0000]      Key: path                      --> /test-path                          
INFO[0000]      ---------------------------------------------------------------------- 
INFO[0007] + Host and path found in my-ns-pre project: Host: test-path.com Path: /test-path Status: &RouteStatus{Ingress:[{test-path.com my [{Admitted True   2021-08-02 13:45:37 +0200 CEST}] None}],} 
INFO[0007] + Host and path found in my-ns-pre project: Host: test-path.com Path: /test-path Status: &RouteStatus{Ingress:[{test-path.com my [{Admitted False HostAlreadyClaimed route test-duplicate already exposes test-path.com and is older 2021-08-02 13:46:24 +0200 CEST}] None }],} 
INFO[0116] -> Number of projects analyzed: 590. Total number of Routes processed: 4316. 
INFO[0116] Cli Command Execution took 1m56.517531912s 
```

#### Cluster <a name="cluster"></a>

##### Login <a name="login"></a>

The login command converts the connection data from data/cluster.yml into the command needed to login to the different 
clusters.

```shell
PS [..]\GolandProjects\k8s-playground> .\k8s-playground.exe cluster login

██╗  ██╗ █████╗ ███████╗      ██████╗ ██╗      █████╗ ██╗   ██╗ ██████╗ ██████╗  ██████╗ ██╗   ██╗███╗   ██╗██████╗
██║ ██╔╝██╔══██╗██╔════╝      ██╔══██╗██║     ██╔══██╗╚██╗ ██╔╝██╔════╝ ██╔══██╗██╔═══██╗██║   ██║████╗  ██║██╔══██╗
█████╔╝ ╚█████╔╝███████╗█████╗██████╔╝██║     ███████║ ╚████╔╝ ██║  ███╗██████╔╝██║   ██║██║   ██║██╔██╗ ██║██║  ██║
██╔═██╗ ██╔══██╗╚════██║╚════╝██╔═══╝ ██║     ██╔══██║  ╚██╔╝  ██║   ██║██╔══██╗██║   ██║██║   ██║██║╚██╗██║██║  ██║
██║  ██╗╚█████╔╝███████║      ██║     ███████╗██║  ██║   ██║   ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║ ╚████║██████╔╝
╚═╝  ╚═╝ ╚════╝ ╚══════╝      ╚═╝     ╚══════╝╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝╚═════╝
                                                                                                   Version-0.3.0
+------------------+---------------------+-----------------------------------------------------------------------------------------------------------------------------------------+
SUPRAENVIRONMENT        CLUSTER                 LOGIN
+------------------+---------------------+-----------------------------------------------------------------------------------------------------------------------------------------+
| des              | micluster1          | oc login --insecure-skip-tls-verify --username xxxxxxxxxxxxxxxx --password xxxxxxxxxxxx micluster.com:443                               |
+                  +---------------------+                                                                                                                                         +
|                  | micluster2          |                                                                                                                                         |
+                  +---------------------+                                                                                                                                         +
|                  | micluster3          |                                                                                                                                         |
+                  +---------------------+                                                                                                                                         +
|                  | micluster4          |                                                                                                                                         |
+                  +---------------------+                                                                                                                                         +
|                  | micluster5          |                                                                                                                                         |
+                  +---------------------+                                                                                                                                         +
|                  | micluster6          |                                                                                                                                         |
+                  +---------------------+-----------------------------------------------------------------------------------------------------------------------------------------+
|                  | micluster7          | oc login --insecure-skip-tls-verify --username xxxxxxxxxxxxxxxx --password xxxxxxxxxxxx micluster.com:443                               |
+------------------+---------------------+-----------------------------------------------------------------------------------------------------------------------------------------+
| pre              | micluster1          | oc login --insecure-skip-tls-verify --username xxxxxxxxxxxxxxxx --password xxxxxxxxxxxxxxxx micluster.com:443                           |
+                  +---------------------+                                                                                                                                         +
|                  | micluster2          |                                                                                                                                         |
+                  +---------------------+                                                                                                                                         +
|                  | micluster3          |                                                                                                                                         |
+                  +---------------------+                                                                                                                                         +
|                  | micluster4          |                                                                                                                                         |
+                  +---------------------+                                                                                                                                         +
|                  | micluster5          |                                                                                                                                         |
+                  +---------------------+-----------------------------------------------------------------------------------------------------------------------------------------+
|                  | micluster6          | oc login --insecure-skip-tls-verify --username xxxxxxxxxxxxxxxx --password xxxxxxxxxxxxxxxx micluster.com:443                           |
+                  +---------------------+-----------------------------------------------------------------------------------------------------------------------------------------+
|                  | micluster7          | oc login --insecure-skip-tls-verify --username xxxxxxxxxxxxxxxx --password xxxxxxxxxxxxxxxx micluster.com:443                           |
+                  +---------------------+-----------------------------------------------------------------------------------------------------------------------------------------+
|                  | micluster8          | oc login --insecure-skip-tls-verify --username xxxxxxxxxxxxx --password xxxxxxxxxxxxx micluster.com:443                                 |
+                  +---------------------+                                                                                                                                         +
|                  | micluster9          |                                                                                                                                         |
+------------------+---------------------+-----------------------------------------------------------------------------------------------------------------------------------------+
| pro              | micluster1          | oc login --insecure-skip-tls-verify --username xxxxxxxxxxxxx --password xxxxxxxxxxxxx micluster.com:443                                 |
+                  +---------------------+                                                                                                                                         +
|                  | micluster2          |                                                                                                                                         |
+                  +---------------------+                                                                                                                                         +
|                  | micluster3          |                                                                                                                                         |
+                  +---------------------+                                                                                                                                         +
|                  | micluster4          |                                                                                                                                         |
+                  +---------------------+                                                                                                                                         +
|                  | micluster5          |                                                                                                                                         |
+                  +---------------------+-----------------------------------------------------------------------------------------------------------------------------------------+
|                  | micluster6          | oc login --insecure-skip-tls-verify --username xxxxxxxxxxxxx --password xxxxxxxxxxxxx micluster.com:443                                 |
+                  +---------------------+-----------------------------------------------------------------------------------------------------------------------------------------+
|                  | micluster7          | oc login --insecure-skip-tls-verify --username xxxxxxxxxxxxx --password xxxxxxxxxxxxx micluster.com:443                                 |
+                  +---------------------+                                                                                                                                         +
|                  | micluster7          |                                                                                                                                         |
+                  +---------------------+-----------------------------------------------------------------------------------------------------------------------------------------+
|                  | micluster8          | oc login --insecure-skip-tls-verify --username xxxxxxxxxxxxx --password xxxxxxxxxxxxx micluster.com:443                                 |
+                  +---------------------+-----------------------------------------------------------------------------------------------------------------------------------------+
|                  | micluster9          | oc login --insecure-skip-tls-verify --username xxxxxxxxxxxxx --password xxxxxxxxxxxxx micluster.com:443                                 |
+                  +---------------------+-----------------------------------------------------------------------------------------------------------------------------------------+
|                  | micluster10         | oc login --insecure-skip-tls-verify --username xxxxxxxxxxxxx --password xxxxxxxxxxxxx micluster.com:443                                 |
+                  +---------------------+-----------------------------------------------------------------------------------------------------------------------------------------+
|                  | micluster11         | oc login --insecure-skip-tls-verify --username xxxxxxxxxxxxx --password xxxxxxxxxxxxx micluster.com:443                                 |
+                  +---------------------+-----------------------------------------------------------------------------------------------------------------------------------------+
|                  | micluster12         | oc login --insecure-skip-tls-verify --username xxxxxxxxxxxxx --password xxxxxxxxxxxxx micluster.com:443                                 |
+                  +---------------------+-----------------------------------------------------------------------------------------------------------------------------------------+
|                  | micluster13         | oc login --insecure-skip-tls-verify --username xxxxxxxxxxxxx --password xxxxxxxxxxxxx micluster.com:443                                 |
+                  +---------------------+-----------------------------------------------------------------------------------------------------------------------------------------+
|                  | micluster14         | oc login --insecure-skip-tls-verify --username xxxxxxxxxxxxx --password xxxxxxxxxxxxx micluster.com:443                                 |
+                  +---------------------+-----------------------------------------------------------------------------------------------------------------------------------------+
|                  | micluster15         | oc login --insecure-skip-tls-verify --username xxxxxxxxxxxxx --password xxxxxxxxxxxxx micluster.com:443                                 |
+                  +---------------------+-----------------------------------------------------------------------------------------------------------------------------------------+
|                  | micluster16         | oc login --insecure-skip-tls-verify --username xxxxxxxxxxxxx --password xxxxxxxxxxxxx micluster.com:443                                 |
+------------------+---------------------+-----------------------------------------------------------------------------------------------------------------------------------------+
```

## Changelog <a name="changelog"></a>

### v.0.4.0

- Add cluster login new command: `k8s-playground cluster login`.

### v.0.3.0

- Add getNamespace in Route command: `k8s-playground route checkDuplicates`.
- Add route package.

### v.0.2.0

- Add checkConnection command: `k8s-playground networking checkConnection`.

### v.0.1.0

- Implement CLI with [Urfave v2 library](https://github.com/urfave/cli/blob/master/docs/v2/manual.md).
- Add networkpolicies backup command example.
- Add project, service, pod amd networkpolicy packages with basic functionality.  
- Add util package.
- Add Readme.

