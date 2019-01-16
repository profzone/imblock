package script

import (
	"bytes"
	"encoding"
	"errors"

	git_chinawayltd_com_golib_tools_courier_enumeration "git.chinawayltd.com/golib/tools/courier/enumeration"
)

var InvalidOpCode = errors.New("invalid OpCode")

func init() {
	git_chinawayltd_com_golib_tools_courier_enumeration.RegisterEnums("OpCode", map[string]string{
		"0":         "0",
		"0NOTEQUAL": "146",
		"1":         "81 - AKA OP_TRUE",
		"10":        "90",
		"11":        "91",
		"12":        "92",
		"13":        "93",
		"14":        "94",
		"15":        "95",
		"16":        "96",
		"1ADD":      "139",
		"1NEGATE":   "79",
		"1SUB":      "140",
		"2":         "82",
		"2DIV":      "142",
		"2DROP":     "109",
		"2DUP":      "110",
		"2MUL":      "141",
		"2OVER":     "112",
		"2ROT":      "113",
		"2SWAP":     "114",
		"3":         "83",
		"3DUP":      "111",
		"4":         "84",
		"5":         "85",
		"6":         "86",
		"7":         "87",
		"8":         "88",
		"9":         "89",
		"ABS":       "144",
		"ADD":       "147",
		"AND":       "132",
		"BOOLAND":   "154",
		"BOOLOR":    "155",
		"CAT":       "126",
		"CHECKLOCKTIMEVERIFY": "177 - AKA OP_NOP2",
		"CHECKMULTISIG":       "174",
		"CHECKMULTISIGVERIFY": "175",
		"CHECKSEQUENCEVERIFY": "178 - AKA OP_NOP3",
		"CHECKSIG":            "172",
		"CHECKSIGVERIFY":      "173",
		"CODESEPARATOR":       "171",
		"DATA_1":              "1",
		"DATA_10":             "10",
		"DATA_11":             "11",
		"DATA_12":             "12",
		"DATA_13":             "13",
		"DATA_14":             "14",
		"DATA_15":             "15",
		"DATA_16":             "16",
		"DATA_17":             "17",
		"DATA_18":             "18",
		"DATA_19":             "19",
		"DATA_2":              "2",
		"DATA_20":             "20",
		"DATA_21":             "21",
		"DATA_22":             "22",
		"DATA_23":             "23",
		"DATA_24":             "24",
		"DATA_25":             "25",
		"DATA_26":             "26",
		"DATA_27":             "27",
		"DATA_28":             "28",
		"DATA_29":             "29",
		"DATA_3":              "3",
		"DATA_30":             "30",
		"DATA_31":             "31",
		"DATA_32":             "32",
		"DATA_33":             "33",
		"DATA_34":             "34",
		"DATA_35":             "35",
		"DATA_36":             "36",
		"DATA_37":             "37",
		"DATA_38":             "38",
		"DATA_39":             "39",
		"DATA_4":              "4",
		"DATA_40":             "40",
		"DATA_41":             "41",
		"DATA_42":             "42",
		"DATA_43":             "43",
		"DATA_44":             "44",
		"DATA_45":             "45",
		"DATA_46":             "46",
		"DATA_47":             "47",
		"DATA_48":             "48",
		"DATA_49":             "49",
		"DATA_5":              "5",
		"DATA_50":             "50",
		"DATA_51":             "51",
		"DATA_52":             "52",
		"DATA_53":             "53",
		"DATA_54":             "54",
		"DATA_55":             "55",
		"DATA_56":             "56",
		"DATA_57":             "57",
		"DATA_58":             "58",
		"DATA_59":             "59",
		"DATA_6":              "6",
		"DATA_60":             "60",
		"DATA_61":             "61",
		"DATA_62":             "62",
		"DATA_63":             "63",
		"DATA_64":             "64",
		"DATA_65":             "65",
		"DATA_66":             "66",
		"DATA_67":             "67",
		"DATA_68":             "68",
		"DATA_69":             "69",
		"DATA_7":              "7",
		"DATA_70":             "70",
		"DATA_71":             "71",
		"DATA_72":             "72",
		"DATA_73":             "73",
		"DATA_74":             "74",
		"DATA_75":             "75",
		"DATA_8":              "8",
		"DATA_9":              "9",
		"DEPTH":               "116",
		"DIV":                 "150",
		"DROP":                "117",
		"DUP":                 "118",
		"ELSE":                "103",
		"ENDIF":               "104",
		"EQUAL":               "135",
		"EQUALVERIFY":         "136",
		"FALSE":               "0 - AKA OP_0",
		"FROMALTSTACK":        "108",
		"GREATERTHAN":         "160",
		"GREATERTHANOREQUAL":  "162",
		"HASH160":             "169",
		"HASH256":             "170",
		"IF":                  "99",
		"IFDUP":               "115",
		"INVALIDOPCODE":       "255 - bitcoin core internal",
		"INVERT":              "131",
		"LEFT":                "128",
		"LESSTHAN":            "159",
		"LESSTHANOREQUAL":     "161",
		"LSHIFT":              "152",
		"MAX":                 "164",
		"MIN":                 "163",
		"MOD":                 "151",
		"MUL":                 "149",
		"NEGATE":              "143",
		"NIP":                 "119",
		"NOP":                 "97",
		"NOP1":                "176",
		"NOP10":               "185",
		"NOP2":                "177",
		"NOP3":                "178",
		"NOP4":                "179",
		"NOP5":                "180",
		"NOP6":                "181",
		"NOP7":                "182",
		"NOP8":                "183",
		"NOP9":                "184",
		"NOT":                 "145",
		"NOTIF":               "100",
		"NUMEQUAL":            "156",
		"NUMEQUALVERIFY":      "157",
		"NUMNOTEQUAL":         "158",
		"OR":                  "133",
		"OVER":                "120",
		"PICK":                "121",
		"PUBKEY":              "254 - bitcoin core internal",
		"PUBKEYHASH":          "253 - bitcoin core internal",
		"PUBKEYS":             "251 - bitcoin core internal",
		"PUSHDATA1":           "76",
		"PUSHDATA2":           "77",
		"PUSHDATA4":           "78",
		"RESERVED":            "80",
		"RESERVED1":           "137",
		"RESERVED2":           "138",
		"RETURN":              "106",
		"RIGHT":               "129",
		"RIPEMD160":           "166",
		"ROLL":                "122",
		"ROT":                 "123",
		"RSHIFT":              "153",
		"SHA1":                "167",
		"SHA256":              "168",
		"SIZE":                "130",
		"SMALLINTEGER":        "250 - bitcoin core internal",
		"SUB":                 "148",
		"SUBSTR":              "127",
		"SWAP":                "124",
		"TOALTSTACK":          "107",
		"TRUE":                "81",
		"TUCK":                "125",
		"UNKNOWN186":          "186",
		"UNKNOWN187":          "187",
		"UNKNOWN188":          "188",
		"UNKNOWN189":          "189",
		"UNKNOWN190":          "190",
		"UNKNOWN191":          "191",
		"UNKNOWN192":          "192",
		"UNKNOWN193":          "193",
		"UNKNOWN194":          "194",
		"UNKNOWN195":          "195",
		"UNKNOWN196":          "196",
		"UNKNOWN197":          "197",
		"UNKNOWN198":          "198",
		"UNKNOWN199":          "199",
		"UNKNOWN200":          "200",
		"UNKNOWN201":          "201",
		"UNKNOWN202":          "202",
		"UNKNOWN203":          "203",
		"UNKNOWN204":          "204",
		"UNKNOWN205":          "205",
		"UNKNOWN206":          "206",
		"UNKNOWN207":          "207",
		"UNKNOWN208":          "208",
		"UNKNOWN209":          "209",
		"UNKNOWN210":          "210",
		"UNKNOWN211":          "211",
		"UNKNOWN212":          "212",
		"UNKNOWN213":          "213",
		"UNKNOWN214":          "214",
		"UNKNOWN215":          "215",
		"UNKNOWN216":          "216",
		"UNKNOWN217":          "217",
		"UNKNOWN218":          "218",
		"UNKNOWN219":          "219",
		"UNKNOWN220":          "220",
		"UNKNOWN221":          "221",
		"UNKNOWN222":          "222",
		"UNKNOWN223":          "223",
		"UNKNOWN224":          "224",
		"UNKNOWN225":          "225",
		"UNKNOWN226":          "226",
		"UNKNOWN227":          "227",
		"UNKNOWN228":          "228",
		"UNKNOWN229":          "229",
		"UNKNOWN230":          "230",
		"UNKNOWN231":          "231",
		"UNKNOWN232":          "232",
		"UNKNOWN233":          "233",
		"UNKNOWN234":          "234",
		"UNKNOWN235":          "235",
		"UNKNOWN236":          "236",
		"UNKNOWN237":          "237",
		"UNKNOWN238":          "238",
		"UNKNOWN239":          "239",
		"UNKNOWN240":          "240",
		"UNKNOWN241":          "241",
		"UNKNOWN242":          "242",
		"UNKNOWN243":          "243",
		"UNKNOWN244":          "244",
		"UNKNOWN245":          "245",
		"UNKNOWN246":          "246",
		"UNKNOWN247":          "247",
		"UNKNOWN248":          "248",
		"UNKNOWN249":          "249",
		"UNKNOWN252":          "252",
		"VER":                 "98",
		"VERIF":               "101",
		"VERIFY":              "105",
		"VERNOTIF":            "102",
		"WITHIN":              "165",
		"XOR":                 "134",
	})
}

