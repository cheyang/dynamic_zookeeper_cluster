package deployer

import (
	"io"
	"fmt"
	"strconv"
	"os"
)

const (
	PARTICIPANT = "participant"
	OBSERVER = "observer"
	MYID_FILE = "/tmp/zookeeper/myid"
	MYID_DIR ="/tmp/zookeeper"
	MYID = "MYID"
	ZK_DIR = "/opt/zookeeper"
	ZK_CLI = ZKR_DIR + "/bin/zkCli.sh"
	ZK_DYNAMIC_CONF = "/opt/zookeeper/conf/zoo.cfg.dynamic"
)

type ServerEntry struct{
	Entry string
}

type Deployer struct{
	Entries []ServerEntry
	MyID string
}

func (this *Deployer) BuildEntries(){
	
}

func (this *Deployer) GenerateTemplate() error {
	
	dynamicFile := os.Getenv()
	
	file, err := os.Create(dynamicFile)
	
	if err != nil{
		fmt.Println("Failed to create the configuration file: ", dynamicFile)
		return err
	}
	
	defer file.close()
	
	
	for _, entry := range this.Entries{
		
		if entry.Entry != nil{
			file.WriteString(entry.Entry+"\n")
		}else{
			file.WriteString(entry.PeerUrl+":"+PARTICIPANT+";0.0.0.0:"+strconv.Itoa(entry.PeerUrl)+"\n")
		}
	}
	
	return nil
}

func (this *Deployer) Deploy(dynamicFile string){
	
	err := this.CreateMyID()
	
	if err != nil{
		fmt.Println(err)
		
		return
	}
	
	
	this.BuildEntries()
	
	err := this.GenerateTemplate(dynamicFile)
	
	if err != nil{
		fmt.Println(err)
		
		return
	}
	
}

func (this *Deployer) CreateMyID() error{
	if _, err := os.Stat(MYID_FILE);  os.IsExist(err){
		fmt.Printf("file exists; returning...")
		return nil
	}
	
	mydir, err := os.Stat(MYID_DIR)
	
	if err != nil{
		if !mydir.IsDir() {
			os.MkdirAll(mydir, 0775)
		}
	}
	
	
	file, err := os.Create(MYID_FILE)
	
	
	if err != nil{
		fmt.Println("Failed to create the configuration file: ", MYID)
		return err
	}
	
	defer file.close()
	
	myid := os.Getenv(MYID)
	
	file.WriteString(myid)
	
	return nil
	
}

/**
func Deploy(deployType string){

	if _, err := os.Stat(dynamicFile); err == nil {
	
	}
	
	var deployer Deployer
	
	switch Strings.ToLower(deployType){
		
		case "static":
			
	}
}
*/