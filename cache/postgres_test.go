package cache_test

import (
	"testing"

	"github.com/mu-box/shaman/cache"
	"github.com/mu-box/shaman/config"
	shaman "github.com/mu-box/shaman/core/common"
)

// test postgres cache init
func TestPostgresInitialize(t *testing.T) {
	config.L2Connect = "postgres://postgres@127.0.0.1?sslmode=disable" // default
	err := cache.Initialize()
	config.L2Connect = "postgresql://postgres@127.0.0.1:9999?sslmode=disable" // unable to init?
	err2 := cache.Initialize()
	if err != nil || err2 != nil {
		t.Errorf("Failed to initalize postgres cacher - %v%v", err, err2)
	}
}

// test postgres cache addRecord
func TestPostgresAddRecord(t *testing.T) {
	postgresReset()
	err := cache.AddRecord(&micropack)
	if err != nil {
		t.Errorf("Failed to add record to postgres cacher - %v", err)
	}

	err = cache.AddRecord(&micropack)
	if err != nil {
		t.Errorf("Failed to add record to postgres cacher - %v", err)
	}
}

// test postgres cache getRecord
func TestPostgresGetRecord(t *testing.T) {
	postgresReset()
	cache.AddRecord(&micropack)
	_, err := cache.GetRecord("microbox.cloud.")
	_, err2 := cache.GetRecord("microbox.cloud")
	if err == nil || err2 != nil {
		t.Errorf("Failed to get record from postgres cacher - %v%v", err, err2)
	}
}

// test postgres cache updateRecord
func TestPostgresUpdateRecord(t *testing.T) {
	postgresReset()
	err := cache.UpdateRecord("microbox.cloud", &micropack)
	err2 := cache.UpdateRecord("microbox.cloud", &micropack)
	if err != nil || err2 != nil {
		t.Errorf("Failed to update record in postgres cacher - %v%v", err, err2)
	}
}

// test postgres cache deleteRecord
func TestPostgresDeleteRecord(t *testing.T) {
	postgresReset()
	err := cache.DeleteRecord("microbox.cloud")
	cache.AddRecord(&micropack)
	err2 := cache.DeleteRecord("microbox.cloud")
	if err != nil || err2 != nil {
		t.Errorf("Failed to delete record from postgres cacher - %v%v", err, err2)
	}
}

// test postgres cache resetRecords
func TestPostgresResetRecords(t *testing.T) {
	postgresReset()
	err := cache.ResetRecords(&microBoth)
	if err != nil {
		t.Errorf("Failed to reset records in postgres cacher - %v", err)
	}
}

// test postgres cache listRecords
func TestPostgresListRecords(t *testing.T) {
	postgresReset()
	_, err := cache.ListRecords()
	cache.ResetRecords(&microBoth)
	_, err2 := cache.ListRecords()
	if err != nil || err2 != nil {
		t.Errorf("Failed to list records in postgres cacher - %v%v", err, err2)
	}
}

func postgresReset() {
	config.L2Connect = "postgres://postgres@127.0.0.1?sslmode=disable"
	cache.Initialize()
	blank := make([]shaman.Resource, 0, 0)
	cache.ResetRecords(&blank)
}
