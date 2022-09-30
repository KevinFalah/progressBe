package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	dto "waysgallery/dto/result"
)

func UploadFile(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Upload file
		// FormFile returns the first file for the given key `myFile`
		// it also returns the FileHeader so we can get the Filename,
		// the Header and the size of the file
		file, _, err := r.FormFile("photo")

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode("Error Retrieving the File")
			return
		}
		defer file.Close()

		// setup file type filtering
		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		filetype := http.DetectContentType(buff)
		fmt.Println(filetype)
		if filetype != "image/jpeg" && filetype != "image/png" && filetype != "image/jpg" && filetype != "video/mp4" && filetype != "application/pdf" {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "The provided file format is not allowed. Please upload a JPEG or PNG image"}
			json.NewEncoder(w).Encode(response)
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		// fmt.Printf("File Size: %+v\n", handler.Size)
		// fmt.Printf("MIME Header: %+v\n", handler.Header)
		const MAX_UPLOAD_SIZE = 100 << 20 // 1MB
		// Parse our multipart form, 10 << 20 specifies a maximum
		// upload of 1 MB files.
		r.ParseMultipartForm(MAX_UPLOAD_SIZE)
		if r.ContentLength > MAX_UPLOAD_SIZE {
			w.WriteHeader(http.StatusBadRequest)
			response := Result{Code: http.StatusBadRequest, Message: "Max size in 1mb"}
			json.NewEncoder(w).Encode(response)
			return
		}

		// Create a temporary file within our temp-images directory that follows
		// a particular naming pattern
		// image/png
		// video/mp4
		fileTypeSplit := strings.Split(filetype, "/")
		tempFile, err := ioutil.TempFile("uploads", fileTypeSplit[0]+"-*."+fileTypeSplit[1])
		if err != nil {
			fmt.Println(err)
			fmt.Println("path upload error")
			json.NewEncoder(w).Encode(err)
			return
		}
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		// write this byte array to our temporary file
		tempFile.Write(fileBytes)

		data := tempFile.Name()
		fmt.Println(data)    // uploads/image-89312783912.jpg
		filename := data[8:] // split uploads/

		// add filename to ctx
		ctx := context.WithValue(r.Context(), "dataFile", filename)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
