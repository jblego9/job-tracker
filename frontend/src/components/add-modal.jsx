import { useEffect, useRef, useState } from "react";
import * as Request from "../request.js";

export default function AddModal({ onAdd, onCancel }) {
    const [formData, setFormData] = useState({
        Id: 0,
        Company: "",
        Role: "",
        Location: "",
        Date: "",
        Status: "AWAITING_RESPONSE",
    });

    const dialogRef = useRef(null);
    const statusRef = useRef(null);

    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleConfirm = async (e) => {
        try {
            e.preventDefault();
            const response = await Request.postData(Request.apiUrl, formData);
            const { Id } = await response.json();
            const createdApp = { ...formData, Id };
            onAdd(createdApp);
        } catch (error) {
            statusRef.current.innerText = "Failed to create application";
            console.error(error);
        }
    };

    useEffect(() => {
        dialogRef.current.showModal();
    }, []);

    return (
        <dialog className="modal" ref={dialogRef} onClose={onCancel}>
            <form className="modal-form">
                <h1 className="no-bottom-margin">Add Application</h1>

                <h2>Job Details</h2>
                <label htmlFor="company">Company</label>
                <input
                    type="text"
                    id="company"
                    name="Company"
                    autoComplete="organization"
                    defaultValue=""
                    onChange={handleChange}
                    required
                />

                <label htmlFor="role">Role</label>
                <input
                    type="text"
                    id="role"
                    name="Role"
                    defaultValue=""
                    onChange={handleChange}
                    required
                />

                <label htmlFor="location">Location</label>
                <input
                    type="text"
                    id="location"
                    name="Location"
                    defaultValue=""
                    onChange={handleChange}
                    required
                />

                <h2>Application Details</h2>
                <label htmlFor="date">Date</label>
                <input
                    type="date"
                    id="date"
                    name="Date"
                    defaultValue={null}
                    onChange={handleChange}
                />

                <label htmlFor="status">Status</label>
                <select
                    id="status"
                    name="Status"
                    defaultValue="AWAITING_RESPONSE"
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
                    <div></div>
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
                            Add Application
                        </button>
                    </div>
                </div>
                <p ref={statusRef} className="modal-form__action-message"></p>
            </form>
        </dialog>
    );
}
