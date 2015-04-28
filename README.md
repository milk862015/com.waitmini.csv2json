#使用说明
##csv数据填写说明
csv的
第一行是描述，用于方便知道该列数据的意义。<br>
第二行是导出的json的关键key的名字。<br>
第三行是导出的该列的数据类型。分别有string,number,array-string,array-number。<br>
具体可参考项目中的round.csv<br>
然后运行如下命令<br>
csv2json [csv文件名] [导出保存的数据的文件名]<br>
例如<br>csv2json round.csv round.json<br>