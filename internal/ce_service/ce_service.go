package ce_service

import (
	"fmt"
	"strings"

	"go.arcalot.io/imagebuilder/internal/docker"
)

type ContainerEngineService interface {
	Build(filepath string, name string, tags []string, quay_img_exp string) error
	Tag(image_tag string, destination string) error
	Push(destination string, username string, password string, registry_address string) error
}

func NewContainerEngineService(choice string) (ContainerEngineService, error) {
	choice = strings.ToLower(choice)
	switch choice {
	case "podman":
		return nil, fmt.Errorf("podman is not supported yet")
	default: // docker
		return docker.NewCEClient()
	}
}