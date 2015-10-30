package deployer

import (
	"os"
	"regexp"
	"strings"
	"sort"
	"fmt"
)

type StaticDeployer struct{
	Deployer
}


func NewStaticDeployer() *StaticDeployer{
	
	
	return &StaticDeployer{entries: nil}
	
}

/**
* Build entries from environment variables
*/
func (this *StaticDeployer) BuildEntries(){
	
	var zookeeperIndex []int = make([]int, len(os.Environ))
	
	re := regexp.MustCompile("ADDITIONAL_ZOOKEEPER_[\\d]+$")
	
	for _, env := range os.Environ() {
		
		if re.MatchString(env){
			re_d := regexp.MustCompile("[\\d]+$")
			
			i, err := strconv.Atoi(re_d.FindString(env))
			
			if err !=nil{
				 fmt.Println(err)
				 return
				}
			
			append(zookeeperIndex, i)
				
		}
	}
	
	
	sort.Ints(zookeeperIndex)
	
	this.Entries = make([]ServerEntry, len(zookeeperIndex))
	
	
	for _, key := range zookeeperIndex {
		 value = os.Getenv("ADDITIONAL_ZOOKEEPER_"+strconv.itoa(key))
		 
		 append(this.Entries, ServerEntry{entry: value})
	}
	
}

func (this *StaticDeployer) GenerateTemplate(){
	
}

func (this *StaticDeployer) Deploy(){
	
}

