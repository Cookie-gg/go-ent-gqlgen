// Code generated by ent, DO NOT EDIT.

package ent

import (
	"errors"
	"fmt"
	"go-ent-gqlgen/ent/post"
	"go-ent-gqlgen/ent/predicate"
	"go-ent-gqlgen/ent/user"
	"time"
)

// PostWhereInput represents a where input for filtering Post queries.
type PostWhereInput struct {
	Predicates []predicate.Post  `json:"-"`
	Not        *PostWhereInput   `json:"not,omitempty"`
	Or         []*PostWhereInput `json:"or,omitempty"`
	And        []*PostWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "title" field predicates.
	Title             *string  `json:"title,omitempty"`
	TitleNEQ          *string  `json:"titleNEQ,omitempty"`
	TitleIn           []string `json:"titleIn,omitempty"`
	TitleNotIn        []string `json:"titleNotIn,omitempty"`
	TitleGT           *string  `json:"titleGT,omitempty"`
	TitleGTE          *string  `json:"titleGTE,omitempty"`
	TitleLT           *string  `json:"titleLT,omitempty"`
	TitleLTE          *string  `json:"titleLTE,omitempty"`
	TitleContains     *string  `json:"titleContains,omitempty"`
	TitleHasPrefix    *string  `json:"titleHasPrefix,omitempty"`
	TitleHasSuffix    *string  `json:"titleHasSuffix,omitempty"`
	TitleEqualFold    *string  `json:"titleEqualFold,omitempty"`
	TitleContainsFold *string  `json:"titleContainsFold,omitempty"`

	// "message" field predicates.
	Message             *string  `json:"message,omitempty"`
	MessageNEQ          *string  `json:"messageNEQ,omitempty"`
	MessageIn           []string `json:"messageIn,omitempty"`
	MessageNotIn        []string `json:"messageNotIn,omitempty"`
	MessageGT           *string  `json:"messageGT,omitempty"`
	MessageGTE          *string  `json:"messageGTE,omitempty"`
	MessageLT           *string  `json:"messageLT,omitempty"`
	MessageLTE          *string  `json:"messageLTE,omitempty"`
	MessageContains     *string  `json:"messageContains,omitempty"`
	MessageHasPrefix    *string  `json:"messageHasPrefix,omitempty"`
	MessageHasSuffix    *string  `json:"messageHasSuffix,omitempty"`
	MessageEqualFold    *string  `json:"messageEqualFold,omitempty"`
	MessageContainsFold *string  `json:"messageContainsFold,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *time.Time  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *time.Time  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *time.Time  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *time.Time  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *time.Time  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *time.Time  `json:"updatedAtLTE,omitempty"`

	// "likes" field predicates.
	Likes      *int  `json:"likes,omitempty"`
	LikesNEQ   *int  `json:"likesNEQ,omitempty"`
	LikesIn    []int `json:"likesIn,omitempty"`
	LikesNotIn []int `json:"likesNotIn,omitempty"`
	LikesGT    *int  `json:"likesGT,omitempty"`
	LikesGTE   *int  `json:"likesGTE,omitempty"`
	LikesLT    *int  `json:"likesLT,omitempty"`
	LikesLTE   *int  `json:"likesLTE,omitempty"`

	// "user" edge predicates.
	HasUser     *bool             `json:"hasUser,omitempty"`
	HasUserWith []*UserWhereInput `json:"hasUserWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *PostWhereInput) AddPredicates(predicates ...predicate.Post) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the PostWhereInput filter on the PostQuery builder.
func (i *PostWhereInput) Filter(q *PostQuery) (*PostQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyPostWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyPostWhereInput is returned in case the PostWhereInput is empty.
var ErrEmptyPostWhereInput = errors.New("ent: empty predicate PostWhereInput")

// P returns a predicate for filtering posts.
// An error is returned if the input is empty or invalid.
func (i *PostWhereInput) P() (predicate.Post, error) {
	var predicates []predicate.Post
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, post.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Post, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, post.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Post, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, post.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, post.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, post.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, post.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, post.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, post.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, post.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, post.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, post.IDLTE(*i.IDLTE))
	}
	if i.Title != nil {
		predicates = append(predicates, post.TitleEQ(*i.Title))
	}
	if i.TitleNEQ != nil {
		predicates = append(predicates, post.TitleNEQ(*i.TitleNEQ))
	}
	if len(i.TitleIn) > 0 {
		predicates = append(predicates, post.TitleIn(i.TitleIn...))
	}
	if len(i.TitleNotIn) > 0 {
		predicates = append(predicates, post.TitleNotIn(i.TitleNotIn...))
	}
	if i.TitleGT != nil {
		predicates = append(predicates, post.TitleGT(*i.TitleGT))
	}
	if i.TitleGTE != nil {
		predicates = append(predicates, post.TitleGTE(*i.TitleGTE))
	}
	if i.TitleLT != nil {
		predicates = append(predicates, post.TitleLT(*i.TitleLT))
	}
	if i.TitleLTE != nil {
		predicates = append(predicates, post.TitleLTE(*i.TitleLTE))
	}
	if i.TitleContains != nil {
		predicates = append(predicates, post.TitleContains(*i.TitleContains))
	}
	if i.TitleHasPrefix != nil {
		predicates = append(predicates, post.TitleHasPrefix(*i.TitleHasPrefix))
	}
	if i.TitleHasSuffix != nil {
		predicates = append(predicates, post.TitleHasSuffix(*i.TitleHasSuffix))
	}
	if i.TitleEqualFold != nil {
		predicates = append(predicates, post.TitleEqualFold(*i.TitleEqualFold))
	}
	if i.TitleContainsFold != nil {
		predicates = append(predicates, post.TitleContainsFold(*i.TitleContainsFold))
	}
	if i.Message != nil {
		predicates = append(predicates, post.MessageEQ(*i.Message))
	}
	if i.MessageNEQ != nil {
		predicates = append(predicates, post.MessageNEQ(*i.MessageNEQ))
	}
	if len(i.MessageIn) > 0 {
		predicates = append(predicates, post.MessageIn(i.MessageIn...))
	}
	if len(i.MessageNotIn) > 0 {
		predicates = append(predicates, post.MessageNotIn(i.MessageNotIn...))
	}
	if i.MessageGT != nil {
		predicates = append(predicates, post.MessageGT(*i.MessageGT))
	}
	if i.MessageGTE != nil {
		predicates = append(predicates, post.MessageGTE(*i.MessageGTE))
	}
	if i.MessageLT != nil {
		predicates = append(predicates, post.MessageLT(*i.MessageLT))
	}
	if i.MessageLTE != nil {
		predicates = append(predicates, post.MessageLTE(*i.MessageLTE))
	}
	if i.MessageContains != nil {
		predicates = append(predicates, post.MessageContains(*i.MessageContains))
	}
	if i.MessageHasPrefix != nil {
		predicates = append(predicates, post.MessageHasPrefix(*i.MessageHasPrefix))
	}
	if i.MessageHasSuffix != nil {
		predicates = append(predicates, post.MessageHasSuffix(*i.MessageHasSuffix))
	}
	if i.MessageEqualFold != nil {
		predicates = append(predicates, post.MessageEqualFold(*i.MessageEqualFold))
	}
	if i.MessageContainsFold != nil {
		predicates = append(predicates, post.MessageContainsFold(*i.MessageContainsFold))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, post.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, post.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, post.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, post.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, post.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, post.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, post.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, post.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, post.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, post.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, post.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, post.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, post.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, post.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, post.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, post.UpdatedAtLTE(*i.UpdatedAtLTE))
	}
	if i.Likes != nil {
		predicates = append(predicates, post.LikesEQ(*i.Likes))
	}
	if i.LikesNEQ != nil {
		predicates = append(predicates, post.LikesNEQ(*i.LikesNEQ))
	}
	if len(i.LikesIn) > 0 {
		predicates = append(predicates, post.LikesIn(i.LikesIn...))
	}
	if len(i.LikesNotIn) > 0 {
		predicates = append(predicates, post.LikesNotIn(i.LikesNotIn...))
	}
	if i.LikesGT != nil {
		predicates = append(predicates, post.LikesGT(*i.LikesGT))
	}
	if i.LikesGTE != nil {
		predicates = append(predicates, post.LikesGTE(*i.LikesGTE))
	}
	if i.LikesLT != nil {
		predicates = append(predicates, post.LikesLT(*i.LikesLT))
	}
	if i.LikesLTE != nil {
		predicates = append(predicates, post.LikesLTE(*i.LikesLTE))
	}

	if i.HasUser != nil {
		p := post.HasUser()
		if !*i.HasUser {
			p = post.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasUserWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasUserWith))
		for _, w := range i.HasUserWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasUserWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, post.HasUserWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyPostWhereInput
	case 1:
		return predicates[0], nil
	default:
		return post.And(predicates...), nil
	}
}

