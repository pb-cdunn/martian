//
// Copyright (c) 2014 10X Genomics, Inc. All rights reserved.
//
// Martian formatter tests.
//

package syntax

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	SetEnforcementLevel(EnforceError)
	os.Exit(m.Run())
}

// Checks that the source can be parsed and compiled
func testGood(t *testing.T, src string) *Ast {
	t.Helper()
	if ast, err := yaccParse([]byte(src), new(SourceFile),
		makeStringIntern()); err != nil {
		t.Fatal(err.Error())
		return nil
	} else if err := ast.compile(); err != nil {
		t.Errorf("Failed to compile src: %v\n%s", err, err.Error())
		return nil
	} else {
		return ast
	}
}

func TestSimplePipe(t *testing.T) {
	t.Parallel()
	if ast := testGood(t, `
# File comment

# Stage comment
# This describes the stage.
stage SUM_SQUARES(
    in  float[] values,
    # sum comment
    out float   sum,
    src py      "stages/sum_squares",
)

pipeline SUM_SQUARE_PIPELINE(
    in  float[] values,
    out float   sum,
)
{
    call SUM_SQUARES(
        values = self.values,
    )
    return (
        sum = SUM_SQUARES.sum,
    )
}

call SUM_SQUARE_PIPELINE(
    values = [00.0e-10, 10.0, 2.0e1, 3.0e+1, 400.0e-1, 5e1, 6e+1, 700e-1],
)
`); ast != nil {
		if s := ast.Callables.Table["SUM_SQUARES"]; s == nil {
			t.Error("No callable named SUM_SQUARES found")
		} else {
			if len(s.getNode().Comments) != 2 {
				t.Errorf("Incorrect stage comment count %d", len(s.getNode().Comments))
			} else {
				if s.getNode().Comments[0] != "# Stage comment" {
					t.Errorf("Expected comment '# Stage comment', got %s",
						s.getNode().Comments[0])
				}
				if s.getNode().Comments[1] != "# This describes the stage." {
					t.Errorf("Expected comment '# This describes the stage.', got %s",
						s.getNode().Comments[1])
				}
			}
			o := s.(*Stage).OutParams
			if len(o.Table) != 1 {
				t.Errorf("Incorrect out param count %d", len(o.Table))
			} else if p := o.Table["sum"]; p == nil {
				t.Error("No out param 'sum' found")
			} else {
				if len(p.getNode().Comments) != 1 {
					t.Errorf("Incorrect param comment count %d", len(p.getNode().Comments))
				} else {
					if p.getNode().Comments[0] != "# sum comment" {
						t.Errorf("Expected comment '# sum comment', got %s",
							p.getNode().Comments[0])
					}
				}
			}
		}
		if ast.Call == nil {
			t.Error("Expected a call.")
		} else {
			if ast.Call.DecId != ast.Call.Id {
				t.Errorf("Expected callable name %s to match alias %s",
					ast.Call.DecId, ast.Call.Id)
			}
			if ast.Call.Bindings == nil || len(ast.Call.Bindings.List) != 1 {
				t.Errorf("Expected 1 binding.")
			} else if binding := ast.Call.Bindings.List[0]; binding.Id != "values" {
				t.Errorf("Expected binding to 'values', got %s",
					binding.Id)
			} else if binding.Exp.getKind() != KindArray {
				t.Errorf("Expected array binding, got %v",
					binding.Exp.getKind())
			} else if val := binding.Exp.ToInterface(); val == nil {
				t.Errorf("Could not get binding value.")
			} else if arr, ok := val.([]interface{}); !ok {
				t.Errorf("Could not convert expression of type %T to array.",
					val)
			} else if len(arr) != 8 {
				t.Errorf("Expected 8-element array, got %d elements.",
					len(arr))
			} else {
				for i, v := range arr {
					if f, ok := v.(float64); !ok {
						t.Errorf("Expected floating point number, got %T",
							v)
					} else if f != float64(10*i) {
						t.Errorf("Expected %d, got %f", 10*i, f)
					}
				}
			}
		}
	}
}

