package collision

import (
	"bacteria/component"

	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
)

func AddToSpace(space *donburi.Entry, objects ...*donburi.Entry) {
	for _, obj := range objects {
		component.Space.Get(space).Add(GetObject(obj))
	}
}

func SetObject(entry *donburi.Entry, obj *resolv.Object) {
	component.CollideBox.Set(entry, obj)
}

func GetObject(entry *donburi.Entry) *resolv.Object {
	return component.CollideBox.Get(entry)
}

func Remove(space *donburi.Entry, objects ...*donburi.Entry) {
	for _, obj := range objects {
		component.Space.Get(space).Remove(GetObject(obj))
	}
}
