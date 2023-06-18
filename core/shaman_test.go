package shaman_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/jcelliott/lumber"

	"github.com/mu-box/shaman/config"
	"github.com/mu-box/shaman/core"
	sham "github.com/mu-box/shaman/core/common"
)

var (
	micropack  = sham.Resource{Domain: "microbox.cloud.", Records: []sham.Record{{Address: "127.0.0.1"}}}
	micropack2 = sham.Resource{Domain: "microbox.cloud.", Records: []sham.Record{{Address: "127.0.0.3"}}}
	microbox   = sham.Resource{Domain: "microbox.cloud.", Records: []sham.Record{{Address: "127.0.0.2"}}}
	microBoth  = []sham.Resource{micropack, microbox}
)

func TestMain(m *testing.M) {
	shamanClear()
	// manually configure
	config.Log = lumber.NewConsoleLogger(lumber.LvlInt("FATAL"))

	// run tests
	rtn := m.Run()

	os.Exit(rtn)
}

func TestAddRecord(t *testing.T) {
	shamanClear()
	err := shaman.AddRecord(&micropack)
	err = shaman.AddRecord(&micropack)
	err2 := shaman.AddRecord(&micropack2)
	if err != nil || err2 != nil {
		t.Errorf("Failed to add record - %v%v", err, err2)
	}
}

func TestGetRecord(t *testing.T) {
	shamanClear()
	_, err := shaman.GetRecord("microbox.cloud")
	shaman.AddRecord(&micropack)
	_, err2 := shaman.GetRecord("microbox.cloud")
	if err == nil || err2 != nil {
		// t.Errorf("Failed to get record - %v%v", err, "hi")
		t.Errorf("Failed to get record - %v%v", err, err2)
	}
}

func TestUpdateRecord(t *testing.T) {
	shamanClear()
	err := shaman.UpdateRecord("microbox.cloud", &micropack)
	err2 := shaman.UpdateRecord("microbox.cloud", &micropack)
	if err != nil || err2 != nil {
		t.Errorf("Failed to update record - %v%v", err, err2)
	}
}

func TestDeleteRecord(t *testing.T) {
	shamanClear()
	err := shaman.DeleteRecord("microbox.cloud")
	shaman.AddRecord(&micropack)
	err2 := shaman.DeleteRecord("microbox.cloud")
	if err != nil || err2 != nil {
		t.Errorf("Failed to delete record - %v%v", err, err2)
	}
}

func TestResetRecords(t *testing.T) {
	shamanClear()
	err := shaman.ResetRecords(&microBoth)
	err2 := shaman.ResetRecords(&microBoth, true)
	if err != nil || err2 != nil {
		t.Errorf("Failed to reset records - %v%v", err, err2)
	}
}

func TestListDomains(t *testing.T) {
	shamanClear()
	domains := shaman.ListDomains()
	if fmt.Sprint(domains) != "[]" {
		t.Errorf("Failed to list domains - %+q", domains)
	}
	shaman.ResetRecords(&microBoth)
	domains = shaman.ListDomains()
	if len(domains) != 2 {
		t.Errorf("Failed to list domains - %+q", domains)
	}
}

func TestListRecords(t *testing.T) {
	shamanClear()
	resources := shaman.ListRecords()
	if fmt.Sprint(resources) != "[]" {
		t.Errorf("Failed to list records - %+q", resources)
	}
	shaman.ResetRecords(&microBoth)
	resources = shaman.ListRecords()
	if len(resources) == 2 && (resources[0].Domain != "microbox.cloud." && resources[0].Domain != "microbox.cloud.") {
		t.Errorf("Failed to list records - %+q", resources)
	}
}

func TestExists(t *testing.T) {
	shamanClear()
	if shaman.Exists("microbox.cloud") {
		t.Errorf("Failed to list records")
	}
	shaman.AddRecord(&micropack)
	if !shaman.Exists("microbox.cloud") {
		t.Errorf("Failed to list records")
	}
}

func shamanClear() {
	shaman.Answers = make(map[string]sham.Resource, 0)
}
