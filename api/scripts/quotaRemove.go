package scripts

import (
	"time"

	"alif/quota/app/models"
)

func RunQuotaRemove() {
	timer := time.NewTicker(time.Second * 30)

	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			models.RemoveOldQuota(time.Now().Add(-time.Minute))

		case <-stopScripts:
			return
		}

	}
}
