import React, { useContext } from 'react'
import { AppContext } from '../../utils/GlobalContext'
import {
  COMPONENT_DEFINITIONS,
  FORM_ELEMENTS,
  NESTED_TYPES,
} from '../../utils/constants'
import FormInput from './FormInput'
import FormSelect from './FormSelect'
import useCollapse from 'react-collapsed'

export default function FormBody({
  componentDefinition,
  definitionId,
  type,
  elementIndex,
}) {
  const { definitions } = useContext(AppContext)

  const formSchema = componentDefinition.formSchema

  const elementIndexProp =
    elementIndex >= 0
      ? {
          elementIndex: elementIndex,
        }
      : {}

  const form = Object.keys(formSchema).map((fieldName, index) => {
    const fieldSchema = formSchema[fieldName]
    let networkPolicy = false
    if (componentDefinition.name === COMPONENT_DEFINITIONS.networking.name) {
      networkPolicy =
        definitions[definitionId][componentDefinition.payloadName][
          fieldSchema.payloadName
        ].policy_mode
    }

    if (fieldSchema.hasOwnProperty('nestedFields')) {
      const { getCollapseProps, getToggleProps, isExpanded } = useCollapse()
      const nestedFieldElements = Object.keys(fieldSchema.nestedFields).map(
        (nestedFieldName, Index) => {
          return (
            <div key={`${nestedFieldName}-${Index}`} className="my-2">
              <FormInput
                parentField={fieldSchema}
                fieldName={nestedFieldName}
                fieldSchema={fieldSchema.nestedFields[nestedFieldName]}
                componentDefinition={componentDefinition}
                payloadName={
                  fieldSchema.nestedFields[nestedFieldName].payloadName
                }
                definitionId={definitionId}
                type={type}
                {...elementIndexProp}
              />
            </div>
          )
        },
      )
      return (
        //parent container
        <div key={`${fieldName}-${index}`}>
          {fieldSchema.elementType === FORM_ELEMENTS.input && (
            <FormInput
              fieldName={fieldName}
              fieldSchema={fieldSchema}
              componentDefinition={componentDefinition}
              payloadName={fieldSchema.payloadName}
              definitionId={definitionId}
              type={type}
              toggleProps={getToggleProps}
              isExpanded={isExpanded}
              {...elementIndexProp}
            />
          )}
          {fieldSchema.elementType === FORM_ELEMENTS.select && (
            <FormSelect
              fieldName={fieldName}
              fieldSchema={fieldSchema}
              definitionId={definitionId}
              componentDefinition={componentDefinition}
              type={type}
              {...elementIndexProp}
            />
          )}
          {/* child container */}
          {fieldSchema.nestedType === NESTED_TYPES.collapsable && (
            <div {...getCollapseProps()}>{nestedFieldElements}</div>
          )}
          {fieldSchema.nestedType === NESTED_TYPES.conditional &&
            //temporary solution for network component only
            networkPolicy === fieldSchema.options.custom && (
              <div>{nestedFieldElements}</div>
            )}
        </div>
      )
    } else {
      return (
        <div key={`${fieldName}-${index}`} className="my-2">
          {fieldSchema.elementType === FORM_ELEMENTS.select && (
            <FormSelect
              fieldName={fieldName}
              fieldSchema={fieldSchema}
              type={type}
              definitionId={definitionId}
              componentDefinition={componentDefinition}
              {...elementIndexProp}
            />
          )}
          {fieldSchema.elementType === FORM_ELEMENTS.input && (
            <FormInput
              fieldName={fieldName}
              fieldSchema={fieldSchema}
              componentDefinition={componentDefinition}
              payloadName={fieldSchema.payloadName}
              definitionId={definitionId}
              type={type}
              {...elementIndexProp}
            />
          )}
        </div>
      )
    }
  })
  return <div>{form}</div>
}
