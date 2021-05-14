// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testCircleMembers(t *testing.T) {
	t.Parallel()

	query := CircleMembers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCircleMembersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CircleMember{}
	if err = randomize.Struct(seed, o, circleMemberDBTypes, true, circleMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CircleMembers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCircleMembersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CircleMember{}
	if err = randomize.Struct(seed, o, circleMemberDBTypes, true, circleMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := CircleMembers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CircleMembers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCircleMembersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CircleMember{}
	if err = randomize.Struct(seed, o, circleMemberDBTypes, true, circleMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CircleMemberSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CircleMembers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCircleMembersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CircleMember{}
	if err = randomize.Struct(seed, o, circleMemberDBTypes, true, circleMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CircleMemberExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if CircleMember exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CircleMemberExists to return true, but got false.")
	}
}

func testCircleMembersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CircleMember{}
	if err = randomize.Struct(seed, o, circleMemberDBTypes, true, circleMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	circleMemberFound, err := FindCircleMember(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if circleMemberFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCircleMembersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CircleMember{}
	if err = randomize.Struct(seed, o, circleMemberDBTypes, true, circleMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = CircleMembers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCircleMembersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CircleMember{}
	if err = randomize.Struct(seed, o, circleMemberDBTypes, true, circleMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := CircleMembers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCircleMembersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	circleMemberOne := &CircleMember{}
	circleMemberTwo := &CircleMember{}
	if err = randomize.Struct(seed, circleMemberOne, circleMemberDBTypes, false, circleMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}
	if err = randomize.Struct(seed, circleMemberTwo, circleMemberDBTypes, false, circleMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = circleMemberOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = circleMemberTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CircleMembers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCircleMembersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	circleMemberOne := &CircleMember{}
	circleMemberTwo := &CircleMember{}
	if err = randomize.Struct(seed, circleMemberOne, circleMemberDBTypes, false, circleMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}
	if err = randomize.Struct(seed, circleMemberTwo, circleMemberDBTypes, false, circleMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = circleMemberOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = circleMemberTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CircleMembers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func circleMemberBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *CircleMember) error {
	*o = CircleMember{}
	return nil
}

func circleMemberAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *CircleMember) error {
	*o = CircleMember{}
	return nil
}

func circleMemberAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *CircleMember) error {
	*o = CircleMember{}
	return nil
}

func circleMemberBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *CircleMember) error {
	*o = CircleMember{}
	return nil
}

func circleMemberAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *CircleMember) error {
	*o = CircleMember{}
	return nil
}

func circleMemberBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *CircleMember) error {
	*o = CircleMember{}
	return nil
}

func circleMemberAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *CircleMember) error {
	*o = CircleMember{}
	return nil
}

func circleMemberBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *CircleMember) error {
	*o = CircleMember{}
	return nil
}

func circleMemberAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *CircleMember) error {
	*o = CircleMember{}
	return nil
}

func testCircleMembersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &CircleMember{}
	o := &CircleMember{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, circleMemberDBTypes, false); err != nil {
		t.Errorf("Unable to randomize CircleMember object: %s", err)
	}

	AddCircleMemberHook(boil.BeforeInsertHook, circleMemberBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	circleMemberBeforeInsertHooks = []CircleMemberHook{}

	AddCircleMemberHook(boil.AfterInsertHook, circleMemberAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	circleMemberAfterInsertHooks = []CircleMemberHook{}

	AddCircleMemberHook(boil.AfterSelectHook, circleMemberAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	circleMemberAfterSelectHooks = []CircleMemberHook{}

	AddCircleMemberHook(boil.BeforeUpdateHook, circleMemberBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	circleMemberBeforeUpdateHooks = []CircleMemberHook{}

	AddCircleMemberHook(boil.AfterUpdateHook, circleMemberAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	circleMemberAfterUpdateHooks = []CircleMemberHook{}

	AddCircleMemberHook(boil.BeforeDeleteHook, circleMemberBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	circleMemberBeforeDeleteHooks = []CircleMemberHook{}

	AddCircleMemberHook(boil.AfterDeleteHook, circleMemberAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	circleMemberAfterDeleteHooks = []CircleMemberHook{}

	AddCircleMemberHook(boil.BeforeUpsertHook, circleMemberBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	circleMemberBeforeUpsertHooks = []CircleMemberHook{}

	AddCircleMemberHook(boil.AfterUpsertHook, circleMemberAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	circleMemberAfterUpsertHooks = []CircleMemberHook{}
}

func testCircleMembersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CircleMember{}
	if err = randomize.Struct(seed, o, circleMemberDBTypes, true, circleMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CircleMembers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCircleMembersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CircleMember{}
	if err = randomize.Struct(seed, o, circleMemberDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(circleMemberColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := CircleMembers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCircleMemberToOneCircleUsingCircle(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local CircleMember
	var foreign Circle

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, circleMemberDBTypes, false, circleMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, circleDBTypes, false, circleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Circle struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.CircleID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Circle().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := CircleMemberSlice{&local}
	if err = local.L.LoadCircle(ctx, tx, false, (*[]*CircleMember)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Circle == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Circle = nil
	if err = local.L.LoadCircle(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Circle == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCircleMemberToOneUserUsingMember(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local CircleMember
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, circleMemberDBTypes, false, circleMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.MemberID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Member().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := CircleMemberSlice{&local}
	if err = local.L.LoadMember(ctx, tx, false, (*[]*CircleMember)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Member == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Member = nil
	if err = local.L.LoadMember(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Member == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCircleMemberToOneSetOpCircleUsingCircle(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CircleMember
	var b, c Circle

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, circleMemberDBTypes, false, strmangle.SetComplement(circleMemberPrimaryKeyColumns, circleMemberColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, circleDBTypes, false, strmangle.SetComplement(circlePrimaryKeyColumns, circleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, circleDBTypes, false, strmangle.SetComplement(circlePrimaryKeyColumns, circleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Circle{&b, &c} {
		err = a.SetCircle(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Circle != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.CircleMembers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CircleID != x.ID {
			t.Error("foreign key was wrong value", a.CircleID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CircleID))
		reflect.Indirect(reflect.ValueOf(&a.CircleID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CircleID != x.ID {
			t.Error("foreign key was wrong value", a.CircleID, x.ID)
		}
	}
}
func testCircleMemberToOneSetOpUserUsingMember(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CircleMember
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, circleMemberDBTypes, false, strmangle.SetComplement(circleMemberPrimaryKeyColumns, circleMemberColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetMember(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Member != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.MemberCircleMembers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.MemberID != x.ID {
			t.Error("foreign key was wrong value", a.MemberID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.MemberID))
		reflect.Indirect(reflect.ValueOf(&a.MemberID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.MemberID != x.ID {
			t.Error("foreign key was wrong value", a.MemberID, x.ID)
		}
	}
}

func testCircleMembersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CircleMember{}
	if err = randomize.Struct(seed, o, circleMemberDBTypes, true, circleMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCircleMembersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CircleMember{}
	if err = randomize.Struct(seed, o, circleMemberDBTypes, true, circleMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CircleMemberSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCircleMembersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CircleMember{}
	if err = randomize.Struct(seed, o, circleMemberDBTypes, true, circleMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CircleMembers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	circleMemberDBTypes = map[string]string{`ID`: `int`, `CircleID`: `char`, `MemberID`: `char`, `CreatedAt`: `datetime`, `UpdatedAt`: `datetime`, `DeletedAt`: `datetime`}
	_                   = bytes.MinRead
)

func testCircleMembersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(circleMemberPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(circleMemberAllColumns) == len(circleMemberPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CircleMember{}
	if err = randomize.Struct(seed, o, circleMemberDBTypes, true, circleMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CircleMembers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, circleMemberDBTypes, true, circleMemberPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCircleMembersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(circleMemberAllColumns) == len(circleMemberPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CircleMember{}
	if err = randomize.Struct(seed, o, circleMemberDBTypes, true, circleMemberColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CircleMembers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, circleMemberDBTypes, true, circleMemberPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(circleMemberAllColumns, circleMemberPrimaryKeyColumns) {
		fields = circleMemberAllColumns
	} else {
		fields = strmangle.SetComplement(
			circleMemberAllColumns,
			circleMemberPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := CircleMemberSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCircleMembersUpsert(t *testing.T) {
	t.Parallel()

	if len(circleMemberAllColumns) == len(circleMemberPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLCircleMemberUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := CircleMember{}
	if err = randomize.Struct(seed, &o, circleMemberDBTypes, false); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CircleMember: %s", err)
	}

	count, err := CircleMembers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, circleMemberDBTypes, false, circleMemberPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CircleMember struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CircleMember: %s", err)
	}

	count, err = CircleMembers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}