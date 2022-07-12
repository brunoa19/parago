import React, { useContext } from 'react'
import { AppContext } from '../../utils/GlobalContext'
import { COMPONENT_DEFINITIONS } from '../../utils/constants'
import ComponentTile from '../ComponentTile'
import { Button } from 'antd'
import { CloseCircleOutlined } from '@ant-design/icons'
import { removeAppBlock } from '../work-area/WorkArea'

function TileGroup({ definitionId, componentName, type }) {
  const { definitions, setDefinitions } = useContext(AppContext)
  const appComponent = definitions[definitionId][componentName]

  const elements = appComponent.map((itemName, index) => {
    return (
      <ComponentTile
        definitionId={definitionId}
        definitions={definitions}
        setDefinitions={setDefinitions}
        componentDefinition={COMPONENT_DEFINITIONS[componentName]}
        isFirst={index === 0 ? true : false}
        elementIndex={index}
        type={type}
        key={index}
      />
    )
  })
  return elements
}

export default function BlockPresentational({
  droppableIsOver,
  droppableIsDragging,
  id,
  type,
}) {
  const { definitions, setDefinitions } = useContext(AppContext)
  function handleRemove() {
    removeAppBlock(definitions, setDefinitions, id)
  }

  const tailwindClasses = `pb-5 border-2 border-solid rounded-md w-96 ${
    droppableIsOver
      ? 'bg-white/80 border-shipaGrey/50'
      : 'bg-white/60 border-shipaGrey/10'
  }
  ${droppableIsDragging ? 'bg-white' : 'bg-white/60'}`

  return (
    <div>
      {/* button to delete app block */}
      <div className="flex flex-row-reverse pb-0.5">
        <Button
          className="border-shipaDarkBlue/10 hover:border-shipaDarkBlue/40 border text-shipaDarkBlue/60 hover:text-shipaDarkBlue rounded-md"
          icon={<CloseCircleOutlined />}
          onClick={handleRemove}
        ></Button>
      </div>
      <div
        className={tailwindClasses}
        style={{
          minHeight: '200px',
        }}
      >
        {Object.keys(definitions[id]).map((componentName, index) => {
          if (componentName === 'position') {
            return null
          }
          if (COMPONENT_DEFINITIONS[componentName].multiple) {
            return (
              <TileGroup
                definitionId={id}
                componentName={componentName}
                type={type}
                key={`${componentName}-${index}`}
              />
            )
          } else {
            return (
              <ComponentTile
                definitionId={id}
                definitions={definitions}
                setDefinitions={setDefinitions}
                componentDefinition={COMPONENT_DEFINITIONS[componentName]}
                isFirst={true}
                type={type}
                key={`${componentName}-${index}`}
              />
            )
          }
        })}
      </div>
    </div>
  )
}
