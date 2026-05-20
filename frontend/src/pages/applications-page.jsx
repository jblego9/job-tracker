import { useEffect, useState } from "react";
import * as Request from "../request.js"
import ApplicationsList from "../components/applications-list.jsx";
import EditModal from "../components/edit-modal.jsx";

export default function ApplicationsPage() {
    const [applications, setApplications] = useState([]);
    const [appToEdit, setAppToEdit] = useState(null);
    const [pageNumber, setPageNumber] = useState(1);
    const [pageSize, setPageSize] = useState(25);

    useEffect(() => {
        const getPaginationParams = (firstParams) => {
            return `${Request.getParamsPrefix(firstParams)}pageNum=${pageNumber-1}&pageSize=${pageSize}`;
        }

        const getApplications = async () => {
            try {
                return await Request.getDataJson(`${Request.apiUrl}${getPaginationParams(true)}`);
            } catch (error) {
                console.error(error.message);
                return [];
            }
        };

        getApplications().then(setApplications);
    }, [pageNumber, pageSize]);

    return (
        <>
            {appToEdit && (
                <EditModal />
            )}
            <ApplicationsList apps={applications} onEdit={setAppToEdit}/>
        </>
    );
}

