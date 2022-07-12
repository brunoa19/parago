import { getNetworkPolicyOptions } from './compOptions'

export const ROUTER_PATHS = {
  designPage: '/',
  detailsPage: '/details',
  searchPage: '/search',
  loginPage: '/login',
}

export const REQUEST_URLS = {
  // urls with '/' at the end need a value to be appended before the request
  POST: {
    login: `${process.env.REACT_APP_BACKEND_BASE_URL}/auth/login`,
    logout: `${process.env.REACT_APP_BACKEND_BASE_URL}/auth/logout`,
    frameworksGen: `${process.env.REACT_APP_BACKEND_BASE_URL}/shipa-gen/frameworks`,
    appsGen: `${process.env.REACT_APP_BACKEND_BASE_URL}/shipa-gen/apps`,
    save: `${process.env.REACT_APP_BACKEND_BASE_URL}/config`,
    update: `${process.env.REACT_APP_BACKEND_BASE_URL}/config/`,
  },
  GET: {
    users: `${process.env.REACT_APP_BACKEND_BASE_URL}/auth/user`,
    shipaCloud: `${process.env.REACT_APP_BACKEND_BASE_URL}/shipa-server/`,
    search: `${process.env.REACT_APP_BACKEND_BASE_URL}/config/search`,
    id: `${process.env.REACT_APP_BACKEND_BASE_URL}/config/`,
    ownedConfigs: `${process.env.REACT_APP_BACKEND_BASE_URL}/config`,
    publicConfigs: `${process.env.REACT_APP_BACKEND_BASE_URL}/config/list/public`,
    orgConfigs: `${process.env.REACT_APP_BACKEND_BASE_URL}/config/list/organization`,
    privateConfigs: `${process.env.REACT_APP_BACKEND_BASE_URL}/config/list/private`,
    clone: `${process.env.REACT_APP_BACKEND_BASE_URL}/config/clone/`,
    metrics: `${process.env.REACT_APP_BACKEND_BASE_URL}/config/`,
  },
  DELETE: {
    delete: `${process.env.REACT_APP_BACKEND_BASE_URL}/config/`,
  },
}

export const LOCAL_STORAGE = {
  loadedConfig: 'loadedConfig',
  definitions: 'definitions',
  authToken: 'authToken',
}

export const POP_UP_TYPES = {
  save: 'save',
  update: 'update',
}

export const APP_BLOCK_TYPES = {
  appDeployment: 'appDeployment',
  framework: 'framework',
}

export const ACCESS_TYPES = {
  owned: 'owned',
  public: 'public',
  private: 'private',
  organization: 'organization',
}

export const FORM_ELEMENTS = {
  input: 'input',
  select: 'select',
}

export const NESTED_TYPES = {
  conditional: 'conditional',
  collapsable: 'collapsable',
}

export const NETWORK_POLICY_OPTIONS = {
  allowAll: 'allow_all',
  denyAll: 'deny_all',
  custom: 'custom',
}

export const PROVIDERS = {
  Terraform: 'terraform',
  Crossplane: 'crossplane',
  Pulumi: 'pulumi',
  Ansible: 'ansible',
  'Github Actions': 'github',
  CloudFormation: 'cloudformation',
  Gitlab: 'gitlab',
  'Helm chart': 'helm_chart'
}

export const DND_TYPES = {
  block: 'Block',
  appComponent: 'appComponent',
  policyComponent: 'policyComponent',
  workArea: 'workArea',
}

