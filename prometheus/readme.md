## 一、安装

1、下载yaml文件

```shell
git clone https://github.com/coreos/kube-prometheus.git
```

2、安装

```shell
# 第一步
kubectl apply -f manifests/setup/

# 第二步
kubectl apply -f manifests/
```

3、查看资源

```shell
pdf@pg:~/kube-prometheus$ kubectl get po -n monitoring -o wide
NAME                                  READY   STATUS    RESTARTS   AGE     IP                NODE     NOMINATED NODE   READINESS GATES
alertmanager-main-0                   2/2     Running   0          5m57s   10.244.0.117      pg       <none>           <none>
alertmanager-main-1                   2/2     Running   0          5m57s   10.244.2.12       node01   <none>           <none>
alertmanager-main-2                   2/2     Running   0          5m57s   10.244.1.45       node00   <none>           <none>
grafana-58dc7468d7-n9w6f              1/1     Running   0          6m2s    10.244.0.116      pg       <none>           <none>
kube-state-metrics-78b46c84d8-bh8nm   3/3     Running   0          6m2s    10.244.1.44       node00   <none>           <none>
node-exporter-jbnph                   2/2     Running   0          6m1s    192.168.131.137   node01   <none>           <none>
node-exporter-l6x2c                   2/2     Running   0          6m1s    192.168.131.140   node00   <none>           <none>
node-exporter-xpmpz                   2/2     Running   0          6m1s    192.168.131.141   pg       <none>           <none>
prometheus-adapter-5cd5798d96-8msfw   1/1     Running   0          6m2s    10.244.2.11       node01   <none>           <none>
prometheus-k8s-0                      3/3     Running   1          5m47s   10.244.0.118      pg       <none>           <none>
prometheus-k8s-1                      3/3     Running   1          5m47s   10.244.1.46       node00   <none>           <none>
prometheus-operator-99dccdc56-qsv7b   1/1     Running   0          6m16s   10.244.0.115      pg       <none>           <none>

```

默认启动的 `SVC` 是 `ClusterIP` 形式，如果要用 `NodePort` 则需要手动修改



## 二、Prometheus

查看 `svc` :

```shell
pdf@pg:~$ kubectl get svc -n monitoring 
NAME                    TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                      AGE
alertmanager-main       ClusterIP   10.100.170.205   <none>        9093/TCP                     4h45m
alertmanager-operated   ClusterIP   None             <none>        9093/TCP,9094/TCP,9094/UDP   4h45m
grafana                 ClusterIP   10.102.12.102    <none>        3000/TCP                     4h45m
kube-state-metrics      ClusterIP   None             <none>        8443/TCP,9443/TCP            4h45m
node-exporter           ClusterIP   None             <none>        9100/TCP                     4h45m
prometheus-adapter      ClusterIP   10.96.67.144     <none>        443/TCP                      4h45m
prometheus-k8s          ClusterIP   10.104.132.91    <none>        9090/TCP                     4h45m
prometheus-operated     ClusterIP   None             <none>        9090/TCP                     4h44m
prometheus-operator     ClusterIP   None             <none>        8080/TCP                     4h45m

```

访问 `http://10.104.132.91:9090` 可以访问 `prometheus` 的 webUI 界面。访问 `http://10.104.132.91:9090/metrics` 可以看到普罗米修斯采集的数据。

 

