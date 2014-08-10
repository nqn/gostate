package zookeeper

import (
  "github.com/samuel/go-zookeeper/zk"
)

type ZookeeperStorage struct {}

func New() {
  return &ZookeeperStorage {}
}
