package lib

import (
	"github.com/diamondburned/arikawa/v3/discord"
)

// Userish is an interface that captures the common methods you may want to call
// on either a discord.Member or discord.User, including a display name.
type Userish interface {
	ID() discord.UserID
	Username() string
	DisplayName() string // The added method for getting the display name
}

// memberUser adapts a discord.Member to the Userish interface.
type memberUser struct {
	*discord.Member
}

func (mu memberUser) ID() discord.UserID {
	return mu.User.ID
}

func (mu memberUser) Username() string {
	return mu.User.Username
}

func (mu memberUser) DisplayName() string {
	// If Nick is set, return it as the display name, otherwise return Username
	if mu.Member.Nick != "" {
		return mu.Member.Nick
	}
	return mu.User.Username
}

// directUser adapts a discord.User to the Userish interface.
type directUser struct {
	*discord.User
}

func (du directUser) ID() discord.UserID {
	return du.User.ID
}

func (du directUser) Username() string {
	return du.User.Username
}

func (du directUser) DisplayName() string {
	// For a direct user, the display name is just the username since no nickname is available.
	return du.User.Username
}

// GetUserObject takes an interaction event and returns a Userish, which may be
// either a discord.Member or a discord.User, but exposes it through a consistent interface.
func GetUserObject(event discord.InteractionEvent) Userish {
	if event.Member != nil {
		return memberUser{event.Member}
	} else {
		return directUser{event.User}
	}
}