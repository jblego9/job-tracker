import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './styles.css'
import ApplicationsTable from './components/applications-table.jsx'

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <ApplicationsTable applications={[
      {
        Id: 1,
        Date: "2026-05-04",
        Company: "Company1",
        Role: "Role1",
        Location: "Location1",
        Status: "REJECTED",
      },
      {
        Id: 2,
        Date: "2026-05-04",
        Company: "Company2",
        Role: "Role2",
        Location: "Location2",
        Status: "NO_RESPONSE",
      },
    ]} />
  </StrictMode>,
)
