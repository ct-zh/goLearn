package _78

var badVersion int

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	l := 1
	r := n
	for l < r {
		mid := (r-l)/2 + l
		if isBadVersion(mid) { // 该版本是坏版本
			r = mid
		} else { // 该版本是正常的版本
			l = mid + 1
		}
	}
	return l
}

func isBadVersion(version int) bool {
	return version >= badVersion
}
