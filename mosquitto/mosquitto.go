package mosquitto

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"syscall"
)

const defaultFailedCode = 1

func RunCommand(name string, args ...string) (stdout string, stderr string, exitCode int) {
	log.Println("run command:", name, args)
	var outbuf, errbuf bytes.Buffer
	cmd := exec.Command(name, args...)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	stdout = outbuf.String()
	stderr = errbuf.String()

	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		} else {
			log.Printf("Could not get exit code for failed program: %v, %v", name, args)
			exitCode = defaultFailedCode
			if stderr == "" {
				stderr = err.Error()
			}
		}
	} else {
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
	}
	log.Printf("command result, stdout: %v, stderr: %v, exitCode: %v", stdout, stderr, exitCode)
	return
}

func WriteToAclFile (username string) {
	f, err := os.OpenFile("mosquitto.acl",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	data := "user " + username + "\r\ntopic " + username + "/#\r\n"
	if _, err := f.WriteString(data); err != nil {
		log.Println(err)
	}
}

func DeleteFromAclFile (username string) {

	file, err := os.Open("mosquitto.acl")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	tmp, err := ioutil.TempFile("d:/main/working/mosquitto", "replace-*")
	if err != nil {
		log.Fatal(err)
	}
	defer tmp.Close()

	if err := replace(file, tmp, username); err != nil {
		log.Fatal(err)
	}

	if err := tmp.Close(); err != nil {
		log.Fatal(err)
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

	if err := os.Rename(tmp.Name(), "mosquitto.acl"); err != nil {
		log.Fatal(err)
	}
}

func replace(r io.Reader, w io.Writer, username string) error {
	firstLine := "user " + username
	secondLine := "topic " + username +"/#"
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		if line != firstLine && line != secondLine {
			if _, err := io.WriteString(w, line+"\n"); err != nil {
				return err
			}
		}
	}
	return sc.Err()
}