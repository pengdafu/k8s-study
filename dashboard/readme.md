
### 生成证书

```shell script
openssl req -x509 -nodes -days 1000 -newkey rsa:2048 -keyout kube-dashboard.key -out kube-dashboard.crt -subj="/CN=k8s.dashboard.cn/O=k8s.dashboard.cn"
```

### 生成 secret

```shell script
kubectl create secret tls kube-dashboard-ssl -n kubernetes-dashboard --key kube-dashboard.key --cert kube-dashboard.crt
```

### 获取登录 token

```shell script
kubectl -n kubernetes-dashboard describe secret $(kubectl -n kubernetes-dashboard get secret | grep ui| awk '{print $1}')
```