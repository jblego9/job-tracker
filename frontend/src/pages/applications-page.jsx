import { useEffect, useState } from "react";
import * as Request from "../request.js";
import ApplicationsList from "../components/applications-list.jsx";

export default function ApplicationsPage() {
    const [applications, setApplications] = useState([]);
    const [updatedApplications, setUpdatedApplications] = useState(false);
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
    }, [updatedApplications, pageNumber, pageSize]);

    return (
        <ApplicationsList
            applications={applications}
            applicationsUpdated={setUpdatedApplications}
            pageNumber={pageNumber}
            pageSize={pageSize}
            onPageNumberChange={setPageNumber}
            onPageSizeChange={setPageSize}
        />
    );
}