func TestBinding(t *testing.T) {
	t.Parallel()
	if ast := testGood(t, `
stage SUM_SQUARES(
    in  float[] values,
    in  int     threads,
    in  bool    local,
    out float   sum,
    src py      "stages/sum_squares",
)

stage REPORT(
    in  float[] values,
    in  float   sum,
    src py      "stages/report",
)

pipeline SUM_SQUARE_PIPELINE(
    in  float[] values,
    out float   sum,
)
{
    call SUM_SQUARES(
        values  = self.values,
        threads = 1,
        local   = false,
    )
    call REPORT(
        values = self.values,
        sum    = SUM_SQUARES.sum,
    )

    return (
        sum = SUM_SQUARES.sum,
    )
}

call SUM_SQUARE_PIPELINE(
    values = [0.1, 2.0e-1, 3e-1],
)
`); ast != nil {
		if ast.Call == nil {
			t.Error("Expected a call.")
		} else {
			if ast.Call.DecId != ast.Call.Id {
				t.Errorf("Expected callable name %s to match alias %s",
					ast.Call.DecId, ast.Call.Id)
			}
			if ast.Call.Bindings == nil || len(ast.Call.Bindings.List) != 1 {
				t.Errorf("Expected 1 binding.")
			} else if binding := ast.Call.Bindings.List[0]; binding.Id != "values" {
				t.Errorf("Expected binding to 'values', got %s",
					binding.Id)
			} else if binding.Exp.getKind() != KindArray {
				t.Errorf("Expected array binding, got %v",
					binding.Exp.getKind())
			} else if val := binding.Exp.ToInterface(); val == nil {
				t.Errorf("Could not get binding value.")
			} else if arr, ok := val.([]interface{}); !ok {
				t.Errorf("Could not convert expression of type %T to array.",
					val)
			} else if len(arr) != 3 {
				t.Errorf("Expected 3-element array, got %d elements.",
					len(arr))
			} else {
				for i, v := range arr {
					if f, ok := v.(float64); !ok {
						t.Errorf("Expected floating point number, got %T",
							v)
					} else if f != float64(i+1)/10 {
						t.Errorf("Expected 0.%d, got %f", i+1, f)
					}
				}
			}
		}
	}
}

func TestSubPipe(t *testing.T) {
	t.Parallel()
	testGood(t, `
stage SUM_SQUARES(
    in  float[] values,
    out float   sum,
    src py      "stages/sum_squares",
)

stage REPORT(
    in  float[] values,
    in  float   sum,
    src py      "stages/report",
)

pipeline STUFF(
    in float[] values,
    out float sum,
)
{
    call SUM_SQUARES(
        values = self.values,
    )
    return (
        sum = SUM_SQUARES.sum,
    )
}

pipeline SUM_SQUARE_PIPELINE(
    in  float[] values,
    out float   sum,
    out float   sum2,
)
{
    call STUFF(
        values = self.values,
    )
    call SUM_SQUARES(
        values = self.values,
    )
    call REPORT(
        values = self.values,
        sum    = SUM_SQUARES.sum,
    )

    return (
        sum = STUFF.sum,
        sum2 = SUM_SQUARES.sum,
    )
}

call SUM_SQUARE_PIPELINE(
    values = [1.0, 2.0, 3.0],
	)
`)
}

func TestTopoSort(t *testing.T) {
	t.Parallel()
	if ast := testGood(t, `
stage SUM_SQUARES(
    in  float[] values,
    out float   sum,
    src py      "stages/sum_squares",
)

stage REPORT(
    in  float[] values,
    in  float   sum,
    src py      "stages/report",
)

pipeline SUM_SQUARE_PIPELINE(
    in  float[] values,
    out float   sum,
)
{
    call REPORT(
        values = self.values,
        sum    = SUM_SQUARES.sum,
    )
    call SUM_SQUARES(
        values = self.values,
    )

    return (
        sum = SUM_SQUARES.sum,
    )
}

call SUM_SQUARE_PIPELINE(
    values = [1.0, 2.0, 3.0],
)
`); ast != nil {
		if len(ast.Pipelines) != 1 {
			t.Fatal("Incorrect pipeline count", len(ast.Pipelines))
		} else if calls := ast.Pipelines[0].Calls; len(calls) != 2 {
			t.Fatal("Incorrect call count", len(calls))
		} else if calls[0].Id != "SUM_SQUARES" {
			t.Error("Incorrect stage ordering.")
		}
	}
}

