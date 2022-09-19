package main

import (
	"github.com/afdesk/poc-fullclassify/classificator"
	"io"
	"log"
	"os"
	"path"
	"runtime/pprof"
)

const folder = "./perl-licenses"

func handleAllLicenses() {
	files, err := os.ReadDir(folder)
	if err != nil {
		log.Fatalf("can't ReadDir %q. error: %v", folder, err)
	}
	for _, f := range files {
		lic, err := os.Open(path.Join(folder, f.Name()))
		if err != nil {
			log.Printf("[ERROR] can't open %q: %v", f.Name(), err)
			continue
		}
		classificator.Classify(lic)
	}
}

func handleFullClassify() {
	files, err := os.ReadDir(folder)
	if err != nil {
		log.Fatalf("can't ReadDir %q. error: %v", folder, err)
	}
	for _, f := range files {
		filepath := path.Join(folder, f.Name())
		lic, err := os.Open(filepath)
		if err != nil {
			log.Printf("[ERROR] can't open %q: %v", f.Name(), err)
			continue
		}
		content, err := io.ReadAll(lic)
		if err != nil {
			log.Printf("[ERROR] can't open %q: %v", f.Name(), err)
			continue
		}
		classificator.FullClassify(filepath, content)
	}

}

func main() {
	handleAllLicenses()
	handleFullClassify()

	fMem, err := os.Create("mem.profile")
	if err != nil {
		panic("could not create memory profile: " + err.Error())
	}
	defer fMem.Close() // error handling omitted for example
	if err := pprof.WriteHeapProfile(fMem); err != nil {
		panic("could not write memory profile: " + err.Error())
	}
}
