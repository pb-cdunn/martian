//line martian/syntax/grammar.y:2

//
// Copyright (c) 2014 10X Genomics, Inc. All rights reserved.
//
// MRO grammar.
//

package syntax

import __yyfmt__ "fmt"

//line martian/syntax/grammar.y:8
import (
	"strings"
)

//line martian/syntax/grammar.y:16
type mmSymType struct {
	yys       int
	global    *Ast
	srcfile   *SourceFile
	arr       int16
	loc       int
	val       []byte
	modifiers *Modifiers
	dec       Dec
	decs      []Dec
	inparam   *InParam
	outparam  *OutParam
	retains   []*RetainParam
	stretains *RetainParams
	i_params  *InParams
	o_params  *OutParams
	res       *Resources
	par_tuple paramsTuple
	src       *SrcParam
	exp       Exp
	exps      []Exp
	rexp      *RefExp
	vexp      *ValExp
	kvpairs   map[string]Exp
	call      *CallStm
	calls     []*CallStm
	binding   *BindStm
	bindings  *BindStms
	retstm    *ReturnStm
	plretains *PipelineRetains
	reflist   []*RefExp
	includes  []*Include
	intern    *stringIntern
}

const SKIP = 57346
const COMMENT = 57347
const INVALID = 57348
const SEMICOLON = 57349
const COLON = 57350
const COMMA = 57351
const EQUALS = 57352
const LBRACKET = 57353
const RBRACKET = 57354
const LPAREN = 57355
const RPAREN = 57356
const LBRACE = 57357
const RBRACE = 57358
const SWEEP = 57359
const RETURN = 57360
const SELF = 57361
const FILETYPE = 57362
const STAGE = 57363
const PIPELINE = 57364
const CALL = 57365
const SPLIT = 57366
const USING = 57367
const RETAIN = 57368
const LOCAL = 57369
const PREFLIGHT = 57370
const VOLATILE = 57371
const DISABLED = 57372
const STRICT = 57373
const IN = 57374
const OUT = 57375
const SRC = 57376
const AS = 57377
const THREADS = 57378
const MEM_GB = 57379
const SPECIAL = 57380
const ID = 57381
const LITSTRING = 57382
const NUM_FLOAT = 57383
const NUM_INT = 57384
const DOT = 57385
const PY = 57386
const EXEC = 57387
const COMPILED = 57388
const MAP = 57389
const INT = 57390
const STRING = 57391
const FLOAT = 57392
const PATH = 57393
const BOOL = 57394
const TRUE = 57395
const FALSE = 57396
const NULL = 57397
const DEFAULT = 57398
const INCLUDE_DIRECTIVE = 57399

var mmToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"SKIP",
	"COMMENT",
	"INVALID",
	"SEMICOLON",
	"COLON",
	"COMMA",
	"EQUALS",
	"LBRACKET",
	"RBRACKET",
	"LPAREN",
	"RPAREN",
	"LBRACE",
	"RBRACE",
	"SWEEP",
	"RETURN",
	"SELF",
	"FILETYPE",
	"STAGE",
	"PIPELINE",
	"CALL",
	"SPLIT",
	"USING",
	"RETAIN",
	"LOCAL",
	"PREFLIGHT",
	"VOLATILE",
	"DISABLED",
	"STRICT",
	"IN",
	"OUT",
	"SRC",
	"AS",
	"THREADS",
	"MEM_GB",
	"SPECIAL",
	"ID",
	"LITSTRING",
	"NUM_FLOAT",
	"NUM_INT",
	"DOT",
	"PY",
	"EXEC",
	"COMPILED",
	"MAP",
	"INT",
	"STRING",
	"FLOAT",
	"PATH",
	"BOOL",
	"TRUE",
	"FALSE",
	"NULL",
	"DEFAULT",
	"INCLUDE_DIRECTIVE",
}
var mmStatenames = [...]string{}

const mmEofCode = 1
const mmErrCode = 2
const mmInitialStackSize = 16

//line martian/syntax/grammar.y:725

//line yacctab:1
var mmExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 44,
	13, 111,
	35, 111,
	-2, 70,
	-1, 45,
	13, 113,
	35, 113,
	-2, 71,
	-1, 46,
	13, 120,
	35, 120,
	-2, 72,
}

const mmPrivate = 57344

const mmLast = 607

