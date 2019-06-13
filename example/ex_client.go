// modeminfo collects and displays information related to the modem and its
// current configuration.
//
// This serves as an example of how interact with a modem, as well as
// providing information which may be useful for debugging.
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/wkarasz/goat-modem/at"
	"github.com/wkarasz/goat-modem/serial"
	"github.com/wkarasz/goat-modem/trace"
)

func main() {
	dev := flag.String("d", "/dev/ttyUSB0", "path to modem device")
	baud := flag.Int("b", 115200, "baud rate")
	timeout := flag.Duration("t", 400*time.Millisecond, "command timeout period")
	verbose := flag.Bool("v", false, "log modem interactions")
	flag.Parse()
	m, err := serial.New(*dev, *baud)
	if err != nil {
		log.Println(err)
		return
	}
	defer m.Close()
	var mio io.ReadWriter = m
	if *verbose {
		mio = trace.New(m, log.New(os.Stdout, "", log.LstdFlags))
	}
	a := at.New(mio)
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	err = a.Init(ctx)
	cancel()
	if err != nil {
		log.Println(err)
		return
	}
	cmds := []string{
		"I",
		"+GCAP",
		"+CMEE=2",
		"+CGMI",
		"+CGMM",
		"+CGMR",
		"+CGSN",
		"+CSQ",
		"+CIMI",
		"+CREG?",
		"+CNUM",
		"+CPIN?",
		"+CEER",
		"+CSCA?",
		"+CSMS?",
		"+CSMS=?",
		"+CPMS=?",
		"+CNMI?",
		"+CNMI=?",
		"+CNMA=?",
		"+CMGF=?",
	}
	for _, cmd := range cmds {
		ctx, cancel := context.WithTimeout(context.Background(), *timeout)
		info, err := a.Command(ctx, cmd)
		cancel()
		fmt.Println("AT" + cmd)
		if err != nil {
			fmt.Printf(" %s\n", err)
		} else {
			for _, l := range info {
				fmt.Printf(" %s\n", l)
			}
		}
	}
	
	/* Works but costs money $0.05/text
	ctx, cancel = context.WithTimeout(context.Background(), *timeout)
	info, err := a.SMSCommand(ctx, "+cmgs=\"+15165551234\"", "hello go")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		for _, l := range info {
			fmt.Printf("Info: %s\n", l)
		}
	}
	*/
}
