package classpath

import "io/ioutil"
import "path/filepath"

//用于存放目录的绝对路径，和Java语言 不同，Go结构体不需要显示实现接口，只要方法匹配即可。
type DirEntry struct {
	absDir string
}

//Go没有 专门的构造函数，本书统一使用new开头的函数来创建结构体实 例，并把这类函数称为构造函数
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	//使用& 视为对该实例进行new对实例操作
	return &DirEntry{absDir}
}
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}
func (self *DirEntry) String() string {
	return self.absDir
}
