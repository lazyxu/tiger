b print # 下断点
r # 运行
disassemble # 显示汇编代码
set disassembly-flavor intel # 设置为intel风格汇编

x/nfu 0×300098  显示指定地址的内存数据
n 显示内存单位，长度
f 格式(除了 print 格式外，还有 字符串s 和 汇编 i)
u 内存单位(b: 1字节; h: 2字节; w: 4字节; g: 8字节)

x/10x $sp 打印stack的前10个元素
