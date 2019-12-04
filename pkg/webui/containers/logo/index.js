// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

import LogoComponent from '../../components/logo'

import {
  selectAssetsRootPath,
  selectBrandingRootPath,
  selectBrandingText,
  selectBrandingClusterID,
  selectApplicationSiteName,
} from '../../lib/selectors/env'

const logo = {
  src: `${selectAssetsRootPath()}/ttes.svg`,
  alt: `${selectApplicationSiteName()} Logo`,
}
const hasCustomBranding = selectBrandingRootPath() !== selectAssetsRootPath()
const secondaryLogo = hasCustomBranding
  ? {
      src: `${selectBrandingRootPath()}/logo.svg`,
      alt: 'Logo',
    }
  : undefined

const Logo = props => (
  <LogoComponent
    logo={logo}
    secondaryLogo={secondaryLogo}
    {...props}
    text={selectBrandingText()}
    clusterTag={selectBrandingClusterID()}
  />
)

export default Logo
