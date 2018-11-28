package network

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	uns "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	legacyconfigv1 "github.com/openshift/api/legacyconfig/v1"
	cpv1 "github.com/openshift/api/openshiftcontrolplane/v1"
	netv1 "github.com/openshift/cluster-network-operator/pkg/apis/networkoperator/v1"
	"github.com/openshift/cluster-network-operator/pkg/render"
)

// NodeNameMagicString is substituted at runtime for the
// real nodename
const NodeNameMagicString = "%%NODENAME%%"

// renderOVNKubernetes returns the manifests for the ovn-kubernetes.
// This creates
// - the ClusterNetwork object
// - the ovn-kubernetes namespace
// - the ovn-kubeernetes setup
// - the ovnkube-node daemonset
// - the ovnkube-master deployment
// and some other small things.
func renderOVNKubernetes(conf *netv1.NetworkConfigSpec, manifestDir string) ([]*uns.Unstructured, error) {
	c := conf.DefaultNetwork.OVNKubernetesConfig

	objs := []*uns.Unstructured{}

	// render the manifests on disk
	data := render.MakeRenderData()
	data.Data["OvnImage"] = os.Getenv("OVN_IMAGE")
	data.Data["OvnImagePolicy"] = "IfNotPresent"
	data.Data["HypershiftImage"] = os.Getenv("HYPERSHIFT_IMAGE")
	data.Data["GenevePort"] = c.GenevePort
	data.Data["MTU"] = c.MTU

	operCfg, err := controllerOVNConfig(conf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build controller config")
	}
	data.Data["NetworkControllerConfig"] = operCfg

	nodeCfg, err := nodeOVNConfig(conf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build node config")
	}
	data.Data["NodeConfig"] = nodeCfg

	manifests, err := render.RenderDir(filepath.Join(manifestDir, "network/ovn-kubernetes"), &data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to render manifests")
	}

	objs = append(objs, manifests...)
	return objs, nil
}

// validateOVNKubernetes checks that the openshift-sdn specific configuration
// is basically sane.
func validateOVNKubernetes(conf *netv1.NetworkConfigSpec) []error {
	out := []error{}
	oc := conf.DefaultNetwork.OVNKubernetesConfig
	if oc == nil {
		out = append(out, errors.Errorf("OVNKubernetesConfig cannot be nil"))
		return out
	}

	if len(conf.ClusterNetworks) != 1 {
		out = append(out, errors.Errorf("ClusterNetworks cannot be empty"))
	}

	if networkPluginName() == "" {
		out = append(out, errors.Errorf("invalid network plugin name"))
	}

	if oc.GenevePort != nil && (*oc.GenevePort < 1 || *oc.GenevePort > 65535) {
		out = append(out, errors.Errorf("invalid GenevePort %d", *oc.GenevePort))
	}

	if oc.MTU != nil && (*oc.MTU < 576 || *oc.MTU > 65536) {
		out = append(out, errors.Errorf("invalid MTU %d", *oc.MTU))
	}

	return out
}

// isOVNKubernetesChangeSafe currently returns an error if any changes are made.
// In the future, we may support rolling out MTU or other alterations.
func isOVNKubernetesChangeSafe(prev, next *netv1.NetworkConfigSpec) []error {
	pn := prev.DefaultNetwork.OVNKubernetesConfig
	nn := next.DefaultNetwork.OVNKubernetesConfig

	if reflect.DeepEqual(pn, nn) {
		return []error{}
	}
	return []error{errors.Errorf("cannot change ovn-kubernetes configuration")}
}