// UserWhereInput represents a where input for filtering User queries.
type UserWhereInput struct {
	Predicates []predicate.User  `json:"-"`
	Not        *UserWhereInput   `json:"not,omitempty"`
	Or         []*UserWhereInput `json:"or,omitempty"`
	And        []*UserWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *time.Time  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *time.Time  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *time.Time  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *time.Time  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *time.Time  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *time.Time  `json:"updatedAtLTE,omitempty"`

	// "posts" edge predicates.
	HasPosts     *bool             `json:"hasPosts,omitempty"`
	HasPostsWith []*PostWhereInput `json:"hasPostsWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *UserWhereInput) AddPredicates(predicates ...predicate.User) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the UserWhereInput filter on the UserQuery builder.
func (i *UserWhereInput) Filter(q *UserQuery) (*UserQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyUserWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyUserWhereInput is returned in case the UserWhereInput is empty.
var ErrEmptyUserWhereInput = errors.New("ent: empty predicate UserWhereInput")

// P returns a predicate for filtering users.
// An error is returned if the input is empty or invalid.
func (i *UserWhereInput) P() (predicate.User, error) {
	var predicates []predicate.User
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, user.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.User, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, user.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.User, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, user.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, user.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, user.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, user.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, user.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, user.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, user.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, user.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, user.IDLTE(*i.IDLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, user.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, user.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, user.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, user.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, user.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, user.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, user.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, user.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, user.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, user.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, user.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, user.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, user.NameContainsFold(*i.NameContainsFold))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, user.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, user.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, user.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, user.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, user.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, user.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, user.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, user.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, user.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, user.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, user.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, user.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, user.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, user.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, user.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, user.UpdatedAtLTE(*i.UpdatedAtLTE))
	}

	if i.HasPosts != nil {
		p := user.HasPosts()
		if !*i.HasPosts {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasPostsWith) > 0 {
		with := make([]predicate.Post, 0, len(i.HasPostsWith))
		for _, w := range i.HasPostsWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasPostsWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasPostsWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyUserWhereInput
	case 1:
		return predicates[0], nil
	default:
		return user.And(predicates...), nil
	}
}
