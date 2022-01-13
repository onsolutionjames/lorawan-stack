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

import React, { Component } from 'react'
import { defineMessages } from 'react-intl'
import bind from 'autobind-decorator'

import Form from '@ttn-lw/components/form'
import Input from '@ttn-lw/components/input'
import SubmitBar from '@ttn-lw/components/submit-bar'
import SubmitButton from '@ttn-lw/components/submit-button'
import Notification from '@ttn-lw/components/notification'
import KeyValueMap from '@ttn-lw/components/key-value-map'
import ModalButton from '@ttn-lw/components/button/modal-button'
import PortalledModal from '@ttn-lw/components/modal/portalled'
import Checkbox from '@ttn-lw/components/checkbox'

import WebhookTemplateInfo from '@console/components/webhook-template-info'

import WebhookFormatSelector from '@console/containers/webhook-formats-select'

import Yup from '@ttn-lw/lib/yup'
import { url as urlRegexp, id as webhookIdRegexp } from '@ttn-lw/lib/regexp'
import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'
import tooltipIds from '@ttn-lw/lib/constants/tooltip-ids'

import { apiKey as webhookAPIKeyRegexp } from '@console/lib/regexp'

import { mapWebhookToFormValues, mapFormValuesToWebhook, blankValues } from './mapping'

const pathPlaceholder = '/path/to/webhook'

const m = defineMessages({
  idPlaceholder: 'my-new-webhook',
  messageInfo:
    'For each enabled message type, an optional path can be defined which will be appended to the base URL',
  deleteWebhook: 'Delete Webhook',
  modalWarning:
    'Are you sure you want to delete webhook "{webhookId}"? Deleting a webhook cannot be undone.',
  additionalHeaders: 'Additional headers',
  headersKeyPlaceholder: 'Authorization',
  headersValuePlaceholder: 'Bearer my-auth-token',
  headersAdd: 'Add header entry',
  headersValidateRequired: 'All header entry values are required. Please remove empty entries.',
  downlinkAPIKey: 'Downlink API key',
  downlinkAPIKeyDesc:
    'The API key will be provided to the endpoint using the "X-Downlink-Apikey" header',
  templateInformation: 'Template information',
  enabledMessages: 'Enabled messages',
  endpointSettings: 'Endpoint settings',
  updateErrorTitle: 'Could not update webhook',
  createErrorTitle: 'Could not create webhook',
  reactivateButtonMessage: 'Reactivate',
  suspendedWebhookMessage:
    'This webhook has been deactivated due to several unsuccessful forwarding attempts. It will be automatically reactivated after 24 hours. If you wish to reactivate right away, you can use the button below.',
  pendingInfo:
    'This webhook is currently pending until attempting its first regular request attempt. Note that webhooks can be deactivated if they encounter too many request failures.',
  messagePathValidateTooLong: 'Enabled message path must be at most 64 characters',
  basicAuthCheckbox: 'Use basic access authentication (basic auth)',
  requestBasicAuth: 'Request authentication',
  hasAuthorization:
    'This webhook already has a header with the same key. Adding another will overwrite the one that is already there',
})

const headerCheck = headers =>
  headers.length === 0 || headers.every(({ value, key }) => value !== '' && key !== '')
const messageCheck = message => {
  if (message && message.enabled) {
    const { value } = message

    return value ? value.length <= 64 : true
  }

  return true
}

