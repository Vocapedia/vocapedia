package snowflake

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func InitSnowflake() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Snowflake is ready")
}

func GenerateID() int64 {
	return node.Generate().Int64()
}
