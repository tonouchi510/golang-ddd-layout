// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Circles", testCircles)
	t.Run("CircleMembers", testCircleMembers)
	t.Run("Users", testUsers)
}

func TestSoftDelete(t *testing.T) {
	t.Run("Circles", testCirclesSoftDelete)
	t.Run("CircleMembers", testCircleMembersSoftDelete)
	t.Run("Users", testUsersSoftDelete)
}

func TestQuerySoftDeleteAll(t *testing.T) {
	t.Run("Circles", testCirclesQuerySoftDeleteAll)
	t.Run("CircleMembers", testCircleMembersQuerySoftDeleteAll)
	t.Run("Users", testUsersQuerySoftDeleteAll)
}

func TestSliceSoftDeleteAll(t *testing.T) {
	t.Run("Circles", testCirclesSliceSoftDeleteAll)
	t.Run("CircleMembers", testCircleMembersSliceSoftDeleteAll)
	t.Run("Users", testUsersSliceSoftDeleteAll)
}

func TestDelete(t *testing.T) {
	t.Run("Circles", testCirclesDelete)
	t.Run("CircleMembers", testCircleMembersDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Circles", testCirclesQueryDeleteAll)
	t.Run("CircleMembers", testCircleMembersQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Circles", testCirclesSliceDeleteAll)
	t.Run("CircleMembers", testCircleMembersSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Circles", testCirclesExists)
	t.Run("CircleMembers", testCircleMembersExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Circles", testCirclesFind)
	t.Run("CircleMembers", testCircleMembersFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Circles", testCirclesBind)
	t.Run("CircleMembers", testCircleMembersBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Circles", testCirclesOne)
	t.Run("CircleMembers", testCircleMembersOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Circles", testCirclesAll)
	t.Run("CircleMembers", testCircleMembersAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Circles", testCirclesCount)
	t.Run("CircleMembers", testCircleMembersCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Circles", testCirclesHooks)
	t.Run("CircleMembers", testCircleMembersHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Circles", testCirclesInsert)
	t.Run("Circles", testCirclesInsertWhitelist)
	t.Run("CircleMembers", testCircleMembersInsert)
	t.Run("CircleMembers", testCircleMembersInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("CircleMemberToCircleUsingCircle", testCircleMemberToOneCircleUsingCircle)
	t.Run("CircleMemberToUserUsingMember", testCircleMemberToOneUserUsingMember)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("CircleToCircleMembers", testCircleToManyCircleMembers)
	t.Run("UserToMemberCircleMembers", testUserToManyMemberCircleMembers)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("CircleMemberToCircleUsingCircleMembers", testCircleMemberToOneSetOpCircleUsingCircle)
	t.Run("CircleMemberToUserUsingMemberCircleMembers", testCircleMemberToOneSetOpUserUsingMember)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("CircleToCircleMembers", testCircleToManyAddOpCircleMembers)
	t.Run("UserToMemberCircleMembers", testUserToManyAddOpMemberCircleMembers)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Circles", testCirclesReload)
	t.Run("CircleMembers", testCircleMembersReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Circles", testCirclesReloadAll)
	t.Run("CircleMembers", testCircleMembersReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Circles", testCirclesSelect)
	t.Run("CircleMembers", testCircleMembersSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Circles", testCirclesUpdate)
	t.Run("CircleMembers", testCircleMembersUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Circles", testCirclesSliceUpdateAll)
	t.Run("CircleMembers", testCircleMembersSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
