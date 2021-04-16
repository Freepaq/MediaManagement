package MediaUtils

import (
	"encoding/json"
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type MyMapping map[string]interface{}

func ReadVideoMeta(fname string, fileStr *FileStruct) {
	//	fmt.Println(os.Getwd())
	cmd, _ := exec.Command("bin/mediainfo/MediaInfo.exe", "--Output=JSON", fname).Output()

	resultingMap := MyMapping{}
	//	file, _ := ioutil.ReadFile("temp.txt")
	if err := json.Unmarshal(cmd, &resultingMap); err != nil {
		log.Println("json.Compact:", err)
		if serr, ok := err.(*json.SyntaxError); ok {
			log.Println("Occurred at offset:", serr.Offset)
		}
	}
	fmt.Println("Current File :" + fname)
	var encodeDate = search(resultingMap, []string{"Encoded_Date", "File_Created_Date"})
	if "" == encodeDate {
		readFromFile(fname, fileStr)
	} else {
		time, err := time.Parse("2006-01-02T15:04:05", encodeDate)
		if nil != err {
			fmt.Println(err)
		} else {
			fileStr.CreationDate = time
			fileStr.MetaOrigin = METAORIGINMETA
		}
	}
}
func search(str MyMapping, value []string) string {
	var result = ""
	var media MyMapping
	media = str["media"].(map[string]interface{})
	//	fmt.Println(reflect.TypeOf(media))
	if nil != media {
		var track []interface{}
		track = media["track"].([]interface{})
		if nil != track {
			for _, m := range track {
				for key, v := range m.(map[string]interface{}) {
					if Contains(value, key) {

						var date = v.(string)
						if "" != date {
							dateArray := strings.Split(date, " ")
							return dateArray[1] + "T" + dateArray[2]
						}
					}
				}
			}
		}
	}
	return result
}

func ReadPhotoMeta(fname string, fileStr *FileStruct) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	f.Seek(0, 0)
	x, err := exif.Decode(f)
	if err != nil {
		readFromFile(fname, fileStr)
	} else {
		fileStr.CreationDate, _ = x.DateTime()
		fileStr.MetaOrigin = METAORIGINMETA
	}
}

func readFromFile(fname string, fileStr *FileStruct) {
	fileStat, err := os.Stat(fname)
	if err != nil {
		log.Fatal(err)
	}
	fileStr.CreationDate = fileStat.ModTime()
	fileStr.MetaOrigin = METAORIGINFILE
}
