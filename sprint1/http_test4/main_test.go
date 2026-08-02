package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserViewHandler(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	users := map[string]User{
		"u1": {
			ID:        "u1",
			FirstName: "Misha",
			LastName:  "Popov",
		},
		"u2": {
			ID:        "u2",
			FirstName: "Sasha",
			LastName:  "Popov",
		},
	}
	handler := UserViewHandler(users)

	type want struct {
		code        int
		user        User
		contentType string
	}

	tests := []struct {
		name   string
		userId string
		want   want
	}{
		{
			name: "Bad request test",
			want: want{
				code: http.StatusBadRequest,
			},
		},
		{
			name:   "Not found test",
			userId: "u3",
			want: want{
				code: http.StatusNotFound,
			},
		},
		{
			name:   "Success test",
			userId: "u1",
			want: want{
				code: http.StatusOK,
				user: users["u1"],
			},
		},
	}
	for _, tt := range tests { // цикл по всем тестам
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/status", nil)

			// добавляем GET параметры
			if tt.userId != "" {
				q := req.URL.Query()
				q.Add("user_id", tt.userId)
				req.URL.RawQuery = q.Encode()
			}

			// создаём новый Recorder
			w := httptest.NewRecorder()
			handler(w, req)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(tt.want.code, res.StatusCode, res.StatusCode)
			// получаем и проверяем тело запроса
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			require.NoError(err)
			if res.StatusCode == http.StatusOK {
				assert.Equal(tt.want.contentType, res.Header.Get("Content-Type"))

				var user User
				err = json.Unmarshal(resBody, &user)
				require.NoError(err)

				assert.Equal(tt.want.user, user)
			}
		})
	}
}
