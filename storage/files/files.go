package files

import (
	e "TelegramBot/lib/error"
	"TelegramBot/storage"
	"os"
	"path/filepath"
)

type Storage struct {
	basePath string
}

const defaultPerm = 0774

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}
func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = e.WrapIfErr("Can't save", err) }()
	filePath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(filePath, defaultPerm); err != nil {
		return err
	}

}
