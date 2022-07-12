import React from 'react'
import {
  CaretDownOutlined,
  CaretUpOutlined,
  CloseCircleOutlined,
} from '@ant-design/icons'
import { Button } from 'antd'
import useCollapse from 'react-collapsed'
import FormBody from './formComponents/FormBody'
import Icon from './presentational/Icon'
import { removeComponent } from './work-area/WorkArea'

export default function ComponentTile({
  definitionId,
  definitions,
  setDefinitions,
  componentDefinition,
  isFirst,
  elementIndex,
  type,
}) {
  const { getCollapseProps, getToggleProps, isExpanded } = useCollapse()

  return (
    <div>
      {/* tile header */}
      {/* //render compact tile header */}
      <div
        className={`flex flex-row items-center ${
          !isFirst ? 'w-full px-2 py-2' : 'justify-between p-2 text-lg'
        } border-b border-shipaGrey/10 bg-shipaDarkBlue/5`}
        {...getToggleProps()}
      >
        <div className={`flex flex-row w-full justify-between`}>
          <div className="flex flex-row items-center">
            <Icon componentName={componentDefinition.name} />
            {!isFirst ? (
              <span />
            ) : (
              <h1 className="px-2">{componentDefinition.name}</h1>
            )}

            {isExpanded ? <CaretUpOutlined /> : <CaretDownOutlined />}
          </div>
          {/* add remove button */}
          {!componentDefinition.mandatory && (
            <Button
              type="text"
              size="small"
              icon={<CloseCircleOutlined />}
              onClick={() => {
                removeComponent(
                  definitions,
                  setDefinitions,
                  definitionId,
                  componentDefinition,
                  elementIndex,
                )
              }}
            ></Button>
          )}
        </div>
      </div>
      {/* tile body */}
      <form {...getCollapseProps()}>
        <div className="flex flex-col text-lg">
          <FormBody
            type={type}
            componentDefinition={componentDefinition}
            definitionId={definitionId}
            elementIndex={elementIndex}
          />
        </div>
      </form>
    </div>
  )
}
