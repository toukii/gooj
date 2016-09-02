package utils

import (
	"testing"
)

var (
	fail = `=== RUN   TestOj
TestCase:[]int{3, 5, 1}, RunResult:[]int{3, 5, 1}
ResultWanted:[]int{3, 3, 1}
--- FAIL: TestOj (0.00s)
        reverse_test.go:17: FAILED
FAIL
exit status 1
FAIL    github.com/toukii/gooj/goojt    0.243s`

	fail_pass = `=== RUN   TestOj
=== RUN   TestOj
cost: 106.179µs
--- PASS: TestOj (0.00s)
PASS
ok  	_/gopath/goojle/56782	0.003s
=== RUN   TestOj
cost: 106.179µs
--- PASS: TestOj (0.00s)
PASS
ok  	_/gopath/goojle/56782	0.003s

TestCase: [2 1 3], wanted:[1 2 3]
--- FAIL: TestOj (0.00s)
	sort_test.go:17: Got:[2 1 3]
FAIL
exit status 1
FAIL	_/gopath/goojle/20417	0.003s`

	setup_fail = `# github.com/toukii/gooj/goojle/51102
percentString_test.go:51:40: missing ',' before newline in composite literal
FAIL	github.com/toukii/gooj/goojle/51102 [setup failed]`

	build_fail = `# _/gopath/goojle/20417
./sort.go:10: undefined: l
FAIL	_/gopath/goojle/20417 [build failed]`

	pass = `=== RUN   TestOj
cost: 106.179µs
--- PASS: TestOj (0.00s)
PASS
ok  	_/gopath/goojle/56782	0.003s`

	pass_fail = `=== RUN   TestOj
fff=== RUN   TestOj
fffcost: 106.179µs
fff--- PASS: TestOj (0.00s)
fffPASS
fffok  	_/gopath/goojle/56782	0.003s
fff=== RUN   TestOj
fffcost: 106.179µs
fff--- PASS: TestOj (0.00s)
fffPASS
fffok  	_/gopath/goojle/56782	0.003s
cost: 211.759µs
--- PASS: TestOj (0.00s)
PASS
ok  	_/gopath/goojle/23327	0.003s`
)

func TestAnalyse(t *testing.T) {
	Analyse(setup_fail)
	Analyse(build_fail)
	Analyse(fail)
	Analyse(fail_pass)
	Analyse(pass)
	Analyse(pass_fail)
}
