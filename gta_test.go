package main

import (
	"fmt"
	"testing"

	"github.com/mitchellh/go-homedir"
)

func TestStorage(t *testing.T) {
	// initStorage()
	var home, _ = homedir.Dir()
	var store = initStorage(home + "/.gt-storage")

	if store == nil {
		t.Error("store is nil")
	}

	if store.aliasList() == "" {
		t.Error("aliasList is empty")
	}

	// add alias
	store.addAlias("test", "/home/xxx/test")
	if store.getAlias("test") != "/home/xxx/test" {
		t.Error("alias not added")
	}

	// rename alias
	store.renameAlias("test", "test2")
	if store.getAlias("test2") != "/home/xxx/test" {
		t.Error("alias not renamed")
	}

	// delete alias
	store.delAlias("test2")
	if store.getAlias("test2") != nil {
		t.Error("alias not deleted")
	}

	// clear alias
	store.addAlias("test", "/home/xxx/test")
	store.clearAlias()
	fmt.Println(store.aliasList())
	if store.aliasList() != "{}" {
		t.Error("alias not cleared")
	}

	store.addAlias("t", "/home/xxx/test")

	fmt.Println(store.aliasList())
	// > { "t": "/home/xxx/test" }
}
