package builder

import (
	"github.com/chriskery/hadoop-cluster-operator/pkg/apis/kubecluster.org/v1alpha1"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type Builder interface {
	SetupWithManager(mgr manager.Manager, recorder record.EventRecorder)
	Build(cluster *v1alpha1.HadoopCluster, status *v1alpha1.HadoopClusterStatus) error
}

func ResourceBuilders(mgr manager.Manager, recorder record.EventRecorder) []Builder {
	var builders = []Builder{&HdfsBuilder{}}
	for _, builder := range builders {
		builder.SetupWithManager(mgr, recorder)
	}
	return builders
}