func ParseOpCodeFromString(s string) (OpCode, error) {
	switch s {
	case "0":
		return OP_CODE__0, nil
	case "0NOTEQUAL":
		return OP_CODE__0NOTEQUAL, nil
	case "1":
		return OP_CODE__1, nil
	case "10":
		return OP_CODE__10, nil
	case "11":
		return OP_CODE__11, nil
	case "12":
		return OP_CODE__12, nil
	case "13":
		return OP_CODE__13, nil
	case "14":
		return OP_CODE__14, nil
	case "15":
		return OP_CODE__15, nil
	case "16":
		return OP_CODE__16, nil
	case "1ADD":
		return OP_CODE__1ADD, nil
	case "1NEGATE":
		return OP_CODE__1NEGATE, nil
	case "1SUB":
		return OP_CODE__1SUB, nil
	case "2":
		return OP_CODE__2, nil
	case "2DIV":
		return OP_CODE__2DIV, nil
	case "2DROP":
		return OP_CODE__2DROP, nil
	case "2DUP":
		return OP_CODE__2DUP, nil
	case "2MUL":
		return OP_CODE__2MUL, nil
	case "2OVER":
		return OP_CODE__2OVER, nil
	case "2ROT":
		return OP_CODE__2ROT, nil
	case "2SWAP":
		return OP_CODE__2SWAP, nil
	case "3":
		return OP_CODE__3, nil
	case "3DUP":
		return OP_CODE__3DUP, nil
	case "4":
		return OP_CODE__4, nil
	case "5":
		return OP_CODE__5, nil
	case "6":
		return OP_CODE__6, nil
	case "7":
		return OP_CODE__7, nil
	case "8":
		return OP_CODE__8, nil
	case "9":
		return OP_CODE__9, nil
	case "ABS":
		return OP_CODE__ABS, nil
	case "ADD":
		return OP_CODE__ADD, nil
	case "AND":
		return OP_CODE__AND, nil
	case "BOOLAND":
		return OP_CODE__BOOLAND, nil
	case "BOOLOR":
		return OP_CODE__BOOLOR, nil
	case "CAT":
		return OP_CODE__CAT, nil
	case "CHECKLOCKTIMEVERIFY":
		return OP_CODE__CHECKLOCKTIMEVERIFY, nil
	case "CHECKMULTISIG":
		return OP_CODE__CHECKMULTISIG, nil
	case "CHECKMULTISIGVERIFY":
		return OP_CODE__CHECKMULTISIGVERIFY, nil
	case "CHECKSEQUENCEVERIFY":
		return OP_CODE__CHECKSEQUENCEVERIFY, nil
	case "CHECKSIG":
		return OP_CODE__CHECKSIG, nil
	case "CHECKSIGVERIFY":
		return OP_CODE__CHECKSIGVERIFY, nil
	case "CODESEPARATOR":
		return OP_CODE__CODESEPARATOR, nil
	case "DATA_1":
		return OP_CODE__DATA_1, nil
	case "DATA_10":
		return OP_CODE__DATA_10, nil
	case "DATA_11":
		return OP_CODE__DATA_11, nil
	case "DATA_12":
		return OP_CODE__DATA_12, nil
	case "DATA_13":
		return OP_CODE__DATA_13, nil
	case "DATA_14":
		return OP_CODE__DATA_14, nil
	case "DATA_15":
		return OP_CODE__DATA_15, nil
	case "DATA_16":
		return OP_CODE__DATA_16, nil
	case "DATA_17":
		return OP_CODE__DATA_17, nil
	case "DATA_18":
		return OP_CODE__DATA_18, nil
	case "DATA_19":
		return OP_CODE__DATA_19, nil
	case "DATA_2":
		return OP_CODE__DATA_2, nil
	case "DATA_20":
		return OP_CODE__DATA_20, nil
	case "DATA_21":
		return OP_CODE__DATA_21, nil
	case "DATA_22":
		return OP_CODE__DATA_22, nil
	case "DATA_23":
		return OP_CODE__DATA_23, nil
	case "DATA_24":
		return OP_CODE__DATA_24, nil
	case "DATA_25":
		return OP_CODE__DATA_25, nil
	case "DATA_26":
		return OP_CODE__DATA_26, nil
	case "DATA_27":
		return OP_CODE__DATA_27, nil
	case "DATA_28":
		return OP_CODE__DATA_28, nil
	case "DATA_29":
		return OP_CODE__DATA_29, nil
	case "DATA_3":
		return OP_CODE__DATA_3, nil
	case "DATA_30":
		return OP_CODE__DATA_30, nil
	case "DATA_31":
		return OP_CODE__DATA_31, nil
	case "DATA_32":
		return OP_CODE__DATA_32, nil
	case "DATA_33":
		return OP_CODE__DATA_33, nil
	case "DATA_34":
		return OP_CODE__DATA_34, nil
	case "DATA_35":
		return OP_CODE__DATA_35, nil
	case "DATA_36":
		return OP_CODE__DATA_36, nil
	case "DATA_37":
		return OP_CODE__DATA_37, nil
	case "DATA_38":
		return OP_CODE__DATA_38, nil
	case "DATA_39":
		return OP_CODE__DATA_39, nil
	case "DATA_4":
		return OP_CODE__DATA_4, nil
	case "DATA_40":
		return OP_CODE__DATA_40, nil
	case "DATA_41":
		return OP_CODE__DATA_41, nil
	case "DATA_42":
		return OP_CODE__DATA_42, nil
	case "DATA_43":
		return OP_CODE__DATA_43, nil
	case "DATA_44":
		return OP_CODE__DATA_44, nil
	case "DATA_45":
		return OP_CODE__DATA_45, nil
	case "DATA_46":
		return OP_CODE__DATA_46, nil
	case "DATA_47":
		return OP_CODE__DATA_47, nil
	case "DATA_48":
		return OP_CODE__DATA_48, nil
	case "DATA_49":
		return OP_CODE__DATA_49, nil
	case "DATA_5":
		return OP_CODE__DATA_5, nil
	case "DATA_50":
		return OP_CODE__DATA_50, nil
	case "DATA_51":
		return OP_CODE__DATA_51, nil
	case "DATA_52":
		return OP_CODE__DATA_52, nil
	case "DATA_53":
		return OP_CODE__DATA_53, nil
	case "DATA_54":
		return OP_CODE__DATA_54, nil
	case "DATA_55":
		return OP_CODE__DATA_55, nil
	case "DATA_56":
		return OP_CODE__DATA_56, nil
	case "DATA_57":
		return OP_CODE__DATA_57, nil
	case "DATA_58":
		return OP_CODE__DATA_58, nil
	case "DATA_59":
		return OP_CODE__DATA_59, nil
	case "DATA_6":
		return OP_CODE__DATA_6, nil
	case "DATA_60":
		return OP_CODE__DATA_60, nil
	case "DATA_61":
		return OP_CODE__DATA_61, nil
	case "DATA_62":
		return OP_CODE__DATA_62, nil
	case "DATA_63":
		return OP_CODE__DATA_63, nil
	case "DATA_64":
		return OP_CODE__DATA_64, nil
	case "DATA_65":
		return OP_CODE__DATA_65, nil
	case "DATA_66":
		return OP_CODE__DATA_66, nil
	case "DATA_67":
		return OP_CODE__DATA_67, nil
	case "DATA_68":
		return OP_CODE__DATA_68, nil
	case "DATA_69":
		return OP_CODE__DATA_69, nil
	case "DATA_7":
		return OP_CODE__DATA_7, nil
	case "DATA_70":
		return OP_CODE__DATA_70, nil
	case "DATA_71":
		return OP_CODE__DATA_71, nil
	case "DATA_72":
		return OP_CODE__DATA_72, nil
	case "DATA_73":
		return OP_CODE__DATA_73, nil
	case "DATA_74":
		return OP_CODE__DATA_74, nil
	case "DATA_75":
		return OP_CODE__DATA_75, nil
	case "DATA_8":
		return OP_CODE__DATA_8, nil
	case "DATA_9":
		return OP_CODE__DATA_9, nil
	case "DEPTH":
		return OP_CODE__DEPTH, nil
	case "DIV":
		return OP_CODE__DIV, nil
	case "DROP":
		return OP_CODE__DROP, nil
	case "DUP":
		return OP_CODE__DUP, nil
	case "ELSE":
		return OP_CODE__ELSE, nil
	case "ENDIF":
		return OP_CODE__ENDIF, nil
	case "EQUAL":
		return OP_CODE__EQUAL, nil
	case "EQUALVERIFY":
		return OP_CODE__EQUALVERIFY, nil
	case "FALSE":
		return OP_CODE__FALSE, nil
	case "FROMALTSTACK":
		return OP_CODE__FROMALTSTACK, nil
	case "GREATERTHAN":
		return OP_CODE__GREATERTHAN, nil
	case "GREATERTHANOREQUAL":
		return OP_CODE__GREATERTHANOREQUAL, nil
	case "HASH160":
		return OP_CODE__HASH160, nil
	case "HASH256":
		return OP_CODE__HASH256, nil
	case "IF":
		return OP_CODE__IF, nil
	case "IFDUP":
		return OP_CODE__IFDUP, nil
	case "INVALIDOPCODE":
		return OP_CODE__INVALIDOPCODE, nil
	case "INVERT":
		return OP_CODE__INVERT, nil
	case "LEFT":
		return OP_CODE__LEFT, nil
	case "LESSTHAN":
		return OP_CODE__LESSTHAN, nil
	case "LESSTHANOREQUAL":
		return OP_CODE__LESSTHANOREQUAL, nil
	case "LSHIFT":
		return OP_CODE__LSHIFT, nil
	case "MAX":
		return OP_CODE__MAX, nil
	case "MIN":
		return OP_CODE__MIN, nil
	case "MOD":
		return OP_CODE__MOD, nil
	case "MUL":
		return OP_CODE__MUL, nil
	case "NEGATE":
		return OP_CODE__NEGATE, nil
	case "NIP":
		return OP_CODE__NIP, nil
	case "NOP":
		return OP_CODE__NOP, nil
	case "NOP1":
		return OP_CODE__NOP1, nil
	case "NOP10":
		return OP_CODE__NOP10, nil
	case "NOP2":
		return OP_CODE__NOP2, nil
	case "NOP3":
		return OP_CODE__NOP3, nil
	case "NOP4":
		return OP_CODE__NOP4, nil
	case "NOP5":
		return OP_CODE__NOP5, nil
	case "NOP6":
		return OP_CODE__NOP6, nil
	case "NOP7":
		return OP_CODE__NOP7, nil
	case "NOP8":
		return OP_CODE__NOP8, nil
	case "NOP9":
		return OP_CODE__NOP9, nil
	case "NOT":
		return OP_CODE__NOT, nil
	case "NOTIF":
		return OP_CODE__NOTIF, nil
	case "NUMEQUAL":
		return OP_CODE__NUMEQUAL, nil
	case "NUMEQUALVERIFY":
		return OP_CODE__NUMEQUALVERIFY, nil
	case "NUMNOTEQUAL":
		return OP_CODE__NUMNOTEQUAL, nil
	case "OR":
		return OP_CODE__OR, nil
	case "OVER":
		return OP_CODE__OVER, nil
	case "PICK":
		return OP_CODE__PICK, nil
	case "PUBKEY":
		return OP_CODE__PUBKEY, nil
	case "PUBKEYHASH":
		return OP_CODE__PUBKEYHASH, nil
	case "PUBKEYS":
		return OP_CODE__PUBKEYS, nil
	case "PUSHDATA1":
		return OP_CODE__PUSHDATA1, nil
	case "PUSHDATA2":
		return OP_CODE__PUSHDATA2, nil
	case "PUSHDATA4":
		return OP_CODE__PUSHDATA4, nil
	case "RESERVED":
		return OP_CODE__RESERVED, nil
	case "RESERVED1":
		return OP_CODE__RESERVED1, nil
	case "RESERVED2":
		return OP_CODE__RESERVED2, nil
	case "RETURN":
		return OP_CODE__RETURN, nil
	case "RIGHT":
		return OP_CODE__RIGHT, nil
	case "RIPEMD160":
		return OP_CODE__RIPEMD160, nil
	case "ROLL":
		return OP_CODE__ROLL, nil
	case "ROT":
		return OP_CODE__ROT, nil
	case "RSHIFT":
		return OP_CODE__RSHIFT, nil
	case "SHA1":
		return OP_CODE__SHA1, nil
	case "SHA256":
		return OP_CODE__SHA256, nil
	case "SIZE":
		return OP_CODE__SIZE, nil
	case "SMALLINTEGER":
		return OP_CODE__SMALLINTEGER, nil
	case "SUB":
		return OP_CODE__SUB, nil
	case "SUBSTR":
		return OP_CODE__SUBSTR, nil
	case "SWAP":
		return OP_CODE__SWAP, nil
	case "TOALTSTACK":
		return OP_CODE__TOALTSTACK, nil
	case "TRUE":
		return OP_CODE__TRUE, nil
	case "TUCK":
		return OP_CODE__TUCK, nil
	case "UNKNOWN186":
		return OP_CODE__UNKNOWN186, nil
	case "UNKNOWN187":
		return OP_CODE__UNKNOWN187, nil
	case "UNKNOWN188":
		return OP_CODE__UNKNOWN188, nil
	case "UNKNOWN189":
		return OP_CODE__UNKNOWN189, nil
	case "UNKNOWN190":
		return OP_CODE__UNKNOWN190, nil
	case "UNKNOWN191":
		return OP_CODE__UNKNOWN191, nil
	case "UNKNOWN192":
		return OP_CODE__UNKNOWN192, nil
	case "UNKNOWN193":
		return OP_CODE__UNKNOWN193, nil
	case "UNKNOWN194":
		return OP_CODE__UNKNOWN194, nil
	case "UNKNOWN195":
		return OP_CODE__UNKNOWN195, nil
	case "UNKNOWN196":
		return OP_CODE__UNKNOWN196, nil
	case "UNKNOWN197":
		return OP_CODE__UNKNOWN197, nil
	case "UNKNOWN198":
		return OP_CODE__UNKNOWN198, nil
	case "UNKNOWN199":
		return OP_CODE__UNKNOWN199, nil
	case "UNKNOWN200":
		return OP_CODE__UNKNOWN200, nil
	case "UNKNOWN201":
		return OP_CODE__UNKNOWN201, nil
	case "UNKNOWN202":
		return OP_CODE__UNKNOWN202, nil
	case "UNKNOWN203":
		return OP_CODE__UNKNOWN203, nil
	case "UNKNOWN204":
		return OP_CODE__UNKNOWN204, nil
	case "UNKNOWN205":
		return OP_CODE__UNKNOWN205, nil
	case "UNKNOWN206":
		return OP_CODE__UNKNOWN206, nil
	case "UNKNOWN207":
		return OP_CODE__UNKNOWN207, nil
	case "UNKNOWN208":
		return OP_CODE__UNKNOWN208, nil
	case "UNKNOWN209":
		return OP_CODE__UNKNOWN209, nil
	case "UNKNOWN210":
		return OP_CODE__UNKNOWN210, nil
	case "UNKNOWN211":
		return OP_CODE__UNKNOWN211, nil
	case "UNKNOWN212":
		return OP_CODE__UNKNOWN212, nil
	case "UNKNOWN213":
		return OP_CODE__UNKNOWN213, nil
	case "UNKNOWN214":
		return OP_CODE__UNKNOWN214, nil
	case "UNKNOWN215":
		return OP_CODE__UNKNOWN215, nil
	case "UNKNOWN216":
		return OP_CODE__UNKNOWN216, nil
	case "UNKNOWN217":
		return OP_CODE__UNKNOWN217, nil
	case "UNKNOWN218":
		return OP_CODE__UNKNOWN218, nil
	case "UNKNOWN219":
		return OP_CODE__UNKNOWN219, nil
	case "UNKNOWN220":
		return OP_CODE__UNKNOWN220, nil
	case "UNKNOWN221":
		return OP_CODE__UNKNOWN221, nil
	case "UNKNOWN222":
		return OP_CODE__UNKNOWN222, nil
	case "UNKNOWN223":
		return OP_CODE__UNKNOWN223, nil
	case "UNKNOWN224":
		return OP_CODE__UNKNOWN224, nil
	case "UNKNOWN225":
		return OP_CODE__UNKNOWN225, nil
	case "UNKNOWN226":
		return OP_CODE__UNKNOWN226, nil
	case "UNKNOWN227":
		return OP_CODE__UNKNOWN227, nil
	case "UNKNOWN228":
		return OP_CODE__UNKNOWN228, nil
	case "UNKNOWN229":
		return OP_CODE__UNKNOWN229, nil
	case "UNKNOWN230":
		return OP_CODE__UNKNOWN230, nil
	case "UNKNOWN231":
		return OP_CODE__UNKNOWN231, nil
	case "UNKNOWN232":
		return OP_CODE__UNKNOWN232, nil
	case "UNKNOWN233":
		return OP_CODE__UNKNOWN233, nil
	case "UNKNOWN234":
		return OP_CODE__UNKNOWN234, nil
	case "UNKNOWN235":
		return OP_CODE__UNKNOWN235, nil
	case "UNKNOWN236":
		return OP_CODE__UNKNOWN236, nil
	case "UNKNOWN237":
		return OP_CODE__UNKNOWN237, nil
	case "UNKNOWN238":
		return OP_CODE__UNKNOWN238, nil
	case "UNKNOWN239":
		return OP_CODE__UNKNOWN239, nil
	case "UNKNOWN240":
		return OP_CODE__UNKNOWN240, nil
	case "UNKNOWN241":
		return OP_CODE__UNKNOWN241, nil
	case "UNKNOWN242":
		return OP_CODE__UNKNOWN242, nil
	case "UNKNOWN243":
		return OP_CODE__UNKNOWN243, nil
	case "UNKNOWN244":
		return OP_CODE__UNKNOWN244, nil
	case "UNKNOWN245":
		return OP_CODE__UNKNOWN245, nil
	case "UNKNOWN246":
		return OP_CODE__UNKNOWN246, nil
	case "UNKNOWN247":
		return OP_CODE__UNKNOWN247, nil
	case "UNKNOWN248":
		return OP_CODE__UNKNOWN248, nil
	case "UNKNOWN249":
		return OP_CODE__UNKNOWN249, nil
	case "UNKNOWN252":
		return OP_CODE__UNKNOWN252, nil
	case "VER":
		return OP_CODE__VER, nil
	case "VERIF":
		return OP_CODE__VERIF, nil
	case "VERIFY":
		return OP_CODE__VERIFY, nil
	case "VERNOTIF":
		return OP_CODE__VERNOTIF, nil
	case "WITHIN":
		return OP_CODE__WITHIN, nil
	case "XOR":
		return OP_CODE__XOR, nil
	}
	return OP_CODE__0, InvalidOpCode
}

