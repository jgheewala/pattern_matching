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

func TestIsMatch(t *testing.T) {
	if IsMatch("aa", "a") == true {
		t.Error("IsMatching aa & a failed")
		return
	}
	//fmt.Println("aa & a done")

	if IsMatch("aa", "aa") == false {
		t.Error("IsMatching aa & aa failed")
		return
	}
	//fmt.Println("aa & aa done")

	if IsMatch("aaa", "aa") == true {
		t.Error("IsMatching aaa & aa failed")
		return
	}
	//fmt.Println("aaa & aa done")

	if IsMatch("aa", "a*") == false {
		t.Error("IsMatching aa & a* failed")
		return
	}
	//fmt.Println("aa & a* done")

	if IsMatch("aa", ".*") == false {
		t.Error("IsMatching aa & .* failed")
		return
	}
	//fmt.Println("aa & .* done")

	if IsMatch("ab", ".*") == false {
		t.Error("IsMatching ab & .* failed")
		return
	}
	//fmt.Println("ab & .* done")

	if IsMatch("aab", "c*a*b") == false {
		t.Error("IsMatching aab & c*a*b failed")
		return
	}
	//fmt.Println("aab & c*a*b* done")

	if IsMatch("aaa", "a.a") == false {
		t.Error("IsMatching aaa & a.a failed")
		return
	}
	//fmt.Println("aaa & a.a done")
}