export const COMPONENT_DEFINITIONS = {
  config: {
    name: 'Config',
    payloadName: 'config',
    componentType: DND_TYPES.appComponent,
    multiple: false,
    mandatory: true,
    formSchema: {
      'Application Name': {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: true,
        payloadName: 'appName',
      },
      Framework: {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: true,
        required: true,
        payloadName: 'framework',
      },
      Team: {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: true,
        required: true,
        payloadName: 'team',
      },
      Tags: {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: false,
        payloadName: 'tags',
      },
    },
    payloadDefinition: {
      appName: '',
      framework: '',
      team: '',
      tags: '',
    },
  },
  deployment: {
    name: 'Deployment',
    payloadName: 'deployment',
    componentType: DND_TYPES.appComponent,
    multiple: false,
    mandatory: true,
    formSchema: {
      Image: {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: true,
        payloadName: 'image',
      },
      Plan: {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: true,
        required: false,
        payloadName: 'plan',
      },
      'Private Registry': {
        elementType: FORM_ELEMENTS.input,
        type: 'checkbox',
        nestedType: NESTED_TYPES.collapsable,
        nestedFields: {
          Secret: {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'registrySecret',
          },
          User: {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'registryUser',
          },
        },
        required: false,
      },
      'Custom Port': {
        elementType: FORM_ELEMENTS.input,
        type: 'checkbox',
        nestedType: NESTED_TYPES.collapsable,
        nestedFields: {
          Port: {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'port',
          },
        },
        required: false,
      },
      Canary: {
        elementType: FORM_ELEMENTS.input,
        type: 'checkbox',
        nestedType: NESTED_TYPES.collapsable,
        nestedFields: {
          Steps: {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'steps',
          },
          'Step Interval': {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'stepInterval',
          },
          'Step Weight': {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'stepWeight',
          },
        },
        required: false,
      },
    },
    payloadDefinition: {
      image: '',
      plan: '',
      registryUser: '',
      registrySecret: '',
      port: '',
      steps: '',
      stepInterval: '',
      stepWeight: '',
    },
  },
  cname: {
    name: 'CNAME',
    payloadName: 'cname',
    componentType: DND_TYPES.appComponent,
    multiple: false,
    mandatory: false,
    formSchema: {
      CNAME: {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: true,
        payloadName: 'cname',
      },
      Encrypt: {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        type: 'checkbox',
        required: false,
        payloadName: 'encrypt',
      },
    },
    payloadDefinition: {
      cname: '',
      encrypt: false,
    },
  },
  dependency: {
    name: 'Dependency',
    payloadName: 'dependency',
    componentType: DND_TYPES.appComponent,
    multiple: false,
    mandatory: false,
    formSchema: {
      'Depends On': {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: false,
        payloadName: 'dependsOn',
      },
    },
    payloadDefinition: {
      dependsOn: '',
    },
  },
  envs: {
    name: 'Environment Variables',
    payloadName: 'envs',
    componentType: DND_TYPES.appComponent,
    multiple: true,
    mandatory: false,
    formSchema: {
      Name: {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: true,
        payloadName: 'name',
      },
      Value: {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: true,
        payloadName: 'value',
      },
    },
    payloadDefinition: {
      name: '',
      value: '',
    },
  },
  autoScale: {
    name: 'Auto Scale',
    payloadName: 'autoScale',
    componentType: DND_TYPES.appComponent,
    multiple: false,
    mandatory: false,
    formSchema: {
      'Max Replicas': {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: true,
        payloadName: 'maxReplicas',
      },
      'Min Replicas': {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: true,
        payloadName: 'minReplicas',
      },
      'Target CPU Utilization': {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: true,
        payloadName: 'targetCpuUtilization',
      },
    },
    payloadDefinition: {
      maxReplicas: '',
      minReplicas: '',
      targetCpuUtilization: '',
    },
  },
  networking: {
    name: 'Networking',
    payloadName: 'networking',
    componentType: DND_TYPES.appComponent,
    multiple: false,
    mandatory: false,
    formSchema: {
      Ingress: {
        elementType: FORM_ELEMENTS.select,
        required: false,
        payloadName: 'ingress',
        options: getNetworkPolicyOptions(),
        nestedType: NESTED_TYPES.conditional,
        nestedFields: {
          ID: {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'id',
          },
          Description: {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'description',
          },
          'Allowed apps': {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'allowed_apps',
          },
          'Allowed frameworks': {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'allowed_frameworks',
          },
          'Allowed ports': {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'ports',
          },
        },
      },
      Egress: {
        elementType: FORM_ELEMENTS.select,
        required: false,
        payloadName: 'egress',
        options: getNetworkPolicyOptions(),
        nestedType: NESTED_TYPES.conditional,
        nestedFields: {
          ID: {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'id',
          },
          Description: {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'description',
          },
          'Allowed apps': {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'allowed_apps',
          },
          'Allowed frameworks': {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'allowed_frameworks',
          },
          'Allowed ports': {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'ports',
          },
        },
      },
    },
    payloadDefinition: {
      ingress: {
        policy_mode: Object.entries(getNetworkPolicyOptions())[0][1],
        // id: '',
        // description: '',
        // allowed_apps: '',
        // allowed_frameworks: '',
        // ports: '',
      },
      egress: {
        policy_mode: Object.entries(getNetworkPolicyOptions())[0][1],
        // id: '',
        // description: '',
        // allowed_apps: '',
        // allowed_frameworks: '',
        // ports: '',
      },
      restart_app: true,
    },
  },
  volumes: {
    name: 'Volume',
    payloadName: 'volumes',
    componentType: DND_TYPES.appComponent,
    multiple: true,
    mandatory: false,
    formSchema: {
      Volume: {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: true,
        required: false,
        payloadName: 'name',
      },
      'Mount Path': {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: false,
        payloadName: 'mountPath',
      },
    },
    payloadDefinition: {
      name: '',
      mountPath: '',
    },
  },
  namespace: {
    name: 'Namespace',
    payloadName: 'namespace',
    componentType: DND_TYPES.policyComponent,
    multiple: false,
    mandatory: false,
    formSchema: {
      Name: {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: false,
        payloadName: 'kubernetesNamespace',
      },
    },
    payloadDefinition: {
      kubernetesNamespace: '',
    },
  },
  plan: {
    name: 'Resource Limit',
    payloadName: 'plan',
    componentType: DND_TYPES.policyComponent,
    multiple: false,
    mandatory: false,
    formSchema: {
      Plan: {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: false,
        payloadName: 'plan',
      },
    },
    payloadDefinition: {
      plan: '',
    },
  },
  security: {
    name: 'Security Scanning',
    payloadName: 'security',
    componentType: DND_TYPES.policyComponent,
    multiple: false,
    mandatory: false,
    formSchema: {
      'Enable Scan': {
        elementType: FORM_ELEMENTS.input,
        type: 'checkbox',
        nestedType: NESTED_TYPES.collapsable,
        nestedFields: {
          'Ignore Components': {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: true,
            payloadName: 'ignoreComponents',
          },
          'Ignore CVEs': {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: true,
            payloadName: 'ignoreCVEs',
          },
        },
        required: false,
      },
    },
    payloadDefinition: {
      ignoreComponents: '',
      ignoreCVEs: '',
    },
  },
  access: {
    name: 'RBAC',
    payloadName: 'access',
    componentType: DND_TYPES.policyComponent,
    multiple: false,
    mandatory: false,
    formSchema: {
      Teams: {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: true,
        payloadName: 'teams',
      },
    },
    payloadDefinition: {
      teams: '',
    },
  },
  containerPolicy: {
    name: 'Registry Control',
    payloadName: 'containerPolicy',
    componentType: DND_TYPES.policyComponent,
    multiple: false,
    mandatory: false,
    formSchema: {
      'Allowed Registries': {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: false,
        payloadName: 'allowedRegistries',
      },
    },
    payloadDefinition: {
      allowedRegistries: '',
    },
  },
  nodeSelectors: {
    name: 'Node Selector',
    payloadName: 'nodeSelectors',
    componentType: DND_TYPES.policyComponent,
    multiple: true,
    mandatory: false,
    formSchema: {
      Label: {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: false,
        payloadName: 'label',
      },
      Value: {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: false,
        payloadName: 'value',
      },
    },
    payloadDefinition: {
      label: '',
      value: '',
    },
  },
  cnameControl: {
    name: 'CNAME Control',
    payloadName: 'cnameControl',
    componentType: DND_TYPES.policyComponent,
    multiple: false,
    mandatory: false,
    formSchema: {
      'Allowed CNAMEs': {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: false,
        payloadName: 'allowedCnames',
      },
    },
    payloadDefinition: {
      allowedCnames: '',
    },
  },
  policyAutoScale: {
    name: 'Auto Scale',
    payloadName: 'policyAutoScale',
    componentType: DND_TYPES.policyComponent,
    multiple: false,
    mandatory: false,
    formSchema: {
      'Max Replicas': {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: true,
        payloadName: 'maxReplicas',
      },
      'Min Replicas': {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: true,
        payloadName: 'minReplicas',
      },
      'Target CPU Utilization': {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: true,
        payloadName: 'targetCpuUtilization',
      },
    },
    payloadDefinition: {
      maxReplicas: '',
      minReplicas: '',
      targetCpuUtilization: '',
    },
  },
  policyNetworking: {
    name: 'Networking',
    payloadName: 'policyNetworking',
    componentType: DND_TYPES.policyComponent,
    multiple: false,
    mandatory: false,
    formSchema: {
      Ingress: {
        elementType: FORM_ELEMENTS.select,
        required: false,
        payloadName: 'ingress',
        options: {
          ...NETWORK_POLICY_OPTIONS,
        },
        nestedType: NESTED_TYPES.conditional,
        nestedFields: {
          ID: {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'id',
          },
          Description: {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'description',
          },
          'Allowed apps': {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'allowed_apps',
          },
          'Allowed frameworks': {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'allowed_frameworks',
          },
          'Allowed ports': {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'ports',
          },
        },
      },
      Egress: {
        elementType: FORM_ELEMENTS.select,
        required: false,
        payloadName: 'egress',
        options: {
          ...NETWORK_POLICY_OPTIONS,
        },
        nestedType: NESTED_TYPES.conditional,
        nestedFields: {
          ID: {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'id',
          },
          Description: {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'description',
          },
          'Allowed apps': {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'allowed_apps',
          },
          'Allowed frameworks': {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'allowed_frameworks',
          },
          'Allowed ports': {
            elementType: FORM_ELEMENTS.input,
            optionsEndpoint: false,
            required: false,
            payloadName: 'ports',
          },
        },
      },
    },
    payloadDefinition: {
      ingress: {
        policy_mode: NETWORK_POLICY_OPTIONS.allowAll,
        // id: '',
        // description: '',
        // allowed_apps: '',
        // allowed_frameworks: '',
        // ports: '',
      },
      egress: {
        policy_mode: NETWORK_POLICY_OPTIONS.allowAll,
        // id: '',
        // description: '',
        // allowed_apps: '',
        // allowed_frameworks: '',
        // ports: '',
      },
      restart_app: true,
    },
  },
  policyConfig: {
    name: 'Config',
    payloadName: 'policyConfig',
    componentType: DND_TYPES.policyComponent,
    multiple: false,
    mandatory: true,
    formSchema: {
      'Policy Name': {
        elementType: FORM_ELEMENTS.input,
        optionsEndpoint: false,
        required: true,
        payloadName: 'policyName',
      },
      Public: {
        elementType: FORM_ELEMENTS.input,
        type: 'checkbox',
        required: false,
        payloadName: 'public',
      },
      Default: {
        elementType: FORM_ELEMENTS.input,
        type: 'checkbox',
        required: false,
        payloadName: 'default',
      },
    },
    payloadDefinition: {
      policyName: '',
      default: false,
      public: false,
    },
  },
}
