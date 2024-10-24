package user

import (
	"backend/app/rpc/user/kitex_gen/user"
	"backend/library/tools"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"math/rand"
	"testing"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	//"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestLoginWithUsername(t *testing.T) {
	h := server.Default()
	h.POST("/v1/user/login/username", LoginWithUsername)
	path := "/v1/user/login/username"                         // todo: you can customize query
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} // todo: you can customize body
	header := ut.Header{}                                     // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "POST", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	// assert.DeepEqual(t, 200, resp.StatusCode())
	// assert.DeepEqual(t, "null", string(resp.Body()))
}

func TestLoginWithEmail(t *testing.T) {
	h := server.Default()
	h.POST("/v1/user/login/email", LoginWithEmail)
	path := "/v1/user/login/email"                            // todo: you can customize query
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} // todo: you can customize body
	header := ut.Header{}                                     // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "POST", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	// assert.DeepEqual(t, 200, resp.StatusCode())
	// assert.DeepEqual(t, "null", string(resp.Body()))
}

func TestLoginWithPhone(t *testing.T) {
	h := server.Default()
	h.POST("/v1/user/login/phone", LoginWithPhone)
	path := "/v1/user/login/phone"                            // todo: you can customize query
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} // todo: you can customize body
	header := ut.Header{}                                     // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "POST", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	// assert.DeepEqual(t, 200, resp.StatusCode())
	// assert.DeepEqual(t, "null", string(resp.Body()))
}

func TestRegister(t *testing.T) {
	h := server.Default()
	h.POST("/v1/user/register", Register)

	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(10000)

	// 定义多个测试用例
	tests := []struct {
		name           string
		request        user.RegisterReq
		expectedStatus int
		expectedBody   tools.Response
	}{
		{
			name: "Valid registration",
			request: user.RegisterReq{
				Username: fmt.Sprintf("test%d", randomInt), // 生成随机用户名
				Password: "Asdasd9277!!!",                  // 假设一个有效的密码
			},
			expectedStatus: 200,
			expectedBody: tools.Response{
				Code: 0,
				Msg:  "ok",
			},
		},
	}

	// 遍历每个测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 将请求数据转换为 JSON 格式
			reqBody, err := json.Marshal(tt.request)
			if err != nil {
				t.Fatalf("Failed to marshal request: %v", err)
			}
			fmt.Sprintf("Request Body: %s", reqBody)

			// 构造请求体
			body := &ut.Body{Body: bytes.NewBuffer(reqBody), Len: len(reqBody)}

			// 构造请求头
			header := ut.Header{
				Key:   "Content-Type",
				Value: "application/json",
			}

			// 发起请求
			w := ut.PerformRequest(h.Engine, "POST", "/v1/user/register", body, header)

			// 获取响应
			resp := w.Result()
			respBody := string(resp.Body())

			t.Logf("Response Status: %d", resp.StatusCode())
			t.Logf("Response Body: %s", respBody)

			assert.DeepEqual(t, tt.expectedStatus, resp.StatusCode())

			assert.DeepEqual(t, tt.expectedBody, respBody)
		})
	}
}

func TestUpdateUserInfo(t *testing.T) {
	h := server.Default()
	h.PUT("/v1/user/info", UpdateUserInfo)
	path := "/v1/user/info"                                   // todo: you can customize query
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} // todo: you can customize body
	header := ut.Header{}                                     // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "PUT", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	// assert.DeepEqual(t, 200, resp.StatusCode())
	// assert.DeepEqual(t, "null", string(resp.Body()))
}

func TestLogout(t *testing.T) {
	h := server.Default()
	h.POST("/v1/user/logout", Logout)
	path := "/v1/user/logout"                                 // todo: you can customize query
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} // todo: you can customize body
	header := ut.Header{}                                     // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "POST", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	// assert.DeepEqual(t, 200, resp.StatusCode())
	// assert.DeepEqual(t, "null", string(resp.Body()))
}

func TestFollowUser(t *testing.T) {
	h := server.Default()
	h.POST("/v1/user/follow", FollowUser)
	path := "/v1/user/follow"                                 // todo: you can customize query
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} // todo: you can customize body
	header := ut.Header{}                                     // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "POST", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	// assert.DeepEqual(t, 200, resp.StatusCode())
	// assert.DeepEqual(t, "null", string(resp.Body()))
}

func TestFollowerList(t *testing.T) {
	h := server.Default()
	h.GET("/v1/user/followers", FollowerList)
	path := "/v1/user/followers"                              // todo: you can customize query
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} // todo: you can customize body
	header := ut.Header{}                                     // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "GET", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	// assert.DeepEqual(t, 200, resp.StatusCode())
	// assert.DeepEqual(t, "null", string(resp.Body()))
}

func TestFollowingList(t *testing.T) {
	h := server.Default()
	h.GET("/v1/user/following", FollowingList)
	path := "/v1/user/following"                              // todo: you can customize query
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} // todo: you can customize body
	header := ut.Header{}                                     // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "GET", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	// assert.DeepEqual(t, 200, resp.StatusCode())
	// assert.DeepEqual(t, "null", string(resp.Body()))
}

func TestFriendList(t *testing.T) {
	h := server.Default()
	h.GET("/v1/user/friends", FriendList)
	path := "/v1/user/friends"                                // todo: you can customize query
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} // todo: you can customize body
	header := ut.Header{}                                     // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "GET", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	// assert.DeepEqual(t, 200, resp.StatusCode())
	// assert.DeepEqual(t, "null", string(resp.Body()))
}

func TestUserUploadFile(t *testing.T) {
	h := server.Default()
	h.POST("/v1/user/upload", UserUploadFile)
	path := "/v1/user/upload"                                 // todo: you can customize query
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} // todo: you can customize body
	header := ut.Header{}                                     // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "POST", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	// assert.DeepEqual(t, 200, resp.StatusCode())
	// assert.DeepEqual(t, "null", string(resp.Body()))
}
