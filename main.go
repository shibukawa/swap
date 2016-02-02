package main
/*
#include <mach/mach.h>
#include <mach/vm_page_size.h>

mach_port_t myHost;
vm_statistics64_data_t vm_stat;

void Init()
{
    myHost = mach_host_self();
}

int64_t SwapCount()
{
	unsigned int count = HOST_VM_INFO64_COUNT;
	kern_return_t ret;
	if ((ret = host_statistics64(myHost, HOST_VM_INFO64, (host_info64_t)&vm_stat, &count) != KERN_SUCCESS)) {
        return -1;
	}
    return (int64_t)(vm_stat.swapins + vm_stat.swapouts);
}
 */
import "C"

import (
    "fmt"
    "time"
)

func main() {
    tick := time.Tick(time.Duration(25) * time.Millisecond)
    C.Init()
    lastSwap := C.SwapCount()
    for _ = range tick {
        swap := C.SwapCount()
        if swap != lastSwap {
            fmt.Println(swap)
            lastSwap = swap
        }
    }
}
