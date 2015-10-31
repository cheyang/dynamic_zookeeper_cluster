package deployer

import (

)

type StaticDeployer struct{
	Deployer
}

func (this *StaticDeployer) BuildEntries(){
	
	var zookeeperIndex []int = make([]int, 0, len(os.Environ()))
	
	}