const validationSchema = Yup.object().shape({
  webhook_id: Yup.string()
    .min(3, Yup.passValues(sharedMessages.validateTooShort))
    .max(36, Yup.passValues(sharedMessages.validateTooLong))
    .matches(webhookIdRegexp, Yup.passValues(sharedMessages.validateIdFormat))
    .required(sharedMessages.validateRequired),
  format: Yup.string().required(sharedMessages.validateRequired),
  basic_auth: Yup.object().shape({
    value: Yup.boolean(),
    username: Yup.string().default(''),
    password: Yup.string().default(''),
  }),
  headers: Yup.array()
    .of(
      Yup.object({
        key: Yup.string().default(''),
        value: Yup.string().default(''),
      }),
    )
    .test('has no empty string values', m.headersValidateRequired, headerCheck)
    .default([]),
  base_url: Yup.string()
    .matches(urlRegexp, Yup.passValues(sharedMessages.validateUrl))
    .required(sharedMessages.validateRequired),
  downlink_api_key: Yup.string().matches(
    webhookAPIKeyRegexp,
    Yup.passValues(sharedMessages.validateApiKey),
  ),
  uplink_message: Yup.object()
    .shape({
      enabled: Yup.boolean(),
      value: Yup.string(),
    })
    .test('has path length at most 64 characters', m.messagePathValidateTooLong, messageCheck),
  join_accept: Yup.object()
    .shape({
      enabled: Yup.boolean(),
      value: Yup.string(),
    })
    .test('has path length at most 64 characters', m.messagePathValidateTooLong, messageCheck),
  downlink_ack: Yup.object()
    .shape({
      enabled: Yup.boolean(),
      value: Yup.string(),
    })
    .test('has path length at most 64 characters', m.messagePathValidateTooLong, messageCheck),
  downlink_nack: Yup.object()
    .shape({
      enabled: Yup.boolean(),
      value: Yup.string(),
    })
    .test('has path length at most 64 characters', m.messagePathValidateTooLong, messageCheck),
  downlink_sent: Yup.object()
    .shape({
      enabled: Yup.boolean(),
      value: Yup.string(),
    })
    .test('has path length at most 64 characters', m.messagePathValidateTooLong, messageCheck),
  downlink_failed: Yup.object()
    .shape({
      enabled: Yup.boolean(),
      value: Yup.string(),
    })
    .test('has path length at most 64 characters', m.messagePathValidateTooLong, messageCheck),
  downlink_queued: Yup.object()
    .shape({
      enabled: Yup.boolean(),
      value: Yup.string(),
    })
    .test('has path length at most 64 characters', m.messagePathValidateTooLong, messageCheck),
  downlink_queue_invalidated: Yup.object()
    .shape({
      enabled: Yup.boolean(),
      value: Yup.string(),
    })
    .test('has path length at most 64 characters', m.messagePathValidateTooLong, messageCheck),
  location_solved: Yup.object()
    .shape({
      enabled: Yup.boolean(),
      value: Yup.string(),
    })
    .test('has path length at most 64 characters', m.messagePathValidateTooLong, messageCheck),
  service_data: Yup.object()
    .shape({
      enabled: Yup.boolean(),
      value: Yup.string(),
    })
    .test('has path length at most 64 characters', m.messagePathValidateTooLong, messageCheck),
})

export default class WebhookForm extends Component {
  static propTypes = {
    appId: PropTypes.string,
    buttonStyle: PropTypes.shape({ activateWebhookButton: PropTypes.shape({}) }),
    existCheck: PropTypes.func,
    initialWebhookValue: PropTypes.shape({
      basic_auth: PropTypes.bool,
      ids: PropTypes.shape({
        webhook_id: PropTypes.string,
      }),
      health_status: PropTypes.shape({
        unhealthy: PropTypes.shape({}),
      }),
      headers: PropTypes.shape({
        Authorization: PropTypes.string,
      }),
    }),
    onDelete: PropTypes.func,
    onDeleteFailure: PropTypes.func,
    onDeleteSuccess: PropTypes.func,
    onReactivate: PropTypes.func,
    onReactivateSuccess: PropTypes.func,
    onSubmit: PropTypes.func.isRequired,
    onSubmitFailure: PropTypes.func,
    onSubmitSuccess: PropTypes.func.isRequired,
    update: PropTypes.bool.isRequired,
    webhookTemplate: PropTypes.webhookTemplate,
  }

  static defaultProps = {
    appId: undefined,
    initialWebhookValue: undefined,
    onReactivate: () => null,
    onReactivateSuccess: () => null,
    onSubmitFailure: () => null,
    onDeleteFailure: () => null,
    onDeleteSuccess: () => null,
    onDelete: () => null,
    webhookTemplate: undefined,
    existCheck: () => false,
    buttonStyle: undefined,
  }

  form = React.createRef()
  modalResolve = () => null
  modalReject = () => null

  constructor(props) {
    super(props)
    const { initialWebhookValue } = this.props

    this.state = {
      error: undefined,
      displayOverwriteModal: false,
      existingId: undefined,
      mayShowCredentialsInput:
        initialWebhookValue &&
        initialWebhookValue.headers &&
        initialWebhookValue.headers.Authorization &&
        initialWebhookValue.headers.Authorization.startsWith('Basic'),
      hasAuthorization: false,
    }
  }

