package controllers_test

import (
	controllers "go-blog/controller"
	"go-blog/controller/testdata"
	"testing"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}