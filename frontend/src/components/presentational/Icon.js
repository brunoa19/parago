import React from 'react'
import {
  SettingOutlined,
  CodeSandboxOutlined,
  LinkOutlined,
  GlobalOutlined,
  RadarChartOutlined,
  ArrowsAltOutlined,
  DatabaseOutlined,
  DeploymentUnitOutlined,
  AliyunOutlined,
  FilterOutlined,
  SecurityScanOutlined,
  UserOutlined,
  FileTextOutlined,
  ForkOutlined,
} from '@ant-design/icons'
import { COMPONENT_DEFINITIONS } from '../../utils/constants'

export default function Icon({ componentName }) {
  switch (componentName) {
    case COMPONENT_DEFINITIONS.config.name:
      return <SettingOutlined />
    case COMPONENT_DEFINITIONS.deployment.name:
      return <CodeSandboxOutlined />
    case COMPONENT_DEFINITIONS.dependency.name:
      return <DeploymentUnitOutlined />
    case COMPONENT_DEFINITIONS.cname.name:
      return <LinkOutlined />
    case COMPONENT_DEFINITIONS.envs.name:
      return <GlobalOutlined />
    case COMPONENT_DEFINITIONS.autoScale.name:
      return <ArrowsAltOutlined />
    case COMPONENT_DEFINITIONS.networking.name:
      return <RadarChartOutlined />
    case COMPONENT_DEFINITIONS.volumes.name:
      return <DatabaseOutlined />
    case COMPONENT_DEFINITIONS.namespace.name:
      return <AliyunOutlined />
    case COMPONENT_DEFINITIONS.plan.name:
      return <FilterOutlined />
    case COMPONENT_DEFINITIONS.security.name:
      return <SecurityScanOutlined />
    case COMPONENT_DEFINITIONS.access.name:
      return <UserOutlined />
    case COMPONENT_DEFINITIONS.containerPolicy.name:
      return <FileTextOutlined />
    case COMPONENT_DEFINITIONS.nodeSelectors.name:
      return <ForkOutlined />
    case COMPONENT_DEFINITIONS.cnameControl.name:
      return <LinkOutlined />
    case COMPONENT_DEFINITIONS.policyConfig.name:
      return <SettingOutlined />
    default:
      return <></>
  }
}
