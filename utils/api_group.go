package utils

import (
	"fmt"
	"go-starterkit-project/config"
)

/**
This function is used to create grouping functions like version api

ex: /api/v1
*/
func SetupApiGroup() string {
	str := fmt.Sprintf("/%s/%s", config.Config("API_WRAP"), config.Config("API_VERSION"))
	return str
}
