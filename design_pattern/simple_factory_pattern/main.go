package main

import (
	"fmt"
	mario "github.com/mats9693/study/design_pattern/simple_factory_pattern/operation"
	"log"
)

func main() {
	facIns := &mario.OperationFactory{}

	log.Println("请输入需要计算的表达式（形如‘1 + 1’，符号前后应有空格）：")
	if n, err := fmt.Scanln(&facIns.NumberA, &facIns.Operate, &facIns.NumberB); n != 3 || err != nil {
		log.Printf("表达式解析错误！\n成功解析数：%d，错误：%v\n使用默认参数：1 + 1\n", n, err)
		facIns.NumberA = 1
		facIns.Operate = "+"
		facIns.NumberA = 1
	}

	result, err := facIns.NewOperation().CalculateResult()
	if err != nil {
		log.Fatalln("计算错误：", err)
	}

	log.Println("结果：", result)

	return
}
