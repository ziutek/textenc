// Package textenc provides functions to decode non-UTF8 text strings
package textenc

var iso6937 = []rune("" +
	"\u00a0¡¢£ ¥ §¤‘“«←↑→↓" +
	"°±²³×µ¶·÷’”»¼½¾¿" +
	"                " +
	"―¹®©™♪¬¦    ⅛⅜⅝⅞" +
	"ΩÆĐªĦ ĲĿŁØŒºÞŦŊŉ" +
	"ĸæđðħıĳŀłøœßþŧŋ\u00AD")

var iso8859_1 = []rune("" +
	"\u00a0¡¢£¤¥¦§¨©ª«¬\u00AD®¯" +
	"°±²³´µ¶·¸¹º»¼½¾¿" +
	"ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏ" +
	"ÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞß" +
	"àáâãäåæçèéêëìíîï" +
	"ðñòóôõö÷øùúûüýþÿ")

var iso8859_2 = []rune("" +
	"\u00a0Ą˘Ł¤ĽŚ§¨ŠŞŤŹ\u00ADŽŻ" +
	"°ą˛ł´ľśˇ¸šşťź˝žż" +
	"ŔÁÂĂÄĹĆÇČÉĘËĚÍÎĎ" +
	"ĐŃŇÓÔŐÖ×ŘŮÚŰÜÝŢß" +
	"ŕáâăäĺćçčéęëěíîď" +
	"đńňóôőö÷řůúűüýţ˙")

var iso8859_3 = []rune("" +
	"\u00a0Ħ˘£¤ Ĥ§¨İŞĞĴ\u00AD Ż" +
	"°ħ²³´µĥ·¸ışğĵ½ ż" +
	"ÀÁÂ ÄĊĈÇÈÉÊËÌÍÎÏ" +
	" ÑÒÓÔĠÖ×ĜÙÚÛÜŬŜß" +
	"àáâ äċĉçèéêëìíîï" +
	" ñòóôġö÷ĝùúûüŭŝ˙")

var iso8859_4 = []rune("" +
	"\u00a0ĄĸŖ¤ĨĻ§¨ŠĒĢŦ\u00ADŽ¯" +
	"°ą˛ŗ´ĩļˇ¸šēģŧŊžŋ" +
	"ĀÁÂÃÄÅÆĮČÉĘËĖÍÎĪ" +
	"ĐŅŌĶÔÕÖ×ØŲÚÛÜŨŪß" +
	"āáâãäåæįčéęëėíîī" +
	"đņōķôõö÷øųúûüũū˙")

var iso8859_5 = []rune("" +
	"\u00a0ЁЂЃЄЅІЇЈЉЊЋЌ\u00ADЎЏ" +
	"АБВГДЕЖЗИЙКЛМНОП" +
	"РСТУФХЦЧШЩЪЫЬЭЮЯ" +
	"абвгдежзийклмноп" +
	"рстуфхцчшщъыьэюя" +
	"№ёђѓєѕіїјљњћќ§ўџ")

var iso8859_6 = []rune("" +
	"\u00a0   ¤       ،\u00AD  " +
	"           ؛   ؟" +
	" ءآأؤإئابةتثجحخد" +
	"ذرزسشصضطظعغ     " +
	"ـفقكلمنهوىي     " +
	"                ")

var iso8859_7 = []rune("" +
	"\u00a0ʽʼ£  ¦§¨© «¬\u00AD ―" +
	"°±²³΄΅Ά·ΈΉΊ»Ό½ΎΏ" +
	"ΐΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟ" +
	"ΠΡ ΣΤΥΦΧΨΩΪΫάέήί" +
	"ΰαβγδεζηθικλμνξο" +
	"πρςστυφχψωϊϋόύώ ")

var iso8859_8 = []rune("" +
	"\u00a0 ¢£¤¥¦§¨©×«¬\u00AD®‾" +
	"°±²³´µ¶·¸¹÷»¼½¾ " +
	"                " +
	"               ‗" +
	"אבגדהוזחטיךכלםמן" +
	"נסעףפץצקרשת     ")

var iso8859_9 = []rune("" +
	"\u00a0¡¢£¤¥¦§¨©ª«¬\u00AD®¯" +
	"°±²³´µ¶·¸¹º»¼½¾¿" +
	"ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏ" +
	"ĞÑÒÓÔÕÖ×ØÙÚÛÜİŞß" +
	"àáâãäåæçèéêëìíîï" +
	"ğñòóôõö÷øùúûüışÿ")

var iso8859_10 = []rune("" +
	"\u00a0ĄĒĢĪĨĶ§ĻĐŠŦŽ\u00ADŪŊ" +
	"°ąēģīĩķ·ļđšŧž―ūŋ" +
	"ĀÁÂÃÄÅÆĮČÉĘËĖÍÎÏ" +
	"ÐŅŌÓÔÕÖŨØŲÚÛÜÝÞß" +
	"āáâãäåæįčéęëėíîï" +
	"ðņōóôõöũøųúûüýþĸ")

