package deployer

import (
    "testing"
    "fmt"
)

func Test_ImportExistingServerEntries(t *testing.T) {
	
	deployer := NewDynmaicDeployer("12")

	err := deployer.ImportExistingServerEntries()
	
	fmt.Println(err)
	
	fmt.Println(deployer.Entries)
}