  @bind
  async handleSubmit(values, { resetForm }) {
    const { appId, onSubmit, onSubmitSuccess, onSubmitFailure, existCheck, update } = this.props
    // Ensure that the basic auth header is deleted properly.
    if (values.basic_auth && values.basic_auth.value === false) {
      values.headers = values.headers.filter(
        header => header.key !== 'Authorizarion' && !header.value.startsWith('Basic'),
      )
    }
    const webhook = mapFormValuesToWebhook(values, appId)
    await this.setState({ error: '' })

    try {
      if (!update) {
        const webhookId = webhook.ids.webhook_id
        const exists = await existCheck(webhookId)
        if (exists) {
          this.setState({ displayOverwriteModal: true, existingId: webhookId })
          await new Promise((resolve, reject) => {
            this.modalResolve = resolve
            this.modalReject = reject
          })
        }
      }
      const result = await onSubmit(webhook)

      resetForm({ values })
      await onSubmitSuccess(result)
    } catch (error) {
      resetForm({ values })
      await this.setState({ error })
      await onSubmitFailure(error)
    }
  }

  @bind
  async handleDelete() {
    const { onDelete, onDeleteSuccess, onDeleteFailure } = this.props
    try {
      await onDelete()
      this.form.current.resetForm()
      onDeleteSuccess()
    } catch (error) {
      await this.setState({ error })
      onDeleteFailure()
    }
  }

  @bind
  handleReplaceModalDecision(mayReplace) {
    if (mayReplace) {
      this.modalResolve()
    } else {
      this.modalReject()
    }
    this.setState({ displayOverwriteModal: false })
  }

  @bind
  async handleReactivate() {
    const { onReactivate, onReactivateSuccess } = this.props
    const healthStatus = {
      health_status: null,
    }

    try {
      await onReactivate(healthStatus)
      onReactivateSuccess()
    } catch (error) {
      this.setState({ error })
    }
  }

  @bind
  handleRequestAuthenticationChange(event) {
    const { initialWebhookValue } = this.props
    this.setState({ mayShowCredentialsInput: event.target.checked })

    if (event.target.checked) {
      this.setState({ hasAuthorization: Boolean(initialWebhookValue.headers.Authorization) })
    }
  }

  @bind
  handleHeadersChange(event) {
    // Check if the headers have repeated keys.
    const headerKeys = event.map(header => header.key)
    this.setState({ hasAuthorization: event.length !== new Set(headerKeys).size })
  }