var iso8859_11 = []rune("" +
	" กขฃคฅฆงจฉชซฌญฎฏ" +
	"ฐฑฒณดตถทธนบปผฝพฟ" +
	"ภมยรฤลฦวศษสหฬอฮฯ" +
	"ะ า            ฿" +
	"เแโใไๅๆ        ๏" +
	"๐๑๒๓๔๕๖๗๘๙๚๛    ")

var iso8859_13 = []rune("" +
	"\u00a0”¢£¤„¦§Ø©Ŗ«¬\u00AD®Æ" +
	"°±²³“µ¶·ø¹ŗ»¼½¾æ" +
	"ĄĮĀĆÄÅĘĒČÉŹĖĢĶĪĻ" +
	"ŠŃŅÓŌÕÖ×ŲŁŚŪÜŻŽß" +
	"ąįāćäåęēčéźėģķīļ" +
	"šńņóōõö÷ųłśūüżž’")

var iso8859_14 = []rune("" +
	"\u00a0Ḃḃ£ĊċḊ§Ẁ©ẂḋỲ\u00AD®Ÿ" +
	"ḞḟĠġṀṁ¶ṖẁṗẃṠỳẄẅṡ" +
	"ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏ" +
	"ŴÑÒÓÔÕÖṪØÙÚÛÜÝŶß" +
	"àáâãäåæçèéêëìíîï" +
	"ŵñòóôõöṫøùúûüýŷÿ")

var iso8859_15 = []rune("" +
	"\u00a0¡¢£€¥Š§š©ª«¬\u00ad®¯" +
	"°±²³Žµ¶·ž¹º»ŒœŸ¿" +
	"ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏ" +
	"ÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞß" +
	"àáâãäåæçèéêëìíîï" +
	"ðñòóôõö÷øùúûüýþÿ")

var iso6937twoBytes = []map[byte]rune{
	{ // Grave 0xC1
		'A': 'À', 'E': 'È', 'I': 'Ì', 'O': 'Ò', 'U': 'Ù',
		'a': 'à', 'e': 'è', 'i': 'ì', 'o': 'ò', 'u': 'ù',
	}, { // Acute 0xC2
		'A': 'Á', 'C': 'Ć', 'E': 'É', 'I': 'Í', 'L': 'Ĺ', 'N': 'Ń',
		'O': 'Ó', 'R': 'Ŕ', 'S': 'Ś', 'U': 'Ú', 'Y': 'Ý', 'Z': 'Ź',
		'a': 'á', 'c': 'ć', 'e': 'é', 'g': 'ģ', 'i': 'í', 'l': 'ĺ',
		'n': 'ń', 'o': 'ó', 'r': 'ŕ', 's': 'ś', 'u': 'ú', 'y': 'ý', 'z': 'ź',
	}, { // Circumflex 0xC3
		'A': 'Â', 'C': 'Ĉ', 'E': 'Ê', 'G': 'Ĝ', 'H': 'Ĥ', 'I': 'Î',
		'J': 'Ĵ', 'O': 'Ô', 'S': 'Ŝ', 'U': 'Û', 'W': 'Ŵ', 'Y': 'Ŷ',
		'a': 'â', 'c': 'ĉ', 'e': 'ê', 'g': 'ĝ', 'h': 'ĥ', 'i': 'î',
		'j': 'ĵ', 'o': 'ô', 's': 'ŝ', 'u': 'û', 'w': 'ŵ', 'y': 'ŷ',
	}, { // Tilde 0xC4
		'A': 'Ã', 'I': 'Ĩ', 'N': 'Ñ', 'O': 'Õ', 'U': 'Ũ',
		'a': 'ã', 'i': 'ĩ', 'n': 'ñ', 'o': 'õ', 'u': 'ũ',
	}, { // Macron 0xC5
		'A': 'Ā', 'E': 'Ē', 'I': 'Ī', 'O': 'Ō', 'U': 'Ū',
		'a': 'ā', 'e': 'ē', 'i': 'ī', 'o': 'ō', 'u': 'ū',
	}, { //Breve 0xC6
		'A': 'Ă', 'G': 'Ğ', 'U': 'Ŭ', 'a': 'ă', 'g': 'ğ', 'u': 'ŭ',
	}, { // Dot 0xC7
		'C': 'Ċ', 'E': 'Ė', 'G': 'Ġ', 'I': 'İ', 'Z': 'Ż',
		'c': 'ċ', 'e': 'ė', 'g': 'ġ', 'z': 'ż',
	}, { // Umlaut 0xC8
		'A': 'Ä', 'E': 'Ë', 'I': 'Ï', 'O': 'Ö', 'U': 'Ü', 'Y': 'Ÿ',
		'a': 'ä', 'e': 'ë', 'i': 'ï', 'o': 'ö', 'u': 'ü', 'y': 'ÿ',
	}, { // unused 0xC9
	}, { // Rsg 0xCA
		'A': 'Å', 'U': 'Ů', 'a': 'å', 'u': 'ů',
	}, { // Cedilla 0xCB
		'C': 'Ç', 'G': 'Ģ', 'K': 'Ķ', 'L': 'Ļ', 'N': 'Ņ', 'R': 'Ŗ', 'S': 'Ş',
		'T': 'Ţ',
		'c': 'ç', 'k': 'ķ', 'l': 'ļ', 'n': 'ņ', 'r': 'ŗ', 's': 'ş', 't': 'ţ',
	}, { //  unused 0xCC
	}, { //  DoubleAcute 0xCD
		'O': 'Ő', 'U': 'Ű', 'o': 'ő', 'u': 'ű',
	}, { // Ogonek 0xCE
		'A': 'Ą', 'E': 'Ę', 'I': 'Į', 'U': 'Ų',
		'a': 'ą', 'e': 'ę', 'i': 'į', 'u': 'ų',
	}, { // Caron 0xCF
		'C': 'Č', 'D': 'Ď', 'E': 'Ě', 'L': 'Ľ', 'N': 'Ň', 'R': 'Ř', 'S': 'Š',
		'T': 'Ť', 'Z': 'Ž',
		'c': 'č', 'd': 'ď', 'e': 'ě', 'l': 'ľ', 'n': 'ň', 'r': 'ř', 's': 'š',
		't': 'ť', 'z': 'ž',
	},
}

