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

// SetDefaults_TaskSpec sets default values for TaskSpec.
// If MinAvailable is nil:
// - If PartitionPolicy is specified with MinPartitions > 0, it defaults to MinPartitions * PartitionSize.
// - Otherwise, it defaults to task.Replicas.
func SetDefaults_TaskSpec(task *TaskSpec) {
	if task == nil {
		return
	}
	if task.MinAvailable == nil {
		if task.PartitionPolicy != nil && task.PartitionPolicy.MinPartitions > 0 {
			minAvailable := task.PartitionPolicy.MinPartitions * task.PartitionPolicy.PartitionSize
			task.MinAvailable = &minAvailable
		} else {
			minAvailable := task.Replicas
			task.MinAvailable = &minAvailable
		}
	}
}

// SetDefaults_JobSpec sets default values for JobSpec tasks and calculates MinAvailable.
// If JobSpec.MinAvailable is 0, it calculates the default as the sum of all tasks' MinAvailable.
func SetDefaults_JobSpec(spec *JobSpec) {
	if spec == nil {
		return
	}
	var totalMinAvailable int32
	for i := range spec.Tasks {
		SetDefaults_TaskSpec(&spec.Tasks[i])
		if spec.Tasks[i].MinAvailable != nil {
			totalMinAvailable += *spec.Tasks[i].MinAvailable
		}
	}
	if spec.MinAvailable == 0 {
		spec.MinAvailable = totalMinAvailable
	}
}

// SetDefaults_Job applies API-level defaults to a Volcano Job.
func SetDefaults_Job(job *Job) {
	if job == nil {
		return
	}
	SetDefaults_JobSpec(&job.Spec)
}
