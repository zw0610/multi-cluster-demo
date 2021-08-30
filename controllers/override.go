/*
Copyright 2021.

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

package controllers

import (
	"context"
	"fmt"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *KubeflowJobReconciler) Get(ctx context.Context, key client.ObjectKey, obj client.Object) error {
	cluster := ctx.Value(ClusterKey)
	if cluster == nil {
		return fmt.Errorf("key %s not found in context", ClusterKey)
	}

	c, ok := r.clientMapper[cluster.(string)]
	if !ok {
		return fmt.Errorf("cluster %s not found in client mapper", cluster)
	}

	return c.Get(ctx, key, obj)
}

func (r *KubeflowJobReconciler) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	cluster := ctx.Value(ClusterKey)
	if cluster == nil {
		return fmt.Errorf("key %s not found in context", ClusterKey)
	}

	c, ok := r.clientMapper[cluster.(string)]
	if !ok {
		return fmt.Errorf("cluster %s not found in client mapper", cluster)
	}

	return c.Create(ctx, obj, opts...)
}