func TestTopoSort2(t *testing.T) {
	t.Parallel()
	if ast := testGood(t, `
stage STAGE (
    in  int input,
    out int output,
    src py  "stages/code",
)

stage STAGE_3 (
    in  int in1,
	in  int in2,
    out int output,
    src py  "stages/code",
)


pipeline PIPELINE(
    in  int input,
    out int output1,
    out int output2,
)
{
    call STAGE_3(
        in1 = self.input,
        in2 = STAGE_2.output,
    )

    call STAGE as STAGE_1(
        input = self.input,
    )

    call STAGE as STAGE_2(
        input = self.input,
    )

    call STAGE as STAGE_4(
        input = STAGE_2.output,
    )

    call STAGE_3 as STAGE_5(
        in1 = self.input,
        in2 = STAGE_2.output,
    )

    return (
        output1 = STAGE_3.output,
        output2 = STAGE_5.output,
    )
}

call PIPELINE(
    input = 1,
)
`); ast != nil {
		if len(ast.Pipelines) != 1 {
			t.Fatal("Incorrect pipeline count", len(ast.Pipelines))
		} else if calls := ast.Pipelines[0].Calls; len(calls) != 5 {
			t.Fatal("Incorrect call count", len(calls))
		} else {
			if calls[0].Id != "STAGE_1" {
				t.Errorf("Incorrect stage ordering: %s was first", calls[0].Id)
			}
			if calls[1].Id != "STAGE_2" {
				t.Errorf("Incorrect stage ordering: %s was second", calls[1].Id)
			}
			if calls[2].Id != "STAGE_3" {
				t.Errorf("Incorrect stage ordering: %s was third", calls[2].Id)
			}
			if calls[3].Id != "STAGE_4" {
				t.Errorf("Incorrect stage ordering: %s was fourth", calls[3].Id)
			}
			if calls[4].Id != "STAGE_5" {
				t.Errorf("Incorrect stage ordering: %s was fifth", calls[4].Id)
			}
		}
	}
}

func TestArrayBind(t *testing.T) {
	t.Parallel()
	testGood(t, `
filetype json;

stage ADD_KEY(
    in  string key,
    in  string value,
    in  json   start,
    out json   result,
    src py     "stages/add_key",
)

stage MERGE_JSON(
    in  json[] inputs,
    out json   result,
    src py     "stages/merge_json",
)

pipeline STUFF(
    in string key1,
    in string value1,
    out json outfile,
)
{
    call ADD_KEY as ADD_KEY1(
        key   = self.key1,
        value = self.value1,
        start = null,
    )
    call ADD_KEY as ADD_KEY2(
        key   = "key2",
        value = "value2",
        start = ADD_KEY1.result,
    )
    call ADD_KEY as ADD_KEY3(
        key   = "key3",
        value = "value3",
        start = ADD_KEY1.result,
    )
    call MERGE_JSON(
        inputs = [
            ADD_KEY2.result,
            ADD_KEY3.result,
        ],
    )
    return (
        outfile = MERGE_JSON.result,
    )
}
`)
}

func TestUserType(t *testing.T) {
	t.Parallel()
	testGood(t, `
filetype goodness;

stage SUM_SQUARES(
    in  float[]  values,
    out goodness sum,
    src py       "stages/sum_squares",
)
`)
}

func TestUserTypeAsFile(t *testing.T) {
	t.Parallel()
	testGood(t, `
filetype goodness;

stage SUM_SQUARES(
    in  float[]  values,
    out goodness sum,
    src py       "stages/sum_squares",
)

pipeline PIPE(
    in  float[] values,
    out file sum,
)
{
    call SUM_SQUARES(
        values = self.values,
    )
    return (
        sum = SUM_SQUARES.sum,
    )
}
`)
}

func TestFileAsUserType(t *testing.T) {
	t.Parallel()
	testGood(t, `
filetype goodness;

stage SUM_SQUARES(
    in  float[] values,
    out file    sum,
    src py      "stages/sum_squares",
)

pipeline PIPE(
    in  float[]  values,
    out goodness sum,
)
{
    call SUM_SQUARES(
        values = self.values,
    )
    return (
        sum = SUM_SQUARES.sum,
    )
}
`)
}

