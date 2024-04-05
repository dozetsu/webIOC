package main

import (
    "fmt"
    "net/http"
    "os"
    "D:\go\"

    "github.com/kardianos/service"
)

type program struct{}

func (p *program) Start(s service.Service) error {
    go p.run()
    return nil
}

func (p *program) Stop(s service.Service) error {
    return nil
}

func (p *program) run() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Serve the text files in the directory
        dir := "./files"
        filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
            if !info.IsDir() {
                http.ServeFile(w, r, path)
            }
            return nil
        })
    })

    fmt.Println("Server started on :8080")
    http.ListenAndServe(":8080", nil)
}

func main() {
    svcConfig := &service.Config{
        Name:        "webIOC",
        DisplayName: "webIOC",
        Description: "A simple web server to host webIOC",
    }

    prg := &program{}
    svc, err := service.New(prg, svcConfig)
    if err != nil {
        fmt.Println(err)
        return
    }

    if len(os.Args) > 1 {
        err := service.Control(svc, os.Args[1])
        if err != nil {
            fmt.Println("Valid actions: install, uninstall, start, stop")
            fmt.Println(err)
            return
        }
        return
    }

    err = svc.Run()
    if err != nil {
        fmt.Println(err)
    }
}
