package advancedfunctions

/*
Defer - is a unique feature of go.
it allows a function to execute automatically just before its enclosing function returns.
deferred functions are used to execute the close of db connections, file closes etc.

example:
-------
func CopyFile(dstName, srcName string) {
src, err := os.Open(srcName)
if err != nil {
	return
}

defer src.Close()

dst, err := os.Open(dstName)
if err != nil {
	return
}
defer dst.Close()

}
*/
