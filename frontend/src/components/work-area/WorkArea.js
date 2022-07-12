import React, { useContext, useState } from 'react'
import { AppContext } from '../../utils/GlobalContext'
import { Button } from 'antd'
import { PlusCircleFilled } from '@ant-design/icons'
import { useDndMonitor } from '@dnd-kit/core'
import Droppable from '../dnd/Droppable'
import Block from '../operational/AppBlock'
import cloneDeep from 'lodash.clonedeep'
import { COMPONENT_DEFINITIONS, DND_TYPES } from '../../utils/constants'
import GeneratePopUp from '../GeneratePopUp'
import { updatePayload } from '../../utils/payload'
import AddDefinitionDropdown from './AddDefinitionDropdown'
import ConfigInfoPopup from '../header/ConfigInfoPopup'

//this needs to be in a utility file
export function addComponent(
  definitions,
  setDefinitions,
  appDefinitionId,
  componentName,
) {
  const defs = { ...definitions }
  const compToAdd = COMPONENT_DEFINITIONS[componentName]
  const payloadDefinition = cloneDeep(compToAdd.payloadDefinition)
  if (!defs[appDefinitionId].hasOwnProperty(componentName)) {
    if (compToAdd.multiple) {
      defs[appDefinitionId][componentName] = [payloadDefinition]
    } else {
      defs[appDefinitionId][componentName] = payloadDefinition
    }
  } else {
    if (compToAdd.multiple) {
      defs[appDefinitionId][componentName] = [
        ...defs[appDefinitionId][componentName],
        payloadDefinition,
      ]
    }
  }
  setDefinitions(defs)
}

//should be in a util file
export function removeComponent(
  definitions,
  setDefinitions,
  definitionId,
  componentDefinition,
  elementIndex,
) {
  const defs = { ...definitions }
  //mandatory components can not be removed
  if (!componentDefinition.mandatory) {
    if (!componentDefinition.multiple) {
      if (defs[definitionId].hasOwnProperty(componentDefinition.payloadName)) {
        delete defs[definitionId][componentDefinition.payloadName]
      }
    } else {
      const arr = defs[definitionId][componentDefinition.payloadName].filter(
        (item, index) => index !== elementIndex,
      )
      defs[definitionId][componentDefinition.payloadName] = [...arr]
    }
  }
  setDefinitions(defs)
}

//should be in a util file
export function removeAppBlock(definitions, setDefinitions, definitionId) {
  const defs = { ...definitions }
  delete defs[definitionId]
  if (Object.keys(defs).length === 1 && defs.type) {
    delete defs.type
  }
  setDefinitions(defs)
}

function Container({ children }) {
  return <div className="bg-slate-500">{children}</div>
}

export default function WorkArea({ configInfoPopup, setConfigInfoPopup }) {
  const {
    definitions,
    setDefinitions,
    setPayload,
    validationArray,
    setAreFieldsValidated,
    userData,
  } = useContext(AppContext)

  //we want to create our component when drag has ended but by that time, info on the drag variables is lost. So we store them.
  const [dropTarget, setDropTarget] = useState(null)
  const [dragObject, setDragObject] = useState(null)
  const [genPopup, setGenPopup] = useState(false)

  const addDefinitionClasses = `h-14 text-lg fixed top-18 left-68 rounded-md`
  useDndMonitor({
    onDragStart(event) {
      setDragObject(event.active)
    },
    onDragMove(event) {},
    onDragOver(event) {
      setDropTarget(event.over)
    },
    onDragEnd(event) {
      if (
        dropTarget &&
        dropTarget.data.current.supportedTypes.includes(
          dragObject.data.current.type,
        )
      ) {
        if (
          dragObject.data.current.type === DND_TYPES.appComponent ||
          dragObject.data.current.type === DND_TYPES.policyComponent
        ) {
          addComponent(
            definitions,
            setDefinitions,
            dropTarget.id,
            dragObject.id,
          )
        }
      }
      setDropTarget(null)
      setDragObject(null)
    },
    onDragCancel(event) {
      setDropTarget(null)
      setDragObject(null)
    },
  })

  return (
    <Droppable disabled={false}>
      <Container>
        {/* render definition blocks */}
        {Object.keys(definitions).map(id => {
          if (id === 'type') {
            return <div key={id}></div>
          }
          if (definitions.type) {
            return <Block type={definitions.type} id={id} key={id} />
          }
          return <div key={id}></div>
        })}

        <div className={addDefinitionClasses}>
          <AddDefinitionDropdown />
        </div>
        <Button
          className="h-14 text-xl fixed bottom-0 right-0 mb-6 mr-4 rounded-xl"
          type="dashed"
          size="large"
          disabled={
            Object.keys(definitions).length > 0 && userData ? false : true
          }
          icon={<PlusCircleFilled className="text-blue-500" />}
          onClick={() => {
            if (validationArray.length > 0) {
              setAreFieldsValidated(false)
            } else {
              updatePayload(definitions, setPayload)
              setGenPopup(true)
            }
          }}
        >
          Generate
        </Button>
        {genPopup && <GeneratePopUp setGenPopup={setGenPopup} />}
        {configInfoPopup && (
          <ConfigInfoPopup setConfigInfoPopup={setConfigInfoPopup} />
        )}
      </Container>
    </Droppable>
  )
}
