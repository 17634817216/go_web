package cmd

import (
	"github.com/gin-gonic/gin"
	"go_admin_api/internal/controller/system"
)

func System(group *gin.RouterGroup) {
	permissionController := system.NewPermissionController()
	rolesController := system.NewRoleController()
	OrganizationController := system.NewOrganizationController()
	UserServiceController := system.NewUserController()
	permissions := group.Group("/permissions")
	{
		permissions.POST("", permissionController.Create)
		permissions.PUT("/:permission_id", permissionController.Update)
		permissions.DELETE("/:id", permissionController.Delete)
		permissions.GET("/tree/:is_backstage", permissionController.GetTree)
		//permissions.GET("/list", permissionController.GetList)
	}

	roles := group.Group("/roles")
	{
		roles.POST("", rolesController.Create)
		roles.GET("", rolesController.GetRoles)
		roles.PUT("/:role_id", rolesController.UpadteRole)
		roles.DELETE("/:role_id", rolesController.DeleteRoles)
		roles.GET("/rolemenu/:role_id", rolesController.GetRolesPermissions)
	}
	organs := group.Group("/organ")
	{
		organs.POST("", OrganizationController.Create)
		organs.PUT("/:org_id", OrganizationController.Update)
		organs.DELETE("/:org_id", OrganizationController.DeleteOrg)
		organs.GET("/tree", OrganizationController.GetOrgTree)
		//permissions.GET("/list", permissionController.GetList)
	}
	users := group.Group("/users")
	{
		users.GET("", UserServiceController.GetlistPage)
		users.POST("", UserServiceController.Create)
		users.PUT(":user_id", UserServiceController.Update)
		users.DELETE(":user_id", UserServiceController.Delete)
	}
}
