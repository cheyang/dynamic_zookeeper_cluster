package deployer

import (
	"os"
	"regexp"
	"strings"
	"strconv"
	"sort"
	"fmt"
)

type StaticDeployer struct{
	Deployer
}


func NewStaticDeployer() *StaticDeployer{
	
	
	return &StaticDeployer{Deployer{Entries: nil}}
	
}

/**
* Build entries from environment variables
*/
func (this *StaticDeployer) BuildEntries(){
	
	var zookeeperIndex []int = make([]int, 0)
	
	re := regexp.MustCompile("ADDITIONAL_ZOOKEEPER_[\\d]+$")
	
	for _, env := range os.Environ() {
		
		envPair := strings.Split(env,  "=")
		
		if re.MatchString(envPair[0]){
			re_d := regexp.MustCompile("[\\d]+$")
			
			i, err := strconv.Atoi(re_d.FindString(envPair[0]))
			
			if err !=nil{
				 fmt.Println(err)
				 return
				}
			
			
			zookeeperIndex = append(zookeeperIndex, i)
				
		}
	}
		
	
	sort.Ints(zookeeperIndex)
	
	fmt.Println(len(zookeeperIndex))
	
	this.Entries = make([]ServerEntry, 4)
	
	
	for _, key := range zookeeperIndex {
		 value := os.Getenv("ADDITIONAL_ZOOKEEPER_"+strconv.Itoa(key))
		 
		 fmt.Println(value)
		 
		 this.Entries= append(this.Entries, ServerEntry{Entry: value})
	}
	
}

func (this *StaticDeployer) GenerateTemplate(){
	
}

func (this *StaticDeployer) Deploy(){
	
}

