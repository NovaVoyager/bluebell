package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func TestCreatePostHandler(t *testing.T) {
	url := "/api/v1/post/create"
	r := gin.Default()
	r.POST(url, CreatePostHandler)
	post := `{"title":"德玛西亚","content":"人在塔在","community_id":1}`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(post)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "token无效")
}
