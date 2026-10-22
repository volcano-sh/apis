/*
Copyright 2026 The Volcano Authors.

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

package helpers

import (
	"context"
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"

	vcbatch "volcano.sh/apis/pkg/apis/batch/v1alpha1"
)

func TestCreateOrUpdateSecretUpdatesOnNonConfigKeyChange(t *testing.T) {
	job := &vcbatch.Job{
		ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "job1"},
	}

	client := fake.NewSimpleClientset(&v1.Secret{
		ObjectMeta: metav1.ObjectMeta{Namespace: job.Namespace, Name: "secret1"},
		Data: map[string][]byte{
			"config":      []byte("same"),
			"ssh.private": []byte("old-key"),
		},
	})

	newData := map[string][]byte{
		"config":      []byte("same"),
		"ssh.private": []byte("new-key"),
	}

	if err := CreateOrUpdateSecret(job, client, newData, "secret1"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got, err := client.CoreV1().Secrets(job.Namespace).Get(context.TODO(), "secret1", metav1.GetOptions{})
	if err != nil {
		t.Fatalf("unexpected error getting secret: %v", err)
	}

	if string(got.Data["ssh.private"]) != "new-key" {
		t.Errorf("expected secret data to be updated to %q, got %q", "new-key", got.Data["ssh.private"])
	}
}
