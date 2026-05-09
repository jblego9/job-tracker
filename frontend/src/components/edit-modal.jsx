import { useEffect, useRef } from "react";

export default function EditModal({ app, onClose }) {
    const dialogRef = useRef(null);

    useEffect(() => {
        dialogRef.current.showModal();
    }, []);

    return (
        <dialog className="modal" ref={dialogRef} onClose={onClose}>
            <form className="modal-form">
                <h1 className="no-bottom-margin">Edit Application</h1>
                <p className="edit-modal-application-id">Application ID: {app.Id}</p>

                <h2>Job Details</h2>
                <label htmlFor="Company">Company</label>
                <input type="text" id="Company" name="Company" defaultValue={app.Company} required />
                <label htmlFor="Role">Role</label>
                <input type="text" id="Role" name="Role" defaultValue={app.Role} required />
                <label htmlFor="Location">Location</label>
                <input type="text" id="Location" name="Location" defaultValue={app.Location} required />

                <h2>Application Details</h2>
                <label htmlFor="Date">Date</label>
                <input type="date" id="Date" name="Date" defaultValue={app.Date} />
                <label htmlFor="Status">Status</label>
                <select id="Status" name="Status" defaultValue={app.Status} required>
                    <option value="AWAITING_RESPONSE">Awaiting Response</option>
                    <option value="NO_RESPONSE">No Response</option>
                    <option value="AWAITING_INTERVIEW">Awaiting Interview</option>
                    <option value="REJECTED">Rejected</option>
                    <option value="ACCEPTED">Accepted</option>
                </select>

                <div className="modal-form__actions">
                    <p className="modal-form__action-message"></p>
                    <button 
                        className="modal-form__button"
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
                    >
                        Edit Application
                    </button>
                </div>
            </form>
        </dialog>
    );
}
