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
	for i := 0; i < *numPtr; i++ {
		if _, err := f.Write([]byte(generateIP(*maskPtr) + "\n")); err != nil {
			log.Fatal(err)
		}
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	log.Printf("Generation took %s", elapsed)
}

func generateIP(mask string) string {
	elems := strings.Split(mask, ".")
	if len(elems) != 4 {
		log.Fatal("mask length is not 4")
	}

	result := make([]string, 4)
	for i := 0; i < 4; i++ {
		if elems[i] != "*" {
			if n, err := strconv.Atoi(elems[i]); err != nil || n > 255 || n < 0 {
				log.Fatal("mask is not valid")
			}

			result[i] = elems[i]
			continue
		}
		result[i] = strconv.Itoa(randRange(0, 255))
	}

	return strings.Join(result, ".")
}

func randRange(min, max int) int {
	return rand.IntN(max+1-min) + min
}
