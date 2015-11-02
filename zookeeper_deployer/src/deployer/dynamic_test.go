package deployer

import (
    "testing"
    "fmt"
)

func Test_ImportExistingServerEntries(t *testing.T) {
	
	deployer := NewDynmaicDeployer("12")

	deployer.ImportExistingServerEntries()
	
	fmt.Println(deployer.Entries)
}

