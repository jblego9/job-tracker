import { StrictMode } from "react"
import { createRoot } from "react-dom/client"
import { BrowserRouter, Routes, Route } from "react-router"
import "./styles.css"
import Header from "./components/header.jsx"
import Sidebar from "./components/sidebar.jsx"
import ApplicationsPage from "./pages/applications-page.jsx"
import SummaryPage from "./pages/summary-page.jsx"

createRoot(document.getElementById("root")).render(
    <StrictMode>
        <BrowserRouter>
            <Header />
            <div className="page-body">
                <Sidebar />
                <main className="content">
                    <Routes>
                        <Route path="/" element={<ApplicationsPage />} />
                        <Route path="/summary" element={<SummaryPage />} />
                    </Routes>
                </main>
            </div>
        </BrowserRouter>
    </StrictMode>
)