func ParseOpCodeFromLabelString(s string) (OpCode, error) {
	switch s {
	case "0":
		return OP_CODE__0, nil
	case "146":
		return OP_CODE__0NOTEQUAL, nil
	case "81 - AKA OP_TRUE":
		return OP_CODE__1, nil
	case "90":
		return OP_CODE__10, nil
	case "91":
		return OP_CODE__11, nil
	case "92":
		return OP_CODE__12, nil
	case "93":
		return OP_CODE__13, nil
	case "94":
		return OP_CODE__14, nil
	case "95":
		return OP_CODE__15, nil
	case "96":
		return OP_CODE__16, nil
	case "139":
		return OP_CODE__1ADD, nil
	case "79":
		return OP_CODE__1NEGATE, nil
	case "140":
		return OP_CODE__1SUB, nil
	case "82":
		return OP_CODE__2, nil
	case "142":
		return OP_CODE__2DIV, nil
	case "109":
		return OP_CODE__2DROP, nil
	case "110":
		return OP_CODE__2DUP, nil
	case "141":
		return OP_CODE__2MUL, nil
	case "112":
		return OP_CODE__2OVER, nil
	case "113":
		return OP_CODE__2ROT, nil
	case "114":
		return OP_CODE__2SWAP, nil
	case "83":
		return OP_CODE__3, nil
	case "111":
		return OP_CODE__3DUP, nil
	case "84":
		return OP_CODE__4, nil
	case "85":
		return OP_CODE__5, nil
	case "86":
		return OP_CODE__6, nil
	case "87":
		return OP_CODE__7, nil
	case "88":
		return OP_CODE__8, nil
	case "89":
		return OP_CODE__9, nil
	case "144":
		return OP_CODE__ABS, nil
	case "147":
		return OP_CODE__ADD, nil
	case "132":
		return OP_CODE__AND, nil
	case "154":
		return OP_CODE__BOOLAND, nil
	case "155":
		return OP_CODE__BOOLOR, nil
	case "126":
		return OP_CODE__CAT, nil
	case "177 - AKA OP_NOP2":
		return OP_CODE__CHECKLOCKTIMEVERIFY, nil
	case "174":
		return OP_CODE__CHECKMULTISIG, nil
	case "175":
		return OP_CODE__CHECKMULTISIGVERIFY, nil
	case "178 - AKA OP_NOP3":
		return OP_CODE__CHECKSEQUENCEVERIFY, nil
	case "172":
		return OP_CODE__CHECKSIG, nil
	case "173":
		return OP_CODE__CHECKSIGVERIFY, nil
	case "171":
		return OP_CODE__CODESEPARATOR, nil
	case "1":
		return OP_CODE__DATA_1, nil
	case "10":
		return OP_CODE__DATA_10, nil
	case "11":
		return OP_CODE__DATA_11, nil
	case "12":
		return OP_CODE__DATA_12, nil
	case "13":
		return OP_CODE__DATA_13, nil
	case "14":
		return OP_CODE__DATA_14, nil
	case "15":
		return OP_CODE__DATA_15, nil
	case "16":
		return OP_CODE__DATA_16, nil
	case "17":
		return OP_CODE__DATA_17, nil
	case "18":
		return OP_CODE__DATA_18, nil
	case "19":
		return OP_CODE__DATA_19, nil
	case "2":
		return OP_CODE__DATA_2, nil
	case "20":
		return OP_CODE__DATA_20, nil
	case "21":
		return OP_CODE__DATA_21, nil
	case "22":
		return OP_CODE__DATA_22, nil
	case "23":
		return OP_CODE__DATA_23, nil
	case "24":
		return OP_CODE__DATA_24, nil
	case "25":
		return OP_CODE__DATA_25, nil
	case "26":
		return OP_CODE__DATA_26, nil
	case "27":
		return OP_CODE__DATA_27, nil
	case "28":
		return OP_CODE__DATA_28, nil
	case "29":
		return OP_CODE__DATA_29, nil
	case "3":
		return OP_CODE__DATA_3, nil
	case "30":
		return OP_CODE__DATA_30, nil
	case "31":
		return OP_CODE__DATA_31, nil
	case "32":
		return OP_CODE__DATA_32, nil
	case "33":
		return OP_CODE__DATA_33, nil
	case "34":
		return OP_CODE__DATA_34, nil
	case "35":
		return OP_CODE__DATA_35, nil
	case "36":
		return OP_CODE__DATA_36, nil
	case "37":
		return OP_CODE__DATA_37, nil
	case "38":
		return OP_CODE__DATA_38, nil
	case "39":
		return OP_CODE__DATA_39, nil
	case "4":
		return OP_CODE__DATA_4, nil
	case "40":
		return OP_CODE__DATA_40, nil
	case "41":
		return OP_CODE__DATA_41, nil
	case "42":
		return OP_CODE__DATA_42, nil
	case "43":
		return OP_CODE__DATA_43, nil
	case "44":
		return OP_CODE__DATA_44, nil
	case "45":
		return OP_CODE__DATA_45, nil
	case "46":
		return OP_CODE__DATA_46, nil
	case "47":
		return OP_CODE__DATA_47, nil
	case "48":
		return OP_CODE__DATA_48, nil
	case "49":
		return OP_CODE__DATA_49, nil
	case "5":
		return OP_CODE__DATA_5, nil
	case "50":
		return OP_CODE__DATA_50, nil
	case "51":
		return OP_CODE__DATA_51, nil
	case "52":
		return OP_CODE__DATA_52, nil
	case "53":
		return OP_CODE__DATA_53, nil
	case "54":
		return OP_CODE__DATA_54, nil
	case "55":
		return OP_CODE__DATA_55, nil
	case "56":
		return OP_CODE__DATA_56, nil
	case "57":
		return OP_CODE__DATA_57, nil
	case "58":
		return OP_CODE__DATA_58, nil
	case "59":
		return OP_CODE__DATA_59, nil
	case "6":
		return OP_CODE__DATA_6, nil
	case "60":
		return OP_CODE__DATA_60, nil
	case "61":
		return OP_CODE__DATA_61, nil
	case "62":
		return OP_CODE__DATA_62, nil
	case "63":
		return OP_CODE__DATA_63, nil
	case "64":
		return OP_CODE__DATA_64, nil
	case "65":
		return OP_CODE__DATA_65, nil
	case "66":
		return OP_CODE__DATA_66, nil
	case "67":
		return OP_CODE__DATA_67, nil
	case "68":
		return OP_CODE__DATA_68, nil
	case "69":
		return OP_CODE__DATA_69, nil
	case "7":
		return OP_CODE__DATA_7, nil
	case "70":
		return OP_CODE__DATA_70, nil
	case "71":
		return OP_CODE__DATA_71, nil
	case "72":
		return OP_CODE__DATA_72, nil
	case "73":
		return OP_CODE__DATA_73, nil
	case "74":
		return OP_CODE__DATA_74, nil
	case "75":
		return OP_CODE__DATA_75, nil
	case "8":
		return OP_CODE__DATA_8, nil
	case "9":
		return OP_CODE__DATA_9, nil
	case "116":
		return OP_CODE__DEPTH, nil
	case "150":
		return OP_CODE__DIV, nil
	case "117":
		return OP_CODE__DROP, nil
	case "118":
		return OP_CODE__DUP, nil
	case "103":
		return OP_CODE__ELSE, nil
	case "104":
		return OP_CODE__ENDIF, nil
	case "135":
		return OP_CODE__EQUAL, nil
	case "136":
		return OP_CODE__EQUALVERIFY, nil
	case "0 - AKA OP_0":
		return OP_CODE__FALSE, nil
	case "108":
		return OP_CODE__FROMALTSTACK, nil
	case "160":
		return OP_CODE__GREATERTHAN, nil
	case "162":
		return OP_CODE__GREATERTHANOREQUAL, nil
	case "169":
		return OP_CODE__HASH160, nil
	case "170":
		return OP_CODE__HASH256, nil
	case "99":
		return OP_CODE__IF, nil
	case "115":
		return OP_CODE__IFDUP, nil
	case "255 - bitcoin core internal":
		return OP_CODE__INVALIDOPCODE, nil
	case "131":
		return OP_CODE__INVERT, nil
	case "128":
		return OP_CODE__LEFT, nil
	case "159":
		return OP_CODE__LESSTHAN, nil
	case "161":
		return OP_CODE__LESSTHANOREQUAL, nil
	case "152":
		return OP_CODE__LSHIFT, nil
	case "164":
		return OP_CODE__MAX, nil
	case "163":
		return OP_CODE__MIN, nil
	case "151":
		return OP_CODE__MOD, nil
	case "149":
		return OP_CODE__MUL, nil
	case "143":
		return OP_CODE__NEGATE, nil
	case "119":
		return OP_CODE__NIP, nil
	case "97":
		return OP_CODE__NOP, nil
	case "176":
		return OP_CODE__NOP1, nil
	case "185":
		return OP_CODE__NOP10, nil
	case "177":
		return OP_CODE__NOP2, nil
	case "178":
		return OP_CODE__NOP3, nil
	case "179":
		return OP_CODE__NOP4, nil
	case "180":
		return OP_CODE__NOP5, nil
	case "181":
		return OP_CODE__NOP6, nil
	case "182":
		return OP_CODE__NOP7, nil
	case "183":
		return OP_CODE__NOP8, nil
	case "184":
		return OP_CODE__NOP9, nil
	case "145":
		return OP_CODE__NOT, nil
	case "100":
		return OP_CODE__NOTIF, nil
	case "156":
		return OP_CODE__NUMEQUAL, nil
	case "157":
		return OP_CODE__NUMEQUALVERIFY, nil
	case "158":
		return OP_CODE__NUMNOTEQUAL, nil
	case "133":
		return OP_CODE__OR, nil
	case "120":
		return OP_CODE__OVER, nil
	case "121":
		return OP_CODE__PICK, nil
	case "254 - bitcoin core internal":
		return OP_CODE__PUBKEY, nil
	case "253 - bitcoin core internal":
		return OP_CODE__PUBKEYHASH, nil
	case "251 - bitcoin core internal":
		return OP_CODE__PUBKEYS, nil
	case "76":
		return OP_CODE__PUSHDATA1, nil
	case "77":
		return OP_CODE__PUSHDATA2, nil
	case "78":
		return OP_CODE__PUSHDATA4, nil
	case "80":
		return OP_CODE__RESERVED, nil
	case "137":
		return OP_CODE__RESERVED1, nil
	case "138":
		return OP_CODE__RESERVED2, nil
	case "106":
		return OP_CODE__RETURN, nil
	case "129":
		return OP_CODE__RIGHT, nil
	case "166":
		return OP_CODE__RIPEMD160, nil
	case "122":
		return OP_CODE__ROLL, nil
	case "123":
		return OP_CODE__ROT, nil
	case "153":
		return OP_CODE__RSHIFT, nil
	case "167":
		return OP_CODE__SHA1, nil
	case "168":
		return OP_CODE__SHA256, nil
	case "130":
		return OP_CODE__SIZE, nil
	case "250 - bitcoin core internal":
		return OP_CODE__SMALLINTEGER, nil
	case "148":
		return OP_CODE__SUB, nil
	case "127":
		return OP_CODE__SUBSTR, nil
	case "124":
		return OP_CODE__SWAP, nil
	case "107":
		return OP_CODE__TOALTSTACK, nil
	case "81":
		return OP_CODE__TRUE, nil
	case "125":
		return OP_CODE__TUCK, nil
	case "186":
		return OP_CODE__UNKNOWN186, nil
	case "187":
		return OP_CODE__UNKNOWN187, nil
	case "188":
		return OP_CODE__UNKNOWN188, nil
	case "189":
		return OP_CODE__UNKNOWN189, nil
	case "190":
		return OP_CODE__UNKNOWN190, nil
	case "191":
		return OP_CODE__UNKNOWN191, nil
	case "192":
		return OP_CODE__UNKNOWN192, nil
	case "193":
		return OP_CODE__UNKNOWN193, nil
	case "194":
		return OP_CODE__UNKNOWN194, nil
	case "195":
		return OP_CODE__UNKNOWN195, nil
	case "196":
		return OP_CODE__UNKNOWN196, nil
	case "197":
		return OP_CODE__UNKNOWN197, nil
	case "198":
		return OP_CODE__UNKNOWN198, nil
	case "199":
		return OP_CODE__UNKNOWN199, nil
	case "200":
		return OP_CODE__UNKNOWN200, nil
	case "201":
		return OP_CODE__UNKNOWN201, nil
	case "202":
		return OP_CODE__UNKNOWN202, nil
	case "203":
		return OP_CODE__UNKNOWN203, nil
	case "204":
		return OP_CODE__UNKNOWN204, nil
	case "205":
		return OP_CODE__UNKNOWN205, nil
	case "206":
		return OP_CODE__UNKNOWN206, nil
	case "207":
		return OP_CODE__UNKNOWN207, nil
	case "208":
		return OP_CODE__UNKNOWN208, nil
	case "209":
		return OP_CODE__UNKNOWN209, nil
	case "210":
		return OP_CODE__UNKNOWN210, nil
	case "211":
		return OP_CODE__UNKNOWN211, nil
	case "212":
		return OP_CODE__UNKNOWN212, nil
	case "213":
		return OP_CODE__UNKNOWN213, nil
	case "214":
		return OP_CODE__UNKNOWN214, nil
	case "215":
		return OP_CODE__UNKNOWN215, nil
	case "216":
		return OP_CODE__UNKNOWN216, nil
	case "217":
		return OP_CODE__UNKNOWN217, nil
	case "218":
		return OP_CODE__UNKNOWN218, nil
	case "219":
		return OP_CODE__UNKNOWN219, nil
	case "220":
		return OP_CODE__UNKNOWN220, nil
	case "221":
		return OP_CODE__UNKNOWN221, nil
	case "222":
		return OP_CODE__UNKNOWN222, nil
	case "223":
		return OP_CODE__UNKNOWN223, nil
	case "224":
		return OP_CODE__UNKNOWN224, nil
	case "225":
		return OP_CODE__UNKNOWN225, nil
	case "226":
		return OP_CODE__UNKNOWN226, nil
	case "227":
		return OP_CODE__UNKNOWN227, nil
	case "228":
		return OP_CODE__UNKNOWN228, nil
	case "229":
		return OP_CODE__UNKNOWN229, nil
	case "230":
		return OP_CODE__UNKNOWN230, nil
	case "231":
		return OP_CODE__UNKNOWN231, nil
	case "232":
		return OP_CODE__UNKNOWN232, nil
	case "233":
		return OP_CODE__UNKNOWN233, nil
	case "234":
		return OP_CODE__UNKNOWN234, nil
	case "235":
		return OP_CODE__UNKNOWN235, nil
	case "236":
		return OP_CODE__UNKNOWN236, nil
	case "237":
		return OP_CODE__UNKNOWN237, nil
	case "238":
		return OP_CODE__UNKNOWN238, nil
	case "239":
		return OP_CODE__UNKNOWN239, nil
	case "240":
		return OP_CODE__UNKNOWN240, nil
	case "241":
		return OP_CODE__UNKNOWN241, nil
	case "242":
		return OP_CODE__UNKNOWN242, nil
	case "243":
		return OP_CODE__UNKNOWN243, nil
	case "244":
		return OP_CODE__UNKNOWN244, nil
	case "245":
		return OP_CODE__UNKNOWN245, nil
	case "246":
		return OP_CODE__UNKNOWN246, nil
	case "247":
		return OP_CODE__UNKNOWN247, nil
	case "248":
		return OP_CODE__UNKNOWN248, nil
	case "249":
		return OP_CODE__UNKNOWN249, nil
	case "252":
		return OP_CODE__UNKNOWN252, nil
	case "98":
		return OP_CODE__VER, nil
	case "101":
		return OP_CODE__VERIF, nil
	case "105":
		return OP_CODE__VERIFY, nil
	case "102":
		return OP_CODE__VERNOTIF, nil
	case "165":
		return OP_CODE__WITHIN, nil
	case "134":
		return OP_CODE__XOR, nil
	}
	return OP_CODE__0, InvalidOpCode
}

