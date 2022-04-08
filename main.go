package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func localAddresses() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
		return
	}
	for _, i := range ifaces {
		fmt.Println(i.Name)
		fmt.Println(i.HardwareAddr.String())
		fmt.Println(i.Flags.String())
		fmt.Println(net.FlagUp.String())
		fmt.Println("my-test", strings.Contains(i.Flags.String(), net.FlagUp.String()))

		addrs, err := i.Addrs()
		if err != nil {
			fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
			continue
		}
		for _, a := range addrs {
			switch v := a.(type) {
			case *net.IPAddr:
				fmt.Printf("%v : %s (%s)\n", i.Name, v, v.IP.DefaultMask())
			}

		}
	}
}

const (
	EthPrefix                    = "eth"
	CentOSNetworkScriptsTemplate = "/etc/sysconfig/network-scripts/ifcfg-%s"
	NetworkCarSettingTemplate    = `DEVICE=%s
BOOTPROTO=dhcp
ONBOOT=yes
TYPE=Ethernet
`
)

func FindMaxEthIndex() (int, error) {
	var max int
	ifaces, err := net.Interfaces()
	if err != nil {
		return 0, err
	}
	for _, i := range ifaces {
		if strings.HasPrefix(i.Name, EthPrefix) {
			index, _ := strconv.Atoi(strings.TrimPrefix(i.Name, EthPrefix))
			if index > max {
				max = index
			}
		}
	}
	return max, nil
}

func main() {
	localAddresses()

	max, err := FindMaxEthIndex()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("max eth index:", max)
	}

	err = CreateOrUpdateNetworkScripts("eth1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("call CreateOrUpdateNetworkScripts success")

	filePath := fmt.Sprintf(CentOSNetworkScriptsTemplate, "eth2")
	err = os.Remove(filePath)
	if err != nil {
		fmt.Println(err)
	}

}

func CreateOrUpdateNetworkScripts(ethName string) error {
	filePath := fmt.Sprintf(CentOSNetworkScriptsTemplate, ethName)
	content := fmt.Sprintf(NetworkCarSettingTemplate, ethName)
	return os.WriteFile(filePath, []byte(content), 0644)
}
