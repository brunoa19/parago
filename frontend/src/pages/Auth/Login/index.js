import React, { useContext, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { Form, Input, Button, Spin } from 'antd'
import { UserOutlined, LockOutlined } from '@ant-design/icons'
import shipaLogo from '../../../assets/images/shipalogo.png'
import { LOCAL_STORAGE, ROUTER_PATHS } from '../../../utils/constants'
import { AppContext } from '../../../utils/GlobalContext'
import { authenticateUser } from '../../../utils/auth'
import {
  getFromLocalStorage,
  setInLocalStorage,
} from '../../../utils/localStorage'
import cloneDeep from 'lodash.clonedeep'

export default function LoginPage() {
  const [loading, setLoading] = useState(false)
  const { setUserData, setDefinitions } = useContext(AppContext)

  let navigate = useNavigate()

  const onFinish = async (values: any) => {
    setLoading(true)
    const authData = await authenticateUser(values.email, values.password)
    if (authData) {
      setInLocalStorage(LOCAL_STORAGE.authToken, authData.token)
      setUserData(authData)

      //reload config data
      if (getFromLocalStorage(LOCAL_STORAGE.loadedConfig)) {
        const configData = cloneDeep(
          getFromLocalStorage(LOCAL_STORAGE.loadedConfig).definition,
        )
        setDefinitions({
          ...configData,
        })
      }

      navigate(ROUTER_PATHS.designPage, { replace: true })
    }
  }
  return (
    <div className="flex h-screen">
      <div className="m-auto px-16 py-10 rounded-md drop-shadow-xl bg-gray-50">
        <div>
          <div className="flex justify-center mb-4">
            <img style={{ height: '70px' }} src={shipaLogo} alt="logo" />
          </div>
          <Form
            name="normal_login"
            className="login-form w-72"
            initialValues={{ remember: true }}
            onFinish={onFinish}
          >
            <Form.Item
              name="email"
              rules={[
                {
                  type: 'email',
                  message: 'The input is not valid E-mail!',
                },
                { required: true, message: 'Please input your email!' },
              ]}
            >
              <Input
                prefix={<UserOutlined className="site-form-item-icon" />}
                placeholder="email"
              />
            </Form.Item>
            <Form.Item
              name="password"
              rules={[
                { required: true, message: 'Please input your Password!' },
              ]}
            >
              <Input
                prefix={<LockOutlined className="site-form-item-icon" />}
                type="password"
                placeholder="Password"
              />
            </Form.Item>
            {/* <Form.Item>
              <Form.Item name="remember" valuePropName="checked" noStyle>
                <Checkbox>Remember me</Checkbox>
              </Form.Item>
              <a className="login-form-forgot" href="">
                Forgot password
              </a>
            </Form.Item> */}

            <Form.Item>
              <div className="flex justify-center">
                {loading ? (
                  <Spin />
                ) : (
                  <div className="w-full space-y-2">
                    <Button
                      block
                      htmlType="submit"
                      className="login-form-button"
                    >
                      Log in
                    </Button>
                    <Button
                      block
                      className="login-form-button"
                      onClick={() => {
                        window.location.href =
                          'https://apps.shipa.cloud/sign-up'
                      }}
                    >
                      Sign up
                    </Button>
                  </div>
                )}
              </div>
            </Form.Item>
          </Form>
        </div>
      </div>
    </div>
  )
}
