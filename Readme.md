go test ./... -cover
or
go test ./... -coverprofile=cover.out && go tool cover -html=cover.out

mockery --dir=features/target --name=TargetDataInterface --filename=TargetData.go --structname=TargetData
