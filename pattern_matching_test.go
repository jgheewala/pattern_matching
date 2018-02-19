package pattern_matching

import (
	"testing"
)

func TestWordPatternI290LeetCode(t *testing.T) {
	if IsWordPattern290LeetCode("abba", "dog cat cat dog") != true {
		t.Error("pattern matching for abba & 'dog cat cat dog' failed")
		return
	}

	if IsWordPattern290LeetCode("abba", "dog cat cat fish") != false {
		t.Error("pattern matching for abba & 'dog cat cat fish' failed")
		return
	}
}

func TestWordPatternII291LeetCode(t *testing.T) {
	if IsWordPatternII290LeetCode("abba", "redblueredblue") != false {
		t.Error("pattern matching for abba & redblueredblue failed")
		return
	}
	if IsWordPatternII290LeetCode("abba", "redbluebluered") != true {
		t.Error("pattern matching for abba & redblueredblue failed")
		return
	}
	if IsWordPatternII290LeetCode("aaaa", "asdasdasdasd") != true {
		t.Error("pattern matching for abba & redblueredblue failed")
		return
	}
	if IsWordPatternII290LeetCode("aabb", "xyzabcxzyabc") != false {
		t.Error("pattern matching for abba & redblueredblue failed")
		return
	}
}

func TestIsWildCardMatching(t *testing.T) {
	if IsWildCardMatching("aa", "a") != false {
		t.Error("wildcard matching failed for 'aa' & a")
		return
	}
	if IsWildCardMatching("aa", "aa") != true {
		t.Error("wildcard matching failed for 'aa' & 'aa'")
		return
	}
	if IsWildCardMatching("aaa", "aa") != false {
		t.Error("wildcard matching failed for 'aaa' & 'aa'")
		return
	}
	if IsWildCardMatching("aa", "*") != true {
		t.Error("wildcard matching failed for 'aa' & '*'")
		return
	}
	if IsWildCardMatching("aa", "a*") != true {
		t.Error("wildcard matching failed for 'aa' & 'a*'")
		return
	}
	if IsWildCardMatching("aa", "?a") != true {
		t.Error("wildcard matching failed for 'aa' & '?a'")
		return
	}
	if IsWildCardMatching("aab", "c*a*b") != false {
		t.Error("wildcard matching failed for 'aab' & 'c*a*b'")
		return
	}
	if IsWildCardMatching("abc", "a******b*********c") != true {
		t.Error("wildcard matching failed for 'abc' & 'a******b*********c'")
		return
	}
}

func TestIsRegularExpressionMatch(t *testing.T) {
	if IsRegularExpressionMatch("aa", "a") == true {
		t.Error("IsRegularExpressionMatching aa & a failed")
		return
	}
	//fmt.Println("aa & a done")

	if IsRegularExpressionMatch("aa", "aa") == false {
		t.Error("IsRegularExpressionMatching aa & aa failed")
		return
	}
	//fmt.Println("aa & aa done")

	if IsRegularExpressionMatch("aaa", "aa") == true {
		t.Error("IsRegularExpressionMatching aaa & aa failed")
		return
	}
	//fmt.Println("aaa & aa done")

	if IsRegularExpressionMatch("aa", "a*") == false {
		t.Error("IsRegularExpressionMatching aa & a* failed")
		return
	}
	//fmt.Println("aa & a* done")

	if IsRegularExpressionMatch("aa", ".*") == false {
		t.Error("IsRegularExpressionMatching aa & .* failed")
		return
	}
	//fmt.Println("aa & .* done")

	if IsRegularExpressionMatch("ab", ".*") == false {
		t.Error("IsRegularExpressionMatching ab & .* failed")
		return
	}
	//fmt.Println("ab & .* done")

	if IsRegularExpressionMatch("aab", "c*a*b") == false {
		t.Error("IsRegularExpressionMatching aab & c*a*b failed")
		return
	}
	//fmt.Println("aab & c*a*b* done")

	if IsRegularExpressionMatch("aaa", "a.a") == false {
		t.Error("IsRegularExpressionMatching aaa & a.a failed")
		return
	}
	//fmt.Println("aaa & a.a done")
}
