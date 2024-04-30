package main

import (
	"sync"
	"time"

	"github.com/zzucch/http-load-tester/internal/config"
	"github.com/zzucch/http-load-tester/internal/load"
)

func main() {
	var wg sync.WaitGroup
	load.StartTime = time.Now()

	for i := 0; i < config.Config.BombardiersAmount; i++ {
		wg.Add(1)
		go load.RequestBombardier(&wg, i)
	}

	wg.Wait()
}
