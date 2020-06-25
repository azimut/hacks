package main

import (
	"fmt"
	"time"

	"github.com/tomsteele/go-nmap"
)

// hostLine returns a bareminimun line foran nmap host
func hostLine(host nmap.Host, start string) (string, error) {
	address := " "
	hostname := " "
	if len(host.Addresses) > 0 {
		address = host.Addresses[0].Addr
	}
	if len(host.Hostnames) > 0 {
		hostname = host.Hostnames[0].Name
	}
	parsedDate, err := time.Parse("Mon Jan  2 15:04:05 2006", start)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d\t%s\t%s\t%s",
		parsedDate.Unix(), //FIXME: might need timezone, before
		host.Status.State,
		address,
		hostname), nil
}

// printNmapFile ...
func printNmap(parsed *nmap.NmapRun) error {
	for _, host := range parsed.Hosts {
		modeline, err := hostLine(host, parsed.StartStr)
		if err != nil {
			return err
		}
		if len(host.Ports) > 0 {
			for _, port := range host.Ports {
				fmt.Printf("%s\t%s\t%s\t%d\t%s\t%s\n",
					modeline,
					port.State.State,
					port.Protocol,
					port.PortId,
					port.Service.Name,
					port.Service.Product)
			}
		} else {
			fmt.Println(modeline)
		}
	}
	return nil
}
