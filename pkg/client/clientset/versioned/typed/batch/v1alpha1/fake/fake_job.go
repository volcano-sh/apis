/*
Copyright The Volcano Authors.

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

package fake

import (
	gentype "k8s.io/client-go/gentype"
	v1alpha1 "volcano.sh/apis/pkg/apis/batch/v1alpha1"
	batchv1alpha1 "volcano.sh/apis/pkg/client/applyconfiguration/batch/v1alpha1"
	typedbatchv1alpha1 "volcano.sh/apis/pkg/client/clientset/versioned/typed/batch/v1alpha1"
)

// fakeJobs implements JobInterface
type fakeJobs struct {
	*gentype.FakeClientWithListAndApply[*v1alpha1.Job, *v1alpha1.JobList, *batchv1alpha1.JobApplyConfiguration]
	Fake *FakeBatchV1alpha1
}

func newFakeJobs(fake *FakeBatchV1alpha1, namespace string) typedbatchv1alpha1.JobInterface {
	return &fakeJobs{
		gentype.NewFakeClientWithListAndApply[*v1alpha1.Job, *v1alpha1.JobList, *batchv1alpha1.JobApplyConfiguration](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("jobs"),
			v1alpha1.SchemeGroupVersion.WithKind("Job"),
			func() *v1alpha1.Job { return &v1alpha1.Job{} },
			func() *v1alpha1.JobList { return &v1alpha1.JobList{} },
			func(dst, src *v1alpha1.JobList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.JobList) []*v1alpha1.Job { return gentype.ToPointerSlice(list.Items) },
			func(list *v1alpha1.JobList, items []*v1alpha1.Job) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
