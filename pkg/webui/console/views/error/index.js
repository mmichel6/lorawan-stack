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
import { defineMessages } from 'react-intl'
import { Container, Row, Col } from 'react-grid-system'

import Button from '../../../components/button'
import Message from '../../../lib/components/message'
import ErrorMessage from '../../../lib/components/error-message'
import { withEnv } from '../../../lib/components/env'
import IntlHelmet from '../../../lib/components/intl-helmet'
import Footer from '../../../components/footer'

import Header from '../../containers/header'

import { httpStatusCode, isUnknown as isUnknownError, isNotFoundError } from '../../../lib/errors/utils'

import errorMessages from '../../../lib/errors/status-code-messages'

import style from './full-view.styl'

const m = defineMessages({
  errorTitle: 'Error',
  unknownErrorTitle: 'An unknown error occurred.',
  unknownErrorMessage: 'Try refreshing this page. If the error persists, please contact an administrator.',
  takeMeBack: 'Take me back',
  refresh: 'Refresh page',
  genericNotFound: 'The page you requested cannot be found.',
})

const reload = () => location.reload()

const FullViewErrorInner = function ({ error, env }) {

  const isUnknown = isUnknownError(error)
  const statusCode = httpStatusCode(error)
  const isNotFound = isNotFoundError(error)

  let errorTitleMessage = m.unknownErrorTitle
  let errorMessageMessage = m.unknownErrorMessage
  if (!isUnknown) {
    errorMessageMessage = error
  } else if (isNotFound) {
    errorMessageMessage = m.genericNotFound
  }
  if (statusCode) {
    errorTitleMessage = errorMessages[statusCode]
  }

  return (
    <div className={style.fullViewError}>
      <Container>
        <Row>
          <Col>
            <IntlHelmet title={m.fullViewErrorTitle} />
            <Message
              className={style.fullViewErrorHeader}
              component="h2"
              content={errorTitleMessage}
            />
            <ErrorMessage
              className={style.fullViewErrorSub}
              content={errorMessageMessage}
            />
            { isNotFoundError(error)
              ? (
                <Button.AnchorLink
                  icon="keyboard_arrow_left"
                  message={m.takeMeBack}
                  href={env.app_root}
                />
              )
              : (
                <Button.AnchorLink
                  icon="refresh"
                  message={m.refresh}
                  onClick={reload}
                />
              )
            }
          </Col>
        </Row>
      </Container>
    </div>
  )
}

const FullViewErrorInnerWithEnv = withEnv(FullViewErrorInner)

const FullViewError = function ({ error }) {
  return (
    <div className={style.wrapper}>
      <Header className={style.header} anchored />
      <FullViewErrorInnerWithEnv error={error} />
      <Footer />
    </div>
  )
}

export { FullViewError as default, FullViewErrorInnerWithEnv as FullViewErrorInner }
