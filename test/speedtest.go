package main

import (
    "context"
    "errors"
    "fmt"
    "github.com/jimsyyap/speedtest/transport"
    "gopkg.in/alecthomas/kingpin.v2"
    "io"
    "log"
    "os"
    "strconv"
    "strings"
    "sync/atomic"
    "time"
    "github.com/jimsyyap/speedtest"
)

var (
    showList = kingpin.Flag("list", "Show available servers").Short('l').Bool()
    serverIds = kingpin.Flag("server", "Select server id to speed test").Short('s').Ints()
    customURL = kingpin.Flag("url", "Custom speed test server url").String()
    savingMode = kingpin.Flag("saving-mode", "Test with few resources, though low accuracy (especially > 30Mbps)").Bool()
)
