package model

import (
	"fmt"
	"strconv"
	"strings"
)

type Container struct {
	category    string
	subCategory string
	variant     *int
	label       *string
	isModified  bool
}

// Chirper.POLICIES[Survivalist Movement]:1 (with label and variant)
// Assets.UPGRADE_NAME[PostSortingFacility01 Mail Storage Extension] (with label)
// Assets.CITIZEN_SURNAME_HOUSEHOLD:100 (with variant)
// Editor.SELECT_DIRECTORY (only category and sub)

func (container *Container) Parse(val string) error {
	ct1 := strings.SplitN(val, ".", 2)
	container.category = ct1[0] // mandatory

	ct2 := strings.SplitN(ct1[1], ":", 2)
	if len(ct2) == 2 {
		v, err := strconv.Atoi(ct2[1])
		if err != nil {
			return err
		}

		container.variant = &v
	}

	ct3 := strings.SplitN(ct2[0], "[", 2)
	container.subCategory = ct3[0]

	if len(ct3) == 2 {
		v := strings.TrimSuffix(ct3[1], "]")
		container.label = &v
	}

	return nil
}

// TODO
func (container *Container) String() string {
	var (
		variant string
		label   string
	)

	if container.variant != nil {
		variant = fmt.Sprintf(":%d", *container.variant)
	}

	if container.label != nil {
		label = fmt.Sprintf("[%s]", *container.label)
	}

	return fmt.Sprintf("%s.%s%s%s", container.category, container.subCategory, label, variant)
}

func (container *Container) Category() string {
	return container.category
}

func (container *Container) SubCategory() string {
	return container.subCategory
}

func (container *Container) Variant() *int {
	return container.variant
}

func (container *Container) Label() *string {
	return container.label
}

func (container *Container) IsModified() bool {
	return container.isModified
}
