package hlogger

import (
	"fmt"
	"os"
	"sync"

	"github.com/logrusorgru/aurora"
)

var (
	mutex = &sync.Mutex{}
)

func log(level, text string) {
	if level == "INFO" {
		text = aurora.Blue("[INFO] ").String() + text
	}
	if level == "WEAK" {
		text = aurora.Yellow("[WEAK] ").String() + text
	}
	if level == "VULN" {
		text = aurora.Red("[VULN] ").String() + text
	}
	if level == "SYSTEM" {
		text = aurora.White("[*] ").String() + text
	}
	mutex.Lock()
	fmt.Fprintf(os.Stderr, text)
	mutex.Unlock()

}
