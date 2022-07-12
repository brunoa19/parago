import React, { useEffect, useState, useContext } from 'react'
import { useDraggable } from '@dnd-kit/core'
import { AppContext } from '../../utils/GlobalContext'
import { DND_TYPES } from '../../utils/constants'

export default function Draggable({
  children,
  id,
  data,
  disabled,
  initialTransform,
  shouldTranslate,
  type,
}) {
  const {
    active,
    attributes,
    listeners,
    setNodeRef,
    transform,
    isDragging,
    node,
    over,
  } = useDraggable({
    //id must be unique
    id: id,
    data: data
      ? { ...data, type: type ? type : null }
      : { type: type ? type : null },
    disabled: disabled ? disabled : false,
  })

  const [currentOrigin, setCurrentOrigin] = useState({ x: 0, y: 0 })
  const [lastTranform, setLastTransform] = useState(null)
  const { definitions, setDefinitions } = useContext(AppContext)

  useEffect(() => {
    if (initialTransform) {
      setCurrentOrigin({ x: initialTransform.x, y: initialTransform.y })
    }
  }, [])

  useEffect(() => {
    const defs = { ...definitions }
    if (type === DND_TYPES.block) {
      defs[id].position = { ...currentOrigin }
      setDefinitions(defs)
    }
  }, [currentOrigin])

  useEffect(() => {
    if (transform) {
      setLastTransform(transform)
    }
  }, [transform])

  useEffect(() => {
    if (!isDragging && lastTranform && shouldTranslate) {
      //restrict movement across screen edges
      if (currentOrigin.x + lastTranform.x < 257) {
        if (currentOrigin.y + lastTranform.y > 65) {
          setCurrentOrigin({
            x: 256,
            y: currentOrigin.y + lastTranform.y,
          })
        } else {
          setCurrentOrigin({
            x: 256,
            y: 65,
          })
        }
      } else {
        if (currentOrigin.y + lastTranform.y > 65) {
          setCurrentOrigin({
            x: currentOrigin.x + lastTranform.x,
            y: currentOrigin.y + lastTranform.y,
          })
        } else {
          setCurrentOrigin({
            x: currentOrigin.x + lastTranform.x,
            y: 65,
          })
        }
      }
    }
  }, [isDragging])

  const style = isDragging
    ? {
        transform: `translate3d(${currentOrigin.x +
          transform.x}px, ${currentOrigin.y + transform.y}px, 0)`,
      }
    : currentOrigin && {
        transform: `translate3d(${currentOrigin.x}px, ${currentOrigin.y}px, 0)`,
      }

  const childProps = React.Children.map(children, child => {
    return React.cloneElement(child, {
      dragActive: active,
      dragIsDragging: isDragging,
      dragTransform: transform,
      dragNode: node,
      dragOver: over,
      dragType: type,
    })
  })

  return (
    <div
      ref={setNodeRef}
      style={{
        ...style,
        position: `${initialTransform ? 'absolute' : 'relative'}`,
      }}
      {...listeners}
      {...attributes}
    >
      {childProps}
    </div>
  )
}
