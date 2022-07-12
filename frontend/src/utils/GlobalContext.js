import React, { useState, createContext, useEffect } from 'react'
import cloneDeep from 'lodash.clonedeep'
import isEmpty from 'lodash.isempty'
import { useNavigate } from 'react-router-dom'
import { APP_BLOCK_TYPES, LOCAL_STORAGE, ROUTER_PATHS } from './constants'
import { getFromLocalStorage, setInLocalStorage } from './localStorage'
import { getUser } from './auth'

export const AppContext = createContext()

export default function GlobalContext(props) {
  //definitions can be either appDeployments or policies
  const [definitions, setDefinitions] = useState(
    getFromLocalStorage(LOCAL_STORAGE.definitions)
      ? getFromLocalStorage(LOCAL_STORAGE.definitions)
      : {},
  )
  //payload for calling generation api
  const [payload, setPayload] = useState(
    definitions
      ? definitions.type === APP_BLOCK_TYPES.appDeployment
        ? { apps: [] }
        : definitions.type === APP_BLOCK_TYPES.framework
        ? { frameworks: [] }
        : null
      : null,
  )
  const [userData, setUserData] = useState(null)
  const [owner, setOwner] = useState(true)
  const [areFieldsValidated, setAreFieldsValidated] = useState(true)
  const [validationArray, setValidationArray] = useState([])

  const navigate = useNavigate()

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const Usrdata = await getUser(
          getFromLocalStorage(LOCAL_STORAGE.authToken),
        )
        setUserData(Usrdata.data)
      } catch (e) {
        console.log('Access token expired')
        navigate(`${ROUTER_PATHS.loginPage}`, { replace: true })
      }
    }
    if (getFromLocalStorage(LOCAL_STORAGE.authToken)) {
      fetchUserData()
    } else {
      setUserData(null)
    }
  }, [getFromLocalStorage(LOCAL_STORAGE.authToken)])

  useEffect(() => {
    //set appDefinitions in local storage
    setInLocalStorage(LOCAL_STORAGE.definitions, cloneDeep(definitions))

    if (isEmpty(definitions)) {
      setInLocalStorage(LOCAL_STORAGE.loadedConfig, null)
      setValidationArray([])
    }
  }, [definitions])

  return (
    <AppContext.Provider
      value={{
        definitions: definitions,
        setDefinitions: setDefinitions,
        payload: payload,
        setPayload: setPayload,
        userData: userData,
        setUserData: setUserData,
        owner: owner,
        setOwner: setOwner,
        validationArray: validationArray,
        setValidationArray: setValidationArray,
        areFieldsValidated: areFieldsValidated,
        setAreFieldsValidated: setAreFieldsValidated,
      }}
    >
      {props.children}
    </AppContext.Provider>
  )
}
