import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './styles.css'
import getTestData from './test-data.js'
import ApplicationsList from './components/applications-list.jsx'

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <ApplicationsList applications={getTestData()} />
  </StrictMode>,
)
