## k8s 学习

该篇项目是学习k8s的时候写的demo，涉及到 `deploy` `sts` `ingress` `service`  `ingress-nginx` `dashboard` `prometheus`

### ingress

ingress 区分namespace，并且使用 `ingress-nginx` 作为 `ingress-controller`。

运行 `ingress-nginx`:

```shell script
kubectl apply -f ingress-nginx.yaml
```

由于这里对官方的yaml文件做了更改，使用的是 `daemonset` 并且使用 `affinity.nodeAffinity` 将pod绑定在master节点。

运行 `ingress` :

```shell script
kubectl apply -f monitoring-ingress.yaml
```

访问 `local.prometheus.io` 即可访问 普罗米修斯 的首页。 grafana 同理。