package constant

type ColumnType string

const (
	CTChar       ColumnType = "char"       //0-255 bytes    		定长字符串
	CTVarchar    ColumnType = "varchar"    //0-65535 bytes    		变长字符串
	CTTinyblob   ColumnType = "tinyblob"   //0-255 bytes    		不超过 255 个字符的二进制字符串
	CTTinytext   ColumnType = "tinytext"   //0-255 bytes    		短文本字符串
	CTBlob       ColumnType = "blob"       //0-65 535 bytes    		二进制形式的长文本数据
	CTText       ColumnType = "text"       //0-65 535 bytes    		长文本数据
	CTMediumblob ColumnType = "mediumblob" //0-16 777 215 bytes     二进制形式的中等长度文本数据
	CTMediumtext ColumnType = "mediumtext" //0-16 777 215 bytes     中等长度文本数据
	CTLongblob   ColumnType = "longblob"   //0-4 294 967 295 bytes  二进制形式的极大文本数据
	CTLongtext   ColumnType = "longtext"   //0-4 294 967 295 bytes  极大文本数据
	CTTinyint    ColumnType = "tinyint"    //1 byte	                小整数值
	CTSmallint   ColumnType = "smallint"   //2 bytes                大整数值
	CTMediumint  ColumnType = "mediumint"  //3 bytes                大整数值
	CTInt        ColumnType = "int"        //4 bytes                大整数值
	CTInteger    ColumnType = "integer"    //4 bytes                大整数值
	CTDecimal    ColumnType = "decimal"
	CTBigint     ColumnType = "bigint"    //8 bytes                极大整数值
	CTFloat      ColumnType = "float"     //4 bytes                单精度 浮点数值
	CTDouble     ColumnType = "double"    //8 bytes                双精度 浮点数值
	CTDate       ColumnType = "date"      //YYYY-MM-DD	            日期值
	CTTime       ColumnType = "time"      //HH:MM:SS	            时间值或持续时间
	CTYear       ColumnType = "year"      //YYYY	 				年份值
	CTDatetime   ColumnType = "datetime"  //YYYY-MM-DD HH:MM:SS	混合日期和时间值
	CTTimestamp  ColumnType = "timestamp" //YYYYMMDD HHMMSS		混合日期和时间值，时间戳
)
