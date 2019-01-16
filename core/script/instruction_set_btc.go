package script

func newBtcInstructionSet() [256]step {
	return [256]step{
		OP_CODE__FALSE: {
			exec:            opFalse,
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_1: {
			exec:            makeOpPush(1),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_2: {
			exec:            makeOpPush(2),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_3: {
			exec:            makeOpPush(3),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_4: {
			exec:            makeOpPush(4),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_5: {
			exec:            makeOpPush(5),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_6: {
			exec:            makeOpPush(6),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_7: {
			exec:            makeOpPush(7),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_8: {
			exec:            makeOpPush(8),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_9: {
			exec:            makeOpPush(9),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_10: {
			exec:            makeOpPush(10),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_11: {
			exec:            makeOpPush(11),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_12: {
			exec:            makeOpPush(12),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_13: {
			exec:            makeOpPush(13),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_14: {
			exec:            makeOpPush(14),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_15: {
			exec:            makeOpPush(15),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_16: {
			exec:            makeOpPush(16),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_17: {
			exec:            makeOpPush(17),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_18: {
			exec:            makeOpPush(18),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_19: {
			exec:            makeOpPush(19),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_20: {
			exec:            makeOpPush(20),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_21: {
			exec:            makeOpPush(21),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_22: {
			exec:            makeOpPush(22),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_23: {
			exec:            makeOpPush(23),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_24: {
			exec:            makeOpPush(24),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_25: {
			exec:            makeOpPush(25),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_26: {
			exec:            makeOpPush(26),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_27: {
			exec:            makeOpPush(27),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_28: {
			exec:            makeOpPush(28),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_29: {
			exec:            makeOpPush(29),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_30: {
			exec:            makeOpPush(30),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_31: {
			exec:            makeOpPush(31),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_32: {
			exec:            makeOpPush(32),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_33: {
			exec:            makeOpPush(33),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_34: {
			exec:            makeOpPush(34),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_35: {
			exec:            makeOpPush(35),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_36: {
			exec:            makeOpPush(36),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_37: {
			exec:            makeOpPush(37),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_38: {
			exec:            makeOpPush(38),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_39: {
			exec:            makeOpPush(39),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_40: {
			exec:            makeOpPush(40),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_41: {
			exec:            makeOpPush(41),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_42: {
			exec:            makeOpPush(42),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_43: {
			exec:            makeOpPush(43),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_44: {
			exec:            makeOpPush(44),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_45: {
			exec:            makeOpPush(45),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_46: {
			exec:            makeOpPush(46),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_47: {
			exec:            makeOpPush(47),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_48: {
			exec:            makeOpPush(48),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_49: {
			exec:            makeOpPush(49),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_50: {
			exec:            makeOpPush(50),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_51: {
			exec:            makeOpPush(51),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_52: {
			exec:            makeOpPush(52),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_53: {
			exec:            makeOpPush(53),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_54: {
			exec:            makeOpPush(54),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_55: {
			exec:            makeOpPush(55),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_56: {
			exec:            makeOpPush(56),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_57: {
			exec:            makeOpPush(57),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_58: {
			exec:            makeOpPush(58),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_59: {
			exec:            makeOpPush(59),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_60: {
			exec:            makeOpPush(60),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_61: {
			exec:            makeOpPush(61),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_62: {
			exec:            makeOpPush(62),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_63: {
			exec:            makeOpPush(63),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_64: {
			exec:            makeOpPush(64),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_65: {
			exec:            makeOpPush(65),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_66: {
			exec:            makeOpPush(66),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_67: {
			exec:            makeOpPush(67),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_68: {
			exec:            makeOpPush(68),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_69: {
			exec:            makeOpPush(69),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_70: {
			exec:            makeOpPush(70),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_71: {
			exec:            makeOpPush(71),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_72: {
			exec:            makeOpPush(72),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_73: {
			exec:            makeOpPush(73),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_74: {
			exec:            makeOpPush(74),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DATA_75: {
			exec:            makeOpPush(75),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSHDATA1: {
			exec:            makeOpPushData(1),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSHDATA2: {
			exec:            makeOpPushData(2),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSHDATA4: {
			exec:            makeOpPushData(4),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__1NEGATE: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__RESERVED: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__TRUE: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__2: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__3: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__4: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__5: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__6: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__7: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__8: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__9: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__10: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__11: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__12: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__13: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__14: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__15: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__16: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__NOP: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__VER: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__IF: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__NOTIF: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__VERIF: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__VERNOTIF: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__ELSE: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__ENDIF: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__VERIFY: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__RETURN: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__TOALTSTACK: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__FROMALTSTACK: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__2DROP: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__2DUP: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__3DUP: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__2OVER: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__2ROT: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__2SWAP: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__IFDUP: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__DEPTH: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__DROP: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__DUP: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__NIP: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__OVER: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__PICK: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__ROLL: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__ROT: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__SWAP: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__TUCK: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__CAT: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__SUBSTR: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__LEFT: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__RIGHT: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__SIZE: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__INVERT: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__AND: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__OR: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__XOR: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__EQUAL: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__EQUALVERIFY: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__RESERVED1: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__RESERVED2: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__1ADD: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__1SUB: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__2MUL: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__2DIV: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__NEGATE: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__ABS: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__NOT: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__0NOTEQUAL: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__ADD: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__SUB: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__MUL: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__DIV: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__MOD: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__LSHIFT: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__RSHIFT: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__BOOLAND: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__BOOLOR: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__NUMEQUAL: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__NUMEQUALVERIFY: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__NUMNOTEQUAL: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__LESSTHAN: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__GREATERTHAN: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__LESSTHANOREQUAL: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__GREATERTHANOREQUAL: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__MIN: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__MAX: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__WITHIN: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__RIPEMD160: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__SHA1: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__SHA256: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__HASH160: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__HASH256: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__CODESEPARATOR: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__CHECKSIG: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__CHECKSIGVERIFY: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__CHECKMULTISIG: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__CHECKMULTISIGVERIFY: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__NOP1: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__CHECKLOCKTIMEVERIFY: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__CHECKSEQUENCEVERIFY: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__NOP4: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__NOP5: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__NOP6: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__NOP7: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__NOP8: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__NOP9: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__NOP10: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN186: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN187: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN188: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN189: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN190: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN191: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN192: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN193: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN194: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN195: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN196: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN197: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN198: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN199: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN200: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN201: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN202: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN203: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN204: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN205: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN206: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN207: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN208: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN209: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN210: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN211: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN212: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN213: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN214: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN215: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN216: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN217: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN218: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN219: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN220: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN221: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN222: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN223: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN224: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN225: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN226: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN227: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN228: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN229: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN230: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN231: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN232: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN233: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN234: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN235: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN236: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN237: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN238: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN239: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN240: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN241: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN242: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN243: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN244: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN245: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN246: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN247: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN248: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN249: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__SMALLINTEGER: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__PUBKEYS: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__UNKNOWN252: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__PUBKEYHASH: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__PUBKEY: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
		OP_CODE__INVALIDOPCODE: {
			exec:            nil,
			stackValidation: makeStackValidationFunc(0, 0),
			valid:           true,
		},
	}
}
