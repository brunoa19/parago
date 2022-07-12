import React, { useContext, useEffect, useState } from 'react'
import { Button, Layout, Switch, Pagination } from 'antd'
import shipaCorpLogo from '../../assets/images/shipa-corp-logo.png'
import { AppContext } from '../../utils/GlobalContext'
import {
  listOrgConfigs,
  listOwnedConfigs,
  listPrivateConfigs,
  listPublicConfigs,
} from '../../utils/configRequests'
import { useNavigate } from 'react-router-dom'
import { ACCESS_TYPES, ROUTER_PATHS } from '../../utils/constants'

function ConfigTile({ id, name, org, description, resourceType }) {
  const navigate = useNavigate()

  return (
    <div className="flex flex-col justify-between bg-white border border-black/30 shadow-md w-80 h-40 rounded-md m-3 p-2">
      <div className="flex flex-row justify-between space-x-2">
        <div className="flex flex-row items-center space-x-2">
          <img
            className="w-5 h-5 mt-1 rounded-full"
            src={shipaCorpLogo}
            alt="logo"
          />
          <span>
            {org}/{name}
          </span>
        </div>
        <span className="font-semibold">{resourceType}</span>
      </div>
      <div className="flex flex-row items-center m-2 pl-4 text-xs line-clamp-4 text-shipaGrey">
        <p>{description}</p>
      </div>
      <div className="flex flex-row justify-end m-1">
        <Button
          type="text"
          size="small"
          onClick={() => {
            //update url param
            navigate(`${ROUTER_PATHS.detailsPage}/${id}`)
          }}
        >
          <span className="underline text-xs">details ></span>
        </Button>
      </div>
    </div>
  )
}

