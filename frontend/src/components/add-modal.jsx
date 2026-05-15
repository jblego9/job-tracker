import { useEffect, useRef, useState } from "react";
import * as Request from "../request.js";

export default function AddModal({ applicationsUpdated, onClose }) {
    const [formData, setFormData] = useState({
        id: 0,
        company: "",
        role: "",
        location: "",
        date: "",
        status: "AWAITING_RESPONSE",
    });

    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const dialogRef = useRef(null);

    useEffect(() => {
        dialogRef.current.showModal();
    }, []);

    const statusRef = useRef(null);

    const handleConfirm = async (e) => {
        try {
            e.preventDefault();
            await Request.postData(Request.apiUrl, formData);
            onClose();
            applicationsUpdated(true);
        } catch (error) {
            statusRef.current.innerText = "Failed to add application";
            console.error(error);
        }
    };

    return (
        <dialog className="modal" ref={dialogRef} onClose={onClose}>
            <form className="modal-form">
                <h1 className="no-bottom-margin">Add Application</h1>

                <h2>Job Details</h2>
                <label htmlFor="company">Company</label>
                <input
                    type="text"
                    id="company"
                    name="company"
                    autoComplete="organization"
                    defaultValue=""
                    onChange={handleChange}
                    required
                />

                <label htmlFor="role">Role</label>
                <input
                    type="text"
                    id="role"
                    name="role"
                    defaultValue=""
                    onChange={handleChange}
                    required
                />

                <label htmlFor="location">Location</label>
                <input
                    type="text"
                    id="location"
                    name="location"
                    defaultValue=""
                    onChange={handleChange}
                    required
                />

                <h2>Application Details</h2>
                <label htmlFor="date">Date</label>
                <input
                    type="date"
                    id="date"
                    name="date"
                    defaultValue={null}
                    onChange={handleChange}
                />

                <label htmlFor="status">Status</label>
                <select
                    id="status"
                    name="status"
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
                            onClick={() => {
                                event.preventDefault();
                                dialogRef.current.close();
                            }
                            }
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
    )
}
