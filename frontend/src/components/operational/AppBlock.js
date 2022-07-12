import React, { useContext, useEffect } from 'react'
import { AppContext } from '../../utils/GlobalContext'
import { APP_BLOCK_TYPES, COMPONENT_DEFINITIONS } from '../../utils/constants'
import { DND_TYPES } from '../../utils/constants'
import Draggable from '../dnd/Draggable'
import Droppable from '../dnd/Droppable'
import BlockPresentational from '../presentational/AppBlock'
import { addComponent } from '../work-area/WorkArea'

function DroppableBlock({ id, type }) {
  return (
    <Droppable
      supportedTypes={
        type === APP_BLOCK_TYPES.appDeployment
          ? [DND_TYPES.appComponent]
          : [DND_TYPES.policyComponent]
      }
      id={id}
      disabled={false}
      data={{ blockType: type }}
    >
      <BlockPresentational id={id} type={type} />
    </Droppable>
  )
}

export default function Block({ id, type }) {
  const { definitions, setDefinitions } = useContext(AppContext)

  const position = definitions[id].position
    ? { ...definitions[id].position }
    : {
        x: 250 + Math.random() * 200,
        y: 250 + Math.random() * 200,
      }

  useEffect(() => {
    //add position to appDeps
    const appDeps = { ...definitions }
    appDeps[id].position = { ...position }
    setDefinitions(appDeps)
  }, [])

  useEffect(() => {
    //add mandatory components
    for (const component in COMPONENT_DEFINITIONS) {
      if (COMPONENT_DEFINITIONS[component].mandatory) {
        //only add APPLICATION components when type is appDeployment
        if (
          COMPONENT_DEFINITIONS[component].componentType ===
            DND_TYPES.appComponent &&
          type === APP_BLOCK_TYPES.appDeployment
        ) {
          addComponent(definitions, setDefinitions, id, component)
        }
        //only add POLICY components when type is policy
        if (
          COMPONENT_DEFINITIONS[component].componentType ===
            DND_TYPES.policyComponent &&
          type === APP_BLOCK_TYPES.framework
        ) {
          addComponent(definitions, setDefinitions, id, component)
        }
      }
    }
  }, [])

  return (
    <Draggable
      initialTransform={position}
      id={id}
      disabled={false}
      shouldTranslate={true}
      type={DND_TYPES.block}
    >
      <DroppableBlock id={id} type={type} />
    </Draggable>
  )
}
