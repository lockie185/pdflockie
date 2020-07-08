/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package textencoding

const (
	baseMacExpert = "MacExpertEncoding"
	baseMacRoman  = "MacRomanEncoding"
)

var (
	macExpert = newSimpleMapping(baseMacExpert, macExpertCharToRune)
	macRoman  = newSimpleMapping(baseMacRoman, macRomanCharToRune)
)

func init() {
	RegisterSimpleEncoding(baseMacExpert, NewMacExpertEncoder)
	RegisterSimpleEncoding(baseMacRoman, NewMacRomanEncoder)
}

// NewMacExpertEncoder returns a SimpleEncoder that implements MacExpertEncoding.
func NewMacExpertEncoder() SimpleEncoder {
	return macExpert.NewEncoder()
}

// NewMacRomanEncoder returns a SimpleEncoder that implements MacRomanEncoding.
func NewMacRomanEncoder() SimpleEncoder {
	return macRoman.NewEncoder()
}

var macExpertCharToRune = map[byte]rune{ // 165 entries
	0x20: ' ', 0x21: '\uf721', 0x22: '\uf6f8', 0x23: '\uf7a2', 0x24: '\uf724',
	0x25: '\uf6e4', 0x26: '\uf726', 0x27: '\uf7b4', 0x28: '⁽', 0x29: '⁾',
	0x2a: '‥', 0x2b: '․', 0x2c: ',', 0x2d: '-', 0x2e: '.',
	0x2f: '⁄', 0x30: '\uf730', 0x31: '\uf731', 0x32: '\uf732', 0x33: '\uf733',
	0x34: '\uf734', 0x35: '\uf735', 0x36: '\uf736', 0x37: '\uf737', 0x38: '\uf738',
	0x39: '\uf739', 0x3a: ':', 0x3b: ';', 0x3d: '\uf6de', 0x3f: '\uf73f',
	0x44: '\uf7f0', 0x47: '¼', 0x48: '½', 0x49: '¾', 0x4a: '⅛',
	0x4b: '⅜', 0x4c: '⅝', 0x4d: '⅞', 0x4e: '⅓', 0x4f: '⅔',
	0x56: 'ﬀ', 0x57: 'ﬁ', 0x58: 'ﬂ', 0x59: 'ﬃ', 0x5a: 'ﬄ',
	0x5b: '₍', 0x5d: '₎', 0x5e: '\uf6f6', 0x5f: '\uf6e5', 0x60: '\uf760',
	0x61: '\uf761', 0x62: '\uf762', 0x63: '\uf763', 0x64: '\uf764', 0x65: '\uf765',
	0x66: '\uf766', 0x67: '\uf767', 0x68: '\uf768', 0x69: '\uf769', 0x6a: '\uf76a',
	0x6b: '\uf76b', 0x6c: '\uf76c', 0x6d: '\uf76d', 0x6e: '\uf76e', 0x6f: '\uf76f',
	0x70: '\uf770', 0x71: '\uf771', 0x72: '\uf772', 0x73: '\uf773', 0x74: '\uf774',
	0x75: '\uf775', 0x76: '\uf776', 0x77: '\uf777', 0x78: '\uf778', 0x79: '\uf779',
	0x7a: '\uf77a', 0x7b: '₡', 0x7c: '\uf6dc', 0x7d: '\uf6dd', 0x7e: '\uf6fe',
	0x81: '\uf6e9', 0x82: '\uf6e0', 0x87: '\uf7e1', 0x88: '\uf7e0', 0x89: '\uf7e2',
	0x8a: '\uf7e4', 0x8b: '\uf7e3', 0x8c: '\uf7e5', 0x8d: '\uf7e7', 0x8e: '\uf7e9',
	0x8f: '\uf7e8', 0x90: '\uf7ea', 0x91: '\uf7eb', 0x92: '\uf7ed', 0x93: '\uf7ec',
	0x94: '\uf7ee', 0x95: '\uf7ef', 0x96: '\uf7f1', 0x97: '\uf7f3', 0x98: '\uf7f2',
	0x99: '\uf7f4', 0x9a: '\uf7f6', 0x9b: '\uf7f5', 0x9c: '\uf7fa', 0x9d: '\uf7f9',
	0x9e: '\uf7fb', 0x9f: '\uf7fc', 0xa1: '⁸', 0xa2: '₄', 0xa3: '₃',
	0xa4: '₆', 0xa5: '₈', 0xa6: '₇', 0xa7: '\uf6fd', 0xa9: '\uf6df',
	0xaa: '₂', 0xac: '\uf7a8', 0xae: '\uf6f5', 0xaf: '\uf6f0', 0xb0: '₅',
	0xb2: '\uf6e1', 0xb3: '\uf6e7', 0xb4: '\uf7fd', 0xb6: '\uf6e3', 0xb9: '\uf7fe',
	0xbb: '₉', 0xbc: '₀', 0xbd: '\uf6ff', 0xbe: '\uf7e6', 0xbf: '\uf7f8',
	0xc0: '\uf7bf', 0xc1: '₁', 0xc2: '\uf6f9', 0xc9: '\uf7b8', 0xcf: '\uf6fa',
	0xd0: '‒', 0xd1: '\uf6e6', 0xd6: '\uf7a1', 0xd8: '\uf7ff', 0xda: '¹',
	0xdb: '²', 0xdc: '³', 0xdd: '⁴', 0xde: '⁵', 0xdf: '⁶',
	0xe0: '⁷', 0xe1: '⁹', 0xe2: '⁰', 0xe4: '\uf6ec', 0xe5: '\uf6f1',
	0xe6: '\uf6f3', 0xe9: '\uf6ed', 0xea: '\uf6f2', 0xeb: '\uf6eb', 0xf1: '\uf6ee',
	0xf2: '\uf6fb', 0xf3: '\uf6f4', 0xf4: '\uf7af', 0xf5: '\uf6ea', 0xf6: 'ⁿ',
	0xf7: '\uf6ef', 0xf8: '\uf6e2', 0xf9: '\uf6e8', 0xfa: '\uf6f7', 0xfb: '\uf6fc',
}

