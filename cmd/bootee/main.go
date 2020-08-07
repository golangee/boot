package main

import (
	"flag"
	"github.com/golangee/boot"
	"log"
	"os"
)

func main() {
	dir := flag.String("dir", "", "The directory root of the go module which should be processed.")
	help := flag.Bool("help", false, "shows this help")
	force := flag.Bool("force", false, "if true, forces a template deployment and overwrites your files")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	if *dir != "" {
		if err := os.Chdir(*dir); err != nil {
			log.Fatal(err)
		}
	}

	if err := boot.GenerateWithOptions(*force); err != nil {
		log.Fatal(err)
	}
}
