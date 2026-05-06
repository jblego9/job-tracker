import { StrictMode } from "react"
import { createRoot } from "react-dom/client"
import "./styles.css"
import getTestData from "./test-data.js"
import ApplicationsList from "./components/applications-list.jsx"
import Header from "./components/header.jsx"

createRoot(document.getElementById("root")).render(
    <StrictMode>
        <Header />
        <ApplicationsList applications={getTestData()} />
    </StrictMode>
)
