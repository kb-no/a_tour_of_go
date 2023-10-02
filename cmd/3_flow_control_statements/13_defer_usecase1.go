// どういう場面でdeferが使われるか
// ファイルをオープンして編集する場合

func CopyFile(dstName, srcName string) (written int64, err error) {
  src, err := os.Open(srcName) // ←ここでファイルをオープンしたら、
  if err != nil {
		return
  }
  defer src.Close() // ←ここでファイルをクローズする。deferがないとerrがあった場合にファイルが閉じられなくなる。

  dst, err := os.Create(dstName)
  if err != nil {
		return
  }
  defer dst.Close()

  return io.Copy(dst, src)
}
