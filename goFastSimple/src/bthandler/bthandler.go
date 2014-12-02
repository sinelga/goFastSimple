package bthandler

import (
	"clean_pathinfo"
//	"createfirstgz"
	"createpage"
	"log/syslog"
	"net/http"
	"strings"
	"sitemaphandler"
	"robots_txt"
)

func BTrequestHandler(golog syslog.Writer, resp http.ResponseWriter, req *http.Request, locale string, themes string, site string, pathinfo string, bot string, startparameters []string, blocksite bool,variant string) {

	pathinfoclean := clean_pathinfo.CleanPath(golog, pathinfo)
	
//	golog.Info(site+" "+bot)

//	if blocksite {
//
//		golog.Info("BTrequestHandler:will block-> " + site+pathinfo)
//
//	}

	var bytepage []byte
	if strings.HasSuffix(pathinfoclean, ".html")  {

		bytepage = createpage.CreateHtmlPage(golog, locale, themes,site, bot, startparameters,blocksite,variant)

		resp.Write(bytepage)

	} else if strings.HasSuffix(pathinfoclean, "sitemap.xml")  {
		
		
		bytepage =sitemaphandler.Create(golog, locale, themes,site,startparameters)
		resp.Header().Add("Content-type","text/xml")
		resp.Write(bytepage)
		
	} else if strings.HasSuffix(pathinfoclean, "robots.txt") {
		
		
		bytepage =robots_txt.Create(golog, locale, themes,site)
		resp.Header().Add("Content-type","text/plain")
		resp.Write(bytepage)
		
		
	} else {
		
		resp.WriteHeader(404)
		
	}

//	if strings.HasSuffix(pathinfoclean, ".html") {
//
//		go createfirstgz.Creategzhtml(golog, locale, themes, site, pathinfoclean, bytepage)
//
//	}

}
