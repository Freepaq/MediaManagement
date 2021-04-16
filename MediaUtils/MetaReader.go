package MediaUtils

import (
	"github.com/rwcarlsen/goexif/exif"
	"log"
	"os"
)

func ReadVideoMeta(fname string, fileStr *FileStruct) {
//	fmt.Println(os.Getwd())
//	cmd := exec.Command("bin/mediainfo/MediaInfo.exe", " --Output=JSON ", fname)

//	var arr []string
//	json.Unmarshal(cmd.Output(), &arr)

	readFromFile(fname, fileStr)
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
