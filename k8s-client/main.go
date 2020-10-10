package main

import (
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"os"
	"path/filepath"

	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {
	var (
		kubeconfig *string
		dropNamespace = make(map[string]struct{})
	)
	dropNamespace["wuzh"] = struct {}{}
	dropNamespace["velero"] = struct {}{}
	dropNamespace["test"] = struct {}{}
	dropNamespace["redis"] = struct {}{}
	dropNamespace["postgres"] = struct {}{}
	dropNamespace["kube-system"] = struct {}{}
	dropNamespace["kube-public"] = struct {}{}
	dropNamespace["jenkins"] = struct {}{}
	dropNamespace["integration"] = struct {}{}
	dropNamespace["eolinker"] = struct {}{}
	dropNamespace["default"] = struct {}{}
	dropNamespace["cicd"] = struct {}{}
	dropNamespace["cattle-system"] = struct {}{}
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	//pods, err := clientset.CoreV1().Pods("fstest").List(metav1.ListOptions{})
	//if err != nil {
	//	panic(err.Error())
	//}
	//fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	// 获取namespace
	namespaceList, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	var deploymentInfo = make(map[string][]int32)

	for _,nsList := range namespaceList.Items {
		if _,ok := dropNamespace[nsList.Name]; ok {
			continue
		}

		deployment,err := clientset.AppsV1().Deployments(nsList.Name).List(metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		// 遍历deployment
		for _,i := range deployment.Items {
			var b int32 = 0
			if i.Spec.Replicas != &b {
				deploymentInfo[getDeploymentName(nsList.Name,i.Labels["app"])] = nil
			}
		}

		// 遍历svc
		svc,err := clientset.CoreV1().Services(nsList.Name).List(metav1.ListOptions{})
		if err != nil {
			panic(err)
		}

		for _,i := range svc.Items {
			if _,ok := deploymentInfo[getDeploymentName(nsList.Name,i.Spec.Selector["app"])];ok {
				for _,v := range i.Spec.Ports{
					deploymentInfo[getDeploymentName(nsList.Name,i.Spec.Selector["app"])] = append(deploymentInfo[getDeploymentName(nsList.Name,i.Spec.Selector["app"])],v.NodePort)
				}
			}
		}

	}
	//data,err := yaml.Marshal(deploymentInfo)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(data))
	for _,v := range deploymentInfo {
		for _,v1 := range v {
			fmt.Println(v1)
		}
	}
}

func getDeploymentName(namespace,deployName string) string {
	return fmt.Sprintf("%s/%s",namespace,deployName)
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
