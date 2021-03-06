// Copyright 2019 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package prometheus

import (
	"context"
	monv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/kiegroup/kogito-cloud-operator/pkg/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceMonitorInterface has functions that interact with Service Monitor instances in the Kubernetes cluster
type ServiceMonitorInterface interface {
	List(namespace string) (*monv1.ServiceMonitorList, error)
}

type serviceMonitor struct {
	client *client.Client
}

func newServiceMonitor(c *client.Client) ServiceMonitorInterface {
	client.MustEnsureClient(c)
	return &serviceMonitor{
		client: c,
	}
}

func (s *serviceMonitor) List(namespace string) (*monv1.ServiceMonitorList, error) {
	log.Debugf("List service monitor instances from namespace %s", namespace)

	pList, err := s.client.PrometheusCli.ServiceMonitors(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	return pList, nil
}
