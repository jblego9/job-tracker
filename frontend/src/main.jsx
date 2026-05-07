import { StrictMode } from "react"
import { createRoot } from "react-dom/client"
import { BrowserRouter, Routes, Route } from "react-router"
import "./styles.css"
import getTestData from "./test-data.js"
import Header from "./components/header.jsx"
import Sidebar from "./components/sidebar.jsx"
import ApplicationsList from "./components/applications-list.jsx"
import Summary from "./components/summary.jsx"

createRoot(document.getElementById("root")).render(
    <StrictMode>
        <BrowserRouter>
            <Header />
            <div className="page-body">
                <Sidebar />
                <main className="content">
                    <Routes>
                        <Route path="/" element={<ApplicationsList applications={getTestData()} />} />
                        <Route path="/summary" element={<Summary />} />
                    </Routes>
                </main>
            </div>
        </BrowserRouter>
    </StrictMode>
)
