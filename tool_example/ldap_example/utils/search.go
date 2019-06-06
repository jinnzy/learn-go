package utils

import (
	"fmt"
	"gopkg.in/ldap.v3"
	"log"
)

func ExampleConn_Search() {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "192.168.56.122", 389))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	searchRequest := ldap.NewSearchRequest(
		"dc=devops,dc=com", // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		//"(&(objectClass=organizationalPerson))", // The filter to apply
		//"(&(cn=*))", // The filter to apply
		"(&(cn=goAdmin))", // The filter to apply
		//[]string{"dn", "cn","member"},                    // A list attributes to retrieve
		[]string{},                    // A list attributes to retrieve
		nil,
	)

	bindusername := "cn=admin,dc=devops,dc=com"
	bindpassword := "123456"

	err = l.Bind(bindusername, bindpassword)
	if err != nil {
		log.Fatal("err",err)
	}
	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	result := make(map[string]interface{})
	for _, entry := range sr.Entries {
		//fmt.Println(entry.GetAttributeValue("members"))
		//fmt.Printf("%s: %v\n", entry.DN, entry.GetAttributeValue("cn"))
		//result["DN"] = entry.DN
		//for _,v := range entry.Attributes {
		//	result[v.Name] = v.Values
		//}
		for _,v := range entry.GetAttributeValues("memberUid"){
			result[v] = struct {}{}
		}
	}
	fmt.Println(result)

	if _, ok := result["jinzheyu"]; ok {
		//存在
		fmt.Println(ok)
	}else {
		fmt.Println(ok)
		fmt.Println("不存在")
	}
}
