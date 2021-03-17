package classpath

import "os"
import "strings"

const pathListSeparator = string(os.PathListSeparator)

//Entry-接口类型名，接口名一般er结尾
//Entry接口有4个实现，分别是DirEntry、ZipEntry、 CompositeEntry和WildcardEntry
type Entry interface {
	//方法名首字母应该大写，这里没大写
	//class文件的相对路径
	//传 入的参数应该是java/lang/Object.class
	readClass(className string) ([]byte, Entry, error)
	String() string
}

//根据参数创建不同类型的Entry实例
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)

}
