import { useEffect, useRef, useState } from "react";
import * as Request from "../request.js";

export default function EditModal({ app, onEdit, onDelete, onCancel }) {
    const [formData, setFormData] = useState({
        Id: app.Id,
        Company: app.Company,
        Role: app.Role,
        Location: app.Location,
        Date: app.Date,
        Status: app.Status,
    });

    const dialogRef = useRef(null);
    const statusRef = useRef(null);

    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleConfirm = async (e) => {
        try {
            e.preventDefault();
            await Request.putData(Request.apiUrl, formData);
            onEdit(formData);
        } catch (error) {
            statusRef.current.innerText = "Failed to edit application";
            console.error(error);
        }
    };

    const handleDelete = async (e) => {
        try {
            e.preventDefault();
            await Request.deleteData(`${Request.apiUrl}/${app.Id}`);
            onDelete(formData.Id);
        } catch (error) {
            statusRef.current.innerText = "Failed to delete application";
            console.error(error);
        }
    };

    useEffect(() => {
        dialogRef.current.showModal();
    }, []);

    return (
        <dialog className="modal" ref={dialogRef} onClose={onCancel}>
            <form className="modal-form">
                <h1 className="no-bottom-margin">Edit Application</h1>
                <p className="edit-modal-application-id">Application ID: {app.Id}</p>

                <h2>Job Details</h2>
                <label htmlFor="company">Company</label>
                <input
                    type="text"
                    id="company"
                    name="Company"
                    autoComplete="organization"
                    defaultValue={app.Company}
                    onChange={handleChange}
                    required
                />

                <label htmlFor="role">Role</label>
                <input
                    type="text"
                    id="role"
                    name="Role"
                    defaultValue={app.Role}
                    onChange={handleChange}
                    required
                />

                <label htmlFor="location">Location</label>
                <input
                    type="text"
                    id="location"
                    name="Location"
                    defaultValue={app.Location}
                    onChange={handleChange}
                    required
                />

                <h2>Application Details</h2>
                <label htmlFor="date">Date</label>
                <input
                    type="date"
                    id="date"
                    name="Date"
                    defaultValue={app.Date}
                    onChange={handleChange}
                />

                <label htmlFor="status">Status</label>
                <select
                    id="status"
                    name="Status"
                    defaultValue={app.Status}
                    onChange={handleChange}
                    required
                >
                    <option value="AWAITING_RESPONSE">Awaiting Response</option>
                    <option value="NO_RESPONSE">No Response</option>
                    <option value="AWAITING_INTERVIEW">Awaiting Interview</option>
                    <option value="REJECTED">Rejected</option>
                    <option value="ACCEPTED">Accepted</option>
                </select>

                <div className="modal-form__actions">
                    <button
                        className="modal-form__button modal-form__button--delete"
                        type="button"
                        onClick={handleDelete}
                    >
                        Delete Application    
                    </button>
                    <div>
                        <button 
                            className="modal-form__button"
                            type="button"
                            onClick={onCancel}
                        >
                            Cancel
                        </button>
                        <button
                            className="modal-form__button modal-form__button--primary"
                            type="submit"
                            onSubmit={handleConfirm}
                            onClick={handleConfirm}
                        >
                            Edit Application
                        </button>
                    </div>
                </div>
                <p ref={statusRef} className="modal-form__action-message"></p>
            </form>
        </dialog>
    );
}
