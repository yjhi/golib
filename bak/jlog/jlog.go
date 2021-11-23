package jlog

/*******************************************************************
*
* Changed By yjh 20210407
*
*
********************************************************************/

import (
	"log"
	"os"
	"time"
)

type Log struct {
	SaveFile   bool
	FilePath   string
	FileHandle *os.File
}

func New(isfile bool, filepre string) *Log {

	j := &Log{
		SaveFile: isfile,
		FilePath: filepre + time.Now().Format("2006-01-02") + ".log",
	}

	return j
}

func NewLog(isfile bool, filepre string) *Log {
	return New(isfile, filepre)

}

func NewLogOpen(isfile bool, filepre string) (*Log, bool) {

	j := &Log{
		SaveFile: isfile,
		FilePath: filepre + time.Now().Format("2006-01-02") + ".log",
	}

	return j, j.init()
}

func (c *Log) Open() bool {
	return c.init()
}

func (c *Log) init() bool {

	FileHandle, err := os.OpenFile(c.FilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
		return false
	}

	if c.SaveFile {
		log.SetOutput(FileHandle)
	}

	return true

}

func (c *Log) Write(data string, flag string) {
	log.Println("["+flag+"]", data)
}

func (c *Log) Debug(data string) {
	c.Write(data, "DEBUG")
}
func (c *Log) Info(data string) {
	c.Write(data, "INFO")
}
func (c *Log) Error(data string) {
	c.Write(data, "ERROR")
}
func (c *Log) Line() {
	log.Println("-----------------------------------------------------------------")
}

func (c *Log) Close() {
	if c.FileHandle != nil {
		c.FileHandle.Close()
	}
}
