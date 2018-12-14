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
	expectedCutCollection := &meta.Collection{OpenParenthesisComment, OpenParenthesisToken}

	// action

	actualCutCollection := actualCollection.Cut(
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

	if !reflect.DeepEqual(expectedCutCollection, actualCutCollection) {
		diff := pretty.Compare(expectedCutCollection, actualCutCollection)

		if diff != "" {
			t.Errorf("\nexpected and actual Cut collections are not equal\ndiff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("\nexpected and actual Cut collections are not equal\n")
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
	expectedCutCollection := &meta.Collection{whiteSpace, OpenParenthesisComment}

	// action

	actualCutCollection := actualCollection.Cut(meta.OrFilter(
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

	if !reflect.DeepEqual(expectedCutCollection, actualCutCollection) {
		diff := pretty.Compare(expectedCutCollection, actualCutCollection)

		if diff != "" {
			t.Errorf("\nexpected and actual Cut collections are not equal\ndiff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("\nexpected and actual Cut collections are not equal\n")
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
	expectedCutCollection := &meta.Collection{
		OpenParenthesisComment,
	}

	// action

	actualCutCollection := actualCollection.Cut(meta.AndFilter(
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

	if !reflect.DeepEqual(expectedCutCollection, actualCutCollection) {
		diff := pretty.Compare(expectedCutCollection, actualCutCollection)

		if diff != "" {
			t.Errorf("\nexpected and actual Cut collections are not equal\ndiff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("\nexpected and actual Cut collections are not equal\n")
		}
	}
}

func TestCollectionCutByValue(t *testing.T) {

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
		OpenParenthesisComment,
	}
	expectedCutCollection := &meta.Collection{
		OpenParenthesisToken,
		CloseParenthesisToken,
	}

	// action

	actualCutCollection := actualCollection.Cut(meta.ValueFilter("(", ")"))

	// check

	if !reflect.DeepEqual(expectedCollection, actualCollection) {
		diff := pretty.Compare(expectedCollection, actualCollection)

		if diff != "" {
			t.Errorf("\nexpected and actual collections are not equal\ndiff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("\nexpected and actual collections are not equal\n")
		}
	}

	if !reflect.DeepEqual(expectedCutCollection, actualCutCollection) {
		diff := pretty.Compare(expectedCutCollection, actualCutCollection)

		if diff != "" {
			t.Errorf("\nexpected and actual Cut collections are not equal\ndiff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("\nexpected and actual Cut collections are not equal\n")
		}
	}
}

func TestCollectionCutUntilFirstToken(t *testing.T) {

	// prepare

	whiteSpace := &meta.Data{
		Type:      meta.WhiteSpaceType,
		Value:     "\n",
		TokenName: meta.NodeStart,
	}
	OpenParenthesisComment := &meta.Data{
		Type:      meta.CommentType,
		Value:     "// some comment",
		TokenName: meta.NodeStart,
	}
	OpenParenthesisToken := &meta.Data{
		Type:      meta.TokenType,
		Value:     "(",
		TokenName: meta.OpenParenthesisToken,
	}
	CloseParenthesisComment := &meta.Data{
		Type:      meta.WhiteSpaceType,
		Value:     "// some comment",
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
		CloseParenthesisComment,
		CloseParenthesisToken,
	}

	expectedCutCollection := &meta.Collection{
		whiteSpace,
		OpenParenthesisComment,
	}
	expectedCollection := meta.Collection{
		OpenParenthesisToken,
		CloseParenthesisComment, // must not be cut
		CloseParenthesisToken,
	}

	// action

	actualCutCollection := actualCollection.Cut(
		meta.StopOnFailureFilter(
			meta.NotFilter(
				meta.TypeFilter(meta.TokenType),
			),
		),
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

	if !reflect.DeepEqual(expectedCutCollection, actualCutCollection) {
		diff := pretty.Compare(expectedCutCollection, actualCutCollection)

		if diff != "" {
			t.Errorf("\nexpected and actual Cut collections are not equal\ndiff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("\nexpected and actual Cut collections are not equal\n")
		}
	}
}

func TestCollectionFindByType(t *testing.T) {

	// prepare

	whiteSpace := &meta.Data{
		Type:      meta.WhiteSpaceType,
		Value:     "\n",
		TokenName: meta.NodeStart,
	}
	OpenParenthesisComment := &meta.Data{
		Type:      meta.CommentType,
		Value:     "// some comment",
		TokenName: meta.NodeStart,
	}
	OpenParenthesisToken := &meta.Data{
		Type:      meta.TokenType,
		Value:     "(",
		TokenName: meta.OpenParenthesisToken,
	}
	CloseParenthesisComment := &meta.Data{
		Type:      meta.WhiteSpaceType,
		Value:     "// some comment",
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
		CloseParenthesisComment,
		CloseParenthesisToken,
	}

	expectedCollection := meta.Collection{
		whiteSpace,
		OpenParenthesisComment,
		OpenParenthesisToken,
		CloseParenthesisComment,
		CloseParenthesisToken,
	}

	expectedFoundCollection := meta.Collection{
		OpenParenthesisToken,
		CloseParenthesisToken,
	}

	// action

	actualFoundCollection := actualCollection.FindBy(meta.TypeFilter(meta.TokenType))

	// check

	if !reflect.DeepEqual(expectedCollection, actualCollection) {
		diff := pretty.Compare(expectedCollection, actualCollection)

		if diff != "" {
			t.Errorf("\nexpected and actual collections are not equal\ndiff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("\nexpected and actual collections are not equal\n")
		}
	}

	if !reflect.DeepEqual(expectedFoundCollection, actualFoundCollection) {
		diff := pretty.Compare(expectedFoundCollection, actualFoundCollection)

		if diff != "" {
			t.Errorf("\nexpected and actual Cut collections are not equal\ndiff: (-expected +actual)\n%s", diff)
		} else {
			t.Errorf("\nexpected and actual Cut collections are not equal\n")
		}
	}
}
