package main

import (
	"bytes"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	//"github.com/google/uuid"
	"path/filepath"
	"strings"
	"sync"

	//"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)
type Ret struct{
	Code int
	Param string
	Msg string
	Data []Note
}

var Notes []Note

var Type = [6]string{"note","word","pdf","ppt","excel"}
var Color = [...]string{"66CCCC","FF99CC","FF9999","FFCC99","FF6666","99CC66","666699","FF9999","99CC33","FF9900","FFCC00","FF0033","FF9966","FF9900","CC3399","99CC33","FF6600","993366","CCCC33","666633","66CCCC","666699","CC9999","666699","FF9900","0099CC","CCCC99","CC3399","99CC00","FF6666","3399CC","CC6600","999999","CCCC33","FF9933","009933","0099CC","CCCCCC","FF6666","FF6600","009966","CC6633","FFCC99","CC6600","CC0066","009999","FFCC33"}
var Title = [...]string{"进入大厂的面试经验（P7）","吃透 Vue 项目开发实践｜16个方面深入前端工程化开发技巧《下》","大厂面试题分享：如何让(a===1&&a===2&&a===3)的值为true?","2020 年，Vue 受欢迎程度是否会超过 React？","前端高级进阶：前端部署的发展历程","前端早早聊|堂主 - 如何推动前端团队的基础设施建设","（4.5w字🔥建议收藏）“可能是” 2020最值得看的小程序系列实战教程","使用vue实现HTML页面生成图片","2019年度最常见的JavaScript面试题和答案","前端高级进阶：使用 docker 高效部署你的前端应用","字节跳动今日头条前端面经（4轮技术面+hr面）","Java堆内存是线程共享的！面试官：你确定吗？","CSS中容易被忽视的 position属性sticky","简单通俗的理解Vue3.0中的Proxy","做了5年iOS，靠着这份面试题跟答案，我从12K变成了30K","127个常用的JS代码片段，每段代码花30秒就能看懂（四）","60行代码，造一个动画库轮子（超详细）","一个不容错过的Spring Cloud实战项目！","SpringBoot中处理校验逻辑的两种方式，真的很机智！","面试官问：Node 与底层之间如何执行异步 I/O 调用","深入探索Android稳定性优化","2017-2020历年字节跳动Android面试真题解析","RSA初探，聊聊怎么破解HTTPS","使用uni-app开发叮咚买菜的一些笔记","React Hooks完全上手指南","写个日志请求切面，前后端甩锅更方便","基于Redis实现分布式锁之前，这些坑你一定得知道","2020 校招，我是如何拿到小米、京东、字节大厂前端offer","JavaScript实现图片合成下载","async原理解析","Fuchsia 的开发语言政策与解读","学习 koa 源码的整体架构，浅析koa洋葱模型原理和co原理","Insert into select语句引发的生产事故","百万级商品数据实时同步，查询结果秒出","来聊聊Activity的显示原理","Java 8 Optional 良心指南，建议收藏","【从零冲击音视频开发】移动环境搭建","javascript关键的十大优化知识点【建议大忙人的你看看】","2020互联网Java后端面试必备解析—Redis23题","记一次vue-element-admin 的动态路由权限管理和菜单渲染的学习","🔥🔥🔥面向 Android 高级工程师的一份面试宝典 (持续更新)🔥🔥🔥","装饰器(Decorator)","近2万字详解JAVA NIO2文件操作，过瘾！","这样学习Servlet，会事半功倍！！","想前进的前端进行的自我反省","从零到部署：用 Vue 和 Express 实现迷你全栈电商应用（七）","一个成功的程序员，自然要懂微服务，汇总微服务架构的15钟框架！","SpringBoot图文教程14—阿里开源EasyExcel「为百万数据读写设计」","炫酷！从未见过如此Q弹的Switcher","基于Java实现的人脸识别功能，一切都为了宠粉（附源码）","我与花呗与前端技术","NEI 开源预告","从前端现状到数据校验史","Go语言基础（五）—— 并发编程","前端劝退预警：JavaScript 工具链不完全指南","副业收入是我做程序媛的3倍，工作外的B面人生是怎样的？","缓存穿透了怎么办？","'&'和'&&' 往年没考,今年肯定要考","🔥🔥 用vue全家桶实现mac版微信（不断更新...）","微服务落地实践 - 经验分享","Ant Design核心作者谈4.0版本更新背后的故事","详解ES6中的class","webpack优化实践(遇到很奇葩的问题)","小册上新 | React SSR 服务端渲染原理解析与实践","编写高质量JavaScript模块的4个最佳实践","【小技巧】巧用CSS属性值正则匹配选择器","面试题及日常小结（一）-基础篇","axios发起http请求的一些细节","全面分析总结JS内存模型","面试题：Redis的应用场景核心设计，看完面试不在慌！","前端工程师的 LeetCode 之旅 -- 174周赛","RedisTemplate：我不背锅，是你用错了","前端早早聊|Scott - 如何在人单力薄时立项推动基建","面试被问分布式事务（2PC、3PC、TCC），这样解释没毛病！","VUE 3.0 学习探索入门系列 - 用几个 demo 认识 vue3（3）","Now in Android：02 - 欢迎使用 Android Studio 4.0 ！","开源字体不香吗？五款 GitHub 上的爆红字体任君选","一个iOS老兵说，疫情当前，风险和机遇并存！","Flutter 拖拽控件Draggable看这一篇就够了","简单，真诚「阿里新零售CRO技术部大前端团队」春招实习！","30分钟带你了解Web工程师必知的Docker知识","Swift开发小记（含面试题）","前端高级进阶：CICD 下前端的多特性分支环境部署","可能是最好的 BFC 解析了... ...","字节跳动面经 前端实习 (1+2+3+hr)","【背上Jetpack之Fragment】从源码角度看 Fragment 生命周期 AndroidX Fragment1.2.2源码分析","Now in Android #13 - 最新 Android 动态分享","NutUI 3.0 中单元测试的探索和实践","[深入17] HTTP 和 HTTPS","Android性能优化-你的lottie动画今天跳帧了吗？","使用Electron实现一个iPic","前端工程师的 LeetCode 之旅 - 夜喵专场（21）","双非本科！2020年Java应届生秋招面试回顾！","JavaScript 中的内存管理是如何工作的","Android事件分发底层原理","【Java面试题】关于String，最近被问到了这2道面试题","看过无数Java GC文章，这5个问题你也未必知道！","Flutter、iOS混合开发实践","flutter实现B站播放器暂停时的header效果","使用视图绑定替代 findViewById","想要拿到大厂offer之前，先了解一下这些吧！","Spring Security 实战干货：过滤器链的机制和特性","仿抖音上下滑动播放视频","微信小程序生成图片分享朋友圈","这几道JS面试刁钻题，你能答对吗😎","2020面试，已拿到头条跟阿里offer","这样谈gc，面试官将被你吊打","做一个能在线编程+视频对话的视频面试应用","基于Canal和Kafka实现MySQL的Binlog近实时同步","「蚂蚁金服 AntV」G6 3.4 又双叒叕来啦","现在卖网课的都已经这么浮躁了吗？","是时候体验一下github action的魅力了","面试官：听说你熟悉OkHttp原理？","4000字干货长文！从校招和社招的角度说说如何准备大厂面试？","复盘MySQL分页查询优化方案","图编辑引擎 X6 让你们久等了","Elasticsearch 索引设计实战指南","未来魔法校的微前端实践","VS Code提高前端开发效率插件","闲鱼高效投放背后的秘密——鲲鹏","处理线上RabbitMQ队列阻塞","自从有了它，我就再也没有用过函数重写","Android Studio 3.6 稳定版发布","从原理谈闭包简单易懂，不看别后悔！","通过objc_msgSend实现iOS方法耗时监控","可计算、高性能的自定义用户字段方案","如何渲染几万条数据并不卡住界面","Chrome浏览器版本升级带来的跨域访问问题","使用纯Java实现一个WebSSH项目","技术开源项目从零到一的心路历程","原生JavaScript实现无缝轮播效果","JS中数组去重的三种方法","3月编程语言排行及程序员工资，快看看你在哪个等级！","前端该了解的HTTP、HTTPS及TCP协议（精简版）","iOS 查漏补缺 - PerformSelector","全网最易懂版：什么是立即执行函数？有什么作用？","小程序中使用Echart","教妹学Java：Spring 基础篇","爬虫管理平台 Crawlab 新功能介绍 - 用 Git 做 CI/CD","聊聊老板让我删除日志文件那些事儿","Leet-Code每日一题持续更新","2019年，我的个人总结","探讨SWIFT 5.2的新功能特性","使用消息中间件时，如何保证消息仅仅被消费一次？","从借呗双十一游戏看蚂蚁金服 Oasis 3D 工作流","jenkins + docker-compose 实现一键构建部署","VUE 3.0 学习探索入门系列 - 回顾 vue2 辉煌一生（2）","使用Prometheus + Grafana搭建监控系统","Spring Cloud 配置加载（一）—— 如何从配置中心读取配置","从实习到全干，还能走多久 - 开篇","2.时间复杂度与空间复杂度","【JavaScript】（附面试题）一文搞懂面向对象编程","Android指纹识别浅析","服务端面试经","Python一键转Jar包，Java调用Python新姿势！","五分钟新概念之宏任务与微任务","NVM中完美解决npm不好使","关于字体对齐那些事","Dart语言详解（一）——详细介绍","拒绝JavaScript，这三个CSS技巧你一定用的上","React之CLODOP实战：(一)手写并打印快递单","iTerm2 都不会用，还敢自称老司机？（上）","前端基础——HTML、浏览器篇","Vuex操作实录","🔥JVM从入门到入土之JVM的类加载机制","基于vue-cli3/cli4解决前端使用axios跨域问题","js将数值转为万，亿，万亿并保留一位小数","一文解读微服务架构的服务与发现—Spring Cloud","项目中数组循环的一个优化案例","DeepDelta：一种通过深度学习自动修复编译错误的方法","面试代码题（vivo）数位之积","《程序员有故事》怎样渡过这个漫长的假期？","Flutter 中的图文混排与原理解析","JavaScript异步机制知识整理","Service Worker简易教程","iOS探索 KVC原理及自定义","Flutter之引擎启动流程","iOS底层学习 - KVO探索之路","【SpringBoot】为什么我的定时任务不执行？","LayoutInflater--布局文件的加载","不知道javaagent是什么,运行个hello world就知道了","NodeJs入门，带你获取头条热搜形成好看的标签云","React错题本---题1——使用useState更新状态失败","Android APP/AMS/WMS之间交互总结","没用过这些IDEA插件？怪不得写代码头疼","Java的三魂七魄 —— 反射","《面试》SpringBoot启动做了哪些事？","左耳朵耗子：技术人如何更好地把控发展趋势？","外部配置属性值是如何被绑定到XxxProperties类属性上的？--SpringBoot源码（五）","Android View 的绘制流程分析及其源码调用追踪","Android框架组件 ViewBinding详解及使用","提升用户体验之局部过渡解决方案","Serverless 多环境配置方案探索","Android实现清理缓存功能","springboot+mybatis多数据源实现原理","运营平台系统在菜单权限的基础上进一步细化到按钮权限管理（Vue）","iOS 基于Service构建组件通信的思考","手把手带你画出一份属于自己的JavaScript原型图","线上一个数据库死锁问题的分析","浅谈前端性能优化（html+css篇）","一文搞定前端包管理工具","JS小案例——获取随机验证码","你应该知道的 HTTP 基础知识","源码分析Android系统后台应用启动服务crash","[译] 理解 CSS 网格布局：创建一个网格容器","vue中的echarts地图(没有找不到的地图，只有你不会的地图,内有一套地图的万能模板)","还在使用集合类完成这些功能？不妨来看看 Guava 集合类!!!","webpack中tapable原理详解，一起学习任务流程管理","前端科普系列（2）：Node.js 换个角度看世界","volatile引发的一个有趣的测试","【译】推荐你使用Vim的三个半理由","前端性能监控-window.performance解读","Node.js 入门 —— 基于 egg.js 和 socket.io 的聊天小应用","记一次zookeeper网络异常引发的dubbo服务provider丢失事故","专治MySQL乱码问题","黑客入侵攻击下的银行 App，数据安全何去何从？","《面试》SpringBoot面试题","JS中时间格式化的三种方法","责任链模式与lambda重构责任链模式","JavaScript中的this到底指向谁","电商sku组合查询状态细究与实现","『Android Q 源码分析』- Android 10.0 WatchDog源码解析","手写一个范围选择的日历组件","你一定不知道的有关HttpServletResponse和HttpServletRequest取值的两个坑","Python——五分钟带你弄懂迭代器与生成器，夯实代码能力","事件总线方案实践","进程无故消失的破案历程","在腾讯，有多少技术Leader在写代码","贼，夜入豪宅，可偷之物甚多，而负重能力有限，偷哪些才更加不枉此行？","你真的理解Promise吗？","https原理初探","浅谈Android动画的那些事","Android Studio Design Tools 中的 UX 更改 — Split View","2020了，你们还兼容IE吗？记一次React页面跳转在IE上的神奇问题","如何建立架构师的立体化思维？【1】","Android一次完美的跨进程服务共享实践","Vue->>Router版本太新，导致重复点击路由控制台报错解决方案","nginx+node+vue+node部署个简单的前后端分离应用","前端微服务简单实践"}

type Note struct {
	ID          int
	Cid         int
	Type        string
	Like        string
	Status      bool
	Title       string
	Color       string
	Keywords    string
	Description string
	CreatedAt   string
	UpdatedAt   string
}
var db = &sql.DB{}

// 动作
var action string

func init() {
	path := "./note.db"
	_, err := os.Stat(path)

	db, _ = sql.Open("sqlite3", path)
	flag.StringVar(&action, "action", "run", "动作：运行 （run）或 初始化 (initialization)")
	flag.Parse()

	if err != nil {
		initialization()
		time.Sleep(5 * 1e9)
	}
}

var noteType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Note",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"color": &graphql.Field{
				Type: graphql.String,
			},
			"cid": &graphql.Field{
				Type: graphql.Int,
			},
			"type": &graphql.Field{
				Type: graphql.String,
			},
			"like": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"note": &graphql.Field{
			Type: noteType,
			Description: "Get Note By ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "Note ID",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}

				querySql := "SELECT id, cid, type, status, title, keywords, description, created_at, updated_at FROM notes where id= ?;"
				row := db.QueryRow(querySql,id)

				var cid int
				var Type string
				var path string
				var status bool
				var title string
				var keywords string
				var description string
				var created_at string
				var updated_at string
				var like string

				_ = row.Scan(&id, &cid, &Type, &status, &title, &keywords, &description, &created_at, &updated_at)
				if Type == "note" || Type == ""  {
					Type = "note"
				}else {
					querySql := "SELECT path FROM note_file where note_id= ?;"
					row := db.QueryRow(querySql,id)
					_ = row.Scan(&path)
					Type = "file"
					like = "http://localhost:8090/" + path
				}

				Note :=  Note{ID: id, Cid:cid,Type:Type,Status:status,Title:title,Like:like}

				return Note, nil
				},
		},
		"list": &graphql.Field{
			Type: graphql.NewList(noteType),
			Description: "Get Nav List",
			Args: graphql.FieldConfigArgument{
				"cid": &graphql.ArgumentConfig{
					Description: "Note Cid",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				cid, err := strconv.Atoi(p.Args["cid"].(string))
				if err != nil {
					return nil, err
				}
				var Notes []Note


				querySql := "SELECT id, cid, type, status, title, keywords, description, created_at, updated_at FROM notes where cid= ?;"
				rows, err := db.Query(querySql,cid)

				if rows == nil {
					return nil, nil
				}
				var id int
				var Type string
				var status bool
				var title string
				var keywords string
				var description string
				var created_at string
				var updated_at string

				for rows.Next() {
					_ = rows.Scan(&id, &cid, &Type, &status, &title, &keywords, &description, &created_at, &updated_at)
					tutorial :=  Note{ID: id, Cid:cid,Type:Type,Status:true,Color: Color[rand.Intn(len(Color))],Title:title}

					Notes = append(Notes, tutorial)
				}
				return Notes, nil
			},
		},
	},
})

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        noteType,
			Description: "Create a new Tutorial",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"cid": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 0,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				note, _ := db.Exec("INSERT INTO notes('cid','title','type','status') values(?,?,?,?)", params.Args["cid"],params.Args["title"].(string) , "",1)
				LastInsertId, _ := note.LastInsertId()
				id, _ := strconv.Atoi(strconv.FormatInt(LastInsertId,10))
				return Note{ID: id, Cid:0,Type:"",Status:true,Title:params.Args["title"].(string)}, nil
			},
		},
	},
})

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    QueryType,
		Mutation: MutationType,
	})

	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.Dir("./dist")))
	http.Handle("/js", http.FileServer(http.Dir("./dist/js")))
	http.Handle("/css", http.FileServer(http.Dir("./dist/css")))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	//http.HandleFunc("/api", handler(schema))
	http.HandleFunc("/api/", handler(schema))
	http.HandleFunc("/uploadFile/", uploadFile)
	//http.HandleFunc("/", api)
	//http.HandleFunc("/", handler)
	fmt.Println("http://127.0.0.1:8090")
	err = http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(schema graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Set("content-type", "application/json")             //返回数据格式是json
		query, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//fmt.Print(1,string(query))
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: string(query),
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)

	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("file")
	id := r.PostFormValue("id")
	if err != nil {
		//fmt.Println("Error Retrieving the File")
		fmt.Println(1,err)
		return
	}
	defer file.Close()

	srt := filepath.Ext(handler.Filename)

	imgDir := "./public/" + id + "/"
	now := time.Now()

	imgDir = imgDir + strconv.Itoa(now.Year()) + "/" + "/" + strconv.Itoa(now.Day()) + "/"

	fileDir, err := os.Stat(imgDir)
	if err != nil || !fileDir.IsDir() {
		dirErr := os.MkdirAll(imgDir, os.ModePerm)
		if dirErr != nil {
			fmt.Println("create dir failed")
			os.Exit(1)
		}
	}

	filename := imgDir + strconv.Itoa(int(now.Unix())) + strconv.Itoa(int(rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))) + srt
	exists := checkExists(filename)
	if exists {
		log.Println("read data failed:", filename)
	}


	fSrc, err := ioutil.ReadAll(file)
	Ext := GetFileType(fSrc[:10])

	file, _, err = r.FormFile("file")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer file.Close()

	fp, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer fp.Close()

	size, err := io.Copy(fp, file)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	var noteId int64 = 0
	if Ext == "wps" || Ext == "docx" || Ext == "pdf" {
		note, _ := db.Exec("INSERT INTO notes('cid','title','type','status') values(?,?,?,?)", id,handler.Filename , "office",1)
		noteId, _ = note.LastInsertId()
	}

	if len(id) > 0 {
		if noteId > 0 {
			id = strconv.FormatInt(noteId,10)
		}
		_, err = db.Exec("INSERT INTO note_file('note_id','name', 'path', 'size', 'ext', 'type') values(?,?,?,?,?,?)", id,handler.Filename,filename , strconv.FormatInt(size/1024, 10) + "KB",srt,"")
	}

	if err != nil {
		log.Fatal(err)
	}

	var list = make(map[string]string)
	list["link"] = "http://127.0.0.1:8090/" + filename
	list["size"] = strconv.FormatInt(size/1024, 10) + "KB"
	list["action"] = "call-upload-result"
	list["ext"] = Ext
	jsonStr, _ := json.Marshal(list)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonStr))
}

func checkExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
func api(w http.ResponseWriter, req *http.Request) {
	data := Note{ID: 1, Cid:2,Type:"file",Status:true,Title:"测试"}

	ret := new(Ret)
	id := req.FormValue("id")
	//id := req.PostFormValue('id')

	ret.Code = 0
	ret.Param = id
	ret.Msg = "success"
	ret.Data = append(ret.Data, data)
	ret.Data = append(ret.Data, data)
	ret.Data = append(ret.Data, data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ret)
}

func initialization() {
	markSql := "CREATE TABLE notes (" +
		"id          INTEGER     NOT NULL," +
		"cid         INTEGER (4) NOT NULL," +
		"type        TEXT (255)  NOT NULL," +
		"status      INTEGER (1) NOT NULL," +
		"title       TEXT (255)," +
		"keywords    TEXT (255)," +
		"description TEXT (255)," +
		"created_at  DATETIME," +
		"updated_at  DATETIME," +
		"PRIMARY KEY (" +
		"id" +
		")" +
		");"
	smt, err := db.Prepare(markSql)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = smt.Exec()

	markSql = "CREATE INDEX note_index ON notes ( id DESC, cid, type, status, title, keywords );"
	smt, err = db.Prepare(markSql)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = smt.Exec()
	markSql = "CREATE TABLE note_data (" +
		"id         INTEGER     NOT NULL," +
		"note_id    INTEGER (4) NOT NULL," +
		"content    TEXT," +
		"created_at DATETIME," +
		"updated_at DATETIME," +
		"PRIMARY KEY (id)" +
		");"
	smt, err = db.Prepare(markSql)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = smt.Exec()

	markSql = "CREATE INDEX note_data_index ON note_data ( id DESC, note_id);"
	smt, err = db.Prepare(markSql)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = smt.Exec()

	markSql = "CREATE TABLE note_file (" +
		"id         INTEGER     NOT NULL," +
		"note_id    INTEGER (4) NOT NULL," +
		"name    TEXT," +
		"path    TEXT," +
		"size    INTEGER," +
		"ext    TEXT," +
		"type    TEXT," +
		"created_at DATETIME," +
		"updated_at DATETIME," +
		"PRIMARY KEY (id)" +
		");"
	smt, err = db.Prepare(markSql)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = smt.Exec()

	markSql = "CREATE INDEX note_data_index ON note_data ( id DESC, note_id);"
	smt, err = db.Prepare(markSql)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = smt.Exec()
}

