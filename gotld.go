/**
 *
 *
 */
package gotld

import (
	"errors"
	"strings"
)

// tld item
type TldItem struct {
	Id                     int32
	Country, Tld, Category string
	Lables                 int
}

const (
	GOTLD_VERSION = "gotld V1.2"
)

var tldMap = make(map[string]TldItem)

// Initialization Top Level Domain Table
func init() {
	initTld()
}

//
//
func GetTld(url string) (tld TldItem, domain string, err error) {

	dm := strings.Split(url, ".")

	size := len(dm)
	if size == 1 {
		tld, _ = tldMap[url]
	}

	for i := 1; i < size; i++ {
		value, ok := tldMap[strings.Join(dm[size-i:size], ".")]
		if ok {
			tld = value
			domain = strings.Join(dm[size-i-1:size], ".")
		}
	}

	if tld.Tld == "" {
		err = errors.New("Can't get tld from " + url)
	} else {
		tld.Lables = size
	}

	return tld, domain, err
}

func GetVersion() string {
	return GOTLD_VERSION
}
