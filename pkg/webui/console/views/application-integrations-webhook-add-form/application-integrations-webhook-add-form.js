// Copyright © 2021 The Things Network Foundation, The Things Industries B.V.
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
import { Container, Col, Row } from 'react-grid-system'
import { defineMessages } from 'react-intl'

import tts from '@console/api/tts'

import PageTitle from '@ttn-lw/components/page-title'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'

import NotFoundRoute from '@ttn-lw/lib/components/not-found-route'

import WebhookForm from '@console/components/webhook-form'
import WebhookTemplateForm from '@console/components/webhook-template-form'

import { isNotFoundError } from '@ttn-lw/lib/errors/utils'
import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'

const m = defineMessages({
  addCustomWebhook: 'Add custom webhook',
  addWebhookViaTemplate: 'Add {templateName} webhook',
  customWebhook: 'Custom webhook',
})

const ApplicationWebhookAddForm = props => {
  const { templateId, isCustom, webhookTemplate, appId, navigateToList } = props

  let breadcrumbContent = m.customWebhook
  if (!templateId) {
    breadcrumbContent = sharedMessages.add
  } else if (!isCustom && name) {
    breadcrumbContent = name
  }

  useBreadcrumbs(
    'apps.single.integrations.webhooks.various.add',
    <Breadcrumb
      path={`/applications/${appId}/integrations/webhooks/add/template/${templateId}`}
      content={breadcrumbContent}
    />,
  )

  let pageTitle = m.addCustomWebhook
  if (!webhookTemplate) {
    pageTitle = sharedMessages.addWebhook
  } else if (isCustom) {
    pageTitle = {
      ...m.addWebhookViaTemplate,
      values: {
        templateName: webhookTemplate.name,
      },
    }
  }

  const handleSubmit = React.useCallback(
    async webhook => {
      await tts.Applications.Webhooks.create(appId, webhook)
    },
    [appId],
  )
  const handleSubmitSuccess = React.useCallback(() => {
    navigateToList(appId)
  }, [appId, navigateToList])

  const existCheck = React.useCallback(
    async webhookId => {
      try {
        await tts.Applications.Webhooks.getById(appId, webhookId, [])
        return true
      } catch (error) {
        if (isNotFoundError(error)) {
          return false
        }

        throw error
      }
    },
    [appId],
  )

  // Render Not Found error when the template was not found.
  if (!isCustom && templateId && !webhookTemplate) {
    return <NotFoundRoute />
  }

  return (
    <Container>
      <PageTitle title={pageTitle} />
      <Row>
        <Col lg={8} md={12}>
          {isCustom ? (
            <WebhookForm
              update={false}
              onSubmit={handleSubmit}
              onSubmitSuccess={handleSubmitSuccess}
              existCheck={existCheck}
            />
          ) : (
            <WebhookTemplateForm
              appId={appId}
              templateId={templateId}
              onSubmit={handleSubmit}
              onSubmitSuccess={handleSubmitSuccess}
              webhookTemplate={webhookTemplate}
              existCheck={existCheck}
            />
          )}
        </Col>
      </Row>
    </Container>
  )
}

ApplicationWebhookAddForm.propTypes = {
  appId: PropTypes.string.isRequired,
  isCustom: PropTypes.bool.isRequired,
  navigateToList: PropTypes.func.isRequired,
  templateId: PropTypes.string,
  webhookTemplate: PropTypes.webhookTemplate,
}

ApplicationWebhookAddForm.defaultProps = {
  templateId: undefined,
  webhookTemplate: undefined,
}

export default ApplicationWebhookAddForm
