package deployer

import (
    "testing"
    "fmt"
    "os"
)

func Test_ImportExistingServerEntries(t *testing.T) {
	
	os.Setenv(ZK_LEADER_URL, "localhost:2181")
	
	deployer := NewDynmaicDeployer("12")

	err := deployer.ImportExistingServerEntries()
	
	fmt.Println(err)
	
	fmt.Println(deployer.Entries)
}

