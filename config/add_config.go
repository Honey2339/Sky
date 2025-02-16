package config

import (
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

func interval() {
	time.Sleep(1 * time.Second)
}

func Add_config() (RootConfigSchema, error) {
    config := RootConfigSchema{}
    
    var PORT string
    var err error
    defaultWorker := runtime.NumCPU()
    defaultWorkerStr := strconv.Itoa(defaultWorker)

    huh.NewInput().
        Title("Port").
        Prompt("?").
        Value(&PORT).
        Run()
    
    config.Server.Listen, err = strconv.Atoi(PORT)

    _ = spinner.New().
        Title("Config...").
        Action(interval).
        Run()

    huh.NewInput().
        Title("How many worker do you need").
        Prompt("?").
        Value(&defaultWorkerStr).
        Run()

    config.Server.Workers, err = strconv.Atoi(defaultWorkerStr)

    _ = spinner.New().
        Title("Initializing Workers...").
        Action(interval).
        Run()
    
    for i := 0; i < config.Server.Workers; i++ {
        var name string
        var url string

        huh.NewInput().
            Title(fmt.Sprintf("Worker %d name", (i + 1))).
            Prompt("?").
            Value(&name).
            Run()

        huh.NewInput().
            Title("Url of your upstream").
            Prompt("?").
            Value(&url).
            Run()

        config.Server.Upstreams = append(config.Server.Upstreams, Upstream{
            ID: name,
            URL: url,
        })

        _ = spinner.New().
            Title("Configuring...").
            Action(interval).
            Run()
    }

    //Rules
    // var flag bool = false
    // huh.NewConfirm().
    //     Title("Do you config rules?").
    //     Affirmative("Yes!").
    //     Negative("No!").
    //     Value(&flag).
    //     Run()
    
    // if(flag){
        
    // }

    return config, err
}