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
                <label for="Company">Company</label>
                <input type="text" name="Company" required />
                <label for="Role">Role</label>
                <input type="text" name="Role" required />
                <label for="Location">Location</label>
                <input type="text" name="Location" required />

                <h2>Application Details</h2>
                <label for="Date">Date</label>
                <input type="date" name="Date"/>
                <label for="Status">Status</label>
                <select name="Status" required>
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
                        onClick={() => dialogRef.current.close()}
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
