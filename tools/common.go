package tools

import (
	"github.com/latifrons/commongo/files"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
)

type FolderConfig struct {
	Root    string
	Log     string
	Data    string
	Config  string
	Private string
}

func EnsureFolder(folder string, perm os.FileMode) {
	err := files.MkDirPermIfNotExists(folder, perm)
	if err != nil {
		logrus.WithError(err).WithField("path", folder).Fatal("failed to create folder")
	}
}

func defaultPath(givenPath string, defaultRoot string, suffix string) string {
	if givenPath == "" {
		return path.Join(defaultRoot, suffix)
	}
	if path.IsAbs(givenPath) {
		return givenPath
	}
	return path.Join(defaultRoot, givenPath)
}

func EnsureFolders() FolderConfig {
	config := FolderConfig{
		Root:    viper.GetString("dir.root"),
		Log:     defaultPath(viper.GetString("dir.log"), viper.GetString("dir.root"), "log"),
		Data:    defaultPath(viper.GetString("dir.data"), viper.GetString("dir.root"), "data"),
		Config:  defaultPath(viper.GetString("dir.config"), viper.GetString("dir.root"), "config"),
		Private: defaultPath(viper.GetString("dir.private"), viper.GetString("dir.root"), "private"),
	}
	EnsureFolder(config.Root, 0755)
	EnsureFolder(config.Log, 0755)
	EnsureFolder(config.Data, 0755)
	EnsureFolder(config.Config, 0755)
	EnsureFolder(config.Private, 0700)
	return config

}
