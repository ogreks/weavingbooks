package system

import (
    "bufio"
    "io"
    "log"
    "strings"
)

type Config interface {
	SetKey(key, value string)
	GetKey(key string) interface{}
}

type BaseConfig struct {

}

func OpenConfig(path string) map[string]string {
    config := make(map[string]string)
    file, err := OpenFile(path)
    if err != nil {
        log.Fatalln(err)
        panic(err)
    }
    defer file.Close()
    r := bufio.NewReader(file)
    for {
        line, _, err := r.ReadLine()
        if err != nil {
            if err == io.EOF {
                break
            }
            panic(err)
        }
        content := strings.TrimSpace(string(line))
        index := strings.Index(content,"=")
        if index < 0 {
            continue
        }
        key := strings.TrimSpace(content[:index])
        if len(key) == 0 {
            continue
        }
        value := strings.TrimSpace(content[index+1:])
        if len(value) == 0 {
            continue
        }
        config[key] = value
    }
    return config
}