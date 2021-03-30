package misc

import (
	"log"

	"github.com/melbahja/goph"
)

const (
	HOST   = "rvo"
	USER   = "root"
	PASSWD = "soft4net"
)

func main() {

	client, err := goph.New(USER, HOST, goph.Password(PASSWD))
	if err != nil {
		log.Println("ssh connection failed")
		log.Fatal(err)
	}
	defer client.Close()
	log.Println("Create logs archive... ")
	log.Println("Get journalctl ...")
	_, err = client.Run("journalctl -xb > /var/log/bitbucket.org/storingio/storingio/journalctl.log")
	if err != nil {
		log.Println("Get journalctl info failed")
		log.Fatal(err)
	}
	_, err = client.Run("tar -C /var/log/bitbucket.org/storingio --overwrite -cvf /tmp/logs.zip storingio && chmod 777 /tmp/logs.zip")
	if err != nil {
		log.Println("Create logs archive failed")
		log.Fatal(err)
	}

	err = client.Download("/tmp/logs.zip", "logs.zip")
	if err != nil {
		log.Println("Download logs error")
		log.Fatal(err)
	}
	log.Println("Done, check logs.zip folder")
}
