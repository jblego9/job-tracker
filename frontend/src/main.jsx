import { StrictMode } from "react"
import { createRoot } from "react-dom/client"
import { BrowserRouter, Routes, Route } from "react-router"
import "./styles.css"
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
                        <Route path="/" element={<ApplicationsList applications={await getApplications()} />} />
                        <Route path="/summary" element={<Summary />} />
                    </Routes>
                </main>
            </div>
        </BrowserRouter>
    </StrictMode>
)

async function getJSON(url) {
    const response = await fetch(url, { method: "GET" });

    if (!response.ok)
        throw new Error(`Failed to get JSON data from '${url}'. Response status: ${response.status}`);

    return await response.json();
}

async function getApplications() {
    try {
        return await getJSON("api/v1/applications");
    } catch (error) {
        console.error(error.message);
        return [];
    }
}
