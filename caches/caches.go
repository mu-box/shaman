package caches

// TODO:
//  - implement caching backends
//  - add logging
//  - test

import (
	"fmt"
	"net/url"
	// "github.com/nanopack/shaman/config"
)

type Cacher interface {
	GetRecord(string) (string, error)
	SetRecord(string, string) error
	ReviseRecord(string, string) error
	DeleteRecord(string) error
}

var (
	l1 Cacher
	l2 Cacher
)

func initializeCacher(connection string, expires int) (Cacher, error) {
	u, err := url.Parse(connection)
	if err != nil {

	}
	switch u.Scheme {
	case "redis":
		cacher, err := NewRedisCacher(connection, expires)
		if err != nil {
			return nil, err
		}
		return cacher, nil
	case "postgres":
		cacher, err := NewPostgresCacher(connection, expires)
		if err != nil {
			return nil, err
		}
		return cacher, nil
	}
	return nil, nil
}

func Init() {

}

func Key(domain string, rtype uint16) string {
	return fmt.Sprintf("%d-%s", rtype, domain)
}

func AddRecord(key string, value string) error {
	if l2 != nil {
		err := l2.SetRecord(key, value)
		if err != nil {
			return nil
		}
	}
	if l1 != nil {
		err := l1.SetRecord(key, value)
		if err != nil {
			return nil
		}
	}
	return nil
}

func RemoveRecord(key string) error {
	if l2 != nil {
		err := l2.DeleteRecord(key)
		if err != nil {
			return err
		}
	}
	if l1 != nil {
		err := l1.DeleteRecord(key)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateRecord(key string, value string) error {
	if l2 != nil {
		err := l2.ReviseRecord(key, value)
		if err != nil {
			return err
		}
	}
	if l1 != nil {
		err := l1.ReviseRecord(key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func FindRecord(key string) (string, error) {
	var record string
	if l1 != nil {
		record, err := l1.GetRecord(key)
		if err != nil {
			return record, err
		}
	}
	if record != "" {
		return record, nil
	}
	if l2 != nil {
		record, err := l2.GetRecord(key)
		if record != "" {
			l1.SetRecord(key, record)
			return record, err
		}
	}
	return "", nil
}