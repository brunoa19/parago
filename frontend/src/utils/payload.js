import {
  APP_BLOCK_TYPES,
  COMPONENT_DEFINITIONS,
  NETWORK_POLICY_OPTIONS,
  REQUEST_URLS,
} from './constants'

function getMultipleAppFields(componentPayload, componentSchema) {
  if (componentSchema.multiple) {
    let activeFields = {}
    const fieldsContainer = {
      [componentSchema.payloadName]: [],
    }
    //check if all fields contain input value
    componentPayload.forEach((fieldSet, index) => {
      let isPopulated = false
      for (const fieldName in fieldSet) {
        if (fieldSet[fieldName] !== '') {
          isPopulated = true
        } else {
          isPopulated = false
        }
      }
      isPopulated && fieldsContainer[componentSchema.payloadName].push(fieldSet)
    })

    //check if any fieldSets were added
    fieldsContainer[componentSchema.payloadName].length > 0
      ? (activeFields = {
          ...activeFields,
          ...fieldsContainer,
        })
      : (activeFields = {
          ...activeFields,
        })

    return activeFields
  }
}
function getAppFields(componentPayload, componentSchema) {
  if (!componentSchema.multiple) {
    let activeFields = {}
    if (componentSchema.name === COMPONENT_DEFINITIONS.networking.name) {
      activeFields = getNetworkFields(componentPayload, componentSchema)
    } else {
      for (const fieldName in componentPayload) {
        if (fieldName === 'encrypt' && componentPayload[fieldName] === '') {
          activeFields[fieldName] = false
        }
        if (componentPayload[fieldName] !== '') {
          if (fieldName === 'port') {
            activeFields[fieldName] = parseInt(componentPayload[fieldName])
          } else {
            activeFields[fieldName] = componentPayload[fieldName]
          }
        }
      }
    }
    if (componentSchema.name === COMPONENT_DEFINITIONS.dependency.name) {
      //  convert string into Array
      const array = componentPayload.dependsOn
        .replace(/\s/g, '')
        .split(',')
        .filter(Boolean)
      activeFields.dependsOn = array
    }
    return activeFields
  }
}
function getPolicyFields(componentPayload, componentSchema) {
  const activeFields = {}
  activeFields.resources = { general: {} }
  switch (componentSchema.name) {
    case COMPONENT_DEFINITIONS.policyConfig.name: //config component
      activeFields.shipaFramework = componentPayload.policyName
      activeFields.name = componentPayload.policyName
      activeFields.resources.general.setup = {
        default:
          componentPayload.default === null ||
          componentPayload.default === false
            ? false
            : true,
        public: componentPayload.public,
      }
      break
    case COMPONENT_DEFINITIONS.plan.name: //plan component
      activeFields.resources.general.plan = {
        name: componentPayload.plan,
      }
      break
    case COMPONENT_DEFINITIONS.security.name: //security component
      activeFields.resources.general.security = {
        disableScan:
          componentPayload.ignoreComponents === '' ||
          componentPayload.ignoreCves === ''
            ? false
            : true,
        ignoreComponents: componentPayload.ignoreComponents //convert string into array
          .replace(/\s/g, '')
          .split(',')
          .filter(Boolean),
        ignoreCves: componentPayload.ignoreCVEs //convert string into array
          .replace(/\s/g, '')
          .split(',')
          .filter(Boolean),
      }
      break
    case COMPONENT_DEFINITIONS.access.name: //access component
      activeFields.resources.general.access = {
        append: componentPayload.teams //convert string into array
          .replace(/\s/g, '')
          .split(',')
          .filter(Boolean),
      }
      break
    case COMPONENT_DEFINITIONS.containerPolicy.name: //containerPolicy component
      activeFields.resources.general.containerPolicy = {
        allowedHosts: componentPayload.allowedRegistries //convert string into array
          .replace(/\s/g, '')
          .split(',')
          .filter(Boolean),
      }
      break
    case COMPONENT_DEFINITIONS.policyAutoScale.name: //policyAutoScale component
      activeFields.resources.general.podAutoScaler = {
        minReplicas: parseInt(componentPayload.minReplicas),
        maxReplicas: parseInt(componentPayload.maxReplicas),
        targetCPUUtilizationPercentage: parseInt(
          componentPayload.targetCpuUtilization,
        ),
        disableAppOverride: true,
      }
      break
    case COMPONENT_DEFINITIONS.cnameControl.name: //cnameControl component
      activeFields.resources.general.domainPolicy = {
        allowedCnames: componentPayload.allowedCnames //convert string into array
          .replace(/\s/g, '')
          .split(',')
          .filter(Boolean),
      }
      break
    case COMPONENT_DEFINITIONS.policyNetworking.name: //policyNetworking component
      activeFields.resources.general.networkPolicy = getNetworkFields(
        componentPayload,
        componentSchema,
      )
      break
    default:
      break
  }

  return activeFields
}
function getNetworkFields(networkComponentPayload, networkComponentSchema) {
  let activeFields = { 'network-policy': { restart_app: true } }
  for (const fieldName in networkComponentPayload) {
    if (
      networkComponentPayload[fieldName].policy_mode ===
      NETWORK_POLICY_OPTIONS.custom
    ) {
      const fields = {
        policy_mode: 'allow-custom-rules-only',
        custom_rules: [{ enabled: true }],
      }
      for (const itemName in networkComponentPayload[fieldName]) {
        if (itemName !== 'policy_mode') {
          fields.custom_rules[0][itemName] =
            networkComponentPayload[fieldName][itemName]
          if (
            (itemName === 'allowed_apps' ||
              itemName === 'allowed_frameworks') &&
            networkComponentPayload[fieldName][itemName]
          ) {
            //  convert string into Array
            const array = networkComponentPayload[fieldName][itemName]
              .replace(/\s/g, '')
              .split(',')
              .filter(Boolean)
            fields.custom_rules[0][itemName] = array
          }
          if (
            itemName === 'ports' &&
            networkComponentPayload[fieldName][itemName]
          ) {
            const ports = networkComponentPayload[fieldName][itemName]
              .replace(/\s/g, '')
              .split(',')
              .filter(Boolean)
            const array = []
            ports.forEach(port => {
              array.push({
                protocol: 'TCP',
                port: parseInt(port),
              })
            })
            fields.custom_rules[0][itemName] = array
          }
        }
      }
      activeFields['network-policy'][fieldName] = { ...fields }
    } else if (fieldName !== 'restart_app') {
      //only get policy_mode
      const policy = {
        policy_mode: networkComponentPayload[fieldName].policy_mode,
      }
      activeFields['network-policy'][fieldName] = { ...policy }
    }
  }
  return activeFields
}
function getActiveFields(componentPayload, componentSchema, type) {
  let activeFields = {}
  if (type === APP_BLOCK_TYPES.appDeployment) {
    activeFields = {
      ...activeFields,
      ...getAppFields(componentPayload, componentSchema),
      ...getMultipleAppFields(componentPayload, componentSchema),
    }
  }
  if (type === APP_BLOCK_TYPES.framework) {
    activeFields = {
      ...getPolicyFields(componentPayload, componentSchema),
    }
  }

  return activeFields
}

