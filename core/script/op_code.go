package script

//go:generate tools gen enum OpCode
// swagger:enum
type OpCode byte

const (
	OP_CODE__0                   OpCode = 0x00 // 0
	OP_CODE__FALSE               OpCode = 0x00 // 0 - AKA OP_0
	OP_CODE__DATA_1              OpCode = 0x01 // 1
	OP_CODE__DATA_2              OpCode = 0x02 // 2
	OP_CODE__DATA_3              OpCode = 0x03 // 3
	OP_CODE__DATA_4              OpCode = 0x04 // 4
	OP_CODE__DATA_5              OpCode = 0x05 // 5
	OP_CODE__DATA_6              OpCode = 0x06 // 6
	OP_CODE__DATA_7              OpCode = 0x07 // 7
	OP_CODE__DATA_8              OpCode = 0x08 // 8
	OP_CODE__DATA_9              OpCode = 0x09 // 9
	OP_CODE__DATA_10             OpCode = 0x0a // 10
	OP_CODE__DATA_11             OpCode = 0x0b // 11
	OP_CODE__DATA_12             OpCode = 0x0c // 12
	OP_CODE__DATA_13             OpCode = 0x0d // 13
	OP_CODE__DATA_14             OpCode = 0x0e // 14
	OP_CODE__DATA_15             OpCode = 0x0f // 15
	OP_CODE__DATA_16             OpCode = 0x10 // 16
	OP_CODE__DATA_17             OpCode = 0x11 // 17
	OP_CODE__DATA_18             OpCode = 0x12 // 18
	OP_CODE__DATA_19             OpCode = 0x13 // 19
	OP_CODE__DATA_20             OpCode = 0x14 // 20
	OP_CODE__DATA_21             OpCode = 0x15 // 21
	OP_CODE__DATA_22             OpCode = 0x16 // 22
	OP_CODE__DATA_23             OpCode = 0x17 // 23
	OP_CODE__DATA_24             OpCode = 0x18 // 24
	OP_CODE__DATA_25             OpCode = 0x19 // 25
	OP_CODE__DATA_26             OpCode = 0x1a // 26
	OP_CODE__DATA_27             OpCode = 0x1b // 27
	OP_CODE__DATA_28             OpCode = 0x1c // 28
	OP_CODE__DATA_29             OpCode = 0x1d // 29
	OP_CODE__DATA_30             OpCode = 0x1e // 30
	OP_CODE__DATA_31             OpCode = 0x1f // 31
	OP_CODE__DATA_32             OpCode = 0x20 // 32
	OP_CODE__DATA_33             OpCode = 0x21 // 33
	OP_CODE__DATA_34             OpCode = 0x22 // 34
	OP_CODE__DATA_35             OpCode = 0x23 // 35
	OP_CODE__DATA_36             OpCode = 0x24 // 36
	OP_CODE__DATA_37             OpCode = 0x25 // 37
	OP_CODE__DATA_38             OpCode = 0x26 // 38
	OP_CODE__DATA_39             OpCode = 0x27 // 39
	OP_CODE__DATA_40             OpCode = 0x28 // 40
	OP_CODE__DATA_41             OpCode = 0x29 // 41
	OP_CODE__DATA_42             OpCode = 0x2a // 42
	OP_CODE__DATA_43             OpCode = 0x2b // 43
	OP_CODE__DATA_44             OpCode = 0x2c // 44
	OP_CODE__DATA_45             OpCode = 0x2d // 45
	OP_CODE__DATA_46             OpCode = 0x2e // 46
	OP_CODE__DATA_47             OpCode = 0x2f // 47
	OP_CODE__DATA_48             OpCode = 0x30 // 48
	OP_CODE__DATA_49             OpCode = 0x31 // 49
	OP_CODE__DATA_50             OpCode = 0x32 // 50
	OP_CODE__DATA_51             OpCode = 0x33 // 51
	OP_CODE__DATA_52             OpCode = 0x34 // 52
	OP_CODE__DATA_53             OpCode = 0x35 // 53
	OP_CODE__DATA_54             OpCode = 0x36 // 54
	OP_CODE__DATA_55             OpCode = 0x37 // 55
	OP_CODE__DATA_56             OpCode = 0x38 // 56
	OP_CODE__DATA_57             OpCode = 0x39 // 57
	OP_CODE__DATA_58             OpCode = 0x3a // 58
	OP_CODE__DATA_59             OpCode = 0x3b // 59
	OP_CODE__DATA_60             OpCode = 0x3c // 60
	OP_CODE__DATA_61             OpCode = 0x3d // 61
	OP_CODE__DATA_62             OpCode = 0x3e // 62
	OP_CODE__DATA_63             OpCode = 0x3f // 63
	OP_CODE__DATA_64             OpCode = 0x40 // 64
	OP_CODE__DATA_65             OpCode = 0x41 // 65
	OP_CODE__DATA_66             OpCode = 0x42 // 66
	OP_CODE__DATA_67             OpCode = 0x43 // 67
	OP_CODE__DATA_68             OpCode = 0x44 // 68
	OP_CODE__DATA_69             OpCode = 0x45 // 69
	OP_CODE__DATA_70             OpCode = 0x46 // 70
	OP_CODE__DATA_71             OpCode = 0x47 // 71
	OP_CODE__DATA_72             OpCode = 0x48 // 72
	OP_CODE__DATA_73             OpCode = 0x49 // 73
	OP_CODE__DATA_74             OpCode = 0x4a // 74
	OP_CODE__DATA_75             OpCode = 0x4b // 75
	OP_CODE__PUSHDATA1           OpCode = 0x4c // 76
	OP_CODE__PUSHDATA2           OpCode = 0x4d // 77
	OP_CODE__PUSHDATA4           OpCode = 0x4e // 78
	OP_CODE__1NEGATE             OpCode = 0x4f // 79
	OP_CODE__RESERVED            OpCode = 0x50 // 80
	OP_CODE__1                   OpCode = 0x51 // 81 - AKA OP_TRUE
	OP_CODE__TRUE                OpCode = 0x51 // 81
	OP_CODE__2                   OpCode = 0x52 // 82
	OP_CODE__3                   OpCode = 0x53 // 83
	OP_CODE__4                   OpCode = 0x54 // 84
	OP_CODE__5                   OpCode = 0x55 // 85
	OP_CODE__6                   OpCode = 0x56 // 86
	OP_CODE__7                   OpCode = 0x57 // 87
	OP_CODE__8                   OpCode = 0x58 // 88
	OP_CODE__9                   OpCode = 0x59 // 89
	OP_CODE__10                  OpCode = 0x5a // 90
	OP_CODE__11                  OpCode = 0x5b // 91
	OP_CODE__12                  OpCode = 0x5c // 92
	OP_CODE__13                  OpCode = 0x5d // 93
	OP_CODE__14                  OpCode = 0x5e // 94
	OP_CODE__15                  OpCode = 0x5f // 95
	OP_CODE__16                  OpCode = 0x60 // 96
	OP_CODE__NOP                 OpCode = 0x61 // 97
	OP_CODE__VER                 OpCode = 0x62 // 98
	OP_CODE__IF                  OpCode = 0x63 // 99
	OP_CODE__NOTIF               OpCode = 0x64 // 100
	OP_CODE__VERIF               OpCode = 0x65 // 101
	OP_CODE__VERNOTIF            OpCode = 0x66 // 102
	OP_CODE__ELSE                OpCode = 0x67 // 103
	OP_CODE__ENDIF               OpCode = 0x68 // 104
	OP_CODE__VERIFY              OpCode = 0x69 // 105
	OP_CODE__RETURN              OpCode = 0x6a // 106
	OP_CODE__TOALTSTACK          OpCode = 0x6b // 107
	OP_CODE__FROMALTSTACK        OpCode = 0x6c // 108
	OP_CODE__2DROP               OpCode = 0x6d // 109
	OP_CODE__2DUP                OpCode = 0x6e // 110
	OP_CODE__3DUP                OpCode = 0x6f // 111
	OP_CODE__2OVER               OpCode = 0x70 // 112
	OP_CODE__2ROT                OpCode = 0x71 // 113
	OP_CODE__2SWAP               OpCode = 0x72 // 114
	OP_CODE__IFDUP               OpCode = 0x73 // 115
	OP_CODE__DEPTH               OpCode = 0x74 // 116
	OP_CODE__DROP                OpCode = 0x75 // 117
	OP_CODE__DUP                 OpCode = 0x76 // 118
	OP_CODE__NIP                 OpCode = 0x77 // 119
	OP_CODE__OVER                OpCode = 0x78 // 120
	OP_CODE__PICK                OpCode = 0x79 // 121
	OP_CODE__ROLL                OpCode = 0x7a // 122
	OP_CODE__ROT                 OpCode = 0x7b // 123
	OP_CODE__SWAP                OpCode = 0x7c // 124
	OP_CODE__TUCK                OpCode = 0x7d // 125
	OP_CODE__CAT                 OpCode = 0x7e // 126
	OP_CODE__SUBSTR              OpCode = 0x7f // 127
	OP_CODE__LEFT                OpCode = 0x80 // 128
	OP_CODE__RIGHT               OpCode = 0x81 // 129
	OP_CODE__SIZE                OpCode = 0x82 // 130
	OP_CODE__INVERT              OpCode = 0x83 // 131
	OP_CODE__AND                 OpCode = 0x84 // 132
	OP_CODE__OR                  OpCode = 0x85 // 133
	OP_CODE__XOR                 OpCode = 0x86 // 134
	OP_CODE__EQUAL               OpCode = 0x87 // 135
	OP_CODE__EQUALVERIFY         OpCode = 0x88 // 136
	OP_CODE__RESERVED1           OpCode = 0x89 // 137
	OP_CODE__RESERVED2           OpCode = 0x8a // 138
	OP_CODE__1ADD                OpCode = 0x8b // 139
	OP_CODE__1SUB                OpCode = 0x8c // 140
	OP_CODE__2MUL                OpCode = 0x8d // 141
	OP_CODE__2DIV                OpCode = 0x8e // 142
	OP_CODE__NEGATE              OpCode = 0x8f // 143
	OP_CODE__ABS                 OpCode = 0x90 // 144
	OP_CODE__NOT                 OpCode = 0x91 // 145
	OP_CODE__0NOTEQUAL           OpCode = 0x92 // 146
	OP_CODE__ADD                 OpCode = 0x93 // 147
	OP_CODE__SUB                 OpCode = 0x94 // 148
	OP_CODE__MUL                 OpCode = 0x95 // 149
	OP_CODE__DIV                 OpCode = 0x96 // 150
	OP_CODE__MOD                 OpCode = 0x97 // 151
	OP_CODE__LSHIFT              OpCode = 0x98 // 152
	OP_CODE__RSHIFT              OpCode = 0x99 // 153
	OP_CODE__BOOLAND             OpCode = 0x9a // 154
	OP_CODE__BOOLOR              OpCode = 0x9b // 155
	OP_CODE__NUMEQUAL            OpCode = 0x9c // 156
	OP_CODE__NUMEQUALVERIFY      OpCode = 0x9d // 157
	OP_CODE__NUMNOTEQUAL         OpCode = 0x9e // 158
	OP_CODE__LESSTHAN            OpCode = 0x9f // 159
	OP_CODE__GREATERTHAN         OpCode = 0xa0 // 160
	OP_CODE__LESSTHANOREQUAL     OpCode = 0xa1 // 161
	OP_CODE__GREATERTHANOREQUAL  OpCode = 0xa2 // 162
	OP_CODE__MIN                 OpCode = 0xa3 // 163
	OP_CODE__MAX                 OpCode = 0xa4 // 164
	OP_CODE__WITHIN              OpCode = 0xa5 // 165
	OP_CODE__RIPEMD160           OpCode = 0xa6 // 166
	OP_CODE__SHA1                OpCode = 0xa7 // 167
	OP_CODE__SHA256              OpCode = 0xa8 // 168
	OP_CODE__HASH160             OpCode = 0xa9 // 169
	OP_CODE__HASH256             OpCode = 0xaa // 170
	OP_CODE__CODESEPARATOR       OpCode = 0xab // 171
	OP_CODE__CHECKSIG            OpCode = 0xac // 172
	OP_CODE__CHECKSIGVERIFY      OpCode = 0xad // 173
	OP_CODE__CHECKMULTISIG       OpCode = 0xae // 174
	OP_CODE__CHECKMULTISIGVERIFY OpCode = 0xaf // 175
	OP_CODE__NOP1                OpCode = 0xb0 // 176
	OP_CODE__NOP2                OpCode = 0xb1 // 177
	OP_CODE__CHECKLOCKTIMEVERIFY OpCode = 0xb1 // 177 - AKA OP_NOP2
	OP_CODE__NOP3                OpCode = 0xb2 // 178
	OP_CODE__CHECKSEQUENCEVERIFY OpCode = 0xb2 // 178 - AKA OP_NOP3
	OP_CODE__NOP4                OpCode = 0xb3 // 179
	OP_CODE__NOP5                OpCode = 0xb4 // 180
	OP_CODE__NOP6                OpCode = 0xb5 // 181
	OP_CODE__NOP7                OpCode = 0xb6 // 182
	OP_CODE__NOP8                OpCode = 0xb7 // 183
	OP_CODE__NOP9                OpCode = 0xb8 // 184
	OP_CODE__NOP10               OpCode = 0xb9 // 185
	OP_CODE__UNKNOWN186          OpCode = 0xba // 186
	OP_CODE__UNKNOWN187          OpCode = 0xbb // 187
	OP_CODE__UNKNOWN188          OpCode = 0xbc // 188
	OP_CODE__UNKNOWN189          OpCode = 0xbd // 189
	OP_CODE__UNKNOWN190          OpCode = 0xbe // 190
	OP_CODE__UNKNOWN191          OpCode = 0xbf // 191
	OP_CODE__UNKNOWN192          OpCode = 0xc0 // 192
	OP_CODE__UNKNOWN193          OpCode = 0xc1 // 193
	OP_CODE__UNKNOWN194          OpCode = 0xc2 // 194
	OP_CODE__UNKNOWN195          OpCode = 0xc3 // 195
	OP_CODE__UNKNOWN196          OpCode = 0xc4 // 196
	OP_CODE__UNKNOWN197          OpCode = 0xc5 // 197
	OP_CODE__UNKNOWN198          OpCode = 0xc6 // 198
	OP_CODE__UNKNOWN199          OpCode = 0xc7 // 199
	OP_CODE__UNKNOWN200          OpCode = 0xc8 // 200
	OP_CODE__UNKNOWN201          OpCode = 0xc9 // 201
	OP_CODE__UNKNOWN202          OpCode = 0xca // 202
	OP_CODE__UNKNOWN203          OpCode = 0xcb // 203
	OP_CODE__UNKNOWN204          OpCode = 0xcc // 204
	OP_CODE__UNKNOWN205          OpCode = 0xcd // 205
	OP_CODE__UNKNOWN206          OpCode = 0xce // 206
	OP_CODE__UNKNOWN207          OpCode = 0xcf // 207
	OP_CODE__UNKNOWN208          OpCode = 0xd0 // 208
	OP_CODE__UNKNOWN209          OpCode = 0xd1 // 209
	OP_CODE__UNKNOWN210          OpCode = 0xd2 // 210
	OP_CODE__UNKNOWN211          OpCode = 0xd3 // 211
	OP_CODE__UNKNOWN212          OpCode = 0xd4 // 212
	OP_CODE__UNKNOWN213          OpCode = 0xd5 // 213
	OP_CODE__UNKNOWN214          OpCode = 0xd6 // 214
	OP_CODE__UNKNOWN215          OpCode = 0xd7 // 215
	OP_CODE__UNKNOWN216          OpCode = 0xd8 // 216
	OP_CODE__UNKNOWN217          OpCode = 0xd9 // 217
	OP_CODE__UNKNOWN218          OpCode = 0xda // 218
	OP_CODE__UNKNOWN219          OpCode = 0xdb // 219
	OP_CODE__UNKNOWN220          OpCode = 0xdc // 220
	OP_CODE__UNKNOWN221          OpCode = 0xdd // 221
	OP_CODE__UNKNOWN222          OpCode = 0xde // 222
	OP_CODE__UNKNOWN223          OpCode = 0xdf // 223
	OP_CODE__UNKNOWN224          OpCode = 0xe0 // 224
	OP_CODE__UNKNOWN225          OpCode = 0xe1 // 225
	OP_CODE__UNKNOWN226          OpCode = 0xe2 // 226
	OP_CODE__UNKNOWN227          OpCode = 0xe3 // 227
	OP_CODE__UNKNOWN228          OpCode = 0xe4 // 228
	OP_CODE__UNKNOWN229          OpCode = 0xe5 // 229
	OP_CODE__UNKNOWN230          OpCode = 0xe6 // 230
	OP_CODE__UNKNOWN231          OpCode = 0xe7 // 231
	OP_CODE__UNKNOWN232          OpCode = 0xe8 // 232
	OP_CODE__UNKNOWN233          OpCode = 0xe9 // 233
	OP_CODE__UNKNOWN234          OpCode = 0xea // 234
	OP_CODE__UNKNOWN235          OpCode = 0xeb // 235
	OP_CODE__UNKNOWN236          OpCode = 0xec // 236
	OP_CODE__UNKNOWN237          OpCode = 0xed // 237
	OP_CODE__UNKNOWN238          OpCode = 0xee // 238
	OP_CODE__UNKNOWN239          OpCode = 0xef // 239
	OP_CODE__UNKNOWN240          OpCode = 0xf0 // 240
	OP_CODE__UNKNOWN241          OpCode = 0xf1 // 241
	OP_CODE__UNKNOWN242          OpCode = 0xf2 // 242
	OP_CODE__UNKNOWN243          OpCode = 0xf3 // 243
	OP_CODE__UNKNOWN244          OpCode = 0xf4 // 244
	OP_CODE__UNKNOWN245          OpCode = 0xf5 // 245
	OP_CODE__UNKNOWN246          OpCode = 0xf6 // 246
	OP_CODE__UNKNOWN247          OpCode = 0xf7 // 247
	OP_CODE__UNKNOWN248          OpCode = 0xf8 // 248
	OP_CODE__UNKNOWN249          OpCode = 0xf9 // 249
	OP_CODE__SMALLINTEGER        OpCode = 0xfa // 250 - bitcoin core internal
	OP_CODE__PUBKEYS             OpCode = 0xfb // 251 - bitcoin core internal
	OP_CODE__UNKNOWN252          OpCode = 0xfc // 252
	OP_CODE__PUBKEYHASH          OpCode = 0xfd // 253 - bitcoin core internal
	OP_CODE__PUBKEY              OpCode = 0xfe // 254 - bitcoin core internal
	OP_CODE__INVALIDOPCODE       OpCode = 0xff // 255 - bitcoin core internal
)