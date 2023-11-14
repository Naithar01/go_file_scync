package global

import (
	"log"
	"net"
	"os"
)

func SetRootPath(root_path string) error {
	err := os.Setenv("RootPath", root_path)
	if err != nil {
		return err
	}

	return nil
}

func GetRootPath() string {
	root_path_string := os.Getenv("RootPath")
	return root_path_string
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
