// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/19 18:51:50                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base                                                                                                        *
// * File: base.go                                                                                                     *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base

import "unsafe"

const (
	MinUseRune byte = ' '
	MaxUseRune byte = '~'
)

type IEncoding interface {
	Encode(src []byte) ([]byte, error)
	Decode(src []byte) ([]byte, error)
}

func StringToBytes(src string) []byte {
	return unsafe.Slice(unsafe.StringData(src), len(src))
}

func BytesToString(src []byte) string {
	return unsafe.String(&src[0], len(src))
}

func HasRepeatElem[T comparable](array []T) bool {
	mp := make(map[T]struct{})
	for _, v := range array {
		mp[v] = struct{}{}
	}
	return len(mp) == len(array)
}

func IsIllegalCharacter(c rune) bool {
	return !(c >= rune(MinUseRune) && c <= rune(MaxUseRune))
}
