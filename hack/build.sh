export KUBECONFIG=/home/pliu/go/src/github.com/openshift/installer/test/auth/kubeconfig

rm  _output/linux/amd64/cluster-network-operator
bash hack/build-go.sh
env POD_NAME=LOCAL \
	$(cat images) \
	MULTUS_IMAGE="quay.io/openshift/origin-multus-cni:v4.0.0" \
	OVN_IMAGE="docker.io/openshift/origin-ovn-kubernetes:v4.0.0" \
	OVN_IMAGE="docker.io/ovnkube/ovn-daemonset:latest" \
       	_output/linux/amd64/cluster-network-operator
