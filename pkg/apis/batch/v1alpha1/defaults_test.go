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

package v1alpha1

import (
	"testing"
)

func int32Ptr(val int32) *int32 {
	return &val
}

func TestSetDefaults_TaskSpec(t *testing.T) {
	testCases := []struct {
		name     string
		input    *TaskSpec
		expected *int32
	}{
		{
			name:     "nil task pointer handles gracefully",
			input:    nil,
			expected: nil,
		},
		{
			name: "task with explicit MinAvailable preserved",
			input: &TaskSpec{
				Name:         "task-explicit",
				Replicas:     5,
				MinAvailable: int32Ptr(2),
			},
			expected: int32Ptr(2),
		},
		{
			name: "task without partition policy defaults to Replicas",
			input: &TaskSpec{
				Name:     "task-default-replicas",
				Replicas: 4,
			},
			expected: int32Ptr(4),
		},
		{
			name: "task with PartitionPolicy and MinPartitions > 0 defaults to MinPartitions * PartitionSize",
			input: &TaskSpec{
				Name:     "task-partition-policy",
				Replicas: 10,
				PartitionPolicy: &PartitionPolicySpec{
					MinPartitions: 2,
					PartitionSize: 3,
				},
			},
			expected: int32Ptr(6),
		},
		{
			name: "task with PartitionPolicy but MinPartitions == 0 falls back to Replicas",
			input: &TaskSpec{
				Name:     "task-partition-policy-zero-min",
				Replicas: 3,
				PartitionPolicy: &PartitionPolicySpec{
					MinPartitions: 0,
					PartitionSize: 5,
				},
			},
			expected: int32Ptr(3),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			SetDefaults_TaskSpec(tc.input)
			if tc.input == nil {
				return
			}
			if tc.expected == nil {
				if tc.input.MinAvailable != nil {
					t.Fatalf("expected nil MinAvailable, got %v", *tc.input.MinAvailable)
				}
				return
			}
			if tc.input.MinAvailable == nil {
				t.Fatalf("expected MinAvailable %d, got nil", *tc.expected)
			}
			if *tc.input.MinAvailable != *tc.expected {
				t.Fatalf("expected MinAvailable %d, got %d", *tc.expected, *tc.input.MinAvailable)
			}
		})
	}
}

func TestSetDefaults_Job(t *testing.T) {
	testCases := []struct {
		name                    string
		job                     *Job
		expectedJobMinAvailable int32
		expectedTaskMins        []int32
	}{
		{
			name:                    "nil job handles gracefully",
			job:                     nil,
			expectedJobMinAvailable: 0,
			expectedTaskMins:        nil,
		},
		{
			name: "job with MinAvailable=0 calculates sum of task defaults",
			job: &Job{
				Spec: JobSpec{
					Tasks: []TaskSpec{
						{Name: "task-1", Replicas: 2},
						{Name: "task-2", Replicas: 3},
					},
				},
			},
			expectedJobMinAvailable: 5,
			expectedTaskMins:        []int32{2, 3},
		},
		{
			name: "job with explicit MinAvailable preserves job MinAvailable while defaulting tasks",
			job: &Job{
				Spec: JobSpec{
					MinAvailable: 1,
					Tasks: []TaskSpec{
						{Name: "task-1", Replicas: 4},
					},
				},
			},
			expectedJobMinAvailable: 1,
			expectedTaskMins:        []int32{4},
		},
		{
			name: "job with mixed task configurations and partition policies",
			job: &Job{
				Spec: JobSpec{
					Tasks: []TaskSpec{
						{
							Name:     "worker",
							Replicas: 8,
							PartitionPolicy: &PartitionPolicySpec{
								MinPartitions: 2,
								PartitionSize: 2,
							},
						},
						{
							Name:         "driver",
							Replicas:     1,
							MinAvailable: int32Ptr(1),
						},
					},
				},
			},
			expectedJobMinAvailable: 5,
			expectedTaskMins:        []int32{4, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			SetDefaults_Job(tc.job)
			if tc.job == nil {
				return
			}
			if tc.job.Spec.MinAvailable != tc.expectedJobMinAvailable {
				t.Fatalf("expected Job Spec.MinAvailable %d, got %d", tc.expectedJobMinAvailable, tc.job.Spec.MinAvailable)
			}
			for i, exp := range tc.expectedTaskMins {
				task := tc.job.Spec.Tasks[i]
				if task.MinAvailable == nil {
					t.Fatalf("task %d: expected MinAvailable %d, got nil", i, exp)
				}
				if *task.MinAvailable != exp {
					t.Fatalf("task %d: expected MinAvailable %d, got %d", i, exp, *task.MinAvailable)
				}
			}
		})
	}
}
