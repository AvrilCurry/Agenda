# Agenda
A CLI Program

Agenda程序可以分为三部分来看，第一部分就是命令行的设计，第二部分就是后台数据的存储，第三部分就是
命令行读写后台数据。命令行的设计就是包括在cmd文件夹里面的所有.go文件，每个文件都是各自命令
行的设计，同时也包括了在调用该命令后的消息处理。第二部分就是后台数据的存储，这里数据都是以Json的
格式存储在文件中，分别是userDatabase.data和meetingDatabase.data文件中。第三部分就是命令行读写
后台数据。这部分是通过在entity文件夹里面的data.go文件实现的，data.go里面包含了cmd文件中所有（除了help）
.go文件的处理方式，而命令行也是通过这些函数接口来实现对后台数据的读写。

在TestCobra文件夹里面还包含了log文件夹，里面的error.log显示命令行运行的信息以及结果。

在Agenda文件夹里面的operation.txt是我测试时调用的命令行的顺序。

命令行的设计是在Agenda文件夹里面的design.md。

User和Meeting两个数据结构是在Agenda文件夹里面的struct.md。

最后结果的截图是放在Screenshot文件夹里面，分别对应相应的命令。
