package operconfig

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	uns "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"

	operv1 "github.com/openshift/api/operator/v1"
	cnoclient "github.com/openshift/cluster-network-operator/pkg/client"
)

var gvrEgressFirewall = schema.GroupVersionResource{Group: "k8s.ovn.org", Version: "v1", Resource: "egressfirewalls"}
var gvrEgressNetworkPolicy = schema.GroupVersionResource{Group: "network.openshift.io", Version: "v1", Resource: "egressnetworkpolicies"}

func migrateEgressFirewallCRs(ctx context.Context, operConfig *operv1.Network, client cnoclient.Client) error {
	if string(operConfig.Spec.DefaultNetwork.Type) == operConfig.Spec.Migration.NetworkType {
		return nil
	}
	switch operConfig.Spec.DefaultNetwork.Type {
	case operv1.NetworkTypeOpenShiftSDN:
		return convertEgressNetworkPolicyToEgressFirewall(ctx, client)
	case operv1.NetworkTypeOVNKubernetes:
		return convertEgressFirewallToEgressNetworkPolicy(ctx, client)
	}

	return nil
}

func convertEgressNetworkPolicyToEgressFirewall(ctx context.Context, client cnoclient.Client) error {
	enpList, err := client.Default().Dynamic().Resource(gvrEgressNetworkPolicy).List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, enp := range enpList.Items {
		log.Printf("Convert EgressNetworkPolicy %s/%s", enp.GetNamespace(), enp.GetName())

		egressFirewall := &uns.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "k8s.ovn.org/v1",
				"kind":       "EgressFirewall",
				"metadata": map[string]interface{}{
					"name": enp.GetName(),
				},
				"spec": enp.Object["spec"],
			},
		}
		_, err = client.Default().Dynamic().Resource(gvrEgressFirewall).Namespace(enp.GetNamespace()).Create(ctx, egressFirewall, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func convertEgressFirewallToEgressNetworkPolicy(ctx context.Context, client cnoclient.Client) error {
	efList, err := client.Default().Dynamic().Resource(gvrEgressFirewall).List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, ef := range efList.Items {
		log.Printf("Convert EgressNetworkPolicy %s/%s", ef.GetNamespace(), ef.GetName())

		egressNetworkPolicy := &uns.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "network.openshift.io/v1",
				"kind":       "EgressNetworkPolicy",
				"metadata": map[string]interface{}{
					"name": ef.GetName(),
				},
				"spec": ef.Object["spec"],
			},
		}
		_, err = client.Default().Dynamic().Resource(gvrEgressNetworkPolicy).Namespace(ef.GetNamespace()).Create(ctx, egressNetworkPolicy, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}
