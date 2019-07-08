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
    
    //无需鉴权的地址，独立放在鉴权方法的上面
    //apiv1.GET("/tags", v1.GetTags)
        //新建文章
        apiv1.POST("/articles", v1.AddArticle)
    //代码块中的地址需要鉴权
    apiv1.Use(jwt.JWT())
    {
        //获取标签列表
        apiv1.GET("/tags", v1.GetTags)
        //新建标签
        apiv1.POST("/tags", v1.AddTags)
        //更新指定标签
        apiv1.PUT("/tags/:id", v1.EditTags)
        //删除指定标签
        apiv1.DELETE("/tags/:id", v1.DeleteTags)

        //获取文章列表
        apiv1.GET("/articles", v1.GetArticles)
        //获取指定文章
        apiv1.GET("/articles/:id", v1.GetArticle)

        //更新指定文章
        apiv1.PUT("/articles/:id", v1.EditArticle)
        //删除指定文章
        apiv1.DELETE("/articles/:id", v1.DeleteArticle)
    }
    return r
}
