// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'
import getByPath from '@ttn-lw/lib/get-by-path'

import messages from '../messages'

import DescriptionList from './shared/description-list'

const JoinResponsePreview = React.memo(({ event }) => {
  const { identifiers, data } = event
  const ids = identifiers[0].device_ids

  const sessionKeyId = getByPath(data, 'session_keys.session_key_id')

  return (
    <DescriptionList>
      <DescriptionList.Byte title={messages.devAddr} data={ids.dev_addr} />
      <DescriptionList.Byte title={sharedMessages.joinEUI} data={ids.join_eui} />
      <DescriptionList.Byte title={sharedMessages.devEUI} data={ids.dev_eui} />
      <DescriptionList.Byte title={messages.sessionKeyId} data={sessionKeyId} convertToHex />
    </DescriptionList>
  )
})

JoinResponsePreview.propTypes = {
  event: PropTypes.event.isRequired,
}

export default JoinResponsePreview
