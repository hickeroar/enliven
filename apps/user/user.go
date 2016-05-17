package user

//go:generate go-bindata -o templates.go -pkg user templates/...

import (
	"strconv"

	"github.com/hickeroar/enliven"
	"github.com/jinzhu/gorm"
)

// User describes the user database structure.
type User struct {
	gorm.Model

	DisplayName      string
	Age              int
	Login            string `gorm:"type:varchar(100);unique_index;"`
	Email            string `gorm:"type:varchar(100);unique_index;"`
	Password         string `gorm:"type:varchar(100);"`
	VerificationCode string `gorm:"type:varchar(100);unique_index;"`
	Group            Group  `gorm:"not null"`
	Superuser        bool
}

// HasPermission checks if a user has a specific permission
func (u *User) HasPermission(name string) bool {
	if u.Superuser {
		return true
	}

	var groupStack []string
	return u.hasPermission(name, &u.Group, groupStack)
}

// hasPermission recursively looks through a group's inheritance chain to look for a permission
func (u *User) hasPermission(name string, group *Group, groupStack []string) bool {
	// Checking if this group has a permission matching the one we're looking for
	for _, perm := range group.Permisions {
		if name == perm.Name {
			return true
		}
	}

	// If this group inherits from another groups, we check that group for a permission
	if group.Inherits != nil && group.Name != group.Inherits.Name {

		// We're avoiding infinite group loops here by making sure we don't check a group more than once.
		for _, stackGroup := range groupStack {
			if stackGroup == group.Inherits.Name {
				return false
			}
		}
		groupStack = append(groupStack, group.Inherits.Name)

		// Recursively checking the group's inheritance chain.
		return u.hasPermission(name, group.Inherits, groupStack)
	}

	return false
}

// Group describes the user group database structure
type Group struct {
	gorm.Model

	Name       string `gorm:"not null;unique;"`
	Inherits   *Group
	Permisions []Permission
}

// Permission describes a permission that can be linked to many groups
type Permission struct {
	gorm.Model

	Name string `gorm:"not_null;unique;"`
}

// NewApp generates and returns an instance of the app
func NewApp() *App {
	return &App{}
}

// App handles adding a route handler for static assets
type App struct{}

// Initialize sets up our app to handle embedded static asset requests
func (ua *App) Initialize(ev *enliven.Enliven) {
	var config = enliven.Config{
		"user.login.route":    "/user/login/",
		"user.logout.route":   "/user/logout/",
		"user.register.route": "/user/register/",
		"user.verify.route":   "/user/verify/",
		"user.password.route": "/user/password/",

		"user.login.template":    "",
		"user.logout.template":   "",
		"user.register.template": "",
		"user.verify.template":   "",
		"user.password.template": "",
	}

	config = enliven.MergeConfig(config, ev.GetConfig())

	// Migrating the user tables
	ev.GetDatabase().AutoMigrate(&User{}, &Group{}, &Permission{})

	// Routing setup
	ev.AddRoute(config["user.login.route"], LoginGetHandler)

	// Handles the setup of context variables to support user session management
	ev.AddMiddlewareFunc(SessionMiddleware)
}

// GetName returns the app's name
func (ua *App) GetName() string {
	return "user"
}

// GetUser returns an instance of the User model
func GetUser(ctx *enliven.Context) *User {
	// If they're not logged in, return 0
	if ctx.Items["UserLoggedIn"] == "0" {
		return nil
	}

	// The user may be cached from an earlier lookup.
	if user, ok := ctx.Storage["User"]; ok {
		return user.(*User)
	}

	var user User
	dbUserID, _ := strconv.Atoi(ctx.Items["UserID"])
	ctx.Enliven.GetDatabase().First(&user, dbUserID)

	// Caching the user lookup for later.
	ctx.Storage["User"] = &user

	return &user
}
