package main

import (
	"fmt"
	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func main()  {
	// New一个gorm adapter
	a,err := gormadapter.NewAdapter("mysql","root:123456@tcp(192.168.56.122:3306)/test_casbin",true)
	if err != nil {
		log.Fatal(err)
	}
	// 加载rbac.conf和gorm adapter
	enforcer,err := casbin.NewEnforcer("E:/gocode/src/github.com/learn-go/tool_example/casbin_test/rbac.conf",a)
	if err != nil {
		log.Fatal(err)
	}
	// 添加策略
	added,err := enforcer.AddPolicy("2", "/api/v1/users", ".*")
	if err != nil {
		log.Fatal(err)
	}
	if added {
		log.Println("添加成功")
	}
	// 添加组策略
	enforcer.AddGroupingPolicy("test", "2")
	//err = enforcer.LoadPolicy()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//succ,err := enforcer.Enforce("superManager","/api/v1/users","get")
	//succ,err := enforcer.Enforce("alice","/api/v1/users","get")
	// 添加组策略，alice继承superManager的权限
	succ,err := enforcer.Enforce("test","/api/v1/users","get")
	if err != nil {
		log.Println(err)
	}
	if succ {
		log.Println("验证通过")
	}else {
		log.Println("验证未通过")

	}
	//fmt.Println(enforcer.GetGroupingPolicy())
	// 添加角色为superManagerRole
	//added2,err := enforcer.AddPolicy("superManagerRole", "/api/v/users", "*","test")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if added2 {
	//	log.Println("添加成功")
	//}

	fmt.Println(enforcer.GetUsersForRole("supertManager"))
	//data := enforcer.GetAllRoles()
	//fmt.Println(data)
}
