package support

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sony/sonyflake"
	"net"
)

var sf *sonyflake.Sonyflake
var IdGeneratorInstance IdGeneratorSupport

type IdGeneratorSupport struct {
}

func init() {
	IdGeneratorInstance = IdGeneratorSupport{}
	//var st sonyflake.Settings
	//st.MachineID = awsutil.AmazonEC2MachineID
	//sf = sonyflake.NewSonyflake(st)
	//if sf == nil {
	//	panic("sonyflake not created")
	//}
}
func localIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	var ip string = "localhost"
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}
	return ip
}
func (igs IdGeneratorSupport) NextId() int64 {
	//currentUsers := support.UserContextInstance.GetCurrentUser()
	//log.Println(currentUsers)
	//var tableName = tpe.(string)
	//
	//return generatorNextId(tableName)
	id, err := sf.NextID()
	if err != nil {
		panic(errors.New("id生成错误:" + err.Error()))
	}
	return int64(id)
}

//func generatorNextId(name string) int64 {
//	//var idGenerator = dao.Use(base.DB).IDGENERATOR
//	//var idGeneratorDao = dao.Use(base.DB).WithContext(context.Background()).IDGENERATOR
//	//idGeneratorList, error := idGeneratorDao.Where(idGenerator.FlagNo.Eq(name)).Find()
//	//if error != nil {
//	//	panic(error)
//	//}
//	//if len(idGeneratorList) > 1 {
//	//	panic(errors.New("id生成器配置错误!模板只能为一个！模板名称:" + name))
//	//}
//	//m.Lock()
//	//m.Unlock()
//	//return 0
//	id, err := sf.NextID()
//	if err != nil {
//
//	}
//	return id.(int64)
//}
