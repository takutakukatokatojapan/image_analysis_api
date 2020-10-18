package service

import (
	"fmt"
	"github.com/takutakukatokatojapan/image_analysis_api/infrastructure/appctx"
	"github.com/takutakukatokatojapan/image_analysis_api/infrastructure/logger"
	"io/ioutil"
)

type (
	FileLoadingService interface {
		Load(ctx *appctx.APPCtx) FileLoadingResultContainer
	}
	FileLoadingServiceImpl struct {
	}
	FileLoadingResultContainer struct {
		FileName string
		Data     string
		Error    error
	}
)

func NewFileLoadingServiceImpl() FileLoadingService {
	return &FileLoadingServiceImpl{}
}

func (f FileLoadingServiceImpl) Load(ctx *appctx.APPCtx) FileLoadingResultContainer {
	file, err := ctx.FormFile("file")
	if err != nil {
		return FileLoadingResultContainer{
			FileName: "",
			Data:     "",
			Error:    fmt.Errorf("happen where read parameter: %w", err),
		}
	}
	fileBinary, err := file.Open()
	if err != nil {
		return FileLoadingResultContainer{
			FileName: "",
			Data:     "",
			Error:    fmt.Errorf("happen where open file: %w", err),
		}
	}
	defer func() {
		if err = fileBinary.Close(); err != nil {
			logger.Default.Warn(ctx.XRequestID, fmt.Sprintf("can not close file: %v", err))
		}
	}()
	b, err := ioutil.ReadAll(fileBinary)
	if err != nil {
		return FileLoadingResultContainer{
			FileName: "",
			Data:     "",
			Error:    fmt.Errorf("happen where read file: %w", err),
		}
	}
	return FileLoadingResultContainer{
		FileName: file.Filename,
		Data:     string(b),
		Error:    nil,
	}
}