func fillOVNKubernetesDefaults(conf *netv1.NetworkConfigSpec) {
	if conf.DeployKubeProxy == nil {
		prox := false
		conf.DeployKubeProxy = &prox
	}

	if conf.KubeProxyConfig == nil {
		conf.KubeProxyConfig = &netv1.ProxyConfig{}
	}
	if conf.KubeProxyConfig.BindAddress == "" {
		conf.KubeProxyConfig.BindAddress = "0.0.0.0"
	}

	sc := conf.DefaultNetwork.OVNKubernetesConfig
	if sc.GenevePort == nil {
		var port uint32 = 6081
		sc.GenevePort = &port
	}
	if sc.MTU == nil {
		var mtu uint32 = 1400
		sc.MTU = &mtu
	}
}

func networkPluginName() string {
	return "ovn-kubernetes"
}

//PHIL -- what is this?
// controllerOVNConfig builds the contents of controller-config.yaml
// for the controller
func controllerOVNConfig(conf *netv1.NetworkConfigSpec) (string, error) {
	c := conf.DefaultNetwork.OVNKubernetesConfig

	// generate master network configuration
	//PHIL need netv1.NetworkConfigSpec version
	ippools := []cpv1.ClusterNetworkEntry{}
	for _, net := range conf.ClusterNetworks {
		ippools = append(ippools, cpv1.ClusterNetworkEntry{CIDR: net.CIDR, HostSubnetLength: net.HostSubnetLength})
	}

	//PHIL we need a ovn version here
	cfg := cpv1.OpenShiftControllerManagerConfig{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "openshiftcontrolplane.config.openshift.io/v1",
			Kind:       "OpenShiftControllerManagerConfig",
		},
		// no ObjectMeta - not an API object

		// This is the OpenShiftSDN configuration.
		// It will be converted into the ovn configuration.
		// ClusterNetworks is extended cidr (ip/mask/subnet)
		// VXLANPort is genevePort
		//PHIL netv1.NetworkConfigSpec.OVNKubernetesConfig
		Network: cpv1.NetworkControllerConfig{
			NetworkPluginName:  networkPluginName(),
			ClusterNetworks:    ippools,
			ServiceNetworkCIDR: conf.ServiceNetwork,
			VXLANPort:          *c.GenevePort,
		},
	}

	buf, err := yaml.Marshal(cfg)
	if err != nil {
		return "", err
	}

	// HACK: danw changed the capitalization of VXLANPort, but it's not yet
	// merged in to origin. So just set both.
	// Remove when origin merges api.
	obj := &uns.Unstructured{}
	err = yaml.Unmarshal(buf, obj)
	if err != nil {
		return "", err
	}
	p := json.Number(fmt.Sprintf("%d", *c.GenevePort))

	uns.SetNestedField(obj.Object, p, "network", "vxLANPort")

	buf, err = yaml.Marshal(obj)
	return string(buf), err
}

// nodeOVNConfig builds the (yaml text of) the NodeConfig object
// consumed by the sdn node process
func nodeOVNConfig(conf *netv1.NetworkConfigSpec) (string, error) {
	c := conf.DefaultNetwork.OVNKubernetesConfig

	result := legacyconfigv1.NodeConfig{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "NodeConfig",
		},
		NodeName: NodeNameMagicString,
		NetworkConfig: legacyconfigv1.NodeNetworkConfig{
			NetworkPluginName: networkPluginName(),
			MTU:               *c.MTU,
		},
		// ServingInfo is used by both the proxy and metrics components
		ServingInfo: legacyconfigv1.ServingInfo{
			ClientCA:    "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt",
			BindAddress: conf.KubeProxyConfig.BindAddress + ":10251", // port is unused but required
		},

		// Openshift-sdn calls the CRI endpoint directly; point it to crio
		KubeletArguments: legacyconfigv1.ExtendedArguments{
			"container-runtime":          {"remote"},
			"container-runtime-endpoint": {"/var/run/crio/crio.sock"},
		},

		IPTablesSyncPeriod: conf.KubeProxyConfig.IptablesSyncPeriod,
		ProxyArguments:     conf.KubeProxyConfig.ProxyArguments,
	}

	buf, err := yaml.Marshal(result)
	return string(buf), err
}
