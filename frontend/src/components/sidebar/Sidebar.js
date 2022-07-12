import React, { useContext } from 'react'
import { Layout } from 'antd'
import Component from '../operational/Component'
import {
  APP_BLOCK_TYPES,
  COMPONENT_DEFINITIONS,
  DND_TYPES,
} from '../../utils/constants'
import { AppContext } from '../../utils/GlobalContext'

export default function Sidebar(props) {
  const { definitions } = useContext(AppContext)

  const { Sider } = Layout
  const COMP_DEFS = Object.keys(COMPONENT_DEFINITIONS)
    .sort()
    .reduce((obj, key) => {
      obj[key] = COMPONENT_DEFINITIONS[key]
      return obj
    }, {})

  const componentContainerClasses =
    'h-fit mt-18 bg-slate-700/5 rounded-md mx-2 pb-2 pt-4'

  return (
    <Sider
      className="h-screen fixed left-0 top-2 scrollbar-thin scrollbar-thumb-slate-700 scrollbar-track-slate-300 overflow-y-scroll overflow-x-hidden -mt-2"
      style={{ zIndex: 10, backgroundColor: '#ececec' }}
      width={256}
    >
      {definitions.type && definitions.type === APP_BLOCK_TYPES.appDeployment && (
        <div className={componentContainerClasses}>
          <p className="text-slate-600 subpixel-antialiased text-base mb-5 pl-5">
            APPLICATION
          </p>
          <div className="px-2">
            {Object.keys(COMP_DEFS).map(
              componentName =>
                COMPONENT_DEFINITIONS[componentName].mandatory === false &&
                COMPONENT_DEFINITIONS[componentName].componentType ===
                  DND_TYPES.appComponent && (
                  <Component
                    componentName={COMPONENT_DEFINITIONS[componentName].name}
                    id={componentName}
                    type={DND_TYPES.appComponent}
                    shouldTranslate={false}
                    key={componentName}
                  ></Component>
                ),
            )}
          </div>
        </div>
      )}
      {definitions.type && definitions.type === APP_BLOCK_TYPES.framework && (
        <div className={componentContainerClasses}>
          <p className="text-slate-600 subpixel-antialiased text-base mb-5 pl-5">
            POLICY
          </p>
          <div className="px-2">
            {Object.keys(COMP_DEFS).map(
              componentName =>
                COMPONENT_DEFINITIONS[componentName].mandatory === false &&
                COMPONENT_DEFINITIONS[componentName].componentType ===
                  DND_TYPES.policyComponent && (
                  <Component
                    componentName={COMPONENT_DEFINITIONS[componentName].name}
                    id={componentName}
                    type={DND_TYPES.policyComponent}
                    shouldTranslate={false}
                    key={componentName}
                  ></Component>
                ),
            )}
          </div>
        </div>
      )}
    </Sider>
  )
}
