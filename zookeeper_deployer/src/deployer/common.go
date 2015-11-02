package deployer

import (
	"os/exec"
	"bytes"
	"fmt"
)


func Exec(cmd string, args []string) (output string, err error){
	command := exec.Command(cmd, args...)
	
	
	w := bytes.NewBuffer(nil)
    command.Stderr = w
    
    w1 := bytes.NewBuffer(nil)
    
    command.Stdout = w1
    
    if err = command.Run(); err != nil {
        fmt.Printf("Run returns: %s\n", err)
        return output, err
    }
    
    fmt.Printf("Stderr: %s\n", string(w.Bytes()))
    
    output = string(w.Bytes())

	return output, err
}
