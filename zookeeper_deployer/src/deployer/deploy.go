package deployer

import (
)

const (
	PARTICIPANT = "participant"
	OBSERVER = "observer"
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

func (this *Deployer) GenerateTemplate(){
	
}

func (this *Deployer) Deploy(){
	this.BuildEntries()
	
	this.GenerateTemplate()
}

/**
func Deploy(deployType string){
	
	var deployer Deployer
	
	switch Strings.ToLower(deployType){
		
		case "static":
			
	}
}
*/