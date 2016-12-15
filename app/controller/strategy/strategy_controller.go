package strategy

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"io/ioutil"

	"github.com/gin-gonic/gin"
	h "github.com/masato25/owl_backend/app/helper"
	f "github.com/masato25/owl_backend/app/model/falcon_portal"
)

func GetStrategy(c *gin.Context) {
	var strategys []f.Strategy
	tidtmp := c.DefaultQuery("tid", "")
	if tidtmp == "" {
		h.JSONR(c, badstatus, "tid is missing")
		return
	}
	tid, err := strconv.Atoi(tidtmp)
	if err != nil {
		h.JSONR(c, badstatus, err)
		return
	}
	dt := db.Falcon.Where("tpl_id = ?", tid).Find(&strategys)
	if dt.Error != nil {
		h.JSONR(c, badstatus, dt.Error)
		return
	}
	h.JSONR(c, strategys)
	return
}

type APICreateStrategyInput struct {
	Metric     string `json:"metric" binding:"required"`
	Tags       string `json:"tags"`
	MaxStep    int    `json:"max_step" binding:"required"`
	Priority   int    `json:"priority" binding:"exists"`
	Func       string `json:"func" binding:"required"`
	Op         string `json:"op" binding:"required"`
	RightValue string `json:"right_value" binding:"required"`
	Note       string `json:"note"`
	RunBegin   string `json:"run_begin"`
	RunEnd     string `json:"run_end"`
	TplId      uint   `json:"tpl_id" binding:"required"`
}

func (this APICreateStrategyInput) CheckFormat() (err error) {
	validOp := regexp.MustCompile(`^(>|=|<|!)(=)?$`)
	validRightValue := regexp.MustCompile(`^\d+$`)
	validTime := regexp.MustCompile(`^\d{2}:\d{2}$`)
	switch {
	case !validOp.MatchString(this.Op):
		err = errors.New("op's formating is not vaild")
	case !validRightValue.MatchString(this.RightValue):
		err = errors.New("right_value's formating is not vaild")
	case !validTime.MatchString(this.RunBegin) && this.RunBegin != "":
		err = errors.New("run_begin's formating is not vaild, please refer ex. 00:00")
	case !validTime.MatchString(this.RunEnd) && this.RunEnd != "":
		err = errors.New("run_end's formating is not vaild, please refer ex. 24:00")
	}
	return
}

func CreateStrategy(c *gin.Context) {
	var inputs APICreateStrategyInput
	if err := c.Bind(&inputs); err != nil {
		h.JSONR(c, badstatus, err)
		return
	}
	if err := inputs.CheckFormat(); err != nil {
		h.JSONR(c, badstatus, err)
		return
	}
	strategy := f.Strategy{
		Metric:     inputs.Metric,
		Tags:       inputs.Tags,
		MaxStep:    inputs.MaxStep,
		Priority:   inputs.Priority,
		Func:       inputs.Func,
		Op:         inputs.Op,
		RightValue: inputs.RightValue,
		Note:       inputs.Note,
		RunBegin:   inputs.RunBegin,
		RunEnd:     inputs.RunEnd,
		TplId:      inputs.TplId,
	}
	dt := db.Falcon.Save(&strategy)
	if dt.Error != nil {
		h.JSONR(c, expecstatus, dt.Error)
		return
	}
	h.JSONR(c, "stragtegy created")
	return
}

func MetricQuery(c *gin.Context) {
	data, err := ioutil.ReadFile("data/metric")
	if err != nil {
		h.JSONR(c, badstatus, err)
		return
	}
	metrics := strings.Split(string(data), "\n")
	h.JSONR(c, metrics)
	return
}