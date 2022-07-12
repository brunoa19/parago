import React from 'react'
import { Routes, Route } from 'react-router-dom'
import routes from './routes'

export default function AppRoutes() {
  return (
    <Routes>
      {routes.map(([key, value]) => (
        <Route
          key={key}
          path={value.path}
          exact={value.exact}
          element={<value.component />}
        />
      ))}
    </Routes>
  )
}
