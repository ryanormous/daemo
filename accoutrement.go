package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

/* COMMAND LINE OPTIONS */
func Getopt() string {
	// CONFIGURATION FILE
	var confpath string
	usage := "OPTIONAL. PATH TO CONFIGURATION FILE."
	flag.StringVar(&confpath, "conf", "", usage)
	// HELP
	var help bool
	msg := "NEVER BE DAUNTED."
	flag.BoolVar(&help, "help", false, msg)
	// PARSE
	flag.Parse()
	// USAGE
	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}
	// CONFIGURATION FILE PATH
	if len(confpath) > 0 {
		confpath, _ = filepath.Abs(confpath)
		if _, err := os.Stat(confpath); os.IsNotExist(err) {
			log.Fatalln("PROBLEM READING FILE:", confpath, err)
		}
	}
	return confpath
}

/* CONFIGURATION */
type Configuration struct {
	Confpath string
	Name     string
	Hostname string
	Logfile  string
	Pidfile  string
	Root     string
	Version  string
}

func (c Configuration) HasField(field string) bool {
	field = strings.Title(field)
	if _, found := reflect.TypeOf(c).FieldByName(field); found {
		return true
	}
	return false
}

func (c *Configuration) SetField(field string, value string) {
	field = strings.Title(field)
	reflect.ValueOf(c).Elem().FieldByName(field).SetString(value)
}

func (c *Configuration) Serialize() []uint8 {
	js, _ := json.Marshal(c)
	return js
}

func (c *Configuration) Print() {
	// SHOW CURRENT CONFIGURATION
	js, _ := json.Marshal(
		map[string]interface{}{"config": c},
	)
	Log.Println(string(js))
}

func (c *Configuration) SetDefaults() {
	// Name
	if c.Name == "" {
		c.Name = "daemo"
	}
	// Hostname
	c.Hostname, _ = os.Hostname()
	// Version
	if c.Version == "" {
		c.Version = "none"
	}
	// Pidfile
	if c.Root != "" && c.Pidfile != "" {
		c.Pidfile = filepath.Join(
			c.Root, "run", c.Pidfile,
		)
	} else {
		c.Pidfile = strings.Join(
			[]string{"/tmp/", c.Name, ".pid"}, "",
		)
	}
	// Logfile
	if c.Root != "" && c.Logfile != "" {
		c.Logfile = filepath.Join(
			c.Root, "log", c.Logfile,
		)
	}
}

func (c *Configuration) Read(opts *map[string]string) {
	// READ CONFIG JSON
	buf, err := ioutil.ReadFile(c.Confpath)
	if err != nil {
		log.Fatalln("PROBLEM READING CONFIG FILE:", c.Confpath, err)
	}
	// PARSE
	err = json.Unmarshal(buf, &opts)
	if err != nil {
		log.Fatalln("bad json:", c.Confpath, err)
	}
}

func (c *Configuration) Load() {
	if len(c.Confpath) > 0 {
		var opts map[string]string
		Conf.Read(&opts)
		for key, val := range opts {
			if Conf.HasField(key) && len(val) > 0 {
				Conf.SetField(key, val)
			}
		}
	}
	Conf.SetDefaults()
}

func ConfInit() {
	// GLOBAL `Conf`
	Conf = &Configuration{}
	Conf.Confpath = Getopt()
	Conf.Load()
}

/* LOGGING */
func LogInit() {
	flag := log.Ldate | log.Ltime | log.Lmicroseconds
	// GLOBAL `Log`
	if Conf.Logfile == "" {
		Log = log.New(os.Stdout, "", flag)
	} else {
		file, err := os.Create(Conf.Logfile)
		if err != nil {
			log.Fatalln(err)
		}
		Log = log.New(file, "", flag)
	}
}
