import React, { useContext, useState } from 'react'
import { AppContext } from '../utils/GlobalContext'
import { APP_BLOCK_TYPES, PROVIDERS } from '../utils/constants'
import { Button } from 'antd'
import { CloseCircleOutlined } from '@ant-design/icons'
import { generatePayload } from '../utils/payload'

export default function GeneratePopUp({ setGenPopup }) {
  const { definitions, payload, setPayload } = useContext(AppContext)
  const [generateDisabled, setGenerateDisabled] = useState(true)

  function handleChange(event) {
    setGenerateDisabled(false)
    //add provider to payload
    const pl = { ...payload }
    if (!pl.hasOwnProperty('provider')) {
      pl.provider = event.target.value
    }
    setPayload(pl)
  }

  return (
    <div
      className="flex backdrop-blur-sm absolute"
      style={{ height: '100vh', width: '100vw', zIndex: 11 }}
    >
      <div className="flex absolute top-1/3 left-1/2 px-4 py-16 bg-white rounded-xl border border-dashed border-green-300 drop-shadow-md">
        <div className="absolute top-0 right-0 text-xl">
          <Button
            className="rounded-md"
            type="text"
            icon={<CloseCircleOutlined />}
            onClick={() => {
              setGenPopup(false)
              //reset payload data
              setPayload(
                definitions
                  ? definitions.type === APP_BLOCK_TYPES.appDeployment
                    ? { apps: [] }
                    : definitions.type === APP_BLOCK_TYPES.framework
                    ? { frameworks: [] }
                    : null
                  : null,
              )
            }}
          ></Button>
        </div>
        <div>
          <label className="pr-16 text-xl">Provider*</label>
          <select
            className="pr-12 bg-slate-100 rounded-md border border-slate-500 text-xl"
            onChange={handleChange}
            defaultValue={'DEFAULT'}
          >
            <option value="DEFAULT" disabled hidden>
              Choose...
            </option>
            {Object.keys(PROVIDERS).map((providerName, index) => {
              return (
                <option
                  value={PROVIDERS[providerName]}
                  key={`${providerName}-${index}`}
                >
                  {providerName}
                </option>
              )
            })}
          </select>
        </div>
        <div className="absolute bottom-0 right-0 pr-4 mb-2">
          <Button
            className="rounded-md"
            type="dashed"
            size="large"
            disabled={generateDisabled}
            onClick={() => {
              generatePayload(definitions, payload, setPayload)
              setGenPopup(false)
            }}
          >
            Generate
          </Button>
        </div>
      </div>
    </div>
  )
}
