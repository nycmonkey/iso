// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: country.proto

package iso // import "github.com/nycmonkey/iso"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// ISO Country codes, plus a couple of FactSet additions
type Country int32

const (
	UNKNOWN_COUNTRY Country = 0
	// Supranational
	SUPRANATIONAL Country = 1
	// N/A
	NOT_APPLICABLE Country = 2
	// Andorra
	AD Country = 3
	// United Arab Emirates
	AE Country = 4
	// Afghanistan
	AF Country = 5
	// Antigua and Barbuda
	AG Country = 6
	// Anguilla
	AI Country = 7
	// Albania
	AL Country = 8
	// Armenia
	AM Country = 9
	// Netherlands Antilles
	AN Country = 10
	// Angola
	AO Country = 11
	// Asia Pacific Region
	AP Country = 12
	// Antarctica
	AQ Country = 13
	// Argentina
	AR Country = 14
	// American Samoa
	AS Country = 15
	// Austria
	AT Country = 16
	// Australia
	AU Country = 17
	// Aruba
	AW Country = 18
	// Aland Islands
	AX Country = 19
	// Azerbaijan
	AZ Country = 20
	// Bosnia and Herzegovina
	BA Country = 21
	// Barbados
	BB Country = 22
	// Bangladesh
	BD Country = 23
	// Belgium
	BE Country = 24
	// Burkina Faso
	BF Country = 25
	// Bulgaria
	BG Country = 26
	// Bahrain
	BH Country = 27
	// Burundi
	BI Country = 28
	// Benin
	BJ Country = 29
	// Saint Barthelemy
	BL Country = 30
	// Bermuda
	BM Country = 31
	// Brunei
	BN Country = 32
	// Bolivia
	BO Country = 33
	// Bonaire, St. Eust., Saba
	BQ Country = 34
	// Brazil
	BR Country = 35
	// Bahamas
	BS Country = 36
	// Bhutan
	BT Country = 37
	// Bouvet Island
	BV Country = 38
	// Botswana
	BW Country = 39
	// Belarus
	BY Country = 40
	// Belize
	BZ Country = 41
	// Canada
	CA Country = 42
	// Cocos (Keeling) Islands
	CC Country = 43
	// Congo (Dem. Rep. of the)
	CD Country = 44
	// Central African Republic
	CF Country = 45
	// Congo (Rep. of)
	CG Country = 46
	// Switzerland
	CH Country = 47
	// Ivory Coast
	CI Country = 48
	// Cook Islands
	CK Country = 49
	// Chile
	CL Country = 50
	// Cameroon
	CM Country = 51
	// China
	CN Country = 52
	// Colombia
	CO Country = 53
	// Costa Rica
	CR Country = 54
	// Serbia & Montenegro
	CS Country = 55
	// Cuba
	CU Country = 56
	// Cape Verde
	CV Country = 57
	// Curacao
	CW Country = 58
	// Christmas Island
	CX Country = 59
	// Cyprus
	CY Country = 60
	// Czech Republic
	CZ Country = 61
	// East Germany
	DD Country = 62
	// Germany
	DE Country = 63
	// Djibouti
	DJ Country = 64
	// Denmark
	DK Country = 65
	// Dominica
	DM Country = 66
	// Dominican Republic
	DO Country = 67
	// Algeria
	DZ Country = 68
	// Ecuador
	EC Country = 69
	// Estonia
	EE Country = 70
	// Egypt
	EG Country = 71
	// Western Sahara
	EH Country = 72
	// Eritrea
	ER Country = 73
	// Spain
	ES Country = 74
	// Ethiopia
	ET Country = 75
	// Europe Region
	EU Country = 76
	// Finland
	FI Country = 77
	// Fiji
	FJ Country = 78
	// Falkland Islands
	FK Country = 79
	// Micronesia
	FM Country = 80
	// Faroe Islands
	FO Country = 81
	// France
	FR Country = 82
	// Gabon
	GA Country = 83
	// United Kingdom
	GB Country = 84
	// Grenada
	GD Country = 85
	// Georgia
	GE Country = 86
	// French Guiana
	GF Country = 87
	// Guernsey
	GG Country = 88
	// Ghana
	GH Country = 89
	// Gibraltar
	GI Country = 90
	// Greenland
	GL Country = 91
	// Gambia
	GM Country = 92
	// Guinea
	GN Country = 93
	// Guadeloupe
	GP Country = 94
	// Equatorial Guinea
	GQ Country = 95
	// Greece
	GR Country = 96
	// So Georgia & So Sand Isl
	GS Country = 97
	// Guatemala
	GT Country = 98
	// Guam
	GU Country = 99
	// Guinea-Bissau
	GW Country = 100
	// Guyana
	GY Country = 101
	// Hong Kong
	HK Country = 102
	// Heard & McDonald Islands
	HM Country = 103
	// Honduras
	HN Country = 104
	// Croatia
	HR Country = 105
	// Haiti
	HT Country = 106
	// Hungary
	HU Country = 107
	// Indonesia
	ID Country = 108
	// Ireland
	IE Country = 109
	// Israel
	IL Country = 110
	// Isle of Man
	IM Country = 111
	// India
	IN Country = 112
	// Brit. Indian Ocean Terr.
	IO Country = 113
	// Iraq
	IQ Country = 114
	// Iran
	IR Country = 115
	// Iceland
	IS Country = 116
	// Italy
	IT Country = 117
	// Jersey
	JE Country = 118
	// Jamaica
	JM Country = 119
	// Jordan
	JO Country = 120
	// Japan
	JP Country = 121
	// Kenya
	KE Country = 122
	// Kyrgyzstan
	KG Country = 123
	// Cambodia
	KH Country = 124
	// Kiribati
	KI Country = 125
	// Comoros
	KM Country = 126
	// Saint Kitts and Nevis
	KN Country = 127
	// North Korea
	KP Country = 128
	// South Korea
	KR Country = 129
	// Kuwait
	KW Country = 130
	// Cayman Islands
	KY Country = 131
	// Kazakhstan
	KZ Country = 132
	// Laos
	LA Country = 133
	// Lebanon
	LB Country = 134
	// St. Lucia
	LC Country = 135
	// Liechtenstein
	LI Country = 136
	// Sri Lanka
	LK Country = 137
	// Liberia
	LR Country = 138
	// Lesotho
	LS Country = 139
	// Lithuania
	LT Country = 140
	// Luxembourg
	LU Country = 141
	// Latvia
	LV Country = 142
	// Libya
	LY Country = 143
	// Morocco
	MA Country = 144
	// Monaco
	MC Country = 145
	// Moldova
	MD Country = 146
	// Montenegro
	ME Country = 147
	// Saint Martin
	MF Country = 148
	// Madagascar
	MG Country = 149
	// Marshall Islands
	MH Country = 150
	// Macedonia
	MK Country = 151
	// Mali
	ML Country = 152
	// Myanmar
	MM Country = 153
	// Mongolia
	MN Country = 154
	// Macau
	MO Country = 155
	// Marianas
	MP Country = 156
	// Martinique
	MQ Country = 157
	// Mauritania
	MR Country = 158
	// Montserrat
	MS Country = 159
	// Malta
	MT Country = 160
	// Mauritius
	MU Country = 161
	// Maldives
	MV Country = 162
	// Malawi
	MW Country = 163
	// Mexico
	MX Country = 164
	// Malaysia
	MY Country = 165
	// Mozambique
	MZ Country = 166
	// Namibia
	NA Country = 167
	// New Caledonia
	NC Country = 168
	// Niger
	NE Country = 169
	// Norfolk Island
	NF Country = 170
	// Nigeria
	NG Country = 171
	// Nicaragua
	NI Country = 172
	// Netherlands
	NL Country = 173
	// Norway
	NO Country = 174
	// Nepal
	NP Country = 175
	// Nauru
	NR Country = 176
	// Niue
	NU Country = 177
	// New Zealand
	NZ Country = 178
	// Oman
	OM Country = 179
	// Panama
	PA Country = 180
	// Peru
	PE Country = 181
	// French Polynesia
	PF Country = 182
	// Papua New Guinea
	PG Country = 183
	// Philippines
	PH Country = 184
	// Pakistan
	PK Country = 185
	// Poland
	PL Country = 186
	// Saint Pierre & Miquelon
	PM Country = 187
	// Pitcairn
	PN Country = 188
	// Puerto Rico
	PR Country = 189
	// Palestine
	PS Country = 190
	// Portugal
	PT Country = 191
	// Palau
	PW Country = 192
	// Paraguay
	PY Country = 193
	// Qatar
	QA Country = 194
	// La Reunion
	RE Country = 195
	// Romania
	RO Country = 196
	// Serbia
	RS Country = 197
	// Russian Federation
	RU Country = 198
	// Rwanda
	RW Country = 199
	// Saudi Arabia
	SA Country = 200
	// Solomon Islands
	SB Country = 201
	// Seychelles
	SC Country = 202
	// Sudan
	SD Country = 203
	// Sweden
	SE Country = 204
	// Singapore
	SG Country = 205
	// Saint Helena
	SH Country = 206
	// Slovenia
	SI Country = 207
	// Svalbard and Jan Mayen
	SJ Country = 208
	// Slovakia
	SK Country = 209
	// Sierra Leone
	SL Country = 210
	// San Marino
	SM Country = 211
	// Senegal
	SN Country = 212
	// Somalia
	SO Country = 213
	// Suriname
	SR Country = 214
	// South Sudan
	SS Country = 215
	// Sao Tome and Principe
	ST Country = 216
	// Soviet Union
	SU Country = 217
	// El Salvador
	SV Country = 218
	// Sint Marten
	SX Country = 219
	// Syria
	SY Country = 220
	// Swaziland
	SZ Country = 221
	// Turks and Caicos Islands
	TC Country = 222
	// Chad
	TD Country = 223
	// French Southern Terr.
	TF Country = 224
	// Togo
	TG Country = 225
	// Thailand
	TH Country = 226
	// Tajikistan
	TJ Country = 227
	// Tokelau
	TK Country = 228
	// Timor-Leste
	TL Country = 229
	// Turkmenistan
	TM Country = 230
	// Tunisia
	TN Country = 231
	// Tonga
	TO Country = 232
	// Turkey
	TR Country = 233
	// Trinidad and Tobago
	TT Country = 234
	// Tuvalu
	TV Country = 235
	// Taiwan
	TW Country = 236
	// Tanzania
	TZ Country = 237
	// Ukraine
	UA Country = 238
	// Uganda
	UG Country = 239
	// U.S. Minor Outlying Isl.
	UM Country = 240
	// United States
	US Country = 241
	// Uruguay
	UY Country = 242
	// Uzbekistan
	UZ Country = 243
	// Vatican City
	VA Country = 244
	// St Vincent & Grenadines
	VC Country = 245
	// North Vietnam
	VD Country = 246
	// Venezuela
	VE Country = 247
	// British Virgin Islands
	VG Country = 248
	// U.S. Virgin Islands
	VI Country = 249
	// Vietnam
	VN Country = 250
	// Vanuatu
	VU Country = 251
	// Wallis & Futuna
	WF Country = 252
	// Samoa
	WS Country = 253
	// Kosovo
	XK Country = 254
	// Afr/Mid-East Unallocated
	XP Country = 255
	// Asia/Pacific Unallocated
	XR Country = 256
	// Americas Unallocated
	XS Country = 257
	// Europe Unallocated
	XT Country = 258
	// Zanzibar
	XU Country = 259
	// Tanganyika
	XV Country = 260
	// Non-Geographic Revenue
	XW Country = 261
	// Non-Operating Country
	XY Country = 262
	// Unspecified
	XZ Country = 263
	// South Yemen
	YD Country = 264
	// Yemen
	YE Country = 265
	// Mayotte
	YT Country = 266
	// Yugoslavia
	YU Country = 267
	// South Africa
	ZA Country = 268
	// Zambia
	ZM Country = 269
	// Zimbabwe
	ZW Country = 270
)