var macRomanCharToRune = map[byte]rune{ // 255 entries
	0x1: '\x01', 0x2: '\x02', 0x3: '\x03', 0x4: '\x04', 0x5: '\x05',
	0x6: '\x06', 0x7: '\a', 0x8: '\b', 0x9: '\t', 0xa: '\n',
	0xb: '\v', 0xc: '\f', 0xd: '\r', 0xe: '\x0e', 0xf: '\x0f',
	0x10: '\x10', 0x11: '\x11', 0x12: '\x12', 0x13: '\x13', 0x14: '\x14',
	0x15: '\x15', 0x16: '\x16', 0x17: '\x17', 0x18: '\x18', 0x19: '\x19',
	0x1a: '\x1a', 0x1b: '\x1b', 0x1c: '\x1c', 0x1d: '\x1d', 0x1e: '\x1e',
	0x1f: '\x1f', 0x20: ' ', 0x21: '!', 0x22: '"', 0x23: '#',
	0x24: '$', 0x25: '%', 0x26: '&', 0x27: '\'', 0x28: '(',
	0x29: ')', 0x2a: '*', 0x2b: '+', 0x2c: ',', 0x2d: '-',
	0x2e: '.', 0x2f: '/', 0x30: '0', 0x31: '1', 0x32: '2',
	0x33: '3', 0x34: '4', 0x35: '5', 0x36: '6', 0x37: '7',
	0x38: '8', 0x39: '9', 0x3a: ':', 0x3b: ';', 0x3c: '<',
	0x3d: '=', 0x3e: '>', 0x3f: '?', 0x40: '@', 0x41: 'A',
	0x42: 'B', 0x43: 'C', 0x44: 'D', 0x45: 'E', 0x46: 'F',
	0x47: 'G', 0x48: 'H', 0x49: 'I', 0x4a: 'J', 0x4b: 'K',
	0x4c: 'L', 0x4d: 'M', 0x4e: 'N', 0x4f: 'O', 0x50: 'P',
	0x51: 'Q', 0x52: 'R', 0x53: 'S', 0x54: 'T', 0x55: 'U',
	0x56: 'V', 0x57: 'W', 0x58: 'X', 0x59: 'Y', 0x5a: 'Z',
	0x5b: '[', 0x5c: '\\', 0x5d: ']', 0x5e: '^', 0x5f: '_',
	0x60: '`', 0x61: 'a', 0x62: 'b', 0x63: 'c', 0x64: 'd',
	0x65: 'e', 0x66: 'f', 0x67: 'g', 0x68: 'h', 0x69: 'i',
	0x6a: 'j', 0x6b: 'k', 0x6c: 'l', 0x6d: 'm', 0x6e: 'n',
	0x6f: 'o', 0x70: 'p', 0x71: 'q', 0x72: 'r', 0x73: 's',
	0x74: 't', 0x75: 'u', 0x76: 'v', 0x77: 'w', 0x78: 'x',
	0x79: 'y', 0x7a: 'z', 0x7b: '{', 0x7c: '|', 0x7d: '}',
	0x7e: '~', 0x7f: '\u007f', 0x80: 'Ä', 0x81: 'Å', 0x82: 'Ç',
	0x83: 'É', 0x84: 'Ñ', 0x85: 'Ö', 0x86: 'Ü', 0x87: 'á',
	0x88: 'à', 0x89: 'â', 0x8a: 'ä', 0x8b: 'ã', 0x8c: 'å',
	0x8d: 'ç', 0x8e: 'é', 0x8f: 'è', 0x90: 'ê', 0x91: 'ë',
	0x92: 'í', 0x93: 'ì', 0x94: 'î', 0x95: 'ï', 0x96: 'ñ',
	0x97: 'ó', 0x98: 'ò', 0x99: 'ô', 0x9a: 'ö', 0x9b: 'õ',
	0x9c: 'ú', 0x9d: 'ù', 0x9e: 'û', 0x9f: 'ü', 0xa0: '†',
	0xa1: '°', 0xa2: '¢', 0xa3: '£', 0xa4: '§', 0xa5: '•',
	0xa6: '¶', 0xa7: 'ß', 0xa8: '®', 0xa9: '©', 0xaa: '™',
	0xab: '´', 0xac: '¨', 0xad: '≠', 0xae: 'Æ', 0xaf: 'Ø',
	0xb0: '∞', 0xb1: '±', 0xb2: '≤', 0xb3: '≥', 0xb4: '¥',
	0xb5: 'µ', 0xb6: '∂', 0xb7: '∑', 0xb8: '∏', 0xb9: 'π',
	0xba: '∫', 0xbb: 'ª', 0xbc: 'º', 0xbd: 'Ω', 0xbe: 'æ',
	0xbf: 'ø', 0xc0: '¿', 0xc1: '¡', 0xc2: '¬', 0xc3: '√',
	0xc4: 'ƒ', 0xc5: '≈', 0xc6: '∆', 0xc7: '«', 0xc8: '»',
	0xc9: '…', 0xca: '\u00a0', 0xcb: 'À', 0xcc: 'Ã', 0xcd: 'Õ',
	0xce: 'Œ', 0xcf: 'œ', 0xd0: '–', 0xd1: '—', 0xd2: '“',
	0xd3: '”', 0xd4: '‘', 0xd5: '’', 0xd6: '÷', 0xd7: '◊',
	0xd8: 'ÿ', 0xd9: 'Ÿ', 0xda: '⁄', 0xdb: '€', 0xdc: '‹',
	0xdd: '›', 0xde: 'ﬁ', 0xdf: 'ﬂ', 0xe0: '‡', 0xe1: '·',
	0xe2: '‚', 0xe3: '„', 0xe4: '‰', 0xe5: 'Â', 0xe6: 'Ê',
	0xe7: 'Á', 0xe8: 'Ë', 0xe9: 'È', 0xea: 'Í', 0xeb: 'Î',
	0xec: 'Ï', 0xed: 'Ì', 0xee: 'Ó', 0xef: 'Ô', 0xf0: '\uf8ff',
	0xf1: 'Ò', 0xf2: 'Ú', 0xf3: 'Û', 0xf4: 'Ù', 0xf5: 'ı',
	0xf6: 'ˆ', 0xf7: '˜', 0xf8: '¯', 0xf9: '˘', 0xfa: '˙',
	0xfb: '˚', 0xfc: '¸', 0xfd: '˝', 0xfe: '˛', 0xff: 'ˇ',
}
