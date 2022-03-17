package utils

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func GetCurrentPath() (dir string, isGoRun bool, err error) {
	dir, err = getPathByExecutable()
	if err != nil {
		return
	}

	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		return getPathByCaller(), true, nil
	}

	return dir, false, nil
}

// getPathByExecutable 获取当前执行文件绝对路径
func getPathByExecutable() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))

	return res, nil
}

// getPathByCaller 获取当前执行文件绝对路径（go run）
func getPathByCaller() string {
	var dir string
	_, filename, _, ok := runtime.Caller(2)
	if ok {
		dir = path.Dir(filename)
	}

	return dir
}
