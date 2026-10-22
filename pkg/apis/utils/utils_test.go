/*
Copyright 2018 The Kubernetes Authors.

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

package utils

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestGetControllerWithControllerRef(t *testing.T) {
	isController := true
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			OwnerReferences: []metav1.OwnerReference{
				{
					UID:        types.UID("owner-uid"),
					Controller: &isController,
				},
			},
		},
	}

	if got := GetController(pod); got != types.UID("owner-uid") {
		t.Errorf("GetController() = %q, want %q", got, "owner-uid")
	}
}

func TestGetControllerWithoutControllerRef(t *testing.T) {
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			OwnerReferences: []metav1.OwnerReference{},
		},
	}

	if got := GetController(pod); got != types.UID("") {
		t.Errorf("GetController() = %q, want empty UID", got)
	}
}

func TestGetControllerNonControllerOwnerIgnored(t *testing.T) {
	isController := false
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			OwnerReferences: []metav1.OwnerReference{
				{
					UID:        types.UID("non-controller-owner"),
					Controller: &isController,
				},
			},
		},
	}

	if got := GetController(pod); got != types.UID("") {
		t.Errorf("GetController() = %q, want empty UID", got)
	}
}

func TestGetControllerAccessorError(t *testing.T) {
	if got := GetController("not an object"); got != types.UID("") {
		t.Errorf("GetController() = %q, want empty UID for a non accessor input", got)
	}
}
