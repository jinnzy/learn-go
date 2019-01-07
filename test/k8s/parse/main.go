package main

import (
	"k8s.io/client-go/kubernetes/scheme"
	"fmt"
)

var deployment = `
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
name: my-nginx
spec:
replicas: 2
template:
  metadata:
    labels:
      run: my-nginx
  spec:
    containers:
    - name: my-nginx
      image: nginx
      ports:
      - containerPort: 80
`

func main()  {
	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj,_,_ := decode([]byte(deployment),nil,nil)
	fmt.Printf("%#v\n",obj)
}
