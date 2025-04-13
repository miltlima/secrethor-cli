package secrethor

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var version = "v0.0.1"

type UsedSecret struct {
	Namespace string
	Name      string
	UsedBy    []string
}

type OrphanedSecret struct {
	Namespace string
	Name      string
}

func Check(namespace string, output string, verbose bool) error {
	PrintBanner(version)
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return fmt.Errorf("failed to build kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create k8s client: %v", err)
	}

	ctx := context.TODO()
	targetNamespace := namespace
	if targetNamespace == "" || targetNamespace == "all" {
		targetNamespace = metav1.NamespaceAll
	}

	secrets, err := clientset.CoreV1().Secrets(targetNamespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("failed to list secrets: %v", err)
	}

	usedSecrets := map[string][]string{}

	workloads := []struct {
		name string
		list func() ([]v1.PodSpec, []metav1.ObjectMeta, error)
	}{
		{"Deployments", func() ([]v1.PodSpec, []metav1.ObjectMeta, error) {
			list, err := clientset.AppsV1().Deployments(targetNamespace).List(ctx, metav1.ListOptions{})
			if err != nil {
				return nil, nil, err
			}
			specs, metas := make([]v1.PodSpec, 0), make([]metav1.ObjectMeta, 0)
			for _, d := range list.Items {
				specs = append(specs, d.Spec.Template.Spec)
				metas = append(metas, d.ObjectMeta)
			}
			return specs, metas, nil
		},
		},
		{"StatefulSets", func() ([]v1.PodSpec, []metav1.ObjectMeta, error) {
			list, err := clientset.AppsV1().StatefulSets(targetNamespace).List(ctx, metav1.ListOptions{})
			if err != nil {
				return nil, nil, err
			}
			specs, metas := make([]v1.PodSpec, 0), make([]metav1.ObjectMeta, 0)
			for _, s := range list.Items {
				specs = append(specs, s.Spec.Template.Spec)
				metas = append(metas, s.ObjectMeta)
			}
			return specs, metas, nil
		},
		},
		{"DaemonSets", func() ([]v1.PodSpec, []metav1.ObjectMeta, error) {
			list, err := clientset.AppsV1().DaemonSets(targetNamespace).List(ctx, metav1.ListOptions{})
			if err != nil {
				return nil, nil, err
			}
			specs, metas := make([]v1.PodSpec, 0), make([]metav1.ObjectMeta, 0)
			for _, d := range list.Items {
				specs = append(specs, d.Spec.Template.Spec)
				metas = append(metas, d.ObjectMeta)
			}
			return specs, metas, nil
		},
		},
		{"ReplicaSets", func() ([]v1.PodSpec, []metav1.ObjectMeta, error) {
			list, err := clientset.AppsV1().ReplicaSets(targetNamespace).List(ctx, metav1.ListOptions{})
			if err != nil {
				return nil, nil, err
			}
			specs, metas := make([]v1.PodSpec, 0), make([]metav1.ObjectMeta, 0)
			for _, r := range list.Items {
				specs = append(specs, r.Spec.Template.Spec)
				metas = append(metas, r.ObjectMeta)
			}
			return specs, metas, nil
		},
		},
		{"CronJobs", func() ([]v1.PodSpec, []metav1.ObjectMeta, error) {
			list, err := clientset.BatchV1().CronJobs(targetNamespace).List(ctx, metav1.ListOptions{})
			if err != nil {
				return nil, nil, err
			}
			specs, metas := make([]v1.PodSpec, 0), make([]metav1.ObjectMeta, 0)
			for _, cj := range list.Items {
				specs = append(specs, cj.Spec.JobTemplate.Spec.Template.Spec)
				metas = append(metas, cj.ObjectMeta)
			}
			return specs, metas, nil
		},
		},
		{"Jobs", func() ([]v1.PodSpec, []metav1.ObjectMeta, error) {
			list, err := clientset.BatchV1().Jobs(targetNamespace).List(ctx, metav1.ListOptions{})
			if err != nil {
				return nil, nil, err
			}
			specs, metas := make([]v1.PodSpec, 0), make([]metav1.ObjectMeta, 0)
			for _, j := range list.Items {
				specs = append(specs, j.Spec.Template.Spec)
				metas = append(metas, j.ObjectMeta)
			}
			return specs, metas, nil
		},
		},
		{"Pods", func() ([]v1.PodSpec, []metav1.ObjectMeta, error) {
			list, err := clientset.CoreV1().Pods(targetNamespace).List(ctx, metav1.ListOptions{})
			if err != nil {
				return nil, nil, err
			}
			specs, metas := make([]v1.PodSpec, 0), make([]metav1.ObjectMeta, 0)
			for _, p := range list.Items {
				specs = append(specs, p.Spec)
				metas = append(metas, p.ObjectMeta)
			}
			return specs, metas, nil
		},
		},
	}

	for _, w := range workloads {
		specs, metas, err := w.list()
		if err != nil {
			return fmt.Errorf("failed to list %s: %v", w.name, err)
		}
		for i := range specs {
			checkRefs(metas[i], specs[i], w.name, usedSecrets, verbose)
		}
	}

	var usedList []UsedSecret
	var orphanedList []OrphanedSecret
	usedCount := 0
	orphanedCount := 0

	for _, s := range secrets.Items {
		key := fmt.Sprintf("%s/%s", s.Namespace, s.Name)
		if refs, ok := usedSecrets[key]; ok {
			usedList = append(usedList, UsedSecret{
				Namespace: "🔒  " + s.Namespace,
				Name:      s.Name,
				UsedBy:    refs,
			})
			usedCount++
		} else {
			orphanedList = append(orphanedList, OrphanedSecret{
				Namespace: "❗  " + s.Namespace,
				Name:      s.Name,
			})
			orphanedCount++
		}
	}

	if len(usedList) > 0 {
		fmt.Println("\nIn-use Secrets")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"NAMESPACE", "NAME", "USED BY"})
		table.SetAutoWrapText(false)
		table.SetBorder(false)
		table.SetRowSeparator(" ")
		table.SetCenterSeparator(" ")
		table.SetColumnSeparator(" ")
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT})
		for _, s := range usedList {
			table.Append([]string{s.Namespace, s.Name, strings.Join(s.UsedBy, ", ")})
		}
		table.Render()
	}

	// ORPHANED TABLE
	if len(orphanedList) > 0 {
		fmt.Println("\nOrphaned Secrets")
		orphanTable := tablewriter.NewWriter(os.Stdout)
		orphanTable.SetHeader([]string{"NAMESPACE", "NAME"})
		orphanTable.SetAutoWrapText(false)
		orphanTable.SetBorder(false)
		orphanTable.SetRowSeparator(" ")
		orphanTable.SetCenterSeparator(" ")
		orphanTable.SetColumnSeparator(" ")
		orphanTable.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		orphanTable.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT})
		for _, o := range orphanedList {
			orphanTable.Append([]string{o.Namespace, o.Name})
		}
		orphanTable.Render()
	}

	fmt.Println("\nSummary")
	fmt.Printf("🔑  Secrets in total:   %d\n", len(secrets.Items))
	fmt.Printf("🔒  Secrets in use:     %d\n", usedCount)
	fmt.Printf("❗  Orphaned secrets:   %d\n", orphanedCount)

	return nil
}

