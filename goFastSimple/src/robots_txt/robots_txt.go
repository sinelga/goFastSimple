package robots_txt

import (
	"bytes"
	"log/syslog"
)

func Create(golog syslog.Writer, locale string, themes string, site string) []byte {

	var buffer bytes.Buffer
	buffer.WriteString("User-agent: *\nAllow: /\nSitemap: http://" + site + "/sitemap.xml\n")

	return buffer.Bytes()
}
