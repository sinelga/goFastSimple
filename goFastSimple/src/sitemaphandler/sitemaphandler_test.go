package sitemaphandler

import (
    "testing"
    "log/syslog"
    "log"
)



func TestCreate(t *testing.T) {
	
	
	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}
	
	locale :="fi_FI"
	themes :="porno"
	
	startparameters := []string{"tcp",":6379"}

	Create(*golog,locale ,themes,"www.test.com",startparameters)

}

