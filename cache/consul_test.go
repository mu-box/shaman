package cache_test

import (
	"testing"

	"github.com/mu-box/shaman/cache"
	"github.com/mu-box/shaman/config"
	shaman "github.com/mu-box/shaman/core/common"
)

// test consul cache init
func TestConsulInitialize(t *testing.T) {
	config.L2Connect = "consul://127.0.0.1:8500"
	err := cache.Initialize()
	cache.Initialize()
	if err != nil {
		t.Errorf("Failed to initalize consul cacher - %v", err)
	}
}

// test consul cache addRecord
func TestConsulAddRecord(t *testing.T) {
	consulReset()
	err := cache.AddRecord(&micropack)
	if err != nil {
		t.Errorf("Failed to add record to consul cacher - %v", err)
	}
}

// test consul cache getRecord
func TestConsulGetRecord(t *testing.T) {
	consulReset()
	cache.AddRecord(&micropack)
	_, err := cache.GetRecord("microbox.cloud")
	_, err2 := cache.GetRecord("microbox.cloud")
	if err == nil || err2 != nil {
		t.Errorf("Failed to get record from consul cacher - %v%v", err, err2)
	}
}

// test consul cache updateRecord
func TestConsulUpdateRecord(t *testing.T) {
	consulReset()
	err := cache.UpdateRecord("microbox.cloud", &micropack)
	err2 := cache.UpdateRecord("microbox.cloud", &micropack)
	if err != nil || err2 != nil {
		t.Errorf("Failed to update record in consul cacher - %v%v", err, err2)
	}
}

// test consul cache deleteRecord
func TestConsulDeleteRecord(t *testing.T) {
	consulReset()
	err := cache.DeleteRecord("microbox.cloud")
	cache.AddRecord(&micropack)
	err2 := cache.DeleteRecord("microbox.cloud")
	if err != nil || err2 != nil {
		t.Errorf("Failed to delete record from consul cacher - %v%v", err, err2)
	}
}

// test consul cache resetRecords
func TestConsulResetRecords(t *testing.T) {
	consulReset()
	err := cache.ResetRecords(&microBoth)
	if err != nil {
		t.Errorf("Failed to reset records in consul cacher - %v", err)
	}
}

// test consul cache listRecords
func TestConsulListRecords(t *testing.T) {
	consulReset()
	_, err := cache.ListRecords()
	cache.ResetRecords(&microBoth)
	_, err2 := cache.ListRecords()
	if err != nil || err2 != nil {
		t.Errorf("Failed to list records in consul cacher - %v%v", err, err2)
	}
}

func consulReset() {
	config.L2Connect = "consul://127.0.0.1:8500"
	cache.Initialize()
	blank := make([]shaman.Resource, 0, 0)
	cache.ResetRecords(&blank)
}
