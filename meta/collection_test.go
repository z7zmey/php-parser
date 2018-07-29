package meta_test

import (
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/z7zmey/php-parser/meta"
)

func TestCollectionSetTokenName(t *testing.T) {

	// prepare

	whiteSpace := &meta.Data{
		Type:  meta.WhiteSpaceType,
		Value: "\n",
	}
	comment := &meta.Data{
		Type:  meta.CommentType,
		Value: "// some comment",
	}

	baseCollection := meta.Collection{whiteSpace, comment}

	// action

	baseCollection.SetTokenName(meta.OpenParenthesisToken)

	// check

	for _, m := range baseCollection {
		if m.TokenName != meta.OpenParenthesisToken {
			t.Error("The TokenName must be set for all Meta objects")
		}
	}
}

func TestCollectionPush(t *testing.T) {

	// prepare

	whiteSpace := &meta.Data{
		Type:  meta.WhiteSpaceType,
		Value: "\n",
	}
	comment := &meta.Data{
		Type:  meta.CommentType,
		Value: "// some comment",
	}

	actualCollection := meta.Collection{whiteSpace}
	expectedCollection := meta.Collection{whiteSpace, comment}

	// action

	actualCollection.Push(comment)

	// check

	if !reflect.DeepEqual(expectedCollection, actualCollection) {
		diff := pretty.Compare(expectedCollection, actualCollection)

		if diff != "" {
			t.Errorf("\nexpected and actual collections are not equal\ndiff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("\nexpected and actual collections are not equal\n")
		}
	}
}

func TestCollectionUnshift(t *testing.T) {

	// prepare

	whiteSpace := &meta.Data{
		Type:  meta.WhiteSpaceType,
		Value: "\n",
	}
	comment := &meta.Data{
		Type:  meta.CommentType,
		Value: "// some comment",
	}

	actualCollection := meta.Collection{comment}
	expectedCollection := meta.Collection{whiteSpace, comment}

	// action

	actualCollection.Unshift(whiteSpace)

	// check

	if !reflect.DeepEqual(expectedCollection, actualCollection) {
		diff := pretty.Compare(expectedCollection, actualCollection)

		if diff != "" {
			t.Errorf("\nexpected and actual collections are not equal\ndiff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("\nexpected and actual collections are not equal\n")
		}
	}
}

func TestCollectionAppendTo(t *testing.T) {

	// prepare

	whiteSpace := &meta.Data{
		Type:  meta.WhiteSpaceType,
		Value: "\n",
	}
	comment := &meta.Data{
		Type:  meta.CommentType,
		Value: "// some comment",
	}

	actualCollection := meta.Collection{whiteSpace}
	expectedCollection := meta.Collection{whiteSpace, comment}

	baseCollection := meta.Collection{comment}

	// action

	baseCollection.AppendTo(&actualCollection)

	// check

	if !reflect.DeepEqual(expectedCollection, actualCollection) {
		diff := pretty.Compare(expectedCollection, actualCollection)

		if diff != "" {
			t.Errorf("\nexpected and actual collections are not equal\ndiff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("\nexpected and actual collections are not equal\n")
		}
	}
}

func TestEmptyCollectionAppendTo(t *testing.T) {

	// prepare

	whiteSpace := &meta.Data{
		Type:  meta.WhiteSpaceType,
		Value: "\n",
	}

	actualCollection := meta.Collection{whiteSpace}
	expectedCollection := meta.Collection{whiteSpace}

	var baseCollection meta.Collection = nil

	// action

	baseCollection.AppendTo(&actualCollection)

	// check

	if !reflect.DeepEqual(expectedCollection, actualCollection) {
		diff := pretty.Compare(expectedCollection, actualCollection)

		if diff != "" {
			t.Errorf("\nexpected and actual collections are not equal\ndiff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("\nexpected and actual collections are not equal\n")
		}
	}
}

func TestCollectionPrependTo(t *testing.T) {

	// prepare

	whiteSpace := &meta.Data{
		Type:  meta.WhiteSpaceType,
		Value: "\n",
	}
	comment := &meta.Data{
		Type:  meta.CommentType,
		Value: "// some comment",
	}

	actualCollection := meta.Collection{comment}
	expectedCollection := meta.Collection{whiteSpace, comment}

	baseCollection := meta.Collection{whiteSpace}

	// action

	baseCollection.PrependTo(&actualCollection)

	// check

	if !reflect.DeepEqual(expectedCollection, actualCollection) {
		diff := pretty.Compare(expectedCollection, actualCollection)

		if diff != "" {
			t.Errorf("\nexpected and actual collections are not equal\ndiff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("\nexpected and actual collections are not equal\n")
		}
	}
}

func TestEmptyCollectionPrependTo(t *testing.T) {

	// prepare

	comment := &meta.Data{
		Type:  meta.CommentType,
		Value: "// some comment",
	}

	actualCollection := meta.Collection{comment}
	expectedCollection := meta.Collection{comment}

	baseCollection := meta.Collection{}

	// action

	baseCollection.PrependTo(&actualCollection)

	// check

	if !reflect.DeepEqual(expectedCollection, actualCollection) {
		diff := pretty.Compare(expectedCollection, actualCollection)

		if diff != "" {
			t.Errorf("\nexpected and actual collections are not equal\ndiff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("\nexpected and actual collections are not equal\n")
		}
	}
}

func TestCollectionCutByTokenName(t *testing.T) {

	// prepare

	whiteSpace := &meta.Data{
		Type:  meta.WhiteSpaceType,
		Value: "\n",
	}
	OpenParenthesisComment := &meta.Data{
		Type:      meta.CommentType,
		Value:     "// some comment",
		TokenName: meta.OpenParenthesisToken,
	}
	OpenParenthesisToken := &meta.Data{
		Type:      meta.TokenType,
		Value:     "(",
		TokenName: meta.OpenParenthesisToken,
	}
	CloseParenthesisToken := &meta.Data{
		Type:      meta.TokenType,
		Value:     ")",
		TokenName: meta.CloseParenthesisToken,
	}

	actualCollection := meta.Collection{whiteSpace, OpenParenthesisComment, OpenParenthesisToken, CloseParenthesisToken}

	expectedCollection := meta.Collection{whiteSpace, CloseParenthesisToken}
	expectedCuttedCollection := &meta.Collection{OpenParenthesisComment, OpenParenthesisToken}

	// action

	actualCuttedCollection := actualCollection.Cut(
		meta.TokenNameFilter(meta.OpenParenthesisToken),
	)

	// check

	if !reflect.DeepEqual(expectedCollection, actualCollection) {
		diff := pretty.Compare(expectedCollection, actualCollection)

		if diff != "" {
			t.Errorf("\nexpected and actual collections are not equal\ndiff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("\nexpected and actual collections are not equal\n")
		}
	}

	if !reflect.DeepEqual(expectedCuttedCollection, actualCuttedCollection) {
		diff := pretty.Compare(expectedCuttedCollection, actualCuttedCollection)

		if diff != "" {
			t.Errorf("\nexpected and actual cutted collections are not equal\ndiff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("\nexpected and actual cutted collections are not equal\n")
		}
	}
}

func TestCollectionCutByTokenTypes(t *testing.T) {

	// prepare

	whiteSpace := &meta.Data{
		Type:  meta.WhiteSpaceType,
		Value: "\n",
	}
	OpenParenthesisComment := &meta.Data{
		Type:      meta.CommentType,
		Value:     "// some comment",
		TokenName: meta.OpenParenthesisToken,
	}
	OpenParenthesisToken := &meta.Data{
		Type:      meta.TokenType,
		Value:     "(",
		TokenName: meta.OpenParenthesisToken,
	}
	CloseParenthesisToken := &meta.Data{
		Type:      meta.TokenType,
		Value:     ")",
		TokenName: meta.CloseParenthesisToken,
	}

	actualCollection := meta.Collection{whiteSpace, OpenParenthesisComment, OpenParenthesisToken, CloseParenthesisToken}

	expectedCollection := meta.Collection{OpenParenthesisToken, CloseParenthesisToken}
	expectedCuttedCollection := &meta.Collection{whiteSpace, OpenParenthesisComment}

	// action

	actualCuttedCollection := actualCollection.Cut(meta.OrFilter(
		meta.TypeFilter(meta.CommentType),
		meta.TypeFilter(meta.WhiteSpaceType)),
	)

	// check

	if !reflect.DeepEqual(expectedCollection, actualCollection) {
		diff := pretty.Compare(expectedCollection, actualCollection)

		if diff != "" {
			t.Errorf("\nexpected and actual collections are not equal\ndiff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("\nexpected and actual collections are not equal\n")
		}
	}

	if !reflect.DeepEqual(expectedCuttedCollection, actualCuttedCollection) {
		diff := pretty.Compare(expectedCuttedCollection, actualCuttedCollection)

		if diff != "" {
			t.Errorf("\nexpected and actual cutted collections are not equal\ndiff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("\nexpected and actual cutted collections are not equal\n")
		}
	}
}

func TestCollectionCutByTokenNameButNotType(t *testing.T) {

	// prepare

	whiteSpace := &meta.Data{
		Type:  meta.WhiteSpaceType,
		Value: "\n",
	}
	OpenParenthesisComment := &meta.Data{
		Type:      meta.CommentType,
		Value:     "// some comment",
		TokenName: meta.OpenParenthesisToken,
	}
	OpenParenthesisToken := &meta.Data{
		Type:      meta.TokenType,
		Value:     "(",
		TokenName: meta.OpenParenthesisToken,
	}
	CloseParenthesisToken := &meta.Data{
		Type:      meta.TokenType,
		Value:     ")",
		TokenName: meta.CloseParenthesisToken,
	}

	actualCollection := meta.Collection{
		whiteSpace,
		OpenParenthesisComment,
		OpenParenthesisToken,
		CloseParenthesisToken,
	}

	expectedCollection := meta.Collection{
		whiteSpace,
		OpenParenthesisToken,
		CloseParenthesisToken,
	}
	expectedCuttedCollection := &meta.Collection{
		OpenParenthesisComment,
	}

	// action

	actualCuttedCollection := actualCollection.Cut(meta.AndFilter(
		meta.TokenNameFilter(meta.OpenParenthesisToken),
		meta.NotFilter(meta.TypeFilter(meta.TokenType))),
	)

	// check

	if !reflect.DeepEqual(expectedCollection, actualCollection) {
		diff := pretty.Compare(expectedCollection, actualCollection)

		if diff != "" {
			t.Errorf("\nexpected and actual collections are not equal\ndiff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("\nexpected and actual collections are not equal\n")
		}
	}

	if !reflect.DeepEqual(expectedCuttedCollection, actualCuttedCollection) {
		diff := pretty.Compare(expectedCuttedCollection, actualCuttedCollection)

		if diff != "" {
			t.Errorf("\nexpected and actual cutted collections are not equal\ndiff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("\nexpected and actual cutted collections are not equal\n")
		}
	}
}