export default function SearchPage() {
  const { userData } = useContext(AppContext)
  const [publicConfigs, setPublicConfigs] = useState(null)
  const [privateConfigs, setPrivateConfigs] = useState(null)
  const [orgConfigs, setOrgConfigs] = useState(null)
  const [ownedConfigs, setOwnedConfigs] = useState(null)
  const [filter, setFilter] = useState({})
  const [pagination, setPagination] = useState({ page: 1, pageSize: 10 })

  useEffect(() => {
    if (userData) {
      setFilter({
        owned: ACCESS_TYPES.owned,
        public: ACCESS_TYPES.public,
      })
    } else {
      setFilter({
        public: ACCESS_TYPES.public,
      })
    }
  }, [userData])

  useEffect(() => {
    async function fetchPublicConfigs() {
      const publicConfigs = await listPublicConfigs(
        pagination.page,
        pagination.pageSize,
      )
      if (publicConfigs.status === 200) {
        setPublicConfigs(publicConfigs.data.results)
      } else {
        setPublicConfigs(null)
      }
    }
    fetchPublicConfigs()
  }, [pagination])

  useEffect(() => {
    async function getConfigs(type) {
      switch (type) {
        case ACCESS_TYPES.owned:
          const ownedConfigs = await listOwnedConfigs()
          ownedConfigs.data && setOwnedConfigs(ownedConfigs.data.results)
          break
        case ACCESS_TYPES.public:
          const publicConfigs = await listPublicConfigs(
            pagination.page,
            pagination.pageSize,
          )
          console.log(publicConfigs)
          publicConfigs.data && setPublicConfigs(publicConfigs.data.results)
          break
        case ACCESS_TYPES.private:
          const privateConfigs = await listPrivateConfigs()
          privateConfigs.data && setPrivateConfigs(privateConfigs.data.results)
          break
        case ACCESS_TYPES.organization:
          const orgConfigs = await listOrgConfigs()
          orgConfigs.data && setOrgConfigs(orgConfigs.data.results)
          break
        default:
          break
      }
    }
    if (filter) {
      Object.keys(ACCESS_TYPES).map(type => {
        if (filter.hasOwnProperty(ACCESS_TYPES[type])) {
          getConfigs(type)
        }
        return true
      })
    }
  }, [filter])

  function updateFilter(value) {
    // if value exists in filter -> remove it,
    // else add it
    const newFilter = { ...filter }
    if (newFilter.hasOwnProperty(value)) {
      delete newFilter[value]
    } else {
      newFilter[value] = value
    }
    setFilter(newFilter)
  }

  return (
    <Layout style={{ backgroundColor: 'white' }}>
      <div
        id="parent-container"
        className="flex flex-row w-full mt-16 justify-between"
      >
        <div
          id="filter-container"
          className="flex flex-col w-fit h-fit m-5 space-y-2 p-4 drop-shadow-sm border border-black/30 bg-white rounded-md"
        >
          <span className="font-bold text-lg subpixel-antialiased">
            Filter configs
          </span>
          {Object.keys(ACCESS_TYPES).map((typeName, index) => {
            return (
              <div
                className="flex flex-row justify-between"
                key={`${typeName}-${index}`}
              >
                <span>{ACCESS_TYPES[typeName]}</span>
                <Switch
                  size="small"
                  className={`${
                    filter.hasOwnProperty(typeName)
                      ? 'bg-slate-500'
                      : 'bg-slate-100'
                  } ml-4`}
                  onChange={e => updateFilter(typeName)}
                  defaultChecked={
                    ACCESS_TYPES[typeName] === ACCESS_TYPES.public ||
                    ACCESS_TYPES[typeName] === ACCESS_TYPES.owned
                      ? true
                      : false
                  }
                  disabled={
                    ACCESS_TYPES[typeName] !== ACCESS_TYPES.public && !userData
                      ? true
                      : false
                  }
                />
              </div>
            )
          })}
        </div>
        <div id="tile-container" className="flex flex-col w-full p-5 mb-5">
          {Object.keys(filter).map(item => {
            return (
              <div key={`${item}`}>
                <div
                  className="flex flex-row justify-center flex-wrap mb-10"
                  id="configs-container"
                >
                  <div className="w-full bg-black/70 text-white rounded-md pl-5 text-lg font-semibold">
                    {filter[item].charAt(0).toUpperCase() +
                      filter[item].slice(1)}{' '}
                    configs
                  </div>
                  {filter[item] === ACCESS_TYPES.owned &&
                    ownedConfigs &&
                    ownedConfigs.map(config => {
                      return (
                        <ConfigTile
                          key={config.id}
                          id={config.id}
                          name={config.name}
                          org={config.org}
                          description={config.description}
                          resourceType={config.resourceType}
                        />
                      )
                    })}
                  {filter[item] === ACCESS_TYPES.public &&
                    publicConfigs &&
                    publicConfigs.map(config => {
                      return (
                        <ConfigTile
                          key={config.id}
                          id={config.id}
                          name={config.name}
                          org={config.org}
                          description={config.description}
                          resourceType={config.resourceType}
                        />
                      )
                    })}
                  {filter[item] === ACCESS_TYPES.private &&
                    privateConfigs &&
                    privateConfigs.map(config => {
                      return (
                        <ConfigTile
                          key={config.id}
                          id={config.id}
                          name={config.name}
                          org={config.org}
                          description={config.description}
                          resourceType={config.resourceType}
                        />
                      )
                    })}
                  {filter[item] === ACCESS_TYPES.organization &&
                    orgConfigs &&
                    orgConfigs.map(config => {
                      return (
                        <ConfigTile
                          key={config.id}
                          id={config.id}
                          name={config.name}
                          org={config.org}
                          description={config.description}
                          resourceType={config.resourceType}
                        />
                      )
                    })}
                </div>
                {filter[item] === ACCESS_TYPES.public && publicConfigs && (
                  <div
                    id="public-pagination-container"
                    className="flex flex-row m-5 justify-center"
                  >
                    <Pagination
                      onChange={value => {
                        setPagination({
                          page: value,
                          pageSize: pagination.pageSize,
                        })
                      }}
                      size="small"
                      hideOnSinglePage
                      current={pagination.page}
                      total={
                        publicConfigs &&
                        Math.ceil(publicConfigs.length / pagination.pageSize)
                      }
                    />
                  </div>
                )}
              </div>
            )
          })}
        </div>
      </div>
    </Layout>
  )
}
