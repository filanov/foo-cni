# foo-cni

## add cni as a multus plugin

build and copy the binary under `/opt/cni/bin/foo-cni`

add foo-cni as a multus plugin

```bash
cat /etc/cni/net.d/00-multus.conf  | jq .
{
  "cniVersion": "0.3.1",
  "name": "multus-cni-network",
  "type": "multus",
  "capabilities": {
    "portMappings": true
  },
  "cniConf": "/host/etc/cni/multus/net.d",
  "kubeconfig": "/etc/cni/net.d/multus.d/multus.kubeconfig",
  "delegates": [
    {
      "cniVersion": "0.3.1",
      "name": "cbr0",
      "plugins": [
        {
          "delegate": {
            "hairpinMode": true,
            "isDefaultGateway": true
          },
          "type": "flannel"
        },
        {
          "capabilities": {
            "portMappings": true
          },
          "type": "portmap"
        },
        {
          "type": "foo-cni"
        }
      ]
    }
  ]
}
```

create a network attachment definition for the cni

```bash
cat <<EOF | kubectl create -f -
apiVersion: k8s.cni.cncf.io/v1
kind: NetworkAttachmentDefinition
metadata:
  name: foo-cni-network
  namespace: default
spec:
  config: '{
    "type": "foo-cni"
  }'
EOF
```


create a dummy pod to try

```bash
cat <<EOF | kubectl create -f -
apiVersion: v1
kind: Pod
metadata:
  name: test-pod
  annotations:
    k8s.v1.cni.cncf.io/networks: foo-cni-network
spec:
  containers:
  - name: test-container
    image: busybox
    command: ["sleep", "3600"]
EOF
```