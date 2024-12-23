// カレントディレクトリにビルド結果が生成されない
// 「go env GOCACHE」でキャッシュディレクトリに格納されていることを確認できる
// 「go clean --modcache」で消せる
package hsd

func StringDistance(lhs, rhs string) int {
	return Distrance([]rune(lhs), []rune(rhs))
}

func Distrance(a, b []rune) int {
	dist := 0
	if len(a) != len(b) {
		return -1
	}
	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist
}
