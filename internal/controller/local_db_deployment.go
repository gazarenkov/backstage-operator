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
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

var (
	DefaultLocalDbDeployment = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: postgres
spec:
  replicas: 1
  selector:
    matchLabels:
      app: postgres
  template:
    metadata:
      labels:
        app: postgres
    spec:
      containers:
        - name: postgres
          image: postgres:13.2-alpine
          imagePullPolicy: 'IfNotPresent'
          ports:
            - containerPort: 5432
          envFrom:
            - secretRef:
                name: postgres-secrets
          volumeMounts:
            - mountPath: /var/lib/postgresql/data
              name: postgresdb
      volumes:
        - name: postgresdb
          persistentVolumeClaim:
            claimName: postgres-storage-claim
`
	DefaultLocalDbService = `apiVersion: v1
kind: Service
metadata:
  name: postgres
spec:
  selector:
    app: postgres
  ports:
    - port: 5432
`
)

func (r *BackstageDeploymentReconciler) applyLocalDbDeployment(ctx context.Context, backstage bs.Backstage, ns string) error {

	//lg := log.FromContext(ctx)

	var deployment *appsv1.Deployment
	if backstage.Spec.LocalDb.Deployment.Kind != "" {
		deployment = &backstage.Spec.LocalDb.Deployment
	} else {
		deployment = &appsv1.Deployment{}
		err := readYaml(DefaultLocalDbDeployment, deployment)
		if err != nil {
			return err
		}
	}

	deployment.Namespace = ns

	// TODO consider apply merge instead?
	deployment.Namespace = ns
	err := r.Get(ctx, types.NamespacedName{Name: deployment.Name, Namespace: ns}, deployment)
	if err != nil {
		if errors.IsNotFound(err) {

		} else {
			return fmt.Errorf("failed to get deployment, reason: %s", err)
		}
	} else {
		return nil
	}

	err = r.Create(ctx, deployment)
	if err != nil {
		return fmt.Errorf("failed to create deplyment, reason: %s", err)
	}

	return nil
}

func (r *BackstageDeploymentReconciler) applyLocalDbService(ctx context.Context, backstage bs.Backstage, ns string) error {

	//lg := log.FromContext(ctx)

	var service *corev1.Service
	if backstage.Spec.LocalDb.Service.Kind != "" {
		service = &backstage.Spec.LocalDb.Service
	} else {
		service = &corev1.Service{}
		err := readYaml(DefaultLocalDbService, service)
		if err != nil {
			return err
		}
	}

	// TODO consider apply merge instead?
	service.Namespace = ns
	err := r.Get(ctx, types.NamespacedName{Name: service.Name, Namespace: ns}, service)
	if err != nil {
		if errors.IsNotFound(err) {
		} else {
			return fmt.Errorf("failed to get service, reason: %s", err)
		}
	} else {
		return nil
	}

	err = r.Create(ctx, service)
	//po := client.PatchOptions{}
	//patch := client.Apply
	//patch.Type()
	//patch.Data(fff)
	//err = r.Patch(ctx, service, patch)
	if err != nil {
		return fmt.Errorf("failed to create service, reason: %s", err)
	}

	return nil
}