var Country_name = map[int32]string{
	0:   "UNKNOWN_COUNTRY",
	1:   "SUPRANATIONAL",
	2:   "NOT_APPLICABLE",
	3:   "AD",
	4:   "AE",
	5:   "AF",
	6:   "AG",
	7:   "AI",
	8:   "AL",
	9:   "AM",
	10:  "AN",
	11:  "AO",
	12:  "AP",
	13:  "AQ",
	14:  "AR",
	15:  "AS",
	16:  "AT",
	17:  "AU",
	18:  "AW",
	19:  "AX",
	20:  "AZ",
	21:  "BA",
	22:  "BB",
	23:  "BD",
	24:  "BE",
	25:  "BF",
	26:  "BG",
	27:  "BH",
	28:  "BI",
	29:  "BJ",
	30:  "BL",
	31:  "BM",
	32:  "BN",
	33:  "BO",
	34:  "BQ",
	35:  "BR",
	36:  "BS",
	37:  "BT",
	38:  "BV",
	39:  "BW",
	40:  "BY",
	41:  "BZ",
	42:  "CA",
	43:  "CC",
	44:  "CD",
	45:  "CF",
	46:  "CG",
	47:  "CH",
	48:  "CI",
	49:  "CK",
	50:  "CL",
	51:  "CM",
	52:  "CN",
	53:  "CO",
	54:  "CR",
	55:  "CS",
	56:  "CU",
	57:  "CV",
	58:  "CW",
	59:  "CX",
	60:  "CY",
	61:  "CZ",
	62:  "DD",
	63:  "DE",
	64:  "DJ",
	65:  "DK",
	66:  "DM",
	67:  "DO",
	68:  "DZ",
	69:  "EC",
	70:  "EE",
	71:  "EG",
	72:  "EH",
	73:  "ER",
	74:  "ES",
	75:  "ET",
	76:  "EU",
	77:  "FI",
	78:  "FJ",
	79:  "FK",
	80:  "FM",
	81:  "FO",
	82:  "FR",
	83:  "GA",
	84:  "GB",
	85:  "GD",
	86:  "GE",
	87:  "GF",
	88:  "GG",
	89:  "GH",
	90:  "GI",
	91:  "GL",
	92:  "GM",
	93:  "GN",
	94:  "GP",
	95:  "GQ",
	96:  "GR",
	97:  "GS",
	98:  "GT",
	99:  "GU",
	100: "GW",
	101: "GY",
	102: "HK",
	103: "HM",
	104: "HN",
	105: "HR",
	106: "HT",
	107: "HU",
	108: "ID",
	109: "IE",
	110: "IL",
	111: "IM",
	112: "IN",
	113: "IO",
	114: "IQ",
	115: "IR",
	116: "IS",
	117: "IT",
	118: "JE",
	119: "JM",
	120: "JO",
	121: "JP",
	122: "KE",
	123: "KG",
	124: "KH",
	125: "KI",
	126: "KM",
	127: "KN",
	128: "KP",
	129: "KR",
	130: "KW",
	131: "KY",
	132: "KZ",
	133: "LA",
	134: "LB",
	135: "LC",
	136: "LI",
	137: "LK",
	138: "LR",
	139: "LS",
	140: "LT",
	141: "LU",
	142: "LV",
	143: "LY",
	144: "MA",
	145: "MC",
	146: "MD",
	147: "ME",
	148: "MF",
	149: "MG",
	150: "MH",
	151: "MK",
	152: "ML",
	153: "MM",
	154: "MN",
	155: "MO",
	156: "MP",
	157: "MQ",
	158: "MR",
	159: "MS",
	160: "MT",
	161: "MU",
	162: "MV",
	163: "MW",
	164: "MX",
	165: "MY",
	166: "MZ",
	167: "NA",
	168: "NC",
	169: "NE",
	170: "NF",
	171: "NG",
	172: "NI",
	173: "NL",
	174: "NO",
	175: "NP",
	176: "NR",
	177: "NU",
	178: "NZ",
	179: "OM",
	180: "PA",
	181: "PE",
	182: "PF",
	183: "PG",
	184: "PH",
	185: "PK",
	186: "PL",
	187: "PM",
	188: "PN",
	189: "PR",
	190: "PS",
	191: "PT",
	192: "PW",
	193: "PY",
	194: "QA",
	195: "RE",
	196: "RO",
	197: "RS",
	198: "RU",
	199: "RW",
	200: "SA",
	201: "SB",
	202: "SC",
	203: "SD",
	204: "SE",
	205: "SG",
	206: "SH",
	207: "SI",
	208: "SJ",
	209: "SK",
	210: "SL",
	211: "SM",
	212: "SN",
	213: "SO",
	214: "SR",
	215: "SS",
	216: "ST",
	217: "SU",
	218: "SV",
	219: "SX",
	220: "SY",
	221: "SZ",
	222: "TC",
	223: "TD",
	224: "TF",
	225: "TG",
	226: "TH",
	227: "TJ",
	228: "TK",
	229: "TL",
	230: "TM",
	231: "TN",
	232: "TO",
	233: "TR",
	234: "TT",
	235: "TV",
	236: "TW",
	237: "TZ",
	238: "UA",
	239: "UG",
	240: "UM",
	241: "US",
	242: "UY",
	243: "UZ",
	244: "VA",
	245: "VC",
	246: "VD",
	247: "VE",
	248: "VG",
	249: "VI",
	250: "VN",
	251: "VU",
	252: "WF",
	253: "WS",
	254: "XK",
	255: "XP",
	256: "XR",
	257: "XS",
	258: "XT",
	259: "XU",
	260: "XV",
	261: "XW",
	262: "XY",
	263: "XZ",
	264: "YD",
	265: "YE",
	266: "YT",
	267: "YU",
	268: "ZA",
	269: "ZM",
	270: "ZW",
}
var Country_value = map[string]int32{
	"UNKNOWN_COUNTRY": 0,
	"SUPRANATIONAL":   1,
	"NOT_APPLICABLE":  2,
	"AD":              3,
	"AE":              4,
	"AF":              5,
	"AG":              6,
	"AI":              7,
	"AL":              8,
	"AM":              9,
	"AN":              10,
	"AO":              11,
	"AP":              12,
	"AQ":              13,
	"AR":              14,
	"AS":              15,
	"AT":              16,
	"AU":              17,
	"AW":              18,
	"AX":              19,
	"AZ":              20,
	"BA":              21,
	"BB":              22,
	"BD":              23,
	"BE":              24,
	"BF":              25,
	"BG":              26,
	"BH":              27,
	"BI":              28,
	"BJ":              29,
	"BL":              30,
	"BM":              31,
	"BN":              32,
	"BO":              33,
	"BQ":              34,
	"BR":              35,
	"BS":              36,
	"BT":              37,
	"BV":              38,
	"BW":              39,
	"BY":              40,
	"BZ":              41,
	"CA":              42,
	"CC":              43,
	"CD":              44,
	"CF":              45,
	"CG":              46,
	"CH":              47,
	"CI":              48,
	"CK":              49,
	"CL":              50,
	"CM":              51,
	"CN":              52,
	"CO":              53,
	"CR":              54,
	"CS":              55,
	"CU":              56,
	"CV":              57,
	"CW":              58,
	"CX":              59,
	"CY":              60,
	"CZ":              61,
	"DD":              62,
	"DE":              63,
	"DJ":              64,
	"DK":              65,
	"DM":              66,
	"DO":              67,
	"DZ":              68,
	"EC":              69,
	"EE":              70,
	"EG":              71,
	"EH":              72,
	"ER":              73,
	"ES":              74,
	"ET":              75,
	"EU":              76,
	"FI":              77,
	"FJ":              78,
	"FK":              79,
	"FM":              80,
	"FO":              81,
	"FR":              82,
	"GA":              83,
	"GB":              84,
	"GD":              85,
	"GE":              86,
	"GF":              87,
	"GG":              88,
	"GH":              89,
	"GI":              90,
	"GL":              91,
	"GM":              92,
	"GN":              93,
	"GP":              94,
	"GQ":              95,
	"GR":              96,
	"GS":              97,
	"GT":              98,
	"GU":              99,
	"GW":              100,
	"GY":              101,
	"HK":              102,
	"HM":              103,
	"HN":              104,
	"HR":              105,
	"HT":              106,
	"HU":              107,
	"ID":              108,
	"IE":              109,
	"IL":              110,
	"IM":              111,
	"IN":              112,
	"IO":              113,
	"IQ":              114,
	"IR":              115,
	"IS":              116,
	"IT":              117,
	"JE":              118,
	"JM":              119,
	"JO":              120,
	"JP":              121,
	"KE":              122,
	"KG":              123,
	"KH":              124,
	"KI":              125,
	"KM":              126,
	"KN":              127,
	"KP":              128,
	"KR":              129,
	"KW":              130,
	"KY":              131,
	"KZ":              132,
	"LA":              133,
	"LB":              134,
	"LC":              135,
	"LI":              136,
	"LK":              137,
	"LR":              138,
	"LS":              139,
	"LT":              140,
	"LU":              141,
	"LV":              142,
	"LY":              143,
	"MA":              144,
	"MC":              145,
	"MD":              146,
	"ME":              147,
	"MF":              148,
	"MG":              149,
	"MH":              150,
	"MK":              151,
	"ML":              152,
	"MM":              153,
	"MN":              154,
	"MO":              155,
	"MP":              156,
	"MQ":              157,
	"MR":              158,
	"MS":              159,
	"MT":              160,
	"MU":              161,
	"MV":              162,
	"MW":              163,
	"MX":              164,
	"MY":              165,
	"MZ":              166,
	"NA":              167,
	"NC":              168,
	"NE":              169,
	"NF":              170,
	"NG":              171,
	"NI":              172,
	"NL":              173,
	"NO":              174,
	"NP":              175,
	"NR":              176,
	"NU":              177,
	"NZ":              178,
	"OM":              179,
	"PA":              180,
	"PE":              181,
	"PF":              182,
	"PG":              183,
	"PH":              184,
	"PK":              185,
	"PL":              186,
	"PM":              187,
	"PN":              188,
	"PR":              189,
	"PS":              190,
	"PT":              191,
	"PW":              192,
	"PY":              193,
	"QA":              194,
	"RE":              195,
	"RO":              196,
	"RS":              197,
	"RU":              198,
	"RW":              199,
	"SA":              200,
	"SB":              201,
	"SC":              202,
	"SD":              203,
	"SE":              204,
	"SG":              205,
	"SH":              206,
	"SI":              207,
	"SJ":              208,
	"SK":              209,
	"SL":              210,
	"SM":              211,
	"SN":              212,
	"SO":              213,
	"SR":              214,
	"SS":              215,
	"ST":              216,
	"SU":              217,
	"SV":              218,
	"SX":              219,
	"SY":              220,
	"SZ":              221,
	"TC":              222,
	"TD":              223,
	"TF":              224,
	"TG":              225,
	"TH":              226,
	"TJ":              227,
	"TK":              228,
	"TL":              229,
	"TM":              230,
	"TN":              231,
	"TO":              232,
	"TR":              233,
	"TT":              234,
	"TV":              235,
	"TW":              236,
	"TZ":              237,
	"UA":              238,
	"UG":              239,
	"UM":              240,
	"US":              241,
	"UY":              242,
	"UZ":              243,
	"VA":              244,
	"VC":              245,
	"VD":              246,
	"VE":              247,
	"VG":              248,
	"VI":              249,
	"VN":              250,
	"VU":              251,
	"WF":              252,
	"WS":              253,
	"XK":              254,
	"XP":              255,
	"XR":              256,
	"XS":              257,
	"XT":              258,
	"XU":              259,
	"XV":              260,
	"XW":              261,
	"XY":              262,
	"XZ":              263,
	"YD":              264,
	"YE":              265,
	"YT":              266,
	"YU":              267,
	"ZA":              268,
	"ZM":              269,
	"ZW":              270,
}

