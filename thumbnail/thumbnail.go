package thumbnail

import (
	"log"
	"os"
	"sync"
)

func ImageFile(infile string) (string, error) {
	return infile, nil
}

// makeThumbnails 生成指定文件的缩略图
func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// makeThumbnails2 注意: 不正确
// 在没有完成想要完成的事之前退出
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f)
	}
}

// makeThumbnails3 并行生成指定文件的缩略图
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f) {
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}
	for range filenames {
		<-ch
	}
}

// makeThumbnials4 为指定文件并行地生成缩略图
// 如果任何步骤出错它返回一个错误
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}
	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}
	return nil
}

// makeThumbnails5 为指定文件并行的生成缩略图
// 它以任意顺序返回生成的文件名
// 如果任何步骤出错就返回一个错误
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}
	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles := append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}

// makeThumbnails6 为从通道上接收的每个文件生成缩略图
// 它返回其生成文件所占用的字节数
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	wg := sync.WaitGroup
	for _, f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}
	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
