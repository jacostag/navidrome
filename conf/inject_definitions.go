package conf

import (
	"github.com/deluan/gosonic/domain"
	"github.com/deluan/gosonic/persistence"
	"github.com/deluan/gosonic/utils"
)

func init() {
	utils.DefineSingleton(new(domain.ArtistIndexRepository), persistence.NewArtistIndexRepository)
	utils.DefineSingleton(new(domain.PropertyRepository), persistence.NewPropertyRepository)
	utils.DefineSingleton(new(domain.MediaFolderRepository), persistence.NewMediaFolderRepository)
}