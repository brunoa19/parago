import React from 'react'
import Icon from './Icon'

export default function ComponentPresentational({ componentName }) {
  return (
    <div className="hover:bg-shipaDarkBlue/10 flex items-center py-2 pl-3 my-2 rounded-full">
      <div className="mr-2 text-xl">
        {<Icon componentName={componentName} />}
      </div>
      <p className="subpixel-antialiased text-base">{componentName}</p>
    </div>
  )
}
