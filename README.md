# foo-cni

## add cni to k8s

build and copy the binary under `/opt/cni/bin/foo-cni`

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