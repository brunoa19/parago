import React, { useContext, useState, useEffect } from 'react'
import { AppContext } from '../../utils/GlobalContext'
import { MinusCircleOutlined, PlusCircleOutlined } from '@ant-design/icons'
import { getFormOptions } from '../../utils/compOptions'
export default function FormInput({
  parentField,
  fieldName,
  fieldSchema,
  componentDefinition,
  definitionId,
  type,
  toggleProps,
  isExpanded,
  elementIndex,
}) {
  const {
    definitions,
    setDefinitions,
    userData,
    owner,
    validationArray,
    areFieldsValidated,
  } = useContext(AppContext)
  //initialize this state with the data from appDeployments TODO convert into simple if else
  const initialValue = componentDefinition.multiple
    ? parentField && parentField.payloadName
      ? definitions[definitionId][componentDefinition.payloadName][
          elementIndex
        ][parentField.payloadName][fieldSchema.payloadName]
      : definitions[definitionId][componentDefinition.payloadName][
          elementIndex
        ][fieldSchema.payloadName]
    : parentField && parentField.payloadName
    ? definitions[definitionId][componentDefinition.payloadName][
        parentField.payloadName
      ][fieldSchema.payloadName]
    : definitions[definitionId][componentDefinition.payloadName][
        fieldSchema.payloadName
      ]

  const [inputValue, setInputValue] = useState(initialValue)
  const [options, setOptions] = useState(null)

  useEffect(() => {
    //get options from api
    async function fetchOptions() {
      const res = await getFormOptions(fieldName)
      setOptions(res)
    }
    userData && fieldSchema.optionsEndpoint && fetchOptions()

    //unload user specific values on logout --TODO this always unloads user specific values.
    // if (!owner) {
    //   if (
    //     fieldSchema.payloadName === 'framework' ||
    //     fieldSchema.payloadName === 'team' ||
    //     fieldSchema.payloadName === 'plan' ||
    //     fieldSchema.payloadName === 'name'
    //   ) {
    //     setInputValue('')
    //   }
    // }
  }, [userData, owner])

  useEffect(() => {
    updateAppDefinitions()
    if (fieldSchema.required && !inputValue) {
      validationArray.push(fieldName)
    }
    if (fieldSchema.required && inputValue) {
      var index = validationArray.indexOf(fieldName)
      if (index !== -1) {
        validationArray.splice(index, 1)
      }
    }
  }, [inputValue])

  function handleChange(event) {
    if (fieldSchema.type === 'checkbox') {
      setInputValue(event.target.checked)
    } else {
      setInputValue(event.target.value)
    }
  }

  function updateAppDefinitions() {
    const defs = { ...definitions }
    const comp = defs[definitionId][componentDefinition.payloadName]
    if (componentDefinition.multiple) {
      parentField && parentField.payloadName //if comp is nested
        ? (comp[elementIndex][parentField.payloadName][
            fieldSchema.payloadName
          ] = inputValue)
        : (comp[elementIndex][fieldSchema.payloadName] = inputValue)
    } else {
      parentField && parentField.payloadName //if comp is nested
        ? (comp[parentField.payloadName][fieldSchema.payloadName] = inputValue)
        : (comp[fieldSchema.payloadName] = inputValue)
    }
    setDefinitions(defs)
  }

  const inputContainerStyle = `${
    fieldSchema.type === 'checkbox' ? 'pr-44' : ''
  } flex flex-col items-center`

  const inputStyle = `${
    fieldSchema.required
      ? !inputValue
        ? 'border-b-rose-500'
        : 'border-b-slate-700'
      : 'border-b-slate-700'
  } border-2 rounded-md`

  const selectStyle = `${
    fieldSchema.required
      ? !inputValue
        ? 'border-b-rose-500'
        : 'border-b-slate-700'
      : 'border-b-slate-700'
  } border-2 rounded-md ${
    fieldSchema.optionsEndpoint ? 'bg-sky-800/10' : 'border-b-slate-700'
  } rounded-md py-0.5 w-50 h-8`

  return (
    <div className="flex flex-row px-2 justify-between mt-2">
      <div>
        <label>{fieldSchema.required ? `${fieldName}*` : fieldName}</label>
      </div>

      <div className={inputContainerStyle}>
        {!areFieldsValidated && !inputValue && fieldSchema.required && (
          <div className={`flex flex-row w-full justify-end text-rose-600`}>
            <span>*required</span>
          </div>
        )}

        {toggleProps ? (
          isExpanded ? (
            <MinusCircleOutlined {...toggleProps()} />
          ) : (
            <PlusCircleOutlined {...toggleProps()} />
          )
        ) : userData && fieldSchema.optionsEndpoint ? (
          options ? (
            <select
              className={selectStyle}
              onChange={handleChange}
              defaultValue={initialValue ? initialValue : ''}
            >
              <option value="" disabled>
                Choose a {fieldName}
              </option>
              {options.map((option, index) => {
                return (
                  <option value={option.name} key={`${option.name}-${index}`}>
                    {option.name}
                  </option>
                )
              })}
            </select>
          ) : (
            <span>No {fieldName}s available</span>
          )
        ) : (
          <input
            className={inputStyle}
            onChange={handleChange}
            //this cant be done with tailwind for some reason
            style={
              fieldSchema.type === 'checkbox'
                ? { height: '22px', width: '22px' }
                : {}
            }
            value={inputValue}
            type={fieldSchema.type ? fieldSchema.type : 'text'}
            {...(inputValue === true ? { checked: true } : { checked: false })}
          />
        )}
      </div>
    </div>
  )
}
