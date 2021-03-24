package main

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"

	capav1 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha3"
	capzv1 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
	capgv1 "sigs.k8s.io/cluster-api-provider-gcp/api/v1alpha3"
)

func main() {
	capaCluster := &capav1.AWSCluster{}
	capzCluster := &capzv1.AzureCluster{}
	capgCluster := &capgv1.GCPCluster{}

	for _, obj := range []runtime.Object{capaCluster, capgCluster, capzCluster} {
		fmt.Println("#v", obj)
	}
}
