import { StrictMode } from "react"
import { createRoot } from "react-dom/client"
import "./styles.css"
import getTestData from "./test-data.js"
import ApplicationsList from "./components/applications-list.jsx"
import Header from "./components/header.jsx"
import Sidebar from "./components/sidebar.jsx"

createRoot(document.getElementById("root")).render(
    <StrictMode>
        <Header />
        <div className="page-body">
            <Sidebar />
            <main className="content">
                <ApplicationsList applications={getTestData()} />
            </main>
        </div>
    </StrictMode>
)
