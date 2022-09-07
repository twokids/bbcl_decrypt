package encrypt

import (
	"bbcl-decrypt/util"
	"crypto/rsa"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"path"
	"sync"
	"time"
)

type User struct {
	ID      int
	Tid     string
	Address string
	Mobile  string
	Name    string
	Remark  string
}

var db *sqlx.DB

// 查询多条数据
func queryList(sql string, pageSize, page int) (error, []User) {
	sql = sql + " limit ?,?"
	//sql := "select id,consignee as name,address,mobile,remark from ks_order_info_v2_copy1 where remark!='ns2291' limit 500,500"
	users := make([]User, 0)

	err := db.Select(&users, sql, page*pageSize, pageSize)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return err, nil
	}
	return nil, users
}

func count(sql string) (error, int) {
	//sql := "select count(*) from ks_order_info_v2_copy1 where remark!='nanshan220901'"
	var ct int
	err := db.Get(&ct, sql)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return err, 0
	}
	return nil, ct
}

func initDB(dsn string) (err error) {
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("connect db failed,err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

func MainEncrypt() {
	//dsn := "test:111111@tcp(127.0.0.1:3306)/test?charset=utf8mb4"
	dsn := "chunjiewang:wcjSa57983@tcp(bbcl-customer-data-services.mysql.polardb.rds.aliyuncs.com:3306)/bbcl_jstdata?charset=utf8mb4"

	initDB(dsn)

	err, ct := count("select count(*) from cust_tm_order_record where remark !='ns2295' and receiver_tel not like '%*%' and receiver_tel!='' ")
	if ct <= 0 || err != nil {
		return
	}
	fmt.Println("待加密数据量：", ct)
	//获取私钥
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("文件创建或打开失败", err)
	}
	filePath1 := path.Join(rootDir, "./etc/public.txt")
	inputPublicStr := util.ReadTxt(filePath1)
	publicKey, _ := util.ParseRsaPublicKeyFromPemStr(inputPublicStr)

	//开启协程
	//读取excel的row。转化解析后的值
	onceCount := 5000
	for i := 0; i <= ct/onceCount; i++ {
		batchEncrypt(onceCount, i, publicKey)
	}
}

func batchEncrypt(pageSize, page int, publicKey *rsa.PublicKey) {
	t1 := time.Now()
	elapsed1 := time.Since(t1)

	err, users := queryList(`select id,tid,receiver_name as name,receiver_address as address,receiver_tel as mobile,remark from cust_tm_order_record where remark !='ns2295' and receiver_tel not like '%*%' order by order_time  `, pageSize, page)
	if err != nil {
		fmt.Println("Query fail")
		return
	}

	// 准备更新语句
	stmt, err := db.Prepare(`update cust_tm_order_record set encrypt_receiver_name=?,encrypt_receiver_tel=?,encrypt_receiver_address=?,remark ='ns229501' where tid=?`)
	if err != nil {
		fmt.Println("Prepare fail")
		return
	}
	defer stmt.Close()
	if len(users) > 0 {
		encryptSave(users, publicKey, stmt)
	}
	elapsed1 = time.Since(t1)
	fmt.Printf("当前执行数据行参 page:%v , pagesize:%v , 执行时长:%v  \n", page, pageSize, elapsed1)
}

func encryptSave(users []User, publicKey *rsa.PublicKey, stmt *sql.Stmt) {
	wg := sync.WaitGroup{}
	ch := make(chan struct{}, 10) // 控制协程数量
	//根据tasklst去执行对应的任务
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("db Begin fail")
		return
	}
	for _, v := range users {
		//if v.Remark!=""{
		//	continue
		//}
		wg.Add(1)
		ch <- struct{}{}
		go func(user User) {
			defer func() {
				<-ch
				wg.Done()
			}()

			tmpRsp := util.Encrypt(publicKey, []string{user.Name, user.Mobile, user.Address})
			_, err := stmt.Exec(tmpRsp[0], tmpRsp[1], tmpRsp[2], user.Tid)
			if err != nil {
				fmt.Println("Exec fail")
				return
			}
		}(v)
	}
	wg.Wait()
	tx.Commit()
	close(ch) //释放ch
}
