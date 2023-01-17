package cobra

import (
	"beide/framework"
	"beide/framework/contract"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

// SetContainer 设置服务容器
func (c *Command) SetContainer(container framework.Container) {
	c.container = container
}

// GetContainer 获取容器
func (c *Command) GetContainer() framework.Container {
	return c.Root().container
}

// CronSpec 保存Cron命令的信息，用于展示
type CronSpec struct {
	Type        string
	Cmd         *Command
	Spec        string
	ServiceName string
}

func (c *Command) SetParantNull() {
	c.parent = nil
}

// AddCronCommand 是用来创建一个Cron任务的
func (c *Command) AddCronCommand(spec string, cmd *Command) {
	// cron结构是挂载在根Command上的
	root := c.Root()
	if root.Cron == nil {
		// 初始化cron
		root.Cron = cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
		root.CronSpecs = []CronSpec{}
	}
	// 增加说明信息
	root.CronSpecs = append(root.CronSpecs, CronSpec{
		Type: "normal-cron",
		Cmd:  cmd,
		Spec: spec,
	})

	// 制作一个rootCommand
	var cronCmd Command
	ctx := root.Context()
	cronCmd = *cmd
	cronCmd.args = []string{}
	cronCmd.SetParantNull()
	cronCmd.SetContainer(root.GetContainer())

	// 增加调用函数
	root.Cron.AddFunc(spec, func() {
		// 如果后续的command出现panic，这里要捕获
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()

		err := cronCmd.ExecuteContext(ctx)
		if err != nil {
			// 打印出err信息
			log.Println(err)
		}
	})
}

// AddDistributedCronCommand 实现一个分布式定时器
// serviceName 这个服务的唯一名字，不允许带有空格
// spec 具体的执行时间
// cmd 具体的执行命令
// holdTime 表示如果我选择上了，这次选择持续的时间，也就是锁释放的时间
func (c *Command) AddDistributedCronCommand(serviceName string, spec string, cmd *Command, holdTime time.Duration) {
	root := c.Root()

	// 初始化cron
	if root.Cron == nil {
		root.Cron = cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
		root.CronSpecs = []CronSpec{}
	}

	// cron命令的注释，这里注意Type为distributed-cron，ServiceName需要填写
	root.CronSpecs = append(root.CronSpecs, CronSpec{
		Type:        "distributed-cron",
		Cmd:         cmd,
		Spec:        spec,
		ServiceName: serviceName,
	})

	appService := root.GetContainer().MustMake(contract.AppKey).(contract.App)
	log.Println("appService:", appService)
	distributeServce := root.GetContainer().MustMake(contract.DistributedKey).(contract.Distributed)
	log.Println("distributeServce:", distributeServce)
	log.Println("contract.DistributedKey:", contract.DistributedKey)
	appID := appService.AppID()

	// 复制要执行的command为cronCmd，并且设置为rootCmd
	var cronCmd Command
	ctx := root.Context()
	cronCmd = *cmd
	cronCmd.args = []string{}
	cronCmd.SetParantNull()

	// cron增加匿名函数
	root.Cron.AddFunc(spec, func() {
		// 防止panic
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()

		// 节点进行选举，返回选举结果
		selectedAppID, err := distributeServce.Select(serviceName, appID, holdTime)
		if err != nil {
			return
		}

		// 如果自己没有被选择到，直接返回
		if selectedAppID != appID {
			return
		}

		// 如果自己被选择到了，执行这个定时任务
		err = cronCmd.ExecuteContext(ctx)
		if err != nil {
			log.Println(err)
		}
	})
}
