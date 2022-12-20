package main

import (
	"encoding/json"
	"log"
	"os"
)

type Job struct {
	Id   int64
	Name string
}

type Result struct {
	Id     int64
	Status string
}

func main() {
	file, _ := os.ReadFile("day5/jobs.json")

	var jobs []Job

	json.Unmarshal(file, &jobs)

	var results []Result
	resultChannel := make(chan Result)
	for _, job := range jobs {
		go execute(job, resultChannel)
		res, _ := <-resultChannel
		results = append(results, res)
	}
	close(resultChannel)

	marshal, _ := json.Marshal(results)
	os.WriteFile("day5/results.json", marshal, 0644)
}

func execute(job Job, result chan<- Result) {
	log.Println("executing ", job.Name)
	result <- Result{Id: job.Id, Status: "SUCCESS"}
}