func (v OpCode) String() string {
	switch v {
	case OP_CODE__0NOTEQUAL:
		return "0NOTEQUAL"
	case OP_CODE__10:
		return "10"
	case OP_CODE__11:
		return "11"
	case OP_CODE__12:
		return "12"
	case OP_CODE__13:
		return "13"
	case OP_CODE__14:
		return "14"
	case OP_CODE__15:
		return "15"
	case OP_CODE__16:
		return "16"
	case OP_CODE__1ADD:
		return "1ADD"
	case OP_CODE__1NEGATE:
		return "1NEGATE"
	case OP_CODE__1SUB:
		return "1SUB"
	case OP_CODE__2:
		return "2"
	case OP_CODE__2DIV:
		return "2DIV"
	case OP_CODE__2DROP:
		return "2DROP"
	case OP_CODE__2DUP:
		return "2DUP"
	case OP_CODE__2MUL:
		return "2MUL"
	case OP_CODE__2OVER:
		return "2OVER"
	case OP_CODE__2ROT:
		return "2ROT"
	case OP_CODE__2SWAP:
		return "2SWAP"
	case OP_CODE__3:
		return "3"
	case OP_CODE__3DUP:
		return "3DUP"
	case OP_CODE__4:
		return "4"
	case OP_CODE__5:
		return "5"
	case OP_CODE__6:
		return "6"
	case OP_CODE__7:
		return "7"
	case OP_CODE__8:
		return "8"
	case OP_CODE__9:
		return "9"
	case OP_CODE__ABS:
		return "ABS"
	case OP_CODE__ADD:
		return "ADD"
	case OP_CODE__AND:
		return "AND"
	case OP_CODE__BOOLAND:
		return "BOOLAND"
	case OP_CODE__BOOLOR:
		return "BOOLOR"
	case OP_CODE__CAT:
		return "CAT"
	case OP_CODE__CHECKLOCKTIMEVERIFY:
		return "CHECKLOCKTIMEVERIFY"
	case OP_CODE__CHECKMULTISIG:
		return "CHECKMULTISIG"
	case OP_CODE__CHECKMULTISIGVERIFY:
		return "CHECKMULTISIGVERIFY"
	case OP_CODE__CHECKSEQUENCEVERIFY:
		return "CHECKSEQUENCEVERIFY"
	case OP_CODE__CHECKSIG:
		return "CHECKSIG"
	case OP_CODE__CHECKSIGVERIFY:
		return "CHECKSIGVERIFY"
	case OP_CODE__CODESEPARATOR:
		return "CODESEPARATOR"
	case OP_CODE__DATA_1:
		return "DATA_1"
	case OP_CODE__DATA_10:
		return "DATA_10"
	case OP_CODE__DATA_11:
		return "DATA_11"
	case OP_CODE__DATA_12:
		return "DATA_12"
	case OP_CODE__DATA_13:
		return "DATA_13"
	case OP_CODE__DATA_14:
		return "DATA_14"
	case OP_CODE__DATA_15:
		return "DATA_15"
	case OP_CODE__DATA_16:
		return "DATA_16"
	case OP_CODE__DATA_17:
		return "DATA_17"
	case OP_CODE__DATA_18:
		return "DATA_18"
	case OP_CODE__DATA_19:
		return "DATA_19"
	case OP_CODE__DATA_2:
		return "DATA_2"
	case OP_CODE__DATA_20:
		return "DATA_20"
	case OP_CODE__DATA_21:
		return "DATA_21"
	case OP_CODE__DATA_22:
		return "DATA_22"
	case OP_CODE__DATA_23:
		return "DATA_23"
	case OP_CODE__DATA_24:
		return "DATA_24"
	case OP_CODE__DATA_25:
		return "DATA_25"
	case OP_CODE__DATA_26:
		return "DATA_26"
	case OP_CODE__DATA_27:
		return "DATA_27"
	case OP_CODE__DATA_28:
		return "DATA_28"
	case OP_CODE__DATA_29:
		return "DATA_29"
	case OP_CODE__DATA_3:
		return "DATA_3"
	case OP_CODE__DATA_30:
		return "DATA_30"
	case OP_CODE__DATA_31:
		return "DATA_31"
	case OP_CODE__DATA_32:
		return "DATA_32"
	case OP_CODE__DATA_33:
		return "DATA_33"
	case OP_CODE__DATA_34:
		return "DATA_34"
	case OP_CODE__DATA_35:
		return "DATA_35"
	case OP_CODE__DATA_36:
		return "DATA_36"
	case OP_CODE__DATA_37:
		return "DATA_37"
	case OP_CODE__DATA_38:
		return "DATA_38"
	case OP_CODE__DATA_39:
		return "DATA_39"
	case OP_CODE__DATA_4:
		return "DATA_4"
	case OP_CODE__DATA_40:
		return "DATA_40"
	case OP_CODE__DATA_41:
		return "DATA_41"
	case OP_CODE__DATA_42:
		return "DATA_42"
	case OP_CODE__DATA_43:
		return "DATA_43"
	case OP_CODE__DATA_44:
		return "DATA_44"
	case OP_CODE__DATA_45:
		return "DATA_45"
	case OP_CODE__DATA_46:
		return "DATA_46"
	case OP_CODE__DATA_47:
		return "DATA_47"
	case OP_CODE__DATA_48:
		return "DATA_48"
	case OP_CODE__DATA_49:
		return "DATA_49"
	case OP_CODE__DATA_5:
		return "DATA_5"
	case OP_CODE__DATA_50:
		return "DATA_50"
	case OP_CODE__DATA_51:
		return "DATA_51"
	case OP_CODE__DATA_52:
		return "DATA_52"
	case OP_CODE__DATA_53:
		return "DATA_53"
	case OP_CODE__DATA_54:
		return "DATA_54"
	case OP_CODE__DATA_55:
		return "DATA_55"
	case OP_CODE__DATA_56:
		return "DATA_56"
	case OP_CODE__DATA_57:
		return "DATA_57"
	case OP_CODE__DATA_58:
		return "DATA_58"
	case OP_CODE__DATA_59:
		return "DATA_59"
	case OP_CODE__DATA_6:
		return "DATA_6"
	case OP_CODE__DATA_60:
		return "DATA_60"
	case OP_CODE__DATA_61:
		return "DATA_61"
	case OP_CODE__DATA_62:
		return "DATA_62"
	case OP_CODE__DATA_63:
		return "DATA_63"
	case OP_CODE__DATA_64:
		return "DATA_64"
	case OP_CODE__DATA_65:
		return "DATA_65"
	case OP_CODE__DATA_66:
		return "DATA_66"
	case OP_CODE__DATA_67:
		return "DATA_67"
	case OP_CODE__DATA_68:
		return "DATA_68"
	case OP_CODE__DATA_69:
		return "DATA_69"
	case OP_CODE__DATA_7:
		return "DATA_7"
	case OP_CODE__DATA_70:
		return "DATA_70"
	case OP_CODE__DATA_71:
		return "DATA_71"
	case OP_CODE__DATA_72:
		return "DATA_72"
	case OP_CODE__DATA_73:
		return "DATA_73"
	case OP_CODE__DATA_74:
		return "DATA_74"
	case OP_CODE__DATA_75:
		return "DATA_75"
	case OP_CODE__DATA_8:
		return "DATA_8"
	case OP_CODE__DATA_9:
		return "DATA_9"
	case OP_CODE__DEPTH:
		return "DEPTH"
	case OP_CODE__DIV:
		return "DIV"
	case OP_CODE__DROP:
		return "DROP"
	case OP_CODE__DUP:
		return "DUP"
	case OP_CODE__ELSE:
		return "ELSE"
	case OP_CODE__ENDIF:
		return "ENDIF"
	case OP_CODE__EQUAL:
		return "EQUAL"
	case OP_CODE__EQUALVERIFY:
		return "EQUALVERIFY"
	case OP_CODE__FALSE:
		return "0"
	case OP_CODE__FROMALTSTACK:
		return "FROMALTSTACK"
	case OP_CODE__GREATERTHAN:
		return "GREATERTHAN"
	case OP_CODE__GREATERTHANOREQUAL:
		return "GREATERTHANOREQUAL"
	case OP_CODE__HASH160:
		return "HASH160"
	case OP_CODE__HASH256:
		return "HASH256"
	case OP_CODE__IF:
		return "IF"
	case OP_CODE__IFDUP:
		return "IFDUP"
	case OP_CODE__INVALIDOPCODE:
		return "INVALIDOPCODE"
	case OP_CODE__INVERT:
		return "INVERT"
	case OP_CODE__LEFT:
		return "LEFT"
	case OP_CODE__LESSTHAN:
		return "LESSTHAN"
	case OP_CODE__LESSTHANOREQUAL:
		return "LESSTHANOREQUAL"
	case OP_CODE__LSHIFT:
		return "LSHIFT"
	case OP_CODE__MAX:
		return "MAX"
	case OP_CODE__MIN:
		return "MIN"
	case OP_CODE__MOD:
		return "MOD"
	case OP_CODE__MUL:
		return "MUL"
	case OP_CODE__NEGATE:
		return "NEGATE"
	case OP_CODE__NIP:
		return "NIP"
	case OP_CODE__NOP:
		return "NOP"
	case OP_CODE__NOP1:
		return "NOP1"
	case OP_CODE__NOP10:
		return "NOP10"
	case OP_CODE__NOP4:
		return "NOP4"
	case OP_CODE__NOP5:
		return "NOP5"
	case OP_CODE__NOP6:
		return "NOP6"
	case OP_CODE__NOP7:
		return "NOP7"
	case OP_CODE__NOP8:
		return "NOP8"
	case OP_CODE__NOP9:
		return "NOP9"
	case OP_CODE__NOT:
		return "NOT"
	case OP_CODE__NOTIF:
		return "NOTIF"
	case OP_CODE__NUMEQUAL:
		return "NUMEQUAL"
	case OP_CODE__NUMEQUALVERIFY:
		return "NUMEQUALVERIFY"
	case OP_CODE__NUMNOTEQUAL:
		return "NUMNOTEQUAL"
	case OP_CODE__OR:
		return "OR"
	case OP_CODE__OVER:
		return "OVER"
	case OP_CODE__PICK:
		return "PICK"
	case OP_CODE__PUBKEY:
		return "PUBKEY"
	case OP_CODE__PUBKEYHASH:
		return "PUBKEYHASH"
	case OP_CODE__PUBKEYS:
		return "PUBKEYS"
	case OP_CODE__PUSHDATA1:
		return "PUSHDATA1"
	case OP_CODE__PUSHDATA2:
		return "PUSHDATA2"
	case OP_CODE__PUSHDATA4:
		return "PUSHDATA4"
	case OP_CODE__RESERVED:
		return "RESERVED"
	case OP_CODE__RESERVED1:
		return "RESERVED1"
	case OP_CODE__RESERVED2:
		return "RESERVED2"
	case OP_CODE__RETURN:
		return "RETURN"
	case OP_CODE__RIGHT:
		return "RIGHT"
	case OP_CODE__RIPEMD160:
		return "RIPEMD160"
	case OP_CODE__ROLL:
		return "ROLL"
	case OP_CODE__ROT:
		return "ROT"
	case OP_CODE__RSHIFT:
		return "RSHIFT"
	case OP_CODE__SHA1:
		return "SHA1"
	case OP_CODE__SHA256:
		return "SHA256"
	case OP_CODE__SIZE:
		return "SIZE"
	case OP_CODE__SMALLINTEGER:
		return "SMALLINTEGER"
	case OP_CODE__SUB:
		return "SUB"
	case OP_CODE__SUBSTR:
		return "SUBSTR"
	case OP_CODE__SWAP:
		return "SWAP"
	case OP_CODE__TOALTSTACK:
		return "TOALTSTACK"
	case OP_CODE__TRUE:
		return "1"
	case OP_CODE__TUCK:
		return "TUCK"
	case OP_CODE__UNKNOWN186:
		return "UNKNOWN186"
	case OP_CODE__UNKNOWN187:
		return "UNKNOWN187"
	case OP_CODE__UNKNOWN188:
		return "UNKNOWN188"
	case OP_CODE__UNKNOWN189:
		return "UNKNOWN189"
	case OP_CODE__UNKNOWN190:
		return "UNKNOWN190"
	case OP_CODE__UNKNOWN191:
		return "UNKNOWN191"
	case OP_CODE__UNKNOWN192:
		return "UNKNOWN192"
	case OP_CODE__UNKNOWN193:
		return "UNKNOWN193"
	case OP_CODE__UNKNOWN194:
		return "UNKNOWN194"
	case OP_CODE__UNKNOWN195:
		return "UNKNOWN195"
	case OP_CODE__UNKNOWN196:
		return "UNKNOWN196"
	case OP_CODE__UNKNOWN197:
		return "UNKNOWN197"
	case OP_CODE__UNKNOWN198:
		return "UNKNOWN198"
	case OP_CODE__UNKNOWN199:
		return "UNKNOWN199"
	case OP_CODE__UNKNOWN200:
		return "UNKNOWN200"
	case OP_CODE__UNKNOWN201:
		return "UNKNOWN201"
	case OP_CODE__UNKNOWN202:
		return "UNKNOWN202"
	case OP_CODE__UNKNOWN203:
		return "UNKNOWN203"
	case OP_CODE__UNKNOWN204:
		return "UNKNOWN204"
	case OP_CODE__UNKNOWN205:
		return "UNKNOWN205"
	case OP_CODE__UNKNOWN206:
		return "UNKNOWN206"
	case OP_CODE__UNKNOWN207:
		return "UNKNOWN207"
	case OP_CODE__UNKNOWN208:
		return "UNKNOWN208"
	case OP_CODE__UNKNOWN209:
		return "UNKNOWN209"
	case OP_CODE__UNKNOWN210:
		return "UNKNOWN210"
	case OP_CODE__UNKNOWN211:
		return "UNKNOWN211"
	case OP_CODE__UNKNOWN212:
		return "UNKNOWN212"
	case OP_CODE__UNKNOWN213:
		return "UNKNOWN213"
	case OP_CODE__UNKNOWN214:
		return "UNKNOWN214"
	case OP_CODE__UNKNOWN215:
		return "UNKNOWN215"
	case OP_CODE__UNKNOWN216:
		return "UNKNOWN216"
	case OP_CODE__UNKNOWN217:
		return "UNKNOWN217"
	case OP_CODE__UNKNOWN218:
		return "UNKNOWN218"
	case OP_CODE__UNKNOWN219:
		return "UNKNOWN219"
	case OP_CODE__UNKNOWN220:
		return "UNKNOWN220"
	case OP_CODE__UNKNOWN221:
		return "UNKNOWN221"
	case OP_CODE__UNKNOWN222:
		return "UNKNOWN222"
	case OP_CODE__UNKNOWN223:
		return "UNKNOWN223"
	case OP_CODE__UNKNOWN224:
		return "UNKNOWN224"
	case OP_CODE__UNKNOWN225:
		return "UNKNOWN225"
	case OP_CODE__UNKNOWN226:
		return "UNKNOWN226"
	case OP_CODE__UNKNOWN227:
		return "UNKNOWN227"
	case OP_CODE__UNKNOWN228:
		return "UNKNOWN228"
	case OP_CODE__UNKNOWN229:
		return "UNKNOWN229"
	case OP_CODE__UNKNOWN230:
		return "UNKNOWN230"
	case OP_CODE__UNKNOWN231:
		return "UNKNOWN231"
	case OP_CODE__UNKNOWN232:
		return "UNKNOWN232"
	case OP_CODE__UNKNOWN233:
		return "UNKNOWN233"
	case OP_CODE__UNKNOWN234:
		return "UNKNOWN234"
	case OP_CODE__UNKNOWN235:
		return "UNKNOWN235"
	case OP_CODE__UNKNOWN236:
		return "UNKNOWN236"
	case OP_CODE__UNKNOWN237:
		return "UNKNOWN237"
	case OP_CODE__UNKNOWN238:
		return "UNKNOWN238"
	case OP_CODE__UNKNOWN239:
		return "UNKNOWN239"
	case OP_CODE__UNKNOWN240:
		return "UNKNOWN240"
	case OP_CODE__UNKNOWN241:
		return "UNKNOWN241"
	case OP_CODE__UNKNOWN242:
		return "UNKNOWN242"
	case OP_CODE__UNKNOWN243:
		return "UNKNOWN243"
	case OP_CODE__UNKNOWN244:
		return "UNKNOWN244"
	case OP_CODE__UNKNOWN245:
		return "UNKNOWN245"
	case OP_CODE__UNKNOWN246:
		return "UNKNOWN246"
	case OP_CODE__UNKNOWN247:
		return "UNKNOWN247"
	case OP_CODE__UNKNOWN248:
		return "UNKNOWN248"
	case OP_CODE__UNKNOWN249:
		return "UNKNOWN249"
	case OP_CODE__UNKNOWN252:
		return "UNKNOWN252"
	case OP_CODE__VER:
		return "VER"
	case OP_CODE__VERIF:
		return "VERIF"
	case OP_CODE__VERIFY:
		return "VERIFY"
	case OP_CODE__VERNOTIF:
		return "VERNOTIF"
	case OP_CODE__WITHIN:
		return "WITHIN"
	case OP_CODE__XOR:
		return "XOR"
	}
	return "UNKNOWN"
}