func (Country) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_country_1a1b89da6bf4440a, []int{0}
}

func init() {
	proto.RegisterEnum("iso.protobuf.Country", Country_name, Country_value)
}
func (x Country) String() string {
	s, ok := Country_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

func init() { proto.RegisterFile("country.proto", fileDescriptor_country_1a1b89da6bf4440a) }

var fileDescriptor_country_1a1b89da6bf4440a = []byte{
	// 1130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x24, 0xd6, 0xf5, 0x7f, 0x5b, 0x0b,
	0x1d, 0xc6, 0xf1, 0xa6, 0xc0, 0x06, 0x61, 0xf2, 0x59, 0x86, 0x0c, 0x0b, 0x6e, 0x03, 0x36, 0x60,
	0xb8, 0x9f, 0x24, 0x27, 0xd2, 0xe4, 0x9c, 0x93, 0x1e, 0x49, 0x9a, 0x20, 0x83, 0x96, 0x49, 0x19,
	0x6b, 0xc6, 0xd6, 0x02, 0x45, 0xe7, 0xbe, 0x31, 0xdc, 0xdd, 0x86, 0xdb, 0x70, 0xb8, 0x2e, 0xbb,
	0xde, 0xeb, 0xbb, 0xbe, 0xa6, 0x57, 0x76, 0x7d, 0xd7, 0x5d, 0x5e, 0x79, 0xbe, 0x3f, 0x7d, 0xde,
	0xaf, 0xd7, 0xf3, 0x0f, 0x3c, 0xe9, 0xb9, 0x23, 0x9d, 0x89, 0xb1, 0xf1, 0xf5, 0x93, 0x4b, 0xd6,
	0xad, 0xef, 0x8c, 0x77, 0x32, 0x73, 0x46, 0x37, 0x74, 0x8c, 0xc3, 0x13, 0x2b, 0x17, 0x1f, 0xca,
	0xa4, 0x67, 0xe7, 0x6d, 0xcf, 0x2c, 0x4c, 0xcf, 0x4f, 0xfc, 0xaa, 0x1f, 0x34, 0xfd, 0xe5, 0xf9,
	0x20, 0xf1, 0xe3, 0xb0, 0x45, 0x5f, 0x66, 0x41, 0x7a, 0x6e, 0x94, 0xd4, 0x43, 0xc7, 0x77, 0xe2,
	0x4a, 0xe0, 0x3b, 0x35, 0x52, 0x99, 0x4c, 0x7a, 0x9e, 0x1f, 0xc4, 0xcb, 0x9d, 0x7a, 0xbd, 0x56,
	0xc9, 0x3b, 0xb9, 0x9a, 0x4b, 0x7f, 0x66, 0x56, 0xba, 0xdf, 0x29, 0xf0, 0x14, 0xd5, 0xe5, 0xa9,
	0x6a, 0x91, 0xa7, 0xa9, 0x25, 0x66, 0xa9, 0x15, 0x66, 0xab, 0x35, 0x9e, 0xae, 0x7a, 0x3c, 0x43,
	0xf5, 0x49, 0xab, 0x01, 0xcf, 0x54, 0xeb, 0xcc, 0x51, 0x07, 0x99, 0xab, 0x86, 0xcc, 0x53, 0x23,
	0xe6, 0xab, 0x31, 0xa8, 0x09, 0x0b, 0xd4, 0x26, 0x19, 0x75, 0x88, 0x85, 0x6a, 0x9b, 0x67, 0xf5,
	0x9a, 0x73, 0x78, 0xb6, 0x9a, 0xe3, 0x39, 0x6a, 0x81, 0xe7, 0xaa, 0x2e, 0x8b, 0xd4, 0x22, 0xcf,
	0x53, 0x4b, 0x3c, 0x5f, 0x2d, 0xf3, 0x02, 0xb5, 0xc2, 0x0b, 0xd5, 0x01, 0x5e, 0xa4, 0xd6, 0xc8,
	0xaa, 0x1e, 0x2f, 0x56, 0x7d, 0x5e, 0xa2, 0x06, 0xbc, 0x54, 0x1d, 0xe4, 0x65, 0x6a, 0xc8, 0xcb,
	0xd5, 0x88, 0x57, 0xa8, 0x31, 0xaf, 0x54, 0x1b, 0xbc, 0x4a, 0x6d, 0xf2, 0x6a, 0xb5, 0xc5, 0x6b,
	0xd4, 0x36, 0xaf, 0xed, 0x35, 0xef, 0xb0, 0x58, 0xcd, 0xf3, 0x3a, 0xb5, 0xc0, 0xeb, 0xd5, 0x22,
	0x6f, 0x50, 0x4b, 0x2c, 0x51, 0xcb, 0x2c, 0x55, 0x2b, 0xbc, 0x51, 0xad, 0xf2, 0x26, 0xb5, 0xc6,
	0x9b, 0x55, 0x8f, 0x65, 0xaa, 0xcf, 0x5b, 0xd4, 0x80, 0xb7, 0xaa, 0x21, 0x6f, 0x53, 0x23, 0xde,
	0xae, 0x26, 0xbc, 0x43, 0x6d, 0xf0, 0x4e, 0xb5, 0xc9, 0xbb, 0xd4, 0x21, 0xde, 0xad, 0xb6, 0x78,
	0x8f, 0xda, 0xe6, 0xbd, 0xbd, 0x16, 0x0a, 0xbc, 0x4f, 0x75, 0x79, 0xbf, 0x3a, 0xc0, 0x07, 0xd4,
	0x2a, 0x8e, 0xea, 0x91, 0x53, 0x03, 0xf2, 0x6a, 0x9b, 0x42, 0xaf, 0x6e, 0x1e, 0x57, 0x75, 0x29,
	0xaa, 0x25, 0x4a, 0x6a, 0x99, 0xb2, 0x1a, 0x52, 0x51, 0x23, 0x06, 0xd4, 0x98, 0xaa, 0x9a, 0x50,
	0xeb, 0xb5, 0x58, 0xc1, 0x53, 0x07, 0xf0, 0xd5, 0x2a, 0x81, 0xea, 0x51, 0x57, 0x03, 0x06, 0xd5,
	0x90, 0xb0, 0xd7, 0x92, 0x43, 0xa4, 0xe6, 0x88, 0xd5, 0x02, 0x89, 0xea, 0xd2, 0x50, 0x8b, 0x34,
	0xd5, 0x12, 0x43, 0x6a, 0x99, 0x96, 0x5a, 0xa1, 0xad, 0xd6, 0xf8, 0xa0, 0xea, 0xf1, 0x21, 0xd5,
	0xe7, 0xc3, 0x6a, 0x9d, 0x8f, 0xa8, 0x83, 0x2c, 0x57, 0x43, 0x3e, 0xaa, 0x46, 0x7c, 0x4c, 0x8d,
	0x19, 0x56, 0x13, 0x46, 0xd4, 0x26, 0x1f, 0x57, 0x5b, 0xac, 0xe8, 0xb5, 0x5c, 0x65, 0xa5, 0xea,
	0xb1, 0x4a, 0xf5, 0x59, 0xad, 0x86, 0x8c, 0xaa, 0x31, 0x9f, 0x50, 0x13, 0xd6, 0xf4, 0x5a, 0x29,
	0xf0, 0x49, 0xd5, 0x65, 0xad, 0x5a, 0x63, 0x4c, 0xf5, 0xe8, 0xa8, 0x3e, 0xeb, 0xd4, 0x80, 0x4f,
	0xa9, 0x83, 0xac, 0x57, 0x43, 0x36, 0xa8, 0x11, 0xe3, 0x6a, 0xcc, 0x44, 0xaf, 0x03, 0x2e, 0x9f,
	0x56, 0x3d, 0x3e, 0xa3, 0x06, 0x7c, 0x56, 0xad, 0x33, 0xd9, 0x6b, 0xd5, 0xe5, 0x73, 0x6a, 0x89,
	0xcf, 0xab, 0x65, 0xbe, 0xa0, 0x56, 0xf8, 0xa2, 0xea, 0xf1, 0x25, 0xd5, 0xe7, 0xcb, 0x99, 0xd9,
	0xe9, 0xfe, 0x6a, 0x9d, 0x8d, 0x29, 0x21, 0x64, 0x93, 0xa1, 0xc9, 0x66, 0x43, 0x8b, 0x2d, 0x86,
	0x36, 0x5b, 0x85, 0x9a, 0xc3, 0x36, 0x43, 0x8e, 0xed, 0x86, 0x3c, 0x3b, 0x0c, 0x15, 0x76, 0x1a,
	0xaa, 0xec, 0x32, 0x84, 0xec, 0x36, 0x44, 0xec, 0x31, 0xc4, 0xec, 0x35, 0x24, 0xec, 0x33, 0x34,
	0xd8, 0x6f, 0x68, 0xf1, 0x15, 0xc1, 0x73, 0x38, 0x60, 0xc8, 0xf3, 0x55, 0x43, 0x81, 0xaf, 0x19,
	0x5c, 0xbe, 0x6e, 0x28, 0xf2, 0x0d, 0x43, 0x89, 0x6f, 0x1a, 0xca, 0x7c, 0xcb, 0x50, 0xe5, 0xdb,
	0x86, 0x1a, 0xdf, 0x31, 0x78, 0x7c, 0xd7, 0xe0, 0xf3, 0x3d, 0x43, 0xc0, 0xf7, 0x0d, 0x75, 0x7e,
	0x60, 0x18, 0xe4, 0x87, 0x86, 0x90, 0x1f, 0x19, 0x22, 0x7e, 0x6c, 0x88, 0xf9, 0x89, 0x21, 0xe1,
	0xa7, 0x86, 0x06, 0x3f, 0x33, 0x34, 0xf9, 0xb9, 0x61, 0x88, 0x5f, 0x18, 0x5a, 0xfc, 0xd2, 0xd0,
	0xe6, 0x57, 0x82, 0xef, 0xf0, 0x6b, 0x43, 0x9e, 0x83, 0x06, 0x97, 0xdf, 0x18, 0x8a, 0xfc, 0xd6,
	0x50, 0xe2, 0x77, 0x86, 0x0a, 0xbf, 0x37, 0xd4, 0xf8, 0x83, 0x21, 0xe0, 0x8f, 0x86, 0x3a, 0x7f,
	0x32, 0x84, 0xfc, 0xd9, 0x90, 0xf0, 0x17, 0x43, 0x9b, 0xbf, 0x0a, 0x81, 0xc7, 0xdf, 0x84, 0xba,
	0xc3, 0x21, 0x83, 0xcb, 0xdf, 0x0d, 0x45, 0xfe, 0x61, 0x28, 0xf1, 0x4f, 0x43, 0x99, 0x7f, 0x19,
	0xaa, 0xfc, 0xdb, 0x50, 0xe3, 0x3f, 0x06, 0x8f, 0xff, 0x1a, 0x7c, 0xfe, 0x67, 0x08, 0xf9, 0xbf,
	0x21, 0xe2, 0x04, 0x43, 0xcc, 0x89, 0x86, 0x26, 0x27, 0x19, 0x5a, 0x9c, 0x2c, 0x0c, 0x3a, 0x9c,
	0x22, 0x84, 0x2e, 0xa7, 0x1a, 0x02, 0x4e, 0x33, 0x44, 0x9c, 0x6e, 0x48, 0x38, 0xc3, 0xd0, 0xe4,
	0x4c, 0x21, 0x72, 0x38, 0x6c, 0xc8, 0x71, 0x96, 0x21, 0xcf, 0xd9, 0x86, 0x02, 0xe7, 0x18, 0x5c,
	0xce, 0x35, 0x94, 0x38, 0xcf, 0x50, 0xe6, 0x7c, 0x43, 0x85, 0x0b, 0x0c, 0x03, 0x4c, 0x19, 0xaa,
	0x5c, 0x68, 0xa8, 0x71, 0x91, 0xc1, 0xe3, 0x62, 0x83, 0xcf, 0x25, 0x86, 0x80, 0x4b, 0x0d, 0x21,
	0x97, 0x19, 0x22, 0x2e, 0x37, 0xc4, 0x1c, 0x31, 0x24, 0x5c, 0x61, 0x68, 0x70, 0xa5, 0x61, 0x88,
	0xab, 0x0c, 0x2d, 0xae, 0x36, 0xb4, 0xb9, 0x46, 0x88, 0xf3, 0x5c, 0x6b, 0x28, 0x70, 0x9d, 0xa1,
	0xc8, 0x51, 0x43, 0x89, 0x69, 0x43, 0x99, 0xae, 0x61, 0x80, 0x19, 0x43, 0x95, 0xeb, 0x0d, 0x35,
	0x6e, 0x30, 0x78, 0xdc, 0x68, 0xf0, 0xb9, 0xc9, 0x10, 0x70, 0xcc, 0x10, 0x72, 0xb3, 0x21, 0xe6,
	0x16, 0x43, 0x83, 0x5b, 0x0d, 0x4d, 0x6e, 0x33, 0xb4, 0xb9, 0x5d, 0x48, 0x1c, 0xee, 0x30, 0x94,
	0xb8, 0xd3, 0xe0, 0x71, 0xdc, 0x10, 0x71, 0x97, 0xa1, 0xc5, 0xdd, 0x86, 0x36, 0xf7, 0x08, 0x0d,
	0x87, 0x7b, 0x0d, 0x79, 0xee, 0x33, 0x14, 0xb8, 0xdf, 0xe0, 0xf2, 0x80, 0xa1, 0xc4, 0x83, 0x86,
	0x0a, 0x0f, 0x19, 0x7c, 0x1e, 0x36, 0x24, 0x3c, 0x22, 0x34, 0x8b, 0x3c, 0x6a, 0x88, 0x78, 0x4c,
	0x18, 0xaa, 0xf2, 0xb8, 0xa1, 0xce, 0x13, 0x86, 0x90, 0x8d, 0xfd, 0x42, 0xc4, 0x26, 0x43, 0xcc,
	0x66, 0x43, 0xc2, 0x16, 0x43, 0x83, 0xad, 0x86, 0x26, 0xdb, 0x0c, 0x2d, 0xb6, 0x1b, 0xda, 0xec,
	0x10, 0x5a, 0x05, 0x76, 0x1a, 0x5c, 0x76, 0x19, 0x62, 0x76, 0x1b, 0x12, 0xf6, 0x08, 0x6d, 0x87,
	0xbd, 0x06, 0x8f, 0x7d, 0x86, 0x26, 0xfb, 0xfb, 0x73, 0xfe, 0xd4, 0x74, 0xb6, 0xef, 0xc8, 0x74,
	0xb6, 0xef, 0xf8, 0x74, 0x36, 0xb5, 0xb1, 0x9b, 0x4d, 0x1d, 0xec, 0x66, 0x53, 0x87, 0xbb, 0xd9,
	0xd4, 0x54, 0x37, 0x9b, 0x3a, 0xda, 0xcd, 0xa6, 0x8e, 0x75, 0xb3, 0x7d, 0xc7, 0xbb, 0xd9, 0xd4,
	0x81, 0x99, 0x6c, 0xdf, 0xd4, 0x4c, 0xb6, 0xef, 0xc8, 0x4c, 0xb6, 0xaf, 0xbd, 0x68, 0xd5, 0xe8,
	0xf8, 0xea, 0x89, 0xe1, 0x25, 0x23, 0x9d, 0xb5, 0x4b, 0xc7, 0x26, 0x47, 0xd6, 0x76, 0xc6, 0xd6,
	0xac, 0x98, 0x5c, 0x3a, 0xba, 0xa1, 0x33, 0x3c, 0x4b, 0x87, 0x6c, 0xd9, 0x93, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x5c, 0xa7, 0x2f, 0xca, 0xac, 0x09, 0x00, 0x00,
}
