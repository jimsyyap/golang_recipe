package main

import (
	"fmt"
	syscall "golang.org/x/sys/unix"
)

const (
	gigabyte = (1024.0 * 1024.0 * 1024.0)
)

func main() {
	var statfs = syscall.Statfs_t{}
	var total uint64
	var used uint64
	var free uint64
	err := syscall.Statfs("/", &statfs)
	if err != nil {
		fmt.Printf("Failed to get disk info: %s\n", err)
		return
	} else {
		total = statfs.Blocks * uint64(statfs.Bsize)
		free = statfs.Bfree * uint64(statfs.Bsize)
		used = total - free
	}

	fmt.Printf("total disk space : %.1f GB\n", float64(total)/gigabyte)
	fmt.Printf("used disk space  : %.1f GB\n", float64(used)/gigabyte)
	fmt.Printf("free disk space  : %.1f GB\n", float64(free)/gigabyte)
}

//output
// total disk space : 100.5 GB
// used disk space  : 43.1 GB
// free disk space  : 57.5 GB
