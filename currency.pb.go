// Code generated by protoc-gen-go. DO NOT EDIT.
// source: currency.proto

package iso // import "github.com/nycmonkey/iso"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Currency int32

const (
	// N/A (active)
	Currency_UNKNOWN_CURRENCY Currency = 0
	// Andorran Peseta (inactive)
	Currency_ADP Currency = 1
	// Andorran Franc (inactive)
	Currency_ADR Currency = 2
	// U.A.E. Dirham (active)
	Currency_AED Currency = 3
	// Afghan Afghani (inactive)
	Currency_AFA Currency = 4
	// Afghan Afghani (active)
	Currency_AFN Currency = 5
	// Albanian Old Lek (inactive)
	Currency_ALK Currency = 6
	// Albanian Lek (active)
	Currency_ALL Currency = 7
	// Armenian Dram (active)
	Currency_AMD Currency = 8
	// Netherlands Antillean Guilder (active)
	Currency_ANG Currency = 9
	// Angolan Kwanza (active)
	Currency_AOA Currency = 10
	// Angolan New Kwanza (inactive)
	Currency_AON Currency = 11
	// Angolan Kwanza (inactive)
	Currency_AOR Currency = 12
	// Argentine Austral (inactive)
	Currency_ARA Currency = 13
	// Argentine Peso Ley (inactive)
	Currency_ARL Currency = 14
	// Argentine Peso Moneda Nacional (inactive)
	Currency_ARM Currency = 15
	// Argentine Peso Argentino (inactive)
	Currency_ARP Currency = 16
	// Argentine Peso (active)
	Currency_ARS Currency = 17
	// Austrian Schilling (inactive)
	Currency_ATS Currency = 18
	// Australian Dollar (active)
	Currency_AUD Currency = 19
	// Aruban Guilder (active)
	Currency_AWG Currency = 20
	// Azerbaijanian Manat (inactive)
	Currency_AZM Currency = 21
	// Azerbaijanian New Manat (active)
	Currency_AZN Currency = 22
	// Bosnian Marka (active)
	Currency_BAM Currency = 23
	// Barbados Dollar (active)
	Currency_BBD Currency = 24
	// Bangladeshi Taka (active)
	Currency_BDT Currency = 25
	// Belgian Franc (inactive)
	Currency_BEF Currency = 26
	// Bulgarian Lev (inactive)
	Currency_BGJ Currency = 27
	// Bulgarian Lev (inactive)
	Currency_BGK Currency = 28
	// Bulgarian Lev (inactive)
	Currency_BGL Currency = 29
	// Bulgarian Lev (active)
	Currency_BGN Currency = 30
	// Bahraini Dinar (active)
	Currency_BHD Currency = 31
	// Burundian Franc (active)
	Currency_BIF Currency = 32
	// Bermudian Dollar (active)
	Currency_BMD Currency = 33
	// Brunei Dollar (active)
	Currency_BND Currency = 34
	// Bolivian Boliviano (active)
	Currency_BOB Currency = 35
	// Bolivian Peso (inactive)
	Currency_BOP Currency = 36
	// Bolivian Mvdol (funds code) (active)
	Currency_BOV Currency = 37
	// Brazilian Cruzeiro Novo (inactive)
	Currency_BRB Currency = 38
	// Brazilian Cruzado (inactive)
	Currency_BRC Currency = 39
	// Brazilian Cruzeiro (inactive)
	Currency_BRE Currency = 40
	// Brazilian Real (active)
	Currency_BRL Currency = 41
	// Brazilian Cruzado Novo (inactive)
	Currency_BRN Currency = 42
	// Brazilian Cruzeiro Real (inactive)
	Currency_BRR Currency = 43
	// Brazilian Cruzeiro (inactive)
	Currency_BRZ Currency = 44
	// Bahamian Dollar (active)
	Currency_BSD Currency = 45
	// Bhutan Ngultrum (active)
	Currency_BTN Currency = 46
	// Botswana Pula (active)
	Currency_BWP Currency = 47
	// Belarus Ruble (inactive)
	Currency_BYB Currency = 48
	// Belarus Ruble (active)
	Currency_BYN Currency = 49
	// Belarus Ruble (inactive)
	Currency_BYR Currency = 50
	// Belize Dollar (active)
	Currency_BZD Currency = 51
	// Canadian Dollar (active)
	Currency_CAD Currency = 52
	// Congolese Franc (active)
	Currency_CDF Currency = 53
	// Swiss Franc (active)
	Currency_CHF Currency = 54
	// Chilean Escudo (inactive)
	Currency_CLE Currency = 55
	// Chilean Unidad de Fomento (active)
	Currency_CLF Currency = 56
	// Chilean Peso (active)
	Currency_CLP Currency = 57
	// Chinese Yuan Renminbi (HK) (active)
	Currency_CNH Currency = 58
	// Chinese Yuan Renminbi (active)
	Currency_CNY Currency = 59
	// Colombian Peso (active)
	Currency_COP Currency = 60
	// Colombian Unidad de Valor Real (active)
	Currency_COU Currency = 61
	// Costa Rican Colon (active)
	Currency_CRC Currency = 62
	// Serbian Dinar (inactive)
	Currency_CSD Currency = 63
	// Czechoslovak Koruna (inactive)
	Currency_CSK Currency = 64
	// Cuban Convertible Peso (active)
	Currency_CUC Currency = 65
	// Cuban Peso (active)
	Currency_CUP Currency = 66
	// Cape Verdean Escudo (active)
	Currency_CVE Currency = 67
	// Cypriot Pound (inactive)
	Currency_CYP Currency = 68
	// Czech Koruna (active)
	Currency_CZK Currency = 69
	// East German Mark (inactive)
	Currency_DDM Currency = 70
	// German Deutsche Mark (inactive)
	Currency_DEM Currency = 71
	// Djibouti Franc (active)
	Currency_DJF Currency = 72
	// Danish Krone (active)
	Currency_DKK Currency = 73
	// Dominican Peso (active)
	Currency_DOP Currency = 74
	// Algerian Dinar (active)
	Currency_DZD Currency = 75
	// Ecuadoran Sucre (inactive)
	Currency_ECS Currency = 76
	// Ecuadoran Unidad de Valor Constante (inactive)
	Currency_ECV Currency = 77
	// Estonian Kroon (inactive)
	Currency_EEK Currency = 78
	// Egyptian Pound (active)
	Currency_EGP Currency = 79
	// Eritrean Nakfa (active)
	Currency_ERN Currency = 80
	// Spanish Peseta (inactive)
	Currency_ESP Currency = 81
	// Ethopian Birr (active)
	Currency_ETB Currency = 82
	// Euro (active)
	Currency_EUR Currency = 83
	// Finnish Markka (inactive)
	Currency_FIM Currency = 84
	// Fiji Dollar (active)
	Currency_FJD Currency = 85
	// Falkland Islands Pound (active)
	Currency_FKP Currency = 86
	// French Franc (inactive)
	Currency_FRF Currency = 87
	// British Pound (active)
	Currency_GBP Currency = 88
	// British Pound (Pence) (active)
	Currency_GBX Currency = 89
	// Georgian Lari (inactive)
	Currency_GEK Currency = 90
	// Georgian Lari (active)
	Currency_GEL Currency = 91
	// Ghanaian Cedi (inactive)
	Currency_GHC Currency = 92
	// Ghanaian Cedi (active)
	Currency_GHS Currency = 93
	// Gibraltar Pound (active)
	Currency_GIP Currency = 94
	// Gambian Daasi (active)
	Currency_GMD Currency = 95
	// Guinean Syli (inactive)
	Currency_GNE Currency = 96
	// Guinea Franc (active)
	Currency_GNF Currency = 97
	// Equatorial Guinean Ekwele (inactive)
	Currency_GQE Currency = 98
	// Greek Drachma (inactive)
	Currency_GRD Currency = 99
	// Guatemalan Quetzal (active)
	Currency_GTQ Currency = 100
	// Guinea-Bissau Peso (inactive)
	Currency_GWP Currency = 101
	// Guyanan Dollar (active)
	Currency_GYD Currency = 102
	// Hong Kong Dollar (active)
	Currency_HKD Currency = 103
	// Honduran Lempira (active)
	Currency_HNL Currency = 104
	// Croatian Kuna (active)
	Currency_HRK Currency = 105
	// Haitian Gourde (active)
	Currency_HTG Currency = 106
	// Hungarian Forint (active)
	Currency_HUF Currency = 107
	// Indonesian Rupiah (active)
	Currency_IDR Currency = 108
	// Irish Pound (inactive)
	Currency_IEP Currency = 109
	// Israeli Lira (inactive)
	Currency_ILP Currency = 110
	// Israeli Shekel (inactive)
	Currency_ILR Currency = 111
	// Israeli New Shekel (active)
	Currency_ILS Currency = 112
	// Israeli Agora (active)
	Currency_ILX Currency = 113
	// Indian Rupee (active)
	Currency_INR Currency = 114
	// Iraqi Dinar (active)
	Currency_IQD Currency = 115
	// Iranian Rial (active)
	Currency_IRR Currency = 116
	// Icelandic Old Krona (inactive)
	Currency_ISJ Currency = 117
	// Icelandic Krona (active)
	Currency_ISK Currency = 118
	// Italian Lira (inactive)
	Currency_ITL Currency = 119
	// Jamaican Dollar (active)
	Currency_JMD Currency = 120
	// Jordanian Dinar (active)
	Currency_JOD Currency = 121
	// Japanese Yen (active)
	Currency_JPY Currency = 122
	// Kenyan Shilling (active)
	Currency_KES Currency = 123
	// Kyrgyzstan Som (active)
	Currency_KGS Currency = 124
	// Cambodian Riel (active)
	Currency_KHR Currency = 125
	// Comoros Franc (active)
	Currency_KMF Currency = 126
	// North Korean Won (active)
	Currency_KPW Currency = 127
	// Korean Won (active)
	Currency_KRW Currency = 128
	// Kuwaiti Dinar (active)
	Currency_KWD Currency = 129
	// Caymanian Dollar (active)
	Currency_KYD Currency = 130
	// Kazakhstan Tenge (active)
	Currency_KZT Currency = 131
	// Lao Kip (inactive)
	Currency_LAJ Currency = 132
	// Lao Kip (active)
	Currency_LAK Currency = 133
	// Lebanese Pound (active)
	Currency_LBP Currency = 134
	// Sri Lankan Rupee (active)
	Currency_LKR Currency = 135
	// Liberian Dollar (active)
	Currency_LRD Currency = 136
	// Lesotho Loti (active)
	Currency_LSL Currency = 137
	// Lithuanian Litas (inactive)
	Currency_LTL Currency = 138
	// Luxembourg Franc (inactive)
	Currency_LUF Currency = 139
	// Latvian Lat (inactive)
	Currency_LVL Currency = 140
	// Libyan Dinar (active)
	Currency_LYD Currency = 141
	// Moroccan Dirham (active)
	Currency_MAD Currency = 142
	// Moroccan Franc (inactive)
	Currency_MAF Currency = 143
	// Monegasque Franc (inactive)
	Currency_MCL Currency = 144
	// Moldova Leu (active)
	Currency_MDL Currency = 145
	// Malagasy Ariary (active)
	Currency_MGA Currency = 146
	// Malagasy Franc (inactive)
	Currency_MGF Currency = 147
	// Macedonian Denar (active)
	Currency_MKD Currency = 148
	// Old Macedonian Denar (inactive)
	Currency_MKN Currency = 149
	// Mali Franc (inactive)
	Currency_MLF Currency = 150
	// Myanmar Kyats (active)
	Currency_MMK Currency = 151
	// Mongolia Togrog (active)
	Currency_MNT Currency = 152
	// Macau Pataca (active)
	Currency_MOP Currency = 153
	// Mauritanian Ouguiya (active)
	Currency_MRO Currency = 154
	// Maltese Lira (inactive)
	Currency_MTL Currency = 155
	// Maltese Pound (inactive)
	Currency_MTP Currency = 156
	// Mauritius Rupee (active)
	Currency_MUR Currency = 157
	// Maldive Rupee (inactive)
	Currency_MVQ Currency = 158
	// Maldive Rufiyaa (active)
	Currency_MVR Currency = 159
	// Malawi Kwacha (active)
	Currency_MWK Currency = 160
	// Mexican Peso (active)
	Currency_MXN Currency = 161
	// Mexican Peso (inactive)
	Currency_MXP Currency = 162
	// Mexican Unidad de Inversion (UDI) (active)
	Currency_MXV Currency = 163
	// Malaysian Ringgit (active)
	Currency_MYR Currency = 164
	// Mozambican Metical (inactive)
	Currency_MZM Currency = 165
	// Mozambican Metical (active)
	Currency_MZN Currency = 166
	// Namibian Dollar (active)
	Currency_NAD Currency = 167
	// Newfoundland Dollar (inactive)
	Currency_NFD Currency = 168
	// Nigerian Naira (active)
	Currency_NGN Currency = 169
	// Nicaraguan Cordoba Oro (active)
	Currency_NIO Currency = 170
	// Dutch Guilder (inactive)
	Currency_NLG Currency = 171
	// Norwegian Krone (active)
	Currency_NOK Currency = 172
	// Nepalese Rupee (active)
	Currency_NPR Currency = 173
	// New Zealand Dollar (active)
	Currency_NZD Currency = 174
	// Omani Rial (active)
	Currency_OMR Currency = 175
	// Panamanian Balboa (active)
	Currency_PAB Currency = 176
	// Peruvian Sol (inactive)
	Currency_PEH Currency = 177
	// Peruvian Inti (inactive)
	Currency_PEI Currency = 178
	// Peruvian Nuevo Sol (active)
	Currency_PEN Currency = 179
	// Papua New Guinea Kina (active)
	Currency_PGK Currency = 180
	// Philippine Peso (active)
	Currency_PHP Currency = 181
	// Pakistani Rupee (active)
	Currency_PKR Currency = 182
	// Polish Zloty (active)
	Currency_PLN Currency = 183
	// Polish Zloty (inactive)
	Currency_PLZ Currency = 184
	// Portuguese Escudo (inactive)
	Currency_PTE Currency = 185
	// Paraguay Guarini (active)
	Currency_PYG Currency = 186
	// Qatari Rial (active)
	Currency_QAR Currency = 187
	// Romanian Leu (inactive)
	Currency_ROK Currency = 188
	// Romanian Leu (inactive)
	Currency_ROL Currency = 189
	// Romanian Leu (active)
	Currency_RON Currency = 190
	// Serbian Dinar (active)
	Currency_RSD Currency = 191
	// Russian Ruble (active)
	Currency_RUB Currency = 192
	// Russian Ruble (inactive)
	Currency_RUR Currency = 193
	// Rwandan Franc (active)
	Currency_RWF Currency = 194
	// Saudi Arabian Riyal (active)
	Currency_SAR Currency = 195
	// Solomon Islands Dollar (active)
	Currency_SBD Currency = 196
	// Seychelle Rupee (active)
	Currency_SCR Currency = 197
	// Sudanese Dinar (inactive)
	Currency_SDD Currency = 198
	// Sudanese Pound (active)
	Currency_SDG Currency = 199
	// Swedish Krona (active)
	Currency_SEK Currency = 200
	// Singapore Dollar (active)
	Currency_SGD Currency = 201
	// Saint Helenian Pound (active)
	Currency_SHP Currency = 202
	// Slovenian Tolar (inactive)
	Currency_SIT Currency = 203
	// Slovakian Koruna (inactive)
	Currency_SKK Currency = 204
	// Sierra Leone Leone (active)
	Currency_SLL Currency = 205
	// San Marinese Lira (inactive)
	Currency_SML Currency = 206
	// Somali Shilling (active)
	Currency_SOS Currency = 207
	// Suriname Dollar (active)
	Currency_SRD Currency = 208
	// Suriname Guilder (inactive)
	Currency_SRG Currency = 209
	// South Sudanese Pound (active)
	Currency_SSP Currency = 210
	// Sao Tome and Principe Dobra (active)
	Currency_STD Currency = 211
	// Soviet Union Ruble (inactive)
	Currency_SUR Currency = 212
	// Salvadoran Colon (inactive)
	Currency_SVC Currency = 213
	// Syrian Pound (active)
	Currency_SYP Currency = 214
	// Swaziland Lilangeni (active)
	Currency_SZL Currency = 215
	// Thai Baht (active)
	Currency_THB Currency = 216
	// Tajikistani Ruble (inactive)
	Currency_TJR Currency = 217
	// Tajikistani Somoni (active)
	Currency_TJS Currency = 218
	// Turkmenistani Manat (inactive)
	Currency_TMM Currency = 219
	// Turkmenistani Manat (active)
	Currency_TMT Currency = 220
	// Tunisian Dinar (active)
	Currency_TND Currency = 221
	// Tongan Pa'anga (active)
	Currency_TOP Currency = 222
	// Portuguese Timorese Escudo (inactive)
	Currency_TPE Currency = 223
	// Turkish Lira (inactive)
	Currency_TRL Currency = 224
	// Turkish Lira (active)
	Currency_TRY Currency = 225
	// Trinidad and Tobago Dollar (active)
	Currency_TTD Currency = 226
	// Taiwan Dollar (active)
	Currency_TWD Currency = 227
	// Tanzanian Shilling (active)
	Currency_TZS Currency = 228
	// Ukrainian Hryvna (active)
	Currency_UAH Currency = 229
	// Ukrainian Karbovanets (inactive)
	Currency_UAK Currency = 230
	// Uganda Shilling (inactive)
	Currency_UGS Currency = 231
	// Uganda Shilling (active)
	Currency_UGX Currency = 232
	// U.S. Dollar (active)
	Currency_USD Currency = 233
	// Uruguay Peso en Unidades Indexadas (active)
	Currency_UYI Currency = 234
	// Uruguayan Old Peso (inactive)
	Currency_UYN Currency = 235
	// Uruguayan Peso (active)
	Currency_UYU Currency = 236
	// Uzbekistan Sum (active)
	Currency_UZS Currency = 237
	// Vatican Lira (inactive)
	Currency_VAL Currency = 238
	// Venezuelan Old Bolivar (inactive)
	Currency_VEB Currency = 239
	// Venezuelan Bolivar (active)
	Currency_VEF Currency = 240
	// Vietnamese Old Dong (inactive)
	Currency_VNC Currency = 241
	// Vietnamese Dong (active)
	Currency_VND Currency = 242
	// Vanuatu Vatu (active)
	Currency_VUV Currency = 243
	// Samoan Tala (active)
	Currency_WST Currency = 244
	// CFA Franc (BEAC) (active)
	Currency_XAF Currency = 245
	// Silver (ounces) (inactive)
	Currency_XAG Currency = 246
	// Gold (ounces) (inactive)
	Currency_XAU Currency = 247
	// East Caribbean Dollar (active)
	Currency_XCD Currency = 248
	// IMF Special Drawing Rts (inactive)
	Currency_XDR Currency = 249
	// European Currency Unit (inactive)
	Currency_XEU Currency = 250
	// Gold Franc (inactive)
	Currency_XFO Currency = 251
	// CFA Franc (BCEAO) (active)
	Currency_XOF Currency = 252
	// Palladium (ounces) (active)
	Currency_XPD Currency = 253
	// CFP Franc (active)
	Currency_XPF Currency = 254
	// Platinum (ounces) (active)
	Currency_XPT Currency = 255
	// South Yemeni Dinar (inactive)
	Currency_YDD Currency = 256
	// Yemeni Rial (active)
	Currency_YER Currency = 257
	// Yugoslav Dinar (inactive)
	Currency_YOU Currency = 258
	// Yugoslav Dinar (inactive)
	Currency_YUD Currency = 259
	// Yugoslav Dinar (inactive)
	Currency_YUG Currency = 260
	// Yugoslav Dinar (inactive)
	Currency_YUM Currency = 261
	// Yugoslav Dinar (inactive)
	Currency_YUN Currency = 262
	// Yugoslav Dinar (inactive)
	Currency_YUR Currency = 263
	// Yugoslav Dinar (inactive)
	Currency_YUS Currency = 264
	// South African Financial Rand (inactive)
	Currency_ZAL Currency = 265
	// South African Rand (active)
	Currency_ZAR Currency = 266
	// Zambian Kwacha (inactive)
	Currency_ZMK Currency = 267
	// Zambian New Kwacha (active)
	Currency_ZMW Currency = 268
	// Zairean Zaire (inactive)
	Currency_ZRN Currency = 269
	// Zairean Zaire (inactive)
	Currency_ZRZ Currency = 270
	// Rhodesian Dollar (inactive)
	Currency_ZWC Currency = 271
	// Zimbabwean Dollar (inactive)
	Currency_ZWD Currency = 272
	// Zimbabwean Dollar (active)
	Currency_ZWL Currency = 273
	// Zimbabwean Dollar (inactive)
	Currency_ZWN Currency = 274
	// Zimbabwean Dollar (inactive)
	Currency_ZWR Currency = 275
)

