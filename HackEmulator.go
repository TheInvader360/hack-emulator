package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/TheInvader360/hack-emulator/client"
	"github.com/TheInvader360/hack-emulator/hack"
	"github.com/TheInvader360/hack-emulator/hack/stub"

	"github.com/faiface/pixel/pixelgl"
)

var (
	path string
	vm   hack.Hack
	c    client.Client
)

func main() {
	flag.StringVar(&path, "path", "./roms/Fill.hack", "path to rom file")
	flag.Parse()
	vm = &stub.Hack{}
	loadRom()
	c = client.Client{VM: &vm}
	pixelgl.Run(c.Run)
}

func loadRom() {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) //TODO load binary string into vm instruction memory
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
