package scripts

import (
	"time"

	"alif/quota/app/models"
)

func RunQuotaRemove() {
	timer := time.NewTicker(time.Minute * 5)

	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			models.RemoveOldQuota(time.Now().Add(-time.Hour))

		case <-stopScripts:
			return
		}

	}
}
