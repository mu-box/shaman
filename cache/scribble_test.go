package cache_test

import (
	"os"
	"testing"

	"github.com/mu-box/shaman/cache"
	"github.com/mu-box/shaman/config"
)

// test scribble cache init
func TestScribbleInitialize(t *testing.T) {
	config.L2Connect = "/tmp/shamanCache" // default
	err := cache.Initialize()
	config.L2Connect = "!@#$%^&*()" // unparse-able
	err2 := cache.Initialize()
	config.L2Connect = "scribble:///roots/file" // unable to init? (test no sudo)
	err3 := cache.Initialize()
	config.L2Connect = "scribble:///" // defaulting to "/var/db"
	cache.Initialize()
	if err != nil || err2 == nil || err3 != nil {
		t.Errorf("Failed to initalize scribble cacher - %v%v%v", err, err2, err3)
	}
}

// test scribble cache addRecord
func TestScribbleAddRecord(t *testing.T) {
	scribbleReset()
	err := cache.AddRecord(&micropack)
	if err != nil {
		t.Errorf("Failed to add record to scribble cacher - %v", err)
	}
}

// test scribble cache getRecord
func TestScribbleGetRecord(t *testing.T) {
	scribbleReset()
	cache.AddRecord(&micropack)
	_, err := cache.GetRecord("microbox.cloud")
	_, err2 := cache.GetRecord("microbox.cloud")
	if err == nil || err2 != nil {
		t.Errorf("Failed to get record from scribble cacher - %v%v", err, err2)
	}
}

// test scribble cache updateRecord
func TestScribbleUpdateRecord(t *testing.T) {
	scribbleReset()
	err := cache.UpdateRecord("microbox.cloud", &micropack)
	err2 := cache.UpdateRecord("microbox.cloud", &micropack)
	if err != nil || err2 != nil {
		t.Errorf("Failed to update record in scribble cacher - %v%v", err, err2)
	}
}

// test scribble cache deleteRecord
func TestScribbleDeleteRecord(t *testing.T) {
	scribbleReset()
	err := cache.DeleteRecord("microbox.cloud")
	cache.AddRecord(&micropack)
	err2 := cache.DeleteRecord("microbox.cloud")
	if err != nil || err2 != nil {
		t.Errorf("Failed to delete record from scribble cacher - %v%v", err, err2)
	}
}

// test scribble cache resetRecords
func TestScribbleResetRecords(t *testing.T) {
	scribbleReset()
	err := cache.ResetRecords(&microBoth)
	if err != nil {
		t.Errorf("Failed to reset records in scribble cacher - %v", err)
	}
}

// test scribble cache listRecords
func TestScribbleListRecords(t *testing.T) {
	scribbleReset()
	_, err := cache.ListRecords()
	cache.ResetRecords(&microBoth)
	_, err2 := cache.ListRecords()
	if err != nil || err2 != nil {
		t.Errorf("Failed to list records in scribble cacher - %v%v", err, err2)
	}
}

func scribbleReset() {
	os.RemoveAll("/tmp/shamanCache")
	config.L2Connect = "scribble:///tmp/shamanCache"
	cache.Initialize()
}
