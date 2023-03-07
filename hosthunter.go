package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Version: "1.0.0",
		Name:    "hosthunter",
		Usage:   "Convert hostnames to IP addresses",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file",
				Aliases:  []string{"f"},
				Usage:    "Path to file with a list of hostnames",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {
			fileName := cCtx.String("file")

			hostsFile, err := os.Open(fileName)
			if err != nil {
				fmt.Println(err)
			}
			defer hostsFile.Close()

			fileScanner := bufio.NewScanner(hostsFile)

			fileScanner.Split(bufio.ScanLines)

			for fileScanner.Scan() {
				host := fileScanner.Text()
				ips, _ := net.LookupIP(host)
				for _, ip := range ips {
					if ipv4 := ip.To4(); ipv4 != nil {
						fmt.Println(ipv4)
					}
				}
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
