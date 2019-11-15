package main

import "fmt"

const origin = "https://www.yongin.go.kr"
const listurl = "/user/bbs/BD_selectBbsList.do?q_bbsCode=1001&q_clCode=1&q_currPage=1"

//user/bbs/BD_selectBbsList.do?q_menu=&q_bbsType=&q_clCode=1&_lwprtClCode=&q_searchKeyTy=sj___1002&q_searchVal=&q_bbsCode=1001&q_bbscttSn=&q_currPage=2&q_sortName=&q_sortOrder=&

func main() {
	data, err := getHTML(origin + listurl)
	if err != nil {
		panic(err)
	}
	items, err := parseList(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(items)
}
