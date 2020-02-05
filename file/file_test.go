package file

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReadFile(t *testing.T) {
	t.Run("Failed to open file", func(t *testing.T) {
		f := &File{}

		fileContent, err := f.ReadFile("")
		assert.NotNil(t, err)
		assert.Empty(t, fileContent)

	})

	t.Run("Failed to open file", func(t *testing.T) {
		f := &File{}

		fileContent, err := f.ReadFile("")
		assert.NotNil(t, err)
		assert.Empty(t, fileContent)

	})

	t.Run("Read file successfully", func(t *testing.T) {
		f := &File{}

		fileContent, err := f.ReadFile("myfile.txt")
		assert.Nil(t, err)

		assert.Equal(t, "this is myfile.txt content", fileContent)
	})
}

func Test_WriteFile(t *testing.T) {

	t.Run("Failed to write to file", func(t *testing.T) {
		f := &File{}

		fileName := "filetocreate.txt"
		content, err := f.WriteFile(fileName, "")
		assert.Empty(t, content)
		assert.NotNil(t, err)
		assert.Equal(t, "Empty content", err.Error())
	})

	t.Run("Create file successfully", func(t *testing.T) {
		f := &File{}

		fileName := "filecreated.txt"
		fileContent, err := f.WriteFile(fileName, "this is filecreated.txt content")
		assert.Nil(t, err)

		assert.Equal(t, "this is filecreated.txt content", fileContent)

		err = os.Remove(fileName)
		assert.Nil(t, err)
	})

}
