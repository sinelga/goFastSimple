package sitemaphandler

import (
	"domains"
//	"fmt"
	"log/syslog"
	"encoding/xml"
//	"os"
	"time"
	"findkeywords"
	"sitemappathes"
	"bytes"
)

func Create(golog syslog.Writer, locale string,themes string,site string,startparameters []string) []byte  {
	
	keywordsarr := findkeywords.GetAll(golog,locale,themes,startparameters)
	
	pathesarr :=sitemappathes.CreatePathes(golog,keywordsarr)

	docList := new(domains.Pages)
	docList.XmlNS ="http://www.sitemaps.org/schemas/sitemap/0.9"

	for path := range pathesarr {

	doc := new(domains.Page)

	doc.Loc = "http://"+site+"/"+path
	doc.Lastmod = time.Now().Local().Format(time.RFC3339)
//	doc.Name = "The Example Times"
//	doc.Language = "en"
//	doc.Title = "Companies A, B in Merger Talks"
//	doc.Keywords = "business, merger, acquisition, A, B"
//	doc.Image = "http://www.google.com/spacer.gif"
	docList.Pages = append(docList.Pages, doc)
	
	}
	
	resultXml, err := xml.MarshalIndent(docList, "", "  ")
	if err != nil {
		
		golog.Crit(err.Error())
	}
	
	var buffer bytes.Buffer
	
	buffer.WriteString(xml.Header)
	buffer.Write(resultXml)
	
//	fmt.Println(buffer.String())

	return buffer.Bytes()

}
