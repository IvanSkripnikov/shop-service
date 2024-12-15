package logger

import (
	"fmt"
	"os"
	"strings"
)

// Получить имя текущего сервиса.
func getContainerName() string {
	// TODO вынести всё в envs
	//containerName := os.Getenv("CONTAINER_NAME")
	containerName := "container-name"

	if len(containerName) == 0 {
		containerName = serviceNameDefault
	}

	return containerName
}

// Получить путь к файлу с логами.
func getLogFilePath() string {
	return fmt.Sprintf("./log/%s.log", getContainerName())
}

// Получить имя хоста.
func getHostName() string {
	var hostName string
	hostNameFile, err := os.ReadFile("/etc/hostname")
	if err != nil {
		serverName, _ := os.Hostname()
		hostName = serverName
	} else {
		hostName = strings.ReplaceAll(string(hostNameFile), "\n", "")
	}

	return hostName
}
