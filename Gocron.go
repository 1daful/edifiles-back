package cron

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func Start() {
	// Create a new scheduler
	s := gocron.NewScheduler(time.UTC)
	// Connect to the database
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=users sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Get the current date
	now := time.Now()
	// Query the database for the users whose birthdays are today
	rows, err := db.Query("SELECT name FROM users WHERE EXTRACT(month FROM birthday) = $1 AND EXTRACT(day FROM birthday) = $2", now.Month(), now.Day())
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	// Loop over the query results and schedule the sendMessage function for each user
	for rows.Next() {
		var user string
		err := rows.Scan(&user)
		if err != nil {
			panic(err)
		}
		s.Every(1).Day().At("10:00").Do(SendMessage, user)
	}
	// Start the scheduler and block the main goroutine
	s.StartBlocking()
}

func StartSchedule(interval int, timeUnit TimeUnit, startDelay int,  task func(...interface{})) {
	s := gocron.NewScheduler(time.UTC)
	
	switch timeUnit {
		case "second": // for a seocnd interval
			s.Every(interval).Second().At(startDelay).Do(task)
		case "minute": // for a minute interval
			s.Every(interval).Minute().At(startDelay)
		case "hour": // for a hour interval
			s.Every(interval).Hour().At(startDelay)
		case "day": // for a daily interval
			s.Every(interval).Day().At(startDelay)
		case "week": // for a weekly interval
			s.Every(interval).Week().At(startDelay)
		case "month": // for a monthly interval
			s.Every(interval).Month().At(startDelay)
		// Add more cases for different intervals as needed
		default:
			s.Every(1).Day() // Default to daily interval
		}

	//s.Every(interval).timeUnit().At(timeToStart).Do(sendMessage, user)
	s.StartAsync()
}


type TimeUnit string

const (
    TimeUnitSecond    TimeUnit = "second"
    TimeUnitMinute TimeUnit = "minute"
    TimeUnitHour   TimeUnit = "hour"
    TimeUnitDay    TimeUnit = "day"
    TimeUnitWeek TimeUnit = "week"
    TimeUnitMonth   TimeUnit = "month"
)