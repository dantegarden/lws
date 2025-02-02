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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "sigs.k8s.io/lws/api/leaderworkerset/v1"
)

// RolloutStrategyApplyConfiguration represents a declarative configuration of the RolloutStrategy type for use
// with apply.
type RolloutStrategyApplyConfiguration struct {
	Type                       *v1.RolloutStrategyType                       `json:"type,omitempty"`
	RollingUpdateConfiguration *RollingUpdateConfigurationApplyConfiguration `json:"rollingUpdateConfiguration,omitempty"`
}

// RolloutStrategyApplyConfiguration constructs a declarative configuration of the RolloutStrategy type for use with
// apply.
func RolloutStrategy() *RolloutStrategyApplyConfiguration {
	return &RolloutStrategyApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *RolloutStrategyApplyConfiguration) WithType(value v1.RolloutStrategyType) *RolloutStrategyApplyConfiguration {
	b.Type = &value
	return b
}

// WithRollingUpdateConfiguration sets the RollingUpdateConfiguration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RollingUpdateConfiguration field is set to the value of the last call.
func (b *RolloutStrategyApplyConfiguration) WithRollingUpdateConfiguration(value *RollingUpdateConfigurationApplyConfiguration) *RolloutStrategyApplyConfiguration {
	b.RollingUpdateConfiguration = value
	return b
}
