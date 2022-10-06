package support

var IdGeneratorInstance IdGeneratorSupport

type IdGeneratorSupport struct {
}

func init() {

}
func (igs IdGeneratorSupport) NextId(tpe interface{}) int64 {
	return 0
}
