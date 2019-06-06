package main

import (
		"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"flag"
		"log"
	"fmt"
)

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

var (
	// Create a summary to track fictional interservice RPC latencies for three
	// distinct services with different latency distributions. These services are
	// differentiated via a "service" label.
	numCont = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "test_for_num",
			Help: "test_for_num_help",
		})
	)

func init() {
	// Register the summary and the histogram with Prometheus's default registry.
	// 注册指标
	prometheus.MustRegister(numCont)
}

type ClusterManager struct {
	Zone string
	NumCountDesc *prometheus.Desc
}

func (c *ClusterManager) GetNumCount() (numCountByhost map[string]int) {
	numCountByhost = map[string]int{
		"test1": int(rand.Int31n(1000)),
		"test2": int(rand.Int31n(3000)),
	}
	return
}
func (c *ClusterManager) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.NumCountDesc
}
func (c *ClusterManager) Collect(ch chan <- prometheus.Metric) {
	numContByhost := c.GetNumCount
	fmt.Println(numContByhost())
	for host,numcont := range numContByhost(){
		fmt.Println(host)
		fmt.Println(numcont)
		ch <- prometheus.MustNewConstMetric(
			c.NumCountDesc,
			prometheus.CounterValue,
			float64(numcont),
			host,
			)
	}
}
func NewClusterManager(zone string) *ClusterManager {
	return &ClusterManager{
		Zone: zone,
		NumCountDesc: prometheus.NewDesc(
			"num_total",
			"num_help",
			[]string{"host"},
			prometheus.Labels{"zone":zone},
			),
	}
}

func main() {
	flag.Parse()
	workerDB := NewClusterManager("db")
	workerCA := NewClusterManager("ca")

	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(workerDB)
	reg.MustRegister(workerCA)

	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(*addr, nil))

}
