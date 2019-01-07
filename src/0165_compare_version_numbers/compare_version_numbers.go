/*
165. Compare Version Numbers
https://leetcode.com/problems/compare-version-numbers/

Compare two version numbers version1 and version2.
If version1 > version2 return 1; if version1 < version2 return -1;otherwise return 0.

You may assume that the version strings are non-empty and contain only digits and the . character.

The . character does not represent a decimal point and is used to separate number sequences.

For instance, 2.5 is not "two and a half" or "half way to version three",
it is the fifth second-level revision of the second first-level revision.

You may assume the default revision number for each level of a version number to be 0.
For example, version number 3.4 has a revision number of 3 and 4 for its first and second level revision number.
Its third and fourth level revision number are both 0.

Note:

Version strings are composed of numeric strings separated by dots . and this numeric strings may have leading zeroes.
Version strings do not start or end with dots, and they will not be two consecutive dots.
*/
// time: 2019-01-07

package compareversionnumbers

import (
	"strconv"
	"strings"
)

// time complexity: O( max n, m )
// space complexity: O(n+m)
func compareVersion(version1 string, version2 string) int {
	var (
		v1 = strings.Split(version1, ".")
		v2 = strings.Split(version2, ".")
	)
	for i := 0; i < len(v1) || i < len(v2); i++ {
		var v1N, v2N int
		if i < len(v1) {
			v1N, _ = strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			v2N, _ = strconv.Atoi(v2[i])
		}

		if v1N > v2N {
			return 1
		} else if v1N < v2N {
			return -1
		}
	}
	return 0
}
