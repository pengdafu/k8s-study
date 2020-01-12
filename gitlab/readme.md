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

### 生成 gitlab 操作k8s集群需要的token

```shell script
kubectl apply -f gitlab-sa.yaml
```

获取secrets的名字
```shell script
kubectl get sa gitlab-sa -n devops -o json|jq -r '.secrets[0].name'
gitlab-sa-token-57qq7
```

获取 ca.key
```shell script
kubectl get secrets -n devops gitlab-sa-token-57qq7 -o json|jq -r '.data["ca.crt"]'|base64 -d

-----BEGIN CERTIFICATE-----
MIICyDCCAbCgAwIBAgIBADANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwprdWJl
cm5ldGVzMB4XDTE5MTEwNjE0Mjg0NVoXDTI5MTEwMzE0Mjg0NVowFTETMBEGA1UE
AxMKa3ViZXJuZXRlczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALUU
3G9iox3TKJM5Dfpb3rOnEkAZXbzGIaLxXr1jaB7NsxRvlX0tHSIzgZ7ZdDAKE/qh
0BLcb0AF9scenRr5JOw0tdadDi310I7JS9b1JFXah8Bj5QxW67chK/U3rWraN78l
v51GFoo21d2yWH8ZMnZGEYrejCI88EXaaCQeJbr6PDcjLy31tR3CnpT8EVBjxkDn
pzM7zelU4NmWjCHibUTq1N2PrHvrO5cnFlvKWvePnB98U6m4zualz4e7JVRANDUS
tjZL1E2FnDENI049q5CnFcRzs2Y9BSP9aOTLo0ZL6rLkMvrYNTGZ2Su3XsLNqG0L
P+9Xr3INQxQ9pqTZFHUCAwEAAaMjMCEwDgYDVR0PAQH/BAQDAgKkMA8GA1UdEwEB
/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBABlZJ7oIZVvEFP5oDsvAfsUqsMm5
rBf9noBiTIy3nImQZa/Ic3Lnd5dgHC2d3xxnVzxpd+2mKiMwuHTJxwz0W6q8d6MO
YEoOCKY4/dqnLxJooDCbnuoOQHR0UE4cjOrtJFzv4fQdwuee3gKPZ2zIo+2EZFBB
o/gZ3eD9DYVSRBacPpEaQMJPtdjbS2HSas8HLm77Mg23GpRzpXODtX2axnN3yCvl
5WV2efMF+V/UIDmhu7ntYJNAO//PLfQvT+wlRjzRfQtd1w9x+2TxKBF0WjJltQXD
LZGiCc1WFR2deNMtf5BpQ/IK8H4fiYwhoIwGVns15b0N6fx4/arakNd03RE=
-----END CERTIFICATE-----

```

获取token
```shell script
kubectl get secrets -n devops gitlab-sa-token-57qq7 -o json|jq -r '.data.token'|base64 -d

eyJhbGciOiJSUzI1NiIsImtpZCI6InhlWUFhRzloVFQ4SWc5WTN6eVZEMnBhT0l1SGpsa0NiTkFmMGJSMkZXcmsifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZXZvcHMiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlY3JldC5uYW1lIjoiZ2l0bGFiLXNhLXRva2VuLTU3cXE3Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImdpdGxhYi1zYSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6ImE2Nzc4ZDk4LTAyZjMtNDlhZC05MmRmLTM3OTM5MDVlNWY5OSIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkZXZvcHM6Z2l0bGFiLXNhIn0.RNW-AOmrsbcOFWPb0mz7_Oh4vHD5Qnu0DaacEKKGOMv0wGyXj5jMx_oYXVAlJVp1ObrXgr22X53aj03XbHVtVuY4R0Y1FrqmZ_nDGQ610haj5J9UgAUBbjuXtnO239s-dKF-877iL2XaES2VZgPzQIj6pWKKYK7wU5oqjlhXlNUj1sKDZw_c5X1OVQcJhDhw-giXleaV7SjmZl2_eSVIvWXuZL14bstEA7dbNICJJntwqyKyIeZAeHO9MvoPZBmYmccoDXdf8wb8oV8jAnyPfXsm1KE7gJPX4MdgJGaxCGPNh7z4OoebpreJ1fE6VqpZVLC_fgGofBO8nmxGO-8FnA
```

### 生成k8s可以拉取镜像的secrets

```shell script
kubectl create secret docker-registry harbor --docker-server=reg.pdf.cn --docker-username=admin --docker-password=pdf0824 --docker-email=15114876417@163.com -n devops
```

### 注意

此时gitlab操作集群的context为 devops, 而不是 default, 所以创建资源的时候,一定要写namespace