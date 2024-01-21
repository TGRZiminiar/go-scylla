package fileHelper

import (
	"context"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

type (
	FileReq struct {
		File        *multipart.FileHeader `form:"file"`
		Destination string                `form:"destination"`
		Extension   string
		FileName    string
	}
	FileRes struct {
		FileName string `json:"filename"`
		Url      string `json:"url"`
	}

	DeleteFileReq struct {
		Destination string `json:"destination"`
	}

	filesPub struct {
		bucket      string
		destination string
		file        *FileRes
	}
)

func UploadToStorage(req []*FileReq) ([]*FileRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	jobsCh := make(chan *FileReq, len(req))
	resultsCh := make(chan *FileRes, len(req))
	errsCh := make(chan error, len(req))

	res := make([]*FileRes, 0)

	for _, r := range req {
		jobsCh <- r
	}
	close(jobsCh)

	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		go uploadToStorageWorker(ctx, jobsCh, resultsCh, errsCh)
	}

	for a := 0; a < len(req); a++ {
		err := <-errsCh
		if err != nil {
			return nil, err
		}

		result := <-resultsCh
		res = append(res, result)
	}
	return res, nil
}

func uploadToStorageWorker(ctx context.Context, jobs <-chan *FileReq, results chan<- *FileRes, errs chan<- error) {
	for job := range jobs {
		cotainer, err := job.File.Open()
		if err != nil {
			errs <- err
			return
		}
		b, err := ioutil.ReadAll(cotainer)
		if err != nil {
			errs <- err
			return
		}

		// Upload an object to storage
		dest := fmt.Sprintf("./assets/images/%s", job.Destination)
		if err := os.WriteFile(dest, b, 0777); err != nil {
			if err := os.MkdirAll("./assets/images/"+strings.Replace(job.Destination, job.FileName, "", 1), 0777); err != nil {
				errs <- fmt.Errorf("mkdir \"./assets/images/%s\" failed: %v", err, job.Destination)
				return
			}
			if err := os.WriteFile(dest, b, 0777); err != nil {
				errs <- fmt.Errorf("write file failed: %v", err)
				return
			}
		}

		newFile := &filesPub{
			file: &FileRes{
				FileName: job.FileName,
				// host and post
				Url: fmt.Sprintf("http://%s:%d/%s", "mix", "5000", job.Destination),
			},
			destination: job.Destination,
		}

		errs <- nil
		results <- newFile.file
	}
}
