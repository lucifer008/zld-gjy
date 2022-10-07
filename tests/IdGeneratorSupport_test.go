package tests

import (
	"log"
	"testing"
	"zld-jy/da/support"
)

func Test_Generator_Id(t *testing.T) {
	var id = support.IdGeneratorInstance.NextId()
	log.Println(id)
}
