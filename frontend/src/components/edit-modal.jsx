import { useEffect, useRef } from "react";

export default function EditModal({ app, onClose }) {
    const dialogRef = useRef(null);

    useEffect(() => {
        dialogRef.current.showModal();
    }, []);

    return (
        <dialog className="edit-modal" ref={dialogRef} onClose={onClose}>
            <h2>Edit Application</h2>
            <p>ID: {app.Id}</p>
            <button onClick={() => dialogRef.current.close()}>Close</button>
        </dialog>
    );
}
