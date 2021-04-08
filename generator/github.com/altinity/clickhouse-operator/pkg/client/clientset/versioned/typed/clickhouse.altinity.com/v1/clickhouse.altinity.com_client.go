/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
	"github.com/altinity/clickhouse-operator/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type ClickhouseV1Interface interface {
	RESTClient() rest.Interface
	ClickHouseInstallationsGetter
	ClickHouseInstallationTemplatesGetter
	ClickHouseOperatorConfigurationsGetter
}

// ClickhouseV1Client is used to interact with features provided by the clickhouse.altinity.com group.
type ClickhouseV1Client struct {
	restClient rest.Interface
}

func (c *ClickhouseV1Client) ClickHouseInstallations(namespace string) ClickHouseInstallationInterface {
	return newClickHouseInstallations(c, namespace)
}

func (c *ClickhouseV1Client) ClickHouseInstallationTemplates(namespace string) ClickHouseInstallationTemplateInterface {
	return newClickHouseInstallationTemplates(c, namespace)
}

func (c *ClickhouseV1Client) ClickHouseOperatorConfigurations(namespace string) ClickHouseOperatorConfigurationInterface {
	return newClickHouseOperatorConfigurations(c, namespace)
}

// NewForConfig creates a new ClickhouseV1Client for the given config.
func NewForConfig(c *rest.Config) (*ClickhouseV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ClickhouseV1Client{client}, nil
}

// NewForConfigOrDie creates a new ClickhouseV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ClickhouseV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ClickhouseV1Client for the given RESTClient.
func New(c rest.Interface) *ClickhouseV1Client {
	return &ClickhouseV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ClickhouseV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}