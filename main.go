package main

import (
	"fmt"
	"log"

	"k8s.io/apimachinery/pkg/runtime"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	ctlrun "sigs.k8s.io/controller-runtime"
	ctlrunmgr "sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/yaml"

	capav1 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha3"
	capzv1 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
	capgv1 "sigs.k8s.io/cluster-api-provider-gcp/api/v1alpha3"
	capiv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
)

func main() {
	capiCluster := &capiv1.Cluster{}
	capaCluster := &capav1.AWSCluster{}
	capzCluster := &capzv1.AzureCluster{}
	capgCluster := &capgv1.GCPCluster{}

	for _, obj := range []runtime.Object{capiCluster, capaCluster, capgCluster, capzCluster} {
		bs, err := yaml.Marshal(obj)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(bs))
	}

	cfg, err := ctlrun.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	mgr, err := ctlrun.NewManager(cfg, ctlrunmgr.Options{})
	if err != nil {
		log.Fatal(err)
	}
	_ = ctlrun.NewControllerManagedBy(mgr)

	errlist := []error{}
	_ = kerrors.NewAggregate(errlist)
}
