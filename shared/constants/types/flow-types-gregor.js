/* @flow */

// This file is auto-generated by client/protocol/Makefile.

export type int = number
export type int64 = number
export type long = number
export type double = number
export type bytes = any
export type RPCError = {
  code: number,
  desc: string
}

export type Body = bytes

export type Category = string

export type DeviceID = bytes

export type Dismissal = {
  msgIDs: Array<MsgID>;
  ranges: Array<MsgRange>;
}

export type DurationMsec = int64

export type InBandMessage = {
  stateUpdate?: ?StateUpdateMessage;
  stateSync?: ?StateSyncMessage;
}

export type Item = {
  category: Category;
  dtime: TimeOrOffset;
  remindTimes: Array<TimeOrOffset>;
  body: Body;
}

export type ItemAndMetadata = {
  md?: ?Metadata;
  item?: ?Item;
}

export type Message = {
  oobm?: ?OutOfBandMessage;
  ibm?: ?InBandMessage;
}

export type Metadata = {
  uid: UID;
  msgID: MsgID;
  ctime: Time;
  deviceID: DeviceID;
  inBandMsgType: int;
}

export type MsgID = bytes

export type MsgRange = {
  endTime: TimeOrOffset;
  category: Category;
}

export type OutOfBandMessage = {
  uid: UID;
  system: System;
  body: Body;
}

export type Reminder = {
  item: ItemAndMetadata;
  seqno: int;
  remindTime: Time;
}

export type ReminderID = {
  uid: UID;
  msgID: MsgID;
  seqno: int;
}

export type ReminderSet = {
  reminders: Array<Reminder>;
  moreRemindersReady: bool;
}

export type SessionID = string

export type SessionToken = string

export type State = {
  items: Array<ItemAndMetadata>;
}

export type StateSyncMessage = {
  md: Metadata;
}

export type StateUpdateMessage = {
  md: Metadata;
  creation?: ?Item;
  dismissal?: ?Dismissal;
}

export type System = string

export type Time = long

export type TimeOrOffset = {
  time: Time;
  offset: DurationMsec;
}

export type UID = bytes

export type incomingCallMapType = {

}
