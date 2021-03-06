package cache

import (
	"fmt"
	"os"

	"encoding/gob"

	"io/ioutil"
	"strings"

	"github.com/mitchellh/go-homedir"
)

var cache string

func Init() error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	cache = home + "/.crazyserver-cache"
	err = os.MkdirAll(cache, 0777)
	if err != nil {
		return err
	}
	return nil
}

func Clear() {
	Init()
	files, _ := ioutil.ReadDir(cache)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".paramcache") || strings.HasSuffix(f.Name(), ".logcache") {
			os.Remove(cache + "/" + f.Name())
		}
	}
}

func LoadParam(crc uint32, e interface{}) error {
	file, err := os.Open(cache + "/" + fmt.Sprintf("%X", crc) + ".paramcache")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(e)
	if err != nil {
		return err
	}
	return nil
}

func SaveParam(crc uint32, e interface{}) error {
	file, err := os.OpenFile(cache+"/"+fmt.Sprintf("%X", crc)+".paramcache", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(e)
	if err != nil {
		return err
	}
	return nil
}

func LoadLog(crc uint32, e interface{}) error {
	file, err := os.Open(cache + "/" + fmt.Sprintf("%X", crc) + ".logcache")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(e)
	if err != nil {
		return err
	}
	return nil
}

func SaveLog(crc uint32, e interface{}) error {
	file, err := os.OpenFile(cache+"/"+fmt.Sprintf("%X", crc)+".logcache", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(e)
	if err != nil {
		return err
	}
	return nil
}
