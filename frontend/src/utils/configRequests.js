import {
  APP_BLOCK_TYPES,
  LOCAL_STORAGE,
  PROVIDERS,
  REQUEST_URLS,
} from './constants'
import axios from 'axios'
import { getFromLocalStorage } from './localStorage'

export async function save(definitions, payload, setPayload, configInfo) {
  try {
    const requestPayload = {
      ...configInfo,
      data: { ...payload, provider: PROVIDERS.Terraform },
      definition: { ...definitions },
    }
    const URL = REQUEST_URLS.POST.save
    const requestOptions = {
      method: 'POST',
      url: URL,
      headers: {
        'X-Auth-Token': getFromLocalStorage(LOCAL_STORAGE.authToken),
      },
      data: requestPayload,
    }
    const res = await axios({ ...requestOptions })
    //reset payload
    setPayload(
      definitions
        ? definitions.type === APP_BLOCK_TYPES.appDeployment
          ? { apps: [] }
          : definitions.type === APP_BLOCK_TYPES.framework
          ? { frameworks: [] }
          : null
        : null,
    )
    return res
  } catch (e) {
    //reset payload
    setPayload(
      definitions
        ? definitions.type === APP_BLOCK_TYPES.appDeployment
          ? { apps: [] }
          : definitions.type === APP_BLOCK_TYPES.framework
          ? { frameworks: [] }
          : null
        : null,
    )
    console.log('error while saving: ', e)
    return e
  }
}
export async function updateById(definitions, payload, setPayload, configInfo) {
  try {
    const loadedConfig = getFromLocalStorage(LOCAL_STORAGE.loadedConfig)
    const URL = `${REQUEST_URLS.POST.update}${loadedConfig.id}`

    const requestPayload = {
      ...configInfo,
      data: { ...payload, provider: PROVIDERS.Terraform },
      definition: { ...definitions },
    }
    const requestOptions = {
      method: 'POST',
      url: URL,
      headers: {
        'X-Auth-Token': getFromLocalStorage(LOCAL_STORAGE.authToken),
      },
      data: requestPayload,
    }
    const res = await axios({ ...requestOptions })
    //reset payload
    setPayload(
      definitions
        ? definitions.type === APP_BLOCK_TYPES.appDeployment
          ? { apps: [] }
          : definitions.type === APP_BLOCK_TYPES.framework
          ? { frameworks: [] }
          : null
        : null,
    )
    return res
  } catch (e) {
    //reset payload
    setPayload(
      definitions
        ? definitions.type === APP_BLOCK_TYPES.appDeployment
          ? { apps: [] }
          : definitions.type === APP_BLOCK_TYPES.framework
          ? { frameworks: [] }
          : null
        : null,
    )
    console.log('error while updating: ', e)
    return e
  }
}
export async function getConfigDetails(id) {
  try {
    const URL = `${REQUEST_URLS.GET.id}${id}`
    const requestOptions = {
      method: 'GET',
      url: URL,
      headers: getFromLocalStorage(LOCAL_STORAGE.authToken)
        ? {
            'X-Auth-Token': getFromLocalStorage(LOCAL_STORAGE.authToken),
          }
        : {},
    }
    const res = await axios({ ...requestOptions })
    return res
  } catch (e) {
    console.log('error while getting config details:', e)
  }
}
export async function searchByName(name) {
  try {
    const URL = `${REQUEST_URLS.GET.search}?q=${name}`
    const requestOptions = {
      method: 'GET',
      url: URL,
      headers: getFromLocalStorage(LOCAL_STORAGE.authToken)
        ? {
            'X-Auth-Token': getFromLocalStorage(LOCAL_STORAGE.authToken),
          }
        : {},
    }
    const res = await axios({ ...requestOptions })
    return res
  } catch (e) {
    console.log('error while searching by name:', e)
  }
}
export async function cloneById(id) {
  try {
    const URL = `${REQUEST_URLS.GET.clone}${id}`

    const requestOptions = {
      method: 'GET',
      url: URL,
      headers: {
        'X-Auth-Token': getFromLocalStorage(LOCAL_STORAGE.authToken),
      },
    }
    const res = await axios({ ...requestOptions })
    return res
  } catch (e) {
    console.log('error while cloning by id:', e)
  }
}
export async function removeById(id, org, user) {
  try {
    const URL = `${REQUEST_URLS.DELETE.delete}${id}`

    const requestOptions = {
      method: 'DELETE',
      url: URL,
      headers: {
        'X-Auth-Token': getFromLocalStorage(LOCAL_STORAGE.authToken),
      },
    }
    const res = await axios({ ...requestOptions })
    return res
  } catch (e) {
    console.log('error while deleting:', e)
  }
}
export async function listPublicConfigs(page, pageSize) {
  try {
    const URL = `${REQUEST_URLS.GET.publicConfigs}?page=${page}&pageSize=${pageSize}`

    const requestOptions = {
      method: 'GET',
      url: URL,
      headers: {
        'X-Auth-Token': getFromLocalStorage(LOCAL_STORAGE.authToken),
      },
    }
    const res = await axios({ ...requestOptions })
    return res
  } catch (e) {
    console.log('error while listing public configs:', e)
    return false
  }
}
export async function listOwnedConfigs() {
  try {
    const URL = REQUEST_URLS.GET.ownedConfigs

    const requestOptions = {
      method: 'GET',
      url: URL,
      headers: {
        'X-Auth-Token': getFromLocalStorage(LOCAL_STORAGE.authToken),
      },
    }
    const res = await axios({ ...requestOptions })
    return res
  } catch (e) {
    console.log('error while listing owned configs:', e)
    return false
  }
}
export async function listOrgConfigs() {
  try {
    const URL = REQUEST_URLS.GET.orgConfigs

    const requestOptions = {
      method: 'GET',
      url: URL,
      headers: {
        'X-Auth-Token': getFromLocalStorage(LOCAL_STORAGE.authToken),
      },
    }
    const res = await axios({ ...requestOptions })
    return res
  } catch (e) {
    console.log('error while listing orgnization configs:', e)
    return false
  }
}
export async function listPrivateConfigs() {
  try {
    const URL = REQUEST_URLS.GET.privateConfigs

    const requestOptions = {
      method: 'GET',
      url: URL,
      headers: {
        'X-Auth-Token': getFromLocalStorage(LOCAL_STORAGE.authToken),
      },
    }
    const res = await axios({ ...requestOptions })
    return res
  } catch (e) {
    console.log('error while listing private configs:', e)
    return false
  }
}
export async function getMetrics(configId) {
  try {
    const URL = `${REQUEST_URLS.GET.metrics}${configId}/metrics`
    const requestOptions = {
      method: 'GET',
      url: URL,
      headers: getFromLocalStorage(LOCAL_STORAGE.authToken)
        ? {
            'X-Auth-Token': getFromLocalStorage(LOCAL_STORAGE.authToken),
          }
        : {},
    }
    const res = await axios({ ...requestOptions })
    return res
  } catch (e) {
    console.log('error while getting metrics:', e)
  }
}
