package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/redis_practice/repo"
)

type Handler struct {
	customRedis repo.RedisClient
}

func New(customRedis *repo.RedisClient) *Handler {
	return &Handler{
		customRedis: *customRedis,
	}
}

func (h *Handler) SetKeyValueHandler(c *gin.Context) {
	key := c.Param("key")
	value := c.Param("value")
	expiration := 3600

	err := h.customRedis.Put(context.Background(), key, value, expiration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set value in Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Value set successfully"})
}

func (h *Handler) GetKeyValueHandler(c *gin.Context) {
	key := c.Param("key")

	result, err := h.customRedis.Get(context.Background(), key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get value from Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"value": result})
}

func (h *Handler) DeleteKeyHandler(c *gin.Context) {
	key := c.Param("key")

	_, _, errs := h.customRedis.Del(context.Background(), key)
	if len(errs) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete key from Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Key deleted successfully"})
}

func (h *Handler) ExistsKeyHandler(c *gin.Context) {
	key := c.Param("key")

	result, err := h.customRedis.Exists(context.Background(), key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check key existence in Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"exists": result})
}

func (h *Handler) AddToSetHandler(c *gin.Context) {
	setname := c.Param("setname")
	member := c.Param("member")

	err := h.customRedis.AddToSet(context.Background(), setname, member)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add member to set in Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Member added to set successfully"})
}

func (h *Handler) GetFromSetHandler(c *gin.Context) {
	setname := c.Param("setname")
	member := c.Param("member")

	result, err := h.customRedis.GetFromSet(context.Background(), setname, member)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check member in set in Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"exists": result})
}

func (h *Handler) RemoveFromSetHandler(c *gin.Context) {
	setname := c.Param("setname")
	member := c.Param("member")

	err := h.customRedis.RemoveFromSet(context.Background(), setname, member)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove member from set in Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Member removed from set successfully"})
}

func (h *Handler) AddToHashHandler(c *gin.Context) {
	hashname := c.Param("hashname")
	key := c.Param("key")
	value := c.Param("value")
	expiration := 3600 // Example expiration time in seconds

	err := h.customRedis.AddToHash(context.Background(), hashname, key, value, expiration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add key to hash in Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Key added to hash successfully"})
}

func (h *Handler) RemoveFromHashHandler(c *gin.Context) {
	hashname := c.Param("hashname")
	key := c.Param("key")

	err := h.customRedis.RemoveFromHash(context.Background(), hashname, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove key from hash in Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Key removed from hash successfully"})
}

func (h *Handler) ExistsInHashHandler(c *gin.Context) {
	hashname := c.Param("hashname")
	key := c.Param("key")

	result, err := h.customRedis.ExistsInHash(context.Background(), hashname, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check key existence in hash in Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"exists": result})
}

func (h *Handler) GetAllFromHashHandler(c *gin.Context) {
	hashname := c.Param("hashname")

	result, err := h.customRedis.GetAllFromHash(context.Background(), hashname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all keys from hash in Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (h *Handler) LeftPushHandler(c *gin.Context) {
	listname := c.Param("listname")
	value := c.Param("value")

	err := h.customRedis.LeftPush(context.Background(), listname, value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to push value to list in Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Value pushed to list successfully"})
}

func (h *Handler) RightPushHandler(c *gin.Context) {
	listname := c.Param("listname")
	value := c.Param("value")

	err := h.customRedis.RightPush(context.Background(), listname, value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to push value to list in Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Value pushed to list successfully"})
}

func (h *Handler) PopLeftHandler(c *gin.Context) {
	listname := c.Param("listname")

	result, err := h.customRedis.PopLeft(context.Background(), listname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to pop value from list in Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"value": result})
}

func (h *Handler) PopRightHandler(c *gin.Context) {
	listname := c.Param("listname")

	result, err := h.customRedis.PopRight(context.Background(), listname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to pop value from list in Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"value": result})
}

func (h *Handler) ListLengthHandler(c *gin.Context) {
	listname := c.Param("listname")

	result, err := h.customRedis.ListLength(context.Background(), listname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get list length from Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"length": result})
}

func (h *Handler) GetRangeElementsHandler(c *gin.Context) {
	listname := c.Param("listname")
	from := c.Param("from")
	to := c.Param("to")

	fromInt64, err := strconv.ParseInt(from, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'from' parameter"})
		return
	}

	toInt64, err := strconv.ParseInt(to, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'to' parameter"})
		return
	}

	result, err := h.customRedis.GetRangeElements(context.Background(), listname, fromInt64, toInt64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get range elements from list in Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"elements": result})
}


