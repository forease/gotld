gotld
=====

Get domain's tld by go


## Install gotld

    go get github.com/jonsen/gotld

## Import gotld

    import "github.com/jonsen/gotld"


## Use gotld

For example.

    tld, domain, err := gotld.GetTld( *url )
    if err != nil {
        fmt.Println( err )
    }
    fmt.Printf( "TLD: %s, Domain: %s\n", tld.Tld, domain )

## About

    北京实易时代科技有限公司
    Beijing ForEase Times Technology Co., Ltd.
    http://www.forease.net
    
## Contact me

    Jonsen Yang
    im16hot#gmail.com (replace # with @)
