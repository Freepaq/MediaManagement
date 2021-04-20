# MediaManagement

[![Build Status](https://travis-ci.org/Freepaq/MediaManagement.svg?branch=main)](https://travis-ci.org/Freepaq/MediaManagement)
[![Go Report Card](https://goreportcard.com/badge/github.com/Freepaq/MediaManagment)](https://goreportcard.com/report/github.com/Freepaq/MediaManagment)


Useful method to manage Videos and Photo

func Contains(actions []string, key string) bool

func Copy(ori *FileStruct, destFoler string, force bool) bool

func Delete(file FileStruct)

func GetListOfFile(folder string, eligibleFiles string) []string

func IsMediEligible(ext string) bool

func IsPhotoEligible(ext string) bool

func IsVideoEligible(ext string) bool

func ReadPhotoMeta(fname string, fileStr *FileStruct)

func ReadVideoMeta(fname string, fileStr *FileStruct) error

func Rename(file *FileStruct)

type FileStruct struct{ ... }

func GetMeta(fname string) (FileStruct, error)

type MyMapping map[string]interface{}
