package eureka

import (
	"strings"
	log "github.com/Sirupsen/logrus"
)

func (c *Client) SendHeartbeat(appId, instanceId string) error {
	values := []string{"apps", appId, instanceId}
	path := strings.Join(values, "/")
	res, err := c.Put(path, nil)
	if res != nil {
		log.WithFields(log.Fields{
			"path":path,
			"status":res.StatusCode,
		}).Info("SendHeartbeat ")
	}
	return err
}
