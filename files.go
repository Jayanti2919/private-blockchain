package main

import (
	"bufio"
	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

func FileUpload(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	sh := shell.NewLocalShell()
	cid, err := sh.Add(file)
	if err != nil {
		print(err)
		return "", err
	}
	fmt.Println("CID: ", cid)

	return cid, nil
}

func FileDownload(cid string) error {
	sh := shell.NewLocalShell()
	outdir := "A:/private_blockchain"
	err2 := sh.Get(cid, outdir)
	if err2 != nil {
		fmt.Print("Error: ", err2)
		return err2
	}

	retrievedFile, err3 := os.Open(outdir + "/" + cid)
	if err3 != nil {
		fmt.Println(err3)
		return err3
	}
	defer retrievedFile.Close()
	scanner := bufio.NewScanner(retrievedFile)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return err
	}
	return nil
}
