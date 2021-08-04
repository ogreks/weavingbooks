package system

type StorageApi interface {
    Writer(path string,content interface{}) bool
    Reader(path string) interface{}
}

type Storage struct {
    Names string
}