package sqlproto

import (
	"errors"
	"log"
	"strings"

	"github.com/jinzhu/inflection"

	"github.com/tx7do/kratos-cli/sql-proto/internal/render"
)

type ProtoField render.ProtoField
type ProtoFieldArray []render.ProtoField

func WriteServiceProto(
	outputPath string,
	serviceType string,
	targetModuleName, sourceModuleName, moduleVersion string,
	tableName string,
	tableComment string,
	protoFields ProtoFieldArray,
) error {
	switch strings.TrimSpace(strings.ToLower(serviceType)) {
	case "grpc":
		data := render.GrpcProtoTemplateData{
			Module:  targetModuleName,
			Version: moduleVersion,

			Name:    inflection.Singular(tableName),
			Comment: render.RemoveTableCommentSuffix(tableComment),
			Fields:  render.ProtoFieldArray(protoFields),
		}
		return render.WriteGrpcServiceProto(outputPath, data)

	case "rest":
		data := render.RestProtoTemplateData{
			SourceModule: sourceModuleName,
			TargetModule: targetModuleName,
			Version:      moduleVersion,

			Name:    inflection.Singular(tableName),
			Comment: render.RemoveTableCommentSuffix(tableComment),
		}
		return render.WriteRestServiceProto(outputPath, data)

	default:
		return errors.New("sqlproto: unsupported service type: " + serviceType)
	}
}

func WriteServicesProto(
	outputPath string,
	serviceType string,
	targetModuleName, sourceModuleName, moduleVersion string,
	tables TableDataArray,
) error {
	var protoFields ProtoFieldArray

	for i := 0; i < len(tables); i++ {
		table := tables[i]

		protoFields = make(ProtoFieldArray, 0, len(table.Fields))
		for n := 0; n < len(table.Fields); n++ {
			field := table.Fields[n]
			protoFields = append(protoFields, render.ProtoField{
				Number:  n + 1,
				Name:    field.Name,
				Comment: field.Comment,
				Type:    field.Type,
			})
		}

		if err := WriteServiceProto(
			outputPath,
			serviceType,
			targetModuleName, sourceModuleName, moduleVersion,
			table.Name, table.Comment,
			protoFields,
		); err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}