func TestResources(t *testing.T) {
	t.Parallel()
	testGood(t, `
stage SUM_SQUARES(
    in  float[] values,
    out float   sum,
    src py      "stages/sum_squares",
) using (
    threads = 2,
    mem_gb = 1,
)
`)
}

func TestResourcesWithSplit(t *testing.T) {
	t.Parallel()
	if ast := testGood(t, `
stage SUM_SQUARES(
    in  float[] values,
    out float   sum,
    src py      "stages/sum_squares",
) split (
    in  int     foo,
    out int     bar,
) using (
    threads = 2,
    mem_gb = 3,
)
`); ast != nil {
		if len(ast.Stages) != 1 {
			t.Fatalf("Incorrect stage count %d", len(ast.Stages))
		} else if res := ast.Stages[0].Resources; res == nil {
			t.Fatal("No resources.")
		} else {
			if res.StrictVolatile {
				t.Error("Should not be strict volatile.")
			}
			if res.Threads != 2 {
				t.Errorf("Expected 2 threads, saw %d",
					res.Threads)
			}
			if res.MemGB != 3 {
				t.Errorf("Expected 3gb, saw %d",
					res.MemGB)
			}
		}
	}
}

func TestStrictVolatile(t *testing.T) {
	t.Parallel()
	if ast := testGood(t, `
stage SUM_SQUARES(
    in  float[] values,
    out float   sum,
    src py      "stages/sum_squares",
) using (
    threads = 2,
    mem_gb = 1,
    volatile = strict,
)
`); ast != nil {
		if len(ast.Stages) != 1 {
			t.Fatalf("Incorrect stage count %d", len(ast.Stages))
		} else if res := ast.Stages[0].Resources; res == nil {
			t.Fatal("No resources.")
		} else if !res.StrictVolatile {
			t.Error("Not volatile.")
		}
	}
}

func TestRetain(t *testing.T) {
	t.Parallel()
	if ast := testGood(t, `
filetype json;

stage SUM_SQUARES(
    in  float[] values,
    out float   sum,
    out file    report,
    out json    report2,
    src py      "stages/sum_squares",
) using (
    threads = 2,
    mem_gb = 1,
    volatile = strict,
) retain (
    report,
)
`); ast != nil {
		if len(ast.Stages) != 1 {
			t.Fatalf("Incorrect stage count %d", len(ast.Stages))
		}
		if ret := ast.Stages[0].Retain; ret == nil {
			t.Error("No retain.")
		} else if len(ret.Params) != 1 {
			t.Errorf("Expected 1 retain, found %d",
				len(ret.Params))
		} else if ret.Params[0].Id != "report" {
			t.Errorf("Expected to retain 'report'.  Saw '%s' instead.",
				ret.Params[0].Id)
		}
		if outs := ast.Stages[0].OutParams.List; len(outs) != 3 {
			t.Errorf("Incorrect output count: expected 3, got %d", len(outs))
		} else {
			if fname := outs[0].GetOutFilename(); fname != "" {
				t.Errorf("Expected no filename for float output, got %q", fname)
			}
			if fname := outs[1].GetOutFilename(); fname != "report" {
				t.Errorf("Expected file named 'report', got %q", fname)
			}
			if fname := outs[2].GetOutFilename(); fname != "report2.json" {
				t.Errorf("Expected file named 'report2.json', got %q", fname)
			}
		}
	}
}

func TestSplit(t *testing.T) {
	t.Parallel()
	testGood(t, `
stage SUM_SQUARES(
    in  float[] values,
    out float   sum,
    src py      "stages/sum_squares",
) split (
    in  int     foo,
    out int     bar,
)
`)
}

// Check that binding inside an array works.
func TestGoodBindArray(t *testing.T) {
	t.Parallel()
	testGood(t, `
stage SQUARES(
    in  float[][] values,
    out float     square,
    src py        "stages/square",
)

pipeline QUARTIC(
	in  float value,
    out float quart,
)
{
    call SQUARES(
        values = [[1], [2, 3], [1, self.value]],
    )
    return (
        quart = SQUARES.square,
    )
}
`)
}

