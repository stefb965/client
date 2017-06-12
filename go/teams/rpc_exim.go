// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

// Export-Import for RPC for Teams

package teams

import (
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
)

func (t *Team) ExportToTeamPlusAllKeys(idTime keybase1.Time) (keybase1.TeamPlusAllKeys, error) {
	var teamPlusAllKeys keybase1.TeamPlusAllKeys
	perTeamKeys := t.Chain.inner.PerTeamKeys

	members, err := t.Members()
	if err != nil {
		return teamPlusAllKeys, err
	}

	writers := make([]keybase1.UserVersion, 0)
	for _, writerString := range members.Writers {
		writer, err := ParseUserVersion(writerString)
		if err != nil {
			return teamPlusAllKeys, err
		}
		writers = append(writers, writer)
	}

	writersSet := make(map[keybase1.UserVersion]bool, 0)
	for _, writer := range writers {
		writersSet[writer] = true
	}

	onlyReaders := make([]keybase1.UserVersion, 0)
	for _, readerString := range members.Readers {
		reader, err := ParseUserVersion(readerString)
		if err != nil {
			return teamPlusAllKeys, err
		}
		_, ok := writersSet[reader]
		if !ok {
			onlyReaders = append(onlyReaders, reader)
		}
	}

	teamPlusAllKeys = keybase1.TeamPlusAllKeys{
		Id:          t.Chain.GetID(),
		Name:        t.Chain.GetName(),
		PerTeamKeys: perTeamKeys,
		Writers:     writers,
		OnlyReaders: onlyReaders,
	}

	return teamPlusAllKeys, nil
}
