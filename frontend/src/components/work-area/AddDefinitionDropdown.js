import React, { useContext } from 'react'
import { AppContext } from '../../utils/GlobalContext'
import { Button, Menu, Dropdown, Space } from 'antd'
import { PlusCircleFilled } from '@ant-design/icons'
import { v4 as uuidv4 } from 'uuid'
import { APP_BLOCK_TYPES } from '../../utils/constants'

export default function AddDefinitionDropdown() {
  const { definitions, setDefinitions } = useContext(AppContext)

  //TODO should be in a util file
  function addDeployment() {
    const id = uuidv4()
    setDefinitions({
      ...definitions,
      [id]: {},
      type: APP_BLOCK_TYPES.appDeployment,
    })
  }
  //TODO should be in a util file
  function addPolicy() {
    const id = uuidv4()
    setDefinitions({
      ...definitions,
      [id]: {},
      type: APP_BLOCK_TYPES.framework,
    })
  }

  const menu = (
    <Menu>
      <Menu.Item key={'addDeployment'}>
        <Button
          className="h-10 rounded-md"
          style={{ textAlign: 'left' }}
          block
          disabled={
            definitions.type &&
            definitions.type === APP_BLOCK_TYPES.framework &&
            true
          }
          type="text"
          onClick={addDeployment}
        >
          Deployment
        </Button>
      </Menu.Item>
      <Menu.Item key={'addPolicy'}>
        <Button
          className="h-10 rounded-md"
          style={{ textAlign: 'left' }}
          block
          disabled={
            definitions.type &&
            definitions.type === APP_BLOCK_TYPES.appDeployment &&
            true
          }
          type="text"
          onClick={addPolicy}
        >
          Policy
        </Button>
      </Menu.Item>
    </Menu>
  )

  return (
    <Dropdown
      overlay={menu}
      className="border-dashed border border-slate-200 rounded-xl p-3 hover:bg-slate-500/10"
    >
      <button onClick={e => e.preventDefault()}>
        <Space>
          <PlusCircleFilled className="text-green-500" />
          Add Definition
        </Space>
      </button>
    </Dropdown>
  )
}
