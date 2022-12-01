// cocoaGenerator ，ModBus 协议自动生成宝具，解放双手，得空摸鱼🐟！
// Powered By Luckykeeper <luckykeeper@luckykeeper.site | https://luckykeeper.site> 2022/11/27

package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xuri/excelize/v2"
	"github.com/xuthus5/BaiduTranslate"
)

// 显示中文
// 设置环境变量   通过go-findfont 寻找simkai.ttf 字体
func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		if strings.Contains(path, "simkai.ttf") {
			fmt.Println(path)
			os.Setenv("FYNE_FONT", path) // 设置环境变量  // 取消环境变量 os.Unsetenv("FYNE_FONT")
			break
		}
	}
}

var (
	Content              *fyne.Container
	ContentTitle         *widget.Label
	ContentTitleBox      *fyne.Container
	contentBox           *fyne.Container
	RunTime              = 0
	LayoutChanged        = false
	sql_initialize_table = `CREATE TABLE cocoa (platform TEXT PRIMARY KEY NOT NULL,appid TEXT NOT NULL,secret TEXT NOT NULL);`
)

func main() {
	// 先初始化数据库
	if exists, _ := PathExists("./cocoaGenerator.db"); exists {
		log.Println("数据库存在")
	} else {
		log.Println("创建数据库！")
		db, _ := sql.Open("sqlite3", "./cocoaGenerator.db")
		defer db.Close()
		db.Exec(sql_initialize_table) //初始化
	}
	// App 基本信息
	a := app.NewWithID("cocoagenerator.luckykeeper.site")
	logo, _ := fyne.LoadResourceFromPath("cocoa.ico")
	a.SetIcon(logo)
	makeTray(a)
	logLifecycle(a)
	w := a.NewWindow("esphomeYAMLGenerator, A software to generate ESPHome YAML File | Powered by Luckykeeper | Build 20221127 | Ver 1.0.0")
	w.SetMainMenu(makeMenu(a, w))

	// 左侧菜单
	menu := container.NewVBox(
		widget.NewButtonWithIcon("Welcome esphomeYAMLGenerator!",
			theme.HomeIcon(),
			welcomeScreen),
		widget.NewButtonWithIcon("翻译 API 设定",
			theme.SettingsIcon(),
			translateAPISetting),
		widget.NewButtonWithIcon("来生成一下 YAML 文件吧~",
			theme.DocumentPrintIcon(),
			database),
		widget.NewButtonWithIcon("奏了奏了~",
			theme.LogoutIcon(),
			func() { fyne.App.Quit(a) }),
	)

	left := container.New(layout.NewHBoxLayout(), menu, widget.NewSeparator())

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			if RunTime == 0 {
				ContentTitle = widget.NewLabel("雷猴哇~<(￣︶￣)↗[GO!]")
				ContentTitleBox = container.New(layout.NewVBoxLayout(), ContentTitle, widget.NewSeparator())

				addImageIcon := canvas.NewImageFromFile("./img/cocoa.png")
				addImageIcon.FillMode = canvas.ImageFillContain
				addImageIcon.SetMinSize(fyne.NewSize(500, 500))

				Content = container.NewCenter(container.NewVBox(
					widget.NewLabelWithStyle("↑心爱酱可爱捏", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
					widget.NewLabelWithStyle("Welcome to esphomeYAMLGenerator, A software to generate ESPHome YAML File", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
					container.NewHBox(
						widget.NewHyperlink("Powered By Luckykeeper", parseURL("https://luckykeeper.site/")),
						widget.NewLabel("-"),
						widget.NewHyperlink("Github", parseURL("https://github.com/luckykeeper/")),
						widget.NewLabel("-"),
						widget.NewHyperlink("Blog", parseURL("https://luckykeeper.site/")),
						widget.NewLabel("-"),
						widget.NewHyperlink("来找心爱酱玩~", parseURL("https://github.com/luckykeeper/LOVE69_renpy_remaster")),
					),
				))
				Content = container.New(layout.NewVBoxLayout(), addImageIcon, Content)

			}
			contentBox = container.New(layout.NewBorderLayout(ContentTitleBox, nil, nil, nil), ContentTitleBox, Content)
			if RunTime == 0 || LayoutChanged {
				contentBox.Refresh()
				RunTime = 1
				LayoutChanged = false
			}
			// 显示主界面，分别：适应宽度，左侧菜单，分割线，右侧内容
			w.SetContent(container.New(layout.NewBorderLayout(nil, nil, left, nil), left, contentBox))
		}
	}()

	// 设置窗口大小
	w.Resize(fyne.NewSize(1280, 720))
	w.SetFixedSize(true)
	// 润！
	w.ShowAndRun()

}

// 欢迎界面
func welcomeScreen() {
	ContentTitle = widget.NewLabel("雷猴哇~<(￣︶￣)↗[GO!]")
	ContentTitleBox = container.New(layout.NewVBoxLayout(), ContentTitle, widget.NewSeparator())

	addImageIcon := canvas.NewImageFromFile("./img/cocoa.png")
	addImageIcon.FillMode = canvas.ImageFillContain
	addImageIcon.SetMinSize(fyne.NewSize(500, 500))

	Content = container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("↑心爱酱可爱捏", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("Welcome to esphomeYAMLGenerator, A software to generate ESPHome YAML File", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		container.NewHBox(
			widget.NewHyperlink("Powered By Luckykeeper", parseURL("https://luckykeeper.site/")),
			widget.NewLabel("-"),
			widget.NewHyperlink("Github", parseURL("https://github.com/luckykeeper/")),
			widget.NewLabel("-"),
			widget.NewHyperlink("Blog", parseURL("https://luckykeeper.site/")),
			widget.NewLabel("-"),
			widget.NewHyperlink("来找心爱酱玩~", parseURL("https://github.com/luckykeeper/LOVE69_renpy_remaster")),
		),
	))

	Content = container.New(layout.NewVBoxLayout(), addImageIcon, Content)

	contentBox = container.New(layout.NewBorderLayout(ContentTitleBox, nil, nil, nil), ContentTitleBox, Content)
	LayoutChanged = true
}

