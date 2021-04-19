# MediaManagement

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