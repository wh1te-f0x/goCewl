package runner

import (
	"flag"
	
)

type Options struct {
	Timeout		int //Time to wait for a response from the server
	Method		string // HTTP Method to use GET/POST
	Target		string // Specify the target
	UserAgent	string // Specify the User Agent to use 
}

func ParseOptions() *Options{

	options := &Options{}
	
	flag.IntVar(&options.Timeout, "t", 30, "Timeout value (Default: 30)")
	flag.StringVar(&options.Method, "m", "GET", "Method to use (Default: GET)")
	flag.StringVar(&options.Target, "u", "", "Url to run against")
	flag.StringVar(&options.UserAgent, "ua", "Mozilla/5.0 (X11; Linux x86_64; rv:76.0) Gecko/20100101 Firefox/76.0", "User Agent to use (Default: Firefox)")

	return options
}

