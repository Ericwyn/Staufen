package cmd

import (
	"fmt"
	"github.com/Ericwyn/Staufen/model/authmodel"
	"github.com/Ericwyn/Staufen/model/bucketmodel"
	"github.com/Ericwyn/Staufen/repo"
	"github.com/Ericwyn/Staufen/repo/bucketrepo"
	"unicode"
)

// 命令行操作

func StartCli() {
	fmt.Println("欢迎来到 Staufen Cli")
	cliHome()
}

func cliHome() {
	for {
		fmt.Println()
		var input = printInputSelect([]inputSelect{
			{
				Key:   "0",
				Value: "list all bucket",
			},
			{
				Key:   "1",
				Value: "create bucket",
			},
			{
				Key:   "2",
				Value: "update bucket",
			},
			{
				Key:   "3",
				Value: "delete bucket",
			},
			{
				Key:   "4",
				Value: "get upload token",
			},
		})
		switch input {
		case "0":
			//cliShowDriver()
			listAllBucket()
			break
		case "1":
			createBucket()
			break
			//cliUploadFile()
		case "2":
			//cliUploadFile()
		case "3":
			deleteBucket()
			break
		case "4":
			getUploadToken()
			break
		case "x":
			return
		default:
			fmt.Println("命令输入错误")
			fmt.Println()
		}
	}
}

func listAllBucket() {
	buckets, err := bucketrepo.ListBucket()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		printOneBucketHeader()
		for _, bucket := range buckets {
			printOneBucket(bucket)
		}
	}
}

func printOneBucketHeader() {
	fmt.Println(
		spaceStr("| ID", 8), "|",
		spaceStr("名称", 20), "|",
		spaceStr("bucketUuid", 40), " |",
		spaceStr("public", 10), " |",
		spaceStr("ApiToken", 40), " |",
	)
}
func printOneBucket(bucket repo.Bucket) {
	fmt.Println(
		spaceStr("| "+fmt.Sprint(bucket.Id), 8), "|",
		spaceStr(bucket.Name, 20), "|",
		spaceStr(bucket.BucketId, 40), " |",
		spaceStr(fmt.Sprint(bucket.Public), 10), " |",
		spaceStr(bucket.ApiToken, 40), " |",
	)
}

func createBucket() {
	fmt.Println("请输入:名称 public(y/n) 压缩(y/n), 如 \"公开bucket y n\"")
	fmt.Print(">>:")

	var bucketName string
	var publicFlag string
	var compressFlag string

	fmt.Scanln(&bucketName, &publicFlag, &compressFlag)
	if (publicFlag != "y" && publicFlag != "n") ||
		(compressFlag != "y" && compressFlag != "n") {
		fmt.Println("输入有误")
		return
	}
	fmt.Println("录入信息如下")
	fmt.Println("bucket名称:\t" + bucketName)
	fmt.Println("public:\t\t" + publicFlag)
	fmt.Println("compress:\t\t" + compressFlag)

	fmt.Println("确认信息(y/n)")
	fmt.Print(">>:")

	var ensureFlag string
	fmt.Scanln(&ensureFlag)
	if ensureFlag != "y" {
		return
	}

	bucket, err := bucketmodel.CreateBucket(bucketName, publicFlag == "y", compressFlag == "y")
	if err != nil {
		fmt.Println(err)
	} else {
		printOneBucketHeader()
		printOneBucket(*bucket)
	}
}

func deleteBucket() {
	fmt.Println("请输入 bucket 的 Id")
	fmt.Print(">>:")
	var bucketId int64
	fmt.Scanf("%d", &bucketId)
	bucket := bucketrepo.GetBucketById(bucketId)
	if bucket == nil {
		fmt.Println("输入错误，无法找到 bucket")
		return
	}
	fmt.Println("是否删除如下 bucket ?")
	printOneBucket(*bucket)

	fmt.Print(">>:")
	var ensureFlag string
	fmt.Scanln(&ensureFlag)
	if ensureFlag == "y" {
		err := bucketmodel.DeleteBucket(*bucket)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getUploadToken() {
	fmt.Println("请输入 bucket 的 Id")
	fmt.Print(">>:")
	var bucketId int64
	fmt.Scanf("%d", &bucketId)
	bucket := bucketrepo.GetBucketById(bucketId)
	if bucket == nil {
		fmt.Println("输入错误，无法找到 bucket")
		return
	}
	fmt.Println(authmodel.GetUploadToken(*bucket))
}

type inputSelect struct {
	Key   string
	Value string
}

func printInputSelect(selects []inputSelect) string {
	fmt.Println("======================================================")
	for _, slt := range selects {
		fmt.Println(slt.Key, ".", slt.Value)
	}
	fmt.Println("======================================================")
	fmt.Print(">>:")
	var input string
	fmt.Scanln(&input)
	return input
}

func spaceStr(msg string, width int) string {
	// 计算字符长度，汉字需要按 2 来计算
	msgLen := 0
	for _, r := range msg {
		if unicode.Is(unicode.Han, r) {
			msgLen += 2
		} else {
			msgLen += 1
		}
	}

	if msgLen < width {
		for i := 0; i < width-msgLen; i++ {
			msg += " "
		}
	}
	return msg
}