func DecodeISO6937(s []byte) string {
	out := make([]rune, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < 0xA0 {
			out = append(out, rune(c))
			continue
		}
		if c >= 0xC1 && c <= 0xCF && c != 0xC9 && c != 0xCC && i+1 < len(s) {
			r, ok := iso6937twoBytes[c-0xC1][s[i+1]]
			if ok {
				out = append(out, r)
				i++
				continue
			}
		}
		out = append(out, iso6937[c-0xA0])
	}
	return string(out)
}

func decodeISO8859(table []rune, s []byte) string {
	out := make([]rune, len(s))
	for i, c := range s {
		if c < 0xa0 {
			out[i] = rune(c)
			continue
		}
		out[i] = table[c-0xa0]
	}
	return string(out)
}
func DecodeISO8859_1(s []byte) string {
	return decodeISO8859(iso8859_1, s)
}

func DecodeISO8859_2(s []byte) string {
	return decodeISO8859(iso8859_2, s)
}

func DecodeISO8859_3(s []byte) string {
	return decodeISO8859(iso8859_3, s)
}

func DecodeISO8859_4(s []byte) string {
	return decodeISO8859(iso8859_4, s)
}

func DecodeISO8859_5(s []byte) string {
	return decodeISO8859(iso8859_5, s)
}

func DecodeISO8859_6(s []byte) string {
	// BUG: two bytes characters not handled
	return decodeISO8859(iso8859_6, s)
}

func DecodeISO8859_7(s []byte) string {
	return decodeISO8859(iso8859_7, s)
}

func DecodeISO8859_8(s []byte) string {
	return decodeISO8859(iso8859_8, s)
}

func DecodeISO8859_9(s []byte) string {
	return decodeISO8859(iso8859_9, s)
}

func DecodeISO8859_10(s []byte) string {
	return decodeISO8859(iso8859_10, s)
}

func DecodeISO8859_11(s []byte) string {
	// BUG: two bytes characters not handled
	return decodeISO8859(iso8859_11, s)
}

func DecodeISO8859_13(s []byte) string {
	return decodeISO8859(iso8859_13, s)
}

func DecodeISO8859_14(s []byte) string {
	return decodeISO8859(iso8859_14, s)
}

func DecodeISO8859_15(s []byte) string {
	return decodeISO8859(iso8859_15, s)
}

var iso8859tab = []func([]byte) string{
	DecodeISO8859_1,
	DecodeISO8859_2,
	DecodeISO8859_3,
	DecodeISO8859_4,
	DecodeISO8859_5,
	DecodeISO8859_6,
	DecodeISO8859_7,
	DecodeISO8859_8,
	DecodeISO8859_9,
	DecodeISO8859_10,
	DecodeISO8859_11,
	nil,
	DecodeISO8859_13,
	DecodeISO8859_14,
	DecodeISO8859_15,
}

// DecodeISO8859 decodes s encoded in ISO8859-part encoding. It panics if
// part < 1 || part == 12 || part > 15.
func DecodeISO8859(part int, s []byte) string {
	if part < 1 || part == 12 || part > 15 {
		panic("not supported table number")
	}
	return iso8859tab[part-1](s)
}
