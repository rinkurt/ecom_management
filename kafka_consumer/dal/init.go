package dal

import (
	"kafka_consumer/dal/ch"
	"kafka_consumer/dal/es"
)

func Init() {
	es.Init()
	ch.Init()
}

func Close() {

}
