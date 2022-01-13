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

describe('Application Webhook', () => {
  const userId = 'create-app-test-user'
  const user = {
    ids: { user_id: userId },
    primary_email_address: 'create-app-test-user@example.com',
    password: 'ABCDefg123!',
    password_confirm: 'ABCDefg123!',
  }
  const appId = 'webhook-test-application'
  const application = {
    ids: {
      application_id: appId,
    },
  }
  const webhookId = 'my-edit-test-webhook'
  const webhookBody = {
    webhook: {
      base_url: 'https://example.com/edit-webhooks-test',
      basic_auth: {
        password: '',
        username: '',
      },
      format: 'json',
      ids: {
        application_ids: {},
        webhook_id: webhookId,
      },
    },
    fieldMasks: ['base_url', 'format', 'ids', 'ids.application_ids', 'ids.webhook_id'],
  }

  before(() => {
    cy.dropAndSeedDatabase()
    cy.createUser(user)
    cy.createApplication(application, userId)
    cy.createWebhook(appId, webhookBody)
  })

  beforeEach(() => {
    cy.loginConsole({ user_id: userId, password: user.password })
    cy.visit(
      `${Cypress.config(
        'consoleRootPath',
      )}/applications/${appId}/integrations/webhooks/${webhookId}`,
    )
  })

  it('succeeds editing webhook', () => {
    const webhook = {
      format: 'Protocol Buffers',
      url: 'https://example.com/webhooks-updated',
      path: 'path/to/webhook',
    }

    cy.findByLabelText('Webhook format').selectOption(webhook.format)
    cy.findByLabelText('Base URL').type(webhook.url)
    cy.get('#uplink_message_checkbox').check()
    cy.findByLabelText('Uplink message').type(webhook.path)

    cy.findByRole('button', { name: 'Save changes' }).click()

    cy.findByTestId('error-notification').should('not.exist')
    cy.findByTestId('toast-notification')
      .should('be.visible')
      .findByText(`Webhook updated`)
      .should('be.visible')
  })

  it('succeeds adding application attributes', () => {
    cy.findByRole('button', { name: /Add header entry/ }).click()

    cy.findByTestId('headers[0].key').type('webhook-test-key')
    cy.findByTestId('headers[0].value').type('webhook-test-value')

    cy.findByRole('button', { name: 'Save changes' }).click()

    cy.findByTestId('error-notification').should('not.exist')
    cy.findByTestId('toast-notification')
      .should('be.visible')
      .findByText(`Webhook updated`)
      .should('be.visible')
  })

  it('succeeds deleting webhook', () => {
    cy.findByRole('button', { name: /Delete Webhook/ }).click()

    cy.findByTestId('modal-window')
      .should('be.visible')
      .within(() => {
        cy.findByText('Delete Webhook', { selector: 'h1' }).should('be.visible')
        cy.findByRole('button', { name: /Delete Webhook/ }).click()
      })

    cy.findByTestId('error-notification').should('not.exist')

    cy.location('pathname').should(
      'eq',
      `${Cypress.config('consoleRootPath')}/applications/${appId}/integrations/webhooks`,
    )

    cy.findByRole('cell', { name: webhookId }).should('not.exist')
  })
})
