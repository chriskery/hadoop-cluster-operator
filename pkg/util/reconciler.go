package util

import (
	"context"
	"github.com/chriskery/hadoop-operator/pkg/apis/kubecluster.org/v1alpha1"
	log "github.com/sirupsen/logrus"
	appv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
	"reflect"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

// OnDependentCreateFunc modify expectations when dependent (pod/service) creation observed.
func OnDependentCreateFunc() func(event.CreateEvent) bool {
	return func(e event.CreateEvent) bool {
		rtype := e.Object.GetLabels()[v1alpha1.ReplicaTypeLabel]
		if len(rtype) == 0 {
			return false
		}

		if controllerRef := metav1.GetControllerOf(e.Object); controllerRef != nil {
			switch e.Object.(type) {
			case *corev1.Pod, *corev1.Service, *corev1.ConfigMap, *appv1.StatefulSet:
				return true
			default:
				return false
			}
		}

		return true
	}
}

// OnDependentUpdateFunc modify expectations when dependent (pod/service) update observed.
func OnDependentUpdateFunc(client client.Client) func(updateEvent event.UpdateEvent) bool {
	return func(e event.UpdateEvent) bool {
		newObj := e.ObjectNew
		oldObj := e.ObjectOld
		if newObj.GetResourceVersion() == oldObj.GetResourceVersion() {
			// Periodic resync will send update events for all known pods.
			// Two different versions of the same pod will always have different RVs.
			return false
		}

		newControllerRef := metav1.GetControllerOf(newObj)
		oldControllerRef := metav1.GetControllerOf(oldObj)
		controllerRefChanged := !reflect.DeepEqual(newControllerRef, oldControllerRef)

		var kind string
		if newControllerRef != nil {
			kind = newControllerRef.Kind
		} else if oldControllerRef != nil {
			kind = oldControllerRef.Kind
		} else {
			return false
		}

		var logger *log.Entry
		switch obj := newObj.(type) {
		case *corev1.Pod, *corev1.Service, *corev1.ConfigMap, *appv1.StatefulSet, *v1alpha1.HadoopCluster:
			logger = LoggerForGenericKind(obj, obj.GetObjectKind().GroupVersionKind().Kind)
		default:
			return false
		}

		if controllerRefChanged && oldControllerRef != nil {
			// The ControllerRef was changed. Sync the old controller, if any.
			if application := resolveControllerRef(kind, oldObj.GetNamespace(), oldControllerRef, client); application != nil {
				logger.Infof("pod/service controller ref updated: %v, %v", newObj, oldObj)
				return true
			}
		}

		// If it has a controller ref, that's all that matters.
		if newControllerRef != nil {
			application := resolveControllerRef(kind, newObj.GetNamespace(), newControllerRef, client)
			if application == nil {
				return false
			}
			logger.Debugf("pod/service has a controller ref: %v, %v", newObj, oldObj)
			return true
		}
		return false
	}
}

// resolveControllerRef returns the application referenced by a ControllerRef,
// or nil if the ControllerRef could not be resolved to a matching application
// of the correct Kind.
func resolveControllerRef(controllerKind string, namespace string, controllerRef *metav1.OwnerReference, client client.Client) metav1.Object {
	// We can't look up by UID, so look up by Name and then verify UID.
	// Don't even try to look up by Name if it's the wrong Kind.
	if controllerRef.Kind != controllerKind {
		return nil
	}

	var object metav1.Object
	if controllerRef.Kind == v1alpha1.HadoopClusterKind {
		hadoopCLuster := &v1alpha1.HadoopCluster{}
		err := client.Get(context.Background(), types.NamespacedName{
			Namespace: namespace, Name: controllerRef.Name,
		}, hadoopCLuster)
		if err != nil {
			return nil
		}
		object = hadoopCLuster
	} else if controllerRef.Kind == v1alpha1.HadoopApplicationKind {
		hadoopApplication := &v1alpha1.HadoopApplication{}
		err := client.Get(context.Background(), types.NamespacedName{
			Namespace: namespace, Name: controllerRef.Name,
		}, hadoopApplication)
		if err != nil {
			return nil
		}
		object = hadoopApplication
	} else {
		return nil
	}

	if object.GetUID() != controllerRef.UID {
		// The controller we found with this Name is not the same one that the
		// ControllerRef points to.
		return nil
	}
	return object
}

// OnDependentDeleteFunc modify expectations when dependent (pod/service) deletion observed.
func OnDependentDeleteFunc() func(event.DeleteEvent) bool {
	return func(e event.DeleteEvent) bool {

		rtype := e.Object.GetLabels()[v1alpha1.ReplicaTypeLabel]
		if len(rtype) == 0 {
			return false
		}

		// logrus.Info("Update on deleting function ", xgbr.ControllerName(), " delete object ", e.Object.GetName())
		if controllerRef := metav1.GetControllerOf(e.Object); controllerRef != nil {
			switch e.Object.(type) {
			case *corev1.Pod, *corev1.Service, *corev1.ConfigMap, *appv1.StatefulSet:
				return true
			default:
				return false
			}
		}
		return true
	}
}
