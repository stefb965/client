package teams

import (
	"encoding/json"
	"strings"

	"golang.org/x/net/context"

	"github.com/keybase/client/go/libkb"
	"github.com/keybase/client/go/protocol/keybase1"
)

func Get(ctx context.Context, g *libkb.GlobalContext, name string) (*Team, error) {
	f := newFinder(g)
	return f.find(ctx, g, name)
}

func GetFromID(ctx context.Context, g *libkb.GlobalContext, id keybase1.TeamID) (*Team, error) {
	f := newFinder(g)
	return f.findFromID(ctx, g, id)
}

type finder struct {
	libkb.Contextified
}

func newFinder(g *libkb.GlobalContext) *finder {
	return &finder{
		Contextified: libkb.NewContextified(g),
	}
}

func (f *finder) findFromRaw(ctx context.Context, g *libkb.GlobalContext, raw *rawTeam) (*Team, error) {
	team := NewTeam(g, strings.Join(raw.Name.Parts, "."))
	team.Box = raw.Box
	team.ReaderKeyMasks = raw.ReaderKeyMasks

	links, err := f.chainLinks(ctx, raw)
	if err != nil {
		return nil, err
	}

	player, err := f.newPlayer(ctx, links)
	if err != nil {
		return nil, err
	}

	state, err := player.GetState()
	if err != nil {
		return nil, err
	}

	team.Chain = &state

	return team, nil
}

func (f *finder) find(ctx context.Context, g *libkb.GlobalContext, name string) (*Team, error) {
	raw, err := f.rawTeam(ctx, name)
	if err != nil {
		return nil, err
	}
	return f.findFromRaw(ctx, g, raw)
}

func (f *finder) findFromID(ctx context.Context, g *libkb.GlobalContext, id keybase1.TeamID) (*Team, error) {
	raw, err := f.rawTeamFromID(ctx, id)
	if err != nil {
		return nil, err
	}
	return f.findFromRaw(ctx, g, raw)
}

func (f *finder) rawTeamFromArgs(ctx context.Context, args libkb.HTTPArgs) (*rawTeam, error) {
	arg := libkb.NewRetryAPIArg("team/get")
	arg.NetContext = ctx
	arg.SessionType = libkb.APISessionTypeREQUIRED
	arg.Args = args
	var rt rawTeam
	if err := f.G().API.GetDecode(arg, &rt); err != nil {
		return nil, err
	}
	return &rt, nil
}

func (f *finder) rawTeam(ctx context.Context, name string) (*rawTeam, error) {
	args := libkb.HTTPArgs{
		"name": libkb.S{Val: name},
	}
	return f.rawTeamFromArgs(ctx, args)
}

func (f *finder) rawTeamFromID(ctx context.Context, id keybase1.TeamID) (*rawTeam, error) {
	args := libkb.HTTPArgs{
		"id": libkb.S{Val: string(id)},
	}
	return f.rawTeamFromArgs(ctx, args)
}

func (f *finder) chainLinks(ctx context.Context, rawTeam *rawTeam) ([]SCChainLink, error) {
	var links []SCChainLink
	for _, raw := range rawTeam.Chain {
		link, err := ParseTeamChainLink(string(raw))
		if err != nil {
			return nil, err
		}
		links = append(links, link)
	}
	return links, nil
}

func (f *finder) newPlayer(ctx context.Context, links []SCChainLink) (*TeamSigChainPlayer, error) {
	uv, err := loadUserVersionByUID(ctx, f.G(), f.G().Env.GetUID())
	if err != nil {
		return nil, err
	}
	player := NewTeamSigChainPlayer(f.G(), uv, false)
	if err := player.AddChainLinks(ctx, links); err != nil {
		return nil, err
	}
	return player, nil
}

type rawTeam struct {
	ID             keybase1.TeamID          `json:"id"`
	Name           keybase1.TeamNameParts   `json:"name"`
	Status         libkb.AppStatus          `json:"status"`
	Chain          []json.RawMessage        `json:"chain"`
	Box            TeamBox                  `json:"box"`
	ReaderKeyMasks []keybase1.ReaderKeyMask `json:"reader_key_masks"`
}

func (r *rawTeam) GetAppStatus() *libkb.AppStatus {
	return &r.Status
}
