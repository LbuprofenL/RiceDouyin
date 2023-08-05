package util

import (
	"github.com/bwmarrin/snowflake"
)

func CreateNode() (*snowflake.Node, error) {
	// 创建一个新的雪花节点，传入一个唯一的Node ID
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}
	return node, nil
}
