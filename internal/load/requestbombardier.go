package load

import (
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/charmbracelet/log"
	"github.com/zzucch/http-load-tester/internal/config"
)

var (
	count     int = 1
	StartTime time.Time
	mutex     sync.Mutex
)

func RequestBombardier(wg *sync.WaitGroup, bombardierId int) {
	defer wg.Done()

	for {
		resp, err := http.Get(config.Config.RequestURL)
		if err != nil {
			log.Fatal("failed to http get", "bombardier", bombardierId, "err", err)
		}

		mutex.Lock()
		count++
		mutex.Unlock()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("failed to read", "bombardier", bombardierId, "err", err)
		} else {
			log.Print("result", "bombardier", bombardierId, "body", string(body))
		}

		mutex.Lock()
		log.Info("current stats:",
			"requests count", count,
			"time since start", time.Since(StartTime),
			"rps", float64(count)/time.Since(StartTime).Seconds())

		mutex.Unlock()
	}
}
