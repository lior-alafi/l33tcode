package service

import (
	"l33tcode/server/pkg/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type testBed struct {
	codeExecutorsMap    map[string]*mocks.MockCodeExecuter
	currentCodeExecutor string
	questionRepo        *mocks.MockQuestionRepository
	logger              *zap.Logger
	languageRepo        *mocks.NewMockLanguageRepository
	svc                 Service
}

func initTestBed(ctrl *gomock.Controller) *testBed {

	tb := &testBed{}
	tb.questionRepo = mocks.NewMockQuestionRepository(ctrl)
	tb.languageRepo = mocks.NewMockLanguageRepository(ctrl)
	tb.logger = zap.NewNop()
	tb.codeExecutorsMap = map[string]*mocks.MockCodeExecuter{
		"a": mocks.MockCodeExecuter(ctrl),
	}
	tb.svc = NewService(tb.logger, tb.questionRepo, tb.languageRepo, tb.codeExecutorsMap, "a")
	return tb
}
func TestSubmitCodeBadRequest(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	tb := initTestBed(ctrl)
	router := gin.Default()
	router.GET("/admin/question/submit", tb.svc.SubmitQuestion)

	req, _ := http.NewRequest("GET", "/admin/question/submit?user=admin", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	require.Equal(t, http.StatusMethodNotAllowed, w.Code)

	// // Assert response body
	// expected := `{"message":"Hello Gopher"}`
}
func TestListCodeExecutors(t *testing.T) {

}

func TestSetCodeExecutor(t *testing.T) {

}

func TestSubmitQuestion(t *testing.T) {

}
func TestGetQuestion(t *testing.T) {

}
func TestRemoveQuestion(t *testing.T) {

}
func TestListQuestions(t *testing.T) {

}
func TestUpdateQuestion(t *testing.T) {

}

func TestListSupportedLanguages(t *testing.T) {

}
func TestGetLanguage(t *testing.T) {

}
