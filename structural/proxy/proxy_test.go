package proxy

import (
	"math/rand"
	"testing"
)

func Test_UserListProxy(t *testing.T) {
	someDatabase := UserList{}
	rand.Seed(2342342)
	for i := 0; i < 10000; i++ {
		n := rand.Int31()
		someDatabase = append(someDatabase, User{ID: n})
	}

	proxy := UserListProxy{
		SomeDatabase:  someDatabase,
		StackCapacity: 2,
		StackCache:    UserList{},
	}

	knownIDs := [3]int32{
		someDatabase[3].ID,
		someDatabase[4].ID,
		someDatabase[5].ID,
	}

	t.Run("FindUser - EmptyCache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		if user.ID != knownIDs[0] {
			t.Error("Returned user name doesn't match with expected")
		}

		if len(proxy.StackCache) != 1 {
			t.Error("After One Successful search in empty cache, the cache of it must be one")
		}

		if proxy.DidLastSearchUsedCache {
			t.Error("No User can be returned from an empty cache")
		}
	})

	t.Run("FindUser User one, ask for the same User", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		if user.ID != knownIDs[0] {
			t.Error("Returned user name doesn't  match with expected")
		}

		if len(proxy.StackCache) != 1 {
			t.Error("Cache must not grow if we asked for an Object that is stored on it")
		}

		if !proxy.DidLastSearchUsedCache {
			t.Error("The User Should have been returned from cache ")
		}

	})

	user1, err := proxy.FindUser(knownIDs[0])
	if err != nil {
		t.Fatal(err)
	}

	user2, _ := proxy.FindUser(knownIDs[1])
	if proxy.DidLastSearchUsedCache {
		t.Error("the last does't store one proxy yet")
	}

	user3, _ := proxy.FindUser(knownIDs[2])
	if proxy.DidLastSearchUsedCache {
		t.Error("the last does't store one proxy yet")
	}

	for i := 0; i < len(proxy.StackCache); i++ {
		if proxy.StackCache[i].ID == user1.ID {
			t.Error("User that should be gone was found")
		}
	}

	if len(proxy.StackCache) != 2 {
		t.Error("After inserting 3 users the cache should not grow more than two")
	}

	for _, v := range proxy.StackCache {
		if v != user2 && v != user3 {
			t.Error("a non expected user was found on the cache ")
		}
	}

}
