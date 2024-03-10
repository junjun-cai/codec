// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/29 21:39:42                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: util                                                                                                        *
// * File: utils.go                                                                                                    *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package util

const (
	MinUseRune byte = ' '
	MaxUseRune byte = '~'
	StdPadding rune = '='
	NotPadding rune = -1
)

func IsIllegalCharacter(c rune) bool {
	if c == NotPadding {
		return true
	}
	return !(c >= rune(MinUseRune) && c <= rune(MaxUseRune))
}
