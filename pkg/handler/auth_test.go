package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/magiconair/properties/assert"
	"github.com/victorvelo/notes/internal/models"
	"github.com/victorvelo/notes/pkg/service"
	mock_service "github.com/victorvelo/notes/pkg/service/mocks"
)

func TestAuth_signup(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockAuthorization, user models.User)

	tester := []struct {
		name                 string
		inputBody            string
		inputUser            models.User
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"login":"firstPerson","password":"1234"}`,
			inputUser: models.User{
				Login:    "firstPerson",
				Password: "1234",
			},
			mockBehaviour: func(s *mock_service.MockAuthorization, user models.User) {
				s.EXPECT().Add(user).Return(nil)
			},
			expectedStatusCode:   201,
			expectedResponseBody: `{"login":"firstPerson","password":"1234"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"login":"firstPerson","password":"1234"}`,
			inputUser: models.User{
				Login:    "firstPerson",
				Password: "1234",
			},
			mockBehaviour: func(s *mock_service.MockAuthorization, user models.User) {
				s.EXPECT().Add(user).Return(errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"error":"something went wrong"}`,
		},
	}

	for _, test := range tester {
		t.Run(test.name, func(t *testing.T) {

			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			test.mockBehaviour(auth, test.inputUser)
			service := &service.Service{Authorization: auth}
			handlerTest := NewHandler(service)
			router := mux.NewRouter()
			router.HandleFunc("/sign-up", handlerTest.signUp).Methods(http.MethodPost)
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/sign-up",
				bytes.NewBufferString(test.inputBody))

			router.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
