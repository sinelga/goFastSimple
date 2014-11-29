package createpage

import (
	"bytes"
	"domains"
	"findfreeparagraph"
	"fmt"
	"html/template"
	"log/syslog"
	"templ_funcmap"
	"time"
)

func CreateHtmlPage(golog syslog.Writer, locale string, themes string, bot string, startparameters []string, blocksite bool, variant string) []byte {

	var paragrapharr []domains.Paragraph

	var base string
	var page string
	var mediablock string

	base = "/home/juno/git/goFastSimple/goFastSimple/templ/" + locale + "/" + themes + "/" + variant + "/base.html"
	page = "/home/juno/git/goFastSimple/goFastSimple/templ/" + locale + "/" + themes + "/" + variant + "/page.html"
	mediablock = "/home/juno/git/goFastSimple/goFastSimple/templ/" + locale + "/" + themes + "/" + variant + "/mediablock.html"

	funcMap := template.FuncMap{
		"FirstWord":                  templ_funcmap.FirstWord,
		"SecondWordFromTitle":		  templ_funcmap.SecondWordFromTitle,	
		"FirstWordFromSenteces":      templ_funcmap.FirstWordFromSenteces,
		"FirstWordFromAllParagraphs": templ_funcmap.FirstWordFromAllParagraphs,
	}

	index, _ := template.New("base").Funcs(funcMap).ParseFiles(
		base,
		page,
		mediablock,
	)

	paragraph := findfreeparagraph.FindFromQ(golog, locale, themes, bot, startparameters)

	if blocksite {

		paragraph.Plocallink = ""
	}

	paragrapharr = append(paragrapharr, paragraph)

	webpage := bytes.NewBuffer(nil)
	
	currenttime := time.Now().Local()
	
	fmt.Println("Pushsite ",paragraph.Pushsite)
	fmt.Println("Ptittle ",paragraph.Ptitle)
	

	htmlpage := domains.Htmlpage{
		Locale:     locale,
		Themes:     themes,
		Variant:    variant,
		Paragraphs: paragrapharr,
		Created: currenttime.Format("2006-01-02 15:04:05"),
		Updated: currenttime.Format("2006-01-02 15:04:05"),
		Pushsite: paragraph.Pushsite,
	}

	if err := index.Execute(webpage, htmlpage); err != nil {
		golog.Err(err.Error())
	}

	webpagebytes := make([]byte, webpage.Len())
	webpagebytes = webpage.Bytes()

	return webpagebytes

}
