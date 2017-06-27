//line expressions.y:2
package expressions

import __yyfmt__ "fmt"

//line expressions.y:2
import (
	"fmt"
	"github.com/osteele/liquid/generics"
	"reflect"
)

func init() {
	// This allows adding and removing references to fmt in the rules below,
	// without having to edit the import statement to avoid erorrs each time.
	_ = fmt.Sprint("")
}

//line expressions.y:16
type yySymType struct {
	yys      int
	name     string
	val      interface{}
	f        func(Context) interface{}
	loopmods LoopModifiers
}

const LITERAL = 57346
const IDENTIFIER = 57347
const KEYWORD = 57348
const RELATION = 57349
const ASSIGN = 57350
const LOOP = 57351
const EQ = 57352
const FOR = 57353
const IN = 57354

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LITERAL",
	"IDENTIFIER",
	"KEYWORD",
	"RELATION",
	"ASSIGN",
	"LOOP",
	"EQ",
	"FOR",
	"IN",
	"'.'",
	"'|'",
	"'<'",
	"'>'",
	"';'",
	"'='",
	"'['",
	"']'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 43

var yyAct = [...]int{

	29, 6, 16, 14, 5, 14, 14, 17, 18, 15,
	32, 15, 15, 19, 36, 9, 24, 25, 26, 27,
	13, 23, 20, 31, 28, 30, 35, 13, 7, 8,
	33, 12, 3, 4, 21, 22, 7, 8, 10, 1,
	34, 11, 2,
}
var yyPact = [...]int{

	24, -1000, -2, 33, 26, 6, -8, -1000, -1000, -1000,
	-5, -1000, 10, 29, 16, 32, 32, 32, 32, 32,
	32, -1000, 32, -1000, -10, -7, -7, -7, 13, -7,
	6, -7, -1000, -1000, 9, -1000, -1000,
}
var yyPgo = [...]int{

	0, 0, 42, 4, 41, 40, 39,
}
var yyR1 = [...]int{

	0, 6, 6, 6, 4, 5, 5, 1, 1, 1,
	1, 3, 3, 3, 2, 2, 2, 2,
}
var yyR2 = [...]int{

	0, 2, 5, 2, 5, 0, 2, 1, 1, 3,
	4, 1, 3, 4, 1, 3, 3, 3,
}
var yyChk = [...]int{

	-1000, -6, -2, 8, 9, -3, -1, 4, 5, 17,
	5, -4, 5, 14, 13, 19, 10, 15, 16, 18,
	12, 5, 6, 5, -1, -1, -1, -1, -3, -1,
	-3, -1, 20, 17, -5, 17, 5,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 14, 11, 7, 8, 1,
	0, 3, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 12, 0, 9, 0, 15, 16, 17, 0, 11,
	5, 13, 10, 2, 0, 4, 6,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 13, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 17,
	15, 18, 16, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 19, 3, 20, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 14,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line expressions.y:32
		{
			yylex.(*lexer).val = yyDollar[1].f
		}
	case 2:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line expressions.y:33
		{
			name, expr := yyDollar[2].name, yyDollar[4].f
			yylex.(*lexer).val = func(ctx Context) interface{} {
				ctx.Set(name, expr(ctx))
				return nil
			}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line expressions.y:40
		{
			yylex.(*lexer).val = yyDollar[2].f
		}
	case 4:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line expressions.y:43
		{
			name, expr, mods := yyDollar[1].name, yyDollar[3].f, yyDollar[4].loopmods
			yyVAL.f = func(ctx Context) interface{} {
				return &Loop{name, expr(ctx), mods}
			}
		}
	case 5:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line expressions.y:51
		{
			yyVAL.loopmods = LoopModifiers{}
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line expressions.y:52
		{
			if yyDollar[2].name != "reversed" {
				panic(ParseError(fmt.Sprintf("undefined loop modifier: %s", yyDollar[2].name)))
			}
			yyDollar[1].loopmods.Reversed = true
			yyVAL.loopmods = yyDollar[1].loopmods
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line expressions.y:62
		{
			val := yyDollar[1].val
			yyVAL.f = func(_ Context) interface{} { return val }
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line expressions.y:63
		{
			name := yyDollar[1].name
			yyVAL.f = func(ctx Context) interface{} { return ctx.Get(name) }
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line expressions.y:64
		{
			e, attr := yyDollar[1].f, yyDollar[3].name
			yyVAL.f = func(ctx Context) interface{} {
				obj := e(ctx)
				ref := reflect.ValueOf(obj)
				switch ref.Kind() {
				case reflect.Map:
					val := ref.MapIndex(reflect.ValueOf(attr))
					if val.Kind() != reflect.Invalid {
						return val.Interface()
					}
				}
				return nil
			}
		}
	case 10:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line expressions.y:79
		{
			e, i := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) interface{} {
				ref := reflect.ValueOf(e(ctx))
				index := reflect.ValueOf(i(ctx))
				switch ref.Kind() {
				case reflect.Array, reflect.Slice:
					switch index.Kind() {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						n := int(index.Int())
						if 0 <= n && n < ref.Len() {
							return ref.Index(n).Interface()
						}
					}
				}
				return nil
			}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line expressions.y:100
		{
			yyVAL.f = makeFilter(yyDollar[1].f, yyDollar[3].name, nil)
		}
	case 13:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line expressions.y:101
		{
			yyVAL.f = makeFilter(yyDollar[1].f, yyDollar[3].name, yyDollar[4].f)
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line expressions.y:106
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) interface{} {
				a, b := fa(ctx), fb(ctx)
				return generics.Equal(a, b)
			}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line expressions.y:113
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) interface{} {
				a, b := fa(ctx), fb(ctx)
				return generics.Less(a, b)
			}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line expressions.y:120
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) interface{} {
				a, b := fa(ctx), fb(ctx)
				return generics.Less(b, a)
			}
		}
	}
	goto yystack /* stack new state and value */
}
