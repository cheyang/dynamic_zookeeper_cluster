package deployer

import (
	"io"
	"fmt"
	"strconv"
	"os/exec"
)

const (
	PARTICIPANT = "participant"
	OBSERVER = "observer"
	MYID_FILE = "/tmp/zookeeper/myid"
	MYID = "MYID"
)

type ServerEntry struct{
	Entry, 	PeerUrl, Role string
	ClientPort int	
}

type Deployer struct{
	Entries []ServerEntry
}

func (this *Deployer) BuildEntries(){
	
}

func (this *Deployer) GenerateTemplate(dynamicFile string) error {
	
	file, err := os.Create(dynamicFile)
	
	if err != nil{
		fmt.Println("Failed to create the configuration file: ", dynamicFile)
		return err
	}
	
	defer file.close()
	
	
	for _, entry range this.Entries{
		
		if entry.Entry != nil{
			file.WriteString(entry.Entry+"\n")
		}else{
			file.WriteString(entry.PeerUrl+":"+PARTICIPANT+";0.0.0.0:"+strconv.Itoa(entry.PeerUrl)+"\n")
		}
	}
	
	return nil
}

func (this *Deployer) Deploy(dynamicFile string){
	
	
	this.BuildEntries()
	
	err := this.GenerateTemplate(dynamicFile)
	
	if err != nil{
		fmt.Println(err)
		
		return
	}
	
}

func (this *Deployer) CreateMyID() error{
	if _, err := os.Stat(MYID_FILE); err == nil{
		fmt.Printf("file exists; returning...")
		return nil
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