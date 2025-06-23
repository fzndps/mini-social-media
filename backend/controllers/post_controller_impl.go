package controllers

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/fzndps/mini-social-media/backend/constant"
	"github.com/fzndps/mini-social-media/backend/helper"
	"github.com/fzndps/mini-social-media/backend/models/web"
	"github.com/fzndps/mini-social-media/backend/services"
	"github.com/julienschmidt/httprouter"
)

type PostControllerImpl struct {
	PostService services.PostService
}

func NewPostController(postService services.PostService) PostController {
	return &PostControllerImpl{
		PostService: postService,
	}
}

func (controller *PostControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// 1. Parse multipart form dulu
	err := request.ParseMultipartForm(20)
	if err != nil {
		log.Println("❌ Gagal parsing multipart form:", err)
		http.Error(writer, "Gagal parsing form", http.StatusBadRequest)
		return
	}

	// 2. Ambil content dari form
	content := request.FormValue("content")
	log.Println("✅ Konten post:", content)

	// 3. Ambil user_id dari context (middleware JWT)
	claimsRaw := request.Context().Value(constant.UserInfoKey)
	claims, ok := claimsRaw.(*helper.JWTClaim)
	if !ok {
		log.Println("❌ GAGAL: JWT context bukan *JWTClaim")
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userId := claims.Id
	log.Println("✅ User ID dari JWT:", userId)

	// 4. Inisialisasi imageURL kosong
	imageURL := ""

	// 5. Coba ambil file gambar (jika ada)
	file, fileHeader, err := request.FormFile("image")
	if err == nil && file != nil {
		defer file.Close()
		log.Println("✅ File diterima:", fileHeader.Filename)

		// 6. Buat folder "uploads" jika belum ada
		err = os.MkdirAll("uploads", os.ModePerm)
		if err != nil {
			log.Println("❌ Gagal buat folder uploads:", err)
			http.Error(writer, "Gagal buat folder uploads", http.StatusInternalServerError)
			return
		}
		log.Println("✅ Folder uploads siap")

		// 7. Simpan file ke disk
		filePath := "./uploads/" + fileHeader.Filename
		fileDest, err := os.Create(filePath)
		if err != nil {
			log.Println("❌ Gagal buat file tujuan:", err)
			http.Error(writer, "Gagal menyimpan file", http.StatusInternalServerError)
			return
		}
		defer fileDest.Close()

		_, err = io.Copy(fileDest, file)
		if err != nil {
			log.Println("❌ Gagal menyalin file:", err)
			http.Error(writer, "Gagal menyimpan file", http.StatusInternalServerError)
			return
		}
		log.Println("✅ File disimpan ke:", filePath)

		// 8. Simpan URL publik (bukan path lokal) ke DB
		imageURL = "http://localhost:3000/uploads/" + fileHeader.Filename
		log.Println("🌍 Image URL publik:", imageURL)
	} else {
		log.Println("ℹ️ Tidak ada file gambar diunggah")
	}

	// 9. Buat request struct
	postCreateRequest := web.PostCreateRequest{
		UserId:   userId,
		Content:  content,
		ImageURL: imageURL,
	}
	log.Println("📦 Request ke Service:", postCreateRequest)

	// 10. Panggil service
	response := controller.PostService.Create(request.Context(), postCreateRequest)
	log.Println("✅ Response dari Service:", response)

	// 11. Kirim response ke client
	helper.WriteResponseBody(writer, response)
}

func (controller *PostControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	postResponse := controller.PostService.FindAll(request.Context())

	log.Println("Ini isi PostResponse", postResponse)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   postResponse,
	}

	log.Println("Ini isi web response", webResponse)

	helper.WriteResponseBody(writer, webResponse)
}
