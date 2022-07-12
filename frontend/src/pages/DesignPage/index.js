import React, { useState } from 'react'
import { Layout } from 'antd'
import {
  DndContext,
  DragOverlay,
  PointerSensor,
  useSensor,
  useSensors,
} from '@dnd-kit/core'
import Sidebar from '../../components/sidebar/Sidebar'
import WorkArea from '../../components/work-area/WorkArea'
import { COMPONENT_DEFINITIONS, DND_TYPES } from '../../utils/constants'
import ComponentPresentational from '../../components/presentational/Component'
import Header from '../../components/header/Header'

export default function DesignPage() {
  const [activeElement, setActiveElement] = useState(null)
  const [configInfoPopup, setConfigInfoPopup] = useState(false)

  function handleDragStart(event) {
    setActiveElement(event.active)
    console.log('drag started 🤚')
  }
  function handleDragEnd() {
    setActiveElement(null)
    console.log('drag ended 🖐')
  }
  const sensors = useSensors(
    useSensor(PointerSensor, {
      activationConstraint: {
        distance: 10,
      },
    }),
  )

  return (
    <Layout style={{ backgroundColor: 'white' }}>
      <Header setConfigInfoPopup={setConfigInfoPopup} />
      <DndContext
        sensors={sensors}
        onDragStart={handleDragStart}
        onDragEnd={handleDragEnd}
        onDragCancel={handleDragEnd}
      >
        <Layout>
          <Sidebar />
          <WorkArea
            configInfoPopup={configInfoPopup}
            setConfigInfoPopup={setConfigInfoPopup}
          />
          {activeElement &&
            activeElement.data.current.type !== DND_TYPES.block && (
              <DragOverlay>
                <ComponentPresentational
                  componentName={COMPONENT_DEFINITIONS[activeElement.id].name}
                />
              </DragOverlay>
            )}
        </Layout>
      </DndContext>
    </Layout>
  )
}
