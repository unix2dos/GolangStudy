## database/sql 接口
		1. sql.Register 注册驱动
		Register(name string, driver driver.Driver)

		func init() {
			sql.Register("sqlite3", &SQLiteDriver{}
		)

		2. driver.Driver
		type Driver interface {
			Open(name string) (Conn, error)
		}

		3. driver.Conn
		Conn只能应用在一个goroutine里面，不能使用在多个goroutine里面
		type Conn interface {
			Prepare(query string) (Stmt, error)//执行语句准备
			Close() error		//关闭连接
			Begin() (Tx, error) //事务处理
		}

		4. driver.Stmt  
		stmt是一种准备好的状态，和Conn相关联，而且只能应用于一个goroutine中，不能应用于多个goroutine。
		type Stmt interface {
			Close() error		//关闭
			NumInput() int		//返回预留参数的个数
			Exec(args []Value) (Result, error) //update insert
			Query(args []Value) (Rows, error)  //select
		}

		5. driver.Tx
		事务处理
		type Tx interface {
			Commit() error
			Rollback() error
		}

		6. driver.Execer
		这是一个Conn可选择实现的接口, 如果这个接口没有定义，那么在调用DB.Exec,就会首先调用Prepare返回Stmt，然后执行Stmt的Exec，然后关闭Stmt。
		type Execer interface {
			Exec(query string, args []Value) (Result, error)
		}

		7. driver.Result
		返回update, insert结果
		type Result interface {
			LastInsertId() (int64, error)
			RowsAffected() (int64, error)
		}
		LastInsertId函数返回由数据库执行插入操作得到的自增ID号。
		RowsAffected函数返回query操作影响的数据条目数。

		8. driver.Rows
		返回select 结果
		type Rows interface {
			Columns() []string //返回查询数据库表的字段信息，这个返回的slice和sql查询的字段一一对应，而不是返回整个表的所有字段
			Close() error  //关闭rows
			Next(dest []Value) error  //返回下一条数据，把数据赋值给dest。dest里面的元素必须是driver.Value的值除了string，返回的数据里面所有的string都必须要转换成[]byte。如果最后没数据了，Next函数最后返回io.EOF。
		}

		9. driver.RowsAffected
		RowsAffected其实就是一个int64的别名，但是他实现了Result接口，用来底层实现Result的表示方式
		type RowsAffected int64
		func (RowsAffected) LastInsertId() (int64, error)
		func (v RowsAffected) RowsAffected() (int64, error)

		10. driver.Value
		Value其实就是一个空接口，他可以容纳任何的数据
		type Value interface{}

		drive的Value是驱动必须能够操作的Value，Value要么是nil，要么是下面的任意一种
		int64
		float64
		bool
		[]byte
		string   [*]除了Rows.Next返回的不能是string, 如果是string要转换成[]byte
		time.Time

		11. driver.ValueConverter
		ValueConverter接口定义了如何把一个普通的值转化成driver.Value的接口
		type ValueConverter interface {
			ConvertValue(v interface{}) (Value, error)
		}

		12. driver.Valuer
		Valuer接口定义了返回一个driver.Value的方式, 很多类型都实现了这个Value方法，用来自身与driver.Value的转化。
		type Valuer interface {
			Value() (Value, error)
		}


