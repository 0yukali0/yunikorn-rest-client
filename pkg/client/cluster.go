package client

import (
	"github.com/apache/yunikorn-core/pkg/webservice/dao"

	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var (
	once              sync.Once
	RestClusterClient *ClusterClient
)

func PrintCluster(w http.ResponseWriter, req *http.Request) {
	c := NewClusterClient("")
	c.Get()
	c.Print()
}

type ClusterClient struct {
	url  string
	data *dao.ClusterDAOInfo
}

func NewClusterClient(url string) *ClusterClient {
	once.Do(func() {
		RestClusterClient = &ClusterClient{
			url:  url,
			data: nil,
		}
	})
	return RestClusterClient
}

func (c *ClusterClient) Get() (*dao.ClusterDAOInfo, error) {
	resp, err := http.Get(c.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var result dao.ClusterDAOInfo
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("Can't unmarshal JSON")
	}
	c.data = &result
	return c.data, nil
}

func (c *ClusterClient) Print() {
	if c.data == nil {
		fmt.Println("Cluster data is empty")
		return
	}
	fmt.Println("###########################")
	fmt.Printf("Cluster %s has partition %s\n",
		c.data.ClusterName,
		c.data.PartitionName,
	)
	fmt.Printf("Created at %v", c.data.StartTime)
	for _, tags := range c.data.RMBuildInformation {
		fmt.Println("------tags------")
		fmt.Println(tags)
		fmt.Println("------end-------")
	}
	fmt.Println("###########################")
	return
}
