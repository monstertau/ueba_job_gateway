package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"job-gateway/service"
	"net/http"
)

type JobHandler struct {
	JobServices *service.Services
}

func NewJobHandler(services *service.Services) *JobHandler {
	return &JobHandler{
		JobServices: services,
	}
}

func (s *JobHandler) MakeHandler(g *gin.RouterGroup) {
	group := g.Group("/jobs")
	group.GET("/profiling", s.getProfilingJobs)
	group.GET("/behavior", s.getBehaviorJob)
	group.GET("/rule", s.getRuleJob)
}

func (s *JobHandler) getProfilingJobs(c *gin.Context) {
	jobs, err := s.JobServices.Profiling.GetJobs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("something gone wrong: %v", err))
		return
	}
	c.JSON(http.StatusOK, jobs)
}

func (s *JobHandler) getBehaviorJob(c *gin.Context) {
	jobs, err := s.JobServices.FlinkSQL.GetBehaviorJobs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("something gone wrong: %v", err))
		return
	}
	c.JSON(http.StatusOK, jobs)
}

func (s *JobHandler) getRuleJob(c *gin.Context) {
	jobs, err := s.JobServices.FlinkSQL.GetRuleJobs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("something gone wrong: %v", err))
		return
	}
	c.JSON(http.StatusOK, jobs)
}
