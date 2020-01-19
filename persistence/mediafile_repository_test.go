package persistence

import (
	"os"
	"path/filepath"

	"github.com/cloudsonic/sonic-server/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MediaFileRepository", func() {
	var repo model.MediaFileRepository

	BeforeEach(func() {
		repo = NewMediaFileRepository()
	})

	Describe("FindByPath", func() {
		It("returns all records from a given ArtistID", func() {
			path := string(os.PathSeparator) + filepath.Join("beatles", "1")
			println("Searching path", path) // TODO Remove
			Expect(repo.FindByPath(path)).To(Equal(model.MediaFiles{
				songComeTogether,
			}))
		})
	})

})