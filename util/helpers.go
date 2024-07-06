package util

import "errors"

func GetRoleByTypeId(TypeId int) (string, error) {

	Roles := map[int]string{
		1: "Admin",
		2: "Customer",
	}

	found, ok := Roles[TypeId]
	if !ok {
		return "", errors.New("Tidak dapat menemukan Role Id")
	}

	return found, nil
}
