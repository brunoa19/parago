import React, { useCallback, useState } from 'react'
import { Input, AutoComplete } from 'antd'
import { searchByName } from '../../utils/configRequests'
import debounce from 'lodash.debounce'
import { useNavigate } from 'react-router-dom'
import { ROUTER_PATHS } from '../../utils/constants'
import shipaCorpLogo from '../../assets/images/shipa-corp-logo.png'

export function SearchInput() {
  const [options, setOptions] = useState([])
  const [inputValue, setInputValue] = useState('')
  const navigate = useNavigate()
  const debouncedHandleOptions = useCallback(debounce(handleOptions, 200), [])

  async function handleOptions(value) {
    if (value === '') {
      setOptions([])
    } else {
      try {
        const { data } = await searchByName(value)

        const opts = []
        data.results.forEach(result => {
          opts.push({
            value: result.id,
            label: (
              <div
                key={result.id}
                className="flex flex-row content-center border-t py-1 space-x-3 -mb-1"
              >
                <img
                  className="w-5 h-5 mt-1 rounded-full"
                  src={shipaCorpLogo}
                  alt="logo"
                />
                <div className="flex flex-col h-14 whitespace-normal">
                  <span className="font-semibold">
                    {result.userInfo.userOrg}/{result.name}
                  </span>
                  <div className="text-shipaGrey text-xs line-clamp-2">
                    {result.description}
                  </div>
                </div>
              </div>
            ),
          })
        })
        setOptions([...opts])
      } catch (e) {
        console.log(e)
      }
    }
  }

  function handleSearch(value) {
    setInputValue(value)
    debouncedHandleOptions(value)
  }
  function handleSelect(value) {
    //clear autoComplete input value
    setInputValue('')
    //update url param
    navigate(`${ROUTER_PATHS.detailsPage}/${value}`)
  }

  return (
    <AutoComplete
      dropdownMatchSelectWidth={'40vw'}
      style={{
        width: '40vw',
      }}
      options={options}
      onSelect={handleSelect}
      onSearch={handleSearch}
      value={inputValue}
    >
      <Input.Search size="large" className="-mt-2" placeholder="search" />
    </AutoComplete>
  )
}
