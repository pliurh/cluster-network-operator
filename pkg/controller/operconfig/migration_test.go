package operconfig

import (
	"context"
	"testing"

	"github.com/openshift/cluster-network-operator/pkg/client/fake"
	"k8s.io/apimachinery/pkg/api/equality"
	uns "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"

)

func TestConvertEgressNetworkPolicyToEgressFirewall(t *testing.T) {
	testCases := []struct {
		name    string
		objects []crclient.Object
		expectObjs []uns.Unstructured
	}{
		{
			name: "Convert EgressNetworkPolicy CR to EgressFirewall CR",
			objects: []crclient.Object{
				&uns.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "network.openshift.io/v1",
						"kind":       "EgressNetworkPolicy",
						"metadata": map[string]interface{}{
							"name":      "default",
							"namespace": "n1",
						},
						"spec": map[string]interface{}{
							"egress": []interface{}{
								map[string]interface{}{
									"to": map[string]interface{}{
										"dnsName": "docs.openshift.com",
									},
									"type": "Allow",
								},
								map[string]interface{}{
									"to": map[string]interface{}{
										"cidrSelector": "8.8.8.8",
									},
									"type": "Deny",
								},
							},
						},
					},
				},
			},
			expectObjs: []uns.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "k8s.ovn.org/v1",
						"kind":       "EgressFirewall",
						"metadata": map[string]interface{}{
							"name":      "default",
							"namespace": "n1",
						},
						"spec": map[string]interface{}{
							"egress": []interface{}{
								map[string]interface{}{
									"to": map[string]interface{}{
										"dnsName": "docs.openshift.com",
									},
									"type": "Allow",
								},
								map[string]interface{}{
									"to": map[string]interface{}{
										"cidrSelector": "8.8.8.8",
									},
									"type": "Deny",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			objs := []*uns.Unstructured{}
			client := fake.NewFakeClient(tc.objects...)

			err := convertEgressNetworkPolicyToEgressFirewall(context.Background(), client, objs)
			if err != nil {
				t.Fatalf("convertEgressNetworkPolicyToEgressFirewall failed: %v", err)
			}
			for _, expectObj := range tc.expectObjs {
				for _, outputObj := range objs {
					if expectObj.GetNamespace() == outputObj.GetNamespace() {
						if !equality.Semantic.DeepEqual(outputObj, expectObj) {
							t.Fatalf("the expected egress firewall CR: %v, got: %v", expectObj, outputObj)
						}
					}
				}
			}
		})
	}
}

func TestConvertEgressFirewallToEgressNetworkPolicy(t *testing.T) {
	testCases := []struct {
		name    string
		objects []crclient.Object
		expectObjs []uns.Unstructured
	}{
		{
			name: "Convert EgressNetworkPolicy CR to EgressFirewall CR",
			objects: []crclient.Object{
				&uns.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "k8s.ovn.org/v1",
						"kind":       "EgressFirewall",
						"metadata": map[string]interface{}{
							"name":      "default",
							"namespace": "n1",
						},
						"spec": map[string]interface{}{
							"egress": []interface{}{
								map[string]interface{}{
									"to": map[string]interface{}{
										"dnsName": "docs.openshift.com",
									},
									"type": "Allow",
								},
								map[string]interface{}{
									"to": map[string]interface{}{
										"cidrSelector": "8.8.8.8",
									},
									"type": "Deny",
								},
							},
						},
					},
				},
			},
			expectObjs: []uns.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "network.openshift.io/v1",
						"kind":       "EgressNetworkPolicy",
						"metadata": map[string]interface{}{
							"name":      "default",
							"namespace": "n1",
						},
						"spec": map[string]interface{}{
							"egress": []interface{}{
								map[string]interface{}{
									"to": map[string]interface{}{
										"dnsName": "docs.openshift.com",
									},
									"type": "Allow",
								},
								map[string]interface{}{
									"to": map[string]interface{}{
										"cidrSelector": "8.8.8.8",
									},
									"type": "Deny",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			objs := []*uns.Unstructured{}
			client := fake.NewFakeClient(tc.objects...)

			err := convertEgressFirewallToEgressNetworkPolicy(context.Background(), client, objs)
			if err != nil {
				t.Fatalf("convertEgressFirewallToEgressNetworkPolicy failed: %v", err)
			}
			for _, expectObj := range tc.expectObjs {
				for _, outputObj := range objs {
					if expectObj.GetNamespace() == outputObj.GetNamespace() {
						if !equality.Semantic.DeepEqual(outputObj, expectObj) {
							t.Fatalf("the expected egress network policy CR: %v, got: %v", expectObj, outputObj)
						}
					}
				}
			}
		})
	}
}

func TestConvertEgressFirewallToEgressNetworkPolicy(t *testing.T) {
	testCases := []struct {
		name    string
		objects []crclient.Object
		expectObjs []uns.Unstructured
	}{
		{
			name: "Convert EgressNetworkPolicy CR to EgressFirewall CR",
			objects: []crclient.Object{
				&uns.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "k8s.ovn.org/v1",
						"kind":       "EgressFirewall",
						"metadata": map[string]interface{}{
							"name":      "default",
							"namespace": "n1",
						},
						"spec": map[string]interface{}{
							"egress": []interface{}{
								map[string]interface{}{
									"to": map[string]interface{}{
										"dnsName": "docs.openshift.com",
									},
									"type": "Allow",
								},
								map[string]interface{}{
									"to": map[string]interface{}{
										"cidrSelector": "8.8.8.8",
									},
									"type": "Deny",
								},
							},
						},
					},
				},
			},
			expectObjs: []uns.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "network.openshift.io/v1",
						"kind":       "EgressNetworkPolicy",
						"metadata": map[string]interface{}{
							"name":      "default",
							"namespace": "n1",
						},
						"spec": map[string]interface{}{
							"egress": []interface{}{
								map[string]interface{}{
									"to": map[string]interface{}{
										"dnsName": "docs.openshift.com",
									},
									"type": "Allow",
								},
								map[string]interface{}{
									"to": map[string]interface{}{
										"cidrSelector": "8.8.8.8",
									},
									"type": "Deny",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			objs := []*uns.Unstructured{}
			client := fake.NewFakeClient(tc.objects...)

			err := convertEgressFirewallToEgressNetworkPolicy(context.Background(), client, objs)
			if err != nil {
				t.Fatalf("convertEgressFirewallToEgressNetworkPolicy failed: %v", err)
			}
			for _, expectObj := range tc.expectObjs {
				for _, outputObj := range objs {
					if expectObj.GetNamespace() == outputObj.GetNamespace() {
						if !equality.Semantic.DeepEqual(outputObj, expectObj) {
							t.Fatalf("the expected egress network policy CR: %v, got: %v", expectObj, outputObj)
						}
					}
				}
			}
		})
	}
}