func checkRefs(objMeta metav1.ObjectMeta, podSpec v1.PodSpec, kind string, usedSecrets map[string][]string, verbose bool) {
	ref := fmt.Sprintf("%s/%s", kind, objMeta.Name)
	containers := append(podSpec.Containers, podSpec.InitContainers...)
	for _, c := range containers {
		for _, env := range c.EnvFrom {
			if env.SecretRef != nil && env.SecretRef.Name != "" {
				key := fmt.Sprintf("%s/%s", objMeta.Namespace, env.SecretRef.Name)
				if verbose {
					fmt.Printf("[DEBUG] SecretRef via EnvFrom: %s → %s\n", key, ref)
				}
				usedSecrets[key] = appendIfMissing(usedSecrets[key], ref)
			}
		}

		for _, env := range c.Env {
			if env.ValueFrom != nil && env.ValueFrom.SecretKeyRef != nil && env.ValueFrom.SecretKeyRef.Name != "" {
				key := fmt.Sprintf("%s/%s", objMeta.Namespace, env.ValueFrom.SecretKeyRef.Name)
				if verbose {
					fmt.Printf("[DEBUG] SecretRef via Env.ValueFrom: %s → %s\n", key, ref)
				}
				usedSecrets[key] = appendIfMissing(usedSecrets[key], ref)
			}
		}
	}

	for _, vol := range podSpec.Volumes {
		if vol.Secret != nil && vol.Secret.SecretName != "" {
			key := fmt.Sprintf("%s/%s", objMeta.Namespace, vol.Secret.SecretName)
			if verbose {
				fmt.Printf("[DEBUG] SecretRef via Volume: %s → %s\n", key, ref)
			}
			usedSecrets[key] = appendIfMissing(usedSecrets[key], ref)
		}
	}

	for _, ips := range podSpec.ImagePullSecrets {
		if ips.Name != "" {
			key := fmt.Sprintf("%s/%s", objMeta.Namespace, ips.Name)
			if verbose {
				fmt.Printf("[DEBUG] SecretRef via ImagePullSecret: %s → %s\n", key, ref)
			}
			usedSecrets[key] = appendIfMissing(usedSecrets[key], ref)
		}
	}
}
