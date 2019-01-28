export KUBECONFIG=/home/pliu/kubeconfig

rm  _output/linux/amd64/cluster-network-operator
bash hack/build-go.sh
env POD_NAME=LOCAL \
	$(cat env.sh) \
	OVN_IMAGE="docker.io/ovnkube/ovn-daemonset:latest" \
       	_output/linux/amd64/cluster-network-operator
