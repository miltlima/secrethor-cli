package secrethor

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func SearchSecret(name string, namespace string) error {
	PrintBanner(version)
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return fmt.Errorf("failed to build kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create k8s client: %v", err)
	}

	ns := namespace
	if ns == "" || ns == "all" {
		ns = metav1.NamespaceAll
	}

	secrets, err := clientset.CoreV1().Secrets(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("failed to list secrets: %v", err)
	}

	found := false
	for _, s := range secrets.Items {
		if s.Name == name {
			fmt.Printf("✅ Found secret: %s/%s\n", s.Namespace, s.Name)
			fmt.Printf("🔐 Type: %s\n", s.Type)
			fmt.Printf("📦 Data keys: %v\n", getDataKeys(s))
			found = true
		}
	}

	if !found {
		fmt.Println("❌ Secret not found.")
	}

	return nil
}

func getDataKeys(secret v1.Secret) []string {
	var keys []string
	for k := range secret.Data {
		keys = append(keys, k)
	}
	return keys
}
