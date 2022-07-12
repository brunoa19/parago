import React from 'react'
import AppRoutes from './routes/AppRoutes'
import { BrowserRouter } from 'react-router-dom'
import GlobalContext from './utils/GlobalContext'
import Header from './components/header/Header'

export default function App() {
  return (
    <BrowserRouter>
      <GlobalContext>
        <Header />
        <AppRoutes />
      </GlobalContext>
    </BrowserRouter>
  )
}
