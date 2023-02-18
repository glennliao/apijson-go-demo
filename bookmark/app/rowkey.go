package app

import (
	"context"
	"github.com/glennliao/apijson-go/config"
	"github.com/glennliao/apijson-go/consts"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/yitter/idgenerator-go/idgen"
	"strconv"
)

func init() {

	var options = idgen.NewIdGeneratorOptions(1)
	options.WorkerIdBitLength = 6
	options.BaseTime = 1591545600000 // 2020-06-08
	idgen.SetIdGenerator(options)

	config.RowKeyGenFunc("idgen", func(ctx context.Context, genParam g.Map, table string, data g.Map) (g.Map, error) {
		return g.Map{
			consts.RowKey: strconv.FormatInt(idgen.NextId(), 10),
		}, nil
	})
}
