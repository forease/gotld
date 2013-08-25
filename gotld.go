/**
 *
 * 
 */
package gotld


import (
    //"fmt"
    "strings"
)


// tld item
type TldItem struct {
    Id int32
    Country, Tld, Category string
}

const (
    GOTLD_VERSION = "gotld V1.0"
)

var tldMap = make(map[string]*TldItem)

// Initialization Top Level Domain Table
func init() {
    initTld()
}

// 
// 
func GetTld( url string ) ( tld *TldItem, domain string, err error ) {
    var (
        tmpTld, tar string
        djump uint
    )

    dm := strings.Split( url, "." )

    for i := len(dm) - 1; i >= 0; i-- {

        tmpTld = dm[i] + tar + tmpTld
        tar = "."

        // 判断当前域名是否为域名
        currTld, ok := tldMap[ tmpTld ]
        if ok {
            tld = currTld
            if i - 1 >= 0 {
                domain = dm[i-1] + "." + tmpTld
            }
        }

        djump++

        if djump > 3 {
            break
        }
    }


    return tld, domain, err
}


func GetVersion() string {
    return GOTLD_VERSION
}