var Currency_name = map[int32]string{
	0:   "UNKNOWN_CURRENCY",
	1:   "ADP",
	2:   "ADR",
	3:   "AED",
	4:   "AFA",
	5:   "AFN",
	6:   "ALK",
	7:   "ALL",
	8:   "AMD",
	9:   "ANG",
	10:  "AOA",
	11:  "AON",
	12:  "AOR",
	13:  "ARA",
	14:  "ARL",
	15:  "ARM",
	16:  "ARP",
	17:  "ARS",
	18:  "ATS",
	19:  "AUD",
	20:  "AWG",
	21:  "AZM",
	22:  "AZN",
	23:  "BAM",
	24:  "BBD",
	25:  "BDT",
	26:  "BEF",
	27:  "BGJ",
	28:  "BGK",
	29:  "BGL",
	30:  "BGN",
	31:  "BHD",
	32:  "BIF",
	33:  "BMD",
	34:  "BND",
	35:  "BOB",
	36:  "BOP",
	37:  "BOV",
	38:  "BRB",
	39:  "BRC",
	40:  "BRE",
	41:  "BRL",
	42:  "BRN",
	43:  "BRR",
	44:  "BRZ",
	45:  "BSD",
	46:  "BTN",
	47:  "BWP",
	48:  "BYB",
	49:  "BYN",
	50:  "BYR",
	51:  "BZD",
	52:  "CAD",
	53:  "CDF",
	54:  "CHF",
	55:  "CLE",
	56:  "CLF",
	57:  "CLP",
	58:  "CNH",
	59:  "CNY",
	60:  "COP",
	61:  "COU",
	62:  "CRC",
	63:  "CSD",
	64:  "CSK",
	65:  "CUC",
	66:  "CUP",
	67:  "CVE",
	68:  "CYP",
	69:  "CZK",
	70:  "DDM",
	71:  "DEM",
	72:  "DJF",
	73:  "DKK",
	74:  "DOP",
	75:  "DZD",
	76:  "ECS",
	77:  "ECV",
	78:  "EEK",
	79:  "EGP",
	80:  "ERN",
	81:  "ESP",
	82:  "ETB",
	83:  "EUR",
	84:  "FIM",
	85:  "FJD",
	86:  "FKP",
	87:  "FRF",
	88:  "GBP",
	89:  "GBX",
	90:  "GEK",
	91:  "GEL",
	92:  "GHC",
	93:  "GHS",
	94:  "GIP",
	95:  "GMD",
	96:  "GNE",
	97:  "GNF",
	98:  "GQE",
	99:  "GRD",
	100: "GTQ",
	101: "GWP",
	102: "GYD",
	103: "HKD",
	104: "HNL",
	105: "HRK",
	106: "HTG",
	107: "HUF",
	108: "IDR",
	109: "IEP",
	110: "ILP",
	111: "ILR",
	112: "ILS",
	113: "ILX",
	114: "INR",
	115: "IQD",
	116: "IRR",
	117: "ISJ",
	118: "ISK",
	119: "ITL",
	120: "JMD",
	121: "JOD",
	122: "JPY",
	123: "KES",
	124: "KGS",
	125: "KHR",
	126: "KMF",
	127: "KPW",
	128: "KRW",
	129: "KWD",
	130: "KYD",
	131: "KZT",
	132: "LAJ",
	133: "LAK",
	134: "LBP",
	135: "LKR",
	136: "LRD",
	137: "LSL",
	138: "LTL",
	139: "LUF",
	140: "LVL",
	141: "LYD",
	142: "MAD",
	143: "MAF",
	144: "MCL",
	145: "MDL",
	146: "MGA",
	147: "MGF",
	148: "MKD",
	149: "MKN",
	150: "MLF",
	151: "MMK",
	152: "MNT",
	153: "MOP",
	154: "MRO",
	155: "MTL",
	156: "MTP",
	157: "MUR",
	158: "MVQ",
	159: "MVR",
	160: "MWK",
	161: "MXN",
	162: "MXP",
	163: "MXV",
	164: "MYR",
	165: "MZM",
	166: "MZN",
	167: "NAD",
	168: "NFD",
	169: "NGN",
	170: "NIO",
	171: "NLG",
	172: "NOK",
	173: "NPR",
	174: "NZD",
	175: "OMR",
	176: "PAB",
	177: "PEH",
	178: "PEI",
	179: "PEN",
	180: "PGK",
	181: "PHP",
	182: "PKR",
	183: "PLN",
	184: "PLZ",
	185: "PTE",
	186: "PYG",
	187: "QAR",
	188: "ROK",
	189: "ROL",
	190: "RON",
	191: "RSD",
	192: "RUB",
	193: "RUR",
	194: "RWF",
	195: "SAR",
	196: "SBD",
	197: "SCR",
	198: "SDD",
	199: "SDG",
	200: "SEK",
	201: "SGD",
	202: "SHP",
	203: "SIT",
	204: "SKK",
	205: "SLL",
	206: "SML",
	207: "SOS",
	208: "SRD",
	209: "SRG",
	210: "SSP",
	211: "STD",
	212: "SUR",
	213: "SVC",
	214: "SYP",
	215: "SZL",
	216: "THB",
	217: "TJR",
	218: "TJS",
	219: "TMM",
	220: "TMT",
	221: "TND",
	222: "TOP",
	223: "TPE",
	224: "TRL",
	225: "TRY",
	226: "TTD",
	227: "TWD",
	228: "TZS",
	229: "UAH",
	230: "UAK",
	231: "UGS",
	232: "UGX",
	233: "USD",
	234: "UYI",
	235: "UYN",
	236: "UYU",
	237: "UZS",
	238: "VAL",
	239: "VEB",
	240: "VEF",
	241: "VNC",
	242: "VND",
	243: "VUV",
	244: "WST",
	245: "XAF",
	246: "XAG",
	247: "XAU",
	248: "XCD",
	249: "XDR",
	250: "XEU",
	251: "XFO",
	252: "XOF",
	253: "XPD",
	254: "XPF",
	255: "XPT",
	256: "YDD",
	257: "YER",
	258: "YOU",
	259: "YUD",
	260: "YUG",
	261: "YUM",
	262: "YUN",
	263: "YUR",
	264: "YUS",
	265: "ZAL",
	266: "ZAR",
	267: "ZMK",
	268: "ZMW",
	269: "ZRN",
	270: "ZRZ",
	271: "ZWC",
	272: "ZWD",
	273: "ZWL",
	274: "ZWN",
	275: "ZWR",
}
var Currency_value = map[string]int32{
	"UNKNOWN_CURRENCY": 0,
	"ADP":              1,
	"ADR":              2,
	"AED":              3,
	"AFA":              4,
	"AFN":              5,
	"ALK":              6,
	"ALL":              7,
	"AMD":              8,
	"ANG":              9,
	"AOA":              10,
	"AON":              11,
	"AOR":              12,
	"ARA":              13,
	"ARL":              14,
	"ARM":              15,
	"ARP":              16,
	"ARS":              17,
	"ATS":              18,
	"AUD":              19,
	"AWG":              20,
	"AZM":              21,
	"AZN":              22,
	"BAM":              23,
	"BBD":              24,
	"BDT":              25,
	"BEF":              26,
	"BGJ":              27,
	"BGK":              28,
	"BGL":              29,
	"BGN":              30,
	"BHD":              31,
	"BIF":              32,
	"BMD":              33,
	"BND":              34,
	"BOB":              35,
	"BOP":              36,
	"BOV":              37,
	"BRB":              38,
	"BRC":              39,
	"BRE":              40,
	"BRL":              41,
	"BRN":              42,
	"BRR":              43,
	"BRZ":              44,
	"BSD":              45,
	"BTN":              46,
	"BWP":              47,
	"BYB":              48,
	"BYN":              49,
	"BYR":              50,
	"BZD":              51,
	"CAD":              52,
	"CDF":              53,
	"CHF":              54,
	"CLE":              55,
	"CLF":              56,
	"CLP":              57,
	"CNH":              58,
	"CNY":              59,
	"COP":              60,
	"COU":              61,
	"CRC":              62,
	"CSD":              63,
	"CSK":              64,
	"CUC":              65,
	"CUP":              66,
	"CVE":              67,
	"CYP":              68,
	"CZK":              69,
	"DDM":              70,
	"DEM":              71,
	"DJF":              72,
	"DKK":              73,
	"DOP":              74,
	"DZD":              75,
	"ECS":              76,
	"ECV":              77,
	"EEK":              78,
	"EGP":              79,
	"ERN":              80,
	"ESP":              81,
	"ETB":              82,
	"EUR":              83,
	"FIM":              84,
	"FJD":              85,
	"FKP":              86,
	"FRF":              87,
	"GBP":              88,
	"GBX":              89,
	"GEK":              90,
	"GEL":              91,
	"GHC":              92,
	"GHS":              93,
	"GIP":              94,
	"GMD":              95,
	"GNE":              96,
	"GNF":              97,
	"GQE":              98,
	"GRD":              99,
	"GTQ":              100,
	"GWP":              101,
	"GYD":              102,
	"HKD":              103,
	"HNL":              104,
	"HRK":              105,
	"HTG":              106,
	"HUF":              107,
	"IDR":              108,
	"IEP":              109,
	"ILP":              110,
	"ILR":              111,
	"ILS":              112,
	"ILX":              113,
	"INR":              114,
	"IQD":              115,
	"IRR":              116,
	"ISJ":              117,
	"ISK":              118,
	"ITL":              119,
	"JMD":              120,
	"JOD":              121,
	"JPY":              122,
	"KES":              123,
	"KGS":              124,
	"KHR":              125,
	"KMF":              126,
	"KPW":              127,
	"KRW":              128,
	"KWD":              129,
	"KYD":              130,
	"KZT":              131,
	"LAJ":              132,
	"LAK":              133,
	"LBP":              134,
	"LKR":              135,
	"LRD":              136,
	"LSL":              137,
	"LTL":              138,
	"LUF":              139,
	"LVL":              140,
	"LYD":              141,
	"MAD":              142,
	"MAF":              143,
	"MCL":              144,
	"MDL":              145,
	"MGA":              146,
	"MGF":              147,
	"MKD":              148,
	"MKN":              149,
	"MLF":              150,
	"MMK":              151,
	"MNT":              152,
	"MOP":              153,
	"MRO":              154,
	"MTL":              155,
	"MTP":              156,
	"MUR":              157,
	"MVQ":              158,
	"MVR":              159,
	"MWK":              160,
	"MXN":              161,
	"MXP":              162,
	"MXV":              163,
	"MYR":              164,
	"MZM":              165,
	"MZN":              166,
	"NAD":              167,
	"NFD":              168,
	"NGN":              169,
	"NIO":              170,
	"NLG":              171,
	"NOK":              172,
	"NPR":              173,
	"NZD":              174,
	"OMR":              175,
	"PAB":              176,
	"PEH":              177,
	"PEI":              178,
	"PEN":              179,
	"PGK":              180,
	"PHP":              181,
	"PKR":              182,
	"PLN":              183,
	"PLZ":              184,
	"PTE":              185,
	"PYG":              186,
	"QAR":              187,
	"ROK":              188,
	"ROL":              189,
	"RON":              190,
	"RSD":              191,
	"RUB":              192,
	"RUR":              193,
	"RWF":              194,
	"SAR":              195,
	"SBD":              196,
	"SCR":              197,
	"SDD":              198,
	"SDG":              199,
	"SEK":              200,
	"SGD":              201,
	"SHP":              202,
	"SIT":              203,
	"SKK":              204,
	"SLL":              205,
	"SML":              206,
	"SOS":              207,
	"SRD":              208,
	"SRG":              209,
	"SSP":              210,
	"STD":              211,
	"SUR":              212,
	"SVC":              213,
	"SYP":              214,
	"SZL":              215,
	"THB":              216,
	"TJR":              217,
	"TJS":              218,
	"TMM":              219,
	"TMT":              220,
	"TND":              221,
	"TOP":              222,
	"TPE":              223,
	"TRL":              224,
	"TRY":              225,
	"TTD":              226,
	"TWD":              227,
	"TZS":              228,
	"UAH":              229,
	"UAK":              230,
	"UGS":              231,
	"UGX":              232,
	"USD":              233,
	"UYI":              234,
	"UYN":              235,
	"UYU":              236,
	"UZS":              237,
	"VAL":              238,
	"VEB":              239,
	"VEF":              240,
	"VNC":              241,
	"VND":              242,
	"VUV":              243,
	"WST":              244,
	"XAF":              245,
	"XAG":              246,
	"XAU":              247,
	"XCD":              248,
	"XDR":              249,
	"XEU":              250,
	"XFO":              251,
	"XOF":              252,
	"XPD":              253,
	"XPF":              254,
	"XPT":              255,
	"YDD":              256,
	"YER":              257,
	"YOU":              258,
	"YUD":              259,
	"YUG":              260,
	"YUM":              261,
	"YUN":              262,
	"YUR":              263,
	"YUS":              264,
	"ZAL":              265,
	"ZAR":              266,
	"ZMK":              267,
	"ZMW":              268,
	"ZRN":              269,
	"ZRZ":              270,
	"ZWC":              271,
	"ZWD":              272,
	"ZWL":              273,
	"ZWN":              274,
	"ZWR":              275,
}

