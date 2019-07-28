package daos

import (
	"fmt"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/config"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProjectDAO_FindAll(t *testing.T) {
	config.Config.DB = test_data.ResetDB()
	dao := NewProjectDAO()

	projects := dao.FindAll()

	fmt.Printf("%v", projects)

	assert.Equal(t, 2, len(projects))
}
