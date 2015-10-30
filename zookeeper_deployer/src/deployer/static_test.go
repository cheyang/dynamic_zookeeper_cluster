package deployer

import (
    "testing"
    "os"
    "fmt"
)

func Test_BuildEntries(t *testing.T) {
		os.Setenv("ADDITIONAL_ZOOKEEPER_4", "server.4=localhost:2888:3888" )
		os.Setenv("ADDITIONAL_ZOOKEEPER_1", "server.1=localhost:2888:3888" )
		os.Setenv("ADDITIONAL_ZOOKEEPER_3", "server.3=localhost:2888:3888" )
		os.Setenv("ADDITIONAL_ZOOKEEPER_2", "server.2=localhost:2888:3888" )
		
		deployer := NewStaticDeployer()
		
		deployer.BuildEntries()
		
		
		to_compare := []ServerEntry{ServerEntry{Entry:"server.1=localhost:2888:3888"},
									ServerEntry{Entry:"server.2=localhost:2888:3888"},
									ServerEntry{Entry:"server.3=localhost:2888:3888"},
									ServerEntry{Entry:"server.4=localhost:2888:3888"}}
		
		fmt.Println(deployer.Entries) 
		
		/**
		if to_compare != deployer.Entries {
			
			t.Error("Build Entries failed")
		}
		*/
		fmt.Println(to_compare) 
}

