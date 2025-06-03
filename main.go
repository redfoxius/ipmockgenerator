package main

import (
	"flag"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	// parse command line flags
	maskPtr := flag.String("mask", "212.121.*.*", "IP mask for gen, a string.")
	numPtr := flag.Int("num", 1000, "number of IPs in file, an int.")
	outPtr := flag.String("out", "ip.txt", "output file name, a string.")
	flag.Parse()

	log.Println("mask:", *maskPtr)
	log.Println("numb:", *numPtr)
	log.Println("file:", *outPtr)

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(*outPtr, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	ips := generateIPs(*maskPtr, *numPtr)
	for i := 0; i < *numPtr; i++ {
		if _, err := f.Write([]byte(ips[i])); err != nil {
			log.Fatal(err)
		}
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	log.Printf("Generation took %s", elapsed)
}

func generateIPs(mask string, num int) []string {
	elems := strings.Split(mask, ".")
	result := make([]string, num)
	if validateMask(elems) {
		for i := 0; i < num; i++ {
			res := append(elems[:0:0], elems...)

			for j := 0; j < 4; j++ {
				if res[j] == "*" {
					res[j] = strconv.Itoa(randRange(0, 255))
				}
			}

			result[i] = strings.Join(res, ".") + "\n"
		}
	}

	return result
}

func randRange(min, max int) int {
	return rand.IntN(max+1-min) + min
}

func validateMask(elems []string) bool {
	if len(elems) != 4 {
		log.Fatal("mask length is not 4")
	}
	for i := 0; i < 4; i++ {
		if elems[i] != "*" {
			if n, err := strconv.Atoi(elems[i]); err != nil || n > 255 || n < 0 {
				log.Fatal("mask is not valid")
			}
		}
	}
	return true
}
