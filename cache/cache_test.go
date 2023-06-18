package cache_test

import (
	"os"
	"testing"

	"github.com/jcelliott/lumber"

	"github.com/mu-box/shaman/cache"
	"github.com/mu-box/shaman/config"
	shaman "github.com/mu-box/shaman/core/common"
)

var (
	micropack = shaman.Resource{Domain: "microbox.cloud.", Records: []shaman.Record{{Address: "127.0.0.1"}}}
	microbox  = shaman.Resource{Domain: "microbox.cloud.", Records: []shaman.Record{{Address: "127.0.0.2"}}}
	microBoth = []shaman.Resource{micropack, microbox}
)

func TestMain(m *testing.M) {
	// manually configure
	// config.Log = lumber.NewConsoleLogger(lumber.LvlInt("trace"))
	config.Log = lumber.NewConsoleLogger(lumber.LvlInt("FATAL"))

	// run tests
	rtn := m.Run()

	os.Exit(rtn)
}

// test nil cache init
func TestNoneInitialize(t *testing.T) {
	config.L2Connect = "none://"
	err := cache.Initialize()
	if err != nil {
		t.Errorf("Failed to initalize none cacher - %v", err)
	}
}

// test nil cache addRecord
func TestNoneAddRecord(t *testing.T) {
	noneReset()
	err := cache.AddRecord(&shaman.Resource{})
	if err != nil {
		t.Errorf("Failed to add record to none cacher - %v", err)
	}
}

// test nil cache getRecord
func TestNoneGetRecord(t *testing.T) {
	noneReset()
	_, err := cache.GetRecord("microbox.cloud")
	if err != nil {
		t.Errorf("Failed to get record from none cacher - %v", err)
	}
}

// test nil cache updateRecord
func TestNoneUpdateRecord(t *testing.T) {
	noneReset()
	err := cache.UpdateRecord("microbox.cloud", &shaman.Resource{})
	if err != nil {
		t.Errorf("Failed to update record in none cacher - %v", err)
	}
}

// test nil cache deleteRecord
func TestNoneDeleteRecord(t *testing.T) {
	noneReset()
	err := cache.DeleteRecord("microbox.cloud")
	if err != nil {
		t.Errorf("Failed to delete record from none cacher - %v", err)
	}
}

// test nil cache resetRecords
func TestNoneResetRecords(t *testing.T) {
	noneReset()
	err := cache.ResetRecords(&[]shaman.Resource{})
	if err != nil {
		t.Errorf("Failed to reset records in none cacher - %v", err)
	}
}

// test nil cache listRecords
func TestNoneListRecords(t *testing.T) {
	noneReset()
	_, err := cache.ListRecords()
	if err != nil {
		t.Errorf("Failed to list records in none cacher - %v", err)
	}
}

func TestNoneExists(t *testing.T) {
	noneReset()
	if cache.Exists() {
		t.Error("Cache exits but shouldn't")
	}
}

func noneReset() {
	config.L2Connect = "none://"
	cache.Initialize()
}
