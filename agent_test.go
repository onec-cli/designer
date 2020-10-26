package designer

import (
	"context"
	sshclient "github.com/v8platform/agent"
	"github.com/v8platform/designer/tests"
	"github.com/v8platform/runner"

	"github.com/stretchr/testify/suite"
	"log"
	"os"
	"testing"
)

type AgentTestSuite struct {
	tests.TestSuite
}

func TestAgent(t *testing.T) {
	suite.Run(t, new(AgentTestSuite))
}

func (a *AgentTestSuite) TestStartAgent() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//
	////t := a.T()
	//
	//go func() {
	//	defer func() error {
	//		if err := recover(); err != nil {
	//			return errors.New(fmt.Sprintf("v8 run agent err:%s", err))
	//		}
	//		return nil
	//	}()
	//
	process, err := runner.Background(ctx, tests.NewFileIB(a.TempIB), AgentModeOptions{
		Visible:        true,
		SSHHostKeyAuto: true,
		BaseDir:        "./"},
	)
	//
	//}()
	if err != nil {
		log.Fatal(err)
	}

	<-process.Ready()

	err = waitAgent(ctx, "localhost:1543")
	if err != nil {
		log.Fatal(err)
	}

	//t.R().NoError(err, errors.GetErrorContext(err))

	//err := ssh.New("localhost").
	//	WithUser("admin").
	//	WithPort("1543").    // Default is 22
	//	RunCommand("\n common shutdown\n", ssh.CommandOptions{
	//		Stdout: os.Stdout,
	//		Stderr: os.Stderr,
	//		Pipein:  os.Pipein,
	//	})
	//if err != nil {
	//	log.Fatal(err)
	//}

	logger := log.New(os.Stdout, "", log.LstdFlags)

	client, err := sshclient.NewAgentClient("", "", "localhost:1543")

	if err != nil {
		logger.Fatalf("create agent client %v", err)
	}

	//err = client.CopyDirTo("./", "./src")
	//if err != nil {
	//	logger.Fatalf("copy file %v", err)
	//}
	//err = client.CopyFileFrom("./src/agent.go", "./tmp/")
	//if err != nil {
	//	logger.Fatalf("copy file %v", err)
	//}
	//logger.Printf("copy dir")

	//err = client.CopyDirFrom("./src", "./tmp/")
	//if err != nil {
	//	logger.Fatalf("copy file %v", err)
	//}
	//err = client.Disconnect()
	//if err != nil {
	//	logger.Fatal(err)
	//}
	//
	err = client.Connect()
	if err != nil {
		logger.Fatal(err)
	}

	//err = client.LoadCfg("./1Cv8.cf")
	//if err != nil {
	//	logger.Fatal(err)
	//}
	//
	//err = client.DumpCfgToFiles("./src", true)
	//if err != nil {
	//	logger.Fatal(err)
	//}

	err = client.Disconnect()
	if err != nil {
		logger.Fatal(err)
	}

	err = client.Shutdown()
	if err != nil {
		logger.Fatal(err)
	}

	//session.WriteChannel(sshclient.CONFIG_COMMAND)
	//str := session.ReadChannelTiming(10)
	//session.WriteChannel("options list")
	//str += session.ReadChannelRegExp(1000, re)
	////logger.Println(str)
	//
	//session.WriteChannel("common connect-ib")
	////time.Sleep(time.Second * 1)
	//session.WriteChannel("common disconnect-ib")
	//str += session.ReadChannelRegExp(1000, re)
	////logger.Println(str)
	////str = session.ReadChannelTiming(10)
	////logger.Println(str)
	//
	//session.WriteChannel("common connect-ib")
	//str += session.ReadChannelRegExp(1000, re)
	////logger.Println(str)
	//
	//session.WriteChannel("common shutdown")
	//str += session.ReadChannelExpect(1000, "\r\n]")
	//logger.Println(str)
	//
}