// 翻译 API 设置界面
func translateAPISetting() {
	baiduDataExists, baiduAppID, baiduSecret := DataExists("baidu")

	ContentTitle = widget.NewLabel("在这里设置翻译 API ！ヾ(≧▽≦*)o")
	ContentTitleBox = container.New(layout.NewVBoxLayout(), ContentTitle, widget.NewSeparator())

	// 添加图片小技巧：container 套 container ，注意 layout.NewVBoxLayout （或者 HBox ）的时候会按照组件最小大小来排列
	// 下面带字的会自动计算最小大小，但是你加的图片不会，所以你需要手动给它一个大小，不然就会被压成一个 1x1 的像素点（乐）
	addImageIcon := canvas.NewImageFromFile("./img/yuuka.jpg")
	addImageIcon.FillMode = canvas.ImageFillContain
	addImageIcon.SetMinSize(fyne.NewSize(464, 329.6))

	inputBaiduAppID := widget.NewEntry()
	inputBaiduAppID.SetPlaceHolder("在这里填写APPID")
	if baiduDataExists {
		inputBaiduAppID.SetPlaceHolder(baiduAppID)
	}

	inputBaiduAppSecret := widget.NewEntry()
	inputBaiduAppSecret.SetPlaceHolder("在这里填写Secret")
	if baiduDataExists {
		inputBaiduAppSecret.SetPlaceHolder(baiduSecret)
	}

	NoneInputBaiduErrorCode := widget.NewEntry()

	explainInputBaiduAppID := widget.NewLabel("百度翻译API APPID：")
	explainInputBaiduSecret := widget.NewLabel("百度翻译API Secret：")
	explainInputBaiduErrorCode := widget.NewLabel("下面的输入框显示API测试结果")

	baiduAppIDBox := container.New(layout.NewGridLayout(2), explainInputBaiduAppID, inputBaiduAppID)
	baiduAppSecretBox := container.New(layout.NewGridLayout(2), explainInputBaiduSecret, inputBaiduAppSecret)

	Content = container.NewCenter(
		container.NewVBox(
			widget.NewLabelWithStyle("在下面设置翻译API，目前支持百度翻译API", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			baiduAppIDBox,
			baiduAppSecretBox,
			container.NewVBox(
				widget.NewButton("填好之后戳这里保存~（请勿重复点击）",
					func() { saveTranslateAPISetting("baidu", inputBaiduAppID.Text, inputBaiduAppSecret.Text) },
				),
				container.NewVBox(
					widget.NewButton("然后在这里测试API",
						func() {
							errorCode := testBaiduAPI()
							if errorCode != "" {
								NoneInputBaiduErrorCode.SetPlaceHolder(errorCode)
								NoneInputBaiduErrorCode.Refresh()
							}
						},
					)),
				explainInputBaiduErrorCode, NoneInputBaiduErrorCode,
			),
			container.NewHBox(
				widget.NewHyperlink("申请API", parseURL("https://api.fanyi.baidu.com/product/11")),
				widget.NewLabel("-"),
				widget.NewHyperlink("可用额度查询", parseURL("https://api.fanyi.baidu.com/api/trans/product/desktop")),
				widget.NewLabel("-"),
				widget.NewHyperlink("错误码说明", parseURL("https://api.fanyi.baidu.com/doc/21"))),
		))

	Content = container.New(layout.NewVBoxLayout(), addImageIcon, Content)

	contentBox = container.New(layout.NewBorderLayout(ContentTitleBox, nil, nil, nil), ContentTitleBox, Content)

	LayoutChanged = true
}

// 保存翻译平台 API 参数到数据库
func saveTranslateAPISetting(platform, appid, secret string) {
	baiduDataExists, _, _ := DataExists("baidu")
	db, _ := sql.Open("sqlite3", "./cocoaGenerator.db")
	defer db.Close()
	if baiduDataExists {
		db.Exec("UPDATE cocoa SET appid='" + appid + "',secret='" + secret + "' WHERE platform='" + platform + "';")
	} else if !baiduDataExists {
		db.Exec("insert into cocoa (platform, appid,secret) values ('" + platform + "','" + appid + "','" + secret + "');")
	}
}

// 判断翻译平台 API 参数是否存在
func DataExists(platform string) (result bool, appid, secret string) {
	db, _ := sql.Open("sqlite3", "./cocoaGenerator.db")
	defer db.Close()
	querySql := "select appid,secret from cocoa where platform='" + platform + "';"
	var data, data1 string
	queryResult := db.QueryRow(querySql).Scan(&data, &data1)
	if queryResult == sql.ErrNoRows {
		result = false
		appid = "nil"
		secret = "nil"
		return
	} else { // 之前设置过相关参数
		result = true
		appid = data
		secret = data1
		return
	}
}

// 测试百度翻译 API
func testBaiduAPI() (errorCode string) {
	db, _ := sql.Open("sqlite3", "./cocoaGenerator.db")
	defer db.Close()
	queryBaiduAccountSql := "select appid,secret from cocoa where platform='baidu';"
	var baiduAppID, baiduSecret string
	queryResult := db.QueryRow(queryBaiduAccountSql).Scan(&baiduAppID, &baiduSecret)
	if queryResult == sql.ErrNoRows {
		return "请先填写上方参数"
	} else { // 数据库存在参数
		bi := BaiduTranslate.BaiduInfo{AppID: baiduAppID, Salt: BaiduTranslate.Salt(12), SecretKey: baiduSecret, From: "auto", To: "en"}
		bi.Text = "你好"
		// log.Println(bi.Translate())
		if bi.Translate() == "Hello" {
			return "百度翻译API测试成功！可正常生成YAML文件！"
		} else {
			return bi.Translate()
		}

	}
}

// 判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 库存界面
func database() {
	ContentTitle = widget.NewLabel("来这里生成YAML文件吧o(*≧▽≦)ツ┏━┓")
	ContentTitleBox = container.New(layout.NewVBoxLayout(), ContentTitle, widget.NewSeparator())

	addImageIcon := canvas.NewImageFromFile("./img/hoshino.jpg")
	addImageIcon.FillMode = canvas.ImageFillContain
	addImageIcon.SetMinSize(fyne.NewSize(259.33, 366.67))

	Content = container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("先到device文件夹填写上位机设备模板【对应的IDF和Arduino框架文件应存在】", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("再去modbus文件夹填写上位机及从站参数，最后使用下面的生成功能吧", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("执行生成前务必保存并关闭相关Excel文件，生成完成后到generate文件夹找生成文件", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("【注意目前仅支持ESPHome+ModBus设备】", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("【使用生成功能会清空generate文件夹下的全部文件，请务必注意备份】", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		container.NewHBox(
			widget.NewButtonWithIcon("生成arduino框架YAML文件（调试环境使用）",
				theme.DocumentIcon(),
				// generateArduino,
				func() {
					cleanGenerate()
					generateArduino()
				},
			),
			widget.NewButtonWithIcon("生成esp-idf框架YAML文件（生产环境使用）",
				theme.DocumentIcon(),
				// generateESPIDF,
				func() {
					cleanGenerate()
					generateESPIDF()
				},
			),
			widget.NewButtonWithIcon("同时生成两种框架的YAML文件",
				theme.DocumentIcon(),
				// generateAll,
				func() {
					cleanGenerate()
					generateArduino()
					generateESPIDF()
				},
			),
		)))
	Content = container.New(layout.NewVBoxLayout(), addImageIcon, Content)

	contentBox = container.New(layout.NewBorderLayout(ContentTitleBox, nil, nil, nil), ContentTitleBox, Content)
	LayoutChanged = true
}

// 生成ardiuno框架YAML文件
func generateArduino() {
	// cleanGenerate()

	// 寻找有效上位机及从站配置文件
	dir, _ := ioutil.ReadDir("./modbus")
	for _, file := range dir {
		configFileName := file.Name()
		configFileNameAll := path.Base(configFileName)
		if ext := path.Ext(configFileName); ext == ".xlsx" {
			if configFileNameAll == "ModBus_Template.xlsx" {
				log.Println("检测到模板文件，跳过！")
			} else {
				// 在这里处理检测到的文件
				log.Println("待处理文件全名:", configFileNameAll)
				// 读取上位机设置参数
				deviceTemplate, deviceName, apiPassword, otaPassword, wifiSSID, wifiPassword, staticIP,
					netGateway, netmask, dnsServer, fallbackSSID, fallbackPassword, esphomeWebPort, esphomeWebUser, esphomeWebPassword := readPLCconfig(configFileNameAll)
				// 调试：输出读取到的上位机设置参数
				// log.Println("设备模板：", deviceTemplate)
				// log.Println("设备名称：", deviceName)
				// log.Println("api密码：", apiPassword)
				// log.Println("ota密码：", otaPassword)
				// log.Println("wifi名称：", wifiSSID)
				// log.Println("wifi密码：", wifiPassword)
				// log.Println("固定IP地址：", staticIP)
				// log.Println("网关地址：", netGateway)
				// log.Println("子网掩码：", netmask)
				// log.Println("DNS服务器：", dnsServer)
				// log.Println("回退WiFi名称：", fallbackSSID)
				// log.Println("回退WiFi密码：", fallbackPassword)
				// log.Println("web 端口：", esphomeWebPort)
				// log.Println("web 用户：", esphomeWebUser)
				// log.Println("web 密码：", esphomeWebPassword)
				log.Println("____________")

				// 读取并处理上位机设置模板（ardiuno框架）
				// 读取并复制（创建）目标文件
				destYAMLFile := "./generate/" + deviceName + "_arduino.yaml"
				copy("./device/arduino/"+deviceTemplate+".txt", destYAMLFile, 64)
				// 替换关键词，完成上位机设置
				tempbuffer, _ := ioutil.ReadFile(destYAMLFile)
				templateContent := string(tempbuffer)
				//替换写入上位机配置参数
				tempContent := strings.Replace(templateContent, "cocoaGenerator", deviceName, -1)
				tempContent = strings.Replace(tempContent, "cocoaApiPassword", apiPassword, -1)
				tempContent = strings.Replace(tempContent, "cocoaOtaPassword", otaPassword, -1)
				tempContent = strings.Replace(tempContent, "cocoaWifiName", wifiSSID, -1)
				tempContent = strings.Replace(tempContent, "cocoaWifiPassword", wifiPassword, -1)
				tempContent = strings.Replace(tempContent, "cocoaDeviceIp", staticIP, -1)
				tempContent = strings.Replace(tempContent, "cocoaDeviceGateway", netGateway, -1)
				tempContent = strings.Replace(tempContent, "cocoaDeviceNetmask", netmask, -1)
				tempContent = strings.Replace(tempContent, "cocoaDeviceDns", dnsServer, -1)
				tempContent = strings.Replace(tempContent, "CocoaFallback", fallbackSSID, -1)
				tempContent = strings.Replace(tempContent, "CocoaEsp32!", fallbackPassword, -1)
				tempContent = strings.Replace(tempContent, "cocoaWebPort", esphomeWebPort, -1)
				tempContent = strings.Replace(tempContent, "cocoawebUser", esphomeWebUser, -1)
				tempContent = strings.Replace(tempContent, "cocoawebPassword", esphomeWebPassword, -1)

				//重新写回
				ioutil.WriteFile(destYAMLFile, []byte(tempContent), 0)
				// —————— 上位机数据处理完成，下面开始处理从站数据 ——————
				// Step1：采集从站配置数据并写入 sqlite
				// Step2：生成 modbus_controller
				// Step3：按照 modbus_controller 顺序生成 sensor
				// Step4: 按照 modbus_controller 顺序生成 binary_sensor
				// 需要进行地址去重，这里把数据写入 sqlite 后借用 sql 语句完成此工作

				// Step1：采集从站配置数据并写入 sqlite
				// 为了后续排序处理方便,统一将十六进制的从站地址转换成十进制处理
				sql_drop_modbus_tables := `DROP TABLE modbus;`
				sql_initialize_modbus_tables := `CREATE TABLE modbus (sensorType TEXT NOT NULL,modbusAddress TEXT NOT NULL,modbusName TEXT NOT NULL,
					functionCode TEXT NOT NULL,registerName TEXT NOT NULL,registerAddress TEXT NOT NULL,registerCount TEXT NOT NULL,unit TEXT,
					dataType TEXT,accuracy TEXT,haClass TEXT,multiply TEXT,invert TEXT);`
				db, _ := sql.Open("sqlite3", "./cocoaGenerator.db")
				defer db.Close()
				db.Exec(sql_drop_modbus_tables)
				db.Exec(sql_initialize_modbus_tables)

				// 依次读取各行从站配置数据
				modbusConfigExcel, _ := excelize.OpenFile("./modbus/" + configFileNameAll)
				for line := 5; line >= 2; line++ {
					sensorType, _ := modbusConfigExcel.GetCellValue("从站设置", "B"+strconv.Itoa(line))    // 传感器类型
					modbusAddress, _ := modbusConfigExcel.GetCellValue("从站设置", "C"+strconv.Itoa(line)) // 从站地址

					// 当从站地址是十六进制时,需要转十进制存储
					if strings.Contains(modbusAddress, "0x") {
						modbusAddressInt64, _ := strconv.ParseUint(modbusAddress[2:], 16, 32)
						modbusAddress = fmt.Sprintf("%d", modbusAddressInt64)
					}

					modbusName, _ := modbusConfigExcel.GetCellValue("从站设置", "D"+strconv.Itoa(line))      // 从站名称
					functionCode, _ := modbusConfigExcel.GetCellValue("从站设置", "E"+strconv.Itoa(line))    // 功能码
					registerName, _ := modbusConfigExcel.GetCellValue("从站设置", "F"+strconv.Itoa(line))    // 寄存器数据名称
					registerAddress, _ := modbusConfigExcel.GetCellValue("从站设置", "G"+strconv.Itoa(line)) // 寄存器地址
					registerCount, _ := modbusConfigExcel.GetCellValue("从站设置", "H"+strconv.Itoa(line))   // 寄存器个数
					unit, _ := modbusConfigExcel.GetCellValue("从站设置", "I"+strconv.Itoa(line))            // 单位
					dataType, _ := modbusConfigExcel.GetCellValue("从站设置", "J"+strconv.Itoa(line))        // 数据类型
					accuracy, _ := modbusConfigExcel.GetCellValue("从站设置", "K"+strconv.Itoa(line))        // 数据精度
					haClass, _ := modbusConfigExcel.GetCellValue("从站设置", "L"+strconv.Itoa(line))         // ha类别
					multiply, _ := modbusConfigExcel.GetCellValue("从站设置", "M"+strconv.Itoa(line))        // 变比
					invert, _ := modbusConfigExcel.GetCellValue("从站设置", "N"+strconv.Itoa(line))          // 反转

					if sensorType == "" {
						line = 0
					} else {
						// // 调试：输出读取到的从站配置参数
						// log.Println("传感器类型:", sensorType)
						// log.Println("从站地址:", modbusAddress)
						// log.Println("从站名称:", modbusName)
						// log.Println("功能码:", functionCode)
						// log.Println("寄存器数据名称：", registerName)
						// log.Println("寄存器地址：", registerAddress)
						// log.Println("寄存器个数：", registerCount)
						// log.Println("单位：", unit)
						// log.Println("数据类型：", dataType)
						// log.Println("数据精度：", accuracy)
						// log.Println("ha类别：", haClass)
						// log.Println("变比：", multiply)
						// log.Println("反转：", invert)
						// log.Println("____________")

						// 写入数据库
						sql_writeTo_modbus_table := "insert into modbus (sensorType, modbusAddress,modbusName,functionCode," +
							"registerName,registerAddress,registerCount,unit,dataType,accuracy,haClass,multiply," +
							"invert) values ('" + sensorType + "','" + modbusAddress + "','" + modbusName + "','" + functionCode + "','" + registerName + "','" +
							registerAddress + "','" + registerCount + "','" + unit + "','" + dataType + "','" + accuracy + "','" +
							haClass + "','" + multiply + "','" + invert + "');"
						// log.Println("sqlWrite:", sql_writeTo_modbus_table)
						db.Exec(sql_writeTo_modbus_table)

					}
				}
				log.Println("Step1完成！")
				// 以上, Step1 完成
				// Step2：生成 modbus_controller

				// ardiuno 框架的 modbus_controller 示例如下

				// modbus_controller:
				// 	# 01 康耐德温湿度传感器
				// 	- id: ${device_name}_modbus_controller_01
				//    modbus_id: ${device_name}_modbus
				//    address: 0x01
				//    command_throttle: 200ms
				//    setup_priority: -10
				//    update_interval: 10s

				// 选取去重并按照 ASCII 升序排序的从站地址
				queryControllerAddressAndItsNameSql := "SELECT DISTINCT modbusAddress,modbusName FROM modbus ORDER BY modbusAddress;"

				addControllerFile, _ := os.OpenFile(destYAMLFile, os.O_APPEND, os.ModePerm)
				defer addControllerFile.Close()
				addControllerFileWriter := bufio.NewWriter(addControllerFile)
				addControllerFileWriter.WriteString("\n")
				addControllerFileWriter.WriteString("modbus_controller:\n")
				rows, _ := db.Query(queryControllerAddressAndItsNameSql)
				//  批量生成 modbus_controller
				for rows.Next() {
					var modbusAddress string
					var modbusName string
					rows.Scan(&modbusAddress, &modbusName)
					addControllerFileWriter.WriteString("  # " + modbusAddress + " " + modbusName + "\n")
					addControllerFileWriter.WriteString("  - id: ${device_name}_modbus_controller_" + modbusAddress + "\n")
					addControllerFileWriter.WriteString("    modbus_id: ${device_name}_modbus\n")
					addControllerFileWriter.WriteString("    address: " + modbusAddress + "\n")
					addControllerFileWriter.WriteString("    command_throttle: 200ms\n")
					addControllerFileWriter.WriteString("    setup_priority: -10\n")
					addControllerFileWriter.WriteString("    update_interval: 10s\n")
					addControllerFileWriter.WriteString("\n")
				}
				addControllerFileWriter.Flush()
				log.Println("Step2完成！")

				// 以上, Step2 完成
				// Step3：按照 modbus_controller 顺序生成 sensor

				// sensor 模板如下：

				// sensor:
				// 	- platform: modbus_controller
				//    modbus_controller_id: ${device_name}_modbus_controller_01
				//    id: ${device_name}_modbus_01_temp
				//    name: "Node 3# Temperature"
				//    address: 0x0258
				//    register_count: 2
				//    unit_of_measurement: "°C"
				//    register_type: holding
				//    value_type: FP32_R
				//    accuracy_decimals: 1
				//    device_class: temperature
				// 	  filters:
				//      - multiply: 0.1

				addControllerFileWriter.WriteString("sensor:\n")

				queryControllerAddressSql := "SELECT DISTINCT modbusAddress FROM modbus ORDER BY modbusAddress;"
				rows1, _ := db.Query(queryControllerAddressSql)
				//  按照地址生成 modbus_controller
				for rows1.Next() {
					var modbusAddress string
					rows1.Scan(&modbusAddress)

					// log.Println("modbusAddress:", modbusAddress)

					getModbusDetailByModbusAddressAndSensorType := "SELECT functionCode,registerName,registerAddress,registerCount,unit,dataType,accuracy,haClass,multiply,invert FROM modbus WHERE modbusAddress='" + modbusAddress + "' AND sensorType='sensor';"

					detailSensorRows, _ := db.Query(getModbusDetailByModbusAddressAndSensorType)
					for detailSensorRows.Next() {
						var functionCode, registerName, registerAddress, registerCount, unit, dataType, accuracy, haClass, multiply, invert string
						detailSensorRows.Scan(&functionCode, &registerName, &registerAddress, &registerCount, &unit, &dataType, &accuracy, &haClass, &multiply, &invert)

						var baiduAppID, baiduSecret string
						queryBaiduAccountSql := "select appid,secret from cocoa where platform='baidu';"
						db.QueryRow(queryBaiduAccountSql).Scan(&baiduAppID, &baiduSecret)
						bi := BaiduTranslate.BaiduInfo{AppID: baiduAppID, Salt: BaiduTranslate.Salt(12), SecretKey: baiduSecret, From: "auto", To: "en"}
						bi.Text = registerName
						translateRawStr := bi.Translate()
						translateStrWithoutSpace := strings.Replace(translateRawStr, " ", "_", -1)
						translateStrWithoutSpace = strings.Replace(translateStrWithoutSpace, "#", "", -1)
						translateStrWithoutSpace = strings.Replace(translateStrWithoutSpace, "__", "_", -1)
						translateStrWithoutSpace = strings.Replace(translateStrWithoutSpace, "-", "_", -1)

						// log.Println("registerName:", registerName)
						// log.Println("translateRawStr:", translateRawStr)
						// log.Println("translateStrWithoutSpace:", translateStrWithoutSpace)

						addControllerFileWriter.WriteString("\n  # " + registerName + "\n")
						addControllerFileWriter.WriteString("  - platform: modbus_controller\n")
						addControllerFileWriter.WriteString("    modbus_controller_id: ${device_name}_modbus_controller_" + modbusAddress + "\n")

						addControllerFileWriter.WriteString("    id: ${device_name}_modbus_" + modbusAddress + "_" + translateStrWithoutSpace + "\n")
						addControllerFileWriter.WriteString("    name: \"" + translateRawStr + "\"\n")
						addControllerFileWriter.WriteString("    address: " + registerAddress + "\n")
						addControllerFileWriter.WriteString("    register_count: " + registerCount + "\n")
						if unit != "" {
							addControllerFileWriter.WriteString("    unit_of_measurement: \"" + unit + "\"\n")
						}
						addControllerFileWriter.WriteString("    register_type: " + functionCode + "\n")
						if dataType != "" {
							addControllerFileWriter.WriteString("    value_type: " + dataType + "\n")
						}
						if accuracy != "" {
							addControllerFileWriter.WriteString("    accuracy_decimals: " + accuracy + "\n")
						}
						if haClass != "" {
							addControllerFileWriter.WriteString("    device_class: " + haClass + "\n")
						}

						// filters -> 针对 Sensor 和 binary_senor 通用判断语句（注意两个是不可能同时出现的）
						if multiply != "" || invert != "" {
							addControllerFileWriter.WriteString("    filters:\n")
							addControllerFileWriter.WriteString("    - ")
							if multiply != "" {
								addControllerFileWriter.WriteString("multiply: " + multiply + "\n")
							} else if invert != "" {
								addControllerFileWriter.WriteString("invert\n")
							}
						}
					}
				}
				addControllerFileWriter.Flush()
				log.Println("Step3完成！")

				// Step4: 按照 modbus_controller 顺序生成 binary_sensor

				// 以下是 binary_sensor 示例

				//   # 1#低压传感器失效
				//  - platform: modbus_controller
				// 	  modbus_controller_id: ${device_name}_modbus_controller_01
				// 	  id: ${device_name}_01_low_pressure_sensor_failure
				// 	  name: "1# low pressure sensor failure"
				// 	  register_type: holding
				// 	  address: 0x0551
				// 	  device_class: problem

				addControllerFileWriter.WriteString("\nbinary_sensor:\n")

				rows2, _ := db.Query(queryControllerAddressSql)
				//  按照地址生成 modbus_controller
				for rows2.Next() {
					var modbusAddress string
					rows2.Scan(&modbusAddress)

					// log.Println("modbusAddress:", modbusAddress)

					queryModbusBinarySensor := "SELECT functionCode,registerName,registerAddress,registerCount,unit,dataType,accuracy,haClass,multiply,invert FROM modbus WHERE modbusAddress='" + modbusAddress + "' AND sensorType='binary_sensor';"

					detailBinarySensorRows, _ := db.Query(queryModbusBinarySensor)

					for detailBinarySensorRows.Next() {
						var functionCode, registerName, registerAddress, registerCount, unit, dataType, accuracy, haClass, multiply, invert string
						detailBinarySensorRows.Scan(&functionCode, &registerName, &registerAddress, &registerCount, &unit, &dataType, &accuracy, &haClass, &multiply, &invert)

						var baiduAppID, baiduSecret string
						queryBaiduAccountSql := "select appid,secret from cocoa where platform='baidu';"
						db.QueryRow(queryBaiduAccountSql).Scan(&baiduAppID, &baiduSecret)
						bi := BaiduTranslate.BaiduInfo{AppID: baiduAppID, Salt: BaiduTranslate.Salt(12), SecretKey: baiduSecret, From: "auto", To: "en"}
						bi.Text = registerName
						translateRawStr := bi.Translate()
						translateStrWithoutSpace := strings.Replace(translateRawStr, " ", "_", -1)
						translateStrWithoutSpace = strings.Replace(translateStrWithoutSpace, "#", "", -1)
						translateStrWithoutSpace = strings.Replace(translateStrWithoutSpace, "__", "_", -1)

						// log.Println("registerName:", registerName)
						// log.Println("translateRawStr:", translateRawStr)
						// log.Println("translateStrWithoutSpace:", translateStrWithoutSpace)

						addControllerFileWriter.WriteString("\n  # " + registerName + "\n")
						addControllerFileWriter.WriteString("  - platform: modbus_controller\n")
						addControllerFileWriter.WriteString("    modbus_controller_id: ${device_name}_modbus_controller_" + modbusAddress + "\n")

						addControllerFileWriter.WriteString("    id: ${device_name}_modbus_" + modbusAddress + "_" + translateStrWithoutSpace + "\n")
						addControllerFileWriter.WriteString("    name: \"" + translateRawStr + "\"\n")
						addControllerFileWriter.WriteString("    register_type: " + functionCode + "\n")
						addControllerFileWriter.WriteString("    address: " + registerAddress + "\n")
						if haClass != "" {
							addControllerFileWriter.WriteString("    device_class: " + haClass + "\n")
						}
						// filters -> 针对 Sensor 和 binary_senor 通用判断语句（注意两个是不可能同时出现的）
						if multiply != "" || invert != "" {
							addControllerFileWriter.WriteString("    filters:\n")
							addControllerFileWriter.WriteString("    - ")
							if multiply != "" {
								addControllerFileWriter.WriteString("multiply: " + multiply + "\n")
							} else if invert != "" {
								addControllerFileWriter.WriteString("invert\n")
							}
						}

					}
				}
				addControllerFileWriter.Flush()
				log.Println("Step4完成！")
				log.Println(destYAMLFile, "->生成完成")
			}
		}

	}
}

// 生成ESP-IDF框架YAML文件
// 考虑到单独使用一种功能的情况，虽然两种生成方法类似但是没有单独提出来
func generateESPIDF() {
	// cleanGenerate()

	// 寻找有效上位机及从站配置文件
	dir, _ := ioutil.ReadDir("./modbus")
	for _, file := range dir {
		configFileName := file.Name()
		configFileNameAll := path.Base(configFileName)
		if ext := path.Ext(configFileName); ext == ".xlsx" {
			if configFileNameAll == "ModBus_Template.xlsx" {
				log.Println("检测到模板文件，跳过！")
			} else {
				// 在这里处理检测到的文件
				log.Println("待处理文件全名:", configFileNameAll)
				// 读取上位机设置参数
				deviceTemplate, deviceName, apiPassword, otaPassword, wifiSSID, wifiPassword, staticIP,
					netGateway, netmask, dnsServer, fallbackSSID, fallbackPassword, esphomeWebPort, esphomeWebUser, esphomeWebPassword := readPLCconfig(configFileNameAll)
				// 调试：输出读取到的上位机设置参数
				// log.Println("设备模板：", deviceTemplate)
				// log.Println("设备名称：", deviceName)
				// log.Println("api密码：", apiPassword)
				// log.Println("ota密码：", otaPassword)
				// log.Println("wifi名称：", wifiSSID)
				// log.Println("wifi密码：", wifiPassword)
				// log.Println("固定IP地址：", staticIP)
				// log.Println("网关地址：", netGateway)
				// log.Println("子网掩码：", netmask)
				// log.Println("DNS服务器：", dnsServer)
				// log.Println("回退WiFi名称：", fallbackSSID)
				// log.Println("回退WiFi密码：", fallbackPassword)
				// log.Println("web 端口：", esphomeWebPort)
				// log.Println("web 用户：", esphomeWebUser)
				// log.Println("web 密码：", esphomeWebPassword)
				log.Println("____________")

				// 读取并处理上位机设置模板（ardiuno框架）[说明：先按照ardiuno框架生成，再转换为IDF]
				// 读取并复制（创建）目标文件
				destYAMLFile := "./generate/" + deviceName + "_espidf.yaml"
				copy("./device/esp-idf/"+deviceTemplate+".txt", destYAMLFile, 64)
				// 替换关键词，完成上位机设置
				tempbuffer, _ := ioutil.ReadFile(destYAMLFile)
				templateContent := string(tempbuffer)
				//替换写入上位机配置参数
				tempContent := strings.Replace(templateContent, "cocoaGenerator", deviceName, -1)
				tempContent = strings.Replace(tempContent, "cocoaApiPassword", apiPassword, -1)
				tempContent = strings.Replace(tempContent, "cocoaOtaPassword", otaPassword, -1)
				tempContent = strings.Replace(tempContent, "cocoaWifiName", wifiSSID, -1)
				tempContent = strings.Replace(tempContent, "cocoaWifiPassword", wifiPassword, -1)
				tempContent = strings.Replace(tempContent, "cocoaDeviceIp", staticIP, -1)
				tempContent = strings.Replace(tempContent, "cocoaDeviceGateway", netGateway, -1)
				tempContent = strings.Replace(tempContent, "cocoaDeviceNetmask", netmask, -1)
				tempContent = strings.Replace(tempContent, "cocoaDeviceDns", dnsServer, -1)
				tempContent = strings.Replace(tempContent, "CocoaFallback", fallbackSSID, -1)
				tempContent = strings.Replace(tempContent, "CocoaEsp32!", fallbackPassword, -1)
				tempContent = strings.Replace(tempContent, "cocoaWebPort", esphomeWebPort, -1)
				tempContent = strings.Replace(tempContent, "cocoawebUser", esphomeWebUser, -1)
				tempContent = strings.Replace(tempContent, "cocoawebPassword", esphomeWebPassword, -1)

				//重新写回
				ioutil.WriteFile(destYAMLFile, []byte(tempContent), 0)
				// —————— 上位机数据处理完成，下面开始处理从站数据 ——————
				// Step1：采集从站配置数据并写入 sqlite
				// Step2：生成 modbus_controller
				// Step3：按照 modbus_controller 顺序生成 sensor
				// Step4: 按照 modbus_controller 顺序生成 binary_sensor
				// 需要进行地址去重，这里把数据写入 sqlite 后借用 sql 语句完成此工作

				// Step1：采集从站配置数据并写入 sqlite
				// 为了后续排序处理方便,统一将十六进制的从站地址转换成十进制处理
				sql_drop_modbus_tables := `DROP TABLE modbus;`
				sql_initialize_modbus_tables := `CREATE TABLE modbus (sensorType TEXT NOT NULL,modbusAddress TEXT NOT NULL,modbusName TEXT NOT NULL,
					functionCode TEXT NOT NULL,registerName TEXT NOT NULL,registerAddress TEXT NOT NULL,registerCount TEXT NOT NULL,unit TEXT,
					dataType TEXT,accuracy TEXT,haClass TEXT,multiply TEXT,invert TEXT);`
				db, _ := sql.Open("sqlite3", "./cocoaGenerator.db")
				defer db.Close()
				db.Exec(sql_drop_modbus_tables)
				db.Exec(sql_initialize_modbus_tables)

				// 依次读取各行从站配置数据
				modbusConfigExcel, _ := excelize.OpenFile("./modbus/" + configFileNameAll)
				for line := 5; line >= 2; line++ {
					sensorType, _ := modbusConfigExcel.GetCellValue("从站设置", "B"+strconv.Itoa(line))    // 传感器类型
					modbusAddress, _ := modbusConfigExcel.GetCellValue("从站设置", "C"+strconv.Itoa(line)) // 从站地址

					// 当从站地址是十六进制时,需要转十进制存储
					if strings.Contains(modbusAddress, "0x") {
						modbusAddressInt64, _ := strconv.ParseUint(modbusAddress[2:], 16, 32)
						modbusAddress = fmt.Sprintf("%d", modbusAddressInt64)
					}

					modbusName, _ := modbusConfigExcel.GetCellValue("从站设置", "D"+strconv.Itoa(line))      // 从站名称
					functionCode, _ := modbusConfigExcel.GetCellValue("从站设置", "E"+strconv.Itoa(line))    // 功能码
					registerName, _ := modbusConfigExcel.GetCellValue("从站设置", "F"+strconv.Itoa(line))    // 寄存器数据名称
					registerAddress, _ := modbusConfigExcel.GetCellValue("从站设置", "G"+strconv.Itoa(line)) // 寄存器地址
					registerCount, _ := modbusConfigExcel.GetCellValue("从站设置", "H"+strconv.Itoa(line))   // 寄存器个数
					unit, _ := modbusConfigExcel.GetCellValue("从站设置", "I"+strconv.Itoa(line))            // 单位
					dataType, _ := modbusConfigExcel.GetCellValue("从站设置", "J"+strconv.Itoa(line))        // 数据类型
					accuracy, _ := modbusConfigExcel.GetCellValue("从站设置", "K"+strconv.Itoa(line))        // 数据精度
					haClass, _ := modbusConfigExcel.GetCellValue("从站设置", "L"+strconv.Itoa(line))         // ha类别
					multiply, _ := modbusConfigExcel.GetCellValue("从站设置", "M"+strconv.Itoa(line))        // 变比
					invert, _ := modbusConfigExcel.GetCellValue("从站设置", "N"+strconv.Itoa(line))          // 反转

					if sensorType == "" {
						line = 0
					} else {
						// // 调试：输出读取到的从站配置参数
						// log.Println("传感器类型:", sensorType)
						// log.Println("从站地址:", modbusAddress)
						// log.Println("从站名称:", modbusName)
						// log.Println("功能码:", functionCode)
						// log.Println("寄存器数据名称：", registerName)
						// log.Println("寄存器地址：", registerAddress)
						// log.Println("寄存器个数：", registerCount)
						// log.Println("单位：", unit)
						// log.Println("数据类型：", dataType)
						// log.Println("数据精度：", accuracy)
						// log.Println("ha类别：", haClass)
						// log.Println("变比：", multiply)
						// log.Println("反转：", invert)
						// log.Println("____________")

						// 写入数据库
						sql_writeTo_modbus_table := "insert into modbus (sensorType, modbusAddress,modbusName,functionCode," +
							"registerName,registerAddress,registerCount,unit,dataType,accuracy,haClass,multiply," +
							"invert) values ('" + sensorType + "','" + modbusAddress + "','" + modbusName + "','" + functionCode + "','" + registerName + "','" +
							registerAddress + "','" + registerCount + "','" + unit + "','" + dataType + "','" + accuracy + "','" +
							haClass + "','" + multiply + "','" + invert + "');"
						// log.Println("sqlWrite:", sql_writeTo_modbus_table)
						db.Exec(sql_writeTo_modbus_table)

					}
				}
				log.Println("Step1完成！")
				// 以上, Step1 完成
				// Step2：生成 modbus_controller

				// ardiuno 框架的 modbus_controller 示例如下

				// modbus_controller:
				// 	# 01 康耐德温湿度传感器
				// 	- id: ${device_name}_modbus_controller_01
				//    modbus_id: ${device_name}_modbus
				//    address: 0x01
				//    command_throttle: 200ms
				//    setup_priority: -10
				//    update_interval: 10s

				// 选取去重并按照 ASCII 升序排序的从站地址
				queryControllerAddressAndItsNameSql := "SELECT DISTINCT modbusAddress,modbusName FROM modbus ORDER BY modbusAddress;"

				addControllerFile, _ := os.OpenFile(destYAMLFile, os.O_APPEND, os.ModePerm)
				defer addControllerFile.Close()
				addControllerFileWriter := bufio.NewWriter(addControllerFile)
				addControllerFileWriter.WriteString("\n")
				addControllerFileWriter.WriteString("modbus_controller:\n")
				rows, _ := db.Query(queryControllerAddressAndItsNameSql)
				//  批量生成 modbus_controller
				for rows.Next() {
					var modbusAddress string
					var modbusName string
					rows.Scan(&modbusAddress, &modbusName)
					addControllerFileWriter.WriteString("  # " + modbusAddress + " " + modbusName + "\n")
					addControllerFileWriter.WriteString("  - id: ${device_name}_modbus_controller_" + modbusAddress + "\n")
					addControllerFileWriter.WriteString("    modbus_id: ${device_name}_modbus\n")
					addControllerFileWriter.WriteString("    address: " + modbusAddress + "\n")
					addControllerFileWriter.WriteString("    command_throttle: 200ms\n")
					addControllerFileWriter.WriteString("    setup_priority: -10\n")
					addControllerFileWriter.WriteString("    update_interval: 10s\n")
					addControllerFileWriter.WriteString("\n")
				}
				addControllerFileWriter.Flush()
				log.Println("Step2完成！")

				// 以上, Step2 完成
				// Step3：按照 modbus_controller 顺序生成 sensor

				// sensor 模板如下：

				// sensor:
				// 	- platform: modbus_controller
				//    modbus_controller_id: ${device_name}_modbus_controller_01
				//    id: ${device_name}_modbus_01_temp
				//    name: "Node 3# Temperature"
				//    address: 0x0258
				//    register_count: 2
				//    unit_of_measurement: "°C"
				//    register_type: holding
				//    value_type: FP32_R
				//    accuracy_decimals: 1
				//    device_class: temperature
				// 	  filters:
				//      - multiply: 0.1

				addControllerFileWriter.WriteString("sensor:\n")

				queryControllerAddressSql := "SELECT DISTINCT modbusAddress FROM modbus ORDER BY modbusAddress;"
				rows1, _ := db.Query(queryControllerAddressSql)
				//  按照地址生成 modbus_controller
				for rows1.Next() {
					var modbusAddress string
					rows1.Scan(&modbusAddress)

					// log.Println("modbusAddress:", modbusAddress)

					getModbusDetailByModbusAddressAndSensorType := "SELECT functionCode,registerName,registerAddress,registerCount,unit,dataType,accuracy,haClass,multiply,invert FROM modbus WHERE modbusAddress='" + modbusAddress + "' AND sensorType='sensor';"

					detailSensorRows, _ := db.Query(getModbusDetailByModbusAddressAndSensorType)
					for detailSensorRows.Next() {
						var functionCode, registerName, registerAddress, registerCount, unit, dataType, accuracy, haClass, multiply, invert string
						detailSensorRows.Scan(&functionCode, &registerName, &registerAddress, &registerCount, &unit, &dataType, &accuracy, &haClass, &multiply, &invert)

						var baiduAppID, baiduSecret string
						queryBaiduAccountSql := "select appid,secret from cocoa where platform='baidu';"
						db.QueryRow(queryBaiduAccountSql).Scan(&baiduAppID, &baiduSecret)
						bi := BaiduTranslate.BaiduInfo{AppID: baiduAppID, Salt: BaiduTranslate.Salt(12), SecretKey: baiduSecret, From: "auto", To: "en"}
						bi.Text = registerName
						translateRawStr := bi.Translate()
						translateStrWithoutSpace := strings.Replace(translateRawStr, " ", "_", -1)
						translateStrWithoutSpace = strings.Replace(translateStrWithoutSpace, "#", "", -1)
						translateStrWithoutSpace = strings.Replace(translateStrWithoutSpace, "__", "_", -1)

						// log.Println("registerName:", registerName)
						// log.Println("translateRawStr:", translateRawStr)
						// log.Println("translateStrWithoutSpace:", translateStrWithoutSpace)

						addControllerFileWriter.WriteString("\n  # " + registerName + "\n")
						addControllerFileWriter.WriteString("  - platform: modbus_controller\n")
						addControllerFileWriter.WriteString("    modbus_controller_id: ${device_name}_modbus_controller_" + modbusAddress + "\n")

						addControllerFileWriter.WriteString("    id: ${device_name}_modbus_" + modbusAddress + "_" + translateStrWithoutSpace + "\n")
						addControllerFileWriter.WriteString("    name: \"" + translateRawStr + "\"\n")
						addControllerFileWriter.WriteString("    address: " + registerAddress + "\n")
						addControllerFileWriter.WriteString("    register_count: " + registerCount + "\n")
						if unit != "" {
							addControllerFileWriter.WriteString("    unit_of_measurement: \"" + unit + "\"\n")
						}
						addControllerFileWriter.WriteString("    register_type: " + functionCode + "\n")
						if dataType != "" {
							addControllerFileWriter.WriteString("    value_type: " + dataType + "\n")
						}
						if accuracy != "" {
							addControllerFileWriter.WriteString("    accuracy_decimals: " + accuracy + "\n")
						}
						if haClass != "" {
							addControllerFileWriter.WriteString("    device_class: " + haClass + "\n")
						}

						// filters -> 针对 Sensor 和 binary_senor 通用判断语句（注意两个是不可能同时出现的）
						if multiply != "" || invert != "" {
							addControllerFileWriter.WriteString("    filters:\n")
							addControllerFileWriter.WriteString("    - ")
							if multiply != "" {
								addControllerFileWriter.WriteString("multiply: " + multiply + "\n")
							} else if invert != "" {
								addControllerFileWriter.WriteString("invert\n")
							}
						}
					}
				}
				addControllerFileWriter.Flush()
				log.Println("Step3完成！")

				// Step4: 按照 modbus_controller 顺序生成 binary_sensor

				// 以下是 binary_sensor 示例

				//   # 1#低压传感器失效
				//  - platform: modbus_controller
				// 	  modbus_controller_id: ${device_name}_modbus_controller_01
				// 	  id: ${device_name}_01_low_pressure_sensor_failure
				// 	  name: "1# low pressure sensor failure"
				// 	  register_type: holding
				// 	  address: 0x0551
				// 	  device_class: problem

				addControllerFileWriter.WriteString("\nbinary_sensor:\n")

				rows2, _ := db.Query(queryControllerAddressSql)
				//  按照地址生成 modbus_controller
				for rows2.Next() {
					var modbusAddress string
					rows2.Scan(&modbusAddress)

					// log.Println("modbusAddress:", modbusAddress)

					queryModbusBinarySensor := "SELECT functionCode,registerName,registerAddress,registerCount,unit,dataType,accuracy,haClass,multiply,invert FROM modbus WHERE modbusAddress='" + modbusAddress + "' AND sensorType='binary_sensor';"

					detailBinarySensorRows, _ := db.Query(queryModbusBinarySensor)

					for detailBinarySensorRows.Next() {
						var functionCode, registerName, registerAddress, registerCount, unit, dataType, accuracy, haClass, multiply, invert string
						detailBinarySensorRows.Scan(&functionCode, &registerName, &registerAddress, &registerCount, &unit, &dataType, &accuracy, &haClass, &multiply, &invert)

						var baiduAppID, baiduSecret string
						queryBaiduAccountSql := "select appid,secret from cocoa where platform='baidu';"
						db.QueryRow(queryBaiduAccountSql).Scan(&baiduAppID, &baiduSecret)
						bi := BaiduTranslate.BaiduInfo{AppID: baiduAppID, Salt: BaiduTranslate.Salt(12), SecretKey: baiduSecret, From: "auto", To: "en"}
						bi.Text = registerName
						translateRawStr := bi.Translate()
						translateStrWithoutSpace := strings.Replace(translateRawStr, " ", "_", -1)
						translateStrWithoutSpace = strings.Replace(translateStrWithoutSpace, "#", "", -1)
						translateStrWithoutSpace = strings.Replace(translateStrWithoutSpace, "__", "_", -1)
						translateStrWithoutSpace = strings.Replace(translateStrWithoutSpace, "-", "_", -1)

						// log.Println("registerName:", registerName)
						// log.Println("translateRawStr:", translateRawStr)
						// log.Println("translateStrWithoutSpace:", translateStrWithoutSpace)

						addControllerFileWriter.WriteString("\n  # " + registerName + "\n")
						addControllerFileWriter.WriteString("  - platform: modbus_controller\n")
						addControllerFileWriter.WriteString("    modbus_controller_id: ${device_name}_modbus_controller_" + modbusAddress + "\n")

						addControllerFileWriter.WriteString("    id: ${device_name}_modbus_" + modbusAddress + "_" + translateStrWithoutSpace + "\n")
						addControllerFileWriter.WriteString("    name: \"" + translateRawStr + "\"\n")
						addControllerFileWriter.WriteString("    register_type: " + functionCode + "\n")
						addControllerFileWriter.WriteString("    address: " + registerAddress + "\n")
						if haClass != "" {
							addControllerFileWriter.WriteString("    device_class: " + haClass + "\n")
						}
						// filters -> 针对 Sensor 和 binary_senor 通用判断语句（注意两个是不可能同时出现的）
						if multiply != "" || invert != "" {
							addControllerFileWriter.WriteString("    filters:\n")
							addControllerFileWriter.WriteString("    - ")
							if multiply != "" {
								addControllerFileWriter.WriteString("multiply: " + multiply + "\n")
							} else if invert != "" {
								addControllerFileWriter.WriteString("invert\n")
							}
						}

					}
				}
				addControllerFileWriter.Flush()
				log.Println("Step4完成！")

				// Step5：关键词替换，将 arduino 文件转换为 IDF格式的文件
				tempbuffer1, _ := ioutil.ReadFile(destYAMLFile)
				templateContent1 := string(tempbuffer1)
				templateContent1 = strings.Replace(templateContent1, "${device_name}", deviceName, -1)

				//重新写回
				ioutil.WriteFile(destYAMLFile, []byte(templateContent1), 0)

				log.Println("Step5完成！")
				log.Println(destYAMLFile, "->生成完成")
			}
		}

	}
}

// 读取上位机设置参数
func readPLCconfig(fileName string) (deviceTemplate, deviceName, apiPassword, otaPassword, wifiSSID, wifiPassword, staticIP,
	netGateway, netmask, dnsServer, fallbackSSID, fallbackPassword, esphomeWebPort, esphomeWebUser, esphomeWebPassword string) {
	// 上位机的设置文件
	plcConfigExcel, _ := excelize.OpenFile("./modbus/" + fileName)
	for line := 4; line >= 2; line++ {
		deviceTemplate, _ := plcConfigExcel.GetCellValue("上位机设置", "B"+strconv.Itoa(line))     // 设备模板
		deviceName, _ := plcConfigExcel.GetCellValue("上位机设置", "C"+strconv.Itoa(line))         // 设备名称
		apiPassword, _ := plcConfigExcel.GetCellValue("上位机设置", "D"+strconv.Itoa(line))        // api密码
		otaPassword, _ := plcConfigExcel.GetCellValue("上位机设置", "E"+strconv.Itoa(line))        // ota密码
		wifiSSID, _ := plcConfigExcel.GetCellValue("上位机设置", "F"+strconv.Itoa(line))           // wifi名称
		wifiPassword, _ := plcConfigExcel.GetCellValue("上位机设置", "G"+strconv.Itoa(line))       // wifi密码
		staticIP, _ := plcConfigExcel.GetCellValue("上位机设置", "H"+strconv.Itoa(line))           // 固定IP地址
		netGateway, _ := plcConfigExcel.GetCellValue("上位机设置", "I"+strconv.Itoa(line))         // 网关地址
		netmask, _ := plcConfigExcel.GetCellValue("上位机设置", "J"+strconv.Itoa(line))            // 子网掩码
		dnsServer, _ := plcConfigExcel.GetCellValue("上位机设置", "K"+strconv.Itoa(line))          // DNS服务器
		fallbackSSID, _ := plcConfigExcel.GetCellValue("上位机设置", "L"+strconv.Itoa(line))       // 回退WiFi名称
		fallbackPassword, _ := plcConfigExcel.GetCellValue("上位机设置", "M"+strconv.Itoa(line))   // 回退WiFi密码
		esphomeWebPort, _ := plcConfigExcel.GetCellValue("上位机设置", "N"+strconv.Itoa(line))     // web 端口
		esphomeWebUser, _ := plcConfigExcel.GetCellValue("上位机设置", "O"+strconv.Itoa(line))     // web 用户
		esphomeWebPassword, _ := plcConfigExcel.GetCellValue("上位机设置", "P"+strconv.Itoa(line)) // web 密码
		if deviceTemplate == "" {
			line = 0
		} else {
			return deviceTemplate, deviceName, apiPassword, otaPassword, wifiSSID, wifiPassword, staticIP,
				netGateway, netmask, dnsServer, fallbackSSID, fallbackPassword, esphomeWebPort, esphomeWebUser, esphomeWebPassword
		}
	}
	return
}

// 清空生成文件夹已有数据并重新创建文件夹
func cleanGenerate() {
	generatePathExist, _ := PathExists("./generate")
	if generatePathExist {
		os.RemoveAll("./generate")
		os.Mkdir("./generate", os.ModePerm)
	} else {
		os.Mkdir("./generate", os.ModePerm)
	}
}

// 生命周期日志
func logLifecycle(a fyne.App) {
	a.Lifecycle().SetOnStarted(func() {
		log.Println("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		log.Println("Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		log.Println("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		log.Println("Lifecycle: Exited Foreground")
	})
}

// 顶部菜单
func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	aboutMenu := fyne.NewMenu("关于",
		fyne.NewMenuItem("访问作者博客", func() {
			u, _ := url.Parse("https://luckykeeper.site")
			_ = a.OpenURL(u)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("访问 Github —— 查看我的其它开源项目", func() {
			u, _ := url.Parse("https://github.com/luckykeeper")
			_ = a.OpenURL(u)
		}),
	)

	main := fyne.NewMainMenu(
		aboutMenu,
	)
	return main
}

// 任务栏托盘
func makeTray(a fyne.App) {
	if desk, ok := a.(desktop.App); ok {
		h := fyne.NewMenuItem("esphomeYAMLGenerator By Luckykeeper", func() {})
		menu := fyne.NewMenu("Hello World", h)
		h.Action = func() {
			log.Println("Hi there!")
			h.Label = "esphomeYAMLGenerator By Luckykeeper"
			u, _ := url.Parse("https://github.com/luckykeeper")
			a.OpenURL(u)
			menu.Refresh()
		}
		desk.SetSystemTrayMenu(menu)
	}
}

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}

// 复制文件
func copy(src, dst string, BUFFERSIZE int64) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		log.Println("src is not a regular file.", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	_, err = os.Stat(dst)
	if err == nil {
		log.Println("dst is not a regular file.", src)
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	if err != nil {
		panic(err)
	}

	buf := make([]byte, BUFFERSIZE)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return err
}
