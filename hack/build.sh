rm  _output/linux/amd64/cluster-network-operator
bash hack/build-go.sh
env POD_NAME=LOCAL \
	$(cat env.sh) \
	OVN_IMAGE="quay.io/pliurh/ovn-daemonset:921fb1fc7b834a7b9e00d9b589f750cbead0b375" \
       	_output/linux/amd64/cluster-network-operator
