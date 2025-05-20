package main

import (
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

var username = "username"
var host = "example.com:22"
var privateKeyFile = "/Users/<USER>/.ssh/id_rsa"

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
			ssh.PublicKeys(privateKey), // Pass 1 or more key
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Fatal("Error dialing server. ", err)
	}

	log.Println(string(client.ClientVersion()))
}
