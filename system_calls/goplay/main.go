package main

import "os"        // for interacting with the operating system
import "flag"      //parses command line arguments
import "fmt"       //formatted IO
import "strings"   //string manipulation
import "debug/elf" //working with ELF files
import "syscall"   //make syscalls, for debugging

// this function prints dynamic strings from the ELF file, including DT_NEEDED, DT_SONAME, DT_RPATH, and DT_RUNPATH
func dump_dynstr(file *elf.File) {
	fmt.Printf("DynStrings:\n")
	dynstrs, _ := file.DynString(elf.DT_NEEDED)
	for _, e := range dynstrs {
		fmt.Printf("\t%s\n", e)
	}
	dynstrs, _ = file.DynString(elf.DT_SONAME)
	for _, e := range dynstrs {
		fmt.Printf("\t%s\n", e)
	}
	dynstrs, _ = file.DynString(elf.DT_RPATH)
	for _, e := range dynstrs {
		fmt.Printf("\t%s\n", e)
	}
	dynstrs, _ = file.DynString(elf.DT_RUNPATH)
	for _, e := range dynstrs {
		fmt.Printf("\t%s\n", e)
	}
}

// this function prints symbols from the ELF file, filtering out symbols with empty names
func dump_symbols(file *elf.File) {
	fmt.Printf("Symbols:\n")
	symbols, _ := file.Symbols()
	for _, e := range symbols {
		if !strings.EqualFold(e.Name, "") {
			fmt.Printf("\t%s\n", e.Name)
		}
	}
}

// this function opens the ELF file and calls dump_dynstr and dump_symbols to print dynamic strings and symbols
func dump_elf(filename string) int {
	file, err := elf.Open(filename)
	if err != nil {
		fmt.Printf("cant open file : \"%s\" as an ELF.\n", filename)
		return 2
	}
	dump_dynstr(file)
	dump_symbols(file)
	return 0
}

// this function starts a new process with ptrace enabled and ataches to it for debugging
func init_debug(filename string) int {
	attr := &os.ProcAttr{Sys: &syscall.SysProcAttr{Ptrace: true}}
	if proc, err := os.StartProcess(filename, []string{"/"}, attr); err == nil {
		proc.Wait()
		foo := syscall.PtraceAttach(proc.Pid)
		fmt.Printf("Started New Process: %v.\n", proc.Pid)
		fmt.Printf("PtraceAttach res: %v.\n", foo)
		return 0
	}
	return 2
}

// command-line flags: the program checks for two flags: filename for the ELF file and action for the action (dump or debug)
// file validation - the program checks if the specified file exists and is not a directory
// action execution - depending on the action, the program calls for init_debug or dump_elf
// usage message - if the arguments are not valid, it prints a usage message and exits with error code
func main() {
	if len(os.Args) > 1 {
		filename := flag.String("filename", "", "A binary ELF file.")
		action := flag.String("action", "", "Action to make: {dump|debug}.")
		flag.Parse()
		if *filename == "" || *action == "" {
			goto Usage
		}
		file, err := os.Stat(*filename)
		if os.IsNotExist(err) {
			fmt.Printf("No such file or directory: %s.\n", *filename)
			goto Error
		} else if mode := file.Mode(); mode.IsDir() {
			fmt.Printf("Parameter must be a file, not a " +
				"directory.\n")
			goto Error
		}
		f, err := os.Open(*filename)
		if err != nil {
			fmt.Printf("Couldn’t open file: \"%s\".\n", *filename)
			goto Error
		}
		f.Close()
		fmt.Printf("Tracing program : \"%s\".\n", *filename)
		fmt.Printf("Action : \"%s\".\n", *action)
		switch *action {
		case "debug":
			os.Exit(init_debug(*filename))
		case "dump":
			os.Exit(dump_elf(*filename))
		}
	} else {
		goto Usage
	}

Usage:
	fmt.Printf("Usage of ./main:\n" +
		"  -action=\"{dump|debug}\": Action to make.\n" +
		"  -filename=\"file\": A binary ELF file.\n")
	goto Error

Error:
	os.Exit(2)
}

// Summary
// dump: Prints dynamic strings and symbols from the ELF file.
// debug: Starts a new process with ptrace enabled for debugging.
// The code combines ELF file manipulation and process tracing to provide basic ELF analysis and debugging capabilities.

// output

// Usage of ./main:
//   -action="{dump|debug}": Action to make.
//   -filename="file": A binary ELF file.
// exit status 2
