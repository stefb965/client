package teams

import (
	"golang.org/x/net/context"
	"testing"

	"github.com/keybase/client/go/kbtest"
	"github.com/keybase/client/go/libkb"
	"github.com/keybase/client/go/protocol/keybase1"
)

func TestTeamPlusAllKeysExim(t *testing.T) {
	tc := libkb.SetupTest(t, "TestTeamPlusAllKeysExim", 1)
	tc.Tp.UpgradePerUserKey = true
	_, err := kbtest.CreateAndSignupFakeUser("team", tc.G)
	if err != nil {
		t.Fatal(err)
	}
	defer tc.Cleanup()

	name := createTeam(tc)
	team, err := Get(context.TODO(), tc.G, name)
	if err != nil {
		t.Fatal(err)
	}

	teamFromID, err := GetFromID(context.TODO(), tc.G, team.Chain.GetID())

	exported, err := teamFromID.ExportToTeamPlusAllKeys(keybase1.Time(0))
	if err != nil {
		t.Errorf("Error during export: %s", err)
	}
	if exported.Name != teamFromID.Name {
		t.Errorf("Got name %s, expected %s", exported.Name, teamFromID.Name)
	}
	if exported.Id != teamFromID.Chain.GetID() {
		t.Errorf("Got id %s, expected %s", exported.Id, teamFromID.Chain.GetID())
	}
	if len(exported.PerTeamKeys) != len(teamFromID.Chain.inner.PerTeamKeys) {
		t.Errorf("Got %s perTeamKeys, expected %s", len(exported.PerTeamKeys), len(teamFromID.Chain.inner.PerTeamKeys))
	}
}
