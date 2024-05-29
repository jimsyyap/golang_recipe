package main

import (
	"github.com/jandre/procfs"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"strconv"
)

func main() {
	processes, _ := procfs.Processes(false)
	table := tablewriter.NewWriter(os.Stdout)

	for _, p := range processes {
		table.Append([]string{strconv.Itoa(p.Pid), p.Exe, p.Cwd})
	}
	table.Render()

	m, _ := procfs.NewMeminfo()
	log.Println(m)

}

/*
* this code retrieves a list of running processes on the system
* it creates a table in the console displaying each process's ID, executable path, and ucrrent workind dir
* fetches and logs system memory information
 */
