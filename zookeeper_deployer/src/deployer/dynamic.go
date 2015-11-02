package deployer

import (
	"os"
	"os/exec"
	"log"
	"strings"
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


func NewDynmaicDeployer(myid string){
	
	return &DynamicDeployer{Deployer{Entries: make([]ServerEntry, 0, DEFAULT_CLUSTER_SIZE), MyID: myid}}
}

func (this *DynamicDeployer) BuildEntries(){
	
	err := this.ImportExistingServerEntries();
	
	if err != nil{
		log.Fatal(err)
	}
	
	
	My_zk_peer_url = os.Getenv(MY_ZK_PEER_URL)
	
	if My_zk_peer_url == nil{
		log.Fatal("Please set the environment "+ MY_ZK_PEER_URL +" before running container")
	}
	
	
	My_client_port = os.Getenv(MY_CLIENT_PORT)
	
	if My_client_port == nil{
		log.Fatal("Please set the environment "+ MY_CLIENT_PORT +" before running container")
	}
	
	My_id = os.Getenv(MYID)
	
	if My_id == nil{
		log.Fatal("Please set the environment "+ MYID +" before running container")
	}
	
	
	entry_value := "server."+this.MyID+"="+My_zk_peer_url+":observer;"+My_client_port
	
	
	this.Entries = append(this.Entries, entry_value)
	}


func (this *DynamicDeployer) ImportExistingServerEntries() error{
	
	cmd := ZOOKEEPER_CLI + " -server " + os.Getenv(ZK_LEADER_URL)  " get /zookeeper/config|grep ^server"
	out, err := exec.Command(cmd).Output()
	
	return err
	
	entries := strings.Split(out, "\n")
	
	
	for _, entry := range entries {
		append(this.Entries, ServerEntry{entry})
	}
	
	return nil
}