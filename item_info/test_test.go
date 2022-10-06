package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
	"time"
)

type Visibility struct {
	Source string
	Status int
	Rate   int
}

const (
	StatusPrivate = 3
	StatusPublic  = 1
)

func CalcItemFinalStatus(vis []Visibility) (Visibility,
	error) {
	return Visibility{Source: "user", Status: 2, Rate: 3}, nil
}

func TestCalcItemFinalStatus(t *testing.T) {
	Convey("Case 1", t, func() {
		vis := []Visibility{
			{Source: "user", Status: StatusPrivate},
			{Source: "mgt", Status: StatusPublic},
		}
		res, _ := CalcItemFinalStatus(vis)
		So(res.Status, ShouldEqual, StatusPrivate)
	})
	// ......
}

func TestCalcItemDiff(t *testing.T) {
	r := rand.Float64() * 50
	time.Sleep(time.Duration(r) * time.Millisecond)
}

func TestUpsertItem(t *testing.T) {
	r := rand.Float64()*100 + 200
	time.Sleep(time.Duration(r) * time.Millisecond)
}

func TestUploadItem(t *testing.T) {
	r := rand.Float64()*100 + 300
	time.Sleep(time.Duration(r) * time.Millisecond)
}

func TestDeleteItem(t *testing.T) {
	r := rand.Float64()*100 + 200
	time.Sleep(time.Duration(r) * time.Millisecond)
}

func TestIncrCount(t *testing.T) {
	r := rand.Float64()*100 + 100
	time.Sleep(time.Duration(r) * time.Millisecond)
}
