[redis](../stateful-set/redis-single.yaml) 已经部署

## 部署过程

postgresql -> gitlab -> gitlab-runner

### 部署 postgresql
```shell script
kubectl apply -f Postgresql.yml
```

### 部署 gitlab

```shell script
kubectl apply -f gitlab.yaml
```

### 部署 gitlab-runner

> 注意,gitlab-runner 采用的是 helm 部署, 需要先安装 helm

```shell script
helm install --name gitlab-runner --namespace devops .
```

