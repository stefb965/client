package teams

import (
	"golang.org/x/net/context"
	"testing"

	"github.com/keybase/client/go/kbtest"
	"github.com/keybase/client/go/libkb"
	"github.com/keybase/client/go/protocol/keybase1"
)

func TestTeamPlusAllKeysExim(t *testing.T) {
	tc := libkb.SetupTest(t, "team", 1)
	tc.Tp.UpgradePerUserKey = true
	defer tc.Cleanup()
	kbtest.CreateAndSignupFakeUser("team", tc.G)
	name := createTeam(tc)
	team, err := Get(context.TODO(), tc.G, name)
	var exported = team.ExportToTeamPlusAllKeys(keybase1.Time(0))
	if exported.Name != name {
		t.Errorf("Got name %s, expected %s", exported.Name, name)
	}
	// check rest of PlusAllKeys
}
