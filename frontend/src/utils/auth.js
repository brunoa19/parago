import axios from 'axios'
import { LOCAL_STORAGE, REQUEST_URLS } from './constants'
import { getFromLocalStorage, setInLocalStorage } from './localStorage'

export async function authenticateUser(email, password) {
  const creds = { login: email, password: password }
  const requestOptions = {
    method: 'POST',
    url: REQUEST_URLS.POST.login,
    headers: {},
    data: creds,
  }
  try {
    const res = await axios({ ...requestOptions })
    if (res.status !== 200) {
      alert('Bad request!')
      return false
    }
    return res.data
  } catch (e) {
    console.log(e)
    return false
  }
}

export async function getUser(token) {
  const res = await axios({
    method: 'GET',
    url: REQUEST_URLS.GET.users,
    headers: {
      'X-Auth-Token': token,
    },
  })
  return res
}

export async function logout(setUserData) {
  const res = await axios({
    method: 'POST',
    url: REQUEST_URLS.POST.logout,
    headers: {
      'X-Auth-Token': getFromLocalStorage(LOCAL_STORAGE.authToken),
    },
  })
  if (res.status === 200) {
    setInLocalStorage(LOCAL_STORAGE.authToken, null)
    setUserData(null)
  }
  return res
}
