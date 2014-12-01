package sitemappathes

import (
//	"fmt"
	"log/syslog"
	"math/rand"
	"time"
)

func CreatePathes(golog syslog.Writer, keywords []string) map[string]struct{}{

	qkeywords := len(keywords)

	pathesmap :=make(map[string]struct{})

	rand.Seed(time.Now().Unix())

	var path string

	for i := 0; i < qkeywords; i++ {

		rand01 := rand.Intn(2)

		if rand01 == 0 {

			path = keywords[rand.Intn(qkeywords)] + ".html"

		} else {

			path1 := keywords[rand.Intn(qkeywords)]
			path2 := keywords[rand.Intn(qkeywords)]
			path =path1+"/"+path2+".html"
			
		}
		pathesmap[path]=struct{}{}
				
	}

	return pathesmap

}
