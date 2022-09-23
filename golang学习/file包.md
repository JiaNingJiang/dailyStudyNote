1.os.Args[0]可以获得当前可执行文件的绝对路径

2.filepath.Dir(os.Args[0]) 可以获得当前可执行文件所在的文件目录(也就是取出路径后的最后一个元素--可执行文件的名称)

3.filepath.Abs()返回传入的文件路径的绝对路径形式

4.string(os.PathSeparator)返回操作系统指定的文件分隔符(windows系统和linux系统不一样)