func (x Currency) String() string {
	return proto.EnumName(Currency_name, int32(x))
}
func (Currency) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_currency_029b0943ac8782e0, []int{0}
}

func init() {
	proto.RegisterEnum("iso.protobuf.Currency", Currency_name, Currency_value)
}

func init() { proto.RegisterFile("currency.proto", fileDescriptor_currency_029b0943ac8782e0) }

var fileDescriptor_currency_029b0943ac8782e0 = []byte{
	// 1221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0xd6, 0x65, 0x9f, 0x24, 0x57,
	0x19, 0x40, 0x71, 0xb2, 0x0b, 0xc9, 0x32, 0x84, 0x70, 0x18, 0x16, 0x08, 0xc1, 0x1d, 0x02, 0xec,
	0x02, 0xc1, 0xbd, 0xba, 0x6e, 0x55, 0x75, 0x77, 0x55, 0xdd, 0xba, 0x73, 0xab, 0xaa, 0x7b, 0xba,
	0x90, 0xc0, 0x0c, 0x9b, 0x64, 0x08, 0x3b, 0x13, 0x66, 0x77, 0x80, 0x41, 0x23, 0xeb, 0x2e, 0xb8,
	0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xc3, 0xe2, 0xee, 0xee, 0xf0, 0xeb, 0xe7, 0xb9, 0xef, 0xfe, 0x2f,
	0xce, 0x07, 0x38, 0x33, 0x67, 0x2d, 0xae, 0xad, 0xae, 0x6e, 0x5b, 0x5e, 0x5c, 0xdf, 0x72, 0xc9,
	0xea, 0xca, 0xce, 0x95, 0xd9, 0x33, 0x97, 0x76, 0xac, 0x28, 0x17, 0xd6, 0x2e, 0x38, 0xf7, 0xd4,
	0xe6, 0x99, 0x4d, 0x71, 0x08, 0x66, 0x37, 0xcf, 0xd0, 0xda, 0xdc, 0x56, 0x63, 0x7b, 0x7e, 0xdc,
	0x7a, 0x9f, 0xd8, 0x78, 0xc2, 0x95, 0x66, 0xcf, 0x98, 0xd9, 0x18, 0x19, 0xc7, 0x69, 0x0a, 0xcf,
	0x06, 0x41, 0x62, 0xd8, 0x28, 0x48, 0x23, 0xae, 0xac, 0xb0, 0x5c, 0x45, 0x50, 0xe4, 0x9c, 0xae,
	0x28, 0x38, 0x43, 0x50, 0x1a, 0x36, 0x09, 0x6c, 0xc6, 0x55, 0x05, 0x55, 0xc4, 0x8c, 0xc2, 0x72,
	0x35, 0x85, 0xe7, 0x4c, 0x81, 0x8f, 0xb8, 0xba, 0xa2, 0xe0, 0x2c, 0x45, 0xc9, 0x35, 0x14, 0x0e,
	0x14, 0x35, 0xd7, 0x14, 0x34, 0x35, 0xb3, 0x82, 0xd6, 0x70, 0x2d, 0xc1, 0x38, 0x63, 0xb3, 0xa0,
	0x2b, 0xb9, 0xb6, 0xc2, 0x72, 0x9d, 0x29, 0x7a, 0x51, 0xc9, 0x75, 0x05, 0x3d, 0xc3, 0xd9, 0x02,
	0xd3, 0x70, 0x3d, 0x41, 0x92, 0x72, 0x8e, 0x20, 0x1b, 0x72, 0x7d, 0x45, 0xce, 0x0d, 0x14, 0x05,
	0x37, 0x54, 0x58, 0x6e, 0x24, 0xe8, 0x1b, 0x6e, 0x2c, 0x18, 0xa4, 0xdc, 0x44, 0x50, 0x1a, 0x6e,
	0x2a, 0xb0, 0x86, 0x9b, 0x09, 0xaa, 0x1e, 0x37, 0x57, 0x38, 0x6e, 0xa1, 0x18, 0x71, 0x4b, 0x81,
	0xef, 0x71, 0x2b, 0x45, 0xcc, 0xad, 0x15, 0x09, 0xb7, 0x51, 0x14, 0xdc, 0x56, 0x61, 0x39, 0x57,
	0xe1, 0xb9, 0x9d, 0xa2, 0xe3, 0xf6, 0x82, 0xda, 0x70, 0x07, 0x41, 0x63, 0xd9, 0x22, 0x18, 0x3b,
	0xb6, 0x0a, 0x26, 0x3d, 0xee, 0xa8, 0xb0, 0xdc, 0x49, 0xe1, 0xb9, 0xb3, 0xa0, 0x33, 0x9c, 0x37,
	0x45, 0x1c, 0x19, 0xee, 0x22, 0x30, 0x29, 0x77, 0x15, 0xf4, 0x53, 0xee, 0x26, 0x28, 0x12, 0xee,
	0xae, 0x48, 0xb9, 0x87, 0xc2, 0x71, 0x4f, 0x81, 0xed, 0x73, 0x2f, 0xc5, 0x84, 0x7b, 0x0b, 0x2a,
	0xc7, 0x7d, 0x14, 0x2d, 0xf7, 0x15, 0xf8, 0x98, 0xfb, 0x09, 0x6a, 0xc3, 0xfd, 0x15, 0x39, 0x0f,
	0x10, 0xb4, 0x31, 0x91, 0xc2, 0xd1, 0x13, 0x8c, 0x12, 0x62, 0xc1, 0xc4, 0x61, 0x04, 0x5d, 0x4e,
	0x32, 0x85, 0x31, 0x25, 0xa9, 0x20, 0x29, 0xc9, 0x04, 0xc3, 0x94, 0xbe, 0x20, 0xcf, 0x19, 0x08,
	0x2a, 0xc7, 0x50, 0xd0, 0x19, 0xf2, 0x29, 0x92, 0xb8, 0xa6, 0x50, 0x8c, 0x28, 0x05, 0x49, 0x8e,
	0x15, 0x64, 0x8e, 0x4a, 0xe0, 0x2d, 0x4e, 0x50, 0x3b, 0xe6, 0x04, 0x4d, 0x0f, 0x2f, 0x68, 0x3d,
	0xf5, 0x14, 0xe9, 0xa0, 0xa4, 0x11, 0x0c, 0x0d, 0xad, 0x20, 0x77, 0x8c, 0x04, 0x3e, 0x65, 0x3c,
	0x45, 0xd6, 0x73, 0xcc, 0x2b, 0xe6, 0x99, 0x08, 0x92, 0x9c, 0x4e, 0x51, 0xf0, 0x40, 0x41, 0x3f,
	0xe6, 0x41, 0x8a, 0x9a, 0x07, 0x0b, 0x06, 0x8e, 0x87, 0x08, 0x4a, 0xc3, 0xf9, 0x02, 0x9b, 0xf0,
	0x50, 0x45, 0xca, 0xc3, 0x04, 0x73, 0x09, 0x0b, 0x02, 0x6f, 0x58, 0x14, 0x34, 0x73, 0x3c, 0x5c,
	0x30, 0x76, 0x6c, 0x13, 0x4c, 0x0c, 0x17, 0x4c, 0xd1, 0xcf, 0x0d, 0x17, 0x0a, 0x6c, 0xc1, 0x45,
	0x02, 0x9f, 0xb3, 0x24, 0x68, 0x32, 0x1e, 0x21, 0x68, 0x53, 0x2e, 0x9e, 0x62, 0x60, 0x3c, 0x8f,
	0x14, 0x24, 0x8e, 0xed, 0x82, 0xc2, 0xb1, 0xac, 0xf0, 0xac, 0x28, 0x6a, 0x2e, 0x51, 0xcc, 0xf3,
	0x28, 0x81, 0xf5, 0xac, 0x0a, 0xe6, 0x0c, 0x3b, 0x04, 0xde, 0xb3, 0x53, 0x50, 0x0f, 0x59, 0x53,
	0xe4, 0x3c, 0x5a, 0xd0, 0x14, 0x3c, 0x66, 0x8a, 0x61, 0x69, 0x78, 0xac, 0xa0, 0x32, 0xac, 0x0b,
	0xdc, 0x84, 0xc7, 0x4d, 0x91, 0x27, 0x35, 0x8f, 0x17, 0x64, 0x35, 0x4f, 0x10, 0xf4, 0x3d, 0x4f,
	0x14, 0x94, 0x29, 0x4f, 0x12, 0xb8, 0x31, 0x4f, 0x9e, 0xdd, 0x34, 0xb3, 0x31, 0xf7, 0x63, 0x2e,
	0x3d, 0x4d, 0x34, 0x36, 0x5c, 0xa6, 0x9a, 0x18, 0x2e, 0x57, 0x75, 0x0d, 0x57, 0x88, 0x8a, 0x68,
	0xc8, 0xae, 0xa0, 0x9c, 0xdd, 0xaa, 0x9e, 0x63, 0x8f, 0x2a, 0xf7, 0xec, 0x55, 0x79, 0xc3, 0x3e,
	0x55, 0x5d, 0xb0, 0x5f, 0xd5, 0x14, 0x1c, 0x50, 0xb5, 0x29, 0x07, 0x55, 0xa3, 0x82, 0x43, 0xaa,
	0x89, 0xe1, 0xb0, 0xa8, 0x8c, 0x0c, 0x47, 0x82, 0x52, 0x8e, 0xaa, 0xe2, 0x82, 0x63, 0x2a, 0x53,
	0x70, 0x5c, 0x95, 0x45, 0x9c, 0x08, 0x4a, 0x39, 0xa9, 0xca, 0x0d, 0x4f, 0x09, 0xb2, 0x3c, 0x55,
	0x55, 0xa4, 0x3c, 0x4d, 0x55, 0xe6, 0x3c, 0x5d, 0x65, 0x1b, 0x9e, 0xa1, 0xaa, 0x1c, 0xcf, 0x54,
	0xf9, 0x8a, 0x67, 0xa9, 0x9a, 0x82, 0x67, 0x07, 0x39, 0x9e, 0xa3, 0x6a, 0x3d, 0xcf, 0x55, 0x8d,
	0xe6, 0x78, 0x5e, 0x90, 0xe7, 0xf9, 0xaa, 0x71, 0xce, 0x0b, 0x54, 0xf3, 0x96, 0x17, 0x06, 0x39,
	0x5e, 0x14, 0x34, 0xe2, 0xc5, 0xaa, 0x89, 0xe7, 0x25, 0xaa, 0xae, 0xe4, 0xa5, 0x41, 0x96, 0x97,
	0x89, 0x6c, 0x64, 0x78, 0xb9, 0x2a, 0x35, 0xbc, 0x42, 0x95, 0x59, 0x5e, 0xa9, 0x1a, 0x54, 0xbc,
	0x4a, 0x55, 0x64, 0xbc, 0x5a, 0x55, 0xe5, 0xbc, 0x46, 0xe5, 0x3c, 0xaf, 0x55, 0x75, 0x86, 0xd7,
	0x89, 0xaa, 0xd2, 0xf3, 0x7a, 0x91, 0x8b, 0x7a, 0xbc, 0x41, 0x95, 0xf4, 0x79, 0x63, 0xd0, 0x80,
	0x37, 0x05, 0x59, 0xde, 0xac, 0xca, 0x72, 0xde, 0xa2, 0xea, 0x3b, 0xde, 0xaa, 0xca, 0x3d, 0x6f,
	0x53, 0x15, 0x96, 0xb7, 0x07, 0x75, 0xbc, 0x43, 0xd5, 0x24, 0xbc, 0x53, 0x35, 0xc9, 0x78, 0x97,
	0x68, 0x2e, 0xf2, 0xbc, 0x5b, 0xe4, 0xab, 0x9c, 0xf7, 0x04, 0x15, 0xbc, 0x37, 0xc8, 0xf2, 0x3e,
	0x55, 0x6d, 0x78, 0xbf, 0xaa, 0xed, 0xf1, 0x81, 0x20, 0xcf, 0x07, 0x55, 0xe3, 0x94, 0x0f, 0x89,
	0xea, 0xc8, 0xf3, 0x61, 0x55, 0xcf, 0xf0, 0x11, 0x55, 0xec, 0xf9, 0xa8, 0xca, 0x18, 0x3e, 0x16,
	0x94, 0xf1, 0x71, 0x55, 0x92, 0xf3, 0x09, 0x55, 0x66, 0xf8, 0xa4, 0xaa, 0xef, 0xf8, 0x94, 0x6a,
	0xd0, 0xf0, 0x69, 0x55, 0x9e, 0xf3, 0x19, 0x55, 0x51, 0xf0, 0x59, 0x55, 0x59, 0xf0, 0x39, 0x55,
	0x55, 0xf3, 0x79, 0x95, 0x37, 0x7c, 0x21, 0x28, 0xe3, 0x8b, 0xaa, 0xda, 0xf1, 0x25, 0x55, 0x63,
	0xf8, 0xb2, 0xaa, 0xf5, 0x7c, 0x45, 0x35, 0x8a, 0xf9, 0xaa, 0x6a, 0xe2, 0xf8, 0x9a, 0xaa, 0x2b,
	0xf8, 0xba, 0xa8, 0xe9, 0xf7, 0xf8, 0x86, 0x6a, 0xe8, 0xf9, 0x66, 0x50, 0xcd, 0xb7, 0x54, 0x65,
	0xc9, 0xb7, 0x83, 0x1a, 0xbe, 0xa3, 0xb2, 0x86, 0xef, 0xaa, 0x2a, 0xc7, 0xf7, 0x54, 0x2e, 0xe1,
	0x94, 0xca, 0x17, 0x7c, 0x3f, 0x68, 0xc2, 0x0f, 0x54, 0x8d, 0xe1, 0x87, 0xaa, 0xb1, 0xe1, 0x47,
	0xaa, 0xae, 0xe6, 0xc7, 0xa2, 0x36, 0xea, 0xf3, 0x93, 0xa0, 0x9c, 0x9f, 0xaa, 0xb2, 0x9a, 0x9f,
	0x05, 0xcd, 0xf3, 0x73, 0x55, 0x6d, 0xf8, 0x85, 0x6a, 0x32, 0xe0, 0x97, 0x41, 0x96, 0x5f, 0x05,
	0xb5, 0xfc, 0x5a, 0xd5, 0xd5, 0xfc, 0x46, 0x34, 0x8a, 0x0a, 0x7e, 0xab, 0x4a, 0x7a, 0xfc, 0x2e,
	0x28, 0xe5, 0xf7, 0x2a, 0x1b, 0xf3, 0x87, 0x20, 0xc3, 0x1f, 0x55, 0xed, 0x88, 0x3f, 0x89, 0xc6,
	0x75, 0xc3, 0x9f, 0x45, 0xf3, 0x51, 0xca, 0x5f, 0x82, 0x32, 0xfe, 0x1a, 0xd4, 0xf2, 0x37, 0x55,
	0x6c, 0xf8, 0xbb, 0xca, 0x78, 0xfe, 0xa1, 0x4a, 0x5a, 0xfe, 0xa9, 0x4a, 0x2b, 0xfe, 0xa5, 0xaa,
	0x52, 0xfe, 0xad, 0x72, 0x86, 0xff, 0x04, 0xa5, 0xfc, 0x37, 0xa8, 0xe1, 0x7f, 0xa2, 0x89, 0x31,
	0x5c, 0xba, 0x41, 0x94, 0x78, 0x2e, 0x53, 0x55, 0x2d, 0x97, 0xab, 0x5a, 0xc3, 0x15, 0x41, 0x19,
	0xbb, 0x82, 0x4a, 0x76, 0x07, 0x59, 0xf6, 0x04, 0x79, 0xf6, 0x06, 0xd5, 0xec, 0x13, 0x75, 0x51,
	0xc1, 0xfe, 0x20, 0xcf, 0x01, 0x55, 0x99, 0x73, 0x30, 0x68, 0xcc, 0x21, 0x95, 0xb7, 0x1c, 0x0e,
	0xea, 0x38, 0xa2, 0x1a, 0xc7, 0x1c, 0x0d, 0x32, 0x1c, 0x0b, 0x2a, 0x38, 0x1e, 0x64, 0x39, 0x11,
	0xe4, 0x39, 0xb9, 0xa1, 0x77, 0x4e, 0x77, 0xf6, 0x85, 0x4b, 0x3b, 0x2f, 0x5a, 0x5b, 0xd8, 0xb2,
	0xb8, 0xb2, 0x7d, 0xeb, 0xf2, 0xfa, 0xe2, 0xf6, 0x95, 0xe5, 0x8b, 0xb7, 0xad, 0x6f, 0x5d, 0xda,
	0xb1, 0xb2, 0x70, 0xba, 0x7c, 0xe8, 0x79, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x34, 0xd0, 0xb5,
	0x09, 0xa4, 0x0a, 0x00, 0x00,
}