  render() {
    const { update, initialWebhookValue, webhookTemplate, buttonStyle } = this.props
    const { error, displayOverwriteModal, existingId } = this.state
    let initialValues = blankValues
    if (update && initialWebhookValue) {
      initialValues = mapWebhookToFormValues(initialWebhookValue)
    }

    const mayReactivate =
      update &&
      initialWebhookValue &&
      initialWebhookValue.health_status &&
      initialWebhookValue.health_status.unhealthy

    const isPending =
      update &&
      initialWebhookValue &&
      (initialWebhookValue.health_status === null ||
        initialWebhookValue.health_status === undefined)

    return (
      <>
        {mayReactivate && (
          <Notification
            warning
            content={m.suspendedWebhookMessage}
            action={this.handleReactivate}
            actionMessage={m.reactivateButtonMessage}
            buttonIcon="refresh"
            className={buttonStyle.activateWebhookButton}
            small
          />
        )}
        {isPending && <Notification info content={m.pendingInfo} small />}
        <PortalledModal
          title={sharedMessages.idAlreadyExists}
          message={{
            ...sharedMessages.webhookAlreadyExistsModalMessage,
            values: { id: existingId },
          }}
          buttonMessage={sharedMessages.replaceWebhook}
          onComplete={this.handleReplaceModalDecision}
          approval
          visible={displayOverwriteModal}
        />
        {Boolean(webhookTemplate) && <WebhookTemplateInfo webhookTemplate={webhookTemplate} />}
        <Form
          onSubmit={this.handleSubmit}
          validationSchema={validationSchema}
          initialValues={initialValues}
          error={error}
          errorTitle={update ? m.updateErrorTitle : m.createErrorTitle}
          formikRef={this.form}
        >
          <Form.SubTitle title={sharedMessages.generalSettings} />
          <Form.Field
            name="webhook_id"
            title={sharedMessages.webhookId}
            placeholder={m.idPlaceholder}
            component={Input}
            required
            autoFocus
            disabled={update}
          />
          <WebhookFormatSelector name="format" required />
          <Form.SubTitle title={m.endpointSettings} />
          <Form.Field
            name="base_url"
            title={sharedMessages.webhookBaseUrl}
            placeholder="https://example.com/webhooks"
            component={Input}
          />
          <Form.Field
            name="downlink_api_key"
            title={m.downlinkAPIKey}
            component={Input}
            description={m.downlinkAPIKeyDesc}
            sensitive
            code
          />
          {this.state.hasAuthorization && (
            <Notification content={m.hasAuthorization} small warning />
          )}
          <Form.Field
            autoFocus
            title={m.requestBasicAuth}
            name="basic_auth.value"
            label={m.basicAuthCheckbox}
            onChange={this.handleRequestAuthenticationChange}
            component={Checkbox}
            tooltipId={tooltipIds.BASIC_AUTH}
          />
          {this.state.mayShowCredentialsInput && (
            <Form.FieldContainer horizontal>
              <Form.Field
                required
                title={sharedMessages.username}
                name="basic_auth.username"
                component={Input}
              />
              <Form.Field
                required
                title={sharedMessages.password}
                name="basic_auth.password"
                component={Input}
                sensitive
              />
            </Form.FieldContainer>
          )}
          <Form.Field
            name="headers"
            title={m.additionalHeaders}
            keyPlaceholder={m.headersKeyPlaceholder}
            valuePlaceholder={m.headersValuePlaceholder}
            addMessage={m.headersAdd}
            onChange={this.handleHeadersChange}
            component={KeyValueMap}
          />
          <Form.SubTitle title={m.enabledMessages} />
          <Notification info content={m.messageInfo} small />
          <Form.Field
            name="uplink_message"
            type="toggled-input"
            title={sharedMessages.uplinkMessage}
            placeholder={pathPlaceholder}
            component={Input.Toggled}
          />
          <Form.Field
            name="join_accept"
            type="toggled-input"
            title={sharedMessages.joinAccept}
            placeholder={pathPlaceholder}
            component={Input.Toggled}
          />
          <Form.Field
            name="downlink_ack"
            type="toggled-input"
            title={sharedMessages.downlinkAck}
            placeholder={pathPlaceholder}
            component={Input.Toggled}
          />
          <Form.Field
            name="downlink_nack"
            type="toggled-input"
            title={sharedMessages.downlinkNack}
            placeholder={pathPlaceholder}
            component={Input.Toggled}
          />
          <Form.Field
            name="downlink_sent"
            type="toggled-input"
            title={sharedMessages.downlinkSent}
            placeholder={pathPlaceholder}
            component={Input.Toggled}
          />
          <Form.Field
            name="downlink_failed"
            type="toggled-input"
            title={sharedMessages.downlinkFailed}
            placeholder={pathPlaceholder}
            component={Input.Toggled}
          />
          <Form.Field
            name="downlink_queued"
            type="toggled-input"
            title={sharedMessages.downlinkQueued}
            placeholder={pathPlaceholder}
            component={Input.Toggled}
          />
          <Form.Field
            name="downlink_queue_invalidated"
            type="toggled-input"
            title={sharedMessages.downlinkQueueInvalidated}
            placeholder={pathPlaceholder}
            component={Input.Toggled}
          />
          <Form.Field
            name="location_solved"
            type="toggled-input"
            title={sharedMessages.locationSolved}
            placeholder={pathPlaceholder}
            component={Input.Toggled}
          />
          <Form.Field
            name="service_data"
            type="toggled-input"
            title={sharedMessages.serviceData}
            placeholder={pathPlaceholder}
            component={Input.Toggled}
          />
          <SubmitBar>
            <Form.Submit
              component={SubmitButton}
              message={update ? sharedMessages.saveChanges : sharedMessages.addWebhook}
            />
            {update && (
              <ModalButton
                type="button"
                icon="delete"
                danger
                naked
                message={m.deleteWebhook}
                modalData={{
                  message: {
                    values: { webhookId: initialWebhookValue.ids.webhook_id },
                    ...m.modalWarning,
                  },
                }}
                onApprove={this.handleDelete}
              />
            )}
          </SubmitBar>
        </Form>
      </>
    )
  }
}
