package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	//
	c := cron.New(cron.WithSeconds())

	entryID1, err := c.AddFunc("* * * * * *", func() {
		fmt.Println(time.Now())
	})
	if err != nil {
		panic(err)
	}

	example1 := NewExampleJob("example-job")
	// Run example-job every second
	entryID2, err := c.AddJob("* * * * * * ", example1)

	c.Start()

	time.Sleep(60 * time.Second)
	// Remove all cron jobs after 60 seconds
	c.Remove(entryID1)
	c.Remove(entryID2)

	select {}
}

// exampleJob
type exampleJob struct {
	Name string
}

func NewExampleJob(name string) *exampleJob {
	ej := &exampleJob{
		Name: name,
	}
	return ej
}

func (ej *exampleJob) Run() {
	fmt.Println("Name: ", ej.Name, time.Now())
}
