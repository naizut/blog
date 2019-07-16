package router

import (
    "github.com/gin-gonic/gin"
    // "github.com/gin-contrib/cors"
    
    "blog/pkg/setting"
    "blog/routers/api"
    "blog/routers/api/v1"
    "blog/middleware/jwt"
)

func InitRouter() *gin.Engine {
    //New()+Logger+Recovery = Default()
    r := gin.New()

    r.Use(gin.Logger())

    r.Use(gin.Recovery())

    // r.Use(cors.Default())

    gin.SetMode(setting.RunMode)

    r.POST("/auth", api.GetAuth)

    apiv1 := r.Group("/api/v1")
    
    //无需鉴权
    apiv1.GET("/tags", v1.GetTags)

    apiv1.GET("/articles", v1.GetArticles)
    apiv1.GET("/articles/:id", v1.GetArticle)

    //代码块中的地址需要鉴权
    apiv1.Use(jwt.JWT())
    {
        //标签增删改
        apiv1.POST("/tags", v1.AddTags)
        apiv1.DELETE("/tags/:id", v1.DeleteTags)
        apiv1.PUT("/tags/:id", v1.EditTags)

        //文章增删改
        apiv1.POST("/articles", v1.AddArticle)
        apiv1.DELETE("/articles/:id", v1.DeleteArticle)
        apiv1.PUT("/articles/:id", v1.EditArticle)
    }
    return r
}
