import React from 'react'
import { useDroppable } from '@dnd-kit/core'

export default function Droppable({
  children,
  disabled,
  data,
  id,
  resizeObserverConfig,
  supportedTypes,
}) {
  const { active, rect, isOver, node, over, setNodeRef } = useDroppable({
    //id must be unique
    id: id,
    disabled: disabled ? disabled : false,
    data: data
      ? { ...data, supportedTypes: supportedTypes ? supportedTypes : null }
      : { supportedTypes: supportedTypes ? supportedTypes : null },
    //maybe pass this as a prop too? (later)
    resizeObserverConfig: resizeObserverConfig
      ? resizeObserverConfig
      : {
          disabled: false,
          updateMeasurementsFor: [],
          timeout: 10,
        },
  })
  //attatches props to childern
  const childProps = React.Children.map(children, child => {
    return React.cloneElement(child, {
      droppableActive: active,
      droppableRect: rect,
      droppableIsOver: isOver,
      droppableNode: node,
      droppableOver: over,
    })
  })
  return <div ref={setNodeRef}>{childProps}</div>
}
