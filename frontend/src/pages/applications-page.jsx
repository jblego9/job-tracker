import { useEffect, useState } from "react";
import * as Request from "../request.js"
import ApplicationsList from "../components/applications-list.jsx";
import EditModal from "../components/edit-modal.jsx";

export default function ApplicationsPage() {
    const [apps, setApps] = useState([]);
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

        getApplications().then(setApps);
    }, [pageNumber, pageSize]);

    const handleEdit = (updated) => {
        setApps(prev => prev.map(app => app.Id === updated.Id ? updated : app));
        setAppToEdit(null);
    };

    const handleDelete = (deletedId) => {
        setApps(prev => prev.filter(app => app.Id !== deletedId));
        setAppToEdit(null);
    };

    return (
        <>
            {appToEdit && (
                <EditModal
                    app={appToEdit}
                    onEdit={handleEdit}
                    onDelete={handleDelete}
                    onCancel={() => setAppToEdit(null)}
                />
            )}
            <ApplicationsList apps={apps} onEdit={setAppToEdit}/>
        </>
    );
}

