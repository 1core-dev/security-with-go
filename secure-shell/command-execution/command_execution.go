package main

import (
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

var username = "username"
var host = "example.com:22"
var privateKeyFile = "/Users/<USER>/.ssh/id_rsa"
var commandToExecute = "hostname"

func getKeySigner(privateKeyFile string) ssh.Signer {
	privateKeyData, err := os.ReadFile(privateKeyFile)
	if err != nil {
		log.Fatal("Error loading private key file. ", err)
	}

	privateKey, err := ssh.ParsePrivateKey(privateKeyData)
	if err != nil {
		log.Fatal("Error parsing private key. ", err)
	}

	return privateKey
}

func main() {
	privateKey := getKeySigner(privateKeyFile)
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(privateKey),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Fatal("Error dialing server. ", err)
	}

	// Multiple sessions per client are allowed
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Pipe the session output directly to standard output
	// Thanks to the convenience of writer interface
	session.Stdout = os.Stdout
	err = session.Run(commandToExecute)
	if err != nil {
		log.Fatal("Error executing command. ", err)
	}
}
