package eureka

import (
	"encoding/json"
	"strings"
	log "github.com/Sirupsen/logrus"
)

func (c *Client) RegisterInstance(appName string, instanceInfo *InstanceInfo) error {
	values := []string{"apps", appName}
	path := strings.Join(values, "/")
	instance := &Instance{
		Instance: instanceInfo,
	}
	body, err := json.Marshal(instance)
	if err != nil {
		return err
	}

	res, err := c.Post(path, body)
	if res != nil {
		log.WithFields(log.Fields{
			"path":path,
			"status":res.StatusCode,
			"body":string(body),
		}).Info("RegisterInstance ")
	}
	return err
}
