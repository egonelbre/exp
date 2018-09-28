// +build ignore

package main

import (
	"encoding/base64"
	"log"

	"github.com/gobs/pretty"
	"raintree/smime"
)

func indefiniteToDefinite(data []byte) []byte {

}

func main() {
	data, err := base64.StdEncoding.DecodeString(envelopedData)
	if err != nil {
		log.Fatal(err)
	}

	val, err := smime.Parse(data)
	if err != nil {
		log.Fatal(err)
	}

	pretty.PrettyPrint(val)
}

const envelopedData = `
MIIBHgYJKoZIhvcNAQcDoIIBDzCCAQsCAQAxgcAwgb0CAQAwJjASMRAwDgYDVQQDEwdDYXJ
sUlNBAhBGNGvHgABWvBHTbi7NXXHQMA0GCSqGSIb3DQEBAQUABIGAC3EN5nGIiJi2lsGPcP
2iJ97a4e8kbKQz36zg6Z2i0yx6zYC4mZ7mX7FBs3IWg+f6KgCLx3M1eCbWx8+MDFbbpXadC
DgO8/nUkUNYeNxJtuzubGgzoyEd8Ch4H/dd9gdzTd+taTEgS0ipdSJuNnkVY4/M652jKKHR
LFf02hosdR8wQwYJKoZIhvcNAQcBMBQGCCqGSIb3DQMHBAgtaMXpRwZRNYAgDsiSf8Z9P43
LrY4OxUk660cu1lXeCSFOSOpOJ7FuVyU=`
