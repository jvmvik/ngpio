package ngpio

import (
	"flag"
	"log"
	"os"
	"os/user"
)

/*
 *
// echo 79 > /sys/class/gpio/export
// echo out > /sys/class/gpio/gpio79/direction
// echo 1 > /sys/class/gpio/gpio79/value
// echo 0 > /sys/class/gpio/gpio79/value
*/

func main() {
	//isRoot()

	physicalOutput := flag.Int("p", -1, "physical output port number on the card")
	on := flag.Bool("on", false, "turn on")
	off := flag.Bool("off", false, "turn off")

	flag.Parse()

	exit := false
	if *physicalOutput == -1 {
		log.Printf("pin must be defined: -p 12")
		exit = true
	}
	if !*on && !*off {
		log.Println("either -on or -off must be provided")
		exit = true
	}
	if exit {
		os.Exit(1)
	}

	spec := nano()
	port, err := spec.FindPortByOutput(*physicalOutput)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Printf("found physical port: %d\n", port.LogicalID)

	port.Output()

	if *on {
		port.High()
	}
	if *off {
		port.Low()
	}

	log.Println("done")
}

func isRoot() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	if u.Username != "root" {
		log.Fatal("you must be logged as root")
		os.Exit(1)
	}
}
