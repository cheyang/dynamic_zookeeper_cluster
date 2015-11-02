package deployer

import (
	"os/exec"
	"bytes"
	"fmt"
	"strings"
)


func Exec(cmd string, args []string) (output string, err error){
	
	to_exec_cmd := cmd + " " + strings.Join(args, " ")
	
	fmt.Println(to_exec_cmd)
	
//	command := exec.Command(cmd, args...)

    command := exec.Command("/bin/sh", "-c", to_exec_cmd)
	
	
	w := bytes.NewBuffer(nil)
    command.Stderr = w
    
    w1 := bytes.NewBuffer(nil)
    
    command.Stdout = w1
    
    if err = command.Run(); err != nil {
        fmt.Printf("Run returns: %s\n", err)
        fmt.Printf("Stderr: %s\n", string(w.Bytes()))
        return output, err
    }
    
    
    
    output = string(w1.Bytes())

	return output, err
}
