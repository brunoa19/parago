import React, { useContext, useState } from 'react'
import { AppContext } from '../../utils/GlobalContext'
import { Form, Input, Button, Select, Spin, notification } from 'antd'
import { CloseCircleOutlined } from '@ant-design/icons'
import { getConfigDetails, save, updateById } from '../../utils/configRequests'
import { APP_BLOCK_TYPES, LOCAL_STORAGE } from '../../utils/constants'
import {
  getFromLocalStorage,
  setInLocalStorage,
} from '../../utils/localStorage'

export default function ConfigInfoPopup({ setConfigInfoPopup }) {
  const { definitions, payload, setPayload, owner } = useContext(AppContext)
  const [loading, setLoading] = useState(false)

  const { Option } = Select
  function openNotificationWithIcon(type, message) {
    if (type === 'success') {
      notification[type]({
        message: message,
      })
    }
    if (type === 'error') {
      notification[type]({
        message: message,
      })
    }
  }

  async function handleFinish(data) {
    setLoading(true)
    if (owner) {
      //update
      const configInfo = {
        accessLevel: data.accessLevel,
        description: data.description,
      }

      const res = await updateById(definitions, payload, setPayload, configInfo)

      if (res && res.status === 200) {
        openNotificationWithIcon('success', 'updated')
      } else {
        openNotificationWithIcon('error', res.response.data.error)
      }
    } else {
      const configInfo = {
        name: data.name,
        accessLevel: data.accessLevel,
        description: data.description,
        provider: data.provider,
      }
      const res = await save(definitions, payload, setPayload, configInfo)
      if (res && res.status === 200) {
        let configDetails = null
        configDetails = await getConfigDetails(res.data.id)
        setInLocalStorage(LOCAL_STORAGE.loadedConfig, configDetails.data)
        openNotificationWithIcon('success', 'saved')
      } else {
        setInLocalStorage(LOCAL_STORAGE.loadedConfig, null)
        openNotificationWithIcon('error', res.response.data.error)
      }
    }

    setLoading(false)
    setConfigInfoPopup(false)
  }

  return (
    <div
      className="flex backdrop-blur-sm absolute"
      style={{ height: '100vh', width: '100vw', zIndex: 11 }}
    >
      <div className="flex absolute top-1/3 left-1/2 px-4 pt-12 bg-white rounded-xl border border-dashed border-green-300 drop-shadow-md">
        <div className="absolute top-0 right-0 text-xl">
          <Button
            className="rounded-md"
            type="text"
            icon={<CloseCircleOutlined />}
            onClick={() => {
              setPayload(
                definitions
                  ? definitions.type === APP_BLOCK_TYPES.appDeployment
                    ? { apps: [] }
                    : definitions.type === APP_BLOCK_TYPES.framework
                    ? { frameworks: [] }
                    : null
                  : null,
              )
              setConfigInfoPopup(false)
            }}
          ></Button>
        </div>
        <div>
          <Form
            name="save_definition"
            className="save_definition w-72"
            onFinish={handleFinish}
          >
            {!owner && (
              <Form.Item
                name="name"
                rules={[
                  {
                    required: true,
                    message: 'Please enter a name for this definition!',
                  },
                ]}
              >
                <Input placeholder="name" />
              </Form.Item>
            )}
            {/* <Form.Item
              name="provider"
              rules={[
                {
                  required: true,
                  message: 'Please select a provider for this definition!',
                },
              ]}
            >
              <Select placeholder="Please select a provider">
                {Object.keys(PROVIDERS).map(providerName => {
                  return (
                    <Option
                      value={PROVIDERS[providerName]}
                      key={PROVIDERS[providerName]}
                    >
                      {providerName}
                    </Option>
                  )
                })}
              </Select>
            </Form.Item> */}
            <Form.Item
              name="accessLevel"
              rules={[
                {
                  required: true,
                  message: 'Please select an access level!',
                },
              ]}
              initialValue={
                getFromLocalStorage(LOCAL_STORAGE.loadedConfig)
                  ? getFromLocalStorage(LOCAL_STORAGE.loadedConfig).accessLevel
                  : 'public'
              }
            >
              <Select placeholder="Select an access level">
                <Option value="public">Public</Option>
                <Option value="organization">Organization</Option>
                <Option value="private">Private</Option>
              </Select>
            </Form.Item>
            <Form.Item
              name="description"
              rules={[
                {
                  required: false,
                },
              ]}
              initialValue={
                getFromLocalStorage(LOCAL_STORAGE.loadedConfig)
                  ? getFromLocalStorage(LOCAL_STORAGE.loadedConfig).description
                  : ''
              }
            >
              <Input.TextArea placeholder="description" rows={5} />
            </Form.Item>
            <Form.Item>
              <div className="flex justify-center">
                {loading ? (
                  <Spin />
                ) : (
                  <Button block htmlType="submit" className="login-form-button">
                    {owner ? 'Update' : 'Save'}
                  </Button>
                )}
              </div>
            </Form.Item>
          </Form>
        </div>
      </div>
    </div>
  )
}
