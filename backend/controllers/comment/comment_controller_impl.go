package comment

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fzndps/mini-social-media/backend/exception"
	"github.com/fzndps/mini-social-media/backend/helper"
	"github.com/fzndps/mini-social-media/backend/models/web"
	"github.com/fzndps/mini-social-media/backend/services/comment"
	"github.com/julienschmidt/httprouter"
)

type CommentControllerImpl struct {
	CommentService comment.CommentService
}

func NewCommentController(commentService comment.CommentService) CommentController {
	return &CommentControllerImpl{
		CommentService: commentService,
	}
}

func (controller *CommentControllerImpl) CreateComment(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	log.Println("🚀 Masuk controller CreateComment")

	userId, err := helper.GetUserIdFromRequest(request)
	log.Println("✅ User ID dari JWT:", userId)
	if err != nil {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}

	postIdParam := params.ByName("postId")
	postId, err := strconv.Atoi(postIdParam)
	if err != nil {
		panic(exception.NewConflictRequestError("Post ID tidak valid"))
	}

	commentCreateRequest := web.CommentCreateRequest{}
	helper.ReadFromRequestBody(request, &commentCreateRequest)

	commentCreateRequest.PostId = postId
	commentCreateRequest.UserId = userId

	commentResponse := controller.CommentService.CreateComment(request.Context(), commentCreateRequest)
	log.Println("comment response pada controller :", commentResponse)
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "COMMENT CREATED",
		Data:   commentResponse,
	}

	log.Println("🚨 DEBUG - PostID to insert comment:", commentCreateRequest.PostId)

	log.Println("web response pada controller :", webResponse)

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CommentControllerImpl) FindPostWithCommentsById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	postIdStr := params.ByName("postId")
	postId, err := strconv.Atoi(postIdStr)
	helper.PanicIfError(err)

	commentResponse, err := controller.CommentService.FindPostWithCommentsById(request.Context(), postId)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "POST NOT FOUND",
			Data:   nil,
		}
		helper.WriteResponseBody(writer, webResponse)
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   commentResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CommentControllerImpl) UpdateComment(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	log.Println("Masuk ke controller")
	commentId, err := strconv.Atoi(params.ByName("commentId"))
	helper.PanicIfError(err)

	userId, err := helper.GetUserIdFromRequest(request)
	helper.PanicIfError(err)

	commentUpdateRequest := web.UpdateCommentRequest{}
	helper.ReadFromRequestBody(request, &commentUpdateRequest)
	log.Println("Comment Update Request :", commentUpdateRequest)

	commentResponse := controller.CommentService.UpdateComment(request.Context(), commentId, userId, commentUpdateRequest)
	log.Println("CommentWebResponse :", commentResponse)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "UPDATED",
		Data:   commentResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CommentControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	comId := params.ByName("commentId")
	commentId, err := strconv.Atoi(comId)
	helper.PanicIfError(err)

	poId := params.ByName("postId")
	postId, err := strconv.Atoi(poId)
	helper.PanicIfError(err)

	userId, err := helper.GetUserIdFromRequest(request)
	helper.PanicIfError(err)

	controller.CommentService.Delete(request.Context(), commentId, postId, userId)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "COMMENT DELETED",
	}

	helper.WriteResponseBody(writer, webResponse)
}