// Check that null as an array value works.
func TestBindArrayNull(t *testing.T) {
	t.Parallel()
	testGood(t, `
stage SQUARES(
    in  float[][] values,
    out float     square,
    src py        "stages/square",
)

pipeline QUARTIC(
    out float quart,
)
{
    call SQUARES(
        values = null,
    )
    return (
        quart = SQUARES.square,
    )
}
`)
}

func TestMultiArray(t *testing.T) {
	t.Parallel()
	testGood(t, `
stage SQUARES(
    in  int[][] values,
    out float   square,
    src py      "stages/square",
)

pipeline QUARTIC(
    out float quart,
)
{
    call SQUARES(
        values = [[1], [2, 3], [1, 4]],
    )
    return (
        quart = SQUARES.square,
    )
}
`)
}

func TestSplitOut(t *testing.T) {
	t.Parallel()
	testGood(t, `
stage SUM_SQUARES(
    in  float[] values,
    out string  sum,
    src py      "stages/sum_squares",
) split using (
    in  float value,
    out float square,
)
`)
}

func TestVolatile(t *testing.T) {
	t.Parallel()
	if ast := testGood(t, `
stage SQUARE(
    in  int   value,
    out float square,
    src py    "stages/square",
)

pipeline SQ_PIPE(
    out float square,
)
{
    call volatile SQUARE(
        value = 1,
    )
    return (
        square = SQUARE.square,
    )
}
`); ast != nil {
		if mods := ast.Pipelines[0].Calls[0].Modifiers; mods == nil {
			t.Errorf("Nil mods")
		} else {
			if mods.Bindings != nil {
				t.Errorf("Expected non-bound volatile")
			}
			if !mods.Volatile {
				t.Errorf("Expected volatile")
			}
		}
	}
}

func TestPreflight(t *testing.T) {
	t.Parallel()
	if ast := testGood(t, `
stage SQUARE(
    in  int   value,
    src py    "stages/square",
)

pipeline SQ_PIPE(
)
{
    call preflight SQUARE(
        value = 1,
    )
    return ()
}
`); ast != nil {
		if mods := ast.Pipelines[0].Calls[0].Modifiers; mods == nil {
			t.Errorf("Nil mods")
		} else {
			if mods.Bindings != nil {
				t.Errorf("Expected non-bound volatile")
			}
			if mods.Volatile {
				t.Errorf("Expected non volatile")
			}
		}
	}
}

func TestVolatilePreflight(t *testing.T) {
	t.Parallel()
	if ast := testGood(t, `
stage SQUARE(
    in  int   value,
    src py    "stages/square",
)

pipeline SQ_PIPE(
)
{
    call preflight volatile SQUARE(
        value = 1,
    )
    return ()
}
`); ast != nil {
		if mods := ast.Pipelines[0].Calls[0].Modifiers; mods == nil {
			t.Errorf("Nil mods")
		} else {
			if mods.Bindings != nil {
				t.Errorf("Expected non-bound volatile")
			}
			if !mods.Volatile {
				t.Errorf("Expected volatile")
			}
			if mods.Local {
				t.Errorf("Expected non-local")
			}
			if !mods.Preflight {
				t.Errorf("Expected preflight")
			}
		}
	}
}

func TestVolatilePreflight2(t *testing.T) {
	t.Parallel()
	if ast := testGood(t, `
stage SQUARE(
    in  int   value,
    src py    "stages/square",
)

pipeline SQ_PIPE(
)
{
    call SQUARE(
        value = 1,
    ) using (
        volatile  = true,
        preflight = true,
    )
    return ()
}
`); ast != nil {
		if mods := ast.Pipelines[0].Calls[0].Modifiers; mods == nil {
			t.Errorf("Nil mods")
		} else {
			if !mods.Volatile {
				t.Errorf("Expected volatile")
			}
			if mods.Local {
				t.Errorf("Expected non-local")
			}
			if !mods.Preflight {
				t.Errorf("Expected preflight")
			}
		}
	}
}

