package agent

import (
	"fmt"
	"log"
	"net"
	"net/url"
)

// ReallyUUID is a common agent/collector helper to generate a UUID for bosh-lite
// that all share the same damned "unique id".
func ReallyUUID(boshTarget string, boshUUID string) string {
	uri, err := url.Parse(boshTarget)
	if err != nil {
		log.Fatalln(err)
	}
	host := boshTarget
	if uri.Host != "" {
		host, _, err = net.SplitHostPort(uri.Host)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return fmt.Sprintf("%s-%s", host, boshUUID)
}
