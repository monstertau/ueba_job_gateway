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
	mockG := group.Group("/mock")
	mockG.GET("/profiling", s.getMockProfilingJobs)
	mockG.GET("/behavior", s.getMockBehaviorJob)
	mockG.GET("/rule", s.getMockRuleJob)

	workerG := group.Group("/worker")
	workerG.GET("/profiling", s.getProfilingJobs)
	workerG.GET("/behavior", s.getBehaviorJob)
	workerG.GET("/rule", s.getRuleJob)
}

func (s *JobHandler) getMockProfilingJobs(c *gin.Context) {
	jobs, err := s.JobServices.Profiling.GetMockJobs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("something gone wrong: %v", err))
		return
	}
	c.JSON(http.StatusOK, jobs)
}

func (s *JobHandler) getMockBehaviorJob(c *gin.Context) {
	jobs, err := s.JobServices.FlinkSQL.GetMockBehaviorJobs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("something gone wrong: %v", err))
		return
	}
	c.JSON(http.StatusOK, jobs)
}

func (s *JobHandler) getMockRuleJob(c *gin.Context) {
	jobs, err := s.JobServices.FlinkSQL.GetMockRuleJobs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("something gone wrong: %v", err))
		return
	}
	c.JSON(http.StatusOK, jobs)
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
