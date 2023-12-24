package register

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler interface {
	CreateUser(c *gin.Context)
	GetUserById(c *gin.Context)
}

type handlerImpl struct {
	svc Service
}

func NewHandler(svc Service) Handler {
	return &handlerImpl{svc}
}

func (h *handlerImpl) CreateUser(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user, "id": user.ID})
}

func (h *handlerImpl) GetUserById(c *gin.Context) {
	userID := c.Param("id")

	// Convert userID to uint, handle error if conversion fails.
	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call the service method to get the user by ID.
	user, err := h.svc.GetUserById(uint(userIDUint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var interestFaculty []InterestedFaculties
	interestFaculty, err = h.svc.GetInterestedFacultiesByUserId(uint(userIDUint))
	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	user.InterestedFaculties = interestFaculty

	c.JSON(http.StatusOK, gin.H{"user": user})
}
