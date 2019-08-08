/**
 *
 *
 */
package gotld

import (
	"bytes"
	"errors"
	"strings"
	"sync"
)

// TldItem tld item
type TldItem struct {
	Id                     int32
	Country, Tld, Category string
	Lables                 int
}

const (
	GOTLD_VERSION = "gotld V1.3"
)

var (
	tldMap = make(map[string]TldItem)
	pool   = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}
)

// Initialization Top Level Domain Table
func init() {
	initTld()
}

// GetTld 获取域名
func GetTld(url string) (tld TldItem, domain string, err error) {
	tld, domain, _, err = subDomain(url, 0)

	return
}

// GetSubdomain 获取子域名
func GetSubdomain(url string, level int) (subdomain, domain, tld string) {
	var t TldItem
	t, domain, subdomain, _ = subDomain(url, level)

	return subdomain, domain, t.Tld
}

func subDomain(url string, level int) (tld TldItem, domain, subDomain string, err error) {

	var (
		buffer = pool.Get().(*bytes.Buffer)
		isTLD  bool
	)

	dm := strings.Split(url, ".")

	size := len(dm)
	if size > 1 {
		idx := 0

		for i := size - 1; i >= 0; i-- {
			// 组合域名
			for j := i; j < size; j++ {
				buffer.WriteString(dm[j])
				if j != size-1 {
					buffer.WriteString(".")
				}
			}
			subDomain = buffer.String()

			// 重置buffer
			buffer.Reset()

			// 判断是否为TLD
			value, ok := tldMap[subDomain]
			if ok {
				tld = value
				isTLD = true
				continue
			}

			// 找出TLD后
			if isTLD {
				if domain == "" {
					domain = subDomain
				}
				if idx >= level {
					break
				}
				idx++
			}
		}
	} else {
		tld, _ = tldMap[url]
	}

	pool.Put(buffer)

	if tld.Tld == "" {
		err = errors.New("Can't get tld from " + url)
	} else {
		tld.Lables = size
	}

	return
}

// GetVersion -
func GetVersion() string {
	return GOTLD_VERSION
}
