package teams

import (
	libkb "github.com/keybase/client/go/libkb"
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	context "golang.org/x/net/context"
)

func LoadUserPlusKeys(ctx context.Context, g *libkb.GlobalContext, id keybase1.TeamID) (*keybase1.TeamPlusAllKeys, error) {
	team, err := GetFromID(ctx, g, id)
	if err != nil {
		return nil, err
	}
	teamPlusAllKeys, err := team.ExportToTeamPlusAllKeys(keybase1.Time(0))
	if err != nil {
		return nil, err
	}
	return teamPlusAllKeys, err
}
