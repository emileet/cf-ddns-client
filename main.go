package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/cloudflare/cloudflare-go"
	_ "github.com/joho/godotenv/autoload"
)

type recordsJSON struct {
	Record []record `json:"records"`
}

type record struct {
	Name string `json:"name"`
	Zone string `json:"zone"`
}

func main() {
	useIPv6, _ := strconv.ParseBool(os.Getenv("IPV6"))
	currentIP := ""

	for {
		newIP, err := getExternalIP(useIPv6)
		if err != nil {
			log.Fatal(err)
		} else if newIP == currentIP {
			continue
		}
		currentIP = newIP

		records, err := readRecords()
		if err != nil {
			log.Fatal(err)
		}

		api, err := cloudflare.NewWithAPIToken(os.Getenv("API_TOKEN"))
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("updating dns records to '%v'\n", newIP)
		for i := range records.Record {
			recordType := "A"
			if useIPv6 {
				recordType = "AAAA"
			}

			err = updateDNS(api, newIP, recordType, records.Record[i].Zone, records.Record[i].Name)
			if err != nil {
				log.Fatal(err)
			}
		}
		log.Print("dns records updated!\n\n")

		time.Sleep(time.Minute)
	}
}

func updateDNS(api *cloudflare.API, newIP string, recordType string, zone string, name string) error {
	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		return err
	}

	records, err := api.DNSRecords(context.Background(), zoneID, cloudflare.DNSRecord{Name: name, Type: recordType})
	if err != nil {
		return err
	} else if len(records) != 1 {
		return errors.New("invalid number of DNS records retrieved")
	}
	records[0].Content = newIP

	err = api.UpdateDNSRecord(context.Background(), zoneID, records[0].ID, records[0])
	if err != nil {
		return err
	}

	log.Printf("updated %v\n", name)

	return nil
}

func readRecords() (recordsJSON, error) {
	var records recordsJSON

	data, err := ioutil.ReadFile("data/records.json")
	if err != nil {
		return records, err
	}

	err = json.Unmarshal(data, &records)
	if err != nil {
		return records, err
	}

	return records, nil
}

func getExternalIP(useIPv6 bool) (string, error) {
	service := "https://v4.ident.me"
	if useIPv6 {
		service = "https://v6.ident.me"
	}

	response, err := http.Get(service)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	currentIP, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(currentIP), nil
}
