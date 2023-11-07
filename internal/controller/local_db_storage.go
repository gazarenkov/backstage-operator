/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package controller

import (
	bs "backstage.io/backstage-deploy-operator/api/v1alpha1"
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

/*
apiVersion: v1
kind: PersistentVolume
metadata:

	name: postgres-storage
	namespace: backstage
	labels:
	  type: local

spec:

	storageClassName: manual
	capacity:
	  storage: 2G
	accessModes:
	  - ReadWriteOnce
	persistentVolumeReclaimPolicy: Retain
	hostPath:
	  path: '/mnt/data'

---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: postgres-storage-claim
spec:
  storageClassName: manual
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 2G
*/

var (
	// taken from https://backstage.io/docs/deployment/k8s/
	// TODO should we move it to the file?

	DefaultLocalPV = corev1.PersistentVolume{
		ObjectMeta: v1.ObjectMeta{
			Name: "postgres-storage",
			//	Namespace:    to initialize,
			Labels: map[string]string{
				"type": "local",
			},
		},
		Spec: corev1.PersistentVolumeSpec{
			StorageClassName: "manual",
			Capacity: corev1.ResourceList{
				corev1.ResourceName(corev1.ResourceStorage): resource.MustParse("2Gi"),
			},

			AccessModes: []corev1.PersistentVolumeAccessMode{
				corev1.ReadWriteOnce,
			},
			PersistentVolumeReclaimPolicy: corev1.PersistentVolumeReclaimRetain,
			PersistentVolumeSource: corev1.PersistentVolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/demo",
				},
			},
		},
	}

	DefaultLocalPVC = corev1.PersistentVolumeClaim{

		ObjectMeta: v1.ObjectMeta{
			Name: fmt.Sprintf("%s-%s", DefaultLocalPV.Name, "claim"),
			//	Namespace:    to initialize,
			Labels: DefaultLocalPV.Labels,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			StorageClassName: &DefaultLocalPV.Spec.StorageClassName,
			AccessModes:      DefaultLocalPV.Spec.AccessModes,
			Resources: corev1.ResourceRequirements{
				Requests: DefaultLocalPV.Spec.Capacity,
			},
		},
	}
)

func (r *BackstageDeploymentReconciler) applyPV(ctx context.Context, backstage bs.Backstage, ns string) error {
	// Postgre PersistentVolume
	lg := log.FromContext(ctx)

	givenPV := backstage.Spec.LocalDb.PersistentVolume
	pv := corev1.PersistentVolume{}
	// find by PV.status.name
	// or do we need to get PV by label ?
	err := r.Get(ctx, types.NamespacedName{Name: backstage.Status.LocalDb.PersistentVolume.Name, Namespace: ns}, &pv)

	if err != nil {

		if errors.IsNotFound(err) {

			// create PV from default or from given (if any)
			pv = DefaultLocalPV
			if givenPV.Kind != "" {
				pv = givenPV
			}

			//addParameters to PV
			if backstage.Spec.LocalDb.Parameters.StorageCapacity != "" {
				pv.Spec.Capacity = corev1.ResourceList{
					corev1.ResourceName(corev1.ResourceStorage): resource.MustParse(
						backstage.Spec.LocalDb.Parameters.StorageCapacity),
				}
			}

			r.labels(pv.ObjectMeta, backstage.Name)
			err := r.Create(ctx, &pv)
			if err != nil {
				//status = fmt.Sprintf("failed to create postgre persistent volume, reason:%s", err)
				return fmt.Errorf("failed to create postgre persistent volume, reason:%s", err)
			}
			//backstage.Status.PostgreState = "failed to create postgre volume"
		} else {
			return fmt.Errorf("failed to get  persistent volume, reason:%s", err)
			//status = fmt.Sprintf("failed to create postgre persistent volume, reason:%s", err)
		}
	} else {
		// TODO update?
		lg.Info("PV update is not supported yet")
	}

	backstage.Status.LocalDb.PersistentVolume.Name = pv.Name
	err = r.Status().Update(ctx, &backstage)
	if err != nil {
		lg.Error(err, "failed to update Backstage Status")
		return fmt.Errorf("failed to update Backstage Status, reason:%s", err)
	}

	return nil
}

func (r *BackstageDeploymentReconciler) applyPVC(ctx context.Context, backstage bs.Backstage, ns string) error {
	// Postgre PersistentVolumeClaim
	lg := log.FromContext(ctx)

	givenPVC := backstage.Spec.LocalDb.PersistentVolumeClaim
	pvc := corev1.PersistentVolumeClaim{}
	pvcList := corev1.PersistentVolumeClaimList{}
	// find by PVC.status.name
	// or do we need to get PV by label ?
	err := r.List(ctx, &pvcList, client.InNamespace(ns), client.MatchingLabels{"app.kubernetes.io/instance": backstage.Name})
	if err != nil {
		return fmt.Errorf("failed to get  persistent volume claim, reason:%s", err)
	}

	if len(pvcList.Items) == 0 {

		// create PV from default or from given (if any)
		pvc = DefaultLocalPVC
		pvc.Namespace = ns
		if givenPVC.Kind != "" {
			pvc = givenPVC
		}

		//addParameters to PVC
		if backstage.Spec.LocalDb.Parameters.StorageCapacity != "" {
			pvc.Spec.Resources.Requests = corev1.ResourceList{
				corev1.ResourceName(corev1.ResourceStorage): resource.MustParse(
					backstage.Spec.LocalDb.Parameters.StorageCapacity),
			}
		}
		r.labels(pvc.ObjectMeta, backstage.Name)
		err := r.Create(ctx, &pvc)
		if err != nil {
			//status = fmt.Sprintf("failed to create postgre persistent volume, reason:%s", err)
			return fmt.Errorf("failed to create postgre persistent volume claim , reason:%s", err)
		}

	} else {
		lg.Info("PVC update is not supported yet")
	}

	return nil
}
