package main

import (
	"os"
	"time"

	"github.com/AspieSoft/go-regex-re2"
	"github.com/AspieSoft/goutil/fs"
)

func main(){
	fixFile := true

	go func(){
		for {
			time.Sleep(1 * time.Second)
	
			if !fixFile {
				continue
			}
	
			time.Sleep(1 * time.Second)
	
			fixFile = false
	
			if file, err := os.ReadFile("/usr/share/applications/nemo.desktop"); err == nil {
				file = regex.Comp(`(?m)^OnlyShowIn=`).RepStr(file, []byte(`#OnlyShowIn=`))
				err = os.WriteFile("/usr/share/applications/nemo.desktop", file, 0)
				tries := 1000
				for err != nil && tries > 0 {
					tries--
					time.Sleep(100 * time.Millisecond)
					err = os.WriteFile("/usr/share/applications/nemo.desktop", file, 0)
				}
			}
		}
	}()

	watcher := fs.Watcher()
	watcher.OnFileChange = func(path, op string) {
		if path == "/usr/share/applications/nemo.desktop" {
			fixFile = true
		}
	}
	watcher.WatchDir("/usr/share/applications")
	watcher.Wait()
}
