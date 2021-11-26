package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/matti/webdriver-watcher/internal/cdp"
	"github.com/matti/webdriver-watcher/internal/checker"
)

func main() {
	var errors = 0
	var cdpErrors = 0

	maxErrors, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln("invalid maxErrors")
	}
	checkEvery, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln("invalid checkEvery")
	}

	notOkExec := os.Args[3]

	for {
		ok, maybe, stage, status := checker.Check("http://localhost:9515")
		fmt.Println("webdriver", ok, maybe, stage, status, errors)
		if !ok {
			errors++
		} else {
			errors = 0
		}

		if errors > maxErrors {
			break
		}

		cdpOk, cdpStatus := cdp.Check("http://localhost:9222/json")
		fmt.Println("cdp", cdpOk, cdpStatus)

		if !cdpOk {
			cdpErrors++
		} else {
			cdpErrors = 0
		}

		if cdpErrors > maxErrors {
			break
		}

		time.Sleep(time.Second * time.Duration(checkEvery))
	}

	fmt.Println("exec")
	cmd := exec.Command(notOkExec)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	err = cmd.Start()
	if err != nil {
		log.Fatalf("command start error %y", err)
	}

	cmd.Wait()
	fmt.Println("exit")
}
