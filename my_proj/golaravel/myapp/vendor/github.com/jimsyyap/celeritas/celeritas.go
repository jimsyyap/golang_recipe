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

    err = c.checkDotEnv(rootPath)
    if err != nil {
        return err
    }

    // read .env
    err = godotenv.Load(rootPath + "/.env")
    if err != nil {
        return err
    }

    // create loggers
    infoLog, errorLog := c.startLoggers()
    c.InfoLog = infoLog
    c.ErrorLog = errorLog
    c.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
    c.Version = version

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

func (c *Celeritas) checkDotEnv(path string) error {
    err := c.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))
    if err != nil {
        return err
    }
    return nil
}

func (c *Celeritas) startLoggers() (*log.Logger, *log.Logger) {
    var infoLog *log.Logger
    var errorLog *log.Logger

    infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    return infoLog, ErrorLog
}
