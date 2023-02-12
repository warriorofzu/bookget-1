package util

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func RunCommand(ctx context.Context, text string) error {
	fmt.Println(text)
	var cmd *exec.Cmd
	if string(os.PathSeparator) == "\\" {
		cmd = exec.CommandContext(ctx, "cmd", "/c", text)
	} else {
		cmd = exec.CommandContext(ctx, "bash", "-c", text)
	}
	//捕获标准输出
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		return err
	}
	// 执行命令cmd.CombinedOutput(),且捕获输出
	//output, err = cmd.CombinedOutput()
	if err = cmd.Start(); err != nil {
		return err
	}
	readout := bufio.NewReader(stdout)
	GetOutput(readout)
	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func GetOutput(reader *bufio.Reader) {
	var sumOutput string //统计屏幕的全部输出内容
	outputBytes := make([]byte, 200)
	for {
		n, err := reader.Read(outputBytes) //获取屏幕的实时输出(并不是按照回车分割，所以要结合sumOutput)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			sumOutput += err.Error()
		}
		output := string(outputBytes[:n])
		fmt.Print(output) //输出屏幕内容
		sumOutput += output
	}
	return
}
