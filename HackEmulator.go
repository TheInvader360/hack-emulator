package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/TheInvader360/hack-emulator/client"
	"github.com/TheInvader360/hack-emulator/hack"
	"github.com/TheInvader360/hack-emulator/hack/impl"

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
	vm = impl.NewHack()
	loadRom()
	c = client.Client{VM: &vm}
	go vmLoop()
	pixelgl.Run(c.Run)
}

func loadRom() {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := []uint16{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word, err := strconv.ParseUint(scanner.Text(), 2, 16)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, uint16(word))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	vm.LoadRom(data)
}

func vmLoop() {
	for {
		time.Sleep(10 * time.Nanosecond)
		vm.Tick()
	}
}
