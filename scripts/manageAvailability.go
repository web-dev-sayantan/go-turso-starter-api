package scripts

import (
	"log"
	"sync"
	"time"

	"github.com/ishanz23/go-turso-starter-api/db"
)

type RateTariff struct {
	ID     string
	Tariff int
}

func GenerateAvailability(startDate, endDate time.Time) error {
	log.Println("Generating availability table data")

	rates := [15]RateTariff{
		{ID: "tph-std-dbl", Tariff: 1199},
		{ID: "tph-std-dbl-xtra", Tariff: 1399},
		{ID: "tph-eco-dbl", Tariff: 999},
		{ID: "tph-eco-dbl-xtra", Tariff: 1199},
		{ID: "tph-duplex", Tariff: 1599},
		{ID: "tph-premium-dbl-xtra", Tariff: 1699},
		{ID: "tph-duplex-trpl", Tariff: 1799},
		{ID: "tph-duplex-quad", Tariff: 1999},
		{ID: "tph-premium-dbl", Tariff: 1499},
		{ID: "tph-dorm", Tariff: 349},
		{ID: "tph-dorm-nr", Tariff: 299},
		{ID: "tph-std-dbl-nr", Tariff: 999},
		{ID: "tph-premium-dbl-nr", Tariff: 1399},
		{ID: "tph-eco-dbl-nr", Tariff: 799},
		{ID: "tph-duplex-quad-nr", Tariff: 1799},
	}

	const workerCount = 16
	var wg sync.WaitGroup
	errChan := make(chan error, workerCount)
	jobChan := make(chan struct {
		date time.Time
		rate RateTariff
	}, workerCount)

	// Start worker pool
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(&wg, jobChan, errChan)
	}

	// Generate jobs
	go func() {
		for d := startDate; d.Before(endDate); d = d.AddDate(0, 0, 1) {
			for _, rate := range rates {
				jobChan <- struct {
					date time.Time
					rate RateTariff
				}{d, rate}
			}
		}
		close(jobChan)
	}()

	// Wait for all workers to finish
	wg.Wait()
	close(errChan)

	// Check for any errors
	for err := range errChan {
		if err != nil {
			log.Printf("Error generating availability: %v\n", err)
			return err
		}
	}

	return nil
}

func worker(wg *sync.WaitGroup, jobs <-chan struct {
	date time.Time
	rate RateTariff
}, errChan chan<- error) {
	defer wg.Done()
	for job := range jobs {
		_, err := db.DB.Exec(
			"INSERT INTO availability(rateId, stayDate, rate) VALUES (?, ?, ?)",
			job.rate.ID,
			job.date.Unix(),
			job.rate.Tariff,
		)
		if err != nil {
			errChan <- err
			return
		}
		log.Printf("Generated availability for %s, rate: %s, tariff: %d\n", job.date.Format("2006-01-02"), job.rate.ID, job.rate.Tariff)
	}
}

func CleanAvailability() error {
	currentDate := time.Now().Unix()

	_, err := db.DB.Exec("DELETE FROM availability WHERE stayDate < ?", currentDate)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
