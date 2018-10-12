// +build e2e

/*
Copyright 2018 Knative Authors LLC
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

package test

import (
	"fmt"
	"strings"
	"testing"

	duckv1alpha1 "github.com/knative/pkg/apis/duck/v1alpha1"
	knativetest "github.com/knative/pkg/test"
	"github.com/knative/pkg/test/logging"
	corev1 "k8s.io/api/core/v1"

	"github.com/knative/build-pipeline/pkg/apis/pipeline/v1alpha1"
)

// TestTaskRun is an integration test that will verify a very simple "hello world" TaskRun can be
// executed.
func TestTaskRun(t *testing.T) {
	logger := logging.GetContextLogger(t.Name())
	c, namespace := setup(t, logger)

	knativetest.CleanupOnInterrupt(func() { tearDown(logger, c.KubeClient, namespace) }, logger)
	defer tearDown(logger, c.KubeClient, namespace)

	logger.Infof("Creating volume %s to collect log output", hwVolumeName)
	if _, err := c.KubeClient.Kube.CoreV1().PersistentVolumeClaims(namespace).Create(getHelloWorldVolumeClaim(namespace)); err != nil {
		t.Fatalf("Failed to create Volume `%s`: %s", hwTaskName, err)
	}

	logger.Infof("Creating Task and TaskRun in namespace %s", namespace)
	if _, err := c.TaskClient.Create(getHelloWorldTaskWithVolume(namespace, []string{"/bin/sh", "-c", fmt.Sprintf("echo %s > %s/%s", taskOutput, logPath, logFile)})); err != nil {
		t.Fatalf("Failed to create Task `%s`: %s", hwTaskName, err)
	}
	if _, err := c.TaskRunClient.Create(getHelloWorldTaskRun(namespace)); err != nil {
		t.Fatalf("Failed to create TaskRun `%s`: %s", hwTaskRunName, err)
	}

	logger.Infof("Waiting for TaskRun %s in namespace %s to complete", hwTaskRunName, namespace)
	if err := WaitForTaskRunState(c, hwTaskRunName, func(tr *v1alpha1.TaskRun) (bool, error) {
		c := tr.Status.GetCondition(duckv1alpha1.ConditionSucceeded)
		if c.Status == corev1.ConditionTrue {
			return true, nil
		} else if c.Status == corev1.ConditionFalse {
			return true, fmt.Errorf("pipeline run %s failed!", hwPipelineRunName)
		}
		return false, nil
	}, "TaskRunSuccess"); err != nil {
		t.Errorf("Error waiting for TaskRun %s to finish: %s", hwTaskRunName, err)
	}

	logger.Infof("Verifying TaskRun %s output in volume %s", hwTaskRunName, hwVolumeName)
	output, err := getBuildOutputFromVolume(logger, c, namespace, taskOutput)
	if err != nil {
		t.Fatalf("Unable to get build output from volume %s: %s", hwVolumeName, err)
	}
	if !strings.Contains(output, taskOutput) {
		t.Fatalf("Expected output %s from pod %s but got %s", buildOutput, hwValidationPodName, output)
	}
}
