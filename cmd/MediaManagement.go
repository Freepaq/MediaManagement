package main

import (
	"MediaManagement/pkg/MediaUtils"
	"MediaManagement/pkg/Setup"
	"fmt"
	"os"
	"strings"
)

var actions []string

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Please check arguments")
		os.Exit(-1)
	}
	fmt.Println("Excution Time :" + Setup.CurrentTime)
	mediaType := os.Args[1]
	action := os.Args[2]
	origin := os.Args[3]
	dest := os.Args[4]
	actions = strings.Split(action, ".")
	rows := MediaUtils.GetListOfFile(origin, mediaType)
	fmt.Println(Setup.SEPARATOR)
	for _, file := range rows {

		meta, err := MediaUtils.GetMeta(file)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Source file : [" + file + "]")
			fmt.Println("Creation : [" + meta.CreationDate.String() + "] taken from " + meta.MetaOrigin)

		}
		if Setup.Contains(actions, "RENAME") {
			MediaUtils.Rename(&meta)
		}
		if Setup.Contains(actions, "COPY") {
			MediaUtils.Copy(&meta, dest, false)
		}
		if Setup.Contains(actions, "REPLACE") {
			MediaUtils.Copy(&meta, dest, true)
		}
		if Setup.Contains(actions, "MOVE") {
			if result := MediaUtils.Copy(&meta, dest, true); result {
				MediaUtils.Delete(meta)
			}
		}
		fmt.Println(Setup.SEPARATOR)
	}

	fmt.Println("End Time :" + Setup.CurrentTime)
}
