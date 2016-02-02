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
	"bytes"
	"fmt"
	"golang.org/x/mobile/exp/audio"
	"log"
	"time"
)

type ReadSeekCloser struct {
	*bytes.Reader
}

func (rsk ReadSeekCloser) Close() error {
	return nil
}

func main() {
	tick := time.Tick(time.Duration(16) * time.Millisecond)
	C.Init()
	lastSwap := C.SwapCount()
	noise := ReadSeekCloser{bytes.NewReader(MustAsset("sound/noise.wav"))}
	player, err := audio.NewPlayer(noise, audio.Mono16, 44100)
	if err != nil {
		log.Fatal(err)
	}
	player.SetVolume(1.0)
	for _ = range tick {
		swap := C.SwapCount()
		if swap != lastSwap {
			fmt.Println("swap count:", swap)
			state := player.State()
			if state == audio.Playing || state == audio.Paused {
				player.Stop()
			}
			player.Seek(0)
			player.Play()
			lastSwap = swap
		}
	}
}
