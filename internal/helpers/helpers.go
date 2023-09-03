package helpers

import (
	"fmt"
	"io"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"runtime/debug"
	"time"

	"github.com/mrkouhadi/chinauptodate/internal/config"
	"golang.org/x/crypto/bcrypt"
)

var app *config.AppConfig

// Newhelpers sets up appconfig for helpers

func Newhelpers(a *config.AppConfig) {
	app = a
}

// ClientError sends errors to the client when something goes wrong on the clientside

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client Error With Status Of ", status)
	http.Error(w, http.StatusText(status), status)
}

// ServerError sends errors to the client when something goes wrong on the serverside
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

type ImageDetails struct {
	Filename  string
	Extension string
	Size      int64
	Header    textproto.MIMEHeader
}

// / upload a file fuction
func UploadFile(w http.ResponseWriter, r *http.Request, name string) (ImageDetails, error) {
	// 1. Parse input, type multipart/form-data
	r.Body = http.MaxBytesReader(w, r.Body, 32<<20+512)
	r.ParseMultipartForm(32 << 20) // means max 32 MB    32 << 20 means max 10 MB
	// 2. retrieve file from posted form-data
	file, handler, err := r.FormFile("img")
	if err != nil {
		fmt.Println("Failed to retrieve file from form-data")
		fmt.Println(err)
		return ImageDetails{}, err
	}
	defer file.Close()

	fmt.Printf("Uploaded File Name : %+v \n", handler.Filename)
	fmt.Printf("File Size : %+v KB \n", handler.Size/1000)
	fmt.Printf("MIME Header : %+v \n", handler.Header)
	ext := filepath.Ext(handler.Filename)
	fmt.Printf("Extension of The file : %+v \n", ext)

	// 3. write temprorary file on our server
	tempoFile, err := os.Create("static/images/" + name + ext)

	if err != nil {
		fmt.Println(err)
		return ImageDetails{}, err
	}
	defer tempoFile.Close()
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return ImageDetails{}, err
	}
	tempoFile.Write(fileBytes)
	return ImageDetails{
		Filename:  handler.Filename,
		Extension: ext,
		Size:      handler.Size / 1000, // KB
		Header:    handler.Header,
	}, nil
}

// ////////// upload multiple files
func UploadMultipleFiles(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 32<<20+512)
	r.ParseMultipartForm(32 << 20) // means max 32 MB    32 << 20 means max 10 MB
	files := r.MultipartForm.File["img"]

	for _, handler := range files {
		fmt.Printf("Uploaded File Name : %+v \n", handler.Filename)
		fmt.Printf("File Size : %+v KB \n", handler.Size/1000)
		fmt.Printf("MIME Header : %+v \n", handler.Header)
		ext := filepath.Ext(handler.Filename)
		fmt.Printf("Extension of The file : %+v \n", ext)

		file, err := handler.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		buff := make([]byte, 512)
		_, _ = file.Read(buff)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		filetype := http.DetectContentType(buff)
		if filetype != "image/jpeg" && filetype != "image/png" {
			http.Error(w, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
			return
		}
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		newFile, err := os.Create(fmt.Sprintf("./static/images/%d%s", time.Now().UnixNano(), filepath.Ext(handler.Filename)))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

// /   Hashing Passwords
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// check if user is authenticated
func IsUserAuthenticated(r *http.Request) bool {
	exists := app.Session.Exists(r.Context(), "user")
	return exists
}
