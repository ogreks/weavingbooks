package system

import (
    "log"
    "os"
    "reflect"
    "strconv"
)

type FilesStorage struct {
    name string
    path string
    Storage
}

func (filesStorage FilesStorage) Writer(content interface{}) bool {
    var fileString string
    file, err := OpenFile(filesStorage.path, os.O_APPEND, 0666)
    if err != nil {
        log.Println(err)
        return false
    }
    defer file.Close()
    switch reflect.TypeOf(content).Kind() {
    case reflect.Int, reflect.Int8, reflect.Int16,
        reflect.Int32, reflect.Int64:
        fileString = strconv.FormatInt(reflect.ValueOf(content).Int(), 10)
        break
    case reflect.Uint, reflect.Uint8, reflect.Uint16,
        reflect.Uint32, reflect.Uint64, reflect.Uintptr:
        fileString = strconv.FormatUint(reflect.ValueOf(content).Uint(), 10)
        break
    case reflect.Bool:
        fileString = strconv.FormatBool(reflect.ValueOf(content).Bool())
        break
    case reflect.String:
        fileString = strconv.Quote(reflect.ValueOf(content).String())
        break
    case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
        fileString = reflect.ValueOf(content).Type().String()
    default:
        panic("未定义当前传入类型")
    }
    _, err = file.WriteString(fileString)
    if err != nil {
        log.Println(err)
        return false
    }
    return true
}

func OpenFile(path string, flag int, perm os.FileMode) (*os.File, error)  {
    file, err  := os.OpenFile(path, flag, perm)
    if err != nil {
        log.Fatalln(err)
        return nil, err
    }
    return file, nil
}

func ExistsFile(path string) bool {
    _, err := os.Stat(path)    //os.Stat获取文件信息
    if err != nil {
        if os.IsExist(err) {
            return true
        }
        return false
    }
    return true
}

func IsDir(path string) bool {
    s, err := os.Stat(path)
    if err != nil {
        return false
    }
    return s.IsDir()
}

func IsFile(path string) bool {
    return !IsDir(path)
}