var mmAct = [...]int{

	96, 117, 140, 65, 171, 63, 55, 150, 138, 22,
	106, 4, 38, 39, 14, 16, 81, 123, 91, 92,
	212, 43, 102, 103, 104, 40, 27, 47, 113, 112,
	33, 36, 31, 28, 30, 37, 25, 34, 8, 11,
	12, 7, 35, 29, 32, 23, 48, 223, 184, 54,
	222, 26, 24, 224, 64, 200, 141, 56, 18, 191,
	68, 173, 170, 48, 75, 155, 128, 41, 22, 19,
	204, 67, 183, 52, 95, 15, 225, 201, 202, 203,
	143, 22, 99, 8, 11, 12, 7, 90, 93, 94,
	172, 218, 152, 172, 177, 53, 152, 105, 80, 79,
	75, 114, 166, 147, 145, 127, 149, 131, 130, 7,
	80, 162, 206, 134, 135, 146, 129, 27, 163, 133,
	5, 33, 36, 31, 28, 30, 37, 25, 34, 89,
	151, 80, 193, 35, 29, 32, 23, 152, 107, 154,
	158, 100, 26, 24, 157, 7, 159, 194, 80, 8,
	11, 12, 7, 179, 169, 57, 186, 153, 180, 174,
	6, 178, 168, 181, 17, 167, 137, 185, 59, 60,
	61, 62, 76, 189, 17, 188, 50, 160, 49, 192,
	161, 217, 181, 42, 195, 216, 215, 214, 98, 72,
	71, 70, 205, 69, 230, 229, 75, 1, 228, 118,
	213, 211, 196, 119, 227, 226, 221, 97, 27, 210,
	220, 207, 33, 36, 31, 28, 30, 37, 25, 34,
	197, 190, 175, 148, 35, 29, 32, 23, 122, 120,
	121, 118, 182, 26, 24, 119, 136, 111, 110, 97,
	27, 91, 92, 124, 33, 36, 31, 28, 30, 37,
	25, 34, 109, 108, 198, 164, 35, 29, 32, 23,
	122, 120, 121, 118, 139, 26, 24, 119, 187, 144,
	156, 97, 27, 91, 92, 124, 33, 36, 31, 28,
	30, 37, 25, 34, 51, 58, 74, 88, 35, 29,
	32, 23, 122, 120, 121, 118, 21, 26, 24, 119,
	132, 115, 142, 97, 27, 91, 92, 124, 33, 36,
	31, 28, 30, 37, 25, 34, 3, 116, 77, 13,
	35, 29, 32, 23, 122, 120, 121, 118, 126, 26,
	24, 119, 176, 208, 165, 97, 27, 91, 92, 124,
	33, 36, 31, 28, 30, 37, 25, 34, 199, 78,
	66, 10, 35, 29, 32, 23, 122, 120, 121, 9,
	20, 26, 24, 101, 2, 0, 0, 0, 27, 91,
	92, 124, 33, 36, 31, 28, 30, 37, 25, 34,
	0, 0, 0, 0, 35, 29, 32, 23, 0, 0,
	0, 0, 0, 26, 24, 87, 82, 83, 85, 84,
	86, 219, 0, 0, 0, 0, 97, 27, 0, 0,
	0, 33, 36, 31, 28, 30, 37, 25, 34, 0,
	0, 0, 0, 35, 29, 32, 23, 0, 209, 0,
	0, 0, 26, 24, 27, 0, 0, 0, 33, 36,
	31, 28, 30, 37, 25, 34, 0, 0, 130, 0,
	35, 29, 32, 23, 0, 0, 0, 27, 0, 26,
	24, 33, 36, 31, 28, 30, 37, 25, 34, 0,
	0, 0, 0, 35, 29, 32, 23, 0, 125, 0,
	0, 0, 26, 24, 27, 0, 0, 0, 33, 36,
	31, 28, 30, 37, 25, 34, 0, 0, 0, 0,
	35, 29, 32, 23, 0, 0, 97, 27, 0, 26,
	24, 33, 36, 31, 28, 30, 37, 25, 34, 0,
	0, 0, 0, 35, 29, 32, 23, 0, 73, 0,
	0, 0, 26, 24, 27, 0, 0, 0, 33, 36,
	31, 28, 30, 37, 25, 34, 0, 0, 0, 0,
	35, 29, 32, 23, 0, 0, 0, 27, 0, 26,
	24, 33, 36, 31, 28, 30, 37, 25, 34, 0,
	0, 0, 0, 35, 29, 32, 23, 0, 0, 0,
	27, 0, 26, 24, 33, 36, 31, 44, 45, 46,
	25, 34, 0, 0, 0, 0, 35, 29, 32, 23,
	0, 0, 0, 0, 0, 26, 24,
}
var mmPact = [...]int{

	63, -1000, 18, 129, 33, 29, -1000, -1000, 537, -1000,
	-1000, 537, 537, 129, 33, 27, 33, -1000, 170, -1000,
	560, 20, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 165, 163,
	33, -1000, -1000, 60, -1000, -1000, -1000, -1000, 537, -1000,
	-1000, 141, -1000, 537, -1000, 39, 39, -1000, -1000, 183,
	181, 180, 179, 514, 159, 65, -1000, 348, 115, -35,
	-35, -35, 487, -1000, -1000, 178, -1000, 127, -1000, -22,
	348, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 3, 123,
	244, -1000, -1000, 243, 229, 228, -14, -15, 284, 464,
	81, 26, -1000, -1000, -1000, -1000, 437, 86, -1000, -1000,
	-1000, -1000, 537, 537, 227, 153, -1000, -1000, 252, 40,
	-1000, -1000, -1000, -1000, -1000, -1000, 79, 90, 214, 97,
	145, 56, 122, 33, -1000, -1000, -1000, 316, 168, -1000,
	-1000, -1000, 102, 247, 76, 152, 149, -1000, -1000, -1000,
	53, 52, -1000, -1000, 213, -1000, 68, 33, 148, 144,
	220, -1000, 32, -1000, 316, -1000, 143, -1000, -1000, 39,
	-1000, 212, -1000, -1000, 50, -1000, 116, 134, -1000, 188,
	211, -1000, -1000, 246, -1000, -1000, -1000, 41, 39, 98,
	-1000, -1000, 202, -1000, -1000, 414, 200, -1000, 316, 6,
	-1000, 177, 176, 175, 171, 77, -1000, -1000, 387, -1000,
	-1000, -1000, -1000, 197, 8, 5, 13, 45, -1000, -1000,
	196, -1000, 195, 189, 186, 185, -1000, -1000, -1000, -1000,
	-1000,
}
var mmPgo = [...]int{

	0, 364, 0, 287, 16, 7, 363, 4, 360, 10,
	160, 359, 351, 316, 350, 349, 348, 334, 333, 332,
	6, 3, 328, 318, 2, 1, 317, 17, 8, 302,
	11, 300, 286, 285, 5, 284, 270, 269, 268, 197,
}
var mmR1 = [...]int{

	0, 39, 39, 39, 39, 39, 39, 1, 1, 13,
	13, 10, 10, 10, 12, 11, 37, 37, 38, 38,
	38, 38, 38, 17, 17, 16, 16, 3, 3, 9,
	9, 20, 20, 14, 14, 21, 21, 15, 15, 15,
	15, 15, 15, 23, 5, 7, 4, 4, 4, 4,
	4, 4, 4, 6, 6, 6, 22, 22, 22, 36,
	19, 19, 18, 18, 31, 31, 30, 30, 30, 8,
	8, 8, 8, 35, 35, 33, 33, 33, 33, 34,
	34, 32, 32, 32, 28, 28, 29, 29, 24, 24,
	26, 26, 26, 26, 26, 26, 26, 26, 26, 26,
	26, 27, 27, 25, 25, 25, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2,
}
var mmR2 = [...]int{

	0, 2, 3, 2, 1, 2, 1, 3, 2, 2,
	1, 3, 1, 1, 11, 10, 0, 4, 0, 5,
	5, 5, 5, 0, 4, 0, 3, 3, 1, 0,
	3, 0, 2, 6, 5, 0, 2, 4, 5, 6,
	5, 6, 7, 4, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 0, 6, 5, 4,
	0, 4, 0, 3, 2, 1, 6, 8, 5, 0,
	2, 2, 2, 0, 2, 4, 4, 4, 4, 0,
	2, 4, 8, 7, 3, 1, 5, 3, 1, 1,
	3, 4, 2, 2, 3, 4, 1, 1, 1, 1,
	1, 1, 1, 3, 1, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1,
}
var mmChk = [...]int{

	-1000, -39, -1, -13, -30, 57, -10, 23, 20, -11,
	-12, 21, 22, -13, -30, 57, -30, -10, 25, 40,
	-8, -3, -2, 39, 46, 30, 45, 20, 27, 37,
	28, 26, 38, 24, 31, 36, 25, 29, -2, -2,
	-30, 40, 13, -2, 27, 28, 29, 7, 43, 13,
	13, -35, 13, 35, -2, -20, -20, 14, -33, 27,
	28, 29, 30, -34, -2, -21, -14, 32, -21, 10,
	10, 10, 10, 14, -32, -2, 13, -23, -15, 34,
	33, -4, 48, 49, 51, 50, 52, 47, -3, 14,
	-27, 53, 54, -27, -27, -25, -2, 19, 10, -34,
	14, -6, 44, 45, 46, -4, -9, 15, 9, 9,
	9, 9, 43, 43, -24, 17, -26, -25, 11, 15,
	41, 42, 40, -27, 55, 14, -22, 24, 40, -9,
	11, -2, -31, -30, -2, -2, 9, 13, -28, 12,
	-24, 16, -29, 40, -37, 25, 25, 13, 9, 9,
	-5, -2, 40, 12, -5, 9, -36, -30, 18, -28,
	9, 12, 9, 16, 8, -17, 26, 13, 13, -20,
	9, -7, 40, 9, -5, 9, -19, 26, 13, 9,
	14, -24, 12, 40, 16, -24, 13, -38, -20, -21,
	9, 9, -7, 16, 13, -34, 14, 9, 8, -16,
	14, 36, 37, 38, 29, -21, 14, 9, -18, 14,
	9, -24, 14, -2, 10, 10, 10, 10, 14, 14,
	-25, 9, 42, 42, 40, 31, 9, 9, 9, 9,
	9,
}
var mmDef = [...]int{

	0, -2, 0, 4, 6, 0, 10, 69, 0, 12,
	13, 0, 0, 1, 3, 0, 5, 9, 0, 8,
	0, 0, 28, 106, 107, 108, 109, 110, 111, 112,
	113, 114, 115, 116, 117, 118, 119, 120, 0, 0,
	2, 7, 73, 0, -2, -2, -2, 11, 0, 31,
	31, 0, 79, 0, 27, 35, 35, 68, 74, 0,
	0, 0, 0, 0, 0, 0, 32, 0, 0, 0,
	0, 0, 0, 66, 80, 0, 79, 0, 36, 0,
	0, 29, 46, 47, 48, 49, 50, 51, 52, 0,
	0, 101, 102, 0, 0, 0, 104, 0, 0, 0,
	56, 0, 53, 54, 55, 29, 0, 0, 75, 76,
	77, 78, 0, 0, 0, 0, 88, 89, 0, 0,
	96, 97, 98, 99, 100, 67, 16, 0, 0, 0,
	0, 0, 0, 65, 103, 105, 81, 0, 0, 92,
	85, 93, 0, 0, 23, 0, 0, 31, 43, 37,
	0, 0, 44, 30, 0, 34, 60, 64, 0, 0,
	0, 90, 0, 94, 0, 15, 0, 18, 31, 35,
	38, 0, 45, 40, 0, 33, 0, 0, 79, 0,
	0, 84, 91, 0, 95, 87, 25, 0, 35, 0,
	39, 41, 0, 14, 62, 0, 0, 83, 0, 0,
	17, 0, 0, 0, 0, 0, 58, 42, 0, 59,
	82, 86, 24, 0, 0, 0, 0, 0, 57, 61,
	0, 26, 0, 0, 0, 0, 63, 19, 20, 21,
	22,
}
var mmTok1 = [...]int{

	1,
}
var mmTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57,
}
var mmTok3 = [...]int{
	0,
}

var mmErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	mmDebug        = 0
	mmErrorVerbose = false
)

type mmLexer interface {
	Lex(lval *mmSymType) int
	Error(s string)
}

type mmParser interface {
	Parse(mmLexer) int
	Lookahead() int
}

type mmParserImpl struct {
	lval  mmSymType
	stack [mmInitialStackSize]mmSymType
	char  int
}

func (p *mmParserImpl) Lookahead() int {
	return p.char
}

func mmNewParser() mmParser {
	return &mmParserImpl{}
}

const mmFlag = -1000

func mmTokname(c int) string {
	if c >= 1 && c-1 < len(mmToknames) {
		if mmToknames[c-1] != "" {
			return mmToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func mmStatname(s int) string {
	if s >= 0 && s < len(mmStatenames) {
		if mmStatenames[s] != "" {
			return mmStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func mmErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !mmErrorVerbose {
		return "syntax error"
	}

	for _, e := range mmErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + mmTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := mmPact[state]
	for tok := TOKSTART; tok-1 < len(mmToknames); tok++ {
		if n := base + tok; n >= 0 && n < mmLast && mmChk[mmAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if mmDef[state] == -2 {
		i := 0
		for mmExca[i] != -1 || mmExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; mmExca[i] >= 0; i += 2 {
			tok := mmExca[i]
			if tok < TOKSTART || mmExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if mmExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += mmTokname(tok)
	}
	return res
}

func mmlex1(lex mmLexer, lval *mmSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = mmTok1[0]
		goto out
	}
	if char < len(mmTok1) {
		token = mmTok1[char]
		goto out
	}
	if char >= mmPrivate {
		if char < mmPrivate+len(mmTok2) {
			token = mmTok2[char-mmPrivate]
			goto out
		}
	}
	for i := 0; i < len(mmTok3); i += 2 {
		token = mmTok3[i+0]
		if token == char {
			token = mmTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = mmTok2[1] /* unknown char */
	}
	if mmDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", mmTokname(token), uint(char))
	}
	return char, token
}

func mmParse(mmlex mmLexer) int {
	return mmNewParser().Parse(mmlex)
}

func (mmrcvr *mmParserImpl) Parse(mmlex mmLexer) int {
	var mmn int
	var mmVAL mmSymType
	var mmDollar []mmSymType
	_ = mmDollar // silence set and not used
	mmS := mmrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	mmstate := 0
	mmrcvr.char = -1
	mmtoken := -1 // mmrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		mmstate = -1
		mmrcvr.char = -1
		mmtoken = -1
	}()
	mmp := -1
	goto mmstack

ret0:
	return 0

ret1:
	return 1

mmstack:
	/* put a state and value onto the stack */
	if mmDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", mmTokname(mmtoken), mmStatname(mmstate))
	}

	mmp++
	if mmp >= len(mmS) {
		nyys := make([]mmSymType, len(mmS)*2)
		copy(nyys, mmS)
		mmS = nyys
	}
	mmS[mmp] = mmVAL
	mmS[mmp].yys = mmstate

mmnewstate:
	mmn = mmPact[mmstate]
	if mmn <= mmFlag {
		goto mmdefault /* simple state */
	}
	if mmrcvr.char < 0 {
		mmrcvr.char, mmtoken = mmlex1(mmlex, &mmrcvr.lval)
	}
	mmn += mmtoken
	if mmn < 0 || mmn >= mmLast {
		goto mmdefault
	}
	mmn = mmAct[mmn]
	if mmChk[mmn] == mmtoken { /* valid shift */
		mmrcvr.char = -1
		mmtoken = -1
		mmVAL = mmrcvr.lval
		mmstate = mmn
		if Errflag > 0 {
			Errflag--
		}
		goto mmstack
	}

mmdefault:
	/* default state action */
	mmn = mmDef[mmstate]
	if mmn == -2 {
		if mmrcvr.char < 0 {
			mmrcvr.char, mmtoken = mmlex1(mmlex, &mmrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if mmExca[xi+0] == -1 && mmExca[xi+1] == mmstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			mmn = mmExca[xi+0]
			if mmn < 0 || mmn == mmtoken {
				break
			}
		}
		mmn = mmExca[xi+1]
		if mmn < 0 {
			goto ret0
		}
	}
	if mmn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			mmlex.Error(mmErrorMessage(mmstate, mmtoken))
			Nerrs++
			if mmDebug >= 1 {
				__yyfmt__.Printf("%s", mmStatname(mmstate))
				__yyfmt__.Printf(" saw %s\n", mmTokname(mmtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for mmp >= 0 {
				mmn = mmPact[mmS[mmp].yys] + mmErrCode
				if mmn >= 0 && mmn < mmLast {
					mmstate = mmAct[mmn] /* simulate a shift of "error" */
					if mmChk[mmstate] == mmErrCode {
						goto mmstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if mmDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", mmS[mmp].yys)
				}
				mmp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if mmDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", mmTokname(mmtoken))
			}
			if mmtoken == mmEofCode {
				goto ret1
			}
			mmrcvr.char = -1
			mmtoken = -1
			goto mmnewstate /* try again in the same state */
		}
	}

	/* reduction by production mmn */
	if mmDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", mmn, mmStatname(mmstate))
	}

	mmnt := mmn
	mmpt := mmp
	_ = mmpt // guard against "declared and not used"

	mmp -= mmR2[mmn]
	// mmp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if mmp+1 >= len(mmS) {
		nyys := make([]mmSymType, len(mmS)*2)
		copy(nyys, mmS)
		mmS = nyys
	}
	mmVAL = mmS[mmp+1]

	/* consult goto table to find next state */
	mmn = mmR1[mmn]
	mmg := mmPgo[mmn]
	mmj := mmg + mmS[mmp].yys + 1

	if mmj >= mmLast {
		mmstate = mmAct[mmg]
	} else {
		mmstate = mmAct[mmj]
		if mmChk[mmstate] != -mmn {
			mmstate = mmAct[mmg]
		}
	}
	// dummy call; replaced with literal code
	switch mmnt {

	case 1:
		mmDollar = mmS[mmpt-2 : mmpt+1]
		//line martian/syntax/grammar.y:94
		{
			{
				global := NewAst(mmDollar[2].decs, nil, mmDollar[2].srcfile)
				global.Includes = mmDollar[1].includes
				mmlex.(*mmLexInfo).global = global
			}
		}
	case 2:
		mmDollar = mmS[mmpt-3 : mmpt+1]
		//line martian/syntax/grammar.y:100
		{
			{
				global := NewAst(mmDollar[2].decs, mmDollar[3].call, mmDollar[2].srcfile)
				global.Includes = mmDollar[1].includes
				mmlex.(*mmLexInfo).global = global
			}
		}
	case 3:
		mmDollar = mmS[mmpt-2 : mmpt+1]
		//line martian/syntax/grammar.y:106
		{
			{
				global := NewAst(nil, mmDollar[2].call, mmDollar[2].srcfile)
				global.Includes = mmDollar[1].includes
				mmlex.(*mmLexInfo).global = global
			}
		}
	case 4:
		mmDollar = mmS[mmpt-1 : mmpt+1]
		//line martian/syntax/grammar.y:112
		{
			{
				global := NewAst(mmDollar[1].decs, nil, mmDollar[1].srcfile)
				mmlex.(*mmLexInfo).global = global
			}
		}
	case 5:
		mmDollar = mmS[mmpt-2 : mmpt+1]
		//line martian/syntax/grammar.y:117
		{
			{
				global := NewAst(mmDollar[1].decs, mmDollar[2].call, mmDollar[1].srcfile)
				mmlex.(*mmLexInfo).global = global
			}
		}
	case 6:
		mmDollar = mmS[mmpt-1 : mmpt+1]
		//line martian/syntax/grammar.y:122
		{
			{
				global := NewAst(nil, mmDollar[1].call, mmDollar[1].srcfile)
				mmlex.(*mmLexInfo).global = global
			}
		}
	case 7:
		mmDollar = mmS[mmpt-3 : mmpt+1]
		//line martian/syntax/grammar.y:130
		{
			{
				mmVAL.includes = append(mmDollar[1].includes, &Include{
					Node:  NewAstNode(mmDollar[2].loc, mmDollar[2].srcfile),
					Value: mmDollar[3].intern.unquote(mmDollar[3].val),
				})
			}
		}
	case 8:
		mmDollar = mmS[mmpt-2 : mmpt+1]
		//line martian/syntax/grammar.y:136
		{
			{
				mmVAL.includes = []*Include{
					{
						Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
						Value: mmDollar[2].intern.unquote(mmDollar[2].val),
					},
				}
			}
		}
	case 9:
		mmDollar = mmS[mmpt-2 : mmpt+1]
		//line martian/syntax/grammar.y:146
		{
			{
				mmVAL.decs = append(mmDollar[1].decs, mmDollar[2].dec)
			}
		}
	case 10:
		mmDollar = mmS[mmpt-1 : mmpt+1]
		//line martian/syntax/grammar.y:148
		{
			{
				mmVAL.decs = []Dec{mmDollar[1].dec}
			}
		}
	case 11:
		mmDollar = mmS[mmpt-3 : mmpt+1]
		//line martian/syntax/grammar.y:153
		{
			{
				mmVAL.dec = &UserType{
					Node: NewAstNode(mmDollar[2].loc, mmDollar[2].srcfile),
					Id:   mmDollar[2].intern.Get(mmDollar[2].val),
				}
			}
		}
	case 14:
		mmDollar = mmS[mmpt-11 : mmpt+1]
		//line martian/syntax/grammar.y:163
		{
			{
				mmVAL.dec = &Pipeline{
					Node:      NewAstNode(mmDollar[2].loc, mmDollar[2].srcfile),
					Id:        mmDollar[2].intern.Get(mmDollar[2].val),
					InParams:  mmDollar[4].i_params,
					OutParams: mmDollar[5].o_params,
					Calls:     mmDollar[8].calls,
					Callables: &Callables{Table: make(map[string]Callable)},
					Ret:       mmDollar[9].retstm,
					Retain:    mmDollar[10].plretains,
				}
			}
		}
	case 15:
		mmDollar = mmS[mmpt-10 : mmpt+1]
		//line martian/syntax/grammar.y:177
		{
			{
				mmVAL.dec = &Stage{
					Node:      NewAstNode(mmDollar[2].loc, mmDollar[2].srcfile),
					Id:        mmDollar[2].intern.Get(mmDollar[2].val),
					InParams:  mmDollar[4].i_params,
					OutParams: mmDollar[5].o_params,
					Src:       mmDollar[6].src,
					ChunkIns:  mmDollar[8].par_tuple.Ins,
					ChunkOuts: mmDollar[8].par_tuple.Outs,
					Split:     mmDollar[8].par_tuple.Present,
					Resources: mmDollar[9].res,
					Retain:    mmDollar[10].stretains,
				}
			}
		}
	case 16:
		mmDollar = mmS[mmpt-0 : mmpt+1]
		//line martian/syntax/grammar.y:194
		{
			{
				mmVAL.res = nil
			}
		}
	case 17:
		mmDollar = mmS[mmpt-4 : mmpt+1]
		//line martian/syntax/grammar.y:196
		{
			{
				mmDollar[3].res.Node = NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile)
				mmVAL.res = mmDollar[3].res
			}
		}
	case 18:
		mmDollar = mmS[mmpt-0 : mmpt+1]
		//line martian/syntax/grammar.y:204
		{
			{
				mmVAL.res = new(Resources)
			}
		}
	case 19:
		mmDollar = mmS[mmpt-5 : mmpt+1]
		//line martian/syntax/grammar.y:206
		{
			{
				n := NewAstNode(mmDollar[2].loc, mmDollar[2].srcfile)
				mmDollar[1].res.ThreadNode = &n
				i := parseInt(mmDollar[4].val)
				mmDollar[1].res.Threads = int16(i)
				mmVAL.res = mmDollar[1].res
			}
		}
	case 20:
		mmDollar = mmS[mmpt-5 : mmpt+1]
		//line martian/syntax/grammar.y:214
		{
			{
				n := NewAstNode(mmDollar[2].loc, mmDollar[2].srcfile)
				mmDollar[1].res.MemNode = &n
				i := parseInt(mmDollar[4].val)
				mmDollar[1].res.MemGB = int16(i)
				mmVAL.res = mmDollar[1].res
			}
		}
	case 21:
		mmDollar = mmS[mmpt-5 : mmpt+1]
		//line martian/syntax/grammar.y:222
		{
			{
				n := NewAstNode(mmDollar[2].loc, mmDollar[2].srcfile)
				mmDollar[1].res.SpecialNode = &n
				mmDollar[1].res.Special = mmDollar[4].intern.unquote(mmDollar[4].val)
				mmVAL.res = mmDollar[1].res
			}
		}
	case 22:
		mmDollar = mmS[mmpt-5 : mmpt+1]
		//line martian/syntax/grammar.y:229
		{
			{
				n := NewAstNode(mmDollar[2].loc, mmDollar[2].srcfile)
				mmDollar[1].res.VolatileNode = &n
				mmDollar[1].res.StrictVolatile = true
				mmVAL.res = mmDollar[1].res
			}
		}
	case 23:
		mmDollar = mmS[mmpt-0 : mmpt+1]
		//line martian/syntax/grammar.y:239
		{
			{
				mmVAL.stretains = nil
			}
		}
	case 24:
		mmDollar = mmS[mmpt-4 : mmpt+1]
		//line martian/syntax/grammar.y:241
		{
			{
				mmVAL.stretains = &RetainParams{
					Node:   NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Params: mmDollar[3].retains,
				}
			}
		}
	case 25:
		mmDollar = mmS[mmpt-0 : mmpt+1]
		//line martian/syntax/grammar.y:251
		{
			{
				mmVAL.retains = nil
			}
		}
	case 26:
		mmDollar = mmS[mmpt-3 : mmpt+1]
		//line martian/syntax/grammar.y:253
		{
			{
				mmVAL.retains = append(mmDollar[1].retains, &RetainParam{
					Node: NewAstNode(mmDollar[2].loc, mmDollar[2].srcfile),
					Id:   mmDollar[2].intern.Get(mmDollar[2].val),
				})
			}
		}
	case 27:
		mmDollar = mmS[mmpt-3 : mmpt+1]
		//line martian/syntax/grammar.y:264
		{
			{
				idd := append(mmDollar[1].val, '.')
				mmVAL.val = append(idd, mmDollar[3].val...)
			}
		}
	case 28:
		mmDollar = mmS[mmpt-1 : mmpt+1]
		//line martian/syntax/grammar.y:269
		{
			{
				// set capacity == length so append doesn't overwrite
				// other parts of the buffer later.
				mmVAL.val = mmDollar[1].val[:len(mmDollar[1].val):len(mmDollar[1].val)]
			}
		}
	case 29:
		mmDollar = mmS[mmpt-0 : mmpt+1]
		//line martian/syntax/grammar.y:278
		{
			{
				mmVAL.arr = 0
			}
		}
	case 30:
		mmDollar = mmS[mmpt-3 : mmpt+1]
		//line martian/syntax/grammar.y:280
		{
			{
				mmVAL.arr++
			}
		}
	case 31:
		mmDollar = mmS[mmpt-0 : mmpt+1]
		//line martian/syntax/grammar.y:285
		{
			{
				mmVAL.i_params = &InParams{Table: make(map[string]*InParam)}
			}
		}
	case 32:
		mmDollar = mmS[mmpt-2 : mmpt+1]
		//line martian/syntax/grammar.y:287
		{
			{
				mmDollar[1].i_params.List = append(mmDollar[1].i_params.List, mmDollar[2].inparam)
				mmVAL.i_params = mmDollar[1].i_params
			}
		}
	case 33:
		mmDollar = mmS[mmpt-6 : mmpt+1]
		//line martian/syntax/grammar.y:295
		{
			{
				mmVAL.inparam = &InParam{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Tname:    mmDollar[2].intern.Get(mmDollar[2].val),
					ArrayDim: mmDollar[3].arr,
					Id:       mmDollar[4].intern.Get(mmDollar[4].val),
					Help:     unquote(mmDollar[5].val),
				}
			}
		}
	case 34:
		mmDollar = mmS[mmpt-5 : mmpt+1]
		//line martian/syntax/grammar.y:303
		{
			{
				mmVAL.inparam = &InParam{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Tname:    mmDollar[2].intern.Get(mmDollar[2].val),
					ArrayDim: mmDollar[3].arr,
					Id:       mmDollar[4].intern.Get(mmDollar[4].val),
				}
			}
		}
	case 35:
		mmDollar = mmS[mmpt-0 : mmpt+1]
		//line martian/syntax/grammar.y:313
		{
			{
				mmVAL.o_params = &OutParams{Table: make(map[string]*OutParam)}
			}
		}
	case 36:
		mmDollar = mmS[mmpt-2 : mmpt+1]
		//line martian/syntax/grammar.y:315
		{
			{
				mmDollar[1].o_params.List = append(mmDollar[1].o_params.List, mmDollar[2].outparam)
				mmVAL.o_params = mmDollar[1].o_params
			}
		}
	case 37:
		mmDollar = mmS[mmpt-4 : mmpt+1]
		//line martian/syntax/grammar.y:323
		{
			{
				mmVAL.outparam = &OutParam{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Tname:    mmDollar[2].intern.Get(mmDollar[2].val),
					ArrayDim: mmDollar[3].arr,
					Id:       default_out_name,
				}
			}
		}
	case 38:
		mmDollar = mmS[mmpt-5 : mmpt+1]
		//line martian/syntax/grammar.y:330
		{
			{
				mmVAL.outparam = &OutParam{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Tname:    mmDollar[2].intern.Get(mmDollar[2].val),
					ArrayDim: mmDollar[3].arr,
					Id:       default_out_name,
					Help:     unquote(mmDollar[4].val),
				}
			}
		}
	case 39:
		mmDollar = mmS[mmpt-6 : mmpt+1]
		//line martian/syntax/grammar.y:338
		{
			{
				mmVAL.outparam = &OutParam{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Tname:    mmDollar[2].intern.Get(mmDollar[2].val),
					ArrayDim: mmDollar[3].arr,
					Id:       default_out_name,
					Help:     unquote(mmDollar[4].val),
					OutName:  mmDollar[5].intern.unquote(mmDollar[5].val),
				}
			}
		}
	case 40:
		mmDollar = mmS[mmpt-5 : mmpt+1]
		//line martian/syntax/grammar.y:347
		{
			{
				mmVAL.outparam = &OutParam{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Tname:    mmDollar[2].intern.Get(mmDollar[2].val),
					ArrayDim: mmDollar[3].arr,
					Id:       mmDollar[4].intern.Get(mmDollar[4].val),
				}
			}
		}
	case 41:
		mmDollar = mmS[mmpt-6 : mmpt+1]
		//line martian/syntax/grammar.y:354
		{
			{
				mmVAL.outparam = &OutParam{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Tname:    mmDollar[2].intern.Get(mmDollar[2].val),
					ArrayDim: mmDollar[3].arr,
					Id:       mmDollar[4].intern.Get(mmDollar[4].val),
					Help:     unquote(mmDollar[5].val),
				}
			}
		}
	case 42:
		mmDollar = mmS[mmpt-7 : mmpt+1]
		//line martian/syntax/grammar.y:362
		{
			{
				mmVAL.outparam = &OutParam{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Tname:    mmDollar[2].intern.Get(mmDollar[2].val),
					ArrayDim: mmDollar[3].arr,
					Id:       mmDollar[4].intern.Get(mmDollar[4].val),
					Help:     unquote(mmDollar[5].val),
					OutName:  mmDollar[6].intern.unquote(mmDollar[6].val),
				}
			}
		}
	case 43:
		mmDollar = mmS[mmpt-4 : mmpt+1]
		//line martian/syntax/grammar.y:374
		{
			{
				stagecodeParts := strings.Split(mmDollar[3].intern.unquote(mmDollar[3].val), " ")
				mmVAL.src = &SrcParam{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Lang: StageLanguage(mmDollar[2].intern.Get(mmDollar[2].val)),
					Path: stagecodeParts[0],
					Args: stagecodeParts[1:],
				}
			}
		}
	case 56:
		mmDollar = mmS[mmpt-0 : mmpt+1]
		//line martian/syntax/grammar.y:409
		{
			{
				mmVAL.par_tuple = paramsTuple{
					Present: false,
					Ins:     &InParams{Table: make(map[string]*InParam)},
					Outs:    &OutParams{Table: make(map[string]*OutParam)},
				}
			}
		}
	case 57:
		mmDollar = mmS[mmpt-6 : mmpt+1]
		//line martian/syntax/grammar.y:417
		{
			{
				mmVAL.par_tuple = paramsTuple{
					Present: true,
					Ins:     mmDollar[4].i_params,
					Outs:    mmDollar[5].o_params,
				}
			}
		}
	case 58:
		mmDollar = mmS[mmpt-5 : mmpt+1]
		//line martian/syntax/grammar.y:423
		{
			{
				mmVAL.par_tuple = paramsTuple{
					Present: true,
					Ins:     mmDollar[3].i_params,
					Outs:    mmDollar[4].o_params,
				}
			}
		}
	case 59:
		mmDollar = mmS[mmpt-4 : mmpt+1]
		//line martian/syntax/grammar.y:432
		{
			{
				mmVAL.retstm = &ReturnStm{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Bindings: mmDollar[3].bindings,
				}
			}
		}
	case 60:
		mmDollar = mmS[mmpt-0 : mmpt+1]
		//line martian/syntax/grammar.y:440
		{
			{
				mmVAL.plretains = nil
			}
		}
	case 61:
		mmDollar = mmS[mmpt-4 : mmpt+1]
		//line martian/syntax/grammar.y:442
		{
			{
				mmVAL.plretains = &PipelineRetains{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Refs: mmDollar[3].reflist,
				}
			}
		}
	case 62:
		mmDollar = mmS[mmpt-0 : mmpt+1]
		//line martian/syntax/grammar.y:449
		{
			{
				mmVAL.reflist = nil
			}
		}
	case 63:
		mmDollar = mmS[mmpt-3 : mmpt+1]
		//line martian/syntax/grammar.y:451
		{
			{
				mmVAL.reflist = append(mmDollar[1].reflist, mmDollar[2].rexp)
			}
		}
	case 64:
		mmDollar = mmS[mmpt-2 : mmpt+1]
		//line martian/syntax/grammar.y:455
		{
			{
				mmVAL.calls = append(mmDollar[1].calls, mmDollar[2].call)
			}
		}
	case 65:
		mmDollar = mmS[mmpt-1 : mmpt+1]
		//line martian/syntax/grammar.y:457
		{
			{
				mmVAL.calls = []*CallStm{mmDollar[1].call}
			}
		}
	case 66:
		mmDollar = mmS[mmpt-6 : mmpt+1]
		//line martian/syntax/grammar.y:462
		{
			{
				id := mmDollar[3].intern.Get(mmDollar[3].val)
				mmVAL.call = &CallStm{
					Node:      NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Modifiers: mmDollar[2].modifiers,
					Id:        id,
					DecId:     id,
					Bindings:  mmDollar[5].bindings,
				}
			}
		}
	case 67:
		mmDollar = mmS[mmpt-8 : mmpt+1]
		//line martian/syntax/grammar.y:471
		{
			{
				mmVAL.call = &CallStm{
					Node:      NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Modifiers: mmDollar[2].modifiers,
					Id:        mmDollar[5].intern.Get(mmDollar[5].val),
					DecId:     mmDollar[3].intern.Get(mmDollar[3].val),
					Bindings:  mmDollar[7].bindings,
				}
			}
		}
	case 68:
		mmDollar = mmS[mmpt-5 : mmpt+1]
		//line martian/syntax/grammar.y:479
		{
			{
				mmDollar[1].call.Modifiers.Bindings = mmDollar[4].bindings
				mmVAL.call = mmDollar[1].call
			}
		}
	case 69:
		mmDollar = mmS[mmpt-0 : mmpt+1]
		//line martian/syntax/grammar.y:487
		{
			{
				mmVAL.modifiers = new(Modifiers)
			}
		}
	case 70:
		mmDollar = mmS[mmpt-2 : mmpt+1]
		//line martian/syntax/grammar.y:489
		{
			{
				mmVAL.modifiers.Local = true
			}
		}
	case 71:
		mmDollar = mmS[mmpt-2 : mmpt+1]
		//line martian/syntax/grammar.y:491
		{
			{
				mmVAL.modifiers.Preflight = true
			}
		}
	case 72:
		mmDollar = mmS[mmpt-2 : mmpt+1]
		//line martian/syntax/grammar.y:493
		{
			{
				mmVAL.modifiers.Volatile = true
			}
		}
	case 73:
		mmDollar = mmS[mmpt-0 : mmpt+1]
		//line martian/syntax/grammar.y:498
		{
			{
				mmVAL.bindings = &BindStms{
					Node:  NewAstNode(mmDollar[0].loc, mmDollar[0].srcfile),
					Table: make(map[string]*BindStm),
				}
			}
		}
	case 74:
		mmDollar = mmS[mmpt-2 : mmpt+1]
		//line martian/syntax/grammar.y:503
		{
			{
				mmDollar[1].bindings.List = append(mmDollar[1].bindings.List, mmDollar[2].binding)
				mmVAL.bindings = mmDollar[1].bindings
			}
		}
	case 75:
		mmDollar = mmS[mmpt-4 : mmpt+1]
		//line martian/syntax/grammar.y:511
		{
			{
				mmVAL.binding = &BindStm{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Id:   local,
					Exp:  mmDollar[3].vexp,
				}
			}
		}
	case 76:
		mmDollar = mmS[mmpt-4 : mmpt+1]
		//line martian/syntax/grammar.y:517
		{
			{
				mmVAL.binding = &BindStm{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Id:   preflight,
					Exp:  mmDollar[3].vexp,
				}
			}
		}
	case 77:
		mmDollar = mmS[mmpt-4 : mmpt+1]
		//line martian/syntax/grammar.y:523
		{
			{
				mmVAL.binding = &BindStm{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Id:   volatile,
					Exp:  mmDollar[3].vexp,
				}
			}
		}
	case 78:
		mmDollar = mmS[mmpt-4 : mmpt+1]
		//line martian/syntax/grammar.y:529
		{
			{
				mmVAL.binding = &BindStm{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Id:   disabled,
					Exp:  mmDollar[3].rexp,
				}
			}
		}
	case 79:
		mmDollar = mmS[mmpt-0 : mmpt+1]
		//line martian/syntax/grammar.y:537
		{
			{
				mmVAL.bindings = &BindStms{
					Node:  NewAstNode(mmDollar[0].loc, mmDollar[0].srcfile),
					Table: make(map[string]*BindStm),
				}
			}
		}
	case 80:
		mmDollar = mmS[mmpt-2 : mmpt+1]
		//line martian/syntax/grammar.y:542
		{
			{
				mmDollar[1].bindings.List = append(mmDollar[1].bindings.List, mmDollar[2].binding)
				mmVAL.bindings = mmDollar[1].bindings
			}
		}
	case 81:
		mmDollar = mmS[mmpt-4 : mmpt+1]
		//line martian/syntax/grammar.y:550
		{
			{
				mmVAL.binding = &BindStm{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Id:   mmDollar[1].intern.Get(mmDollar[1].val),
					Exp:  mmDollar[3].exp,
				}
			}
		}
	case 82:
		mmDollar = mmS[mmpt-8 : mmpt+1]
		//line martian/syntax/grammar.y:556
		{
			{
				mmVAL.binding = &BindStm{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Id:   mmDollar[1].intern.Get(mmDollar[1].val),
					Exp: &ValExp{
						Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
						Kind:  KindArray,
						Value: mmDollar[5].exps,
					},
					Sweep: true,
				}
			}
		}
	case 83:
		mmDollar = mmS[mmpt-7 : mmpt+1]
		//line martian/syntax/grammar.y:567
		{
			{
				mmVAL.binding = &BindStm{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Id:   mmDollar[1].intern.Get(mmDollar[1].val),
					Exp: &ValExp{
						Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
						Kind:  KindArray,
						Value: mmDollar[5].exps,
					},
					Sweep: true,
				}
			}
		}
	case 84:
		mmDollar = mmS[mmpt-3 : mmpt+1]
		//line martian/syntax/grammar.y:581
		{
			{
				mmVAL.exps = append(mmDollar[1].exps, mmDollar[3].exp)
			}
		}
	case 85:
		mmDollar = mmS[mmpt-1 : mmpt+1]
		//line martian/syntax/grammar.y:583
		{
			{
				mmVAL.exps = []Exp{mmDollar[1].exp}
			}
		}
	case 86:
		mmDollar = mmS[mmpt-5 : mmpt+1]
		//line martian/syntax/grammar.y:588
		{
			{
				mmDollar[1].kvpairs[unquote(mmDollar[3].val)] = mmDollar[5].exp
				mmVAL.kvpairs = mmDollar[1].kvpairs
			}
		}
	case 87:
		mmDollar = mmS[mmpt-3 : mmpt+1]
		//line martian/syntax/grammar.y:593
		{
			{
				mmVAL.kvpairs = map[string]Exp{unquote(mmDollar[1].val): mmDollar[3].exp}
			}
		}
	case 88:
		mmDollar = mmS[mmpt-1 : mmpt+1]
		//line martian/syntax/grammar.y:598
		{
			{
				mmVAL.exp = mmDollar[1].vexp
			}
		}
	case 89:
		mmDollar = mmS[mmpt-1 : mmpt+1]
		//line martian/syntax/grammar.y:600
		{
			{
				mmVAL.exp = mmDollar[1].rexp
			}
		}
	case 90:
		mmDollar = mmS[mmpt-3 : mmpt+1]
		//line martian/syntax/grammar.y:604
		{
			{
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindArray,
					Value: mmDollar[2].exps,
				}
			}
		}
	case 91:
		mmDollar = mmS[mmpt-4 : mmpt+1]
		//line martian/syntax/grammar.y:610
		{
			{
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindArray,
					Value: mmDollar[2].exps,
				}
			}
		}
	case 92:
		mmDollar = mmS[mmpt-2 : mmpt+1]
		//line martian/syntax/grammar.y:616
		{
			{
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindArray,
					Value: make([]Exp, 0),
				}
			}
		}
	case 93:
		mmDollar = mmS[mmpt-2 : mmpt+1]
		//line martian/syntax/grammar.y:622
		{
			{
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindMap,
					Value: make(map[string]interface{}, 0),
				}
			}
		}
	case 94:
		mmDollar = mmS[mmpt-3 : mmpt+1]
		//line martian/syntax/grammar.y:628
		{
			{
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindMap,
					Value: mmDollar[2].kvpairs,
				}
			}
		}
	case 95:
		mmDollar = mmS[mmpt-4 : mmpt+1]
		//line martian/syntax/grammar.y:634
		{
			{
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindMap,
					Value: mmDollar[2].kvpairs,
				}
			}
		}
	case 96:
		mmDollar = mmS[mmpt-1 : mmpt+1]
		//line martian/syntax/grammar.y:640
		{
			{ // Lexer guarantees parseable float strings.
				f := parseFloat(mmDollar[1].val)
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindFloat,
					Value: f,
				}
			}
		}
	case 97:
		mmDollar = mmS[mmpt-1 : mmpt+1]
		//line martian/syntax/grammar.y:649
		{
			{ // Lexer guarantees parseable int strings.
				i := parseInt(mmDollar[1].val)
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindInt,
					Value: i,
				}
			}
		}
	case 98:
		mmDollar = mmS[mmpt-1 : mmpt+1]
		//line martian/syntax/grammar.y:658
		{
			{
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindString,
					Value: unquote(mmDollar[1].val),
				}
			}
		}
	case 100:
		mmDollar = mmS[mmpt-1 : mmpt+1]
		//line martian/syntax/grammar.y:665
		{
			{
				mmVAL.vexp = &ValExp{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind: KindNull,
				}
			}
		}
	case 101:
		mmDollar = mmS[mmpt-1 : mmpt+1]
		//line martian/syntax/grammar.y:673
		{
			{
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindBool,
					Value: true,
				}
			}
		}
	case 102:
		mmDollar = mmS[mmpt-1 : mmpt+1]
		//line martian/syntax/grammar.y:679
		{
			{
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindBool,
					Value: false,
				}
			}
		}
	case 103:
		mmDollar = mmS[mmpt-3 : mmpt+1]
		//line martian/syntax/grammar.y:687
		{
			{
				mmVAL.rexp = &RefExp{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:     KindCall,
					Id:       mmDollar[1].intern.Get(mmDollar[1].val),
					OutputId: mmDollar[3].intern.Get(mmDollar[3].val),
				}
			}
		}
	case 104:
		mmDollar = mmS[mmpt-1 : mmpt+1]
		//line martian/syntax/grammar.y:694
		{
			{
				mmVAL.rexp = &RefExp{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:     KindCall,
					Id:       mmDollar[1].intern.Get(mmDollar[1].val),
					OutputId: default_out_name,
				}
			}
		}
	case 105:
		mmDollar = mmS[mmpt-3 : mmpt+1]
		//line martian/syntax/grammar.y:701
		{
			{
				mmVAL.rexp = &RefExp{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind: KindSelf,
					Id:   mmDollar[3].intern.Get(mmDollar[3].val),
				}
			}
		}
	}
	goto mmstack /* stack new state and value */
}
