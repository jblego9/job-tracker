import { useEffect, useState } from "react";
import * as Request from "../request.js"
import ApplicationsList from "../components/applications-list.jsx";
import EditModal from "../components/edit-modal.jsx";
import ControlBar from "../components/control-bar.jsx";

export default function ApplicationsPage() {
    const [apps, setApps] = useState([]);
    const [appToEdit, setAppToEdit] = useState(null);
    const [addingApp, setAddingApp] = useState(false);
    const [pageNumber, setPageNumber] = useState(1);
    const [pageSize, setPageSize] = useState(25);

    useEffect(() => {
        const getPaginationParams = (firstParams) => {
            return `${Request.getParamsPrefix(firstParams)}pageNum=${pageNumber-1}&pageSize=${pageSize}`;
        }

        const getApplications = async () => {
            try {
                const response = await Request.getData(`${Request.apiUrl}${getPaginationParams(true)}`);
                return await response.json();
            } catch (error) {
                console.error(error.message);
                return [];
            }
        };

        getApplications().then(setApps);
    }, [pageNumber, pageSize]);

    const handleAdd = (created) => {
        setApps(prev => [...prev, created]);
        setAddingApp(false);
    };

    const handleEdit = (updated) => {
        setApps(prev => prev.map(app => app.Id === updated.Id ? updated : app));
        setAppToEdit(null);
    };

    const handleDelete = (deletedId) => {
        setApps(prev => prev.filter(app => app.Id !== deletedId));
        setAppToEdit(null);
    };

    const handlePageNumberChange = (value) => {
        setPageNumber((value >= 1) ? value : 1);
    };

    const handlePageSizeChange = (value) => {
        setPageSize((value >= 1) ? value : 1);
    };

    return (
        <>
            <ControlBar
                pageNumber={pageNumber}
                onPageNumberChange={handlePageNumberChange}
                pageSize={pageSize}
                onPageSizeChange={handlePageSizeChange}
                addingApp={addingApp}
                onAddingApp={setAddingApp}
                onAdd={handleAdd}
            />
            <ApplicationsList apps={apps} onEdit={setAppToEdit}/>
            {appToEdit && (
                <EditModal
                    app={appToEdit}
                    onEdit={handleEdit}
                    onDelete={handleDelete}
                    onCancel={() => setAppToEdit(null)}
                />
            )}
        </>
    );
}

