package docker

import (
	"github.com/ugurcsen/sand-panel/internal/core/domain"
	"path"
)

func (d docker) getUserPath(c *domain.Collection) string {
	return path.Join(d.BaseDir, c.UserId)
}

func (d docker) getCollectionPath(c *domain.Collection) string {
	return path.Join(d.getUserPath(c), c.UserId+"_"+c.Id)
}

func (d docker) getComposePath(c *domain.Collection) string {
	return path.Join(d.getCollectionPath(c), "docker-compose.yml")
}
