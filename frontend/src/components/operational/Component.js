import React from 'react'
import Draggable from '../dnd/Draggable'
import ComponentPresentational from '../presentational/Component'
export default function Component({
  componentName,
  id,
  type,
  shouldTranslate,
}) {
  return (
    <Draggable id={id} type={type} shouldTranslate={shouldTranslate}>
      <ComponentPresentational componentName={componentName} />
    </Draggable>
  )
}
