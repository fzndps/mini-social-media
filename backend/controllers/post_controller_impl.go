package controllers

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/fzndps/mini-social-media/backend/helper"
	"github.com/fzndps/mini-social-media/backend/middleware"
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
	// 1. Ambil file dari form
	file, fileHeader, err := request.FormFile("image")
	if err != nil {
		log.Println("❌ Gagal ambil file dari form:", err)
		http.Error(writer, "Gagal ambil file image", http.StatusBadRequest)
		return
	}
	defer file.Close()
	log.Println("✅ File diterima:", fileHeader.Filename)

	request.ParseMultipartForm(20)

	// 2. Buat folder "uploads" jika belum ada
	err = os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		log.Println("❌ Gagal buat folder uploads:", err)
		http.Error(writer, "Gagal buat folder uploads", http.StatusInternalServerError)
		return
	}
	log.Println("✅ Folder uploads siap")

	// 3. Simpan file ke folder "uploads"
	filePath := "./uploads/" + fileHeader.Filename
	fileDestination, err := os.Create(filePath)
	if err != nil {
		log.Println("❌ Gagal buat file tujuan:", err)
		http.Error(writer, "Gagal menyimpan file", http.StatusInternalServerError)
		return
	}
	defer fileDestination.Close()

	_, err = io.Copy(fileDestination, file)
	if err != nil {
		log.Println("❌ Gagal menyalin file:", err)
		http.Error(writer, "Gagal menyimpan file", http.StatusInternalServerError)
		return
	}
	log.Println("✅ File disimpan ke:", filePath)

	// 4. Ambil content dari form
	content := request.PostFormValue("content")
	log.Println("✅ Konten post:", content)

	// 5. Ambil user_id dari context (middleware JWT)
	claimsRaw := request.Context().Value(middleware.UserInfoKey)
	claims, ok := claimsRaw.(*helper.JWTClaim)
	if !ok {
		log.Println("❌ GAGAL: JWT context bukan *JWTClaim")
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userId := claims.Id
	log.Println("✅ User ID dari JWT:", userId)

	// 6. Buat request struct
	postCreateRequest := web.PostCreateRequest{
		UserId:   userId,
		Content:  content,
		ImageURL: filePath,
	}
	log.Println("📦 Request ke Service:", postCreateRequest)

	// 7. Panggil service
	response := controller.PostService.Create(request.Context(), postCreateRequest)
	log.Println("✅ Response dari Service:", response)

	// 8. Kirim response ke client
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
