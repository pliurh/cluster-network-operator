rm  _output/linux/amd64/cluster-network-operator
bash hack/build-go.sh
env POD_NAME=LOCAL \
	$(cat env.sh) \
	OVN_IMAGE="quay.io/pliurh/ovn-daemonset:3e7481d3295279289922c0c672c0348e2fc19fb0" \
       	_output/linux/amd64/cluster-network-operator
