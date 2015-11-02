package deployer

import (
	"os"
	"log"
	"strings"
	"fmt"
)

const (
	MY_ZK_PEER_URL = "MY_ZK_PEER_URL"
	MY_CLIENT_PORT = "MY_CLIENT_PORT"
	ZK_LEADER_URL = "ZK_LEADER_URL"
	DEFAULT_CLUSTER_SIZE = 20
	
)

var (
	My_zk_peer_url, My_client_port string
)

type DynamicDeployer struct{
	Deployer
}


func NewDynmaicDeployer(myid string) *DynamicDeployer{
	
	return &DynamicDeployer{Deployer{Entries: make([]ServerEntry, 0, DEFAULT_CLUSTER_SIZE), MyID: myid}}
}

func (this *DynamicDeployer) BuildEntries(){
	
	err := this.ImportExistingServerEntries();
	
	if err != nil{
		log.Fatal(err)
	}
	
	
	My_zk_peer_url = os.Getenv(MY_ZK_PEER_URL)
	
	if My_zk_peer_url == ""{
		log.Fatal("Please set the environment "+ MY_ZK_PEER_URL +" before running container")
	}
	
	
	My_client_port = os.Getenv(MY_CLIENT_PORT)
	
	if My_client_port == ""{
		log.Fatal("Please set the environment "+ MY_CLIENT_PORT +" before running container")
	}
	
	
	entry_value := "server."+this.MyID+"="+My_zk_peer_url+":observer;"+My_client_port
	
	
	this.Entries = append(this.Entries, ServerEntry{entry_value})
	}


func (this *DynamicDeployer) ImportExistingServerEntries() error{
	
	
	cmd := ZK_CLI
	
	if os.Getenv(ZK_LEADER_URL) == ""{
		return errors.New("Please set the environment "+ ZK_LEADER_URL +" before running container")
	}
	

    args := []string {"-server", os.Getenv(ZK_LEADER_URL),  "get /zookeeper/config|grep ^server" }

	output, err := Exec(cmd, args)
    	
	return err
	
	fmt.Println(string(out))
	
	entries := strings.Split(string(out), "\n")
	
	
	for _, entry := range entries {
		this.Entries = append(this.Entries, ServerEntry{entry})
	}
	
	return nil
}