func (v OpCode) Label() string {
	switch v {
	case OP_CODE__0NOTEQUAL:
		return "146"
	case OP_CODE__10:
		return "90"
	case OP_CODE__11:
		return "91"
	case OP_CODE__12:
		return "92"
	case OP_CODE__13:
		return "93"
	case OP_CODE__14:
		return "94"
	case OP_CODE__15:
		return "95"
	case OP_CODE__16:
		return "96"
	case OP_CODE__1ADD:
		return "139"
	case OP_CODE__1NEGATE:
		return "79"
	case OP_CODE__1SUB:
		return "140"
	case OP_CODE__2:
		return "82"
	case OP_CODE__2DIV:
		return "142"
	case OP_CODE__2DROP:
		return "109"
	case OP_CODE__2DUP:
		return "110"
	case OP_CODE__2MUL:
		return "141"
	case OP_CODE__2OVER:
		return "112"
	case OP_CODE__2ROT:
		return "113"
	case OP_CODE__2SWAP:
		return "114"
	case OP_CODE__3:
		return "83"
	case OP_CODE__3DUP:
		return "111"
	case OP_CODE__4:
		return "84"
	case OP_CODE__5:
		return "85"
	case OP_CODE__6:
		return "86"
	case OP_CODE__7:
		return "87"
	case OP_CODE__8:
		return "88"
	case OP_CODE__9:
		return "89"
	case OP_CODE__ABS:
		return "144"
	case OP_CODE__ADD:
		return "147"
	case OP_CODE__AND:
		return "132"
	case OP_CODE__BOOLAND:
		return "154"
	case OP_CODE__BOOLOR:
		return "155"
	case OP_CODE__CAT:
		return "126"
	case OP_CODE__CHECKLOCKTIMEVERIFY:
		return "177 - AKA OP_NOP2"
	case OP_CODE__CHECKMULTISIG:
		return "174"
	case OP_CODE__CHECKMULTISIGVERIFY:
		return "175"
	case OP_CODE__CHECKSEQUENCEVERIFY:
		return "178 - AKA OP_NOP3"
	case OP_CODE__CHECKSIG:
		return "172"
	case OP_CODE__CHECKSIGVERIFY:
		return "173"
	case OP_CODE__CODESEPARATOR:
		return "171"
	case OP_CODE__DATA_1:
		return "1"
	case OP_CODE__DATA_10:
		return "10"
	case OP_CODE__DATA_11:
		return "11"
	case OP_CODE__DATA_12:
		return "12"
	case OP_CODE__DATA_13:
		return "13"
	case OP_CODE__DATA_14:
		return "14"
	case OP_CODE__DATA_15:
		return "15"
	case OP_CODE__DATA_16:
		return "16"
	case OP_CODE__DATA_17:
		return "17"
	case OP_CODE__DATA_18:
		return "18"
	case OP_CODE__DATA_19:
		return "19"
	case OP_CODE__DATA_2:
		return "2"
	case OP_CODE__DATA_20:
		return "20"
	case OP_CODE__DATA_21:
		return "21"
	case OP_CODE__DATA_22:
		return "22"
	case OP_CODE__DATA_23:
		return "23"
	case OP_CODE__DATA_24:
		return "24"
	case OP_CODE__DATA_25:
		return "25"
	case OP_CODE__DATA_26:
		return "26"
	case OP_CODE__DATA_27:
		return "27"
	case OP_CODE__DATA_28:
		return "28"
	case OP_CODE__DATA_29:
		return "29"
	case OP_CODE__DATA_3:
		return "3"
	case OP_CODE__DATA_30:
		return "30"
	case OP_CODE__DATA_31:
		return "31"
	case OP_CODE__DATA_32:
		return "32"
	case OP_CODE__DATA_33:
		return "33"
	case OP_CODE__DATA_34:
		return "34"
	case OP_CODE__DATA_35:
		return "35"
	case OP_CODE__DATA_36:
		return "36"
	case OP_CODE__DATA_37:
		return "37"
	case OP_CODE__DATA_38:
		return "38"
	case OP_CODE__DATA_39:
		return "39"
	case OP_CODE__DATA_4:
		return "4"
	case OP_CODE__DATA_40:
		return "40"
	case OP_CODE__DATA_41:
		return "41"
	case OP_CODE__DATA_42:
		return "42"
	case OP_CODE__DATA_43:
		return "43"
	case OP_CODE__DATA_44:
		return "44"
	case OP_CODE__DATA_45:
		return "45"
	case OP_CODE__DATA_46:
		return "46"
	case OP_CODE__DATA_47:
		return "47"
	case OP_CODE__DATA_48:
		return "48"
	case OP_CODE__DATA_49:
		return "49"
	case OP_CODE__DATA_5:
		return "5"
	case OP_CODE__DATA_50:
		return "50"
	case OP_CODE__DATA_51:
		return "51"
	case OP_CODE__DATA_52:
		return "52"
	case OP_CODE__DATA_53:
		return "53"
	case OP_CODE__DATA_54:
		return "54"
	case OP_CODE__DATA_55:
		return "55"
	case OP_CODE__DATA_56:
		return "56"
	case OP_CODE__DATA_57:
		return "57"
	case OP_CODE__DATA_58:
		return "58"
	case OP_CODE__DATA_59:
		return "59"
	case OP_CODE__DATA_6:
		return "6"
	case OP_CODE__DATA_60:
		return "60"
	case OP_CODE__DATA_61:
		return "61"
	case OP_CODE__DATA_62:
		return "62"
	case OP_CODE__DATA_63:
		return "63"
	case OP_CODE__DATA_64:
		return "64"
	case OP_CODE__DATA_65:
		return "65"
	case OP_CODE__DATA_66:
		return "66"
	case OP_CODE__DATA_67:
		return "67"
	case OP_CODE__DATA_68:
		return "68"
	case OP_CODE__DATA_69:
		return "69"
	case OP_CODE__DATA_7:
		return "7"
	case OP_CODE__DATA_70:
		return "70"
	case OP_CODE__DATA_71:
		return "71"
	case OP_CODE__DATA_72:
		return "72"
	case OP_CODE__DATA_73:
		return "73"
	case OP_CODE__DATA_74:
		return "74"
	case OP_CODE__DATA_75:
		return "75"
	case OP_CODE__DATA_8:
		return "8"
	case OP_CODE__DATA_9:
		return "9"
	case OP_CODE__DEPTH:
		return "116"
	case OP_CODE__DIV:
		return "150"
	case OP_CODE__DROP:
		return "117"
	case OP_CODE__DUP:
		return "118"
	case OP_CODE__ELSE:
		return "103"
	case OP_CODE__ENDIF:
		return "104"
	case OP_CODE__EQUAL:
		return "135"
	case OP_CODE__EQUALVERIFY:
		return "136"
	case OP_CODE__FALSE:
		return "0 - AKA OP_0"
	case OP_CODE__FROMALTSTACK:
		return "108"
	case OP_CODE__GREATERTHAN:
		return "160"
	case OP_CODE__GREATERTHANOREQUAL:
		return "162"
	case OP_CODE__HASH160:
		return "169"
	case OP_CODE__HASH256:
		return "170"
	case OP_CODE__IF:
		return "99"
	case OP_CODE__IFDUP:
		return "115"
	case OP_CODE__INVALIDOPCODE:
		return "255 - bitcoin core internal"
	case OP_CODE__INVERT:
		return "131"
	case OP_CODE__LEFT:
		return "128"
	case OP_CODE__LESSTHAN:
		return "159"
	case OP_CODE__LESSTHANOREQUAL:
		return "161"
	case OP_CODE__LSHIFT:
		return "152"
	case OP_CODE__MAX:
		return "164"
	case OP_CODE__MIN:
		return "163"
	case OP_CODE__MOD:
		return "151"
	case OP_CODE__MUL:
		return "149"
	case OP_CODE__NEGATE:
		return "143"
	case OP_CODE__NIP:
		return "119"
	case OP_CODE__NOP:
		return "97"
	case OP_CODE__NOP1:
		return "176"
	case OP_CODE__NOP10:
		return "185"
	case OP_CODE__NOP4:
		return "179"
	case OP_CODE__NOP5:
		return "180"
	case OP_CODE__NOP6:
		return "181"
	case OP_CODE__NOP7:
		return "182"
	case OP_CODE__NOP8:
		return "183"
	case OP_CODE__NOP9:
		return "184"
	case OP_CODE__NOT:
		return "145"
	case OP_CODE__NOTIF:
		return "100"
	case OP_CODE__NUMEQUAL:
		return "156"
	case OP_CODE__NUMEQUALVERIFY:
		return "157"
	case OP_CODE__NUMNOTEQUAL:
		return "158"
	case OP_CODE__OR:
		return "133"
	case OP_CODE__OVER:
		return "120"
	case OP_CODE__PICK:
		return "121"
	case OP_CODE__PUBKEY:
		return "254 - bitcoin core internal"
	case OP_CODE__PUBKEYHASH:
		return "253 - bitcoin core internal"
	case OP_CODE__PUBKEYS:
		return "251 - bitcoin core internal"
	case OP_CODE__PUSHDATA1:
		return "76"
	case OP_CODE__PUSHDATA2:
		return "77"
	case OP_CODE__PUSHDATA4:
		return "78"
	case OP_CODE__RESERVED:
		return "80"
	case OP_CODE__RESERVED1:
		return "137"
	case OP_CODE__RESERVED2:
		return "138"
	case OP_CODE__RETURN:
		return "106"
	case OP_CODE__RIGHT:
		return "129"
	case OP_CODE__RIPEMD160:
		return "166"
	case OP_CODE__ROLL:
		return "122"
	case OP_CODE__ROT:
		return "123"
	case OP_CODE__RSHIFT:
		return "153"
	case OP_CODE__SHA1:
		return "167"
	case OP_CODE__SHA256:
		return "168"
	case OP_CODE__SIZE:
		return "130"
	case OP_CODE__SMALLINTEGER:
		return "250 - bitcoin core internal"
	case OP_CODE__SUB:
		return "148"
	case OP_CODE__SUBSTR:
		return "127"
	case OP_CODE__SWAP:
		return "124"
	case OP_CODE__TOALTSTACK:
		return "107"
	case OP_CODE__TRUE:
		return "81"
	case OP_CODE__TUCK:
		return "125"
	case OP_CODE__UNKNOWN186:
		return "186"
	case OP_CODE__UNKNOWN187:
		return "187"
	case OP_CODE__UNKNOWN188:
		return "188"
	case OP_CODE__UNKNOWN189:
		return "189"
	case OP_CODE__UNKNOWN190:
		return "190"
	case OP_CODE__UNKNOWN191:
		return "191"
	case OP_CODE__UNKNOWN192:
		return "192"
	case OP_CODE__UNKNOWN193:
		return "193"
	case OP_CODE__UNKNOWN194:
		return "194"
	case OP_CODE__UNKNOWN195:
		return "195"
	case OP_CODE__UNKNOWN196:
		return "196"
	case OP_CODE__UNKNOWN197:
		return "197"
	case OP_CODE__UNKNOWN198:
		return "198"
	case OP_CODE__UNKNOWN199:
		return "199"
	case OP_CODE__UNKNOWN200:
		return "200"
	case OP_CODE__UNKNOWN201:
		return "201"
	case OP_CODE__UNKNOWN202:
		return "202"
	case OP_CODE__UNKNOWN203:
		return "203"
	case OP_CODE__UNKNOWN204:
		return "204"
	case OP_CODE__UNKNOWN205:
		return "205"
	case OP_CODE__UNKNOWN206:
		return "206"
	case OP_CODE__UNKNOWN207:
		return "207"
	case OP_CODE__UNKNOWN208:
		return "208"
	case OP_CODE__UNKNOWN209:
		return "209"
	case OP_CODE__UNKNOWN210:
		return "210"
	case OP_CODE__UNKNOWN211:
		return "211"
	case OP_CODE__UNKNOWN212:
		return "212"
	case OP_CODE__UNKNOWN213:
		return "213"
	case OP_CODE__UNKNOWN214:
		return "214"
	case OP_CODE__UNKNOWN215:
		return "215"
	case OP_CODE__UNKNOWN216:
		return "216"
	case OP_CODE__UNKNOWN217:
		return "217"
	case OP_CODE__UNKNOWN218:
		return "218"
	case OP_CODE__UNKNOWN219:
		return "219"
	case OP_CODE__UNKNOWN220:
		return "220"
	case OP_CODE__UNKNOWN221:
		return "221"
	case OP_CODE__UNKNOWN222:
		return "222"
	case OP_CODE__UNKNOWN223:
		return "223"
	case OP_CODE__UNKNOWN224:
		return "224"
	case OP_CODE__UNKNOWN225:
		return "225"
	case OP_CODE__UNKNOWN226:
		return "226"
	case OP_CODE__UNKNOWN227:
		return "227"
	case OP_CODE__UNKNOWN228:
		return "228"
	case OP_CODE__UNKNOWN229:
		return "229"
	case OP_CODE__UNKNOWN230:
		return "230"
	case OP_CODE__UNKNOWN231:
		return "231"
	case OP_CODE__UNKNOWN232:
		return "232"
	case OP_CODE__UNKNOWN233:
		return "233"
	case OP_CODE__UNKNOWN234:
		return "234"
	case OP_CODE__UNKNOWN235:
		return "235"
	case OP_CODE__UNKNOWN236:
		return "236"
	case OP_CODE__UNKNOWN237:
		return "237"
	case OP_CODE__UNKNOWN238:
		return "238"
	case OP_CODE__UNKNOWN239:
		return "239"
	case OP_CODE__UNKNOWN240:
		return "240"
	case OP_CODE__UNKNOWN241:
		return "241"
	case OP_CODE__UNKNOWN242:
		return "242"
	case OP_CODE__UNKNOWN243:
		return "243"
	case OP_CODE__UNKNOWN244:
		return "244"
	case OP_CODE__UNKNOWN245:
		return "245"
	case OP_CODE__UNKNOWN246:
		return "246"
	case OP_CODE__UNKNOWN247:
		return "247"
	case OP_CODE__UNKNOWN248:
		return "248"
	case OP_CODE__UNKNOWN249:
		return "249"
	case OP_CODE__UNKNOWN252:
		return "252"
	case OP_CODE__VER:
		return "98"
	case OP_CODE__VERIF:
		return "101"
	case OP_CODE__VERIFY:
		return "105"
	case OP_CODE__VERNOTIF:
		return "102"
	case OP_CODE__WITHIN:
		return "165"
	case OP_CODE__XOR:
		return "134"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*OpCode)(nil)

func (v OpCode) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidOpCode
	}
	return []byte(str), nil
}

func (v *OpCode) UnmarshalText(data []byte) (err error) {
	*v, err = ParseOpCodeFromString(string(bytes.ToUpper(data)))
	return
}
