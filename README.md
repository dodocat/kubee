# Kubee some helpful Kubernetes tools

## Installation

```
go get github.com/dodocat/kubee
```

set PATH if needed

```
export PATH=$PATH:$GOPATH/bin
```

## USAGE

run:

```
kubee -c config.yaml -n default # todo parse flags
```

output:

```
+---------------+---------------------------------+---------------------------------+---------+-------+---------+-------+
|   NAMESPACE   |           DEPLOYMENT            |            CONTAINER            |   CPU   |  CPU  |   MEM   |  MEM  |
|               |                                 |                                 | REQUEST | LIMIT | REQUEST | LIMIT |
+---------------+---------------------------------+---------------------------------+---------+-------+---------+-------+
| kube-system   | alicloud-application-controller | alicloud-application-controller |       0 |     0 |       0 |     0 |
| kube-system   | alicloud-disk-controller        | alicloud-disk-controller        |       0 |     0 |       0 |     0 |
| kube-system   | alicloud-monitor-controller     | alicloud-monitor-controller     |       0 |     0 |       0 |     0 |
| kube-system   | aliyun-acr-credential-helper    | aliyun-acr-credential-helper    |       0 |     0 |       0 |     0 |
| kube-system   | cluster-autoscaler              | cluster-autoscaler              |       1 |     1 |     300 |   300 |
| kube-system   | coredns                         | coredns                         |       1 |     0 |      70 |   170 |
| kube-system   | metrics-server                  | metrics-server                  |       0 |     0 |       0 |     0 |
| kube-system   | nginx-ingress-controller        | nginx-ingress-controller        |       0 |     0 |       0 |     0 |
| kube-system   | tiller-deploy                   | tiller                          |       0 |     0 |       0 |     0 |

```

## TODO

- [ ] add flag -c
- [ ] add flag -n
- [ ] sum resource usage
- [ ] show current status