var fileTypeMap sync.Map

func initFileTypeMap() {
	fileTypeMap.Store("ffd8ffe000104a464946", "jpg")  //JPEG (jpg)
	fileTypeMap.Store("89504e470d0a1a0a0000", "png")  //PNG (png)
	fileTypeMap.Store("47494638396126026f01", "gif")  //GIF (gif)
	fileTypeMap.Store("49492a00227105008037", "tif")  //TIFF (tif)
	fileTypeMap.Store("424d228c010000000000", "bmp")  //16色位图(bmp)
	fileTypeMap.Store("424d8240090000000000", "bmp")  //24位位图(bmp)
	fileTypeMap.Store("424d8e1b030000000000", "bmp")  //256色位图(bmp)
	fileTypeMap.Store("41433130313500000000", "dwg")  //CAD (dwg)
	fileTypeMap.Store("3c21444f435459504520", "html") //HTML (html)   3c68746d6c3e0  3c68746d6c3e0
	fileTypeMap.Store("3c68746d6c3e0", "html")        //HTML (html)   3c68746d6c3e0  3c68746d6c3e0
	fileTypeMap.Store("3c21646f637479706520", "htm")  //HTM (htm)
	fileTypeMap.Store("48544d4c207b0d0a0942", "css")  //css
	fileTypeMap.Store("696b2e71623d696b2e71", "js")   //js
	fileTypeMap.Store("7b5c727466315c616e73", "rtf")  //Rich Text Format (rtf)
	fileTypeMap.Store("38425053000100000000", "psd")  //Photoshop (psd)
	fileTypeMap.Store("46726f6d3a203d3f6762", "eml")  //Email [Outlook Express 6] (eml)
	fileTypeMap.Store("d0cf11e0a1b11ae10000", "doc")  //MS Excel 注意：word、msi 和 excel的文件头一样
	fileTypeMap.Store("d0cf11e0a1b11ae10000", "vsd")  //Visio 绘图
	fileTypeMap.Store("5374616E64617264204A", "mdb")  //MS Access (mdb)
	fileTypeMap.Store("252150532D41646F6265", "ps")
	fileTypeMap.Store("255044462d312e350d0a", "pdf")  //Adobe Acrobat (pdf)
	fileTypeMap.Store("255044462d312e370a25", "pdf")  //Adobe Acrobat (pdf)
	fileTypeMap.Store("2e524d46000000120001", "rmvb") //rmvb/rm相同
	fileTypeMap.Store("464c5601050000000900", "flv")  //flv与f4v相同
	fileTypeMap.Store("00000020667479706d70", "mp4")
	fileTypeMap.Store("49443303000000002176", "mp3")
	fileTypeMap.Store("000001ba210001000180", "mpg") //
	fileTypeMap.Store("3026b2758e66cf11a6d9", "wmv") //wmv与asf相同
	fileTypeMap.Store("52494646e27807005741", "wav") //Wave (wav)
	fileTypeMap.Store("52494646d07d60074156", "avi")
	fileTypeMap.Store("4d546864000000060001", "mid") //MIDI (mid)
	fileTypeMap.Store("504b0304140000000800", "zip")
	fileTypeMap.Store("526172211a0700cf9073", "rar")
	fileTypeMap.Store("235468697320636f6e66", "ini")
	fileTypeMap.Store("504b03040a0000000000", "jar")
	fileTypeMap.Store("4d5a9000030000000400", "exe")        //可执行文件
	fileTypeMap.Store("3c25402070616765206c", "jsp")        //jsp文件
	fileTypeMap.Store("4d616e69666573742d56", "mf")         //MF文件
	fileTypeMap.Store("3c3f786d6c2076657273", "xml")        //xml文件
	fileTypeMap.Store("494e5345525420494e54", "sql")        //xml文件
	fileTypeMap.Store("7061636b616765207765", "java")       //java文件
	fileTypeMap.Store("406563686f206f66660d", "bat")        //bat文件
	fileTypeMap.Store("1f8b0800000000000000", "gz")         //gz文件
	fileTypeMap.Store("6c6f67346a2e726f6f74", "properties") //bat文件
	fileTypeMap.Store("cafebabe0000002e0041", "class")      //bat文件
	fileTypeMap.Store("49545346030000006000", "chm")        //bat文件
	fileTypeMap.Store("04000000010000001300", "mxp")        //bat文件
	fileTypeMap.Store("504b0304140006000800", "docx")       //docx文件
	fileTypeMap.Store("d0cf11e0a1b11ae10000", "wps")        //WPS文字wps、表格et、演示dps都是一样的
	fileTypeMap.Store("6431303a637265617465", "torrent")
	fileTypeMap.Store("6D6F6F76", "mov")         //Quicktime (mov)
	fileTypeMap.Store("FF575043", "wpd")         //WordPerfect (wpd)
	fileTypeMap.Store("CFAD12FEC5FD746F", "dbx") //Outlook Express (dbx)
	fileTypeMap.Store("2142444E", "pst")         //Outlook (pst)
	fileTypeMap.Store("AC9EBD8F", "qdf")         //Quicken (qdf)
	fileTypeMap.Store("E3828596", "pwl")         //Windows Password (pwl)
	fileTypeMap.Store("2E7261FD", "ram")         //Real Audio (ram)
	fileTypeMap.Store("fileTypeMap", "fileTypeMap")
}

// 获取前面结果字节的二进制
func bytesToHexString(src []byte) string {
	res := bytes.Buffer{}
	if src == nil || len(src) <= 0 {
		return ""
	}
	temp := make([]byte, 0)
	for _, v := range src {
		sub := v & 0xFF
		hv := hex.EncodeToString(append(temp, sub))
		if len(hv) < 2 {
			res.WriteString(strconv.FormatInt(int64(0), 10))
		}
		res.WriteString(hv)
	}
	return res.String()
}

// 用文件前面几个字节来判断
// fSrc: 文件字节流（就用前面几个字节）
func GetFileType(fSrc []byte) string {
	var fileType string
	fileCode := bytesToHexString(fSrc)
	_, ok := fileTypeMap.Load("fileTypeMap")

	if !ok {
		initFileTypeMap()
	}

	fileTypeMap.Range(func(key, value interface{}) bool {
		k := key.(string)
		v := value.(string)
		if strings.HasPrefix(fileCode, strings.ToLower(k)) ||
			strings.HasPrefix(k, strings.ToLower(fileCode)) {
			fileType = v
			return false
		}
		return true
	})
	return fileType
}