export function updatePayload(definitions, setPayload) {
  const pl = definitions
    ? definitions.type === APP_BLOCK_TYPES.appDeployment
      ? { apps: [] }
      : definitions.type === APP_BLOCK_TYPES.framework
      ? { frameworks: [] }
      : null
    : null

  for (const name in definitions) {
    var fields = {}
    const definition = definitions[name]
    //skip the fields `type: policy` and `type: appDeployments`
    if (
      definition === APP_BLOCK_TYPES.framework ||
      definition === APP_BLOCK_TYPES.appDeployment
    ) {
      continue
    }
    for (const compDef in COMPONENT_DEFINITIONS) {
      if (
        definition.hasOwnProperty(COMPONENT_DEFINITIONS[compDef].payloadName)
      ) {
        const componentPayload =
          definition[COMPONENT_DEFINITIONS[compDef].payloadName]
        const componentSchema = COMPONENT_DEFINITIONS[compDef]
        const afields = getActiveFields(
          componentPayload,
          componentSchema,
          definitions.type,
        )

        if (definitions.type === APP_BLOCK_TYPES.appDeployment) {
          fields = {
            ...fields,
            ...afields,
          }
        }
        if (definitions.type === APP_BLOCK_TYPES.framework) {
          const { general } = afields.resources
          const prev_general = fields.resources ? fields.resources.general : {}
          fields = {
            ...afields,
            resources: {
              general: {
                ...general,
                ...prev_general,
              },
            },
          }
        }
      }
    }

    if (definitions.type === APP_BLOCK_TYPES.appDeployment) {
      pl.apps.push(fields)
    }
    if (definitions.type === APP_BLOCK_TYPES.framework) {
      pl.frameworks.push(fields)
    }
  }

  if (definitions.type === APP_BLOCK_TYPES.appDeployment) {
    //add temp values
    pl.apps.forEach(item => {
      if (!item.hasOwnProperty('norestart')) {
        item.norestart = true
      }
      if (!item.hasOwnProperty('private')) {
        item.private = true
      }
    })
  }
  setPayload(pl)
}
export async function generatePayload(definitions, payload, setPayload) {
  const data = { ...payload }

  const postURL = data.frameworks
    ? REQUEST_URLS.POST.frameworksGen
    : REQUEST_URLS.POST.appsGen

  const requestOptions = {
    method: 'POST',
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  }

  try {
    const res = await fetch(postURL, requestOptions)
    if (res.status !== 200) {
      alert('Bad request!')
      return
    }
    const resJson = await res.json()

    const url = window.URL.createObjectURL(new Blob([resJson.file.content]))
    const link = document.createElement('a')
    link.href = url
    link.setAttribute('download', resJson.file.name)
    document.body.appendChild(link)
    link.click()
    link.parentNode.removeChild(link)
    //reset payload data
    setPayload(
      definitions
        ? definitions.type === APP_BLOCK_TYPES.appDeployment
          ? { apps: [] }
          : definitions.type === APP_BLOCK_TYPES.framework
          ? { frameworks: [] }
          : null
        : null,
    )
  } catch (e) {
    console.log(e)
  }
}
