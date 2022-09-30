package middleware

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

func UploadMultipleFile(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := r.ParseMultipartForm(200000) // grab the multipart form
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		formdata := r.MultipartForm // ok, no problem so far, read the Form data

		//get the *fileheaders
		files := formdata.File["multiplefiles"] // grab the filenames

		for i, _ := range files { //loop the files
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}

			out, err := os.Create("/tmp/" + files[i].Filename)

			defer out.Close()
			if err != nil {
				fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
				return
			}

			_, err = io.Copy(out, file) //file not files[i]

			if err != nil {
				fmt.Fprintln(w, err)
				return
			}

			fmt.Fprintf(w, "Files uploaded successfully : ")
			fmt.Fprintf(w, files[i].Filename+"\n")
		
		}

		ctx := context.WithValue(r.Context(), "multipleImage", files)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
