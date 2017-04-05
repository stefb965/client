// @flow
import * as Constants from '../../constants/chat'
import Conversation from './index'
import {connect} from 'react-redux'
import {navigateAppend} from '../../actions/route-tree'
import {withState, withHandlers, compose, branch, renderNothing, lifecycle} from 'recompose'

import type {Props} from '.'
import type {TypedState} from '../../constants/reducer'

const mapStateToProps = (state: TypedState, {routePath, routeState}) => {
  const selectedConversationIDKey = routePath.last()
  let finalizeInfo = null
  let rekeyInfo = null
  let supersedes = null
  let supersededBy = null

  if (selectedConversationIDKey !== Constants.nothingSelected) {
    const conversationState = state.chat.get('conversationStates').get(selectedConversationIDKey)
    if (conversationState) {
      rekeyInfo = state.chat.get('rekeyInfos').get(selectedConversationIDKey)
      finalizeInfo = state.chat.get('finalizedState').get(selectedConversationIDKey)
      supersedes = Constants.convSupersedesInfo(selectedConversationIDKey, state.chat)
      supersededBy = Constants.convSupersededByInfo(selectedConversationIDKey, state.chat)
    }
  }

  return {
    finalizeInfo,
    rekeyInfo,
    selectedConversationIDKey,
    supersededBy,
    supersedes,
  }
}

const mapDispatchToProps = (dispatch: Dispatch, {setRouteState, navigateUp}) => ({
  onAttach: (selectedConversation, inputs: Array<Constants.AttachmentInput>) => { dispatch(navigateAppend([{props: {conversationIDKey: selectedConversation, inputs}, selected: 'attachmentInput'}])) },
  onBack: () => dispatch(navigateUp()),
})

const mergeProps = (stateProps, dispatchProps) => {
  return {
    ...stateProps,
    ...dispatchProps,
    onAttach: (inputs: Array<Constants.AttachmentInput>) => dispatchProps.onAttach(stateProps.selectedConversationIDKey, inputs),
    onOpenNewerConversation: stateProps.supersededBy
      ? () => stateProps.onOpenConversation(stateProps.supersededBy.conversationIDKey)
      : () => dispatchProps.startConversation(stateProps.participants.toArray()),
  }
}

export default compose(
  connect(mapStateToProps, mapDispatchToProps, mergeProps),
  withState('sidePanelOpen', 'setSidePanelOpen', false),
  withState('focusInputCounter', 'setFocusInputCounter', 0),
  withState('editLastMessageCounter', 'setEditLastMessageCounter', 0),
  withState('listScrollDownCounter', 'setListScrollDownCounter', 0),
  withHandlers({
    onCloseSidePanel: props => () => props.setSidePanelOpen(false),
    onEditLastMessage: props => () => props.setEditLastMessageCounter(props.editLastMessageCounter + 1),
    onFocusInput: props => () => props.setFocusInputCounter(props.focusInputCounter + 1),
    onScrollDown: props => () => props.setListScrollDownCounter(props.listScrollDownCounter + 1),
    onToggleSidePanel: props => () => props.setSidePanelOpen(!props.sidePanelOpen),
  }),
  branch(props => !props.selectedConversationIDKey,
    renderNothing(),
  ),
  lifecycle({
    componentWillReceiveProps: function (nextProps: Props) {
      if (this.props.selectedConversationIDKey !== nextProps.selectedConversationIDKey) {
        this.props.onCloseSidePanel()
      }
    },
  }),
)(Conversation)
