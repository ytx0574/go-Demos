package main

import (
	"C"
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
#cgo CFLAGS: -x -objective-c
*/

/*
file
*/
func main() {
	/*
	此处打开文本, 因为对\r \n的处理, 所有很多内容是看不到的.  需要注意分别
	"\r\raaa[\r\r]bbb\r\r"
	*/

	strFilePath := "/Users/johnson/Desktop/go_aaaa1.txt"
	file, err := os.Open(strFilePath)  //打开一个文件, 默认里面是只读
	defer file.Close()
	if err != nil {
		fmt.Printf("打开文件:%q出错, err:%v\n", strFilePath, err)
	}
	fmt.Printf("file = %v\n", file)

	//strFilePath = "/Users/johnson/Desktop/go_aaaa122.txt"
	//file, err = os.Open(strFilePath)
	//if err != nil {
	//	fmt.Printf("打开文件:%q出错, err:%v\n", strFilePath, err)
	//}
	//fmt.Printf("file = %v\n", file)

	////一次性读文件
	//strBytes, err := ioutil.ReadFile(strFilePath)   //内部实现为ReadAll
	//if err != nil {
	//	fmt.Printf("一次性读取文件失败:%v\n", err)
	//}
	//fmt.Printf("一次性获取文件内容:\nbytes:%v\nstring:%v\n", strBytes, string(strBytes))
	//fmt.Println("\n")

	//分批次读取 size指定缓冲区的byte数
	var reader *bufio.Reader = bufio.NewReaderSize(file, 18)
	//for {
	//	str, err := reader.ReadString('6')
	//	fmt.Printf("str:%v, err:%v\n", str, err)
	//	if err == io.EOF {
	//		break
	//	}
	//}


	//按切片长度读取, 如果长度大于等于缓冲区, 那么直接返回p, 如果长度小于缓冲区, 把数据写入缓冲区, 然后再从缓冲区拷贝到切片
	strBytes := make([]byte, 9)
	n, err := reader.Read(strBytes) //指定Read长度,
	fmt.Printf("reader.Read strBytes:%v, strBytesLen:%v\n", string(strBytes), len(strBytes))
	fmt.Printf("reader.Read:%v, err:%v\n", n, err)

	//读取一次读取所有内容
	//strBytesAll, err := ioutil.ReadAll(reader)
	//fmt.Printf("ioutil.ReadAll:%v, err:%v\n,", string(strBytesAll), err)

	//按字节读取 返回一个byte
	Byte, err := reader.ReadByte()
	fmt.Printf("ReadByte:%v, err:%v\n", Byte, err)

	//读取到指定字节, 返回[]byte
	strBytes1, err := reader.ReadBytes('m')
	fmt.Printf("ReadBytes:%v\n", string(strBytes1))

	//整行读取, 如果超过buffer长度, 则会停止. 无法读取该行后面的数据
	line, prefix, err := reader.ReadLine()
	fmt.Printf("ReadLine: %v, prefix:%t, err:%v\n", string(line), prefix, err)

	//按int32读取. 可读取整个汉字. (按字读取, 不是按字节读)
	rune, size, err :=  reader.ReadRune()
	fmt.Printf("ReadRune:%v size:%v, err:%v\n", string(rune), size, err)
	rune, size, err =  reader.ReadRune()
	fmt.Printf("ReadRune:%v size:%v, err:%v\n", string(rune), size, err)	//按int32读取. 可读取整个汉字
	rune, size, err =  reader.ReadRune()
	fmt.Printf("ReadRune:%v size:%v, err:%v\n", string(rune), size, err)

	//按字节读取, 直接第一次遇到指定字符.  如果超过buffer长度, 会停止.
	strBytes2, err := reader.ReadSlice('i')
	fmt.Printf("ReadSlice:%v, err:%v\n", string(strBytes2), err)

	bufferd := reader.Buffered()
	fmt.Printf("Buffered:%v\n", bufferd)

	//重置数据源, 重置reader. 充值reader除buf之外的其他参数
	fmt.Printf("reader:%v\n", reader)
	reader.Reset(file)
	fmt.Printf("reader:%v\n", reader)

	str, err := reader.ReadString('8')
	fmt.Printf("ReadString:%v\n", str)

	line, prefix, err = reader.ReadLine()
	fmt.Printf("ReadLine: %v, prefix:%t, err:%v\n", string(line), prefix, err)



	/*
	一般都文件属性标识如下： -rwxrwxrwx
	Unix/Linux的文件权限. 权限数值为八进制. 所以添加权限时, 必须前面加0

	第1位：文件属性，一般常用的是"-"，表示是普通文件；"d"表示是一个目录。
	第2～4位：文件所有者的权限rwx (可读/可写/可执行)。
	第5～7位：文件所属用户组的权限rwx (可读/可写/可执行)。
	第8～10位：其他人的权限rwx (可读/可写/可执行)。
	*/
	fmt.Printf("os.ModePerm:%v\n", os.ModePerm)
	fmt.Printf("os.0666 所有人都可读可写:%v\n", os.FileMode(0666))
	fmt.Printf("os.0777 所有人可读可写可执行:%v\n", os.FileMode(0777))
	fmt.Printf("os.0644 文件所有者可读写, 用户组和其他人只可读:%v\n", os.FileMode(0644))
	fmt.Printf("os.0744 文件所有者可读写执行, 用户组和其他人只可读:%v\n", os.FileMode(0744))

	//写文件  文件所有者权限不能乱写. 否则Unix/Linux下面无法查看
	file1, err := os.OpenFile(strFilePath, os.O_RDWR | os.O_APPEND | os.O_CREATE, os.ModePerm)
	defer file1.Close()

	var writer *bufio.Writer = bufio.NewWriterSize(file1, 21)
	var writeStr = `
用t保存全局的itabTable地址，然后使用t.find函数查找，这么做是为了防止在查找
	过程中itabTable被替换导致错误
	如果未找到，再尝试加锁查找。原因是第一步查找时可能有另一个协程并发写入，从而导致Find函数未找到但实际数据是
	存在的。这时通过加锁防止itabTable被写入，然后在itabTable中查找
	如果扔为找到，此时根据接口类型和数据类型生成一个新的itab插入itabTable中。如果插入失败，则panic 注意这里添加时，申请的内存大
	小为len(inter.mhdr)-1，前面我们知道fun数组大小为1，所以这里再申请内存时只需再申请len(inter.mhdr)-1即可。`
	//writeStr = "12345678900987"
	writebytenum, err := writer.WriteString(writeStr)
	fmt.Printf("WriteString Write Byte Num:%d, err:%v\n", writebytenum, err)
	//将缓存中得数据写入到文件. 因为默认写入带缓存, 如果缓存没有满的时候, 没有满的那一部分不会直接写入到文件的.
	//如果不调用Flush(), 写入缓存的时候, 中文字符因为占用三个字节, 有时缓存大小不足的情况下, 会导致写入乱码.
	//比如这里, 如果上面得size为1024, 那么写入这句话时, 缓存未满, 有一部分不会写入进去
	writer.Flush()

	//写入byte数组 此处同样, 如果写入数据小于缓存大小, 那么不实用Flush无法写入数据
	//p := []byte{'\n', '1', '2', '\n'}
	p := make([]byte, writer.Size() - 10)
	p[0] = '\n'
	p[len(p) - 1] = '\n'
	writer.Write(p)
	//writer.Flush()

	//写入一个字符
	writer.WriteByte('Y')
	//写入一个rune. 包含3个字节的中文
	writer.WriteRune('操')
	writer.Flush()

	//此处注意, Write不会报错. 但是下面的Flush会报错.  因为file的descriptor是一个只读File对象.
	//所以判断是否写入成功, 应该通过flush来判断
	writer1 := bufio.NewWriter(file)
	writebytenum, err = writer1.WriteString("111")
	flushErr := writer1.Flush()

	if flushErr != nil {
		 fmt.Printf("写入失败 flushErr:%v\n", flushErr)
	}else {
		fmt.Printf("写入成功 Write Byte Num:%d\n", writebytenum)
	}

	//判断文件是否存在
	 fileInfo, err := os.Stat(strFilePath)
	 fmt.Printf("fileInfo:%v\n", fileInfo)
	if err == nil {
		fmt.Printf("文件存在, Name:%v Size:%v Mode:%v Sys:%v\n", fileInfo.Name(), fileInfo.Size(), fileInfo.Mode(), fileInfo.Sys())
	}
	if os.IsNotExist(err) {// os.IsExist
		fmt.Printf("文件或文件夹不存在\n")
	}

	//拷贝文件  //文件路径不可包含空格.
	file2, err := os.Open("/Users/johnson/Desktop/go_aaaa2.txt")
	reader1 := bufio.NewReader(file2)
	//此处死循环. 因为writer和reader都指向同一个文件. 一直拷贝同一个文件. 会出错
	//written, err := io.Copy(writer, reader)
	written, err := io.Copy(writer, reader1)
	if err != nil {
		fmt.Printf("拷贝文件失败 err:%v\n", err)
	}else {
		fmt.Printf("拷贝文件成功. 拷贝长度为:%v\n", written)
	}



	////简单执行命令行, 输出到文件
	//file3, err := os.OpenFile("/Users/johnson/Desktop/go_aaaa3.txt", os.O_CREATE | os.O_RDWR, 0777)
	//var pipeReader, pipWriter = io.Pipe()
	//fmt.Printf("%v%v\n", pipeReader, pipWriter)
	//var cmd *exec.Cmd = exec.Command("printenv")
	////cmd.Stdout = os.Stdout
	//cmd.Stdout = file3
	//cmd.Run()

	//io.Pipe() //的简单使用. 返回一对PipeReader PipeWriter
	//reader, writer := io.Pipe()
	//defer writer.Close()
	//lock := make(chan int)
	//// 创建goroutine给reader
	//go func() {
	//	buffer := make([]byte, 100)
	//	reader.Read(buffer)
	//	println(string(buffer))
	//	lock <- 1
	//}()
	//writer.Write([]byte("hello"))
	//<-lock

	//io.TeeReader 将读取的内容写入的一个writer, 并返回一个新的reader
	buffer := make([]byte, 30)
	reader2 := io.TeeReader(reader, writer)
	reader2.Read(buffer)
	fmt.Printf("TeeReader %v\n", string(buffer))
	if reader == reader2 {
		fmt.Printf("TeeReader 直接返回之前的reader\n")
	}else {
		fmt.Printf("TeeReader 返回新的的reader\n")
	}


	//将多个writer 串联起来, 并返回一个新的writer.  返回一个新的multiWriter. 会将写入的数据写入到每=提供的每一个writer
	//io.MultiWriter()

	//将多个reader串联起来. 读取时, 依次获取每个reader的数据 当所有的reader都读取完成, 才会返回io.EOF
	//io.MultiReader()

	createFileWithSize(1024 * 10)
}

//创建指定MB的文件 创建垃圾文件
func createFileWithSize(size int) {
	file, _ := os.OpenFile("/Users/Johnson/Desktop/garbage_file", os.O_RDWR | os.O_APPEND | os.O_CREATE, os.ModePerm)
	defer file.Close()

	str := ""
	for i := 0; i < 1024; i++ {
		str += "\b"
	}

	strBytes := []byte(str)
	for i := 0; i < size * 1024; i++ {
		file.Write(strBytes)
	}
}