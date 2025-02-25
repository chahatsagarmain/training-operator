// Copyright 2024 The Kubeflow Authors
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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v2 "k8s.io/api/autoscaling/v2"
)

// TorchElasticPolicyApplyConfiguration represents a declarative configuration of the TorchElasticPolicy type for use
// with apply.
type TorchElasticPolicyApplyConfiguration struct {
	MaxRestarts *int32          `json:"maxRestarts,omitempty"`
	MinNodes    *int32          `json:"minNodes,omitempty"`
	MaxNodes    *int32          `json:"maxNodes,omitempty"`
	Metrics     []v2.MetricSpec `json:"metrics,omitempty"`
}

// TorchElasticPolicyApplyConfiguration constructs a declarative configuration of the TorchElasticPolicy type for use with
// apply.
func TorchElasticPolicy() *TorchElasticPolicyApplyConfiguration {
	return &TorchElasticPolicyApplyConfiguration{}
}

// WithMaxRestarts sets the MaxRestarts field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxRestarts field is set to the value of the last call.
func (b *TorchElasticPolicyApplyConfiguration) WithMaxRestarts(value int32) *TorchElasticPolicyApplyConfiguration {
	b.MaxRestarts = &value
	return b
}

// WithMinNodes sets the MinNodes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinNodes field is set to the value of the last call.
func (b *TorchElasticPolicyApplyConfiguration) WithMinNodes(value int32) *TorchElasticPolicyApplyConfiguration {
	b.MinNodes = &value
	return b
}

// WithMaxNodes sets the MaxNodes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxNodes field is set to the value of the last call.
func (b *TorchElasticPolicyApplyConfiguration) WithMaxNodes(value int32) *TorchElasticPolicyApplyConfiguration {
	b.MaxNodes = &value
	return b
}

// WithMetrics adds the given value to the Metrics field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Metrics field.
func (b *TorchElasticPolicyApplyConfiguration) WithMetrics(values ...v2.MetricSpec) *TorchElasticPolicyApplyConfiguration {
	for i := range values {
		b.Metrics = append(b.Metrics, values[i])
	}
	return b
}
