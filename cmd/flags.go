/**
  Copyright (c) 2022 Zander Schwid & Co. LLC. All rights reserved.
*/

package main

import (
	"flag"
	proxy "github.com/antihosting/tcp-proxy"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

type ForwardPortFlags []proxy.ForwardPort

var (
	Ports  ForwardPortFlags
	ListenIP = flag.String("ip", "0.0.0.0", "Listen/forward ip address, example '0.0.0.0' or '127.0.0.1'")

	ReadTimeout = flag.String("srt", "30s", "Socket read timeout")
	WriteTimeout = flag.String("swt", "30s", "Socket write timeout")

	BenchmarkTest  = flag.String("b", "", "Run benchmark test [http, socket]")
	BenchmarkSize  = flag.Int("bs", 1 << 20, "Batch size")
	Count = flag.Int("count", 1024, "Count of tests")

	Verbose    = flag.Bool("v", false, "Print logs and debug information")
	Foreground = flag.Bool("f", false, "Indicator that proxy is running in foreground")
	LogFile    = flag.String("log", "stdout", "Write log to file")
)

func init() {
	flag.CommandLine.Var(&Ports, "p", "Forward ports in format src:dst repeatable")
}

func (f *ForwardPortFlags) String() string {
	return "Forward ports in format src:dst repeatable"
}

func (f *ForwardPortFlags) Set(value string) error {
	i := strings.IndexByte(value, ':')
	if i == -1 {
		return errors.Errorf("separator ':' not found in '%s'", value)
	}
	src, err := strconv.ParseInt(value[:i], 10, 32)
	if err != nil {
		return errors.Errorf("parsing of first part '%s' of '%s' was failed with error %v", value[:i], value, err)
	}
	dst, err := strconv.ParseInt(value[i+1:], 10, 32)
	if err != nil {
		return errors.Errorf("parsing of second part '%s' of '%s' was failed with error %v", value[i+1:], value, err)
	}
	*f = append(*f, proxy.ForwardPort{ SrcPort: int(src), DstPort: int(dst) } )
	return nil
}


