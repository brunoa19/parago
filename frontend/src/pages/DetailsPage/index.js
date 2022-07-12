import React, { useContext, useEffect, useState } from 'react'
import { Button, Layout } from 'antd'
import { useNavigate, useParams } from 'react-router-dom'
import {
  cloneById,
  getConfigDetails,
  getMetrics,
  removeById,
} from '../../utils/configRequests'
import shipaCorpLogo from '../../assets/images/shipa-corp-logo.png'
import { AppContext } from '../../utils/GlobalContext'
import { LOCAL_STORAGE, ROUTER_PATHS } from '../../utils/constants'
import { setInLocalStorage } from '../../utils/localStorage'

export default function DetailsPage() {
  const { userData, definitions, setDefinitions } = useContext(AppContext)
  const [owner, setOwner] = useState(false)
  const [details, setDetails] = useState(null)
  const [metrics, setMetrics] = useState(null)
  const { configId } = useParams()
  const navigate = useNavigate()

  async function getDetails() {
    //load details against id via api
    try {
      const { data } = await getConfigDetails(configId)
      setDetails(data)
      if (userData && data.userInfo.userId === userData.userId) {
        setOwner(true)
      } else {
        setOwner(false)
      }
    } catch (e) {
      console.log(e)
    }
    //temp workaround to refresh ui
    setDefinitions({ ...definitions })
  }

  async function fetchMetrics() {
    //load details against id via api
    try {
      const { data } = await getMetrics(configId)
      if (data && data.metrics) {
        setMetrics(data.metrics)
      }
    } catch (e) {}
  }

  useEffect(() => {
    getDetails()
    fetchMetrics()
  }, [configId, userData])

  async function handleLoad() {
    if (!owner) {
      //calling clone for metrics
      try {
        await cloneById(configId)
      } catch (e) {
        console.log('error while cloning:', e)
      }
    }

    const newDefs = { ...details.definition }

    //reset user specific fields
    if (!owner) {
      Object.keys(newDefs).map(def => {
        if (def === 'type') return false
        if (newDefs[def].hasOwnProperty('volumes')) {
          newDefs[def].volumes.forEach(element => {
            element.name = ''
          })
        }
        if (newDefs[def].hasOwnProperty('config')) {
          newDefs[def].config.framework = ''
          newDefs[def].config.team = ''
        }
        if (newDefs[def].hasOwnProperty('deployment')) {
          newDefs[def].deployment.plan = ''
        }
        return true
      })
    }

    setDefinitions({ ...newDefs })
    //set current loaded config id in localstorage
    if (details) {
      setInLocalStorage(LOCAL_STORAGE.loadedConfig, details)
    }

    navigate(ROUTER_PATHS.designPage)
  }

  async function handleDelete() {
    const res = await removeById(configId, userData.orgName, userData.name)
    console.log('delete output:', res)
    navigate(ROUTER_PATHS.designPage)
  }

  return (
    <Layout className="bg-transparent">
      <div className="flex mt-16 mx-20">
        <div className="h-96 w-full flex flex-col">
          <div className="flex flex-row justify-end items-center space-x-2 p-4 w-full h-16">
            {/* clone button container */}
            <Button
              className="hover:text-white hover:bg-shipaOrange hover:border-shipaOrange"
              onClick={handleLoad}
              shape="round"
            >
              {details && (owner ? 'Edit' : 'Clone')}
            </Button>
            {owner && (
              <Button
                className="hover:text-white hover:bg-shipaOrange hover:border-shipaOrange"
                onClick={handleDelete}
                shape="round"
              >
                Delete
              </Button>
            )}
          </div>
          {details ? (
            <div className="flex flex-row w-full h-fit pr-4">
              {/* details container */}
              <div className="flex flex-row grow">
                {/* image and info container */}
                <div className="flex flex-col items-center w-40">
                  {/* image container */}
                  <img className="h-20 w-20" src={shipaCorpLogo} alt="logo" />
                </div>
                <div className="flex flex-col grow w-96">
                  {/* info container  */}
                  <span className="font-bold text-xl subpixel-antialiased">
                    {details.name}
                  </span>
                  <span className="font-medium text-sm subpixel-antialiased">
                    by:{' '}
                    <span className="underline">
                      {details.userInfo.userName}
                    </span>
                  </span>
                  <span className="font-normal text-white subpixel-antialiased mt-4 py-1 px-2 w-fit rounded-full bg-slate-900">
                    {details.userInfo.userOrg}
                  </span>
                  <p
                    className="break-words p-2 bg-slate-400/20 my-4 mr-6 rounded-md"
                    style={{ whiteSpace: 'pre-wrap' }}
                  >
                    {details.description}
                  </p>
                </div>
              </div>
              <div className="flex flex-col w-fit h-fit space-y-2 p-4 drop-shadow-md bg-white border border-shipaDarkBlue/20 rounded-md">
                {/* metrics container */}
                <span className="font-bold text-lg subpixel-antialiased">
                  Provider metrics
                </span>
                <div className="flex flex-row justify-between">
                  <span>clones today</span>
                  <span className="ml-4 px-2 w-fit rounded-full bg-slate-200">
                    {metrics && metrics.clone.perDay}
                  </span>
                </div>
                <div className="flex flex-row justify-between">
                  <span>clones this week</span>
                  <span className="ml-4 px-2 w-fit rounded-full bg-slate-200">
                    {metrics && metrics.clone.perWeek}
                  </span>
                </div>
                <div className="flex flex-row justify-between">
                  <span>clones this month</span>
                  <span className="ml-4 px-2 w-fit rounded-full bg-slate-200">
                    {metrics && metrics.clone.perMonth}
                  </span>
                </div>
                <div className="flex flex-row justify-between">
                  <span>clones this year</span>
                  <span className="ml-4 px-2 w-fit rounded-full bg-slate-200">
                    {metrics && metrics.clone.perYear}
                  </span>
                </div>
                <div className="flex flex-row justify-between">
                  <span>clones over all time</span>
                  <span className="ml-4 px-2 w-fit rounded-full bg-slate-200">
                    {metrics && metrics.clone.total}
                  </span>
                </div>
              </div>
            </div>
          ) : (
            <div className="flex flex-col w-full h-full justify-center items-center rounded-md bg-slate-400/20">
              <h1 className="font-medium text-sm subpixel-antialiased">
                Details not found
              </h1>
              <Button
                className="mt-10 font-medium text-sm subpixel-antialiased hover:text-white hover:bg-shipaOrange hover:border-shipaOrange"
                size="large"
                shape="round"
                onClick={() => {
                  navigate(ROUTER_PATHS.designPage)
                }}
              >
                Go back
              </Button>
            </div>
          )}
        </div>
      </div>
    </Layout>
  )
}
