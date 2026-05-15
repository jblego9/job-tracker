import { useState, useRef } from "react";
import EditModal from "./edit-modal.jsx";
import AddModal from "./add-modal.jsx";

export default function ApplicationsList({ applications, applicationsUpdated, pageNumber, pageSize, onPageNumberChange, onPageSizeChange }) {
    const [appToEdit, setAppToEdit] = useState(null);
    const [addingApp, setAddingApp] = useState(false);

    const pageNumberRef = useRef(null);

    const updatePageNumber = (value) => {
        const validated = (value >= 1) ? value : 1; 
        onPageNumberChange(validated);
        pageNumberRef.current.value = validated;
    };

    const handlePageNumberChange = (e) => {
        updatePageNumber(Number(e.target.value));
    };

    const handlePageSizeChange = (e) => {
        onPageSizeChange(Number(e.target.value));
    };

    return (
        <>
            <div className="applications-control-bar">
                <div className="applications-control-bar__page-controls">
                    <div>
                        <label htmlFor="page-number">Page:</label>
                        <input
                            ref={pageNumberRef}
                            id="page-number"
                            name="page-number"
                            type="number"
                            defaultValue={1}
                            min={1}
                            onChange={handlePageNumberChange}
                        />
                        <button
                            onClick={() => updatePageNumber(pageNumber - 1)}
                        >
                            {"<"}
                        </button>
                        <button
                            onClick={() => updatePageNumber(pageNumber + 1)}
                        >
                            {">"}
                        </button>
                    </div>
                    <div>
                        <label htmlFor="page-size">Applications per page:</label>
                        <select
                            id="page-size"
                            name="page-size"
                            defaultValue={25}
                            onChange={handlePageSizeChange}
                        >
                            <option value={25}>25</option>
                            <option value={50}>50</option>
                            <option value={100}>100</option>
                        </select>
                    </div>
                </div>
                <button className="button" onClick={() => setAddingApp(true)}>Add Application</button>
                { addingApp && (
                    <AddModal applicationsUpdated={applicationsUpdated} onClose={() => setAddingApp(false)}/>
                )}
            </div>
            <ol className="applications-list">
                {applications.map(app => (
                    <li key={app.Id} className="application-card">
                        <div>
                            <div className="application-card__row application-card__row--main">
                                <p>{app.Company}</p>
                                <p>-</p>
                                <p>{app.Role}</p>
                            </div>
                            <div className="application-card__row application-card__row--details">
                                <p>{app.Location}</p>
                            </div>
                            <div className="application-card__row application-card__row--details application-card__row--spaced">
                                {hasDate(app.Date) && (
                                    <>
                                        <p>{app.Date}</p>
                                        <p>-</p>
                                    </>
                                )}
                                <p>{app.Status}</p>
                            </div>
                        </div>
                        <div>
                            <button className="button" onClick={() => setAppToEdit(app)}>Edit</button>
                        </div>
                    </li>
                ))}
            </ol>
            {appToEdit && (
                <EditModal app={appToEdit} applicationsUpdated={applicationsUpdated} onClose={() => setAppToEdit(null)}/>
            )}
        </>
    );
}

function hasDate(date) {
    return date != null && date != "" && date != undefined;
}
