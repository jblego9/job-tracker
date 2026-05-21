import { useEffect, useRef } from "react";

export default function PageControls({ pageNumber, onPageNumberChange, pageSize, onPageSizeChange }) {
    const pageNumberRef = useRef(null);

    useEffect(() => {
        pageNumberRef.current.value = pageNumber;
    }, [pageNumber]);

    return (
        <div className="applications-control-bar__page-controls">
            <div>
                <label htmlFor="page-number">Page:</label>
                <input
                    ref={pageNumberRef}
                    id="page-number"
                    name="page-number"
                    type="number"
                    defaultValue={pageNumber}
                    min={1}
                    onChange={() => onPageNumberChange(Number(event.target.value))}
                />
                <button
                    onClick={() => onPageNumberChange(pageNumber - 1)}
                >
                    {"<"}
                </button>
                <button
                    onClick={() => onPageNumberChange(pageNumber + 1)}
                >
                    {">"}
                </button>
            </div>
            <div>
                <label htmlFor="page-size">Page Size:</label>
                <select
                    id="page-size"
                    name="page-size"
                    defaultValue={pageSize}
                    onChange={() => onPageSizeChange(Number(event.target.value))}
                >
                    <option value={25}>25</option>
                    <option value={50}>50</option>
                    <option value={100}>100</option>
                </select>
            </div>
        </div>
    );
}
