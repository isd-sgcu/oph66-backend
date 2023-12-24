package register

import (
	"github.com/gin-gonic/gin"
	//bcrypt
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

// Handler interface defines methods for handling user-related operations.
type Handler interface {
	CreateUser(c *gin.Context)
	GetUserById(c *gin.Context)
}

// handlerImpl is the implementation of the Handler interface.
type handlerImpl struct {
	svc Service
}

// NewHandler creates a new instance of the Handler interface.
func NewHandler(svc Service) Handler {
	return &handlerImpl{svc}
}

func (h *handlerImpl) CreateUser(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password before saving it to the database.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.HashedPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.HashedPassword = string(hashedPassword)

	// Call the service method to create the user.
	if err := h.svc.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	// Return the user information, including the generated ID.
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user, "id": user.ID})
}

// GetUserById retrieves a user by ID.
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
		// User not found.
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Remove the hashed password from the user object before returning it.
	user.HashedPassword = ""

	//add interestFaculty
	var interestFaculty []InterestedFaculties
	interestFaculty, err = h.svc.GetInterestedFacultiesByUserId(uint(userIDUint))
	if err != nil {
		// User not found.
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	user.InterestedFaculties = interestFaculty

	// User found, return user information.
	c.JSON(http.StatusOK, gin.H{"user": user})
}
