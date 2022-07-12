import React, { useContext, useEffect, useState } from 'react'
import { Button, Layout } from 'antd'
import { SaveOutlined, UnlockOutlined } from '@ant-design/icons'
import { Link, useLocation } from 'react-router-dom'
import shipaLogo from '../../assets/images/shipalogo.png'
import { LOCAL_STORAGE, ROUTER_PATHS } from '../../utils/constants'
import { AppContext } from '../../utils/GlobalContext'
import { logout } from '../../utils/auth'
import { SearchInput } from '../search-input'
import { updatePayload } from '../../utils/payload'
import { getFromLocalStorage } from '../../utils/localStorage'

export default function Header({ setConfigInfoPopup }) {
  const { Header } = Layout
  const {
    definitions,
    setPayload,
    userData,
    setUserData,
    owner,
    setOwner,
    validationArray,
    setAreFieldsValidated,
  } = useContext(AppContext)
  const location = useLocation()
  const [config, setConfig] = useState(
    getFromLocalStorage(LOCAL_STORAGE.loadedConfig),
  )

  useEffect(() => {
    if (userData && config && config.userInfo) {
      if (userData.userId === config.userInfo.userId) {
        setOwner(true)
      } else {
        setOwner(false)
      }
    } else {
      setOwner(false)
    }
  }, [userData, definitions, config])

  useEffect(() => {
    function loadConfig() {
      const item = getFromLocalStorage(LOCAL_STORAGE.loadedConfig)
      if (item) {
        setConfig(item)
      } else {
        setConfig(null)
      }
    }
    window.addEventListener('storage', loadConfig)

    return () => {
      window.removeEventListener('storage', loadConfig)
    }
  }, [])

  function handleSubmit() {
    if (validationArray.length > 0) {
      setAreFieldsValidated(false)
    } else {
      updatePayload(definitions, setPayload)
      setConfigInfoPopup(true)
    }
  }

  return (
    <Header
      className="p-0 fixed top-0 w-full h-fit"
      style={{ zIndex: 11, backgroundColor: '#ececec' }}
    >
      <div className="flex flex-row justify-between items-center">
        {/* logo  */}
        <div className="w-fit">
          <Link to={ROUTER_PATHS.designPage}>
            <div>
              <img className="h-16 w-32" src={shipaLogo} alt="logo" />
            </div>
          </Link>
        </div>
        {/* search input */}
        <div>
          <SearchInput />
        </div>
        {/* menu items */}
        <div className="flex">
          {userData && location.pathname === ROUTER_PATHS.designPage && (
            <Button
              className="h-12 ml-4 mr-2 mt-2 hover:text-white hover:bg-shipaOrange hover:border-shipaOrange"
              onClick={handleSubmit}
              shape="round"
              disabled={definitions.type ? false : true}
              icon={<SaveOutlined />}
            >
              {owner ? 'Update' : 'Save'}
            </Button>
          )}
          {userData ? (
            <Link to={ROUTER_PATHS.designPage}>
              <Button
                className="h-12 ml-2 mr-4 hover:text-white hover:bg-shipaOrange hover:border-shipaOrange"
                onClick={() => {
                  logout(setUserData)
                }}
                shape="round"
                icon={<UnlockOutlined />}
              >
                Log out
              </Button>
            </Link>
          ) : (
            <Link to={ROUTER_PATHS.loginPage}>
              <Button
                className="h-12 ml-4 mr-4 hover:text-white hover:bg-shipaOrange hover:border-shipaOrange"
                shape="round"
                icon={<UnlockOutlined />}
              >
                Login
              </Button>
            </Link>
          )}
        </div>
      </div>
    </Header>
  )
}
