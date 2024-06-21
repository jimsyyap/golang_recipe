package celeritas

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"
    "time"

    "github.com/go-chi/chi/v5"
    "github.com/joho/godotenv"
)
const version = "1.0.0"

type Celeritas struct {
    AppName string
    Debug bool
    Version string
    ErrorLog *log.Logger
    InfoLog *log.Logger
    RootPath string
    Routes *chi.Mux
    config config
}

func (c *Celeritas) New(rootPath string) error {
    pathCofig := initPaths{
        rootPath: rootPath,
        folderNames: []string {
            "handlers",
            "migrations",
            "views",
            "data",
            "public",
            "tmp",
            "logs",
            "middleware",
        },
    }

    err := c.Init(pathCofig)
    if err != nil {
        return err
    }

    return nil
}

func (c *Celeritas) Init(p initPaths) error {
    root := p.rootPath
    for _, path := range p.folderNames {
        //create folder if doessnt exist
        err := c.CreateDirIfNotExist(root + "/" + path)
        if err != nil {
            return err
        }
    }
    return nil
}