func TestDisable(t *testing.T) {
	t.Parallel()
	if ast := testGood(t, `
stage SQUARE(
    in  int   value,
    src py    "stages/square",
)

pipeline SQ_PIPE(
    in bool disable_square,
)
{
    call SQUARE(
        value = 1,
    ) using (
        disabled  = self.disable_square,
        volatile  = true,
        preflight = false,
    )
    return ()
}
`); ast != nil {
		if mods := ast.Pipelines[0].Calls[0].Modifiers; mods == nil {
			t.Errorf("Nil mods")
		} else {
			if !mods.Volatile {
				t.Errorf("Expected volatile")
			}
			if mods.Local {
				t.Errorf("Expected non-local")
			}
			if mods.Preflight {
				t.Errorf("Expected non-preflight")
			}
			if mods.Bindings == nil {
				t.Errorf("Expected bindings.")
			} else if dis := mods.Bindings.Table[disabled]; dis == nil {
				t.Errorf("Expected disable binding.")
			}
		}
	}
}

// Tests that preflights accept pipeline input values.
func TestPreflightDepends(t *testing.T) {
	t.Parallel()
	testGood(t, `
stage PREFLIGHT(
    in  int value,
    src py  "stages/preflight",
)

stage COMPUTE(
    in  int value,
    out int result,
    src py  "stages/compute",
)

pipeline THING(
    in  int  value,
    out int  result,
)
{
    call PREFLIGHT(
        value = self.value,
    ) using (
        preflight = true,
    )

    call COMPUTE(
        value = self.value,
    )

    return (
        result = COMPUTE.result,
    )
}

call THING(
    value        = 1,
)
`)
}

func TestTopCall(t *testing.T) {
	t.Parallel()
	testGood(t, `
stage SQUARE(
    in  int   value,
    src py    "stages/square",
)

pipeline SQ_PIPE(
    in int value,
)
{
    call SQUARE(
        value = self.value,
    )
    return ()
}

call SQ_PIPE(
    value = 1,
)
`)
}

// Tests that one can disable a preflight based on a static value.
func TestDisablePreflightGood(t *testing.T) {
	t.Parallel()
	testGood(t, `
stage PREFLIGHT(
    in  int value,
    src py  "stages/preflight",
)

stage COMPUTE(
    in  int value,
    out int result,
    src py  "stages/compute",
)

pipeline THING(
    in  int  value,
    in  bool no_preflight,
    out int  result,
)
{
    call PREFLIGHT(
        value = self.value,
    ) using (
        disabled  = self.no_preflight,
        preflight = true,
    )
    
    call COMPUTE(
        value = self.value,
    )

    return (
        result = COMPUTE.result,
    )
}

call THING(
    no_preflight = true,
    value        = 1,
)
`)
}

func TestCompileBig(t *testing.T) {
	t.Parallel()
	testGood(t, fmtTestSrc)
}

func TestCheckSrcBad(t *testing.T) {
	t.Parallel()
	if ast := testGood(t, `
stage SQUARE(
    in  int   value,
    src py    "☺_a_path_that_¢ertainly_∂oεsn_t_℮xist_☺",
)
`); ast != nil {
		if err := ast.checkSrcPaths([]string{"."}); err == nil {
			t.Error("Expected source check failure.")
		}
	}
}

func TestCheckSrcGood(t *testing.T) {
	t.Parallel()
	if ast := testGood(t, `
stage SQUARE(
    in  int   value,
    src py    "testdata",
)
`); ast != nil {
		if err := ast.checkSrcPaths([]string{"."}); err != nil {
			t.Error(err)
		}
	}
}

func BenchmarkParse(b *testing.B) {
	srcBytes := []byte(fmtTestSrc)
	srcFile := new(SourceFile)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := yaccParse(srcBytes, srcFile, makeStringIntern()); err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkParseAndCompile(b *testing.B) {
	srcBytes := []byte(fmtTestSrc)
	srcFile := new(SourceFile)
	intern := makeStringIntern()
	// prepopulate the string internment cache.
	yaccParse(srcBytes, srcFile, intern)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if ast, err := yaccParse(srcBytes, srcFile, intern); err != nil {
			b.Error(err.Error())
		} else if err := ast.compile(); err != nil {
			b.Error(err.Error())
		}
	}
}
