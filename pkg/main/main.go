package main

import (
	"github.com/0yukali0/yunikorn-rest-client/pkg/client"

	"net/http"
)

func main() {
	_ = client.NewClusterClient("")
	http.HandleFunc("/cluster", client.PrintCluster)

	http.ListenAndServe(":8090", nil)
}
