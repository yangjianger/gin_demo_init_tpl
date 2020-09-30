package utils

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

//生成分布式唯一uid ---> 具有有序性 雪花算法

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

// 需传⼊入当前的机器器ID
func Init(machineId uint16) (err error) {
	sonyMachineID = machineId
	t, _ := time.Parse("2006-01-02", "2020-01-01")
	settings := sonyflake.Settings{
		StartTime: t,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}

// GetID 返回⽣生成的id值
func GetID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("snoy flake not inited")

		return
	}
	id, err = sonyFlake.NextID()
	return
}
