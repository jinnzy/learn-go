package main

import (
	"fmt"
	yaml2 "github.com/ghodss/yaml"
)

func main()  {
	str := []byte(`- job_name: monitoring/prometheus/0
  honor_labels: false
  kubernetes_sd_configs:
  - role: endpoints
    namespaces:
      names:
      - monitoring
  scrape_interval: 30s
  relabel_configs:
  - action: keep
    source_labels:
    - __meta_kubernetes_service_label_prometheus
    regex: k8s
  - action: keep
    source_labels:
    - __meta_kubernetes_endpoint_port_name
    regex: web
  - source_labels:
    - __meta_kubernetes_endpoint_address_target_kind
    - __meta_kubernetes_endpoint_address_target_name
    separator: ;
    regex: Node;(.*)
    replacement: ${1}
    target_label: node
  - source_labels:
    - __meta_kubernetes_endpoint_address_target_kind
    - __meta_kubernetes_endpoint_address_target_name
    separator: ;
    regex: Pod;(.*)
    replacement: ${1}
    target_label: pod
  - source_labels:
    - __meta_kubernetes_namespace
    target_label: namespace
  - source_labels:
    - __meta_kubernetes_service_name
    target_label: service
  - source_labels:
    - __meta_kubernetes_pod_name
    target_label: pod
  - source_labels:
    - __meta_kubernetes_service_name
    target_label: job
    replacement: ${1}
  - target_label: endpoint
    replacement: web
`)
	j2,err := yaml2.YAMLToJSON(str)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(string(j2))
	y2,err := yaml2.JSONToYAML(j2)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(string(y2))
}
