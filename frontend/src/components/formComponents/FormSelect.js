import React, { useContext } from 'react'
import { AppContext } from '../../utils/GlobalContext'

export default function FormSelect({
  fieldName,
  fieldSchema,
  componentDefinition,
  definitionId,
  elementIndex,
  type,
}) {
  const { definitions, setDefinitions } = useContext(AppContext)

  function handleChange(event) {
    const defs = { ...definitions }
    if (componentDefinition.multiple) {
      defs[definitionId][componentDefinition.payloadName][elementIndex][
        fieldSchema.payloadName
      ].policy_mode = event.target.value
    } else {
      defs[definitionId][componentDefinition.payloadName][
        fieldSchema.payloadName
      ].policy_mode = event.target.value
    }
    setDefinitions(defs)
  }

  const options = Object.keys(fieldSchema.options).map((name, index) => {
    return (
      <option value={fieldSchema.options[name]} key={`${name}-${index}`}>
        {name}
      </option>
    )
  })

  return (
    <div className="flex flex-row px-2 justify-between mt-2">
      <div>
        <label>{fieldSchema.required ? `${fieldName}*` : fieldName}</label>
      </div>
      <div className="flex flex-row">
        <select
          className="border-black border rounded-md py-0.5 w-50 h-8"
          onChange={handleChange}
          //if component is networking, set policy_mode as default value. Else set first value of field object
          defaultValue={
            componentDefinition.payloadName === 'policyNetworking' ||
            componentDefinition.payloadName === 'networking'
              ? definitions[definitionId][componentDefinition.payloadName][
                  fieldSchema.payloadName
                ].policy_mode
              : Object.entries(
                  definitions[definitionId][componentDefinition.payloadName][
                    fieldSchema.payloadName
                  ],
                )[0][1]
          }
        >
          {options}
        </select>
      </div>
    </div>
  )
}
