package MediaUtils

import (
	"github.com/Freepaq/MediaManagement/MediaUtils/Setup"
	"github.com/rwcarlsen/goexif/exif"
	"log"
	"os"
)

func ReadVideoMeta(fname string, fileStr *Setup.FileStruct) {
	readFromFile(fname, fileStr)
}

func ReadPhotoMeta(fname string, fileStr *Setup.FileStruct) {
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
		fileStr.MetaOrigin = Setup.METAORIGINMETA
	}
}

func readFromFile(fname string, fileStr *Setup.FileStruct) {
	fileStat, err := os.Stat(fname)
	if err != nil {
		log.Fatal(err)
	}
	fileStr.CreationDate = fileStat.ModTime()
	fileStr.MetaOrigin = Setup.METAORIGINFILE
}
