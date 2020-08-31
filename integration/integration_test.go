package integration

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"testing"
	"time"
)

func TestCommandService_healthy(t *testing.T) {
	res, err := Cfg.CommandService.health()
	check(t, err)
	_, err = ioutil.ReadAll(res.Body)
	check(t, err)
}
func TestConsumerService_healthy(t *testing.T) {
	res, err := Cfg.ConsumerService.health()
	check(t, err)
	_, err = ioutil.ReadAll(res.Body)
	check(t, err)
}
func TestQueryService_healthy(t *testing.T) {
	res, err := Cfg.QueryService.health()
	check(t, err)
	_, err = ioutil.ReadAll(res.Body)
	check(t, err)
}

func TestE2E_recordCreated(t *testing.T) {
	// COMMAND
	id := submitCommand(t)
	log.Printf("command submitted with id: %s", id)

	// QUERY
	record, duration, err := getTimestampRecord(id)
	if err != nil {
		t.Fatalf("timestamp not written to db within SLA, took: %dms", duration.Milliseconds())
	}
	createdAt, ok := record["created_at"]
	if !ok{
		t.Fatalf("didn't find `created_at` in response")
	}
	eventTimestamp, ok := record["event_timestamp"]
	if !ok{
		t.Fatalf("didn't find `event_timestamp` in response")
	}
	id, ok = record["id"]
	if !ok{
		t.Fatalf("didn't find `id` in response")
	}

	log.Printf("command record retrieved in %dms with id %s event_timestamp %s and created_at %s", duration.Milliseconds(), id, eventTimestamp, createdAt)
}

func TestE2E_bunkRecordsCreated(t *testing.T) {
	totalRecordsToCreate := 1000
	start := time.Now()
	var records []string
	// command
	for i := 0 ; i < totalRecordsToCreate ; i++ {
		records = append(records, submitCommand(t))
	}

	// query
	for i, id := range records {
		if i % 100 == 0{
			log.Printf("queried for record %d created of %d", i, totalRecordsToCreate)
		}
		getTimestampRecord(id)
	}
	elapsed := time.Since(start)
	tps := math.Round(float64(totalRecordsToCreate) / elapsed.Seconds())
	log.Printf("%d records created in %dms, average e2e TPS: %v", totalRecordsToCreate, elapsed.Milliseconds(), tps)
}


func getTimestampRecord(id string) (map[string]string, time.Duration, error) {
	start := time.Now()
	client := resty.New()
	client.SetRetryCount(10)
	client.SetRetryWaitTime(250 * time.Millisecond)
	client.AddRetryCondition(
		func(r *resty.Response, err error) bool {
			retry := r.StatusCode() == http.StatusNotFound
			if retry{
				log.Println("attempting retry")
			}
			return retry
		},
	)
	resp, err := client.R().
		SetHeader("correlationId", uuid.Must(uuid.NewUUID()).String()).
		Get(Cfg.QueryService.endpointPath() + "/api/v1/timestamp/" + id)
	if err != nil {
		return nil, time.Since(start), err
	}

	m := make(map[string]string)
	if resp.StatusCode() == http.StatusOK {
		err := json.Unmarshal(resp.Body(), &m)
		if err != nil {
			return nil, time.Since(start), err
		}
		return m, time.Since(start), err
		//check(t, err)
	} else {
		return m, time.Since(start),
			fmt.Errorf("invalid status code %d from endpoint %s",
				resp.StatusCode(), Cfg.CommandService.endpointPath()+"/app")
	}
}

func submitCommand(t *testing.T) (id string) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("correlationId", uuid.Must(uuid.NewUUID()).String()).
		Post(Cfg.CommandService.endpointPath() + "/app")
	check(t, err)

	if resp.StatusCode() == http.StatusOK {
		id = gjson.Get(string(resp.Body()), "id").String()
	} else {
		t.Fatalf("invalid status code %d from endpoint %s", resp.StatusCode(), Cfg.CommandService.endpointPath()+"/app")
	}
	return id
}
