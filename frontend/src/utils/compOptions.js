import axios from 'axios'
import { COMPONENT_DEFINITIONS, LOCAL_STORAGE, REQUEST_URLS } from './constants'
import { getFromLocalStorage } from './localStorage'

export function getNetworkPolicyOptions() {
  return {
    allowAll: 'allow_all',
    denyAll: 'deny_all',
    custom: 'custom',
  }
}

export async function getFormOptions(fieldName) {
  let shouldRequest = false
  Object.keys(COMPONENT_DEFINITIONS).map(comp_definition => {
    if (COMPONENT_DEFINITIONS[comp_definition].formSchema[fieldName]) {
      shouldRequest = true
    }
    return true
  })

  //fetch options from api
  try {
    if (shouldRequest) {
      const res = await axios({
        method: 'GET',
        url: `${REQUEST_URLS.GET.shipaCloud}${fieldName
          .toLowerCase()
          .concat('s')}`,
        headers: {
          'X-Auth-Token': getFromLocalStorage(LOCAL_STORAGE.authToken),
        },
      })
      const data = res.data
      return data
    }
  } catch (error) {
    console.log('error while getting options from api: ', error)
  }
}
